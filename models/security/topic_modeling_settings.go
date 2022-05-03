package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TopicModelingSettings 
type TopicModelingSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The dynamicallyAdjustTopicCount property
    dynamicallyAdjustTopicCount *bool
    // The ignoreNumbers property
    ignoreNumbers *bool
    // The isEnabled property
    isEnabled *bool
    // The topicCount property
    topicCount *int32
}
// NewTopicModelingSettings instantiates a new topicModelingSettings and sets the default values.
func NewTopicModelingSettings()(*TopicModelingSettings) {
    m := &TopicModelingSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTopicModelingSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTopicModelingSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTopicModelingSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TopicModelingSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDynamicallyAdjustTopicCount gets the dynamicallyAdjustTopicCount property value. The dynamicallyAdjustTopicCount property
func (m *TopicModelingSettings) GetDynamicallyAdjustTopicCount()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.dynamicallyAdjustTopicCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TopicModelingSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dynamicallyAdjustTopicCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDynamicallyAdjustTopicCount(val)
        }
        return nil
    }
    res["ignoreNumbers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIgnoreNumbers(val)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["topicCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopicCount(val)
        }
        return nil
    }
    return res
}
// GetIgnoreNumbers gets the ignoreNumbers property value. The ignoreNumbers property
func (m *TopicModelingSettings) GetIgnoreNumbers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.ignoreNumbers
    }
}
// GetIsEnabled gets the isEnabled property value. The isEnabled property
func (m *TopicModelingSettings) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetTopicCount gets the topicCount property value. The topicCount property
func (m *TopicModelingSettings) GetTopicCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.topicCount
    }
}
// Serialize serializes information the current object
func (m *TopicModelingSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("dynamicallyAdjustTopicCount", m.GetDynamicallyAdjustTopicCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ignoreNumbers", m.GetIgnoreNumbers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("topicCount", m.GetTopicCount())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TopicModelingSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDynamicallyAdjustTopicCount sets the dynamicallyAdjustTopicCount property value. The dynamicallyAdjustTopicCount property
func (m *TopicModelingSettings) SetDynamicallyAdjustTopicCount(value *bool)() {
    if m != nil {
        m.dynamicallyAdjustTopicCount = value
    }
}
// SetIgnoreNumbers sets the ignoreNumbers property value. The ignoreNumbers property
func (m *TopicModelingSettings) SetIgnoreNumbers(value *bool)() {
    if m != nil {
        m.ignoreNumbers = value
    }
}
// SetIsEnabled sets the isEnabled property value. The isEnabled property
func (m *TopicModelingSettings) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
// SetTopicCount sets the topicCount property value. The topicCount property
func (m *TopicModelingSettings) SetTopicCount(value *int32)() {
    if m != nil {
        m.topicCount = value
    }
}
