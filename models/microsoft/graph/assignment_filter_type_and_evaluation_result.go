package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AssignmentFilterTypeAndEvaluationResult struct {
    additionalData map[string]interface{};
    assignmentFilterType *DeviceAndAppManagementAssignmentFilterType;
    evaluationResult *AssignmentFilterEvaluationResult;
}
func NewAssignmentFilterTypeAndEvaluationResult()(*AssignmentFilterTypeAndEvaluationResult) {
    m := &AssignmentFilterTypeAndEvaluationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AssignmentFilterTypeAndEvaluationResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AssignmentFilterTypeAndEvaluationResult) GetAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType) {
    if m == nil {
        return nil
    } else {
        return m.assignmentFilterType
    }
}
func (m *AssignmentFilterTypeAndEvaluationResult) GetEvaluationResult()(*AssignmentFilterEvaluationResult) {
    if m == nil {
        return nil
    } else {
        return m.evaluationResult
    }
}
func (m *AssignmentFilterTypeAndEvaluationResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignmentFilterType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAndAppManagementAssignmentFilterType)
        if err != nil {
            return err
        }
        cast := val.(DeviceAndAppManagementAssignmentFilterType)
        m.SetAssignmentFilterType(&cast)
        return nil
    }
    res["evaluationResult"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAssignmentFilterEvaluationResult)
        if err != nil {
            return err
        }
        cast := val.(AssignmentFilterEvaluationResult)
        m.SetEvaluationResult(&cast)
        return nil
    }
    return res
}
func (m *AssignmentFilterTypeAndEvaluationResult) IsNil()(bool) {
    return m == nil
}
func (m *AssignmentFilterTypeAndEvaluationResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAssignmentFilterType() != nil {
        cast := m.GetAssignmentFilterType().String()
        err := writer.WriteStringValue("assignmentFilterType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEvaluationResult() != nil {
        cast := m.GetEvaluationResult().String()
        err := writer.WriteStringValue("evaluationResult", &cast)
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
func (m *AssignmentFilterTypeAndEvaluationResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AssignmentFilterTypeAndEvaluationResult) SetAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)() {
    m.assignmentFilterType = value
}
func (m *AssignmentFilterTypeAndEvaluationResult) SetEvaluationResult(value *AssignmentFilterEvaluationResult)() {
    m.evaluationResult = value
}