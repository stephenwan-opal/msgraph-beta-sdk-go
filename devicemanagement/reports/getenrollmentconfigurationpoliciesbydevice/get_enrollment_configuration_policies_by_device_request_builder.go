package getenrollmentconfigurationpoliciesbydevice

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetEnrollmentConfigurationPoliciesByDeviceRequestBuilder provides operations to call the getEnrollmentConfigurationPoliciesByDevice method.
type GetEnrollmentConfigurationPoliciesByDeviceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetEnrollmentConfigurationPoliciesByDeviceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetEnrollmentConfigurationPoliciesByDeviceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetEnrollmentConfigurationPoliciesByDeviceRequestBuilderInternal instantiates a new GetEnrollmentConfigurationPoliciesByDeviceRequestBuilder and sets the default values.
func NewGetEnrollmentConfigurationPoliciesByDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetEnrollmentConfigurationPoliciesByDeviceRequestBuilder) {
    m := &GetEnrollmentConfigurationPoliciesByDeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getEnrollmentConfigurationPoliciesByDevice";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetEnrollmentConfigurationPoliciesByDeviceRequestBuilder instantiates a new GetEnrollmentConfigurationPoliciesByDeviceRequestBuilder and sets the default values.
func NewGetEnrollmentConfigurationPoliciesByDeviceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetEnrollmentConfigurationPoliciesByDeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetEnrollmentConfigurationPoliciesByDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getEnrollmentConfigurationPoliciesByDevice
func (m *GetEnrollmentConfigurationPoliciesByDeviceRequestBuilder) CreatePostRequestInformation(body GetEnrollmentConfigurationPoliciesByDevicePostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getEnrollmentConfigurationPoliciesByDevice
func (m *GetEnrollmentConfigurationPoliciesByDeviceRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetEnrollmentConfigurationPoliciesByDevicePostRequestBodyable, requestConfiguration *GetEnrollmentConfigurationPoliciesByDeviceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action getEnrollmentConfigurationPoliciesByDevice
func (m *GetEnrollmentConfigurationPoliciesByDeviceRequestBuilder) Post(body GetEnrollmentConfigurationPoliciesByDevicePostRequestBodyable)([]byte, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action getEnrollmentConfigurationPoliciesByDevice
func (m *GetEnrollmentConfigurationPoliciesByDeviceRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body GetEnrollmentConfigurationPoliciesByDevicePostRequestBodyable, requestConfiguration *GetEnrollmentConfigurationPoliciesByDeviceRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)([]byte, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(requestInfo, "byte", responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.([]byte), nil
}
