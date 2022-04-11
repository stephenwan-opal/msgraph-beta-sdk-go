package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the commsApplication singleton.
type AutoAdmittedUsersType int

const (
    EVERYONEINCOMPANY_AUTOADMITTEDUSERSTYPE AutoAdmittedUsersType = iota
    EVERYONE_AUTOADMITTEDUSERSTYPE
)

func (i AutoAdmittedUsersType) String() string {
    return []string{"EVERYONEINCOMPANY", "EVERYONE"}[i]
}
func ParseAutoAdmittedUsersType(v string) (interface{}, error) {
    result := EVERYONEINCOMPANY_AUTOADMITTEDUSERSTYPE
    switch strings.ToUpper(v) {
        case "EVERYONEINCOMPANY":
            result = EVERYONEINCOMPANY_AUTOADMITTEDUSERSTYPE
        case "EVERYONE":
            result = EVERYONE_AUTOADMITTEDUSERSTYPE
        default:
            return 0, errors.New("Unknown AutoAdmittedUsersType value: " + v)
    }
    return &result, nil
}
func SerializeAutoAdmittedUsersType(values []AutoAdmittedUsersType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}