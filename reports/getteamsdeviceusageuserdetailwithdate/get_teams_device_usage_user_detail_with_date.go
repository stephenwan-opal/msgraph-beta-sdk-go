package getteamsdeviceusageuserdetailwithdate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetTeamsDeviceUsageUserDetailWithDate struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    deletedDate *string;
    isDeleted *bool;
    isLicensed *bool;
    lastActivityDate *string;
    reportPeriod *string;
    reportRefreshDate *string;
    usedAndroidPhone *bool;
    usedChromeOS *bool;
    usediOS *bool;
    usedLinux *bool;
    usedMac *bool;
    usedWeb *bool;
    usedWindows *bool;
    usedWindowsPhone *bool;
    userPrincipalName *string;
}
func NewGetTeamsDeviceUsageUserDetailWithDate()(*GetTeamsDeviceUsageUserDetailWithDate) {
    m := &GetTeamsDeviceUsageUserDetailWithDate{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) GetDeletedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deletedDate
    }
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) GetIsLicensed()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isLicensed
    }
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) GetUsedAndroidPhone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedAndroidPhone
    }
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) GetUsedChromeOS()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedChromeOS
    }
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) GetUsediOS()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usediOS
    }
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) GetUsedLinux()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedLinux
    }
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) GetUsedMac()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedMac
    }
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) GetUsedWeb()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedWeb
    }
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) GetUsedWindows()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedWindows
    }
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) GetUsedWindowsPhone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedWindowsPhone
    }
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["isLicensed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsLicensed(val)
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
    res["usedAndroidPhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUsedAndroidPhone(val)
        return nil
    }
    res["usedChromeOS"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUsedChromeOS(val)
        return nil
    }
    res["usediOS"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUsediOS(val)
        return nil
    }
    res["usedLinux"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUsedLinux(val)
        return nil
    }
    res["usedMac"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUsedMac(val)
        return nil
    }
    res["usedWeb"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUsedWeb(val)
        return nil
    }
    res["usedWindows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUsedWindows(val)
        return nil
    }
    res["usedWindowsPhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUsedWindowsPhone(val)
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
func (m *GetTeamsDeviceUsageUserDetailWithDate) IsNil()(bool) {
    return m == nil
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteBoolValue("isLicensed", m.GetIsLicensed())
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
        err = writer.WriteBoolValue("usedAndroidPhone", m.GetUsedAndroidPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usedChromeOS", m.GetUsedChromeOS())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usediOS", m.GetUsediOS())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usedLinux", m.GetUsedLinux())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usedMac", m.GetUsedMac())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usedWeb", m.GetUsedWeb())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usedWindows", m.GetUsedWindows())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usedWindowsPhone", m.GetUsedWindowsPhone())
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
func (m *GetTeamsDeviceUsageUserDetailWithDate) SetDeletedDate(value *string)() {
    m.deletedDate = value
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) SetIsLicensed(value *bool)() {
    m.isLicensed = value
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) SetUsedAndroidPhone(value *bool)() {
    m.usedAndroidPhone = value
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) SetUsedChromeOS(value *bool)() {
    m.usedChromeOS = value
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) SetUsediOS(value *bool)() {
    m.usediOS = value
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) SetUsedLinux(value *bool)() {
    m.usedLinux = value
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) SetUsedMac(value *bool)() {
    m.usedMac = value
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) SetUsedWeb(value *bool)() {
    m.usedWeb = value
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) SetUsedWindows(value *bool)() {
    m.usedWindows = value
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) SetUsedWindowsPhone(value *bool)() {
    m.usedWindowsPhone = value
}
func (m *GetTeamsDeviceUsageUserDetailWithDate) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}