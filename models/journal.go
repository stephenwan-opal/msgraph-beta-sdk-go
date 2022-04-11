package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Journal 
type Journal struct {
    Entity
    // The account property
    account Accountable;
    // The balancingAccountId property
    balancingAccountId *string;
    // The balancingAccountNumber property
    balancingAccountNumber *string;
    // The code property
    code *string;
    // The displayName property
    displayName *string;
    // The journalLines property
    journalLines []JournalLineable;
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewJournal instantiates a new journal and sets the default values.
func NewJournal()(*Journal) {
    m := &Journal{
        Entity: *NewEntity(),
    }
    return m
}
// CreateJournalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJournalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewJournal(), nil
}
// GetAccount gets the account property value. The account property
func (m *Journal) GetAccount()(Accountable) {
    if m == nil {
        return nil
    } else {
        return m.account
    }
}
// GetBalancingAccountId gets the balancingAccountId property value. The balancingAccountId property
func (m *Journal) GetBalancingAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.balancingAccountId
    }
}
// GetBalancingAccountNumber gets the balancingAccountNumber property value. The balancingAccountNumber property
func (m *Journal) GetBalancingAccountNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.balancingAccountNumber
    }
}
// GetCode gets the code property value. The code property
func (m *Journal) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *Journal) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Journal) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["account"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccount(val.(Accountable))
        }
        return nil
    }
    res["balancingAccountId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBalancingAccountId(val)
        }
        return nil
    }
    res["balancingAccountNumber"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBalancingAccountNumber(val)
        }
        return nil
    }
    res["code"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["journalLines"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJournalLineFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JournalLineable, len(val))
            for i, v := range val {
                res[i] = v.(JournalLineable)
            }
            m.SetJournalLines(res)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    return res
}
// GetJournalLines gets the journalLines property value. The journalLines property
func (m *Journal) GetJournalLines()([]JournalLineable) {
    if m == nil {
        return nil
    } else {
        return m.journalLines
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Journal) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Serialize serializes information the current object
func (m *Journal) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("account", m.GetAccount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("balancingAccountId", m.GetBalancingAccountId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("balancingAccountNumber", m.GetBalancingAccountNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("code", m.GetCode())
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
    if m.GetJournalLines() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetJournalLines()))
        for i, v := range m.GetJournalLines() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("journalLines", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccount sets the account property value. The account property
func (m *Journal) SetAccount(value Accountable)() {
    if m != nil {
        m.account = value
    }
}
// SetBalancingAccountId sets the balancingAccountId property value. The balancingAccountId property
func (m *Journal) SetBalancingAccountId(value *string)() {
    if m != nil {
        m.balancingAccountId = value
    }
}
// SetBalancingAccountNumber sets the balancingAccountNumber property value. The balancingAccountNumber property
func (m *Journal) SetBalancingAccountNumber(value *string)() {
    if m != nil {
        m.balancingAccountNumber = value
    }
}
// SetCode sets the code property value. The code property
func (m *Journal) SetCode(value *string)() {
    if m != nil {
        m.code = value
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Journal) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetJournalLines sets the journalLines property value. The journalLines property
func (m *Journal) SetJournalLines(value []JournalLineable)() {
    if m != nil {
        m.journalLines = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Journal) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}