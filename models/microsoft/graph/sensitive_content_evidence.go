package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SensitiveContentEvidence struct {
    additionalData map[string]interface{};
    length *int32;
    match *string;
    offset *int32;
}
func NewSensitiveContentEvidence()(*SensitiveContentEvidence) {
    m := &SensitiveContentEvidence{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SensitiveContentEvidence) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SensitiveContentEvidence) GetLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.length
    }
}
func (m *SensitiveContentEvidence) GetMatch()(*string) {
    if m == nil {
        return nil
    } else {
        return m.match
    }
}
func (m *SensitiveContentEvidence) GetOffset()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.offset
    }
}
func (m *SensitiveContentEvidence) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["length"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetLength(val)
        return nil
    }
    res["match"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMatch(val)
        return nil
    }
    res["offset"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetOffset(val)
        return nil
    }
    return res
}
func (m *SensitiveContentEvidence) IsNil()(bool) {
    return m == nil
}
func (m *SensitiveContentEvidence) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("length", m.GetLength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("match", m.GetMatch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("offset", m.GetOffset())
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
func (m *SensitiveContentEvidence) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SensitiveContentEvidence) SetLength(value *int32)() {
    m.length = value
}
func (m *SensitiveContentEvidence) SetMatch(value *string)() {
    m.match = value
}
func (m *SensitiveContentEvidence) SetOffset(value *int32)() {
    m.offset = value
}