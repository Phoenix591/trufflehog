// Code generated by go generate; DO NOT EDIT.
package twilio

import "errors"

type Permission int

const (
    NoAccess Permission = iota
    AccountManagementRead Permission = iota
    AccountManagementWrite Permission = iota
    SubaccountConfigurationRead Permission = iota
    SubaccountConfigurationWrite Permission = iota
    KeyManagementRead Permission = iota
    KeyManagementWrite Permission = iota
    ServiceVerificationRead Permission = iota
    ServiceVerificationWrite Permission = iota
    SmsRead Permission = iota
    SmsWrite Permission = iota
    VoiceRead Permission = iota
    VoiceWrite Permission = iota
    MessagingRead Permission = iota
    MessagingWrite Permission = iota
    CallManagementRead Permission = iota
    CallManagementWrite Permission = iota
)

var (
    PermissionStrings = map[Permission]string{
        AccountManagementRead: "account_management:read",
        AccountManagementWrite: "account_management:write",
        SubaccountConfigurationRead: "subaccount_configuration:read",
        SubaccountConfigurationWrite: "subaccount_configuration:write",
        KeyManagementRead: "key_management:read",
        KeyManagementWrite: "key_management:write",
        ServiceVerificationRead: "service_verification:read",
        ServiceVerificationWrite: "service_verification:write",
        SmsRead: "sms:read",
        SmsWrite: "sms:write",
        VoiceRead: "voice:read",
        VoiceWrite: "voice:write",
        MessagingRead: "messaging:read",
        MessagingWrite: "messaging:write",
        CallManagementRead: "call_management:read",
        CallManagementWrite: "call_management:write",
    }

    StringToPermission = map[string]Permission{
        "account_management:read": AccountManagementRead,
        "account_management:write": AccountManagementWrite,
        "subaccount_configuration:read": SubaccountConfigurationRead,
        "subaccount_configuration:write": SubaccountConfigurationWrite,
        "key_management:read": KeyManagementRead,
        "key_management:write": KeyManagementWrite,
        "service_verification:read": ServiceVerificationRead,
        "service_verification:write": ServiceVerificationWrite,
        "sms:read": SmsRead,
        "sms:write": SmsWrite,
        "voice:read": VoiceRead,
        "voice:write": VoiceWrite,
        "messaging:read": MessagingRead,
        "messaging:write": MessagingWrite,
        "call_management:read": CallManagementRead,
        "call_management:write": CallManagementWrite,
    }

    PermissionIDs = map[Permission]int{
        AccountManagementRead: 0,
        AccountManagementWrite: 1,
        SubaccountConfigurationRead: 2,
        SubaccountConfigurationWrite: 3,
        KeyManagementRead: 4,
        KeyManagementWrite: 5,
        ServiceVerificationRead: 6,
        ServiceVerificationWrite: 7,
        SmsRead: 8,
        SmsWrite: 9,
        VoiceRead: 10,
        VoiceWrite: 11,
        MessagingRead: 12,
        MessagingWrite: 13,
        CallManagementRead: 14,
        CallManagementWrite: 15,
    }

    IdToPermission = map[int]Permission{
        0: AccountManagementRead,
        1: AccountManagementWrite,
        2: SubaccountConfigurationRead,
        3: SubaccountConfigurationWrite,
        4: KeyManagementRead,
        5: KeyManagementWrite,
        6: ServiceVerificationRead,
        7: ServiceVerificationWrite,
        8: SmsRead,
        9: SmsWrite,
        10: VoiceRead,
        11: VoiceWrite,
        12: MessagingRead,
        13: MessagingWrite,
        14: CallManagementRead,
        15: CallManagementWrite,
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
