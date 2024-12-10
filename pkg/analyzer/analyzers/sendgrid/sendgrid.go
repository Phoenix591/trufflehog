//go:generate generate_permissions permissions.yaml permissions.go sendgrid

package sendgrid

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	sg "github.com/sendgrid/sendgrid-go"

	"github.com/trufflesecurity/trufflehog/v3/pkg/analyzer/analyzers"
	"github.com/trufflesecurity/trufflehog/v3/pkg/analyzer/config"
	"github.com/trufflesecurity/trufflehog/v3/pkg/context"
)

var _ analyzers.Analyzer = (*Analyzer)(nil)

type Analyzer struct {
	Cfg *config.Config
}

func (Analyzer) Type() analyzers.AnalyzerType { return analyzers.AnalyzerTypeSendgrid }

func (a Analyzer) Analyze(_ context.Context, credInfo map[string]string) (*analyzers.AnalyzerResult, error) {
	key, ok := credInfo["key"]
	if !ok {
		return nil, fmt.Errorf("missing key in credInfo")
	}
	info, err := AnalyzePermissions(a.Cfg, key)
	if err != nil {
		return nil, err
	}
	return secretInfoToAnalyzerResult(info), nil
}

func secretInfoToAnalyzerResult(info *SecretInfo) *analyzers.AnalyzerResult {
	if info == nil {
		return nil
	}

	var keyType string
	if slices.Contains(info.RawScopes, "user.email.read") {
		keyType = "full access"
	} else if slices.Contains(info.RawScopes, "billing.read") {
		keyType = "billing access"
	} else {
		keyType = "restricted access"
	}

	result := analyzers.AnalyzerResult{
		AnalyzerType: analyzers.AnalyzerTypeSendgrid,
		Metadata: map[string]any{
			"key_type":     keyType,
			"2fa_required": slices.Contains(info.RawScopes, "2fa_required"),
		},
		Bindings:           []analyzers.Binding{},
		UnboundedResources: []analyzers.Resource{},
	}

	for _, scope := range info.Scopes {
		resource := getCategoryResource(scope)

		if len(scope.Permissions) == 0 {
			result.UnboundedResources = append(result.UnboundedResources, *resource)
			continue
		}

		for _, permission := range scope.Permissions {
			result.Bindings = append(result.Bindings, analyzers.Binding{
				Resource: *resource,
				Permission: analyzers.Permission{
					Value: permission,
				},
			})
		}
	}

	return &result
}

func getCategoryResource(scope SendgridScope) *analyzers.Resource {
	categoryResource := &analyzers.Resource{
		Name:               scope.Category,
		FullyQualifiedName: scope.Category,
		Type:               "category",
		Metadata:           nil,
	}

	if scope.SubCategory != "" {
		return &analyzers.Resource{
			Name:               scope.SubCategory,
			FullyQualifiedName: fmt.Sprintf("%s/%s", scope.Category, scope.SubCategory),
			Type:               "category",
			Metadata:           nil,
			Parent:             categoryResource,
		}
	}

	return categoryResource
}

type ScopesJSON struct {
	Scopes []string `json:"scopes"`
}

type SecretInfo struct {
	RawScopes []string
	Scopes    []SendgridScope
}

func printPermissions(info *SecretInfo, show_all bool) {
	fmt.Print("\n\n")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	if show_all {
		t.AppendHeader(table.Row{"Scope", "Sub-Scope", "Access", "Permissions"})
	} else {
		t.AppendHeader(table.Row{"Scope", "Sub-Scope", "Access"})
	}
	// Print the scopes
	for _, s := range info.Scopes {
		writer := analyzers.GetWriterFromStatus(s.PermissionType)
		if show_all {
			t.AppendRow([]interface{}{writer(s.Category), writer(s.SubCategory), writer(s.PermissionType), writer(strings.Join(s.Permissions, "\n"))})
		} else if s.PermissionType != analyzers.NONE {
			t.AppendRow([]interface{}{writer(s.Category), writer(s.SubCategory), writer(s.PermissionType)})
		}
	}
	t.Render()
	fmt.Print("\n\n")
}

// getCategoryFromScope returns the category for a given scope.
// It will return the most specific category possible.
// For example, if the scope is "mail.send.read", it will return "Mail Send", not just "Mail"
// since it's searching "mail.send.read" -> "mail.send" -> "mail"
func getScopeIndex(categories []SendgridScope, scope string) int {
	splitScope := strings.Split(scope, ".")
	for i := len(splitScope); i > 0; i-- {
		searchScope := strings.Join(splitScope[:i], ".")
		for i, s := range categories {
			for _, prefix := range s.Prefixes {
				if strings.HasPrefix(searchScope, prefix) {
					return i
				}
			}
		}
	}
	return -1
}

func processPermissions(rawScopes []string) []SendgridScope {
	categoryPermissions := make([]SendgridScope, len(SCOPES))

	// copy all scope categories to the categoryPermissions slice
	copy(categoryPermissions, SCOPES)
	for _, scope := range rawScopes {
		// Skip these scopes since they are not useful for this analysis
		if scope == "2fa_required" || scope == "sender_verification_eligible" {
			continue
		}

		// must be part of generated permissions
		if _, ok := StringToPermission[scope]; !ok {
			continue
		}
		ind := getScopeIndex(categoryPermissions, scope)
		if ind == -1 {
			//color.Red("[!] Scope not found: %v", scope)
			continue
		}
		s := &categoryPermissions[ind]
		s.AddPermission(scope)
	}

	// Run tests to determine the permission type
	for i := range categoryPermissions {
		categoryPermissions[i].RunTests()
	}

	return categoryPermissions
}

func AnalyzeAndPrintPermissions(cfg *config.Config, key string) {
	info, err := AnalyzePermissions(cfg, key)
	if err != nil {
		color.Red("[!] Error: %v", err)
		return
	}

	color.Green("[!] Valid Sendgrid API Key\n\n")

	if slices.Contains(info.RawScopes, "user.email.read") {
		color.Green("[*] Sendgrid Key Type: Full Access Key")
	} else if slices.Contains(info.RawScopes, "billing.read") {
		color.Yellow("[*] Sendgrid Key Type: Billing Access Key")
	} else {
		color.Yellow("[*] Sendgrid Key Type: Restricted Access Key")
	}

	if slices.Contains(info.RawScopes, "2fa_required") {
		color.Yellow("[i] 2FA Required for this account")
	}

	printPermissions(info, cfg.ShowAll)
}

func AnalyzePermissions(cfg *config.Config, key string) (*SecretInfo, error) {
	// Setup custom HTTP client so we can log requests.
	sg.DefaultClient.HTTPClient = analyzers.NewAnalyzeClient(cfg)

	req := sg.GetRequest(key, "/v3/scopes", "https://api.sendgrid.com")
	req.Method = "GET"
	resp, err := sg.API(req)
	if resp.StatusCode == 401 || resp.StatusCode == 403 {
		return nil, fmt.Errorf("Invalid API Key")
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%v", resp.StatusCode)
	}
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response into a struct
	var jsonScopes ScopesJSON
	if err := json.Unmarshal([]byte(resp.Body), &jsonScopes); err != nil {
		return nil, err
	}

	// Now you can access the scopes
	rawScopes := jsonScopes.Scopes

	categoryScope := processPermissions(rawScopes)

	return &SecretInfo{
		RawScopes: rawScopes,
		Scopes:    categoryScope,
	}, nil
}