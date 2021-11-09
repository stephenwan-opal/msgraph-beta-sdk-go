package findmeetingtimes

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type FindMeetingTimesRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    attendees []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AttendeeBase;
    // 
    isOrganizerOptional *bool;
    // 
    locationConstraint *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.LocationConstraint;
    // 
    maxCandidates *int32;
    // 
    meetingDuration *string;
    // 
    minimumAttendeePercentage *float64;
    // 
    returnSuggestionReasons *bool;
    // 
    timeConstraint *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TimeConstraint;
}
// Instantiates a new findMeetingTimesRequestBody and sets the default values.
func NewFindMeetingTimesRequestBody()(*FindMeetingTimesRequestBody) {
    m := &FindMeetingTimesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FindMeetingTimesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the attendees property value. 
func (m *FindMeetingTimesRequestBody) GetAttendees()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AttendeeBase) {
    if m == nil {
        return nil
    } else {
        return m.attendees
    }
}
// Gets the isOrganizerOptional property value. 
func (m *FindMeetingTimesRequestBody) GetIsOrganizerOptional()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOrganizerOptional
    }
}
// Gets the locationConstraint property value. 
func (m *FindMeetingTimesRequestBody) GetLocationConstraint()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.LocationConstraint) {
    if m == nil {
        return nil
    } else {
        return m.locationConstraint
    }
}
// Gets the maxCandidates property value. 
func (m *FindMeetingTimesRequestBody) GetMaxCandidates()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxCandidates
    }
}
// Gets the meetingDuration property value. 
func (m *FindMeetingTimesRequestBody) GetMeetingDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.meetingDuration
    }
}
// Gets the minimumAttendeePercentage property value. 
func (m *FindMeetingTimesRequestBody) GetMinimumAttendeePercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.minimumAttendeePercentage
    }
}
// Gets the returnSuggestionReasons property value. 
func (m *FindMeetingTimesRequestBody) GetReturnSuggestionReasons()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.returnSuggestionReasons
    }
}
// Gets the timeConstraint property value. 
func (m *FindMeetingTimesRequestBody) GetTimeConstraint()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TimeConstraint) {
    if m == nil {
        return nil
    } else {
        return m.timeConstraint
    }
}
// The deserialization information for the current model
func (m *FindMeetingTimesRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attendees"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAttendeeBase() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AttendeeBase, len(val))
            for i, v := range val {
                res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AttendeeBase))
            }
            m.SetAttendees(res)
        }
        return nil
    }
    res["isOrganizerOptional"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOrganizerOptional(val)
        }
        return nil
    }
    res["locationConstraint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewLocationConstraint() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocationConstraint(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.LocationConstraint))
        }
        return nil
    }
    res["maxCandidates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxCandidates(val)
        }
        return nil
    }
    res["meetingDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingDuration(val)
        }
        return nil
    }
    res["minimumAttendeePercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumAttendeePercentage(val)
        }
        return nil
    }
    res["returnSuggestionReasons"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReturnSuggestionReasons(val)
        }
        return nil
    }
    res["timeConstraint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTimeConstraint() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeConstraint(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TimeConstraint))
        }
        return nil
    }
    return res
}
func (m *FindMeetingTimesRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *FindMeetingTimesRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttendees()))
        for i, v := range m.GetAttendees() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("attendees", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isOrganizerOptional", m.GetIsOrganizerOptional())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("locationConstraint", m.GetLocationConstraint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxCandidates", m.GetMaxCandidates())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("meetingDuration", m.GetMeetingDuration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("minimumAttendeePercentage", m.GetMinimumAttendeePercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("returnSuggestionReasons", m.GetReturnSuggestionReasons())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("timeConstraint", m.GetTimeConstraint())
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
func (m *FindMeetingTimesRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the attendees property value. 
// Parameters:
//  - value : Value to set for the attendees property.
func (m *FindMeetingTimesRequestBody) SetAttendees(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AttendeeBase)() {
    m.attendees = value
}
// Sets the isOrganizerOptional property value. 
// Parameters:
//  - value : Value to set for the isOrganizerOptional property.
func (m *FindMeetingTimesRequestBody) SetIsOrganizerOptional(value *bool)() {
    m.isOrganizerOptional = value
}
// Sets the locationConstraint property value. 
// Parameters:
//  - value : Value to set for the locationConstraint property.
func (m *FindMeetingTimesRequestBody) SetLocationConstraint(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.LocationConstraint)() {
    m.locationConstraint = value
}
// Sets the maxCandidates property value. 
// Parameters:
//  - value : Value to set for the maxCandidates property.
func (m *FindMeetingTimesRequestBody) SetMaxCandidates(value *int32)() {
    m.maxCandidates = value
}
// Sets the meetingDuration property value. 
// Parameters:
//  - value : Value to set for the meetingDuration property.
func (m *FindMeetingTimesRequestBody) SetMeetingDuration(value *string)() {
    m.meetingDuration = value
}
// Sets the minimumAttendeePercentage property value. 
// Parameters:
//  - value : Value to set for the minimumAttendeePercentage property.
func (m *FindMeetingTimesRequestBody) SetMinimumAttendeePercentage(value *float64)() {
    m.minimumAttendeePercentage = value
}
// Sets the returnSuggestionReasons property value. 
// Parameters:
//  - value : Value to set for the returnSuggestionReasons property.
func (m *FindMeetingTimesRequestBody) SetReturnSuggestionReasons(value *bool)() {
    m.returnSuggestionReasons = value
}
// Sets the timeConstraint property value. 
// Parameters:
//  - value : Value to set for the timeConstraint property.
func (m *FindMeetingTimesRequestBody) SetTimeConstraint(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TimeConstraint)() {
    m.timeConstraint = value
}
