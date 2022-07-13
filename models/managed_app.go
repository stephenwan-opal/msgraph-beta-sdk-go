package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedApp 
type ManagedApp struct {
    MobileApp
    // A managed (MAM) application's availability.
    appAvailability *ManagedAppAvailability
    // The Application's version.
    version *string
}
// NewManagedApp instantiates a new ManagedApp and sets the default values.
func NewManagedApp()(*ManagedApp) {
    m := &ManagedApp{
        MobileApp: *NewMobileApp(),
    }
    odatatypeValue := "#microsoft.graph.managedApp";
    m.SetType(&odatatypeValue);
    return m
}
// CreateManagedAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.managedAndroidStoreApp":
                        return NewManagedAndroidStoreApp(), nil
                    case "#microsoft.graph.managedIOSStoreApp":
                        return NewManagedIOSStoreApp(), nil
                    case "#microsoft.graph.managedMobileLobApp":
                        return NewManagedMobileLobApp(), nil
                }
            }
        }
    }
    return NewManagedApp(), nil
}
// GetAppAvailability gets the appAvailability property value. A managed (MAM) application's availability.
func (m *ManagedApp) GetAppAvailability()(*ManagedAppAvailability) {
    if m == nil {
        return nil
    } else {
        return m.appAvailability
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["appAvailability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppAvailability)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppAvailability(val.(*ManagedAppAvailability))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetVersion gets the version property value. The Application's version.
func (m *ManagedApp) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// Serialize serializes information the current object
func (m *ManagedApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppAvailability() != nil {
        cast := (*m.GetAppAvailability()).String()
        err = writer.WriteStringValue("appAvailability", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppAvailability sets the appAvailability property value. A managed (MAM) application's availability.
func (m *ManagedApp) SetAppAvailability(value *ManagedAppAvailability)() {
    if m != nil {
        m.appAvailability = value
    }
}
// SetVersion sets the version property value. The Application's version.
func (m *ManagedApp) SetVersion(value *string)() {
    if m != nil {
        m.version = value
    }
}
