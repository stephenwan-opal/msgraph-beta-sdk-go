package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedAndroidLobApp 
type ManagedAndroidLobApp struct {
    ManagedMobileLobApp
    // The Identity Name.
    identityName *string
    // The identity version.
    identityVersion *string
    // The value for the minimum applicable operating system.
    minimumSupportedOperatingSystem AndroidMinimumOperatingSystemable
    // The package identifier.
    packageId *string
    // The version code of managed Android Line of Business (LoB) app.
    versionCode *string
    // The version name of managed Android Line of Business (LoB) app.
    versionName *string
}
// NewManagedAndroidLobApp instantiates a new ManagedAndroidLobApp and sets the default values.
func NewManagedAndroidLobApp()(*ManagedAndroidLobApp) {
    m := &ManagedAndroidLobApp{
        ManagedMobileLobApp: *NewManagedMobileLobApp(),
    }
    return m
}
// CreateManagedAndroidLobAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedAndroidLobAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedAndroidLobApp(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAndroidLobApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedMobileLobApp.GetFieldDeserializers()
    res["identityName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityName(val)
        }
        return nil
    }
    res["identityVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityVersion(val)
        }
        return nil
    }
    res["minimumSupportedOperatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidMinimumOperatingSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedOperatingSystem(val.(AndroidMinimumOperatingSystemable))
        }
        return nil
    }
    res["packageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageId(val)
        }
        return nil
    }
    res["versionCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionCode(val)
        }
        return nil
    }
    res["versionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionName(val)
        }
        return nil
    }
    return res
}
// GetIdentityName gets the identityName property value. The Identity Name.
func (m *ManagedAndroidLobApp) GetIdentityName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.identityName
    }
}
// GetIdentityVersion gets the identityVersion property value. The identity version.
func (m *ManagedAndroidLobApp) GetIdentityVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.identityVersion
    }
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *ManagedAndroidLobApp) GetMinimumSupportedOperatingSystem()(AndroidMinimumOperatingSystemable) {
    if m == nil {
        return nil
    } else {
        return m.minimumSupportedOperatingSystem
    }
}
// GetPackageId gets the packageId property value. The package identifier.
func (m *ManagedAndroidLobApp) GetPackageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.packageId
    }
}
// GetVersionCode gets the versionCode property value. The version code of managed Android Line of Business (LoB) app.
func (m *ManagedAndroidLobApp) GetVersionCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.versionCode
    }
}
// GetVersionName gets the versionName property value. The version name of managed Android Line of Business (LoB) app.
func (m *ManagedAndroidLobApp) GetVersionName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.versionName
    }
}
// Serialize serializes information the current object
func (m *ManagedAndroidLobApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ManagedMobileLobApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("identityName", m.GetIdentityName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identityVersion", m.GetIdentityVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minimumSupportedOperatingSystem", m.GetMinimumSupportedOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("packageId", m.GetPackageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("versionCode", m.GetVersionCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("versionName", m.GetVersionName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIdentityName sets the identityName property value. The Identity Name.
func (m *ManagedAndroidLobApp) SetIdentityName(value *string)() {
    if m != nil {
        m.identityName = value
    }
}
// SetIdentityVersion sets the identityVersion property value. The identity version.
func (m *ManagedAndroidLobApp) SetIdentityVersion(value *string)() {
    if m != nil {
        m.identityVersion = value
    }
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *ManagedAndroidLobApp) SetMinimumSupportedOperatingSystem(value AndroidMinimumOperatingSystemable)() {
    if m != nil {
        m.minimumSupportedOperatingSystem = value
    }
}
// SetPackageId sets the packageId property value. The package identifier.
func (m *ManagedAndroidLobApp) SetPackageId(value *string)() {
    if m != nil {
        m.packageId = value
    }
}
// SetVersionCode sets the versionCode property value. The version code of managed Android Line of Business (LoB) app.
func (m *ManagedAndroidLobApp) SetVersionCode(value *string)() {
    if m != nil {
        m.versionCode = value
    }
}
// SetVersionName sets the versionName property value. The version name of managed Android Line of Business (LoB) app.
func (m *ManagedAndroidLobApp) SetVersionName(value *string)() {
    if m != nil {
        m.versionName = value
    }
}
