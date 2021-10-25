package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SecureScore struct {
    Entity
    activeUserCount *int32;
    averageComparativeScores []AverageComparativeScore;
    azureTenantId *string;
    controlScores []ControlScore;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    currentScore *float64;
    enabledServices []string;
    licensedUserCount *int32;
    maxScore *float64;
    vendorInformation *SecurityVendorInformation;
}
func NewSecureScore()(*SecureScore) {
    m := &SecureScore{
        Entity: *NewEntity(),
    }
    return m
}
func (m *SecureScore) GetActiveUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activeUserCount
    }
}
func (m *SecureScore) GetAverageComparativeScores()([]AverageComparativeScore) {
    if m == nil {
        return nil
    } else {
        return m.averageComparativeScores
    }
}
func (m *SecureScore) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
func (m *SecureScore) GetControlScores()([]ControlScore) {
    if m == nil {
        return nil
    } else {
        return m.controlScores
    }
}
func (m *SecureScore) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *SecureScore) GetCurrentScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.currentScore
    }
}
func (m *SecureScore) GetEnabledServices()([]string) {
    if m == nil {
        return nil
    } else {
        return m.enabledServices
    }
}
func (m *SecureScore) GetLicensedUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.licensedUserCount
    }
}
func (m *SecureScore) GetMaxScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.maxScore
    }
}
func (m *SecureScore) GetVendorInformation()(*SecurityVendorInformation) {
    if m == nil {
        return nil
    } else {
        return m.vendorInformation
    }
}
func (m *SecureScore) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetActiveUserCount(val)
        return nil
    }
    res["averageComparativeScores"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAverageComparativeScore() })
        if err != nil {
            return err
        }
        res := make([]AverageComparativeScore, len(val))
        for i, v := range val {
            res[i] = *(v.(*AverageComparativeScore))
        }
        m.SetAverageComparativeScores(res)
        return nil
    }
    res["azureTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAzureTenantId(val)
        return nil
    }
    res["controlScores"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewControlScore() })
        if err != nil {
            return err
        }
        res := make([]ControlScore, len(val))
        for i, v := range val {
            res[i] = *(v.(*ControlScore))
        }
        m.SetControlScores(res)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["currentScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetCurrentScore(val)
        return nil
    }
    res["enabledServices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetEnabledServices(res)
        return nil
    }
    res["licensedUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetLicensedUserCount(val)
        return nil
    }
    res["maxScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetMaxScore(val)
        return nil
    }
    res["vendorInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityVendorInformation() })
        if err != nil {
            return err
        }
        m.SetVendorInformation(val.(*SecurityVendorInformation))
        return nil
    }
    return res
}
func (m *SecureScore) IsNil()(bool) {
    return m == nil
}
func (m *SecureScore) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("activeUserCount", m.GetActiveUserCount())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAverageComparativeScores()))
        for i, v := range m.GetAverageComparativeScores() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("averageComparativeScores", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureTenantId", m.GetAzureTenantId())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetControlScores()))
        for i, v := range m.GetControlScores() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("controlScores", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("currentScore", m.GetCurrentScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("enabledServices", m.GetEnabledServices())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("licensedUserCount", m.GetLicensedUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("maxScore", m.GetMaxScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("vendorInformation", m.GetVendorInformation())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *SecureScore) SetActiveUserCount(value *int32)() {
    m.activeUserCount = value
}
func (m *SecureScore) SetAverageComparativeScores(value []AverageComparativeScore)() {
    m.averageComparativeScores = value
}
func (m *SecureScore) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
func (m *SecureScore) SetControlScores(value []ControlScore)() {
    m.controlScores = value
}
func (m *SecureScore) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *SecureScore) SetCurrentScore(value *float64)() {
    m.currentScore = value
}
func (m *SecureScore) SetEnabledServices(value []string)() {
    m.enabledServices = value
}
func (m *SecureScore) SetLicensedUserCount(value *int32)() {
    m.licensedUserCount = value
}
func (m *SecureScore) SetMaxScore(value *float64)() {
    m.maxScore = value
}
func (m *SecureScore) SetVendorInformation(value *SecurityVendorInformation)() {
    m.vendorInformation = value
}