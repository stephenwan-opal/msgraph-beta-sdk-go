package windowsupdates

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeploymentState provides operations to manage the admin singleton.
type DeploymentState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies the reasons the deployment has its state value. Read-only.
    reasons []DeploymentStateReasonable;
    // Specifies the requested state of the deployment. Supports a subset of the values for requestedDeploymentStateValue. Possible values are: none, paused, unknownFutureValue.
    requestedValue *RequestedDeploymentStateValue;
    // Specifies the state of the deployment. Supports a subset of the values for deploymentStateValue. Possible values are: scheduled, offering, paused, unknownFutureValue. Read-only.
    value *DeploymentStateValue;
}
// NewDeploymentState instantiates a new deploymentState and sets the default values.
func NewDeploymentState()(*DeploymentState) {
    m := &DeploymentState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeploymentStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeploymentStateFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeploymentState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeploymentState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeploymentState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["reasons"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeploymentStateReasonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeploymentStateReasonable, len(val))
            for i, v := range val {
                res[i] = v.(DeploymentStateReasonable)
            }
            m.SetReasons(res)
        }
        return nil
    }
    res["requestedValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequestedDeploymentStateValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedValue(val.(*RequestedDeploymentStateValue))
        }
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeploymentStateValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val.(*DeploymentStateValue))
        }
        return nil
    }
    return res
}
// GetReasons gets the reasons property value. Specifies the reasons the deployment has its state value. Read-only.
func (m *DeploymentState) GetReasons()([]DeploymentStateReasonable) {
    if m == nil {
        return nil
    } else {
        return m.reasons
    }
}
// GetRequestedValue gets the requestedValue property value. Specifies the requested state of the deployment. Supports a subset of the values for requestedDeploymentStateValue. Possible values are: none, paused, unknownFutureValue.
func (m *DeploymentState) GetRequestedValue()(*RequestedDeploymentStateValue) {
    if m == nil {
        return nil
    } else {
        return m.requestedValue
    }
}
// GetValue gets the value property value. Specifies the state of the deployment. Supports a subset of the values for deploymentStateValue. Possible values are: scheduled, offering, paused, unknownFutureValue. Read-only.
func (m *DeploymentState) GetValue()(*DeploymentStateValue) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *DeploymentState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeploymentState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetReasons() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReasons()))
        for i, v := range m.GetReasons() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("reasons", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRequestedValue() != nil {
        cast := (*m.GetRequestedValue()).String()
        err := writer.WriteStringValue("requestedValue", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetValue() != nil {
        cast := (*m.GetValue()).String()
        err := writer.WriteStringValue("value", &cast)
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
func (m *DeploymentState) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetReasons sets the reasons property value. Specifies the reasons the deployment has its state value. Read-only.
func (m *DeploymentState) SetReasons(value []DeploymentStateReasonable)() {
    if m != nil {
        m.reasons = value
    }
}
// SetRequestedValue sets the requestedValue property value. Specifies the requested state of the deployment. Supports a subset of the values for requestedDeploymentStateValue. Possible values are: none, paused, unknownFutureValue.
func (m *DeploymentState) SetRequestedValue(value *RequestedDeploymentStateValue)() {
    if m != nil {
        m.requestedValue = value
    }
}
// SetValue sets the value property value. Specifies the state of the deployment. Supports a subset of the values for deploymentStateValue. Possible values are: scheduled, offering, paused, unknownFutureValue. Read-only.
func (m *DeploymentState) SetValue(value *DeploymentStateValue)() {
    if m != nil {
        m.value = value
    }
}