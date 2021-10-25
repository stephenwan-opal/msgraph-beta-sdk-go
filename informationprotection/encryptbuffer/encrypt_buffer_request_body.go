package encryptbuffer

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EncryptBufferRequestBody struct {
    additionalData map[string]interface{};
    buffer []byte;
    labelId *string;
}
func NewEncryptBufferRequestBody()(*EncryptBufferRequestBody) {
    m := &EncryptBufferRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *EncryptBufferRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *EncryptBufferRequestBody) GetBuffer()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.buffer
    }
}
func (m *EncryptBufferRequestBody) GetLabelId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.labelId
    }
}
func (m *EncryptBufferRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["buffer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetBuffer(val)
        return nil
    }
    res["labelId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLabelId(val)
        return nil
    }
    return res
}
func (m *EncryptBufferRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *EncryptBufferRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("buffer", m.GetBuffer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("labelId", m.GetLabelId())
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
func (m *EncryptBufferRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *EncryptBufferRequestBody) SetBuffer(value []byte)() {
    m.buffer = value
}
func (m *EncryptBufferRequestBody) SetLabelId(value *string)() {
    m.labelId = value
}