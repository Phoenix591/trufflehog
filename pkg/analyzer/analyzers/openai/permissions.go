// Code generated by go generate; DO NOT EDIT.
package openai

import "errors"

type Permission int

const (
    Invalid Permission = iota
    ModelsRead Permission = iota
    ModelCapabilitiesWrite Permission = iota
    AssistantsRead Permission = iota
    AssistantsWrite Permission = iota
    ThreadsRead Permission = iota
    ThreadsWrite Permission = iota
    FineTuningRead Permission = iota
    FineTuningWrite Permission = iota
    FilesRead Permission = iota
    FilesWrite Permission = iota
)

var (
    PermissionStrings = map[Permission]string{
        ModelsRead: "models:read",
        ModelCapabilitiesWrite: "model_capabilities:write",
        AssistantsRead: "assistants:read",
        AssistantsWrite: "assistants:write",
        ThreadsRead: "threads:read",
        ThreadsWrite: "threads:write",
        FineTuningRead: "fine_tuning:read",
        FineTuningWrite: "fine_tuning:write",
        FilesRead: "files:read",
        FilesWrite: "files:write",
    }

    StringToPermission = map[string]Permission{
        "models:read": ModelsRead,
        "model_capabilities:write": ModelCapabilitiesWrite,
        "assistants:read": AssistantsRead,
        "assistants:write": AssistantsWrite,
        "threads:read": ThreadsRead,
        "threads:write": ThreadsWrite,
        "fine_tuning:read": FineTuningRead,
        "fine_tuning:write": FineTuningWrite,
        "files:read": FilesRead,
        "files:write": FilesWrite,
    }

    PermissionIDs = map[Permission]int{
        ModelsRead: 1,
        ModelCapabilitiesWrite: 2,
        AssistantsRead: 3,
        AssistantsWrite: 4,
        ThreadsRead: 5,
        ThreadsWrite: 6,
        FineTuningRead: 7,
        FineTuningWrite: 8,
        FilesRead: 9,
        FilesWrite: 10,
    }

    IdToPermission = map[int]Permission{
        1: ModelsRead,
        2: ModelCapabilitiesWrite,
        3: AssistantsRead,
        4: AssistantsWrite,
        5: ThreadsRead,
        6: ThreadsWrite,
        7: FineTuningRead,
        8: FineTuningWrite,
        9: FilesRead,
        10: FilesWrite,
    }
)

// ToString converts a Permission enum to its string representation
func (p Permission) ToString() (string, error) {
    if str, ok := PermissionStrings[p]; ok {
        return str, nil
    }
    return "", errors.New("invalid permission")
}

// ToID converts a Permission enum to its ID
func (p Permission) ToID() (int, error) {
    if id, ok := PermissionIDs[p]; ok {
        return id, nil
    }
    return 0, errors.New("invalid permission")
}

// PermissionFromString converts a string representation to its Permission enum
func PermissionFromString(s string) (Permission, error) {
    if p, ok := StringToPermission[s]; ok {
        return p, nil
    }
    return 0, errors.New("invalid permission string")
}

// PermissionFromID converts an ID to its Permission enum
func PermissionFromID(id int) (Permission, error) {
    if p, ok := IdToPermission[id]; ok {
        return p, nil
    }
    return 0, errors.New("invalid permission ID")
}
