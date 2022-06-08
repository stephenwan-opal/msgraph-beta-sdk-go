package getmobileappcountwithstatus

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetMobileAppCountWithStatusRequestBuilder provides operations to call the getMobileAppCount method.
type GetMobileAppCountWithStatusRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetMobileAppCountWithStatusRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetMobileAppCountWithStatusRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetMobileAppCountWithStatusRequestBuilderInternal instantiates a new GetMobileAppCountWithStatusRequestBuilder and sets the default values.
func NewGetMobileAppCountWithStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, status *string)(*GetMobileAppCountWithStatusRequestBuilder) {
    m := &GetMobileAppCountWithStatusRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/microsoft.graph.getMobileAppCount(status='{status}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if status != nil {
        urlTplParams["status"] = *status
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetMobileAppCountWithStatusRequestBuilder instantiates a new GetMobileAppCountWithStatusRequestBuilder and sets the default values.
func NewGetMobileAppCountWithStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetMobileAppCountWithStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetMobileAppCountWithStatusRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getMobileAppCount
func (m *GetMobileAppCountWithStatusRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getMobileAppCount
func (m *GetMobileAppCountWithStatusRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetMobileAppCountWithStatusRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function getMobileAppCount
func (m *GetMobileAppCountWithStatusRequestBuilder) Get()(GetMobileAppCountWithStatusResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function getMobileAppCount
func (m *GetMobileAppCountWithStatusRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GetMobileAppCountWithStatusRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetMobileAppCountWithStatusResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetMobileAppCountWithStatusResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetMobileAppCountWithStatusResponseable), nil
}
