package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TargetedManagedAppProtectionable 
type TargetedManagedAppProtectionable interface {
    ManagedAppProtectionable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppGroupType()(*TargetedManagedAppGroupType)
    GetAssignments()([]TargetedManagedAppPolicyAssignmentable)
    GetIsAssigned()(*bool)
    GetTargetedAppManagementLevels()(*AppManagementLevel)
    SetAppGroupType(value *TargetedManagedAppGroupType)()
    SetAssignments(value []TargetedManagedAppPolicyAssignmentable)()
    SetIsAssigned(value *bool)()
    SetTargetedAppManagementLevels(value *AppManagementLevel)()
}