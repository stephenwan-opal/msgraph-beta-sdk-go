package deployments

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i54250e9a82a71ee0bd81c5f39fef1a69361cc61b0e5d5ee2f19cfce570e73b2d "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatestepversions/item/deployments/ref"
)

// Builds and executes requests for operations under \tenantRelationships\managedTenants\managementTemplateStepVersions\{managementTemplateStepVersion-id}\deployments
type DeploymentsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type DeploymentsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeploymentsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get deployments from tenantRelationships
type DeploymentsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select_escaped []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// Instantiates a new DeploymentsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDeploymentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeploymentsRequestBuilder) {
    m := &DeploymentsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managementTemplateStepVersions/{managementTemplateStepVersion_id}/deployments{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new DeploymentsRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDeploymentsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeploymentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeploymentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get deployments from tenantRelationships
// Parameters:
//  - options : Options for the request
func (m *DeploymentsRequestBuilder) CreateGetRequestInformation(options *DeploymentsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get deployments from tenantRelationships
// Parameters:
//  - options : Options for the request
func (m *DeploymentsRequestBuilder) Get(options *DeploymentsRequestBuilderGetOptions)(*DeploymentsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeploymentsResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*DeploymentsResponse), nil
}
func (m *DeploymentsRequestBuilder) Ref()(*i54250e9a82a71ee0bd81c5f39fef1a69361cc61b0e5d5ee2f19cfce570e73b2d.RefRequestBuilder) {
    return i54250e9a82a71ee0bd81c5f39fef1a69361cc61b0e5d5ee2f19cfce570e73b2d.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}