package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningObjectSummaryable 
type ProvisioningObjectSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(*string)
    GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetChangeId()(*string)
    GetCycleId()(*string)
    GetDurationInMilliseconds()(*int32)
    GetInitiatedBy()(Initiatorable)
    GetJobId()(*string)
    GetModifiedProperties()([]ModifiedPropertyable)
    GetProvisioningAction()(*ProvisioningAction)
    GetProvisioningStatusInfo()(ProvisioningStatusInfoable)
    GetProvisioningSteps()([]ProvisioningStepable)
    GetServicePrincipal()(ProvisioningServicePrincipalable)
    GetSourceIdentity()(ProvisionedIdentityable)
    GetSourceSystem()(ProvisioningSystemable)
    GetStatusInfo()(StatusBaseable)
    GetTargetIdentity()(ProvisionedIdentityable)
    GetTargetSystem()(ProvisioningSystemable)
    GetTenantId()(*string)
    SetAction(value *string)()
    SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetChangeId(value *string)()
    SetCycleId(value *string)()
    SetDurationInMilliseconds(value *int32)()
    SetInitiatedBy(value Initiatorable)()
    SetJobId(value *string)()
    SetModifiedProperties(value []ModifiedPropertyable)()
    SetProvisioningAction(value *ProvisioningAction)()
    SetProvisioningStatusInfo(value ProvisioningStatusInfoable)()
    SetProvisioningSteps(value []ProvisioningStepable)()
    SetServicePrincipal(value ProvisioningServicePrincipalable)()
    SetSourceIdentity(value ProvisionedIdentityable)()
    SetSourceSystem(value ProvisioningSystemable)()
    SetStatusInfo(value StatusBaseable)()
    SetTargetIdentity(value ProvisionedIdentityable)()
    SetTargetSystem(value ProvisioningSystemable)()
    SetTenantId(value *string)()
}