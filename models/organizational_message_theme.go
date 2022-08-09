package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type OrganizationalMessageTheme int

const (
    // Indicates the Update theme
    UPDATE_ORGANIZATIONALMESSAGETHEME OrganizationalMessageTheme = iota
    // Indicates the Training theme
    TRAINING_ORGANIZATIONALMESSAGETHEME
    // Indicates the Welcome to Windows theme
    WELCOMETOWINDOWS_ORGANIZATIONALMESSAGETHEME
)

func (i OrganizationalMessageTheme) String() string {
    return []string{"update", "training", "welcomeToWindows"}[i]
}
func ParseOrganizationalMessageTheme(v string) (interface{}, error) {
    result := UPDATE_ORGANIZATIONALMESSAGETHEME
    switch v {
        case "update":
            result = UPDATE_ORGANIZATIONALMESSAGETHEME
        case "training":
            result = TRAINING_ORGANIZATIONALMESSAGETHEME
        case "welcomeToWindows":
            result = WELCOMETOWINDOWS_ORGANIZATIONALMESSAGETHEME
        default:
            return 0, errors.New("Unknown OrganizationalMessageTheme value: " + v)
    }
    return &result, nil
}
func SerializeOrganizationalMessageTheme(values []OrganizationalMessageTheme) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
