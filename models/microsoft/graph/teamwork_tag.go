package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TeamworkTag struct {
    Entity
    // Tag description as it will appear to the user in Microsoft Teams.
    description *string;
    // Tag name as it will appear to the user in Microsoft Teams.
    displayName *string;
    // The number of users assigned to the tag.
    memberCount *int32;
    // Users assigned to the tag.
    members []TeamworkTagMember;
    // The type of tag. Default is standard.
    tagType *TeamworkTagType;
    // ID of the team in which the tag is defined.
    teamId *string;
}
// Instantiates a new teamworkTag and sets the default values.
func NewTeamworkTag()(*TeamworkTag) {
    m := &TeamworkTag{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the description property value. Tag description as it will appear to the user in Microsoft Teams.
func (m *TeamworkTag) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Tag name as it will appear to the user in Microsoft Teams.
func (m *TeamworkTag) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the memberCount property value. The number of users assigned to the tag.
func (m *TeamworkTag) GetMemberCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.memberCount
    }
}
// Gets the members property value. Users assigned to the tag.
func (m *TeamworkTag) GetMembers()([]TeamworkTagMember) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// Gets the tagType property value. The type of tag. Default is standard.
func (m *TeamworkTag) GetTagType()(*TeamworkTagType) {
    if m == nil {
        return nil
    } else {
        return m.tagType
    }
}
// Gets the teamId property value. ID of the team in which the tag is defined.
func (m *TeamworkTag) GetTeamId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.teamId
    }
}
// The deserialization information for the current model
func (m *TeamworkTag) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["memberCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberCount(val)
        }
        return nil
    }
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkTagMember() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamworkTagMember, len(val))
            for i, v := range val {
                res[i] = *(v.(*TeamworkTagMember))
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["tagType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkTagType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(TeamworkTagType)
            m.SetTagType(&cast)
        }
        return nil
    }
    res["teamId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamId(val)
        }
        return nil
    }
    return res
}
func (m *TeamworkTag) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TeamworkTag) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteInt32Value("memberCount", m.GetMemberCount())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTagType() != nil {
        cast := m.GetTagType().String()
        err = writer.WriteStringValue("tagType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamId", m.GetTeamId())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the description property value. Tag description as it will appear to the user in Microsoft Teams.
// Parameters:
//  - value : Value to set for the description property.
func (m *TeamworkTag) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Tag name as it will appear to the user in Microsoft Teams.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *TeamworkTag) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the memberCount property value. The number of users assigned to the tag.
// Parameters:
//  - value : Value to set for the memberCount property.
func (m *TeamworkTag) SetMemberCount(value *int32)() {
    m.memberCount = value
}
// Sets the members property value. Users assigned to the tag.
// Parameters:
//  - value : Value to set for the members property.
func (m *TeamworkTag) SetMembers(value []TeamworkTagMember)() {
    m.members = value
}
// Sets the tagType property value. The type of tag. Default is standard.
// Parameters:
//  - value : Value to set for the tagType property.
func (m *TeamworkTag) SetTagType(value *TeamworkTagType)() {
    m.tagType = value
}
// Sets the teamId property value. ID of the team in which the tag is defined.
// Parameters:
//  - value : Value to set for the teamId property.
func (m *TeamworkTag) SetTeamId(value *string)() {
    m.teamId = value
}
