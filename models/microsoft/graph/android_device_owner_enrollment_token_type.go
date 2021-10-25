package graph
import (
    "strings"
    "errors"
)
type AndroidDeviceOwnerEnrollmentTokenType int

const (
    DEFAULT_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE AndroidDeviceOwnerEnrollmentTokenType = iota
    CORPORATEOWNEDDEDICATEDDEVICEWITHAZUREADSHAREDMODE_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE
)

func (i AndroidDeviceOwnerEnrollmentTokenType) String() string {
    return []string{"DEFAULT", "CORPORATEOWNEDDEDICATEDDEVICEWITHAZUREADSHAREDMODE"}[i]
}
func ParseAndroidDeviceOwnerEnrollmentTokenType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "DEFAULT":
            return DEFAULT_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE, nil
        case "CORPORATEOWNEDDEDICATEDDEVICEWITHAZUREADSHAREDMODE":
            return CORPORATEOWNEDDEDICATEDDEVICEWITHAZUREADSHAREDMODE_ANDROIDDEVICEOWNERENROLLMENTTOKENTYPE, nil
    }
    return 0, errors.New("Unknown AndroidDeviceOwnerEnrollmentTokenType value: " + v)
}
func SerializeAndroidDeviceOwnerEnrollmentTokenType(values []AndroidDeviceOwnerEnrollmentTokenType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}