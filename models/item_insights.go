package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemInsights casts the previous resource to user.
type ItemInsights struct {
    OfficeGraphInsights
}
// NewItemInsights instantiates a new itemInsights and sets the default values.
func NewItemInsights()(*ItemInsights) {
    m := &ItemInsights{
        OfficeGraphInsights: *NewOfficeGraphInsights(),
    }
    return m
}
// CreateItemInsightsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemInsightsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemInsights(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemInsights) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OfficeGraphInsights.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ItemInsights) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OfficeGraphInsights.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
