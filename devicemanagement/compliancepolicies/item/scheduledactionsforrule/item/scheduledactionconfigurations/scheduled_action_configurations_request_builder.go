package scheduledactionconfigurations

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i270f80d11b93f2267beef99b5035bd7467c1076e161892c73b2b7017a77b77f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/compliancepolicies/item/scheduledactionsforrule/item/scheduledactionconfigurations/ref"
)

// Builds and executes requests for operations under \deviceManagement\compliancePolicies\{deviceManagementCompliancePolicy-id}\scheduledActionsForRule\{deviceManagementComplianceScheduledActionForRule-id}\scheduledActionConfigurations
type ScheduledActionConfigurationsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type ScheduledActionConfigurationsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ScheduledActionConfigurationsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
type ScheduledActionConfigurationsRequestBuilderGetQueryParameters struct {
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
// Instantiates a new ScheduledActionConfigurationsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewScheduledActionConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduledActionConfigurationsRequestBuilder) {
    m := &ScheduledActionConfigurationsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/compliancePolicies/{deviceManagementCompliancePolicy_id}/scheduledActionsForRule/{deviceManagementComplianceScheduledActionForRule_id}/scheduledActionConfigurations{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ScheduledActionConfigurationsRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewScheduledActionConfigurationsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduledActionConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScheduledActionConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// The list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
// Parameters:
//  - options : Options for the request
func (m *ScheduledActionConfigurationsRequestBuilder) CreateGetRequestInformation(options *ScheduledActionConfigurationsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100 elements.
// Parameters:
//  - options : Options for the request
func (m *ScheduledActionConfigurationsRequestBuilder) Get(options *ScheduledActionConfigurationsRequestBuilderGetOptions)(*ScheduledActionConfigurationsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewScheduledActionConfigurationsResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ScheduledActionConfigurationsResponse), nil
}
func (m *ScheduledActionConfigurationsRequestBuilder) Ref()(*i270f80d11b93f2267beef99b5035bd7467c1076e161892c73b2b7017a77b77f6.RefRequestBuilder) {
    return i270f80d11b93f2267beef99b5035bd7467c1076e161892c73b2b7017a77b77f6.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
