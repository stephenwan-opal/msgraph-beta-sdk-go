package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAutopilotDeploymentProfile 
type WindowsAutopilotDeploymentProfile struct {
    Entity
    // The list of assigned devices for the profile.
    assignedDevices []WindowsAutopilotDeviceIdentityable
    // The list of group assignments for the profile.
    assignments []WindowsAutopilotDeploymentProfileAssignmentable
    // Profile creation time
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Description of the profile
    description *string
    // The template used to name the AutoPilot Device. This can be a custom text and can also contain either the serial number of the device, or a randomly generated number. The total length of the text generated by the template can be no more than 15 characters.
    deviceNameTemplate *string
    // The deviceType property
    deviceType *WindowsAutopilotDeviceType
    // Name of the profile
    displayName *string
    // Enable Autopilot White Glove for the profile.
    enableWhiteGlove *bool
    // Enrollment status screen setting
    enrollmentStatusScreenSettings WindowsEnrollmentStatusScreenSettingsable
    // HardwareHash Extraction for the profile
    extractHardwareHash *bool
    // Language configured on the device
    language *string
    // Profile last modified time
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // AzureAD management app ID used during client device-based enrollment discovery
    managementServiceAppId *string
    // Out of box experience setting
    outOfBoxExperienceSettings OutOfBoxExperienceSettingsable
    // Scope tags for the profile.
    roleScopeTagIds []string
}
// NewWindowsAutopilotDeploymentProfile instantiates a new WindowsAutopilotDeploymentProfile and sets the default values.
func NewWindowsAutopilotDeploymentProfile()(*WindowsAutopilotDeploymentProfile) {
    m := &WindowsAutopilotDeploymentProfile{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.windowsAutopilotDeploymentProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsAutopilotDeploymentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsAutopilotDeploymentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.activeDirectoryWindowsAutopilotDeploymentProfile":
                        return NewActiveDirectoryWindowsAutopilotDeploymentProfile(), nil
                    case "#microsoft.graph.azureADWindowsAutopilotDeploymentProfile":
                        return NewAzureADWindowsAutopilotDeploymentProfile(), nil
                }
            }
        }
    }
    return NewWindowsAutopilotDeploymentProfile(), nil
}
// GetAssignedDevices gets the assignedDevices property value. The list of assigned devices for the profile.
func (m *WindowsAutopilotDeploymentProfile) GetAssignedDevices()([]WindowsAutopilotDeviceIdentityable) {
    if m == nil {
        return nil
    } else {
        return m.assignedDevices
    }
}
// GetAssignments gets the assignments property value. The list of group assignments for the profile.
func (m *WindowsAutopilotDeploymentProfile) GetAssignments()([]WindowsAutopilotDeploymentProfileAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Profile creation time
func (m *WindowsAutopilotDeploymentProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. Description of the profile
func (m *WindowsAutopilotDeploymentProfile) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDeviceNameTemplate gets the deviceNameTemplate property value. The template used to name the AutoPilot Device. This can be a custom text and can also contain either the serial number of the device, or a randomly generated number. The total length of the text generated by the template can be no more than 15 characters.
func (m *WindowsAutopilotDeploymentProfile) GetDeviceNameTemplate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceNameTemplate
    }
}
// GetDeviceType gets the deviceType property value. The deviceType property
func (m *WindowsAutopilotDeploymentProfile) GetDeviceType()(*WindowsAutopilotDeviceType) {
    if m == nil {
        return nil
    } else {
        return m.deviceType
    }
}
// GetDisplayName gets the displayName property value. Name of the profile
func (m *WindowsAutopilotDeploymentProfile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEnableWhiteGlove gets the enableWhiteGlove property value. Enable Autopilot White Glove for the profile.
func (m *WindowsAutopilotDeploymentProfile) GetEnableWhiteGlove()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableWhiteGlove
    }
}
// GetEnrollmentStatusScreenSettings gets the enrollmentStatusScreenSettings property value. Enrollment status screen setting
func (m *WindowsAutopilotDeploymentProfile) GetEnrollmentStatusScreenSettings()(WindowsEnrollmentStatusScreenSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentStatusScreenSettings
    }
}
// GetExtractHardwareHash gets the extractHardwareHash property value. HardwareHash Extraction for the profile
func (m *WindowsAutopilotDeploymentProfile) GetExtractHardwareHash()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.extractHardwareHash
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsAutopilotDeploymentProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsAutopilotDeviceIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsAutopilotDeviceIdentityable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsAutopilotDeviceIdentityable)
            }
            m.SetAssignedDevices(res)
        }
        return nil
    }
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsAutopilotDeploymentProfileAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsAutopilotDeploymentProfileAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsAutopilotDeploymentProfileAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["deviceNameTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceNameTemplate(val)
        }
        return nil
    }
    res["deviceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsAutopilotDeviceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceType(val.(*WindowsAutopilotDeviceType))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["enableWhiteGlove"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableWhiteGlove(val)
        }
        return nil
    }
    res["enrollmentStatusScreenSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsEnrollmentStatusScreenSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentStatusScreenSettings(val.(WindowsEnrollmentStatusScreenSettingsable))
        }
        return nil
    }
    res["extractHardwareHash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtractHardwareHash(val)
        }
        return nil
    }
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["managementServiceAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementServiceAppId(val)
        }
        return nil
    }
    res["outOfBoxExperienceSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOutOfBoxExperienceSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutOfBoxExperienceSettings(val.(OutOfBoxExperienceSettingsable))
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    return res
}
// GetLanguage gets the language property value. Language configured on the device
func (m *WindowsAutopilotDeploymentProfile) GetLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.language
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Profile last modified time
func (m *WindowsAutopilotDeploymentProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetManagementServiceAppId gets the managementServiceAppId property value. AzureAD management app ID used during client device-based enrollment discovery
func (m *WindowsAutopilotDeploymentProfile) GetManagementServiceAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementServiceAppId
    }
}
// GetOutOfBoxExperienceSettings gets the outOfBoxExperienceSettings property value. Out of box experience setting
func (m *WindowsAutopilotDeploymentProfile) GetOutOfBoxExperienceSettings()(OutOfBoxExperienceSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.outOfBoxExperienceSettings
    }
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. Scope tags for the profile.
func (m *WindowsAutopilotDeploymentProfile) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// Serialize serializes information the current object
func (m *WindowsAutopilotDeploymentProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignedDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignedDevices()))
        for i, v := range m.GetAssignedDevices() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignedDevices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        err = writer.WriteStringValue("deviceNameTemplate", m.GetDeviceNameTemplate())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceType() != nil {
        cast := (*m.GetDeviceType()).String()
        err = writer.WriteStringValue("deviceType", &cast)
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
        err = writer.WriteBoolValue("enableWhiteGlove", m.GetEnableWhiteGlove())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("enrollmentStatusScreenSettings", m.GetEnrollmentStatusScreenSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("extractHardwareHash", m.GetExtractHardwareHash())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("language", m.GetLanguage())
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
        err = writer.WriteStringValue("managementServiceAppId", m.GetManagementServiceAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("outOfBoxExperienceSettings", m.GetOutOfBoxExperienceSettings())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedDevices sets the assignedDevices property value. The list of assigned devices for the profile.
func (m *WindowsAutopilotDeploymentProfile) SetAssignedDevices(value []WindowsAutopilotDeviceIdentityable)() {
    if m != nil {
        m.assignedDevices = value
    }
}
// SetAssignments sets the assignments property value. The list of group assignments for the profile.
func (m *WindowsAutopilotDeploymentProfile) SetAssignments(value []WindowsAutopilotDeploymentProfileAssignmentable)() {
    if m != nil {
        m.assignments = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Profile creation time
func (m *WindowsAutopilotDeploymentProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. Description of the profile
func (m *WindowsAutopilotDeploymentProfile) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDeviceNameTemplate sets the deviceNameTemplate property value. The template used to name the AutoPilot Device. This can be a custom text and can also contain either the serial number of the device, or a randomly generated number. The total length of the text generated by the template can be no more than 15 characters.
func (m *WindowsAutopilotDeploymentProfile) SetDeviceNameTemplate(value *string)() {
    if m != nil {
        m.deviceNameTemplate = value
    }
}
// SetDeviceType sets the deviceType property value. The deviceType property
func (m *WindowsAutopilotDeploymentProfile) SetDeviceType(value *WindowsAutopilotDeviceType)() {
    if m != nil {
        m.deviceType = value
    }
}
// SetDisplayName sets the displayName property value. Name of the profile
func (m *WindowsAutopilotDeploymentProfile) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEnableWhiteGlove sets the enableWhiteGlove property value. Enable Autopilot White Glove for the profile.
func (m *WindowsAutopilotDeploymentProfile) SetEnableWhiteGlove(value *bool)() {
    if m != nil {
        m.enableWhiteGlove = value
    }
}
// SetEnrollmentStatusScreenSettings sets the enrollmentStatusScreenSettings property value. Enrollment status screen setting
func (m *WindowsAutopilotDeploymentProfile) SetEnrollmentStatusScreenSettings(value WindowsEnrollmentStatusScreenSettingsable)() {
    if m != nil {
        m.enrollmentStatusScreenSettings = value
    }
}
// SetExtractHardwareHash sets the extractHardwareHash property value. HardwareHash Extraction for the profile
func (m *WindowsAutopilotDeploymentProfile) SetExtractHardwareHash(value *bool)() {
    if m != nil {
        m.extractHardwareHash = value
    }
}
// SetLanguage sets the language property value. Language configured on the device
func (m *WindowsAutopilotDeploymentProfile) SetLanguage(value *string)() {
    if m != nil {
        m.language = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Profile last modified time
func (m *WindowsAutopilotDeploymentProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetManagementServiceAppId sets the managementServiceAppId property value. AzureAD management app ID used during client device-based enrollment discovery
func (m *WindowsAutopilotDeploymentProfile) SetManagementServiceAppId(value *string)() {
    if m != nil {
        m.managementServiceAppId = value
    }
}
// SetOutOfBoxExperienceSettings sets the outOfBoxExperienceSettings property value. Out of box experience setting
func (m *WindowsAutopilotDeploymentProfile) SetOutOfBoxExperienceSettings(value OutOfBoxExperienceSettingsable)() {
    if m != nil {
        m.outOfBoxExperienceSettings = value
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. Scope tags for the profile.
func (m *WindowsAutopilotDeploymentProfile) SetRoleScopeTagIds(value []string)() {
    if m != nil {
        m.roleScopeTagIds = value
    }
}
