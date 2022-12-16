package groups

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTeamRequestBuilder builds and executes requests for operations under \groups\{group-id}\team
type ItemTeamRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Channels the channels property
func (m *ItemTeamRequestBuilder) Channels()(*ItemTeamChannelsRequestBuilder) {
    return NewItemTeamChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChannelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.groups.item.team.channels.item collection
func (m *ItemTeamRequestBuilder) ChannelsById(id string)(*ItemTeamChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewItemTeamChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewItemTeamRequestBuilderInternal instantiates a new TeamRequestBuilder and sets the default values.
func NewItemTeamRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamRequestBuilder) {
    m := &ItemTeamRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/team";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemTeamRequestBuilder instantiates a new TeamRequestBuilder and sets the default values.
func NewItemTeamRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamRequestBuilderInternal(urlParams, requestAdapter)
}
// Members provides operations to manage the members property of the microsoft.graph.team entity.
func (m *ItemTeamRequestBuilder) Members()(*ItemTeamMembersRequestBuilder) {
    return NewItemTeamMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById provides operations to manage the members property of the microsoft.graph.team entity.
func (m *ItemTeamRequestBuilder) MembersById(id string)(*ItemTeamMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewItemTeamMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// PrimaryChannel the primaryChannel property
func (m *ItemTeamRequestBuilder) PrimaryChannel()(*ItemTeamPrimaryChannelRequestBuilder) {
    return NewItemTeamPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tags the tags property
func (m *ItemTeamRequestBuilder) Tags()(*ItemTeamTagsRequestBuilder) {
    return NewItemTeamTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TagsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.groups.item.team.tags.item collection
func (m *ItemTeamRequestBuilder) TagsById(id string)(*ItemTeamTagsTeamworkTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTag%2Did"] = id
    }
    return NewItemTeamTagsTeamworkTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
