package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AppManagementConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    keyCredentials []KeyCredentialConfiguration;
    // Collection of password restrictions settings to be applied to an application or service principal
    passwordCredentials []PasswordCredentialConfiguration;
}
// Instantiates a new appManagementConfiguration and sets the default values.
func NewAppManagementConfiguration()(*AppManagementConfiguration) {
    m := &AppManagementConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppManagementConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the keyCredentials property value. 
func (m *AppManagementConfiguration) GetKeyCredentials()([]KeyCredentialConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.keyCredentials
    }
}
// Gets the passwordCredentials property value. Collection of password restrictions settings to be applied to an application or service principal
func (m *AppManagementConfiguration) GetPasswordCredentials()([]PasswordCredentialConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.passwordCredentials
    }
}
// The deserialization information for the current model
func (m *AppManagementConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["keyCredentials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyCredentialConfiguration() })
        if err != nil {
            return err
        }
        res := make([]KeyCredentialConfiguration, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyCredentialConfiguration))
        }
        m.SetKeyCredentials(res)
        return nil
    }
    res["passwordCredentials"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPasswordCredentialConfiguration() })
        if err != nil {
            return err
        }
        res := make([]PasswordCredentialConfiguration, len(val))
        for i, v := range val {
            res[i] = *(v.(*PasswordCredentialConfiguration))
        }
        m.SetPasswordCredentials(res)
        return nil
    }
    return res
}
func (m *AppManagementConfiguration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AppManagementConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetKeyCredentials()))
        for i, v := range m.GetKeyCredentials() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("keyCredentials", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPasswordCredentials()))
        for i, v := range m.GetPasswordCredentials() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("passwordCredentials", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *AppManagementConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the keyCredentials property value. 
// Parameters:
//  - value : Value to set for the keyCredentials property.
func (m *AppManagementConfiguration) SetKeyCredentials(value []KeyCredentialConfiguration)() {
    m.keyCredentials = value
}
// Sets the passwordCredentials property value. Collection of password restrictions settings to be applied to an application or service principal
// Parameters:
//  - value : Value to set for the passwordCredentials property.
func (m *AppManagementConfiguration) SetPasswordCredentials(value []PasswordCredentialConfiguration)() {
    m.passwordCredentials = value
}