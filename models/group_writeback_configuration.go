package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupWritebackConfiguration casts the previous resource to group.
type GroupWritebackConfiguration struct {
    WritebackConfiguration
    // The onPremisesGroupType property
    onPremisesGroupType *string
}
// NewGroupWritebackConfiguration instantiates a new groupWritebackConfiguration and sets the default values.
func NewGroupWritebackConfiguration()(*GroupWritebackConfiguration) {
    m := &GroupWritebackConfiguration{
        WritebackConfiguration: *NewWritebackConfiguration(),
    }
    return m
}
// CreateGroupWritebackConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupWritebackConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupWritebackConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupWritebackConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WritebackConfiguration.GetFieldDeserializers()
    res["onPremisesGroupType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesGroupType(val)
        }
        return nil
    }
    return res
}
// GetOnPremisesGroupType gets the onPremisesGroupType property value. The onPremisesGroupType property
func (m *GroupWritebackConfiguration) GetOnPremisesGroupType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesGroupType
    }
}
// Serialize serializes information the current object
func (m *GroupWritebackConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WritebackConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("onPremisesGroupType", m.GetOnPremisesGroupType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOnPremisesGroupType sets the onPremisesGroupType property value. The onPremisesGroupType property
func (m *GroupWritebackConfiguration) SetOnPremisesGroupType(value *string)() {
    if m != nil {
        m.onPremisesGroupType = value
    }
}
