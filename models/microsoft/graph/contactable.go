package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Contactable 
type Contactable interface {
    OutlookItemable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssistantName()(*string)
    GetBirthday()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetChildren()([]string)
    GetCompanyName()(*string)
    GetDepartment()(*string)
    GetDisplayName()(*string)
    GetEmailAddresses()([]TypedEmailAddressable)
    GetExtensions()([]Extensionable)
    GetFileAs()(*string)
    GetFlag()(FollowupFlagable)
    GetGender()(*string)
    GetGeneration()(*string)
    GetGivenName()(*string)
    GetImAddresses()([]string)
    GetInitials()(*string)
    GetIsFavorite()(*bool)
    GetJobTitle()(*string)
    GetManager()(*string)
    GetMiddleName()(*string)
    GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable)
    GetNickName()(*string)
    GetOfficeLocation()(*string)
    GetParentFolderId()(*string)
    GetPersonalNotes()(*string)
    GetPhones()([]Phoneable)
    GetPhoto()(ProfilePhotoable)
    GetPostalAddresses()([]PhysicalAddressable)
    GetProfession()(*string)
    GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable)
    GetSpouseName()(*string)
    GetSurname()(*string)
    GetTitle()(*string)
    GetWebsites()([]Websiteable)
    GetWeddingAnniversary()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetYomiCompanyName()(*string)
    GetYomiGivenName()(*string)
    GetYomiSurname()(*string)
    SetAssistantName(value *string)()
    SetBirthday(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetChildren(value []string)()
    SetCompanyName(value *string)()
    SetDepartment(value *string)()
    SetDisplayName(value *string)()
    SetEmailAddresses(value []TypedEmailAddressable)()
    SetExtensions(value []Extensionable)()
    SetFileAs(value *string)()
    SetFlag(value FollowupFlagable)()
    SetGender(value *string)()
    SetGeneration(value *string)()
    SetGivenName(value *string)()
    SetImAddresses(value []string)()
    SetInitials(value *string)()
    SetIsFavorite(value *bool)()
    SetJobTitle(value *string)()
    SetManager(value *string)()
    SetMiddleName(value *string)()
    SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)()
    SetNickName(value *string)()
    SetOfficeLocation(value *string)()
    SetParentFolderId(value *string)()
    SetPersonalNotes(value *string)()
    SetPhones(value []Phoneable)()
    SetPhoto(value ProfilePhotoable)()
    SetPostalAddresses(value []PhysicalAddressable)()
    SetProfession(value *string)()
    SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)()
    SetSpouseName(value *string)()
    SetSurname(value *string)()
    SetTitle(value *string)()
    SetWebsites(value []Websiteable)()
    SetWeddingAnniversary(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetYomiCompanyName(value *string)()
    SetYomiGivenName(value *string)()
    SetYomiSurname(value *string)()
}