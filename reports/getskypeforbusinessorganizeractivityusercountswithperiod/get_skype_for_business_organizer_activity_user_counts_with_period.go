package getskypeforbusinessorganizeractivityusercountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    appSharing *int64;
    audioVideo *int64;
    dialInOut3rdParty *int64;
    dialInOutMicrosoft *int64;
    im *int64;
    reportDate *string;
    reportPeriod *string;
    reportRefreshDate *string;
    web *int64;
}
func NewGetSkypeForBusinessOrganizerActivityUserCountsWithPeriod()(*GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) {
    m := &GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) GetAppSharing()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.appSharing
    }
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) GetAudioVideo()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.audioVideo
    }
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) GetDialInOut3rdParty()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.dialInOut3rdParty
    }
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) GetDialInOutMicrosoft()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.dialInOutMicrosoft
    }
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) GetIm()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.im
    }
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) GetWeb()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.web
    }
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appSharing"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetAppSharing(val)
        return nil
    }
    res["audioVideo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetAudioVideo(val)
        return nil
    }
    res["dialInOut3rdParty"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetDialInOut3rdParty(val)
        return nil
    }
    res["dialInOutMicrosoft"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetDialInOutMicrosoft(val)
        return nil
    }
    res["im"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetIm(val)
        return nil
    }
    res["reportDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportDate(val)
        return nil
    }
    res["reportPeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportPeriod(val)
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportRefreshDate(val)
        return nil
    }
    res["web"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetWeb(val)
        return nil
    }
    return res
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("appSharing", m.GetAppSharing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("audioVideo", m.GetAudioVideo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("dialInOut3rdParty", m.GetDialInOut3rdParty())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("dialInOutMicrosoft", m.GetDialInOutMicrosoft())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("im", m.GetIm())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportDate", m.GetReportDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportPeriod", m.GetReportPeriod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportRefreshDate", m.GetReportRefreshDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("web", m.GetWeb())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) SetAppSharing(value *int64)() {
    m.appSharing = value
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) SetAudioVideo(value *int64)() {
    m.audioVideo = value
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) SetDialInOut3rdParty(value *int64)() {
    m.dialInOut3rdParty = value
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) SetDialInOutMicrosoft(value *int64)() {
    m.dialInOutMicrosoft = value
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) SetIm(value *int64)() {
    m.im = value
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod) SetWeb(value *int64)() {
    m.web = value
}