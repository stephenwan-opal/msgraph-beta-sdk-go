package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the compliance singleton.
type DeviceEnrollmentConfigurationType int

const (
    UNKNOWN_DEVICEENROLLMENTCONFIGURATIONTYPE DeviceEnrollmentConfigurationType = iota
    LIMIT_DEVICEENROLLMENTCONFIGURATIONTYPE
    PLATFORMRESTRICTIONS_DEVICEENROLLMENTCONFIGURATIONTYPE
    WINDOWSHELLOFORBUSINESS_DEVICEENROLLMENTCONFIGURATIONTYPE
    DEFAULTLIMIT_DEVICEENROLLMENTCONFIGURATIONTYPE
    DEFAULTPLATFORMRESTRICTIONS_DEVICEENROLLMENTCONFIGURATIONTYPE
    DEFAULTWINDOWSHELLOFORBUSINESS_DEVICEENROLLMENTCONFIGURATIONTYPE
    DEFAULTWINDOWS10ENROLLMENTCOMPLETIONPAGECONFIGURATION_DEVICEENROLLMENTCONFIGURATIONTYPE
    WINDOWS10ENROLLMENTCOMPLETIONPAGECONFIGURATION_DEVICEENROLLMENTCONFIGURATIONTYPE
    DEVICECOMANAGEMENTAUTHORITYCONFIGURATION_DEVICEENROLLMENTCONFIGURATIONTYPE
    SINGLEPLATFORMRESTRICTION_DEVICEENROLLMENTCONFIGURATIONTYPE
    UNKNOWNFUTUREVALUE_DEVICEENROLLMENTCONFIGURATIONTYPE
    ENROLLMENTNOTIFICATIONSCONFIGURATION_DEVICEENROLLMENTCONFIGURATIONTYPE
)

func (i DeviceEnrollmentConfigurationType) String() string {
    return []string{"UNKNOWN", "LIMIT", "PLATFORMRESTRICTIONS", "WINDOWSHELLOFORBUSINESS", "DEFAULTLIMIT", "DEFAULTPLATFORMRESTRICTIONS", "DEFAULTWINDOWSHELLOFORBUSINESS", "DEFAULTWINDOWS10ENROLLMENTCOMPLETIONPAGECONFIGURATION", "WINDOWS10ENROLLMENTCOMPLETIONPAGECONFIGURATION", "DEVICECOMANAGEMENTAUTHORITYCONFIGURATION", "SINGLEPLATFORMRESTRICTION", "UNKNOWNFUTUREVALUE", "ENROLLMENTNOTIFICATIONSCONFIGURATION"}[i]
}
func ParseDeviceEnrollmentConfigurationType(v string) (interface{}, error) {
    result := UNKNOWN_DEVICEENROLLMENTCONFIGURATIONTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "LIMIT":
            result = LIMIT_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "PLATFORMRESTRICTIONS":
            result = PLATFORMRESTRICTIONS_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "WINDOWSHELLOFORBUSINESS":
            result = WINDOWSHELLOFORBUSINESS_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "DEFAULTLIMIT":
            result = DEFAULTLIMIT_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "DEFAULTPLATFORMRESTRICTIONS":
            result = DEFAULTPLATFORMRESTRICTIONS_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "DEFAULTWINDOWSHELLOFORBUSINESS":
            result = DEFAULTWINDOWSHELLOFORBUSINESS_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "DEFAULTWINDOWS10ENROLLMENTCOMPLETIONPAGECONFIGURATION":
            result = DEFAULTWINDOWS10ENROLLMENTCOMPLETIONPAGECONFIGURATION_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "WINDOWS10ENROLLMENTCOMPLETIONPAGECONFIGURATION":
            result = WINDOWS10ENROLLMENTCOMPLETIONPAGECONFIGURATION_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "DEVICECOMANAGEMENTAUTHORITYCONFIGURATION":
            result = DEVICECOMANAGEMENTAUTHORITYCONFIGURATION_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "SINGLEPLATFORMRESTRICTION":
            result = SINGLEPLATFORMRESTRICTION_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DEVICEENROLLMENTCONFIGURATIONTYPE
        case "ENROLLMENTNOTIFICATIONSCONFIGURATION":
            result = ENROLLMENTNOTIFICATIONSCONFIGURATION_DEVICEENROLLMENTCONFIGURATIONTYPE
        default:
            return 0, errors.New("Unknown DeviceEnrollmentConfigurationType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceEnrollmentConfigurationType(values []DeviceEnrollmentConfigurationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
