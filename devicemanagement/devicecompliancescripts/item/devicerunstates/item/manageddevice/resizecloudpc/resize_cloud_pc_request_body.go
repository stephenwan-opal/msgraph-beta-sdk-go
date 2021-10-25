package resizecloudpc

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ResizeCloudPcRequestBody struct {
    additionalData map[string]interface{};
    targetServicePlanId *string;
}
func NewResizeCloudPcRequestBody()(*ResizeCloudPcRequestBody) {
    m := &ResizeCloudPcRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ResizeCloudPcRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ResizeCloudPcRequestBody) GetTargetServicePlanId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetServicePlanId
    }
}
func (m *ResizeCloudPcRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["targetServicePlanId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetServicePlanId(val)
        return nil
    }
    return res
}
func (m *ResizeCloudPcRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *ResizeCloudPcRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("targetServicePlanId", m.GetTargetServicePlanId())
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
func (m *ResizeCloudPcRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ResizeCloudPcRequestBody) SetTargetServicePlanId(value *string)() {
    m.targetServicePlanId = value
}