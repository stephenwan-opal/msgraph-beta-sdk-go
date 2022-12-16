package groups

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTeamTagsTeamworkTagItemRequestBuilder builds and executes requests for operations under \groups\{group-id}\team\tags\{teamworkTag-id}
type ItemTeamTagsTeamworkTagItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemTeamTagsTeamworkTagItemRequestBuilderInternal instantiates a new TeamworkTagItemRequestBuilder and sets the default values.
func NewItemTeamTagsTeamworkTagItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamTagsTeamworkTagItemRequestBuilder) {
    m := &ItemTeamTagsTeamworkTagItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/team/tags/{teamworkTag%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemTeamTagsTeamworkTagItemRequestBuilder instantiates a new TeamworkTagItemRequestBuilder and sets the default values.
func NewItemTeamTagsTeamworkTagItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamTagsTeamworkTagItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamTagsTeamworkTagItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Members provides operations to manage the members property of the microsoft.graph.teamworkTag entity.
func (m *ItemTeamTagsTeamworkTagItemRequestBuilder) Members()(*ItemTeamTagsItemMembersRequestBuilder) {
    return NewItemTeamTagsItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById provides operations to manage the members property of the microsoft.graph.teamworkTag entity.
func (m *ItemTeamTagsTeamworkTagItemRequestBuilder) MembersById(id string)(*ItemTeamTagsItemMembersTeamworkTagMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTagMember%2Did"] = id
    }
    return NewItemTeamTagsItemMembersTeamworkTagMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
