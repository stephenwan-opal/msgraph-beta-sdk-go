package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingService represents a particular service offered by a booking business.
type BookingService struct {
    BookingNamedEntity
    // Additional information that is sent to the customer when an appointment is confirmed.
    additionalInformation *string
    // Contains the set of custom questions associated with a particular service.
    customQuestions []BookingQuestionAssignmentable
    // The default length of the service, represented in numbers of days, hours, minutes, and seconds. For example, P11D23H59M59.999999999999S.
    defaultDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The default physical location for the service.
    defaultLocation Locationable
    // The default monetary price for the service.
    defaultPrice *float64
    // The default way the service is charged. Possible values are: undefined, fixedPrice, startingAt, hourly, free, priceVaries, callUs, notSet, unknownFutureValue.
    defaultPriceType *BookingPriceType
    // The default set of reminders for an appointment of this service. The value of this property is available only when reading this bookingService by its ID.
    defaultReminders []BookingReminderable
    // A text description for the service.
    description *string
    // The isAnonymousJoinEnabled property
    isAnonymousJoinEnabled *bool
    // True means this service is not available to customers for booking.
    isHiddenFromCustomers *bool
    // True indicates that the appointments for the service will be held online. Default value is false.
    isLocationOnline *bool
    // The languageTag property
    languageTag *string
    // The maximum number of customers allowed in a service. If maximumAttendeesCount of the service is greater than 1, pass valid customer IDs while creating or updating an appointment.  To create a customer, use the Create bookingCustomer operation.
    maximumAttendeesCount *int32
    // Additional information about this service.
    notes *string
    // The time to buffer after an appointment for this service ends, and before the next customer appointment can be booked.
    postBuffer *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The time to buffer before an appointment for this service can start.
    preBuffer *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The set of policies that determine how appointments for this type of service should be created and managed.
    schedulingPolicy BookingSchedulingPolicyable
    // True indicates SMS notifications can be sent to the customers for the appointment of the service. Default value is false.
    smsNotificationsEnabled *bool
    // Represents those staff members who provide this service.
    staffMemberIds []string
    // The URL a customer uses to access the service.
    webUrl *string
}
// NewBookingService instantiates a new bookingService and sets the default values.
func NewBookingService()(*BookingService) {
    m := &BookingService{
        BookingNamedEntity: *NewBookingNamedEntity(),
    }
    return m
}
// CreateBookingServiceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingServiceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingService(), nil
}
// GetAdditionalInformation gets the additionalInformation property value. Additional information that is sent to the customer when an appointment is confirmed.
func (m *BookingService) GetAdditionalInformation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.additionalInformation
    }
}
// GetCustomQuestions gets the customQuestions property value. Contains the set of custom questions associated with a particular service.
func (m *BookingService) GetCustomQuestions()([]BookingQuestionAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.customQuestions
    }
}
// GetDefaultDuration gets the defaultDuration property value. The default length of the service, represented in numbers of days, hours, minutes, and seconds. For example, P11D23H59M59.999999999999S.
func (m *BookingService) GetDefaultDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.defaultDuration
    }
}
// GetDefaultLocation gets the defaultLocation property value. The default physical location for the service.
func (m *BookingService) GetDefaultLocation()(Locationable) {
    if m == nil {
        return nil
    } else {
        return m.defaultLocation
    }
}
// GetDefaultPrice gets the defaultPrice property value. The default monetary price for the service.
func (m *BookingService) GetDefaultPrice()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.defaultPrice
    }
}
// GetDefaultPriceType gets the defaultPriceType property value. The default way the service is charged. Possible values are: undefined, fixedPrice, startingAt, hourly, free, priceVaries, callUs, notSet, unknownFutureValue.
func (m *BookingService) GetDefaultPriceType()(*BookingPriceType) {
    if m == nil {
        return nil
    } else {
        return m.defaultPriceType
    }
}
// GetDefaultReminders gets the defaultReminders property value. The default set of reminders for an appointment of this service. The value of this property is available only when reading this bookingService by its ID.
func (m *BookingService) GetDefaultReminders()([]BookingReminderable) {
    if m == nil {
        return nil
    } else {
        return m.defaultReminders
    }
}
// GetDescription gets the description property value. A text description for the service.
func (m *BookingService) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingService) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BookingNamedEntity.GetFieldDeserializers()
    res["additionalInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalInformation(val)
        }
        return nil
    }
    res["customQuestions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingQuestionAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingQuestionAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(BookingQuestionAssignmentable)
            }
            m.SetCustomQuestions(res)
        }
        return nil
    }
    res["defaultDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultDuration(val)
        }
        return nil
    }
    res["defaultLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLocation(val.(Locationable))
        }
        return nil
    }
    res["defaultPrice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultPrice(val)
        }
        return nil
    }
    res["defaultPriceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseBookingPriceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultPriceType(val.(*BookingPriceType))
        }
        return nil
    }
    res["defaultReminders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingReminderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingReminderable, len(val))
            for i, v := range val {
                res[i] = v.(BookingReminderable)
            }
            m.SetDefaultReminders(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["isAnonymousJoinEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAnonymousJoinEnabled(val)
        }
        return nil
    }
    res["isHiddenFromCustomers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHiddenFromCustomers(val)
        }
        return nil
    }
    res["isLocationOnline"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsLocationOnline(val)
        }
        return nil
    }
    res["languageTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguageTag(val)
        }
        return nil
    }
    res["maximumAttendeesCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumAttendeesCount(val)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["postBuffer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostBuffer(val)
        }
        return nil
    }
    res["preBuffer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreBuffer(val)
        }
        return nil
    }
    res["schedulingPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBookingSchedulingPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedulingPolicy(val.(BookingSchedulingPolicyable))
        }
        return nil
    }
    res["smsNotificationsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmsNotificationsEnabled(val)
        }
        return nil
    }
    res["staffMemberIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetStaffMemberIds(res)
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
// GetIsAnonymousJoinEnabled gets the isAnonymousJoinEnabled property value. The isAnonymousJoinEnabled property
func (m *BookingService) GetIsAnonymousJoinEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAnonymousJoinEnabled
    }
}
// GetIsHiddenFromCustomers gets the isHiddenFromCustomers property value. True means this service is not available to customers for booking.
func (m *BookingService) GetIsHiddenFromCustomers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isHiddenFromCustomers
    }
}
// GetIsLocationOnline gets the isLocationOnline property value. True indicates that the appointments for the service will be held online. Default value is false.
func (m *BookingService) GetIsLocationOnline()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isLocationOnline
    }
}
// GetLanguageTag gets the languageTag property value. The languageTag property
func (m *BookingService) GetLanguageTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.languageTag
    }
}
// GetMaximumAttendeesCount gets the maximumAttendeesCount property value. The maximum number of customers allowed in a service. If maximumAttendeesCount of the service is greater than 1, pass valid customer IDs while creating or updating an appointment.  To create a customer, use the Create bookingCustomer operation.
func (m *BookingService) GetMaximumAttendeesCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maximumAttendeesCount
    }
}
// GetNotes gets the notes property value. Additional information about this service.
func (m *BookingService) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// GetPostBuffer gets the postBuffer property value. The time to buffer after an appointment for this service ends, and before the next customer appointment can be booked.
func (m *BookingService) GetPostBuffer()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.postBuffer
    }
}
// GetPreBuffer gets the preBuffer property value. The time to buffer before an appointment for this service can start.
func (m *BookingService) GetPreBuffer()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.preBuffer
    }
}
// GetSchedulingPolicy gets the schedulingPolicy property value. The set of policies that determine how appointments for this type of service should be created and managed.
func (m *BookingService) GetSchedulingPolicy()(BookingSchedulingPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.schedulingPolicy
    }
}
// GetSmsNotificationsEnabled gets the smsNotificationsEnabled property value. True indicates SMS notifications can be sent to the customers for the appointment of the service. Default value is false.
func (m *BookingService) GetSmsNotificationsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.smsNotificationsEnabled
    }
}
// GetStaffMemberIds gets the staffMemberIds property value. Represents those staff members who provide this service.
func (m *BookingService) GetStaffMemberIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.staffMemberIds
    }
}
// GetWebUrl gets the webUrl property value. The URL a customer uses to access the service.
func (m *BookingService) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// Serialize serializes information the current object
func (m *BookingService) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BookingNamedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("additionalInformation", m.GetAdditionalInformation())
        if err != nil {
            return err
        }
    }
    if m.GetCustomQuestions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomQuestions()))
        for i, v := range m.GetCustomQuestions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customQuestions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("defaultDuration", m.GetDefaultDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defaultLocation", m.GetDefaultLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("defaultPrice", m.GetDefaultPrice())
        if err != nil {
            return err
        }
    }
    if m.GetDefaultPriceType() != nil {
        cast := (*m.GetDefaultPriceType()).String()
        err = writer.WriteStringValue("defaultPriceType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefaultReminders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDefaultReminders()))
        for i, v := range m.GetDefaultReminders() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("defaultReminders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAnonymousJoinEnabled", m.GetIsAnonymousJoinEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isHiddenFromCustomers", m.GetIsHiddenFromCustomers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isLocationOnline", m.GetIsLocationOnline())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("languageTag", m.GetLanguageTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumAttendeesCount", m.GetMaximumAttendeesCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("postBuffer", m.GetPostBuffer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("preBuffer", m.GetPreBuffer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schedulingPolicy", m.GetSchedulingPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smsNotificationsEnabled", m.GetSmsNotificationsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetStaffMemberIds() != nil {
        err = writer.WriteCollectionOfStringValues("staffMemberIds", m.GetStaffMemberIds())
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
// SetAdditionalInformation sets the additionalInformation property value. Additional information that is sent to the customer when an appointment is confirmed.
func (m *BookingService) SetAdditionalInformation(value *string)() {
    if m != nil {
        m.additionalInformation = value
    }
}
// SetCustomQuestions sets the customQuestions property value. Contains the set of custom questions associated with a particular service.
func (m *BookingService) SetCustomQuestions(value []BookingQuestionAssignmentable)() {
    if m != nil {
        m.customQuestions = value
    }
}
// SetDefaultDuration sets the defaultDuration property value. The default length of the service, represented in numbers of days, hours, minutes, and seconds. For example, P11D23H59M59.999999999999S.
func (m *BookingService) SetDefaultDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    if m != nil {
        m.defaultDuration = value
    }
}
// SetDefaultLocation sets the defaultLocation property value. The default physical location for the service.
func (m *BookingService) SetDefaultLocation(value Locationable)() {
    if m != nil {
        m.defaultLocation = value
    }
}
// SetDefaultPrice sets the defaultPrice property value. The default monetary price for the service.
func (m *BookingService) SetDefaultPrice(value *float64)() {
    if m != nil {
        m.defaultPrice = value
    }
}
// SetDefaultPriceType sets the defaultPriceType property value. The default way the service is charged. Possible values are: undefined, fixedPrice, startingAt, hourly, free, priceVaries, callUs, notSet, unknownFutureValue.
func (m *BookingService) SetDefaultPriceType(value *BookingPriceType)() {
    if m != nil {
        m.defaultPriceType = value
    }
}
// SetDefaultReminders sets the defaultReminders property value. The default set of reminders for an appointment of this service. The value of this property is available only when reading this bookingService by its ID.
func (m *BookingService) SetDefaultReminders(value []BookingReminderable)() {
    if m != nil {
        m.defaultReminders = value
    }
}
// SetDescription sets the description property value. A text description for the service.
func (m *BookingService) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetIsAnonymousJoinEnabled sets the isAnonymousJoinEnabled property value. The isAnonymousJoinEnabled property
func (m *BookingService) SetIsAnonymousJoinEnabled(value *bool)() {
    if m != nil {
        m.isAnonymousJoinEnabled = value
    }
}
// SetIsHiddenFromCustomers sets the isHiddenFromCustomers property value. True means this service is not available to customers for booking.
func (m *BookingService) SetIsHiddenFromCustomers(value *bool)() {
    if m != nil {
        m.isHiddenFromCustomers = value
    }
}
// SetIsLocationOnline sets the isLocationOnline property value. True indicates that the appointments for the service will be held online. Default value is false.
func (m *BookingService) SetIsLocationOnline(value *bool)() {
    if m != nil {
        m.isLocationOnline = value
    }
}
// SetLanguageTag sets the languageTag property value. The languageTag property
func (m *BookingService) SetLanguageTag(value *string)() {
    if m != nil {
        m.languageTag = value
    }
}
// SetMaximumAttendeesCount sets the maximumAttendeesCount property value. The maximum number of customers allowed in a service. If maximumAttendeesCount of the service is greater than 1, pass valid customer IDs while creating or updating an appointment.  To create a customer, use the Create bookingCustomer operation.
func (m *BookingService) SetMaximumAttendeesCount(value *int32)() {
    if m != nil {
        m.maximumAttendeesCount = value
    }
}
// SetNotes sets the notes property value. Additional information about this service.
func (m *BookingService) SetNotes(value *string)() {
    if m != nil {
        m.notes = value
    }
}
// SetPostBuffer sets the postBuffer property value. The time to buffer after an appointment for this service ends, and before the next customer appointment can be booked.
func (m *BookingService) SetPostBuffer(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    if m != nil {
        m.postBuffer = value
    }
}
// SetPreBuffer sets the preBuffer property value. The time to buffer before an appointment for this service can start.
func (m *BookingService) SetPreBuffer(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    if m != nil {
        m.preBuffer = value
    }
}
// SetSchedulingPolicy sets the schedulingPolicy property value. The set of policies that determine how appointments for this type of service should be created and managed.
func (m *BookingService) SetSchedulingPolicy(value BookingSchedulingPolicyable)() {
    if m != nil {
        m.schedulingPolicy = value
    }
}
// SetSmsNotificationsEnabled sets the smsNotificationsEnabled property value. True indicates SMS notifications can be sent to the customers for the appointment of the service. Default value is false.
func (m *BookingService) SetSmsNotificationsEnabled(value *bool)() {
    if m != nil {
        m.smsNotificationsEnabled = value
    }
}
// SetStaffMemberIds sets the staffMemberIds property value. Represents those staff members who provide this service.
func (m *BookingService) SetStaffMemberIds(value []string)() {
    if m != nil {
        m.staffMemberIds = value
    }
}
// SetWebUrl sets the webUrl property value. The URL a customer uses to access the service.
func (m *BookingService) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
