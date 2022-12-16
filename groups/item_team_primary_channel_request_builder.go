package groups

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTeamPrimaryChannelRequestBuilder builds and executes requests for operations under \groups\{group-id}\team\primaryChannel
type ItemTeamPrimaryChannelRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewItemTeamPrimaryChannelRequestBuilderInternal instantiates a new PrimaryChannelRequestBuilder and sets the default values.
func NewItemTeamPrimaryChannelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimaryChannelRequestBuilder) {
    m := &ItemTeamPrimaryChannelRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/team/primaryChannel";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemTeamPrimaryChannelRequestBuilder instantiates a new PrimaryChannelRequestBuilder and sets the default values.
func NewItemTeamPrimaryChannelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimaryChannelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamPrimaryChannelRequestBuilderInternal(urlParams, requestAdapter)
}
// Members provides operations to manage the members property of the microsoft.graph.channel entity.
func (m *ItemTeamPrimaryChannelRequestBuilder) Members()(*ItemTeamPrimaryChannelMembersRequestBuilder) {
    return NewItemTeamPrimaryChannelMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById provides operations to manage the members property of the microsoft.graph.channel entity.
func (m *ItemTeamPrimaryChannelRequestBuilder) MembersById(id string)(*ItemTeamPrimaryChannelMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewItemTeamPrimaryChannelMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
