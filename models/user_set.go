package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserSet 
type UserSet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
    isBackup *bool;
}
// NewUserSet instantiates a new userSet and sets the default values.
func NewUserSet()(*UserSet) {
    m := &UserSet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUserSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserSet(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserSet) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserSet) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isBackup"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBackup(val)
        }
        return nil
    }
    return res
}
// GetIsBackup gets the isBackup property value. For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
func (m *UserSet) GetIsBackup()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBackup
    }
}
// Serialize serializes information the current object
func (m *UserSet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isBackup", m.GetIsBackup())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserSet) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsBackup sets the isBackup property value. For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
func (m *UserSet) SetIsBackup(value *bool)() {
    if m != nil {
        m.isBackup = value
    }
}