package updaterequest

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type UpdateRequestRequestBody struct {
    additionalData map[string]interface{};
    assignmentState *string;
    decision *string;
    reason *string;
    schedule *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceSchedule;
}
func NewUpdateRequestRequestBody()(*UpdateRequestRequestBody) {
    m := &UpdateRequestRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UpdateRequestRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UpdateRequestRequestBody) GetAssignmentState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignmentState
    }
}
func (m *UpdateRequestRequestBody) GetDecision()(*string) {
    if m == nil {
        return nil
    } else {
        return m.decision
    }
}
func (m *UpdateRequestRequestBody) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
func (m *UpdateRequestRequestBody) GetSchedule()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceSchedule) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
func (m *UpdateRequestRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignmentState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAssignmentState(val)
        return nil
    }
    res["decision"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDecision(val)
        return nil
    }
    res["reason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReason(val)
        return nil
    }
    res["schedule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewGovernanceSchedule() })
        if err != nil {
            return err
        }
        m.SetSchedule(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceSchedule))
        return nil
    }
    return res
}
func (m *UpdateRequestRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *UpdateRequestRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("assignmentState", m.GetAssignmentState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("decision", m.GetDecision())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reason", m.GetReason())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("schedule", m.GetSchedule())
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
func (m *UpdateRequestRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UpdateRequestRequestBody) SetAssignmentState(value *string)() {
    m.assignmentState = value
}
func (m *UpdateRequestRequestBody) SetDecision(value *string)() {
    m.decision = value
}
func (m *UpdateRequestRequestBody) SetReason(value *string)() {
    m.reason = value
}
func (m *UpdateRequestRequestBody) SetSchedule(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceSchedule)() {
    m.schedule = value
}