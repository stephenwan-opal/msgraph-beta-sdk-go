package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/ediscovery"
)

type Tag struct {
    Entity
    childSelectability *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.ChildSelectability;
    childTags []Tag;
    createdBy *IdentitySet;
    description *string;
    displayName *string;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    parent *Tag;
}
func NewTag()(*Tag) {
    m := &Tag{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Tag) GetChildSelectability()(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.ChildSelectability) {
    if m == nil {
        return nil
    } else {
        return m.childSelectability
    }
}
func (m *Tag) GetChildTags()([]Tag) {
    if m == nil {
        return nil
    } else {
        return m.childTags
    }
}
func (m *Tag) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
func (m *Tag) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *Tag) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *Tag) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *Tag) GetParent()(*Tag) {
    if m == nil {
        return nil
    } else {
        return m.parent
    }
}
func (m *Tag) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["childSelectability"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.ParseChildSelectability)
        if err != nil {
            return err
        }
        cast := val.(i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.ChildSelectability)
        m.SetChildSelectability(&cast)
        return nil
    }
    res["childTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTag() })
        if err != nil {
            return err
        }
        res := make([]Tag, len(val))
        for i, v := range val {
            res[i] = *(v.(*Tag))
        }
        m.SetChildTags(res)
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetCreatedBy(val.(*IdentitySet))
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["parent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTag() })
        if err != nil {
            return err
        }
        m.SetParent(val.(*Tag))
        return nil
    }
    return res
}
func (m *Tag) IsNil()(bool) {
    return m == nil
}
func (m *Tag) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChildSelectability() != nil {
        cast := m.GetChildSelectability().String()
        err = writer.WriteStringValue("childSelectability", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetChildTags()))
        for i, v := range m.GetChildTags() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("childTags", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parent", m.GetParent())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Tag) SetChildSelectability(value *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.ChildSelectability)() {
    m.childSelectability = value
}
func (m *Tag) SetChildTags(value []Tag)() {
    m.childTags = value
}
func (m *Tag) SetCreatedBy(value *IdentitySet)() {
    m.createdBy = value
}
func (m *Tag) SetDescription(value *string)() {
    m.description = value
}
func (m *Tag) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *Tag) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *Tag) SetParent(value *Tag)() {
    m.parent = value
}