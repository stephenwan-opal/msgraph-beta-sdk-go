package msgraphbetasdkgo

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i79f2b866e8bec3ee9349dc885ecb3691e94b20459995d83b3dbf9f05341c7a89 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a "github.com/microsoftgraph/msgraph-beta-sdk-go/users"
)

// GraphBaseServiceClient the main entry point of the SDK, exposes the configuration and the fluent API.
type GraphBaseServiceClient struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewGraphBaseServiceClient instantiates a new GraphBaseServiceClient and sets the default values.
func NewGraphBaseServiceClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GraphBaseServiceClient) {
    m := &GraphBaseServiceClient{
    }
    m.pathParameters = make(map[string]string);
    m.urlTemplate = "{+baseurl}";
    m.requestAdapter = requestAdapter;
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.requestAdapter.GetBaseUrl() == "" {
        m.requestAdapter.SetBaseUrl("https://graph.microsoft.com/beta")
    }
    return m
}
// Groups provides operations to manage the collection of group entities.
func (m *GraphBaseServiceClient) Groups()(*i79f2b866e8bec3ee9349dc885ecb3691e94b20459995d83b3dbf9f05341c7a89.GroupsRequestBuilder) {
    return i79f2b866e8bec3ee9349dc885ecb3691e94b20459995d83b3dbf9f05341c7a89.NewGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupsById provides operations to manage the collection of group entities.
func (m *GraphBaseServiceClient) GroupsById(id string)(*i79f2b866e8bec3ee9349dc885ecb3691e94b20459995d83b3dbf9f05341c7a89.GroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["group%2Did"] = id
    }
    return i79f2b866e8bec3ee9349dc885ecb3691e94b20459995d83b3dbf9f05341c7a89.NewGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Users provides operations to manage the collection of user entities.
func (m *GraphBaseServiceClient) Users()(*icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a.UsersRequestBuilder) {
    return icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a.NewUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsersById provides operations to manage the collection of user entities.
func (m *GraphBaseServiceClient) UsersById(id string)(*icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a.UserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["user%2Did"] = id
    }
    return icd01c84a90833c55ac2309fd7034cb1962c60f59eb1ee2b2cf7b04c708402b6a.NewUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
