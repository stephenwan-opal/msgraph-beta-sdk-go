package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PropertyToEvaluate 
type PropertyToEvaluate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Provides the property name.
    propertyName *string;
    // Provides the property value.
    propertyValue *string;
}
// NewPropertyToEvaluate instantiates a new propertyToEvaluate and sets the default values.
func NewPropertyToEvaluate()(*PropertyToEvaluate) {
    m := &PropertyToEvaluate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePropertyToEvaluateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePropertyToEvaluateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPropertyToEvaluate(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PropertyToEvaluate) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PropertyToEvaluate) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["propertyName"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPropertyName(val)
        }
        return nil
    }
    res["propertyValue"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPropertyValue(val)
        }
        return nil
    }
    return res
}
// GetPropertyName gets the propertyName property value. Provides the property name.
func (m *PropertyToEvaluate) GetPropertyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.propertyName
    }
}
// GetPropertyValue gets the propertyValue property value. Provides the property value.
func (m *PropertyToEvaluate) GetPropertyValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.propertyValue
    }
}
// Serialize serializes information the current object
func (m *PropertyToEvaluate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("propertyName", m.GetPropertyName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("propertyValue", m.GetPropertyValue())
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
func (m *PropertyToEvaluate) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPropertyName sets the propertyName property value. Provides the property name.
func (m *PropertyToEvaluate) SetPropertyName(value *string)() {
    if m != nil {
        m.propertyName = value
    }
}
// SetPropertyValue sets the propertyValue property value. Provides the property value.
func (m *PropertyToEvaluate) SetPropertyValue(value *string)() {
    if m != nil {
        m.propertyValue = value
    }
}