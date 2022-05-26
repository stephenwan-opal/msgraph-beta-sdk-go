package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsTab provides operations to manage the collection of chat entities.
type TeamsTab struct {
    Entity
    // Container for custom settings applied to a tab. The tab is considered configured only once this property is set.
    configuration TeamsTabConfigurationable
    // Name of the tab.
    displayName *string
    // The messageId property
    messageId *string
    // Index of the order used for sorting tabs.
    sortOrderIndex *string
    // The application that is linked to the tab.
    teamsApp TeamsAppable
    // The teamsAppId property
    teamsAppId *string
    // Deep link URL of the tab instance. Read only.
    webUrl *string
}
// NewTeamsTab instantiates a new teamsTab and sets the default values.
func NewTeamsTab()(*TeamsTab) {
    m := &TeamsTab{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamsTabFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsTabFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsTab(), nil
}
// GetConfiguration gets the configuration property value. Container for custom settings applied to a tab. The tab is considered configured only once this property is set.
func (m *TeamsTab) GetConfiguration()(TeamsTabConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.configuration
    }
}
// GetDisplayName gets the displayName property value. Name of the tab.
func (m *TeamsTab) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsTab) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsTabConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguration(val.(TeamsTabConfigurationable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["messageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageId(val)
        }
        return nil
    }
    res["sortOrderIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSortOrderIndex(val)
        }
        return nil
    }
    res["teamsApp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsApp(val.(TeamsAppable))
        }
        return nil
    }
    res["teamsAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsAppId(val)
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetMessageId gets the messageId property value. The messageId property
func (m *TeamsTab) GetMessageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.messageId
    }
}
// GetSortOrderIndex gets the sortOrderIndex property value. Index of the order used for sorting tabs.
func (m *TeamsTab) GetSortOrderIndex()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sortOrderIndex
    }
}
// GetTeamsApp gets the teamsApp property value. The application that is linked to the tab.
func (m *TeamsTab) GetTeamsApp()(TeamsAppable) {
    if m == nil {
        return nil
    } else {
        return m.teamsApp
    }
}
// GetTeamsAppId gets the teamsAppId property value. The teamsAppId property
func (m *TeamsTab) GetTeamsAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.teamsAppId
    }
}
// GetWebUrl gets the webUrl property value. Deep link URL of the tab instance. Read only.
func (m *TeamsTab) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// Serialize serializes information the current object
func (m *TeamsTab) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("configuration", m.GetConfiguration())
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
    {
        err = writer.WriteStringValue("messageId", m.GetMessageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sortOrderIndex", m.GetSortOrderIndex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teamsApp", m.GetTeamsApp())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamsAppId", m.GetTeamsAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfiguration sets the configuration property value. Container for custom settings applied to a tab. The tab is considered configured only once this property is set.
func (m *TeamsTab) SetConfiguration(value TeamsTabConfigurationable)() {
    if m != nil {
        m.configuration = value
    }
}
// SetDisplayName sets the displayName property value. Name of the tab.
func (m *TeamsTab) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetMessageId sets the messageId property value. The messageId property
func (m *TeamsTab) SetMessageId(value *string)() {
    if m != nil {
        m.messageId = value
    }
}
// SetSortOrderIndex sets the sortOrderIndex property value. Index of the order used for sorting tabs.
func (m *TeamsTab) SetSortOrderIndex(value *string)() {
    if m != nil {
        m.sortOrderIndex = value
    }
}
// SetTeamsApp sets the teamsApp property value. The application that is linked to the tab.
func (m *TeamsTab) SetTeamsApp(value TeamsAppable)() {
    if m != nil {
        m.teamsApp = value
    }
}
// SetTeamsAppId sets the teamsAppId property value. The teamsAppId property
func (m *TeamsTab) SetTeamsAppId(value *string)() {
    if m != nil {
        m.teamsAppId = value
    }
}
// SetWebUrl sets the webUrl property value. Deep link URL of the tab instance. Read only.
func (m *TeamsTab) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
