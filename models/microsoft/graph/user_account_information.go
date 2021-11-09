package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserAccountInformation struct {
    ItemFacet
    // Shows the age group of user. Allowed values null, minor, notAdult and adult are generated by the directory and cannot be changed.
    ageGroup *string;
    // Contains the two-character country code associated with the users account.
    countryCode *string;
    // 
    preferredLanguageTag *LocaleInfo;
    // The user principal name (UPN) of the user associated with the account.
    userPrincipalName *string;
}
// Instantiates a new userAccountInformation and sets the default values.
func NewUserAccountInformation()(*UserAccountInformation) {
    m := &UserAccountInformation{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// Gets the ageGroup property value. Shows the age group of user. Allowed values null, minor, notAdult and adult are generated by the directory and cannot be changed.
func (m *UserAccountInformation) GetAgeGroup()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ageGroup
    }
}
// Gets the countryCode property value. Contains the two-character country code associated with the users account.
func (m *UserAccountInformation) GetCountryCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryCode
    }
}
// Gets the preferredLanguageTag property value. 
func (m *UserAccountInformation) GetPreferredLanguageTag()(*LocaleInfo) {
    if m == nil {
        return nil
    } else {
        return m.preferredLanguageTag
    }
}
// Gets the userPrincipalName property value. The user principal name (UPN) of the user associated with the account.
func (m *UserAccountInformation) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *UserAccountInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["ageGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAgeGroup(val)
        }
        return nil
    }
    res["countryCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryCode(val)
        }
        return nil
    }
    res["preferredLanguageTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocaleInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredLanguageTag(val.(*LocaleInfo))
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
func (m *UserAccountInformation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserAccountInformation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("ageGroup", m.GetAgeGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("countryCode", m.GetCountryCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("preferredLanguageTag", m.GetPreferredLanguageTag())
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
// Sets the ageGroup property value. Shows the age group of user. Allowed values null, minor, notAdult and adult are generated by the directory and cannot be changed.
// Parameters:
//  - value : Value to set for the ageGroup property.
func (m *UserAccountInformation) SetAgeGroup(value *string)() {
    m.ageGroup = value
}
// Sets the countryCode property value. Contains the two-character country code associated with the users account.
// Parameters:
//  - value : Value to set for the countryCode property.
func (m *UserAccountInformation) SetCountryCode(value *string)() {
    m.countryCode = value
}
// Sets the preferredLanguageTag property value. 
// Parameters:
//  - value : Value to set for the preferredLanguageTag property.
func (m *UserAccountInformation) SetPreferredLanguageTag(value *LocaleInfo)() {
    m.preferredLanguageTag = value
}
// Sets the userPrincipalName property value. The user principal name (UPN) of the user associated with the account.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *UserAccountInformation) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
