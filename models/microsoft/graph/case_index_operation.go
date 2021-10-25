package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CaseIndexOperation struct {
    CaseOperation
}
func NewCaseIndexOperation()(*CaseIndexOperation) {
    m := &CaseIndexOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
func (m *CaseIndexOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    return res
}
func (m *CaseIndexOperation) IsNil()(bool) {
    return m == nil
}
func (m *CaseIndexOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CaseOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}