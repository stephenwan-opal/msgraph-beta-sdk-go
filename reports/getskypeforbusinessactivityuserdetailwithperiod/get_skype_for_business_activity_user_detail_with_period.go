package getskypeforbusinessactivityuserdetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetSkypeForBusinessActivityUserDetailWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    assignedProducts []string;
    deletedDate *string;
    isDeleted *bool;
    lastActivityDate *string;
    organizedConferenceAppSharingCount *int64;
    organizedConferenceAudioVideoCount *int64;
    organizedConferenceAudioVideoMinutes *int64;
    organizedConferenceCloudDialInMicrosoftMinutes *int64;
    organizedConferenceCloudDialInOutMicrosoftCount *int64;
    organizedConferenceCloudDialOutMicrosoftMinutes *int64;
    organizedConferenceDialInOut3rdPartyCount *int64;
    organizedConferenceIMCount *int64;
    organizedConferenceLastActivityDate *string;
    organizedConferenceWebCount *int64;
    participatedConferenceAppSharingCount *int64;
    participatedConferenceAudioVideoCount *int64;
    participatedConferenceAudioVideoMinutes *int64;
    participatedConferenceDialInOut3rdPartyCount *int64;
    participatedConferenceIMCount *int64;
    participatedConferenceLastActivityDate *string;
    participatedConferenceWebCount *int64;
    peerToPeerAppSharingCount *int64;
    peerToPeerAudioCount *int64;
    peerToPeerAudioMinutes *int64;
    peerToPeerFileTransferCount *int64;
    peerToPeerIMCount *int64;
    peerToPeerLastActivityDate *string;
    peerToPeerVideoCount *int64;
    peerToPeerVideoMinutes *int64;
    reportPeriod *string;
    reportRefreshDate *string;
    totalOrganizedConferenceCount *int64;
    totalParticipatedConferenceCount *int64;
    totalPeerToPeerSessionCount *int64;
    userPrincipalName *string;
}
func NewGetSkypeForBusinessActivityUserDetailWithPeriod()(*GetSkypeForBusinessActivityUserDetailWithPeriod) {
    m := &GetSkypeForBusinessActivityUserDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetAssignedProducts()([]string) {
    if m == nil {
        return nil
    } else {
        return m.assignedProducts
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetDeletedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deletedDate
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetOrganizedConferenceAppSharingCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceAppSharingCount
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetOrganizedConferenceAudioVideoCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceAudioVideoCount
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetOrganizedConferenceAudioVideoMinutes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceAudioVideoMinutes
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetOrganizedConferenceCloudDialInMicrosoftMinutes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceCloudDialInMicrosoftMinutes
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetOrganizedConferenceCloudDialInOutMicrosoftCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceCloudDialInOutMicrosoftCount
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetOrganizedConferenceCloudDialOutMicrosoftMinutes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceCloudDialOutMicrosoftMinutes
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetOrganizedConferenceDialInOut3rdPartyCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceDialInOut3rdPartyCount
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetOrganizedConferenceIMCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceIMCount
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetOrganizedConferenceLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceLastActivityDate
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetOrganizedConferenceWebCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.organizedConferenceWebCount
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetParticipatedConferenceAppSharingCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceAppSharingCount
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetParticipatedConferenceAudioVideoCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceAudioVideoCount
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetParticipatedConferenceAudioVideoMinutes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceAudioVideoMinutes
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetParticipatedConferenceDialInOut3rdPartyCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceDialInOut3rdPartyCount
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetParticipatedConferenceIMCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceIMCount
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetParticipatedConferenceLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceLastActivityDate
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetParticipatedConferenceWebCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.participatedConferenceWebCount
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetPeerToPeerAppSharingCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerAppSharingCount
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetPeerToPeerAudioCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerAudioCount
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetPeerToPeerAudioMinutes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerAudioMinutes
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetPeerToPeerFileTransferCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerFileTransferCount
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetPeerToPeerIMCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerIMCount
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetPeerToPeerLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerLastActivityDate
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetPeerToPeerVideoCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerVideoCount
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetPeerToPeerVideoMinutes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.peerToPeerVideoMinutes
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetTotalOrganizedConferenceCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalOrganizedConferenceCount
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetTotalParticipatedConferenceCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalParticipatedConferenceCount
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetTotalPeerToPeerSessionCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalPeerToPeerSessionCount
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedProducts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetAssignedProducts(res)
        return nil
    }
    res["deletedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeletedDate(val)
        return nil
    }
    res["isDeleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDeleted(val)
        return nil
    }
    res["lastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLastActivityDate(val)
        return nil
    }
    res["organizedConferenceAppSharingCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOrganizedConferenceAppSharingCount(val)
        return nil
    }
    res["organizedConferenceAudioVideoCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOrganizedConferenceAudioVideoCount(val)
        return nil
    }
    res["organizedConferenceAudioVideoMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOrganizedConferenceAudioVideoMinutes(val)
        return nil
    }
    res["organizedConferenceCloudDialInMicrosoftMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOrganizedConferenceCloudDialInMicrosoftMinutes(val)
        return nil
    }
    res["organizedConferenceCloudDialInOutMicrosoftCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOrganizedConferenceCloudDialInOutMicrosoftCount(val)
        return nil
    }
    res["organizedConferenceCloudDialOutMicrosoftMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOrganizedConferenceCloudDialOutMicrosoftMinutes(val)
        return nil
    }
    res["organizedConferenceDialInOut3rdPartyCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOrganizedConferenceDialInOut3rdPartyCount(val)
        return nil
    }
    res["organizedConferenceIMCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOrganizedConferenceIMCount(val)
        return nil
    }
    res["organizedConferenceLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOrganizedConferenceLastActivityDate(val)
        return nil
    }
    res["organizedConferenceWebCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOrganizedConferenceWebCount(val)
        return nil
    }
    res["participatedConferenceAppSharingCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetParticipatedConferenceAppSharingCount(val)
        return nil
    }
    res["participatedConferenceAudioVideoCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetParticipatedConferenceAudioVideoCount(val)
        return nil
    }
    res["participatedConferenceAudioVideoMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetParticipatedConferenceAudioVideoMinutes(val)
        return nil
    }
    res["participatedConferenceDialInOut3rdPartyCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetParticipatedConferenceDialInOut3rdPartyCount(val)
        return nil
    }
    res["participatedConferenceIMCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetParticipatedConferenceIMCount(val)
        return nil
    }
    res["participatedConferenceLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetParticipatedConferenceLastActivityDate(val)
        return nil
    }
    res["participatedConferenceWebCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetParticipatedConferenceWebCount(val)
        return nil
    }
    res["peerToPeerAppSharingCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPeerToPeerAppSharingCount(val)
        return nil
    }
    res["peerToPeerAudioCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPeerToPeerAudioCount(val)
        return nil
    }
    res["peerToPeerAudioMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPeerToPeerAudioMinutes(val)
        return nil
    }
    res["peerToPeerFileTransferCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPeerToPeerFileTransferCount(val)
        return nil
    }
    res["peerToPeerIMCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPeerToPeerIMCount(val)
        return nil
    }
    res["peerToPeerLastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPeerToPeerLastActivityDate(val)
        return nil
    }
    res["peerToPeerVideoCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPeerToPeerVideoCount(val)
        return nil
    }
    res["peerToPeerVideoMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPeerToPeerVideoMinutes(val)
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
    res["totalOrganizedConferenceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetTotalOrganizedConferenceCount(val)
        return nil
    }
    res["totalParticipatedConferenceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetTotalParticipatedConferenceCount(val)
        return nil
    }
    res["totalPeerToPeerSessionCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetTotalPeerToPeerSessionCount(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    return res
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) IsNil()(bool) {
    return m == nil
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("assignedProducts", m.GetAssignedProducts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deletedDate", m.GetDeletedDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeleted", m.GetIsDeleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastActivityDate", m.GetLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("organizedConferenceAppSharingCount", m.GetOrganizedConferenceAppSharingCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("organizedConferenceAudioVideoCount", m.GetOrganizedConferenceAudioVideoCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("organizedConferenceAudioVideoMinutes", m.GetOrganizedConferenceAudioVideoMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("organizedConferenceCloudDialInMicrosoftMinutes", m.GetOrganizedConferenceCloudDialInMicrosoftMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("organizedConferenceCloudDialInOutMicrosoftCount", m.GetOrganizedConferenceCloudDialInOutMicrosoftCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("organizedConferenceCloudDialOutMicrosoftMinutes", m.GetOrganizedConferenceCloudDialOutMicrosoftMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("organizedConferenceDialInOut3rdPartyCount", m.GetOrganizedConferenceDialInOut3rdPartyCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("organizedConferenceIMCount", m.GetOrganizedConferenceIMCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("organizedConferenceLastActivityDate", m.GetOrganizedConferenceLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("organizedConferenceWebCount", m.GetOrganizedConferenceWebCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("participatedConferenceAppSharingCount", m.GetParticipatedConferenceAppSharingCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("participatedConferenceAudioVideoCount", m.GetParticipatedConferenceAudioVideoCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("participatedConferenceAudioVideoMinutes", m.GetParticipatedConferenceAudioVideoMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("participatedConferenceDialInOut3rdPartyCount", m.GetParticipatedConferenceDialInOut3rdPartyCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("participatedConferenceIMCount", m.GetParticipatedConferenceIMCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("participatedConferenceLastActivityDate", m.GetParticipatedConferenceLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("participatedConferenceWebCount", m.GetParticipatedConferenceWebCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("peerToPeerAppSharingCount", m.GetPeerToPeerAppSharingCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("peerToPeerAudioCount", m.GetPeerToPeerAudioCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("peerToPeerAudioMinutes", m.GetPeerToPeerAudioMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("peerToPeerFileTransferCount", m.GetPeerToPeerFileTransferCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("peerToPeerIMCount", m.GetPeerToPeerIMCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("peerToPeerLastActivityDate", m.GetPeerToPeerLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("peerToPeerVideoCount", m.GetPeerToPeerVideoCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("peerToPeerVideoMinutes", m.GetPeerToPeerVideoMinutes())
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
        err = writer.WriteInt64Value("totalOrganizedConferenceCount", m.GetTotalOrganizedConferenceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("totalParticipatedConferenceCount", m.GetTotalParticipatedConferenceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("totalPeerToPeerSessionCount", m.GetTotalPeerToPeerSessionCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetAssignedProducts(value []string)() {
    m.assignedProducts = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetDeletedDate(value *string)() {
    m.deletedDate = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetOrganizedConferenceAppSharingCount(value *int64)() {
    m.organizedConferenceAppSharingCount = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetOrganizedConferenceAudioVideoCount(value *int64)() {
    m.organizedConferenceAudioVideoCount = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetOrganizedConferenceAudioVideoMinutes(value *int64)() {
    m.organizedConferenceAudioVideoMinutes = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetOrganizedConferenceCloudDialInMicrosoftMinutes(value *int64)() {
    m.organizedConferenceCloudDialInMicrosoftMinutes = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetOrganizedConferenceCloudDialInOutMicrosoftCount(value *int64)() {
    m.organizedConferenceCloudDialInOutMicrosoftCount = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetOrganizedConferenceCloudDialOutMicrosoftMinutes(value *int64)() {
    m.organizedConferenceCloudDialOutMicrosoftMinutes = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetOrganizedConferenceDialInOut3rdPartyCount(value *int64)() {
    m.organizedConferenceDialInOut3rdPartyCount = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetOrganizedConferenceIMCount(value *int64)() {
    m.organizedConferenceIMCount = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetOrganizedConferenceLastActivityDate(value *string)() {
    m.organizedConferenceLastActivityDate = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetOrganizedConferenceWebCount(value *int64)() {
    m.organizedConferenceWebCount = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetParticipatedConferenceAppSharingCount(value *int64)() {
    m.participatedConferenceAppSharingCount = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetParticipatedConferenceAudioVideoCount(value *int64)() {
    m.participatedConferenceAudioVideoCount = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetParticipatedConferenceAudioVideoMinutes(value *int64)() {
    m.participatedConferenceAudioVideoMinutes = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetParticipatedConferenceDialInOut3rdPartyCount(value *int64)() {
    m.participatedConferenceDialInOut3rdPartyCount = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetParticipatedConferenceIMCount(value *int64)() {
    m.participatedConferenceIMCount = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetParticipatedConferenceLastActivityDate(value *string)() {
    m.participatedConferenceLastActivityDate = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetParticipatedConferenceWebCount(value *int64)() {
    m.participatedConferenceWebCount = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetPeerToPeerAppSharingCount(value *int64)() {
    m.peerToPeerAppSharingCount = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetPeerToPeerAudioCount(value *int64)() {
    m.peerToPeerAudioCount = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetPeerToPeerAudioMinutes(value *int64)() {
    m.peerToPeerAudioMinutes = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetPeerToPeerFileTransferCount(value *int64)() {
    m.peerToPeerFileTransferCount = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetPeerToPeerIMCount(value *int64)() {
    m.peerToPeerIMCount = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetPeerToPeerLastActivityDate(value *string)() {
    m.peerToPeerLastActivityDate = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetPeerToPeerVideoCount(value *int64)() {
    m.peerToPeerVideoCount = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetPeerToPeerVideoMinutes(value *int64)() {
    m.peerToPeerVideoMinutes = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetTotalOrganizedConferenceCount(value *int64)() {
    m.totalOrganizedConferenceCount = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetTotalParticipatedConferenceCount(value *int64)() {
    m.totalParticipatedConferenceCount = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetTotalPeerToPeerSessionCount(value *int64)() {
    m.totalPeerToPeerSessionCount = value
}
func (m *GetSkypeForBusinessActivityUserDetailWithPeriod) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}