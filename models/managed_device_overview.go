package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceOverview 
type ManagedDeviceOverview struct {
    Entity
    // Distribution of Exchange Access State in Intune
    deviceExchangeAccessStateSummary DeviceExchangeAccessStateSummaryable;
    // Device operating system summary.
    deviceOperatingSystemSummary DeviceOperatingSystemSummaryable;
    // The number of devices enrolled in both MDM and EAS
    dualEnrolledDeviceCount *int32;
    // Total enrolled device count. Does not include PC devices managed via Intune PC Agent
    enrolledDeviceCount *int32;
    // Last modified date time of device overview
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Models and Manufactures meatadata for managed devices in the account
    managedDeviceModelsAndManufacturers ManagedDeviceModelsAndManufacturersable;
    // The number of devices enrolled in MDM
    mdmEnrolledCount *int32;
}
// NewManagedDeviceOverview instantiates a new managedDeviceOverview and sets the default values.
func NewManagedDeviceOverview()(*ManagedDeviceOverview) {
    m := &ManagedDeviceOverview{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedDeviceOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDeviceOverview(), nil
}
// GetDeviceExchangeAccessStateSummary gets the deviceExchangeAccessStateSummary property value. Distribution of Exchange Access State in Intune
func (m *ManagedDeviceOverview) GetDeviceExchangeAccessStateSummary()(DeviceExchangeAccessStateSummaryable) {
    if m == nil {
        return nil
    } else {
        return m.deviceExchangeAccessStateSummary
    }
}
// GetDeviceOperatingSystemSummary gets the deviceOperatingSystemSummary property value. Device operating system summary.
func (m *ManagedDeviceOverview) GetDeviceOperatingSystemSummary()(DeviceOperatingSystemSummaryable) {
    if m == nil {
        return nil
    } else {
        return m.deviceOperatingSystemSummary
    }
}
// GetDualEnrolledDeviceCount gets the dualEnrolledDeviceCount property value. The number of devices enrolled in both MDM and EAS
func (m *ManagedDeviceOverview) GetDualEnrolledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.dualEnrolledDeviceCount
    }
}
// GetEnrolledDeviceCount gets the enrolledDeviceCount property value. Total enrolled device count. Does not include PC devices managed via Intune PC Agent
func (m *ManagedDeviceOverview) GetEnrolledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.enrolledDeviceCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceOverview) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceExchangeAccessStateSummary"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceExchangeAccessStateSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceExchangeAccessStateSummary(val.(DeviceExchangeAccessStateSummaryable))
        }
        return nil
    }
    res["deviceOperatingSystemSummary"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceOperatingSystemSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceOperatingSystemSummary(val.(DeviceOperatingSystemSummaryable))
        }
        return nil
    }
    res["dualEnrolledDeviceCount"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDualEnrolledDeviceCount(val)
        }
        return nil
    }
    res["enrolledDeviceCount"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrolledDeviceCount(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["managedDeviceModelsAndManufacturers"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedDeviceModelsAndManufacturersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceModelsAndManufacturers(val.(ManagedDeviceModelsAndManufacturersable))
        }
        return nil
    }
    res["mdmEnrolledCount"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdmEnrolledCount(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last modified date time of device overview
func (m *ManagedDeviceOverview) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetManagedDeviceModelsAndManufacturers gets the managedDeviceModelsAndManufacturers property value. Models and Manufactures meatadata for managed devices in the account
func (m *ManagedDeviceOverview) GetManagedDeviceModelsAndManufacturers()(ManagedDeviceModelsAndManufacturersable) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceModelsAndManufacturers
    }
}
// GetMdmEnrolledCount gets the mdmEnrolledCount property value. The number of devices enrolled in MDM
func (m *ManagedDeviceOverview) GetMdmEnrolledCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.mdmEnrolledCount
    }
}
// Serialize serializes information the current object
func (m *ManagedDeviceOverview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("deviceExchangeAccessStateSummary", m.GetDeviceExchangeAccessStateSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceOperatingSystemSummary", m.GetDeviceOperatingSystemSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("dualEnrolledDeviceCount", m.GetDualEnrolledDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("enrolledDeviceCount", m.GetEnrolledDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("managedDeviceModelsAndManufacturers", m.GetManagedDeviceModelsAndManufacturers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("mdmEnrolledCount", m.GetMdmEnrolledCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceExchangeAccessStateSummary sets the deviceExchangeAccessStateSummary property value. Distribution of Exchange Access State in Intune
func (m *ManagedDeviceOverview) SetDeviceExchangeAccessStateSummary(value DeviceExchangeAccessStateSummaryable)() {
    if m != nil {
        m.deviceExchangeAccessStateSummary = value
    }
}
// SetDeviceOperatingSystemSummary sets the deviceOperatingSystemSummary property value. Device operating system summary.
func (m *ManagedDeviceOverview) SetDeviceOperatingSystemSummary(value DeviceOperatingSystemSummaryable)() {
    if m != nil {
        m.deviceOperatingSystemSummary = value
    }
}
// SetDualEnrolledDeviceCount sets the dualEnrolledDeviceCount property value. The number of devices enrolled in both MDM and EAS
func (m *ManagedDeviceOverview) SetDualEnrolledDeviceCount(value *int32)() {
    if m != nil {
        m.dualEnrolledDeviceCount = value
    }
}
// SetEnrolledDeviceCount sets the enrolledDeviceCount property value. Total enrolled device count. Does not include PC devices managed via Intune PC Agent
func (m *ManagedDeviceOverview) SetEnrolledDeviceCount(value *int32)() {
    if m != nil {
        m.enrolledDeviceCount = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last modified date time of device overview
func (m *ManagedDeviceOverview) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetManagedDeviceModelsAndManufacturers sets the managedDeviceModelsAndManufacturers property value. Models and Manufactures meatadata for managed devices in the account
func (m *ManagedDeviceOverview) SetManagedDeviceModelsAndManufacturers(value ManagedDeviceModelsAndManufacturersable)() {
    if m != nil {
        m.managedDeviceModelsAndManufacturers = value
    }
}
// SetMdmEnrolledCount sets the mdmEnrolledCount property value. The number of devices enrolled in MDM
func (m *ManagedDeviceOverview) SetMdmEnrolledCount(value *int32)() {
    if m != nil {
        m.mdmEnrolledCount = value
    }
}