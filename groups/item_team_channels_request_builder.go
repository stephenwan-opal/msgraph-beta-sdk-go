package groups

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTeamChannelsRequestBuilder builds and executes requests for operations under \groups\{group-id}\team\channels
type ItemTeamChannelsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemTeamChannelsRequestBuilderInternal instantiates a new ChannelsRequestBuilder and sets the default values.
func NewItemTeamChannelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamChannelsRequestBuilder) {
    m := &ItemTeamChannelsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/team/channels";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemTeamChannelsRequestBuilder instantiates a new ChannelsRequestBuilder and sets the default values.
func NewItemTeamChannelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamChannelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamChannelsRequestBuilderInternal(urlParams, requestAdapter)
}
