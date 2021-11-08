package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i233cfcd85fda77bc7b8ed84ae3eab86225b54c43cf6b9995aec4e929ecae0d5f "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignmentrequests/item/roledefinition"
    i3c09d8808940e57cffe2a4b43ec6c03afed85c4d1b8c2a6b6f69170f1707c06e "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignmentrequests/item/cancel"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i781dea20c6497af5c00396ab4ab2df03f2c4e8031835bbd6dffd7f1eb7bc6bbd "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignmentrequests/item/updaterequest"
    i7a382583186c7a5d8b72c51d185f4589248347a161c9a5a7a7d8d885ab571893 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignmentrequests/item/resource"
    ia72138d764aac9bd3a352e3057f1578dd8230b4045684066c5c0d36a2be3badb "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroleassignmentrequests/item/subject"
)

// Builds and executes requests for operations under \governanceRoleAssignmentRequests\{governanceRoleAssignmentRequest-id}
type GovernanceRoleAssignmentRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type GovernanceRoleAssignmentRequestRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type GovernanceRoleAssignmentRequestRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GovernanceRoleAssignmentRequestRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get entity from governanceRoleAssignmentRequests by key
type GovernanceRoleAssignmentRequestRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type GovernanceRoleAssignmentRequestRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceRoleAssignmentRequest;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *GovernanceRoleAssignmentRequestRequestBuilder) Cancel()(*i3c09d8808940e57cffe2a4b43ec6c03afed85c4d1b8c2a6b6f69170f1707c06e.CancelRequestBuilder) {
    return i3c09d8808940e57cffe2a4b43ec6c03afed85c4d1b8c2a6b6f69170f1707c06e.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new GovernanceRoleAssignmentRequestRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewGovernanceRoleAssignmentRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GovernanceRoleAssignmentRequestRequestBuilder) {
    m := &GovernanceRoleAssignmentRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/governanceRoleAssignmentRequests/{governanceRoleAssignmentRequest_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new GovernanceRoleAssignmentRequestRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewGovernanceRoleAssignmentRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GovernanceRoleAssignmentRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGovernanceRoleAssignmentRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete entity from governanceRoleAssignmentRequests
// Parameters:
//  - options : Options for the request
func (m *GovernanceRoleAssignmentRequestRequestBuilder) CreateDeleteRequestInformation(options *GovernanceRoleAssignmentRequestRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get entity from governanceRoleAssignmentRequests by key
// Parameters:
//  - options : Options for the request
func (m *GovernanceRoleAssignmentRequestRequestBuilder) CreateGetRequestInformation(options *GovernanceRoleAssignmentRequestRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// Update entity in governanceRoleAssignmentRequests
// Parameters:
//  - options : Options for the request
func (m *GovernanceRoleAssignmentRequestRequestBuilder) CreatePatchRequestInformation(options *GovernanceRoleAssignmentRequestRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete entity from governanceRoleAssignmentRequests
// Parameters:
//  - options : Options for the request
func (m *GovernanceRoleAssignmentRequestRequestBuilder) Delete(options *GovernanceRoleAssignmentRequestRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get entity from governanceRoleAssignmentRequests by key
// Parameters:
//  - options : Options for the request
func (m *GovernanceRoleAssignmentRequestRequestBuilder) Get(options *GovernanceRoleAssignmentRequestRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceRoleAssignmentRequest, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewGovernanceRoleAssignmentRequest() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceRoleAssignmentRequest), nil
}
// Update entity in governanceRoleAssignmentRequests
// Parameters:
//  - options : Options for the request
func (m *GovernanceRoleAssignmentRequestRequestBuilder) Patch(options *GovernanceRoleAssignmentRequestRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *GovernanceRoleAssignmentRequestRequestBuilder) Resource()(*i7a382583186c7a5d8b72c51d185f4589248347a161c9a5a7a7d8d885ab571893.ResourceRequestBuilder) {
    return i7a382583186c7a5d8b72c51d185f4589248347a161c9a5a7a7d8d885ab571893.NewResourceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GovernanceRoleAssignmentRequestRequestBuilder) RoleDefinition()(*i233cfcd85fda77bc7b8ed84ae3eab86225b54c43cf6b9995aec4e929ecae0d5f.RoleDefinitionRequestBuilder) {
    return i233cfcd85fda77bc7b8ed84ae3eab86225b54c43cf6b9995aec4e929ecae0d5f.NewRoleDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GovernanceRoleAssignmentRequestRequestBuilder) Subject()(*ia72138d764aac9bd3a352e3057f1578dd8230b4045684066c5c0d36a2be3badb.SubjectRequestBuilder) {
    return ia72138d764aac9bd3a352e3057f1578dd8230b4045684066c5c0d36a2be3badb.NewSubjectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GovernanceRoleAssignmentRequestRequestBuilder) UpdateRequest()(*i781dea20c6497af5c00396ab4ab2df03f2c4e8031835bbd6dffd7f1eb7bc6bbd.UpdateRequestRequestBuilder) {
    return i781dea20c6497af5c00396ab4ab2df03f2c4e8031835bbd6dffd7f1eb7bc6bbd.NewUpdateRequestRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
