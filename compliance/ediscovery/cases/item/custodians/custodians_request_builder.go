package custodians

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
    i191ea490108a56b4c3739abc85eafe23cf697b4c891e9defe46b970ab3a3a64d "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/custodians/count"
    icfdf97924ba8a4d9281a6f2a511eb73eea3b32f69a4073b62814f9b2bb76a3ba "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/custodians/removehold"
    ie8b1f413e0c4033bfb87bd144e4d8bb61544e286c96c654a2c31da85aa24d2a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/compliance/ediscovery/cases/item/custodians/applyhold"
)

// CustodiansRequestBuilder provides operations to manage the custodians property of the microsoft.graph.ediscovery.case entity.
type CustodiansRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CustodiansRequestBuilderGetQueryParameters returns a list of case custodian objects for this case.  Nullable.
type CustodiansRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// CustodiansRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustodiansRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CustodiansRequestBuilderGetQueryParameters
}
// CustodiansRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CustodiansRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplyHold the applyHold property
func (m *CustodiansRequestBuilder) ApplyHold()(*ie8b1f413e0c4033bfb87bd144e4d8bb61544e286c96c654a2c31da85aa24d2a6.ApplyHoldRequestBuilder) {
    return ie8b1f413e0c4033bfb87bd144e4d8bb61544e286c96c654a2c31da85aa24d2a6.NewApplyHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCustodiansRequestBuilderInternal instantiates a new CustodiansRequestBuilder and sets the default values.
func NewCustodiansRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustodiansRequestBuilder) {
    m := &CustodiansRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/custodians{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCustodiansRequestBuilder instantiates a new CustodiansRequestBuilder and sets the default values.
func NewCustodiansRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CustodiansRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCustodiansRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *CustodiansRequestBuilder) Count()(*i191ea490108a56b4c3739abc85eafe23cf697b4c891e9defe46b970ab3a3a64d.CountRequestBuilder) {
    return i191ea490108a56b4c3739abc85eafe23cf697b4c891e9defe46b970ab3a3a64d.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation returns a list of case custodian objects for this case.  Nullable.
func (m *CustodiansRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration returns a list of case custodian objects for this case.  Nullable.
func (m *CustodiansRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *CustodiansRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property to custodians for compliance
func (m *CustodiansRequestBuilder) CreatePostRequestInformation(body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to custodians for compliance
func (m *CustodiansRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable, requestConfiguration *CustodiansRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get returns a list of case custodian objects for this case.  Nullable.
func (m *CustodiansRequestBuilder) Get()(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CustodianCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler returns a list of case custodian objects for this case.  Nullable.
func (m *CustodiansRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *CustodiansRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CustodianCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateCustodianCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CustodianCollectionResponseable), nil
}
// Post create new navigation property to custodians for compliance
func (m *CustodiansRequestBuilder) Post(body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler create new navigation property to custodians for compliance
func (m *CustodiansRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable, requestConfiguration *CustodiansRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.CreateCustodianFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.Custodianable), nil
}
// RemoveHold the removeHold property
func (m *CustodiansRequestBuilder) RemoveHold()(*icfdf97924ba8a4d9281a6f2a511eb73eea3b32f69a4073b62814f9b2bb76a3ba.RemoveHoldRequestBuilder) {
    return icfdf97924ba8a4d9281a6f2a511eb73eea3b32f69a4073b62814f9b2bb76a3ba.NewRemoveHoldRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
