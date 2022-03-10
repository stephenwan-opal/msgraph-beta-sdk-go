package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i0a27aefac5e284898c6f31f3cec3cc5ca84d4aa4003ae193a4f82ea4c728e6a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/targetschedule"
    i2a944771b22cb5ea490413365394d6c74a1f803aba545a6f9b42769ec7319d09 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/appscope"
    i44d378692bae641c92a2c30cbfa78e06815a2cd19b7e2e10b8b725b11d8fc1cf "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/principal"
    i644c4ca2b29a486fb1565141c4c7f306a02458cb796d329032b2e0f088479442 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/activatedusing"
    iae59854ee3116fd1456d39cebc6a857e33eb8deea59f76066fe538101bc0f1b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/roledefinition"
    ibdcb7e68595c8a6e7db32816c49d6e162f0cf44239d2990965abb46b25451846 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/directoryscope"
    id18c7cf991b9a22e0a73384bb3123a5de24ee17868f1f7b779a94f0a9c5a6a3a "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/entitlementmanagement/roleassignmentschedulerequests/item/cancel"
)

// UnifiedRoleAssignmentScheduleRequestItemRequestBuilder provides operations to manage the roleAssignmentScheduleRequests property of the microsoft.graph.rbacApplication entity.
type UnifiedRoleAssignmentScheduleRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteOptions options for Delete
type UnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetOptions options for Get
type UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters get roleAssignmentScheduleRequests from roleManagement
type UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchOptions options for Patch
type UnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleRequestable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) ActivatedUsing()(*i644c4ca2b29a486fb1565141c4c7f306a02458cb796d329032b2e0f088479442.ActivatedUsingRequestBuilder) {
    return i644c4ca2b29a486fb1565141c4c7f306a02458cb796d329032b2e0f088479442.NewActivatedUsingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) AppScope()(*i2a944771b22cb5ea490413365394d6c74a1f803aba545a6f9b42769ec7319d09.AppScopeRequestBuilder) {
    return i2a944771b22cb5ea490413365394d6c74a1f803aba545a6f9b42769ec7319d09.NewAppScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Cancel()(*id18c7cf991b9a22e0a73384bb3123a5de24ee17868f1f7b779a94f0a9c5a6a3a.CancelRequestBuilder) {
    return id18c7cf991b9a22e0a73384bb3123a5de24ee17868f1f7b779a94f0a9c5a6a3a.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal instantiates a new UnifiedRoleAssignmentScheduleRequestItemRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    m := &UnifiedRoleAssignmentScheduleRequestItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement/roleAssignmentScheduleRequests/{unifiedRoleAssignmentScheduleRequest_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilder instantiates a new UnifiedRoleAssignmentScheduleRequestItemRequestBuilder and sets the default values.
func NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleAssignmentScheduleRequests for roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) CreateDeleteRequestInformation(options *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get roleAssignmentScheduleRequests from roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) CreateGetRequestInformation(options *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property roleAssignmentScheduleRequests in roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) CreatePatchRequestInformation(options *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property roleAssignmentScheduleRequests for roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Delete(options *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) DirectoryScope()(*ibdcb7e68595c8a6e7db32816c49d6e162f0cf44239d2990965abb46b25451846.DirectoryScopeRequestBuilder) {
    return ibdcb7e68595c8a6e7db32816c49d6e162f0cf44239d2990965abb46b25451846.NewDirectoryScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get roleAssignmentScheduleRequests from roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Get(options *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleRequestable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateUnifiedRoleAssignmentScheduleRequestFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleAssignmentScheduleRequestable), nil
}
// Patch update the navigation property roleAssignmentScheduleRequests in roleManagement
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Patch(options *UnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Principal()(*i44d378692bae641c92a2c30cbfa78e06815a2cd19b7e2e10b8b725b11d8fc1cf.PrincipalRequestBuilder) {
    return i44d378692bae641c92a2c30cbfa78e06815a2cd19b7e2e10b8b725b11d8fc1cf.NewPrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) RoleDefinition()(*iae59854ee3116fd1456d39cebc6a857e33eb8deea59f76066fe538101bc0f1b0.RoleDefinitionRequestBuilder) {
    return iae59854ee3116fd1456d39cebc6a857e33eb8deea59f76066fe538101bc0f1b0.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *UnifiedRoleAssignmentScheduleRequestItemRequestBuilder) TargetSchedule()(*i0a27aefac5e284898c6f31f3cec3cc5ca84d4aa4003ae193a4f82ea4c728e6a4.TargetScheduleRequestBuilder) {
    return i0a27aefac5e284898c6f31f3cec3cc5ca84d4aa4003ae193a4f82ea4c728e6a4.NewTargetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}