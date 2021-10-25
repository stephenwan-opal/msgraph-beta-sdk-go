package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Picture struct {
    Entity
    content []byte;
    contentType *string;
    height *int32;
    width *int32;
}
func NewPicture()(*Picture) {
    m := &Picture{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Picture) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
func (m *Picture) GetContentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
func (m *Picture) GetHeight()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.height
    }
}
func (m *Picture) GetWidth()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.width
    }
}
func (m *Picture) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetContent(val)
        return nil
    }
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContentType(val)
        return nil
    }
    res["height"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetHeight(val)
        return nil
    }
    res["width"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetWidth(val)
        return nil
    }
    return res
}
func (m *Picture) IsNil()(bool) {
    return m == nil
}
func (m *Picture) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("height", m.GetHeight())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("width", m.GetWidth())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Picture) SetContent(value []byte)() {
    m.content = value
}
func (m *Picture) SetContentType(value *string)() {
    m.contentType = value
}
func (m *Picture) SetHeight(value *int32)() {
    m.height = value
}
func (m *Picture) SetWidth(value *int32)() {
    m.width = value
}