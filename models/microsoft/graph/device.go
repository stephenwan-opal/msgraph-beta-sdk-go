package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Device struct {
    DirectoryObject
    // true if the account is enabled; otherwise, false. Required. Default is true.  Supports $filter (eq, ne, NOT, in). Only callers in Global Administrator and Cloud Device Administrator roles can set this property.
    accountEnabled *bool;
    // For internal use only. Not nullable. Supports $filter (eq, NOT, ge, le).
    alternativeSecurityIds []AlternativeSecurityId;
    // The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter (eq, ne, NOT, ge, le) and $orderBy.
    approximateLastSignInDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Set of commands sent to this device.
    commands []Command;
    // The timestamp when the device is no longer deemed compliant. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    complianceExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // User-defined property set by Intune to automatically add devices to groups and simplify managing devices.
    deviceCategory *string;
    // Unique identifier set by Azure Device Registration Service at the time of registration. Supports $filter (eq, ne, NOT, startsWith).
    deviceId *string;
    // For internal use only. Set to null.
    deviceMetadata *string;
    // Ownership of the device. This property is set by Intune. Possible values are: unknown, company, personal.
    deviceOwnership *string;
    // For internal use only.
    deviceVersion *int32;
    // The display name for the device. Required. Supports $filter (eq, ne, NOT, ge, le, in, startsWith), $search, and $orderBy.
    displayName *string;
    // The on-premises domain name of Hybrid Azure AD joined devices. This property is set by Intune.
    domainName *string;
    // Enrollment profile applied to the device. For example, Apple Device Enrollment Profile, Device enrollment - Corporate device identifiers, or Windows Autopilot profile name. This property is set by Intune.
    enrollmentProfileName *string;
    // Enrollment type of the device. This property is set by Intune. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement.
    enrollmentType *string;
    // Contains extension attributes 1-15 for the device. The individual extension attributes are not selectable. These properties are mastered in cloud and can be set during creation or update of a device object in Azure AD. Supports $filter (eq, NOT, startsWith).
    extensionAttributes *OnPremisesExtensionAttributes;
    // The collection of open extensions defined for the device. Read-only. Nullable.
    extensions []Extension;
    // List of hostNames for the device.
    hostnames []string;
    // true if the device complies with Mobile Device Management (MDM) policies; otherwise, false. Read-only. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, NOT).
    isCompliant *bool;
    // true if the device is managed by a Mobile Device Management (MDM) app; otherwise, false. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, NOT).
    isManaged *bool;
    // true if device is rooted; false if device is jail-broken. This can only be updated by Intune.
    isRooted *bool;
    // Form factor of device. Only returned if user signs in with a Microsoft account as part of Project Rome.
    kind *string;
    // Management channel of the device.  This property is set by Intune. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
    managementType *string;
    // Manufacturer of the device. Read-only.
    manufacturer *string;
    // Groups that this device is a member of. Read-only. Nullable. Supports $expand.
    memberOf []DirectoryObject;
    // Model of the device. Read-only.
    model *string;
    // Friendly name of a device. Only returned if user signs in with a Microsoft account as part of Project Rome.
    name *string;
    // The last time at which the object was synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Read-only. Supports $filter (eq, ne, NOT, ge, le, in).
    onPremisesLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Supports $filter (eq, ne, NOT, in).
    onPremisesSyncEnabled *bool;
    // The type of operating system on the device. Required. Supports $filter (eq, ne, NOT, ge, le, startsWith).
    operatingSystem *string;
    // The version of the operating system on the device. Required. Supports $filter (eq, ne, NOT, ge, le, startsWith).
    operatingSystemVersion *string;
    // For internal use only. Not nullable. Supports $filter (eq, NOT, ge, le, startsWith).
    physicalIds []string;
    // Platform of device. Only returned if user signs in with a Microsoft account as part of Project Rome. Only returned if user signs in with a Microsoft account as part of Project Rome.
    platform *string;
    // The profile type of the device. Possible values: RegisteredDevice (default), SecureVM, Printer, Shared, IoT.
    profileType *string;
    // The user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
    registeredOwners []DirectoryObject;
    // Collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
    registeredUsers []DirectoryObject;
    // Date and time of when the device was registered. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    registrationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Device is online or offline. Only returned if user signs in with a Microsoft account as part of Project Rome.
    status *string;
    // List of labels applied to the device by the system.
    systemLabels []string;
    // Groups that the device is a member of. This operation is transitive. Supports $expand.
    transitiveMemberOf []DirectoryObject;
    // Type of trust for the joined device. Read-only. Possible values:  Workplace (indicates bring your own personal devices), AzureAd (Cloud only joined devices), ServerAd (on-premises domain joined devices joined to Azure AD). For more details, see Introduction to device management in Azure Active Directory
    trustType *string;
    // Represents the usage rights a device has been granted.
    usageRights []UsageRight;
}
// Instantiates a new device and sets the default values.
func NewDevice()(*Device) {
    m := &Device{
        DirectoryObject: *NewDirectoryObject(),
    }
    return m
}
// Gets the accountEnabled property value. true if the account is enabled; otherwise, false. Required. Default is true.  Supports $filter (eq, ne, NOT, in). Only callers in Global Administrator and Cloud Device Administrator roles can set this property.
func (m *Device) GetAccountEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.accountEnabled
    }
}
// Gets the alternativeSecurityIds property value. For internal use only. Not nullable. Supports $filter (eq, NOT, ge, le).
func (m *Device) GetAlternativeSecurityIds()([]AlternativeSecurityId) {
    if m == nil {
        return nil
    } else {
        return m.alternativeSecurityIds
    }
}
// Gets the approximateLastSignInDateTime property value. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter (eq, ne, NOT, ge, le) and $orderBy.
func (m *Device) GetApproximateLastSignInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.approximateLastSignInDateTime
    }
}
// Gets the commands property value. Set of commands sent to this device.
func (m *Device) GetCommands()([]Command) {
    if m == nil {
        return nil
    } else {
        return m.commands
    }
}
// Gets the complianceExpirationDateTime property value. The timestamp when the device is no longer deemed compliant. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Device) GetComplianceExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.complianceExpirationDateTime
    }
}
// Gets the deviceCategory property value. User-defined property set by Intune to automatically add devices to groups and simplify managing devices.
func (m *Device) GetDeviceCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceCategory
    }
}
// Gets the deviceId property value. Unique identifier set by Azure Device Registration Service at the time of registration. Supports $filter (eq, ne, NOT, startsWith).
func (m *Device) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the deviceMetadata property value. For internal use only. Set to null.
func (m *Device) GetDeviceMetadata()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceMetadata
    }
}
// Gets the deviceOwnership property value. Ownership of the device. This property is set by Intune. Possible values are: unknown, company, personal.
func (m *Device) GetDeviceOwnership()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceOwnership
    }
}
// Gets the deviceVersion property value. For internal use only.
func (m *Device) GetDeviceVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceVersion
    }
}
// Gets the displayName property value. The display name for the device. Required. Supports $filter (eq, ne, NOT, ge, le, in, startsWith), $search, and $orderBy.
func (m *Device) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the domainName property value. The on-premises domain name of Hybrid Azure AD joined devices. This property is set by Intune.
func (m *Device) GetDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.domainName
    }
}
// Gets the enrollmentProfileName property value. Enrollment profile applied to the device. For example, Apple Device Enrollment Profile, Device enrollment - Corporate device identifiers, or Windows Autopilot profile name. This property is set by Intune.
func (m *Device) GetEnrollmentProfileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentProfileName
    }
}
// Gets the enrollmentType property value. Enrollment type of the device. This property is set by Intune. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement.
func (m *Device) GetEnrollmentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentType
    }
}
// Gets the extensionAttributes property value. Contains extension attributes 1-15 for the device. The individual extension attributes are not selectable. These properties are mastered in cloud and can be set during creation or update of a device object in Azure AD. Supports $filter (eq, NOT, startsWith).
func (m *Device) GetExtensionAttributes()(*OnPremisesExtensionAttributes) {
    if m == nil {
        return nil
    } else {
        return m.extensionAttributes
    }
}
// Gets the extensions property value. The collection of open extensions defined for the device. Read-only. Nullable.
func (m *Device) GetExtensions()([]Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// Gets the hostnames property value. List of hostNames for the device.
func (m *Device) GetHostnames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.hostnames
    }
}
// Gets the isCompliant property value. true if the device complies with Mobile Device Management (MDM) policies; otherwise, false. Read-only. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, NOT).
func (m *Device) GetIsCompliant()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCompliant
    }
}
// Gets the isManaged property value. true if the device is managed by a Mobile Device Management (MDM) app; otherwise, false. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, NOT).
func (m *Device) GetIsManaged()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isManaged
    }
}
// Gets the isRooted property value. true if device is rooted; false if device is jail-broken. This can only be updated by Intune.
func (m *Device) GetIsRooted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRooted
    }
}
// Gets the kind property value. Form factor of device. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) GetKind()(*string) {
    if m == nil {
        return nil
    } else {
        return m.kind
    }
}
// Gets the managementType property value. Management channel of the device.  This property is set by Intune. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
func (m *Device) GetManagementType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementType
    }
}
// Gets the manufacturer property value. Manufacturer of the device. Read-only.
func (m *Device) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// Gets the memberOf property value. Groups that this device is a member of. Read-only. Nullable. Supports $expand.
func (m *Device) GetMemberOf()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.memberOf
    }
}
// Gets the model property value. Model of the device. Read-only.
func (m *Device) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// Gets the name property value. Friendly name of a device. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the onPremisesLastSyncDateTime property value. The last time at which the object was synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Read-only. Supports $filter (eq, ne, NOT, ge, le, in).
func (m *Device) GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesLastSyncDateTime
    }
}
// Gets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Supports $filter (eq, ne, NOT, in).
func (m *Device) GetOnPremisesSyncEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesSyncEnabled
    }
}
// Gets the operatingSystem property value. The type of operating system on the device. Required. Supports $filter (eq, ne, NOT, ge, le, startsWith).
func (m *Device) GetOperatingSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystem
    }
}
// Gets the operatingSystemVersion property value. The version of the operating system on the device. Required. Supports $filter (eq, ne, NOT, ge, le, startsWith).
func (m *Device) GetOperatingSystemVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemVersion
    }
}
// Gets the physicalIds property value. For internal use only. Not nullable. Supports $filter (eq, NOT, ge, le, startsWith).
func (m *Device) GetPhysicalIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.physicalIds
    }
}
// Gets the platform property value. Platform of device. Only returned if user signs in with a Microsoft account as part of Project Rome. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) GetPlatform()(*string) {
    if m == nil {
        return nil
    } else {
        return m.platform
    }
}
// Gets the profileType property value. The profile type of the device. Possible values: RegisteredDevice (default), SecureVM, Printer, Shared, IoT.
func (m *Device) GetProfileType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.profileType
    }
}
// Gets the registeredOwners property value. The user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
func (m *Device) GetRegisteredOwners()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.registeredOwners
    }
}
// Gets the registeredUsers property value. Collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
func (m *Device) GetRegisteredUsers()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.registeredUsers
    }
}
// Gets the registrationDateTime property value. Date and time of when the device was registered. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Device) GetRegistrationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.registrationDateTime
    }
}
// Gets the status property value. Device is online or offline. Only returned if user signs in with a Microsoft account as part of Project Rome.
func (m *Device) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the systemLabels property value. List of labels applied to the device by the system.
func (m *Device) GetSystemLabels()([]string) {
    if m == nil {
        return nil
    } else {
        return m.systemLabels
    }
}
// Gets the transitiveMemberOf property value. Groups that the device is a member of. This operation is transitive. Supports $expand.
func (m *Device) GetTransitiveMemberOf()([]DirectoryObject) {
    if m == nil {
        return nil
    } else {
        return m.transitiveMemberOf
    }
}
// Gets the trustType property value. Type of trust for the joined device. Read-only. Possible values:  Workplace (indicates bring your own personal devices), AzureAd (Cloud only joined devices), ServerAd (on-premises domain joined devices joined to Azure AD). For more details, see Introduction to device management in Azure Active Directory
func (m *Device) GetTrustType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.trustType
    }
}
// Gets the usageRights property value. Represents the usage rights a device has been granted.
func (m *Device) GetUsageRights()([]UsageRight) {
    if m == nil {
        return nil
    } else {
        return m.usageRights
    }
}
// The deserialization information for the current model
func (m *Device) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["accountEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAccountEnabled(val)
        return nil
    }
    res["alternativeSecurityIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAlternativeSecurityId() })
        if err != nil {
            return err
        }
        res := make([]AlternativeSecurityId, len(val))
        for i, v := range val {
            res[i] = *(v.(*AlternativeSecurityId))
        }
        m.SetAlternativeSecurityIds(res)
        return nil
    }
    res["approximateLastSignInDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetApproximateLastSignInDateTime(val)
        return nil
    }
    res["commands"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCommand() })
        if err != nil {
            return err
        }
        res := make([]Command, len(val))
        for i, v := range val {
            res[i] = *(v.(*Command))
        }
        m.SetCommands(res)
        return nil
    }
    res["complianceExpirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetComplianceExpirationDateTime(val)
        return nil
    }
    res["deviceCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceCategory(val)
        return nil
    }
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceId(val)
        return nil
    }
    res["deviceMetadata"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceMetadata(val)
        return nil
    }
    res["deviceOwnership"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceOwnership(val)
        return nil
    }
    res["deviceVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDeviceVersion(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["domainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDomainName(val)
        return nil
    }
    res["enrollmentProfileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEnrollmentProfileName(val)
        return nil
    }
    res["enrollmentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEnrollmentType(val)
        return nil
    }
    res["extensionAttributes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesExtensionAttributes() })
        if err != nil {
            return err
        }
        m.SetExtensionAttributes(val.(*OnPremisesExtensionAttributes))
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExtension() })
        if err != nil {
            return err
        }
        res := make([]Extension, len(val))
        for i, v := range val {
            res[i] = *(v.(*Extension))
        }
        m.SetExtensions(res)
        return nil
    }
    res["hostnames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetHostnames(res)
        return nil
    }
    res["isCompliant"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsCompliant(val)
        return nil
    }
    res["isManaged"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsManaged(val)
        return nil
    }
    res["isRooted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRooted(val)
        return nil
    }
    res["kind"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetKind(val)
        return nil
    }
    res["managementType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagementType(val)
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManufacturer(val)
        return nil
    }
    res["memberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetMemberOf(res)
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetModel(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["onPremisesLastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetOnPremisesLastSyncDateTime(val)
        return nil
    }
    res["onPremisesSyncEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetOnPremisesSyncEnabled(val)
        return nil
    }
    res["operatingSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOperatingSystem(val)
        return nil
    }
    res["operatingSystemVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOperatingSystemVersion(val)
        return nil
    }
    res["physicalIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetPhysicalIds(res)
        return nil
    }
    res["platform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPlatform(val)
        return nil
    }
    res["profileType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProfileType(val)
        return nil
    }
    res["registeredOwners"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetRegisteredOwners(res)
        return nil
    }
    res["registeredUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetRegisteredUsers(res)
        return nil
    }
    res["registrationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetRegistrationDateTime(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStatus(val)
        return nil
    }
    res["systemLabels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetSystemLabels(res)
        return nil
    }
    res["transitiveMemberOf"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDirectoryObject() })
        if err != nil {
            return err
        }
        res := make([]DirectoryObject, len(val))
        for i, v := range val {
            res[i] = *(v.(*DirectoryObject))
        }
        m.SetTransitiveMemberOf(res)
        return nil
    }
    res["trustType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTrustType(val)
        return nil
    }
    res["usageRights"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUsageRight() })
        if err != nil {
            return err
        }
        res := make([]UsageRight, len(val))
        for i, v := range val {
            res[i] = *(v.(*UsageRight))
        }
        m.SetUsageRights(res)
        return nil
    }
    return res
}
func (m *Device) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Device) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accountEnabled", m.GetAccountEnabled())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAlternativeSecurityIds()))
        for i, v := range m.GetAlternativeSecurityIds() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("alternativeSecurityIds", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("approximateLastSignInDateTime", m.GetApproximateLastSignInDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCommands()))
        for i, v := range m.GetCommands() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("commands", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("complianceExpirationDateTime", m.GetComplianceExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceCategory", m.GetDeviceCategory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceMetadata", m.GetDeviceMetadata())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceOwnership", m.GetDeviceOwnership())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceVersion", m.GetDeviceVersion())
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
        err = writer.WriteStringValue("domainName", m.GetDomainName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("enrollmentProfileName", m.GetEnrollmentProfileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("enrollmentType", m.GetEnrollmentType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("extensionAttributes", m.GetExtensionAttributes())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("hostnames", m.GetHostnames())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCompliant", m.GetIsCompliant())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isManaged", m.GetIsManaged())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isRooted", m.GetIsRooted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kind", m.GetKind())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managementType", m.GetManagementType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMemberOf()))
        for i, v := range m.GetMemberOf() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("memberOf", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("onPremisesLastSyncDateTime", m.GetOnPremisesLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("onPremisesSyncEnabled", m.GetOnPremisesSyncEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystem", m.GetOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystemVersion", m.GetOperatingSystemVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("physicalIds", m.GetPhysicalIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("platform", m.GetPlatform())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("profileType", m.GetProfileType())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRegisteredOwners()))
        for i, v := range m.GetRegisteredOwners() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("registeredOwners", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRegisteredUsers()))
        for i, v := range m.GetRegisteredUsers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("registeredUsers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("registrationDateTime", m.GetRegistrationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("systemLabels", m.GetSystemLabels())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTransitiveMemberOf()))
        for i, v := range m.GetTransitiveMemberOf() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("transitiveMemberOf", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("trustType", m.GetTrustType())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUsageRights()))
        for i, v := range m.GetUsageRights() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("usageRights", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the accountEnabled property value. true if the account is enabled; otherwise, false. Required. Default is true.  Supports $filter (eq, ne, NOT, in). Only callers in Global Administrator and Cloud Device Administrator roles can set this property.
// Parameters:
//  - value : Value to set for the accountEnabled property.
func (m *Device) SetAccountEnabled(value *bool)() {
    m.accountEnabled = value
}
// Sets the alternativeSecurityIds property value. For internal use only. Not nullable. Supports $filter (eq, NOT, ge, le).
// Parameters:
//  - value : Value to set for the alternativeSecurityIds property.
func (m *Device) SetAlternativeSecurityIds(value []AlternativeSecurityId)() {
    m.alternativeSecurityIds = value
}
// Sets the approximateLastSignInDateTime property value. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter (eq, ne, NOT, ge, le) and $orderBy.
// Parameters:
//  - value : Value to set for the approximateLastSignInDateTime property.
func (m *Device) SetApproximateLastSignInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.approximateLastSignInDateTime = value
}
// Sets the commands property value. Set of commands sent to this device.
// Parameters:
//  - value : Value to set for the commands property.
func (m *Device) SetCommands(value []Command)() {
    m.commands = value
}
// Sets the complianceExpirationDateTime property value. The timestamp when the device is no longer deemed compliant. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// Parameters:
//  - value : Value to set for the complianceExpirationDateTime property.
func (m *Device) SetComplianceExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.complianceExpirationDateTime = value
}
// Sets the deviceCategory property value. User-defined property set by Intune to automatically add devices to groups and simplify managing devices.
// Parameters:
//  - value : Value to set for the deviceCategory property.
func (m *Device) SetDeviceCategory(value *string)() {
    m.deviceCategory = value
}
// Sets the deviceId property value. Unique identifier set by Azure Device Registration Service at the time of registration. Supports $filter (eq, ne, NOT, startsWith).
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *Device) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the deviceMetadata property value. For internal use only. Set to null.
// Parameters:
//  - value : Value to set for the deviceMetadata property.
func (m *Device) SetDeviceMetadata(value *string)() {
    m.deviceMetadata = value
}
// Sets the deviceOwnership property value. Ownership of the device. This property is set by Intune. Possible values are: unknown, company, personal.
// Parameters:
//  - value : Value to set for the deviceOwnership property.
func (m *Device) SetDeviceOwnership(value *string)() {
    m.deviceOwnership = value
}
// Sets the deviceVersion property value. For internal use only.
// Parameters:
//  - value : Value to set for the deviceVersion property.
func (m *Device) SetDeviceVersion(value *int32)() {
    m.deviceVersion = value
}
// Sets the displayName property value. The display name for the device. Required. Supports $filter (eq, ne, NOT, ge, le, in, startsWith), $search, and $orderBy.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Device) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the domainName property value. The on-premises domain name of Hybrid Azure AD joined devices. This property is set by Intune.
// Parameters:
//  - value : Value to set for the domainName property.
func (m *Device) SetDomainName(value *string)() {
    m.domainName = value
}
// Sets the enrollmentProfileName property value. Enrollment profile applied to the device. For example, Apple Device Enrollment Profile, Device enrollment - Corporate device identifiers, or Windows Autopilot profile name. This property is set by Intune.
// Parameters:
//  - value : Value to set for the enrollmentProfileName property.
func (m *Device) SetEnrollmentProfileName(value *string)() {
    m.enrollmentProfileName = value
}
// Sets the enrollmentType property value. Enrollment type of the device. This property is set by Intune. Possible values are: unknown, userEnrollment, deviceEnrollmentManager, appleBulkWithUser, appleBulkWithoutUser, windowsAzureADJoin, windowsBulkUserless, windowsAutoEnrollment, windowsBulkAzureDomainJoin, windowsCoManagement.
// Parameters:
//  - value : Value to set for the enrollmentType property.
func (m *Device) SetEnrollmentType(value *string)() {
    m.enrollmentType = value
}
// Sets the extensionAttributes property value. Contains extension attributes 1-15 for the device. The individual extension attributes are not selectable. These properties are mastered in cloud and can be set during creation or update of a device object in Azure AD. Supports $filter (eq, NOT, startsWith).
// Parameters:
//  - value : Value to set for the extensionAttributes property.
func (m *Device) SetExtensionAttributes(value *OnPremisesExtensionAttributes)() {
    m.extensionAttributes = value
}
// Sets the extensions property value. The collection of open extensions defined for the device. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the extensions property.
func (m *Device) SetExtensions(value []Extension)() {
    m.extensions = value
}
// Sets the hostnames property value. List of hostNames for the device.
// Parameters:
//  - value : Value to set for the hostnames property.
func (m *Device) SetHostnames(value []string)() {
    m.hostnames = value
}
// Sets the isCompliant property value. true if the device complies with Mobile Device Management (MDM) policies; otherwise, false. Read-only. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, NOT).
// Parameters:
//  - value : Value to set for the isCompliant property.
func (m *Device) SetIsCompliant(value *bool)() {
    m.isCompliant = value
}
// Sets the isManaged property value. true if the device is managed by a Mobile Device Management (MDM) app; otherwise, false. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, NOT).
// Parameters:
//  - value : Value to set for the isManaged property.
func (m *Device) SetIsManaged(value *bool)() {
    m.isManaged = value
}
// Sets the isRooted property value. true if device is rooted; false if device is jail-broken. This can only be updated by Intune.
// Parameters:
//  - value : Value to set for the isRooted property.
func (m *Device) SetIsRooted(value *bool)() {
    m.isRooted = value
}
// Sets the kind property value. Form factor of device. Only returned if user signs in with a Microsoft account as part of Project Rome.
// Parameters:
//  - value : Value to set for the kind property.
func (m *Device) SetKind(value *string)() {
    m.kind = value
}
// Sets the managementType property value. Management channel of the device.  This property is set by Intune. Possible values are: eas, mdm, easMdm, intuneClient, easIntuneClient, configurationManagerClient, configurationManagerClientMdm, configurationManagerClientMdmEas, unknown, jamf, googleCloudDevicePolicyController.
// Parameters:
//  - value : Value to set for the managementType property.
func (m *Device) SetManagementType(value *string)() {
    m.managementType = value
}
// Sets the manufacturer property value. Manufacturer of the device. Read-only.
// Parameters:
//  - value : Value to set for the manufacturer property.
func (m *Device) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// Sets the memberOf property value. Groups that this device is a member of. Read-only. Nullable. Supports $expand.
// Parameters:
//  - value : Value to set for the memberOf property.
func (m *Device) SetMemberOf(value []DirectoryObject)() {
    m.memberOf = value
}
// Sets the model property value. Model of the device. Read-only.
// Parameters:
//  - value : Value to set for the model property.
func (m *Device) SetModel(value *string)() {
    m.model = value
}
// Sets the name property value. Friendly name of a device. Only returned if user signs in with a Microsoft account as part of Project Rome.
// Parameters:
//  - value : Value to set for the name property.
func (m *Device) SetName(value *string)() {
    m.name = value
}
// Sets the onPremisesLastSyncDateTime property value. The last time at which the object was synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Read-only. Supports $filter (eq, ne, NOT, ge, le, in).
// Parameters:
//  - value : Value to set for the onPremisesLastSyncDateTime property.
func (m *Device) SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.onPremisesLastSyncDateTime = value
}
// Sets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Supports $filter (eq, ne, NOT, in).
// Parameters:
//  - value : Value to set for the onPremisesSyncEnabled property.
func (m *Device) SetOnPremisesSyncEnabled(value *bool)() {
    m.onPremisesSyncEnabled = value
}
// Sets the operatingSystem property value. The type of operating system on the device. Required. Supports $filter (eq, ne, NOT, ge, le, startsWith).
// Parameters:
//  - value : Value to set for the operatingSystem property.
func (m *Device) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
// Sets the operatingSystemVersion property value. The version of the operating system on the device. Required. Supports $filter (eq, ne, NOT, ge, le, startsWith).
// Parameters:
//  - value : Value to set for the operatingSystemVersion property.
func (m *Device) SetOperatingSystemVersion(value *string)() {
    m.operatingSystemVersion = value
}
// Sets the physicalIds property value. For internal use only. Not nullable. Supports $filter (eq, NOT, ge, le, startsWith).
// Parameters:
//  - value : Value to set for the physicalIds property.
func (m *Device) SetPhysicalIds(value []string)() {
    m.physicalIds = value
}
// Sets the platform property value. Platform of device. Only returned if user signs in with a Microsoft account as part of Project Rome. Only returned if user signs in with a Microsoft account as part of Project Rome.
// Parameters:
//  - value : Value to set for the platform property.
func (m *Device) SetPlatform(value *string)() {
    m.platform = value
}
// Sets the profileType property value. The profile type of the device. Possible values: RegisteredDevice (default), SecureVM, Printer, Shared, IoT.
// Parameters:
//  - value : Value to set for the profileType property.
func (m *Device) SetProfileType(value *string)() {
    m.profileType = value
}
// Sets the registeredOwners property value. The user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
// Parameters:
//  - value : Value to set for the registeredOwners property.
func (m *Device) SetRegisteredOwners(value []DirectoryObject)() {
    m.registeredOwners = value
}
// Sets the registeredUsers property value. Collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
// Parameters:
//  - value : Value to set for the registeredUsers property.
func (m *Device) SetRegisteredUsers(value []DirectoryObject)() {
    m.registeredUsers = value
}
// Sets the registrationDateTime property value. Date and time of when the device was registered. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// Parameters:
//  - value : Value to set for the registrationDateTime property.
func (m *Device) SetRegistrationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.registrationDateTime = value
}
// Sets the status property value. Device is online or offline. Only returned if user signs in with a Microsoft account as part of Project Rome.
// Parameters:
//  - value : Value to set for the status property.
func (m *Device) SetStatus(value *string)() {
    m.status = value
}
// Sets the systemLabels property value. List of labels applied to the device by the system.
// Parameters:
//  - value : Value to set for the systemLabels property.
func (m *Device) SetSystemLabels(value []string)() {
    m.systemLabels = value
}
// Sets the transitiveMemberOf property value. Groups that the device is a member of. This operation is transitive. Supports $expand.
// Parameters:
//  - value : Value to set for the transitiveMemberOf property.
func (m *Device) SetTransitiveMemberOf(value []DirectoryObject)() {
    m.transitiveMemberOf = value
}
// Sets the trustType property value. Type of trust for the joined device. Read-only. Possible values:  Workplace (indicates bring your own personal devices), AzureAd (Cloud only joined devices), ServerAd (on-premises domain joined devices joined to Azure AD). For more details, see Introduction to device management in Azure Active Directory
// Parameters:
//  - value : Value to set for the trustType property.
func (m *Device) SetTrustType(value *string)() {
    m.trustType = value
}
// Sets the usageRights property value. Represents the usage rights a device has been granted.
// Parameters:
//  - value : Value to set for the usageRights property.
func (m *Device) SetUsageRights(value []UsageRight)() {
    m.usageRights = value
}
