package removehold

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// RemoveHoldRequestBuilder provides operations to call the removeHold method.
type RemoveHoldRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RemoveHoldRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RemoveHoldRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRemoveHoldRequestBuilderInternal instantiates a new RemoveHoldRequestBuilder and sets the default values.
func NewRemoveHoldRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RemoveHoldRequestBuilder) {
    m := &RemoveHoldRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/custodians/microsoft.graph.ediscovery.removeHold";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRemoveHoldRequestBuilder instantiates a new RemoveHoldRequestBuilder and sets the default values.
func NewRemoveHoldRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RemoveHoldRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRemoveHoldRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action removeHold
func (m *RemoveHoldRequestBuilder) CreatePostRequestInformation(body RemoveHoldRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action removeHold
func (m *RemoveHoldRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body RemoveHoldRequestBodyable, requestConfiguration *RemoveHoldRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action removeHold
func (m *RemoveHoldRequestBuilder) Post(body RemoveHoldRequestBodyable)(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action removeHold
func (m *RemoveHoldRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body RemoveHoldRequestBodyable, requestConfiguration *RemoveHoldRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, nil)
    if err != nil {
        return err
    }
    return nil
}
