package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CertificateConnectorDetailsItemGetHealthMetricsRequestBuilder provides operations to call the getHealthMetrics method.
type CertificateConnectorDetailsItemGetHealthMetricsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CertificateConnectorDetailsItemGetHealthMetricsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CertificateConnectorDetailsItemGetHealthMetricsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCertificateConnectorDetailsItemGetHealthMetricsRequestBuilderInternal instantiates a new GetHealthMetricsRequestBuilder and sets the default values.
func NewCertificateConnectorDetailsItemGetHealthMetricsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateConnectorDetailsItemGetHealthMetricsRequestBuilder) {
    m := &CertificateConnectorDetailsItemGetHealthMetricsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/certificateConnectorDetails/{certificateConnectorDetails%2Did}/microsoft.graph.getHealthMetrics";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCertificateConnectorDetailsItemGetHealthMetricsRequestBuilder instantiates a new GetHealthMetricsRequestBuilder and sets the default values.
func NewCertificateConnectorDetailsItemGetHealthMetricsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CertificateConnectorDetailsItemGetHealthMetricsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCertificateConnectorDetailsItemGetHealthMetricsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getHealthMetrics
func (m *CertificateConnectorDetailsItemGetHealthMetricsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body CertificateConnectorDetailsItemGetHealthMetricsPostRequestBodyable, requestConfiguration *CertificateConnectorDetailsItemGetHealthMetricsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action getHealthMetrics
func (m *CertificateConnectorDetailsItemGetHealthMetricsRequestBuilder) Post(ctx context.Context, body CertificateConnectorDetailsItemGetHealthMetricsPostRequestBodyable, requestConfiguration *CertificateConnectorDetailsItemGetHealthMetricsRequestBuilderPostRequestConfiguration)(CertificateConnectorDetailsItemGetHealthMetricsResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, CreateCertificateConnectorDetailsItemGetHealthMetricsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CertificateConnectorDetailsItemGetHealthMetricsResponseable), nil
}
