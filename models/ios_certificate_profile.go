package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosCertificateProfile device Configuration.
type IosCertificateProfile struct {
    DeviceConfiguration
}
// NewIosCertificateProfile instantiates a new iosCertificateProfile and sets the default values.
func NewIosCertificateProfile()(*IosCertificateProfile) {
    m := &IosCertificateProfile{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    return m
}
// CreateIosCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.iosCertificateProfileBase":
                        return NewIosCertificateProfileBase(), nil
                    case "#microsoft.graph.iosImportedPFXCertificateProfile":
                        return NewIosImportedPFXCertificateProfile(), nil
                }
            }
        }
    }
    return NewIosCertificateProfile(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *IosCertificateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
