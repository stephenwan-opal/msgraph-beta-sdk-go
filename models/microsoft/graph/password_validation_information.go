package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PasswordValidationInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies whether the password is valid based on the calculation of the results in the validationResults property. Not nullable. Read-only.
    isValid *bool;
    // The list of password validation rules and whether the password passed those rules. Not nullable. Read-only.
    validationResults []ValidationResult;
}
// Instantiates a new passwordValidationInformation and sets the default values.
func NewPasswordValidationInformation()(*PasswordValidationInformation) {
    m := &PasswordValidationInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PasswordValidationInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the isValid property value. Specifies whether the password is valid based on the calculation of the results in the validationResults property. Not nullable. Read-only.
func (m *PasswordValidationInformation) GetIsValid()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isValid
    }
}
// Gets the validationResults property value. The list of password validation rules and whether the password passed those rules. Not nullable. Read-only.
func (m *PasswordValidationInformation) GetValidationResults()([]ValidationResult) {
    if m == nil {
        return nil
    } else {
        return m.validationResults
    }
}
// The deserialization information for the current model
func (m *PasswordValidationInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isValid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsValid(val)
        }
        return nil
    }
    res["validationResults"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewValidationResult() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ValidationResult, len(val))
            for i, v := range val {
                res[i] = *(v.(*ValidationResult))
            }
            m.SetValidationResults(res)
        }
        return nil
    }
    return res
}
func (m *PasswordValidationInformation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PasswordValidationInformation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isValid", m.GetIsValid())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetValidationResults()))
        for i, v := range m.GetValidationResults() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("validationResults", cast)
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
func (m *PasswordValidationInformation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the isValid property value. Specifies whether the password is valid based on the calculation of the results in the validationResults property. Not nullable. Read-only.
// Parameters:
//  - value : Value to set for the isValid property.
func (m *PasswordValidationInformation) SetIsValid(value *bool)() {
    m.isValid = value
}
// Sets the validationResults property value. The list of password validation rules and whether the password passed those rules. Not nullable. Read-only.
// Parameters:
//  - value : Value to set for the validationResults property.
func (m *PasswordValidationInformation) SetValidationResults(value []ValidationResult)() {
    m.validationResults = value
}
