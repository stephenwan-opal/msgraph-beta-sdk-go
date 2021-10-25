package getyammergroupsactivitydetailwithdate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetYammerGroupsActivityDetailWithDate struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    groupDisplayName *string;
    groupType *string;
    isDeleted *bool;
    lastActivityDate *string;
    likedCount *int64;
    memberCount *int64;
    networkDisplayName *string;
    office365Connected *bool;
    ownerPrincipalName *string;
    postedCount *int64;
    readCount *int64;
    reportPeriod *string;
    reportRefreshDate *string;
}
func NewGetYammerGroupsActivityDetailWithDate()(*GetYammerGroupsActivityDetailWithDate) {
    m := &GetYammerGroupsActivityDetailWithDate{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetYammerGroupsActivityDetailWithDate) GetGroupDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupDisplayName
    }
}
func (m *GetYammerGroupsActivityDetailWithDate) GetGroupType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupType
    }
}
func (m *GetYammerGroupsActivityDetailWithDate) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
func (m *GetYammerGroupsActivityDetailWithDate) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
func (m *GetYammerGroupsActivityDetailWithDate) GetLikedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.likedCount
    }
}
func (m *GetYammerGroupsActivityDetailWithDate) GetMemberCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.memberCount
    }
}
func (m *GetYammerGroupsActivityDetailWithDate) GetNetworkDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.networkDisplayName
    }
}
func (m *GetYammerGroupsActivityDetailWithDate) GetOffice365Connected()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.office365Connected
    }
}
func (m *GetYammerGroupsActivityDetailWithDate) GetOwnerPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerPrincipalName
    }
}
func (m *GetYammerGroupsActivityDetailWithDate) GetPostedCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.postedCount
    }
}
func (m *GetYammerGroupsActivityDetailWithDate) GetReadCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.readCount
    }
}
func (m *GetYammerGroupsActivityDetailWithDate) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetYammerGroupsActivityDetailWithDate) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetYammerGroupsActivityDetailWithDate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["groupDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGroupDisplayName(val)
        return nil
    }
    res["groupType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGroupType(val)
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
    res["likedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetLikedCount(val)
        return nil
    }
    res["memberCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetMemberCount(val)
        return nil
    }
    res["networkDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNetworkDisplayName(val)
        return nil
    }
    res["office365Connected"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetOffice365Connected(val)
        return nil
    }
    res["ownerPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOwnerPrincipalName(val)
        return nil
    }
    res["postedCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPostedCount(val)
        return nil
    }
    res["readCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetReadCount(val)
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
    return res
}
func (m *GetYammerGroupsActivityDetailWithDate) IsNil()(bool) {
    return m == nil
}
func (m *GetYammerGroupsActivityDetailWithDate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("groupDisplayName", m.GetGroupDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupType", m.GetGroupType())
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
        err = writer.WriteInt64Value("likedCount", m.GetLikedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("memberCount", m.GetMemberCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("networkDisplayName", m.GetNetworkDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("office365Connected", m.GetOffice365Connected())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerPrincipalName", m.GetOwnerPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("postedCount", m.GetPostedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("readCount", m.GetReadCount())
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
    return nil
}
func (m *GetYammerGroupsActivityDetailWithDate) SetGroupDisplayName(value *string)() {
    m.groupDisplayName = value
}
func (m *GetYammerGroupsActivityDetailWithDate) SetGroupType(value *string)() {
    m.groupType = value
}
func (m *GetYammerGroupsActivityDetailWithDate) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
func (m *GetYammerGroupsActivityDetailWithDate) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
func (m *GetYammerGroupsActivityDetailWithDate) SetLikedCount(value *int64)() {
    m.likedCount = value
}
func (m *GetYammerGroupsActivityDetailWithDate) SetMemberCount(value *int64)() {
    m.memberCount = value
}
func (m *GetYammerGroupsActivityDetailWithDate) SetNetworkDisplayName(value *string)() {
    m.networkDisplayName = value
}
func (m *GetYammerGroupsActivityDetailWithDate) SetOffice365Connected(value *bool)() {
    m.office365Connected = value
}
func (m *GetYammerGroupsActivityDetailWithDate) SetOwnerPrincipalName(value *string)() {
    m.ownerPrincipalName = value
}
func (m *GetYammerGroupsActivityDetailWithDate) SetPostedCount(value *int64)() {
    m.postedCount = value
}
func (m *GetYammerGroupsActivityDetailWithDate) SetReadCount(value *int64)() {
    m.readCount = value
}
func (m *GetYammerGroupsActivityDetailWithDate) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetYammerGroupsActivityDetailWithDate) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}