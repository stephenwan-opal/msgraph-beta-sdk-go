package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// WorkloadActionDeploymentStatusable 
type WorkloadActionDeploymentStatusable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActionId()(*string)
    GetDeployedPolicyId()(*string)
    GetError()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GenericErrorable)
    GetExcludeGroups()([]string)
    GetIncludeAllUsers()(*bool)
    GetIncludeGroups()([]string)
    GetLastDeploymentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*WorkloadActionStatus)
    SetActionId(value *string)()
    SetDeployedPolicyId(value *string)()
    SetError(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GenericErrorable)()
    SetExcludeGroups(value []string)()
    SetIncludeAllUsers(value *bool)()
    SetIncludeGroups(value []string)()
    SetLastDeploymentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *WorkloadActionStatus)()
}