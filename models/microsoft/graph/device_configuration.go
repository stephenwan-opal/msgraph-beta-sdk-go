package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceConfiguration struct {
    Entity
    // The list of assignments for the device configuration profile.
    assignments []DeviceConfigurationAssignment;
    // DateTime the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Admin provided description of the Device Configuration.
    description *string;
    // The device mode applicability rule for this Policy.
    deviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode;
    // The OS edition applicability for this Policy.
    deviceManagementApplicabilityRuleOsEdition *DeviceManagementApplicabilityRuleOsEdition;
    // The OS version applicability rule for this Policy.
    deviceManagementApplicabilityRuleOsVersion *DeviceManagementApplicabilityRuleOsVersion;
    // Device Configuration Setting State Device Summary
    deviceSettingStateSummaries []SettingStateDeviceSummary;
    // Device configuration installation status by device.
    deviceStatuses []DeviceConfigurationDeviceStatus;
    // Device Configuration devices status overview
    deviceStatusOverview *DeviceConfigurationDeviceOverview;
    // Admin provided name of the device configuration.
    displayName *string;
    // The list of group assignments for the device configuration profile.
    groupAssignments []DeviceConfigurationGroupAssignment;
    // DateTime the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string;
    // Indicates whether or not the underlying Device Configuration supports the assignment of scope tags. Assigning to the ScopeTags property is not allowed when this value is false and entities will not be visible to scoped users. This occurs for Legacy policies created in Silverlight and can be resolved by deleting and recreating the policy in the Azure Portal. This property is read-only.
    supportsScopeTags *bool;
    // Device configuration installation status by user.
    userStatuses []DeviceConfigurationUserStatus;
    // Device Configuration users status overview
    userStatusOverview *DeviceConfigurationUserOverview;
    // Version of the device configuration.
    version *int32;
}
// Instantiates a new deviceConfiguration and sets the default values.
func NewDeviceConfiguration()(*DeviceConfiguration) {
    m := &DeviceConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignments property value. The list of assignments for the device configuration profile.
func (m *DeviceConfiguration) GetAssignments()([]DeviceConfigurationAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the createdDateTime property value. DateTime the object was created.
func (m *DeviceConfiguration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. Admin provided description of the Device Configuration.
func (m *DeviceConfiguration) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the deviceManagementApplicabilityRuleDeviceMode property value. The device mode applicability rule for this Policy.
func (m *DeviceConfiguration) GetDeviceManagementApplicabilityRuleDeviceMode()(*DeviceManagementApplicabilityRuleDeviceMode) {
    if m == nil {
        return nil
    } else {
        return m.deviceManagementApplicabilityRuleDeviceMode
    }
}
// Gets the deviceManagementApplicabilityRuleOsEdition property value. The OS edition applicability for this Policy.
func (m *DeviceConfiguration) GetDeviceManagementApplicabilityRuleOsEdition()(*DeviceManagementApplicabilityRuleOsEdition) {
    if m == nil {
        return nil
    } else {
        return m.deviceManagementApplicabilityRuleOsEdition
    }
}
// Gets the deviceManagementApplicabilityRuleOsVersion property value. The OS version applicability rule for this Policy.
func (m *DeviceConfiguration) GetDeviceManagementApplicabilityRuleOsVersion()(*DeviceManagementApplicabilityRuleOsVersion) {
    if m == nil {
        return nil
    } else {
        return m.deviceManagementApplicabilityRuleOsVersion
    }
}
// Gets the deviceSettingStateSummaries property value. Device Configuration Setting State Device Summary
func (m *DeviceConfiguration) GetDeviceSettingStateSummaries()([]SettingStateDeviceSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceSettingStateSummaries
    }
}
// Gets the deviceStatuses property value. Device configuration installation status by device.
func (m *DeviceConfiguration) GetDeviceStatuses()([]DeviceConfigurationDeviceStatus) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatuses
    }
}
// Gets the deviceStatusOverview property value. Device Configuration devices status overview
func (m *DeviceConfiguration) GetDeviceStatusOverview()(*DeviceConfigurationDeviceOverview) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatusOverview
    }
}
// Gets the displayName property value. Admin provided name of the device configuration.
func (m *DeviceConfiguration) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the groupAssignments property value. The list of group assignments for the device configuration profile.
func (m *DeviceConfiguration) GetGroupAssignments()([]DeviceConfigurationGroupAssignment) {
    if m == nil {
        return nil
    } else {
        return m.groupAssignments
    }
}
// Gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *DeviceConfiguration) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *DeviceConfiguration) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// Gets the supportsScopeTags property value. Indicates whether or not the underlying Device Configuration supports the assignment of scope tags. Assigning to the ScopeTags property is not allowed when this value is false and entities will not be visible to scoped users. This occurs for Legacy policies created in Silverlight and can be resolved by deleting and recreating the policy in the Azure Portal. This property is read-only.
func (m *DeviceConfiguration) GetSupportsScopeTags()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.supportsScopeTags
    }
}
// Gets the userStatuses property value. Device configuration installation status by user.
func (m *DeviceConfiguration) GetUserStatuses()([]DeviceConfigurationUserStatus) {
    if m == nil {
        return nil
    } else {
        return m.userStatuses
    }
}
// Gets the userStatusOverview property value. Device Configuration users status overview
func (m *DeviceConfiguration) GetUserStatusOverview()(*DeviceConfigurationUserOverview) {
    if m == nil {
        return nil
    } else {
        return m.userStatusOverview
    }
}
// Gets the version property value. Version of the device configuration.
func (m *DeviceConfiguration) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *DeviceConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfigurationAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceConfigurationAssignment))
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["deviceManagementApplicabilityRuleDeviceMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementApplicabilityRuleDeviceMode() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceManagementApplicabilityRuleDeviceMode(val.(*DeviceManagementApplicabilityRuleDeviceMode))
        }
        return nil
    }
    res["deviceManagementApplicabilityRuleOsEdition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementApplicabilityRuleOsEdition() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceManagementApplicabilityRuleOsEdition(val.(*DeviceManagementApplicabilityRuleOsEdition))
        }
        return nil
    }
    res["deviceManagementApplicabilityRuleOsVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementApplicabilityRuleOsVersion() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceManagementApplicabilityRuleOsVersion(val.(*DeviceManagementApplicabilityRuleOsVersion))
        }
        return nil
    }
    res["deviceSettingStateSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSettingStateDeviceSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingStateDeviceSummary, len(val))
            for i, v := range val {
                res[i] = *(v.(*SettingStateDeviceSummary))
            }
            m.SetDeviceSettingStateSummaries(res)
        }
        return nil
    }
    res["deviceStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfigurationDeviceStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationDeviceStatus, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceConfigurationDeviceStatus))
            }
            m.SetDeviceStatuses(res)
        }
        return nil
    }
    res["deviceStatusOverview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfigurationDeviceOverview() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceStatusOverview(val.(*DeviceConfigurationDeviceOverview))
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["groupAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfigurationGroupAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationGroupAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceConfigurationGroupAssignment))
            }
            m.SetGroupAssignments(res)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["supportsScopeTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportsScopeTags(val)
        }
        return nil
    }
    res["userStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfigurationUserStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationUserStatus, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceConfigurationUserStatus))
            }
            m.SetUserStatuses(res)
        }
        return nil
    }
    res["userStatusOverview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfigurationUserOverview() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserStatusOverview(val.(*DeviceConfigurationUserOverview))
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
func (m *DeviceConfiguration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceManagementApplicabilityRuleDeviceMode", m.GetDeviceManagementApplicabilityRuleDeviceMode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceManagementApplicabilityRuleOsEdition", m.GetDeviceManagementApplicabilityRuleOsEdition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceManagementApplicabilityRuleOsVersion", m.GetDeviceManagementApplicabilityRuleOsVersion())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceSettingStateSummaries()))
        for i, v := range m.GetDeviceSettingStateSummaries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceSettingStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceStatuses()))
        for i, v := range m.GetDeviceStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceStatusOverview", m.GetDeviceStatusOverview())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroupAssignments()))
        for i, v := range m.GetGroupAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("groupAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("supportsScopeTags", m.GetSupportsScopeTags())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserStatuses()))
        for i, v := range m.GetUserStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userStatusOverview", m.GetUserStatusOverview())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignments property value. The list of assignments for the device configuration profile.
// Parameters:
//  - value : Value to set for the assignments property.
func (m *DeviceConfiguration) SetAssignments(value []DeviceConfigurationAssignment)() {
    m.assignments = value
}
// Sets the createdDateTime property value. DateTime the object was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *DeviceConfiguration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. Admin provided description of the Device Configuration.
// Parameters:
//  - value : Value to set for the description property.
func (m *DeviceConfiguration) SetDescription(value *string)() {
    m.description = value
}
// Sets the deviceManagementApplicabilityRuleDeviceMode property value. The device mode applicability rule for this Policy.
// Parameters:
//  - value : Value to set for the deviceManagementApplicabilityRuleDeviceMode property.
func (m *DeviceConfiguration) SetDeviceManagementApplicabilityRuleDeviceMode(value *DeviceManagementApplicabilityRuleDeviceMode)() {
    m.deviceManagementApplicabilityRuleDeviceMode = value
}
// Sets the deviceManagementApplicabilityRuleOsEdition property value. The OS edition applicability for this Policy.
// Parameters:
//  - value : Value to set for the deviceManagementApplicabilityRuleOsEdition property.
func (m *DeviceConfiguration) SetDeviceManagementApplicabilityRuleOsEdition(value *DeviceManagementApplicabilityRuleOsEdition)() {
    m.deviceManagementApplicabilityRuleOsEdition = value
}
// Sets the deviceManagementApplicabilityRuleOsVersion property value. The OS version applicability rule for this Policy.
// Parameters:
//  - value : Value to set for the deviceManagementApplicabilityRuleOsVersion property.
func (m *DeviceConfiguration) SetDeviceManagementApplicabilityRuleOsVersion(value *DeviceManagementApplicabilityRuleOsVersion)() {
    m.deviceManagementApplicabilityRuleOsVersion = value
}
// Sets the deviceSettingStateSummaries property value. Device Configuration Setting State Device Summary
// Parameters:
//  - value : Value to set for the deviceSettingStateSummaries property.
func (m *DeviceConfiguration) SetDeviceSettingStateSummaries(value []SettingStateDeviceSummary)() {
    m.deviceSettingStateSummaries = value
}
// Sets the deviceStatuses property value. Device configuration installation status by device.
// Parameters:
//  - value : Value to set for the deviceStatuses property.
func (m *DeviceConfiguration) SetDeviceStatuses(value []DeviceConfigurationDeviceStatus)() {
    m.deviceStatuses = value
}
// Sets the deviceStatusOverview property value. Device Configuration devices status overview
// Parameters:
//  - value : Value to set for the deviceStatusOverview property.
func (m *DeviceConfiguration) SetDeviceStatusOverview(value *DeviceConfigurationDeviceOverview)() {
    m.deviceStatusOverview = value
}
// Sets the displayName property value. Admin provided name of the device configuration.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *DeviceConfiguration) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the groupAssignments property value. The list of group assignments for the device configuration profile.
// Parameters:
//  - value : Value to set for the groupAssignments property.
func (m *DeviceConfiguration) SetGroupAssignments(value []DeviceConfigurationGroupAssignment)() {
    m.groupAssignments = value
}
// Sets the lastModifiedDateTime property value. DateTime the object was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *DeviceConfiguration) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
// Parameters:
//  - value : Value to set for the roleScopeTagIds property.
func (m *DeviceConfiguration) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// Sets the supportsScopeTags property value. Indicates whether or not the underlying Device Configuration supports the assignment of scope tags. Assigning to the ScopeTags property is not allowed when this value is false and entities will not be visible to scoped users. This occurs for Legacy policies created in Silverlight and can be resolved by deleting and recreating the policy in the Azure Portal. This property is read-only.
// Parameters:
//  - value : Value to set for the supportsScopeTags property.
func (m *DeviceConfiguration) SetSupportsScopeTags(value *bool)() {
    m.supportsScopeTags = value
}
// Sets the userStatuses property value. Device configuration installation status by user.
// Parameters:
//  - value : Value to set for the userStatuses property.
func (m *DeviceConfiguration) SetUserStatuses(value []DeviceConfigurationUserStatus)() {
    m.userStatuses = value
}
// Sets the userStatusOverview property value. Device Configuration users status overview
// Parameters:
//  - value : Value to set for the userStatusOverview property.
func (m *DeviceConfiguration) SetUserStatusOverview(value *DeviceConfigurationUserOverview)() {
    m.userStatusOverview = value
}
// Sets the version property value. Version of the device configuration.
// Parameters:
//  - value : Value to set for the version property.
func (m *DeviceConfiguration) SetVersion(value *int32)() {
    m.version = value
}
