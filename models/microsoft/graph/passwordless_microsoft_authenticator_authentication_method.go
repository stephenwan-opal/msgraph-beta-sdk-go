package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PasswordlessMicrosoftAuthenticatorAuthenticationMethod struct {
    AuthenticationMethod
    // 
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The timestamp when this method was registered to the user.
    creationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    device *Device;
    // The display name of the mobile device as given by the user.
    displayName *string;
}
// Instantiates a new passwordlessMicrosoftAuthenticatorAuthenticationMethod and sets the default values.
func NewPasswordlessMicrosoftAuthenticatorAuthenticationMethod()(*PasswordlessMicrosoftAuthenticatorAuthenticationMethod) {
    m := &PasswordlessMicrosoftAuthenticatorAuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    return m
}
// Gets the createdDateTime property value. 
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the creationDateTime property value. The timestamp when this method was registered to the user.
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.creationDateTime
    }
}
// Gets the device property value. 
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) GetDevice()(*Device) {
    if m == nil {
        return nil
    } else {
        return m.device
    }
}
// Gets the displayName property value. The display name of the mobile device as given by the user.
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// The deserialization information for the current model
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AuthenticationMethod.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["creationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationDateTime(val)
        }
        return nil
    }
    res["device"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDevice() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevice(val.(*Device))
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    return res
}
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.AuthenticationMethod.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("creationDateTime", m.GetCreationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("device", m.GetDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the createdDateTime property value. 
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the creationDateTime property value. The timestamp when this method was registered to the user.
// Parameters:
//  - value : Value to set for the creationDateTime property.
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.creationDateTime = value
}
// Sets the device property value. 
// Parameters:
//  - value : Value to set for the device property.
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) SetDevice(value *Device)() {
    m.device = value
}
// Sets the displayName property value. The display name of the mobile device as given by the user.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *PasswordlessMicrosoftAuthenticatorAuthenticationMethod) SetDisplayName(value *string)() {
    m.displayName = value
}
