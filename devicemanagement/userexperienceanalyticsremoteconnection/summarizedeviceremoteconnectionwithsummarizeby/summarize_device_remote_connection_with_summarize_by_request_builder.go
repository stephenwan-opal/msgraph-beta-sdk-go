package summarizedeviceremoteconnectionwithsummarizeby

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder provides operations to call the summarizeDeviceRemoteConnection method.
type SummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilderInternal instantiates a new SummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder and sets the default values.
func NewSummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, summarizeBy *string)(*SummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder) {
    m := &SummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsRemoteConnection/microsoft.graph.summarizeDeviceRemoteConnection(summarizeBy='{summarizeBy}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if summarizeBy != nil {
        urlTplParams["summarizeBy"] = *summarizeBy
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder instantiates a new SummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder and sets the default values.
func NewSummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function summarizeDeviceRemoteConnection
func (m *SummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function summarizeDeviceRemoteConnection
func (m *SummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *SummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function summarizeDeviceRemoteConnection
func (m *SummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder) Get()(SummarizeDeviceRemoteConnectionWithSummarizeByResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function summarizeDeviceRemoteConnection
func (m *SummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *SummarizeDeviceRemoteConnectionWithSummarizeByRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(SummarizeDeviceRemoteConnectionWithSummarizeByResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateSummarizeDeviceRemoteConnectionWithSummarizeByResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(SummarizeDeviceRemoteConnectionWithSummarizeByResponseable), nil
}
