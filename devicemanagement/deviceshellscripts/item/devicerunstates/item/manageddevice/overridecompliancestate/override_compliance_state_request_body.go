package overridecompliancestate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type OverrideComplianceStateRequestBody struct {
    additionalData map[string]interface{};
    complianceState *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdministratorConfiguredDeviceComplianceState;
    remediationUrl *string;
}
func NewOverrideComplianceStateRequestBody()(*OverrideComplianceStateRequestBody) {
    m := &OverrideComplianceStateRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *OverrideComplianceStateRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *OverrideComplianceStateRequestBody) GetComplianceState()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdministratorConfiguredDeviceComplianceState) {
    if m == nil {
        return nil
    } else {
        return m.complianceState
    }
}
func (m *OverrideComplianceStateRequestBody) GetRemediationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remediationUrl
    }
}
func (m *OverrideComplianceStateRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["complianceState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ParseAdministratorConfiguredDeviceComplianceState)
        if err != nil {
            return err
        }
        cast := val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdministratorConfiguredDeviceComplianceState)
        m.SetComplianceState(&cast)
        return nil
    }
    res["remediationUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRemediationUrl(val)
        return nil
    }
    return res
}
func (m *OverrideComplianceStateRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *OverrideComplianceStateRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetComplianceState() != nil {
        cast := m.GetComplianceState().String()
        err := writer.WriteStringValue("complianceState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("remediationUrl", m.GetRemediationUrl())
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
func (m *OverrideComplianceStateRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *OverrideComplianceStateRequestBody) SetComplianceState(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdministratorConfiguredDeviceComplianceState)() {
    m.complianceState = value
}
func (m *OverrideComplianceStateRequestBody) SetRemediationUrl(value *string)() {
    m.remediationUrl = value
}