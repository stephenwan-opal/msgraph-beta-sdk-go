package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MobileAppRelationship struct {
    Entity
    // The target mobile app's display name.
    targetDisplayName *string;
    // The target mobile app's display version.
    targetDisplayVersion *string;
    // The target mobile app's app id.
    targetId *string;
    // The target mobile app's publisher.
    targetPublisher *string;
    // The type of relationship indicating whether the target is a parent or child. Possible values are: child, parent.
    targetType *MobileAppRelationshipType;
}
// Instantiates a new mobileAppRelationship and sets the default values.
func NewMobileAppRelationship()(*MobileAppRelationship) {
    m := &MobileAppRelationship{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the targetDisplayName property value. The target mobile app's display name.
func (m *MobileAppRelationship) GetTargetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetDisplayName
    }
}
// Gets the targetDisplayVersion property value. The target mobile app's display version.
func (m *MobileAppRelationship) GetTargetDisplayVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetDisplayVersion
    }
}
// Gets the targetId property value. The target mobile app's app id.
func (m *MobileAppRelationship) GetTargetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetId
    }
}
// Gets the targetPublisher property value. The target mobile app's publisher.
func (m *MobileAppRelationship) GetTargetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetPublisher
    }
}
// Gets the targetType property value. The type of relationship indicating whether the target is a parent or child. Possible values are: child, parent.
func (m *MobileAppRelationship) GetTargetType()(*MobileAppRelationshipType) {
    if m == nil {
        return nil
    } else {
        return m.targetType
    }
}
// The deserialization information for the current model
func (m *MobileAppRelationship) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["targetDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetDisplayName(val)
        }
        return nil
    }
    res["targetDisplayVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetDisplayVersion(val)
        }
        return nil
    }
    res["targetId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetId(val)
        }
        return nil
    }
    res["targetPublisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetPublisher(val)
        }
        return nil
    }
    res["targetType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileAppRelationshipType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(MobileAppRelationshipType)
            m.SetTargetType(&cast)
        }
        return nil
    }
    return res
}
func (m *MobileAppRelationship) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MobileAppRelationship) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("targetDisplayName", m.GetTargetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetDisplayVersion", m.GetTargetDisplayVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetId", m.GetTargetId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetPublisher", m.GetTargetPublisher())
        if err != nil {
            return err
        }
    }
    if m.GetTargetType() != nil {
        cast := m.GetTargetType().String()
        err = writer.WriteStringValue("targetType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the targetDisplayName property value. The target mobile app's display name.
// Parameters:
//  - value : Value to set for the targetDisplayName property.
func (m *MobileAppRelationship) SetTargetDisplayName(value *string)() {
    m.targetDisplayName = value
}
// Sets the targetDisplayVersion property value. The target mobile app's display version.
// Parameters:
//  - value : Value to set for the targetDisplayVersion property.
func (m *MobileAppRelationship) SetTargetDisplayVersion(value *string)() {
    m.targetDisplayVersion = value
}
// Sets the targetId property value. The target mobile app's app id.
// Parameters:
//  - value : Value to set for the targetId property.
func (m *MobileAppRelationship) SetTargetId(value *string)() {
    m.targetId = value
}
// Sets the targetPublisher property value. The target mobile app's publisher.
// Parameters:
//  - value : Value to set for the targetPublisher property.
func (m *MobileAppRelationship) SetTargetPublisher(value *string)() {
    m.targetPublisher = value
}
// Sets the targetType property value. The type of relationship indicating whether the target is a parent or child. Possible values are: child, parent.
// Parameters:
//  - value : Value to set for the targetType property.
func (m *MobileAppRelationship) SetTargetType(value *MobileAppRelationshipType)() {
    m.targetType = value
}
