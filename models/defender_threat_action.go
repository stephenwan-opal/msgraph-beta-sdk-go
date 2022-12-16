package models
import (
    "errors"
)
// Provides operations to manage the collection of user entities.
type DefenderThreatAction int

const (
    // Apply action based on the update definition.
    DEVICEDEFAULT_DEFENDERTHREATACTION DefenderThreatAction = iota
    // Clean the detected threat.
    CLEAN_DEFENDERTHREATACTION
    // Quarantine the detected threat.
    QUARANTINE_DEFENDERTHREATACTION
    // Remove the detected threat.
    REMOVE_DEFENDERTHREATACTION
    // Allow the detected threat.
    ALLOW_DEFENDERTHREATACTION
    // Allow the user to determine the action to take with the detected threat.
    USERDEFINED_DEFENDERTHREATACTION
    // Block the detected threat.
    BLOCK_DEFENDERTHREATACTION
)

func (i DefenderThreatAction) String() string {
    return []string{"deviceDefault", "clean", "quarantine", "remove", "allow", "userDefined", "block"}[i]
}
func ParseDefenderThreatAction(v string) (interface{}, error) {
    result := DEVICEDEFAULT_DEFENDERTHREATACTION
    switch v {
        case "deviceDefault":
            result = DEVICEDEFAULT_DEFENDERTHREATACTION
        case "clean":
            result = CLEAN_DEFENDERTHREATACTION
        case "quarantine":
            result = QUARANTINE_DEFENDERTHREATACTION
        case "remove":
            result = REMOVE_DEFENDERTHREATACTION
        case "allow":
            result = ALLOW_DEFENDERTHREATACTION
        case "userDefined":
            result = USERDEFINED_DEFENDERTHREATACTION
        case "block":
            result = BLOCK_DEFENDERTHREATACTION
        default:
            return 0, errors.New("Unknown DefenderThreatAction value: " + v)
    }
    return &result, nil
}
func SerializeDefenderThreatAction(values []DefenderThreatAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
