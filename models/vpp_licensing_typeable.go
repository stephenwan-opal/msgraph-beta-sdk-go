package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VppLicensingTypeable 
type VppLicensingTypeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSupportDeviceLicensing()(*bool)
    GetSupportsDeviceLicensing()(*bool)
    GetSupportsUserLicensing()(*bool)
    GetSupportUserLicensing()(*bool)
    SetSupportDeviceLicensing(value *bool)()
    SetSupportsDeviceLicensing(value *bool)()
    SetSupportsUserLicensing(value *bool)()
    SetSupportUserLicensing(value *bool)()
}
