package negbinom_dist

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type NegBinom_DistRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    cumulative *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json;
    // 
    numberF *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json;
    // 
    numberS *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json;
    // 
    probabilityS *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json;
}
// Instantiates a new negBinom_DistRequestBody and sets the default values.
func NewNegBinom_DistRequestBody()(*NegBinom_DistRequestBody) {
    m := &NegBinom_DistRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NegBinom_DistRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the cumulative property value. 
func (m *NegBinom_DistRequestBody) GetCumulative()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json) {
    if m == nil {
        return nil
    } else {
        return m.cumulative
    }
}
// Gets the numberF property value. 
func (m *NegBinom_DistRequestBody) GetNumberF()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json) {
    if m == nil {
        return nil
    } else {
        return m.numberF
    }
}
// Gets the numberS property value. 
func (m *NegBinom_DistRequestBody) GetNumberS()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json) {
    if m == nil {
        return nil
    } else {
        return m.numberS
    }
}
// Gets the probabilityS property value. 
func (m *NegBinom_DistRequestBody) GetProbabilityS()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json) {
    if m == nil {
        return nil
    } else {
        return m.probabilityS
    }
}
// The deserialization information for the current model
func (m *NegBinom_DistRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["cumulative"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCumulative(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json))
        }
        return nil
    }
    res["numberF"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberF(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json))
        }
        return nil
    }
    res["numberS"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberS(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json))
        }
        return nil
    }
    res["probabilityS"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewJson() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProbabilityS(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json))
        }
        return nil
    }
    return res
}
func (m *NegBinom_DistRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *NegBinom_DistRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("cumulative", m.GetCumulative())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("numberF", m.GetNumberF())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("numberS", m.GetNumberS())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("probabilityS", m.GetProbabilityS())
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
func (m *NegBinom_DistRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the cumulative property value. 
// Parameters:
//  - value : Value to set for the cumulative property.
func (m *NegBinom_DistRequestBody) SetCumulative(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json)() {
    m.cumulative = value
}
// Sets the numberF property value. 
// Parameters:
//  - value : Value to set for the numberF property.
func (m *NegBinom_DistRequestBody) SetNumberF(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json)() {
    m.numberF = value
}
// Sets the numberS property value. 
// Parameters:
//  - value : Value to set for the numberS property.
func (m *NegBinom_DistRequestBody) SetNumberS(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json)() {
    m.numberS = value
}
// Sets the probabilityS property value. 
// Parameters:
//  - value : Value to set for the probabilityS property.
func (m *NegBinom_DistRequestBody) SetProbabilityS(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Json)() {
    m.probabilityS = value
}
