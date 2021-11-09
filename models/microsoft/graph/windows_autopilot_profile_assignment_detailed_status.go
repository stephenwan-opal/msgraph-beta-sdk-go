package graph
import (
    "strings"
    "errors"
)
// 
type WindowsAutopilotProfileAssignmentDetailedStatus int

const (
    NONE_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS WindowsAutopilotProfileAssignmentDetailedStatus = iota
    HARDWAREREQUIREMENTSNOTMET_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
    SURFACEHUBPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
    HOLOLENSPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
    WINDOWSPCPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
    SURFACEHUB2SPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
    UNKNOWNFUTUREVALUE_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS
)

func (i WindowsAutopilotProfileAssignmentDetailedStatus) String() string {
    return []string{"NONE", "HARDWAREREQUIREMENTSNOTMET", "SURFACEHUBPROFILENOTSUPPORTED", "HOLOLENSPROFILENOTSUPPORTED", "WINDOWSPCPROFILENOTSUPPORTED", "SURFACEHUB2SPROFILENOTSUPPORTED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseWindowsAutopilotProfileAssignmentDetailedStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS, nil
        case "HARDWAREREQUIREMENTSNOTMET":
            return HARDWAREREQUIREMENTSNOTMET_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS, nil
        case "SURFACEHUBPROFILENOTSUPPORTED":
            return SURFACEHUBPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS, nil
        case "HOLOLENSPROFILENOTSUPPORTED":
            return HOLOLENSPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS, nil
        case "WINDOWSPCPROFILENOTSUPPORTED":
            return WINDOWSPCPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS, nil
        case "SURFACEHUB2SPROFILENOTSUPPORTED":
            return SURFACEHUB2SPROFILENOTSUPPORTED_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_WINDOWSAUTOPILOTPROFILEASSIGNMENTDETAILEDSTATUS, nil
    }
    return 0, errors.New("Unknown WindowsAutopilotProfileAssignmentDetailedStatus value: " + v)
}
func SerializeWindowsAutopilotProfileAssignmentDetailedStatus(values []WindowsAutopilotProfileAssignmentDetailedStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
