package getmailtips

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// GetMailTipsPostRequestBody provides operations to call the getMailTips method.
type GetMailTipsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The EmailAddresses property
    emailAddresses []string
    // The MailTipsOptions property
    mailTipsOptions *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailTipsType
}
// NewGetMailTipsPostRequestBody instantiates a new getMailTipsPostRequestBody and sets the default values.
func NewGetMailTipsPostRequestBody()(*GetMailTipsPostRequestBody) {
    m := &GetMailTipsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetMailTipsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetMailTipsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetMailTipsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetMailTipsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEmailAddresses gets the emailAddresses property value. The EmailAddresses property
func (m *GetMailTipsPostRequestBody) GetEmailAddresses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.emailAddresses
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetMailTipsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["emailAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEmailAddresses(res)
        }
        return nil
    }
    res["mailTipsOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseMailTipsType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailTipsOptions(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailTipsType))
        }
        return nil
    }
    return res
}
// GetMailTipsOptions gets the mailTipsOptions property value. The MailTipsOptions property
func (m *GetMailTipsPostRequestBody) GetMailTipsOptions()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailTipsType) {
    if m == nil {
        return nil
    } else {
        return m.mailTipsOptions
    }
}
// Serialize serializes information the current object
func (m *GetMailTipsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEmailAddresses() != nil {
        err := writer.WriteCollectionOfStringValues("emailAddresses", m.GetEmailAddresses())
        if err != nil {
            return err
        }
    }
    if m.GetMailTipsOptions() != nil {
        cast := (*m.GetMailTipsOptions()).String()
        err := writer.WriteStringValue("mailTipsOptions", &cast)
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
func (m *GetMailTipsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEmailAddresses sets the emailAddresses property value. The EmailAddresses property
func (m *GetMailTipsPostRequestBody) SetEmailAddresses(value []string)() {
    if m != nil {
        m.emailAddresses = value
    }
}
// SetMailTipsOptions sets the mailTipsOptions property value. The MailTipsOptions property
func (m *GetMailTipsPostRequestBody) SetMailTipsOptions(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MailTipsType)() {
    if m != nil {
        m.mailTipsOptions = value
    }
}
