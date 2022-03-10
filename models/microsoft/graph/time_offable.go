package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TimeOffable 
type TimeOffable interface {
    ChangeTrackedEntityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDraftTimeOff()(TimeOffItemable)
    GetIsStagedForDeletion()(*bool)
    GetSharedTimeOff()(TimeOffItemable)
    GetUserId()(*string)
    SetDraftTimeOff(value TimeOffItemable)()
    SetIsStagedForDeletion(value *bool)()
    SetSharedTimeOff(value TimeOffItemable)()
    SetUserId(value *string)()
}