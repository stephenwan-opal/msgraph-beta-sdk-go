package uploaddeptoken

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UploadDepTokenRequestBody struct {
    additionalData map[string]interface{};
    appleId *string;
    depToken *string;
}
func NewUploadDepTokenRequestBody()(*UploadDepTokenRequestBody) {
    m := &UploadDepTokenRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UploadDepTokenRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UploadDepTokenRequestBody) GetAppleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appleId
    }
}
func (m *UploadDepTokenRequestBody) GetDepToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.depToken
    }
}
func (m *UploadDepTokenRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["appleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppleId(val)
        return nil
    }
    res["depToken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDepToken(val)
        return nil
    }
    return res
}
func (m *UploadDepTokenRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *UploadDepTokenRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appleId", m.GetAppleId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("depToken", m.GetDepToken())
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
func (m *UploadDepTokenRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UploadDepTokenRequestBody) SetAppleId(value *string)() {
    m.appleId = value
}
func (m *UploadDepTokenRequestBody) SetDepToken(value *string)() {
    m.depToken = value
}