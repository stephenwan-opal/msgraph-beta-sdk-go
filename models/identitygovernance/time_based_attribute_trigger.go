package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeBasedAttributeTrigger 
type TimeBasedAttributeTrigger struct {
    WorkflowExecutionTrigger
    // How many days before or after the time-based attribute specified the workflow should trigger. For example, if the attribute is employeeHireDate and offsetInDays is -1, then the workflow should trigger one day before the employee hire date.
    offsetInDays *int32
    // The timeBasedAttribute property
    timeBasedAttribute *WorkflowTriggerTimeBasedAttribute
}
// NewTimeBasedAttributeTrigger instantiates a new TimeBasedAttributeTrigger and sets the default values.
func NewTimeBasedAttributeTrigger()(*TimeBasedAttributeTrigger) {
    m := &TimeBasedAttributeTrigger{
        WorkflowExecutionTrigger: *NewWorkflowExecutionTrigger(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.timeBasedAttributeTrigger";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTimeBasedAttributeTriggerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeBasedAttributeTriggerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimeBasedAttributeTrigger(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeBasedAttributeTrigger) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WorkflowExecutionTrigger.GetFieldDeserializers()
    res["offsetInDays"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetOffsetInDays)
    res["timeBasedAttribute"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWorkflowTriggerTimeBasedAttribute , m.SetTimeBasedAttribute)
    return res
}
// GetOffsetInDays gets the offsetInDays property value. How many days before or after the time-based attribute specified the workflow should trigger. For example, if the attribute is employeeHireDate and offsetInDays is -1, then the workflow should trigger one day before the employee hire date.
func (m *TimeBasedAttributeTrigger) GetOffsetInDays()(*int32) {
    return m.offsetInDays
}
// GetTimeBasedAttribute gets the timeBasedAttribute property value. The timeBasedAttribute property
func (m *TimeBasedAttributeTrigger) GetTimeBasedAttribute()(*WorkflowTriggerTimeBasedAttribute) {
    return m.timeBasedAttribute
}
// Serialize serializes information the current object
func (m *TimeBasedAttributeTrigger) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WorkflowExecutionTrigger.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("offsetInDays", m.GetOffsetInDays())
        if err != nil {
            return err
        }
    }
    if m.GetTimeBasedAttribute() != nil {
        cast := (*m.GetTimeBasedAttribute()).String()
        err = writer.WriteStringValue("timeBasedAttribute", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOffsetInDays sets the offsetInDays property value. How many days before or after the time-based attribute specified the workflow should trigger. For example, if the attribute is employeeHireDate and offsetInDays is -1, then the workflow should trigger one day before the employee hire date.
func (m *TimeBasedAttributeTrigger) SetOffsetInDays(value *int32)() {
    m.offsetInDays = value
}
// SetTimeBasedAttribute sets the timeBasedAttribute property value. The timeBasedAttribute property
func (m *TimeBasedAttributeTrigger) SetTimeBasedAttribute(value *WorkflowTriggerTimeBasedAttribute)() {
    m.timeBasedAttribute = value
}