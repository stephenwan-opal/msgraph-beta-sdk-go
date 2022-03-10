package root

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i348cfa40d80f2bd346049a604b61e3bbf99a86eaaac1524d02c1cc02ffd86851 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/permissions"
    i4b7ee927e3c2dc792ca09cddc3d0981a89253866d9f177e70356375494e5bd07 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/activities"
    i50e5dbbc7adce54d30748bc54c16ad058df719691e3dd262034b95e19684d59d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/analytics"
    i658e4b44555250a3d896740bfd9a82dcd8af8ce52910c3bc230569041605a43d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/subscriptions"
    i7b7ef26a39790f7760efc19094f8a199de1cdbd8ffb8f10d4309d54d9d571530 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/thumbnails"
    i8f75eaec56c7ab5659de30f31953e0926f050150805e53382e5637b9579e4ea1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/listitem"
    iac52cab0ab0db188a977bc6178beb827af0c6e099c6a5593a0503e97ca41c6a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/content"
    ibff0b4d33209b315966bf59f99fe1478d7383d45cbb30624ed8c4be22ddcff85 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/children"
    id8d5a319b88e388b070ce15327e1b6ba8751257d8c613f3e363927fa1df9b30b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/versions"
    i470c54a9bee9a8ac27ba6d46057802b98035f174bfac53a07d4fb571fa8b4fa9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/thumbnails/item"
    i4873d5650fb447ccd691622a251102eac60f90f73cee1bc40ea7b8a45fe10739 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/activities/item"
    i7c223ca62e6930463b88db6f01fae288e09d61d0e65f94e1bf9b2d7560c3a068 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/children/item"
    i7d632e25daa2136a74aff02a73d2cfc6ab090eca471ab318343c51239e330360 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/versions/item"
    i8aa83421f1959464fd1940ca85cb28228718e593552c3ae1909ed33b6cf13631 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/subscriptions/item"
    ibf85f2aba96b5b04085e8e4e6bc260e595f161ebb0aa75654b03be699f981bd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/root/permissions/item"
)

// RootRequestBuilder provides operations to manage the root property of the microsoft.graph.drive entity.
type RootRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RootRequestBuilderDeleteOptions options for Delete
type RootRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RootRequestBuilderGetOptions options for Get
type RootRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RootRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RootRequestBuilderGetQueryParameters the root folder of the drive. Read-only.
type RootRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// RootRequestBuilderPatchOptions options for Patch
type RootRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DriveItemable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *RootRequestBuilder) Activities()(*i4b7ee927e3c2dc792ca09cddc3d0981a89253866d9f177e70356375494e5bd07.ActivitiesRequestBuilder) {
    return i4b7ee927e3c2dc792ca09cddc3d0981a89253866d9f177e70356375494e5bd07.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.root.activities.item collection
func (m *RootRequestBuilder) ActivitiesById(id string)(*i4873d5650fb447ccd691622a251102eac60f90f73cee1bc40ea7b8a45fe10739.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i4873d5650fb447ccd691622a251102eac60f90f73cee1bc40ea7b8a45fe10739.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Analytics()(*i50e5dbbc7adce54d30748bc54c16ad058df719691e3dd262034b95e19684d59d.AnalyticsRequestBuilder) {
    return i50e5dbbc7adce54d30748bc54c16ad058df719691e3dd262034b95e19684d59d.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *RootRequestBuilder) Children()(*ibff0b4d33209b315966bf59f99fe1478d7383d45cbb30624ed8c4be22ddcff85.ChildrenRequestBuilder) {
    return ibff0b4d33209b315966bf59f99fe1478d7383d45cbb30624ed8c4be22ddcff85.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.root.children.item collection
func (m *RootRequestBuilder) ChildrenById(id string)(*i7c223ca62e6930463b88db6f01fae288e09d61d0e65f94e1bf9b2d7560c3a068.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i7c223ca62e6930463b88db6f01fae288e09d61d0e65f94e1bf9b2d7560c3a068.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/drives/{drive_id}/root{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRootRequestBuilder instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RootRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRootRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *RootRequestBuilder) Content()(*iac52cab0ab0db188a977bc6178beb827af0c6e099c6a5593a0503e97ca41c6a9.ContentRequestBuilder) {
    return iac52cab0ab0db188a977bc6178beb827af0c6e099c6a5593a0503e97ca41c6a9.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for users
func (m *RootRequestBuilder) CreateDeleteRequestInformation(options *RootRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the root folder of the drive. Read-only.
func (m *RootRequestBuilder) CreateGetRequestInformation(options *RootRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property root in users
func (m *RootRequestBuilder) CreatePatchRequestInformation(options *RootRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property root for users
func (m *RootRequestBuilder) Delete(options *RootRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the root folder of the drive. Read-only.
func (m *RootRequestBuilder) Get(options *RootRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDriveItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DriveItemable), nil
}
func (m *RootRequestBuilder) ListItem()(*i8f75eaec56c7ab5659de30f31953e0926f050150805e53382e5637b9579e4ea1.ListItemRequestBuilder) {
    return i8f75eaec56c7ab5659de30f31953e0926f050150805e53382e5637b9579e4ea1.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in users
func (m *RootRequestBuilder) Patch(options *RootRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *RootRequestBuilder) Permissions()(*i348cfa40d80f2bd346049a604b61e3bbf99a86eaaac1524d02c1cc02ffd86851.PermissionsRequestBuilder) {
    return i348cfa40d80f2bd346049a604b61e3bbf99a86eaaac1524d02c1cc02ffd86851.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.root.permissions.item collection
func (m *RootRequestBuilder) PermissionsById(id string)(*ibf85f2aba96b5b04085e8e4e6bc260e595f161ebb0aa75654b03be699f981bd6.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return ibf85f2aba96b5b04085e8e4e6bc260e595f161ebb0aa75654b03be699f981bd6.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Subscriptions()(*i658e4b44555250a3d896740bfd9a82dcd8af8ce52910c3bc230569041605a43d.SubscriptionsRequestBuilder) {
    return i658e4b44555250a3d896740bfd9a82dcd8af8ce52910c3bc230569041605a43d.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.root.subscriptions.item collection
func (m *RootRequestBuilder) SubscriptionsById(id string)(*i8aa83421f1959464fd1940ca85cb28228718e593552c3ae1909ed33b6cf13631.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i8aa83421f1959464fd1940ca85cb28228718e593552c3ae1909ed33b6cf13631.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Thumbnails()(*i7b7ef26a39790f7760efc19094f8a199de1cdbd8ffb8f10d4309d54d9d571530.ThumbnailsRequestBuilder) {
    return i7b7ef26a39790f7760efc19094f8a199de1cdbd8ffb8f10d4309d54d9d571530.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.root.thumbnails.item collection
func (m *RootRequestBuilder) ThumbnailsById(id string)(*i470c54a9bee9a8ac27ba6d46057802b98035f174bfac53a07d4fb571fa8b4fa9.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i470c54a9bee9a8ac27ba6d46057802b98035f174bfac53a07d4fb571fa8b4fa9.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Versions()(*id8d5a319b88e388b070ce15327e1b6ba8751257d8c613f3e363927fa1df9b30b.VersionsRequestBuilder) {
    return id8d5a319b88e388b070ce15327e1b6ba8751257d8c613f3e363927fa1df9b30b.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.root.versions.item collection
func (m *RootRequestBuilder) VersionsById(id string)(*i7d632e25daa2136a74aff02a73d2cfc6ab090eca471ab318343c51239e330360.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i7d632e25daa2136a74aff02a73d2cfc6ab090eca471ab318343c51239e330360.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}