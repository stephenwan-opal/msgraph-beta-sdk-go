package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ServiceHealthIssue struct {
    ServiceAnnouncementBase
    classification *ServiceHealthClassificationType;
    feature *string;
    featureGroup *string;
    impactDescription *string;
    isResolved *bool;
    origin *ServiceHealthOrigin;
    posts []ServiceHealthIssuePost;
    service *string;
    status *ServiceHealthStatus;
}
func NewServiceHealthIssue()(*ServiceHealthIssue) {
    m := &ServiceHealthIssue{
        ServiceAnnouncementBase: *NewServiceAnnouncementBase(),
    }
    return m
}
func (m *ServiceHealthIssue) GetClassification()(*ServiceHealthClassificationType) {
    if m == nil {
        return nil
    } else {
        return m.classification
    }
}
func (m *ServiceHealthIssue) GetFeature()(*string) {
    if m == nil {
        return nil
    } else {
        return m.feature
    }
}
func (m *ServiceHealthIssue) GetFeatureGroup()(*string) {
    if m == nil {
        return nil
    } else {
        return m.featureGroup
    }
}
func (m *ServiceHealthIssue) GetImpactDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.impactDescription
    }
}
func (m *ServiceHealthIssue) GetIsResolved()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isResolved
    }
}
func (m *ServiceHealthIssue) GetOrigin()(*ServiceHealthOrigin) {
    if m == nil {
        return nil
    } else {
        return m.origin
    }
}
func (m *ServiceHealthIssue) GetPosts()([]ServiceHealthIssuePost) {
    if m == nil {
        return nil
    } else {
        return m.posts
    }
}
func (m *ServiceHealthIssue) GetService()(*string) {
    if m == nil {
        return nil
    } else {
        return m.service
    }
}
func (m *ServiceHealthIssue) GetStatus()(*ServiceHealthStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *ServiceHealthIssue) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ServiceAnnouncementBase.GetFieldDeserializers()
    res["classification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceHealthClassificationType)
        if err != nil {
            return err
        }
        cast := val.(ServiceHealthClassificationType)
        m.SetClassification(&cast)
        return nil
    }
    res["feature"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFeature(val)
        return nil
    }
    res["featureGroup"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFeatureGroup(val)
        return nil
    }
    res["impactDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetImpactDescription(val)
        return nil
    }
    res["isResolved"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsResolved(val)
        return nil
    }
    res["origin"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceHealthOrigin)
        if err != nil {
            return err
        }
        cast := val.(ServiceHealthOrigin)
        m.SetOrigin(&cast)
        return nil
    }
    res["posts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServiceHealthIssuePost() })
        if err != nil {
            return err
        }
        res := make([]ServiceHealthIssuePost, len(val))
        for i, v := range val {
            res[i] = *(v.(*ServiceHealthIssuePost))
        }
        m.SetPosts(res)
        return nil
    }
    res["service"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetService(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceHealthStatus)
        if err != nil {
            return err
        }
        cast := val.(ServiceHealthStatus)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *ServiceHealthIssue) IsNil()(bool) {
    return m == nil
}
func (m *ServiceHealthIssue) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ServiceAnnouncementBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassification() != nil {
        cast := m.GetClassification().String()
        err = writer.WriteStringValue("classification", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("feature", m.GetFeature())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("featureGroup", m.GetFeatureGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("impactDescription", m.GetImpactDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isResolved", m.GetIsResolved())
        if err != nil {
            return err
        }
    }
    if m.GetOrigin() != nil {
        cast := m.GetOrigin().String()
        err = writer.WriteStringValue("origin", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPosts()))
        for i, v := range m.GetPosts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("posts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("service", m.GetService())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ServiceHealthIssue) SetClassification(value *ServiceHealthClassificationType)() {
    m.classification = value
}
func (m *ServiceHealthIssue) SetFeature(value *string)() {
    m.feature = value
}
func (m *ServiceHealthIssue) SetFeatureGroup(value *string)() {
    m.featureGroup = value
}
func (m *ServiceHealthIssue) SetImpactDescription(value *string)() {
    m.impactDescription = value
}
func (m *ServiceHealthIssue) SetIsResolved(value *bool)() {
    m.isResolved = value
}
func (m *ServiceHealthIssue) SetOrigin(value *ServiceHealthOrigin)() {
    m.origin = value
}
func (m *ServiceHealthIssue) SetPosts(value []ServiceHealthIssuePost)() {
    m.posts = value
}
func (m *ServiceHealthIssue) SetService(value *string)() {
    m.service = value
}
func (m *ServiceHealthIssue) SetStatus(value *ServiceHealthStatus)() {
    m.status = value
}