package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type LabelPolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    id *string;
    // 
    name *string;
}
// Instantiates a new labelPolicy and sets the default values.
func NewLabelPolicy()(*LabelPolicy) {
    m := &LabelPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LabelPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the id property value. 
func (m *LabelPolicy) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Gets the name property value. 
func (m *LabelPolicy) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// The deserialization information for the current model
func (m *LabelPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
func (m *LabelPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *LabelPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *LabelPolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the id property value. 
// Parameters:
//  - value : Value to set for the id property.
func (m *LabelPolicy) SetId(value *string)() {
    m.id = value
}
// Sets the name property value. 
// Parameters:
//  - value : Value to set for the name property.
func (m *LabelPolicy) SetName(value *string)() {
    m.name = value
}
