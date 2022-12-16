package groups

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTeamTagsRequestBuilder builds and executes requests for operations under \groups\{group-id}\team\tags
type ItemTeamTagsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemTeamTagsRequestBuilderInternal instantiates a new TagsRequestBuilder and sets the default values.
func NewItemTeamTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamTagsRequestBuilder) {
    m := &ItemTeamTagsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/team/tags";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemTeamTagsRequestBuilder instantiates a new TagsRequestBuilder and sets the default values.
func NewItemTeamTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamTagsRequestBuilderInternal(urlParams, requestAdapter)
}
