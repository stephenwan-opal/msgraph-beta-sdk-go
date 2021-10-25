package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserExperienceAnalyticsScoreHistory struct {
    Entity
    startupDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
func NewUserExperienceAnalyticsScoreHistory()(*UserExperienceAnalyticsScoreHistory) {
    m := &UserExperienceAnalyticsScoreHistory{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UserExperienceAnalyticsScoreHistory) GetStartupDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startupDateTime
    }
}
func (m *UserExperienceAnalyticsScoreHistory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["startupDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetStartupDateTime(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsScoreHistory) IsNil()(bool) {
    return m == nil
}
func (m *UserExperienceAnalyticsScoreHistory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("startupDateTime", m.GetStartupDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *UserExperienceAnalyticsScoreHistory) SetStartupDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startupDateTime = value
}