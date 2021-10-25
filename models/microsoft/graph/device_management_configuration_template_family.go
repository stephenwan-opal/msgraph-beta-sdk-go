package graph
import (
    "strings"
    "errors"
)
type DeviceManagementConfigurationTemplateFamily int

const (
    NONE_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY DeviceManagementConfigurationTemplateFamily = iota
    ENDPOINTSECURITYANTIVIRUS_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
    ENDPOINTSECURITYDISKENCRYPTION_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
    ENDPOINTSECURITYFIREWALL_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
    ENDPOINTSECURITYENDPOINTDETECTIONANDRESPONSE_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
    ENDPOINTSECURITYATTACKSURFACEREDUCTION_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
    ENDPOINTSECURITYACCOUNTPROTECTION_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
    ENDPOINTSECURITYAPPLICATIONCONTROL_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY
)

func (i DeviceManagementConfigurationTemplateFamily) String() string {
    return []string{"NONE", "ENDPOINTSECURITYANTIVIRUS", "ENDPOINTSECURITYDISKENCRYPTION", "ENDPOINTSECURITYFIREWALL", "ENDPOINTSECURITYENDPOINTDETECTIONANDRESPONSE", "ENDPOINTSECURITYATTACKSURFACEREDUCTION", "ENDPOINTSECURITYACCOUNTPROTECTION", "ENDPOINTSECURITYAPPLICATIONCONTROL"}[i]
}
func ParseDeviceManagementConfigurationTemplateFamily(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY, nil
        case "ENDPOINTSECURITYANTIVIRUS":
            return ENDPOINTSECURITYANTIVIRUS_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY, nil
        case "ENDPOINTSECURITYDISKENCRYPTION":
            return ENDPOINTSECURITYDISKENCRYPTION_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY, nil
        case "ENDPOINTSECURITYFIREWALL":
            return ENDPOINTSECURITYFIREWALL_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY, nil
        case "ENDPOINTSECURITYENDPOINTDETECTIONANDRESPONSE":
            return ENDPOINTSECURITYENDPOINTDETECTIONANDRESPONSE_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY, nil
        case "ENDPOINTSECURITYATTACKSURFACEREDUCTION":
            return ENDPOINTSECURITYATTACKSURFACEREDUCTION_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY, nil
        case "ENDPOINTSECURITYACCOUNTPROTECTION":
            return ENDPOINTSECURITYACCOUNTPROTECTION_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY, nil
        case "ENDPOINTSECURITYAPPLICATIONCONTROL":
            return ENDPOINTSECURITYAPPLICATIONCONTROL_DEVICEMANAGEMENTCONFIGURATIONTEMPLATEFAMILY, nil
    }
    return 0, errors.New("Unknown DeviceManagementConfigurationTemplateFamily value: " + v)
}
func SerializeDeviceManagementConfigurationTemplateFamily(values []DeviceManagementConfigurationTemplateFamily) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}