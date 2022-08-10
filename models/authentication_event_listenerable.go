package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationEventListenerable 
type AuthenticationEventListenerable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationEventsFlowId()(*string)
    GetConditions()(AuthenticationConditionsable)
    GetPriority()(*int32)
    GetTags()([]KeyValuePairable)
    SetAuthenticationEventsFlowId(value *string)()
    SetConditions(value AuthenticationConditionsable)()
    SetPriority(value *int32)()
    SetTags(value []KeyValuePairable)()
}