package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/termstore"
    i0b982fc6fce1ac2d0c2303097c8efccfbe2a0334914647564fdc68d25e2e8bf5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/termstore/groups/item/sets/item/terms"
    i2e8863714d1e920d7702b4a703714d81cd242645914443f004188922e9fdd7db "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/termstore/groups/item/sets/item/parentgroup"
    i46a937ae66eb746fc0b61e97303d25716a338af8b2a5d4dc9e7c0e2f36eb7c7a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/termstore/groups/item/sets/item/relations"
    i87acd307916195400cf09b0190e7ffc106ecc0a5bdb710da08090017e0f5d38d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/termstore/groups/item/sets/item/children"
    i3f4638c02945c3b951bb6beb6a994164e45df748d9433b5561e4fa4ec9fb8e11 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/termstore/groups/item/sets/item/relations/item"
    i7c7c17bddaf41293a67c772c6cfde4f761310e7bf1903d604c75b8bdccc5f32e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/termstore/groups/item/sets/item/children/item"
    id9bf09181f4b938e9dff34744950bbd02803ebc288f203778ba14a909a7eca97 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/sites/item/termstore/groups/item/sets/item/terms/item"
)

// SetItemRequestBuilder provides operations to manage the sets property of the microsoft.graph.termStore.group entity.
type SetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SetItemRequestBuilderGetQueryParameters all sets under the group in a term [store].
type SetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SetItemRequestBuilderGetQueryParameters
}
// SetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Children the children property
func (m *SetItemRequestBuilder) Children()(*i87acd307916195400cf09b0190e7ffc106ecc0a5bdb710da08090017e0f5d38d.ChildrenRequestBuilder) {
    return i87acd307916195400cf09b0190e7ffc106ecc0a5bdb710da08090017e0f5d38d.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.sites.item.termStore.groups.item.sets.item.children.item collection
func (m *SetItemRequestBuilder) ChildrenById(id string)(*i7c7c17bddaf41293a67c772c6cfde4f761310e7bf1903d604c75b8bdccc5f32e.TermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term%2Did"] = id
    }
    return i7c7c17bddaf41293a67c772c6cfde4f761310e7bf1903d604c75b8bdccc5f32e.NewTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewSetItemRequestBuilderInternal instantiates a new SetItemRequestBuilder and sets the default values.
func NewSetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetItemRequestBuilder) {
    m := &SetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/sites/{site%2Did}/termStore/groups/{group%2Did1}/sets/{set%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSetItemRequestBuilder instantiates a new SetItemRequestBuilder and sets the default values.
func NewSetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property sets for me
func (m *SetItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property sets for me
func (m *SetItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *SetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation all sets under the group in a term [store].
func (m *SetItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration all sets under the group in a term [store].
func (m *SetItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *SetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property sets in me
func (m *SetItemRequestBuilder) CreatePatchRequestInformation(body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Setable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property sets in me
func (m *SetItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Setable, requestConfiguration *SetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property sets for me
func (m *SetItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property sets for me
func (m *SetItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *SetItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get all sets under the group in a term [store].
func (m *SetItemRequestBuilder) Get()(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Setable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler all sets under the group in a term [store].
func (m *SetItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *SetItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Setable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.CreateSetFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Setable), nil
}
// ParentGroup the parentGroup property
func (m *SetItemRequestBuilder) ParentGroup()(*i2e8863714d1e920d7702b4a703714d81cd242645914443f004188922e9fdd7db.ParentGroupRequestBuilder) {
    return i2e8863714d1e920d7702b4a703714d81cd242645914443f004188922e9fdd7db.NewParentGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property sets in me
func (m *SetItemRequestBuilder) Patch(body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Setable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property sets in me
func (m *SetItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Setable, requestConfiguration *SetItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Relations the relations property
func (m *SetItemRequestBuilder) Relations()(*i46a937ae66eb746fc0b61e97303d25716a338af8b2a5d4dc9e7c0e2f36eb7c7a.RelationsRequestBuilder) {
    return i46a937ae66eb746fc0b61e97303d25716a338af8b2a5d4dc9e7c0e2f36eb7c7a.NewRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.sites.item.termStore.groups.item.sets.item.relations.item collection
func (m *SetItemRequestBuilder) RelationsById(id string)(*i3f4638c02945c3b951bb6beb6a994164e45df748d9433b5561e4fa4ec9fb8e11.RelationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation%2Did"] = id
    }
    return i3f4638c02945c3b951bb6beb6a994164e45df748d9433b5561e4fa4ec9fb8e11.NewRelationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Terms the terms property
func (m *SetItemRequestBuilder) Terms()(*i0b982fc6fce1ac2d0c2303097c8efccfbe2a0334914647564fdc68d25e2e8bf5.TermsRequestBuilder) {
    return i0b982fc6fce1ac2d0c2303097c8efccfbe2a0334914647564fdc68d25e2e8bf5.NewTermsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.sites.item.termStore.groups.item.sets.item.terms.item collection
func (m *SetItemRequestBuilder) TermsById(id string)(*id9bf09181f4b938e9dff34744950bbd02803ebc288f203778ba14a909a7eca97.TermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term%2Did"] = id
    }
    return id9bf09181f4b938e9dff34744950bbd02803ebc288f203778ba14a909a7eca97.NewTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
