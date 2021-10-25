package graph
import (
    "strings"
    "errors"
)
type AdministratorConfiguredDeviceComplianceState int

const (
    BASEDONDEVICECOMPLIANCEPOLICY_ADMINISTRATORCONFIGUREDDEVICECOMPLIANCESTATE AdministratorConfiguredDeviceComplianceState = iota
    NONCOMPLIANT_ADMINISTRATORCONFIGUREDDEVICECOMPLIANCESTATE
)

func (i AdministratorConfiguredDeviceComplianceState) String() string {
    return []string{"BASEDONDEVICECOMPLIANCEPOLICY", "NONCOMPLIANT"}[i]
}
func ParseAdministratorConfiguredDeviceComplianceState(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "BASEDONDEVICECOMPLIANCEPOLICY":
            return BASEDONDEVICECOMPLIANCEPOLICY_ADMINISTRATORCONFIGUREDDEVICECOMPLIANCESTATE, nil
        case "NONCOMPLIANT":
            return NONCOMPLIANT_ADMINISTRATORCONFIGUREDDEVICECOMPLIANCESTATE, nil
    }
    return 0, errors.New("Unknown AdministratorConfiguredDeviceComplianceState value: " + v)
}
func SerializeAdministratorConfiguredDeviceComplianceState(values []AdministratorConfiguredDeviceComplianceState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}