package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OmaSetting oMA Settings definition.
type OmaSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Description.
    description *string
    // Display Name.
    displayName *string
    // Indicates whether the value field is encrypted. This property is read-only.
    isEncrypted *bool
    // OMA.
    omaUri *string
    // ReferenceId for looking up secret for decryption. This property is read-only.
    secretReferenceValueId *string
}
// NewOmaSetting instantiates a new omaSetting and sets the default values.
func NewOmaSetting()(*OmaSetting) {
    m := &OmaSetting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOmaSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOmaSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOmaSetting(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OmaSetting) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDescription gets the description property value. Description.
func (m *OmaSetting) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Display Name.
func (m *OmaSetting) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OmaSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["isEncrypted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEncrypted(val)
        }
        return nil
    }
    res["omaUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOmaUri(val)
        }
        return nil
    }
    res["secretReferenceValueId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecretReferenceValueId(val)
        }
        return nil
    }
    return res
}
// GetIsEncrypted gets the isEncrypted property value. Indicates whether the value field is encrypted. This property is read-only.
func (m *OmaSetting) GetIsEncrypted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEncrypted
    }
}
// GetOmaUri gets the omaUri property value. OMA.
func (m *OmaSetting) GetOmaUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.omaUri
    }
}
// GetSecretReferenceValueId gets the secretReferenceValueId property value. ReferenceId for looking up secret for decryption. This property is read-only.
func (m *OmaSetting) GetSecretReferenceValueId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.secretReferenceValueId
    }
}
// Serialize serializes information the current object
func (m *OmaSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEncrypted", m.GetIsEncrypted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("omaUri", m.GetOmaUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secretReferenceValueId", m.GetSecretReferenceValueId())
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
func (m *OmaSetting) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDescription sets the description property value. Description.
func (m *OmaSetting) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Display Name.
func (m *OmaSetting) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsEncrypted sets the isEncrypted property value. Indicates whether the value field is encrypted. This property is read-only.
func (m *OmaSetting) SetIsEncrypted(value *bool)() {
    if m != nil {
        m.isEncrypted = value
    }
}
// SetOmaUri sets the omaUri property value. OMA.
func (m *OmaSetting) SetOmaUri(value *string)() {
    if m != nil {
        m.omaUri = value
    }
}
// SetSecretReferenceValueId sets the secretReferenceValueId property value. ReferenceId for looking up secret for decryption. This property is read-only.
func (m *OmaSetting) SetSecretReferenceValueId(value *string)() {
    if m != nil {
        m.secretReferenceValueId = value
    }
}