package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppAssignment a class containing the properties used for Group Assignment of a Mobile App.
type MobileAppAssignment struct {
    Entity
    // The install intent defined by the admin. Possible values are: available, required, uninstall, availableWithoutEnrollment.
    intent *InstallIntent
    // The settings for target assignment defined by the admin.
    settings MobileAppAssignmentSettingsable
    // The resource type which is the source for the assignment. Possible values are: direct, policySets.
    source *DeviceAndAppManagementAssignmentSource
    // The identifier of the source of the assignment.
    sourceId *string
    // The target group assignment defined by the admin.
    target DeviceAndAppManagementAssignmentTargetable
}
// NewMobileAppAssignment instantiates a new mobileAppAssignment and sets the default values.
func NewMobileAppAssignment()(*MobileAppAssignment) {
    m := &MobileAppAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileAppAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["intent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseInstallIntent)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntent(val.(*InstallIntent))
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMobileAppAssignmentSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(MobileAppAssignmentSettingsable))
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAndAppManagementAssignmentSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*DeviceAndAppManagementAssignmentSource))
        }
        return nil
    }
    res["sourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceId(val)
        }
        return nil
    }
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceAndAppManagementAssignmentTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(DeviceAndAppManagementAssignmentTargetable))
        }
        return nil
    }
    return res
}
// GetIntent gets the intent property value. The install intent defined by the admin. Possible values are: available, required, uninstall, availableWithoutEnrollment.
func (m *MobileAppAssignment) GetIntent()(*InstallIntent) {
    if m == nil {
        return nil
    } else {
        return m.intent
    }
}
// GetSettings gets the settings property value. The settings for target assignment defined by the admin.
func (m *MobileAppAssignment) GetSettings()(MobileAppAssignmentSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// GetSource gets the source property value. The resource type which is the source for the assignment. Possible values are: direct, policySets.
func (m *MobileAppAssignment) GetSource()(*DeviceAndAppManagementAssignmentSource) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
// GetSourceId gets the sourceId property value. The identifier of the source of the assignment.
func (m *MobileAppAssignment) GetSourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceId
    }
}
// GetTarget gets the target property value. The target group assignment defined by the admin.
func (m *MobileAppAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// Serialize serializes information the current object
func (m *MobileAppAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetIntent() != nil {
        cast := (*m.GetIntent()).String()
        err = writer.WriteStringValue("intent", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    if m.GetSource() != nil {
        cast := (*m.GetSource()).String()
        err = writer.WriteStringValue("source", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sourceId", m.GetSourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIntent sets the intent property value. The install intent defined by the admin. Possible values are: available, required, uninstall, availableWithoutEnrollment.
func (m *MobileAppAssignment) SetIntent(value *InstallIntent)() {
    if m != nil {
        m.intent = value
    }
}
// SetSettings sets the settings property value. The settings for target assignment defined by the admin.
func (m *MobileAppAssignment) SetSettings(value MobileAppAssignmentSettingsable)() {
    if m != nil {
        m.settings = value
    }
}
// SetSource sets the source property value. The resource type which is the source for the assignment. Possible values are: direct, policySets.
func (m *MobileAppAssignment) SetSource(value *DeviceAndAppManagementAssignmentSource)() {
    if m != nil {
        m.source = value
    }
}
// SetSourceId sets the sourceId property value. The identifier of the source of the assignment.
func (m *MobileAppAssignment) SetSourceId(value *string)() {
    if m != nil {
        m.sourceId = value
    }
}
// SetTarget sets the target property value. The target group assignment defined by the admin.
func (m *MobileAppAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    if m != nil {
        m.target = value
    }
}
