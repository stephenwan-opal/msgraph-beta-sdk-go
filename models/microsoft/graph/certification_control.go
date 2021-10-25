package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CertificationControl struct {
    additionalData map[string]interface{};
    name *string;
    url *string;
}
func NewCertificationControl()(*CertificationControl) {
    m := &CertificationControl{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *CertificationControl) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *CertificationControl) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *CertificationControl) GetUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.url
    }
}
func (m *CertificationControl) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["url"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUrl(val)
        return nil
    }
    return res
}
func (m *CertificationControl) IsNil()(bool) {
    return m == nil
}
func (m *CertificationControl) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *CertificationControl) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *CertificationControl) SetName(value *string)() {
    m.name = value
}
func (m *CertificationControl) SetUrl(value *string)() {
    m.url = value
}