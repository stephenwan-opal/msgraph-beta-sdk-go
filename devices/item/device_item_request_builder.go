package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i0830d73c05d0cd2e685461a807491311731aedd53e2fe92794ce58ed2cf68cb3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/getmemberobjects"
    i08e90e59dc1f31b2ec2b796cb6007a6d5b03f5b99baeb73d4fb6c3c73e9807cc "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/transitivememberof"
    i0a12eca11799833eefc611cf2bc5d081a6d4ca1f0077140fc3f7859ef5c7c862 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/checkmembergroups"
    i0ff7a5160416ceed56e16358c8af0f5a8032f984461e4e97dde1db62f7207fe4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/commands"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i552a6f26c390b8c7675f028d3a98d5ef0beebdb3ee31db85eea40825ede88db4 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/getmembergroups"
    i8cc492560989153387ef40098b2cfd2301510578bd7681f8a4c9af9aa911f0c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/memberof"
    ic0e516d9ef79746d4d17a660f00104ef0dac16cfa359979aa0dbf10dcce55f06 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/checkmemberobjects"
    ic2da6ff90c2f6b8ecb849d8a1362794682f6d7cfee5f9a38119f672e88e1fefc "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/restore"
    id5a79299b989d24548d743c765056fa70097eae5165493f3dc78187dd251aa93 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/registeredusers"
    ie3938ac8c4c7db4930a30ba4c91d7ca7a7d9d6cac8a9274c997a3420e420b4dc "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/extensions"
    iec69b6c64deeabcd1d010898cbfb5e8a0ca669911e712648a2f05592e7396043 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/registeredowners"
    if010828737e171d055e725c144a741fbeda7f194b68cc132db59f8a94634fc8b "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/usagerights"
    i17b2e41daecab12778ef8849a52a0050afea7ab215b7a063c8d145b656ff350c "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/transitivememberof/item"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i4e492465a29843eed16f5f1c1aed7913b386ec884cccbb7bfed560c3c6f30f75 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/usagerights/item"
    ic49a80ae4774e504e1e936eaa7b3278aa408c882b3991bc5bc8969383c4d4d12 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/commands/item"
    ic5b79b8426b0398a1d37364aa2476ccd846f17a8e2130477d3235a55b6728c2f "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/registeredowners/item"
    id7b0a8873e3dc9a1070cd7d3a137f0ed2968cdf066a85a0810adb726fef41b0e "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/registeredusers/item"
    if0991c114c37c9c6d2fde813890c596fa6c31a4b084d49dd6eb0e8c33c7f8af6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/memberof/item"
    if0a4d871ad7cc2e496e78b2da0daa8ea59cf84cc6bada0cf4faf2f4fa43eb2f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devices/item/extensions/item"
)

// DeviceItemRequestBuilder provides operations to manage the collection of device entities.
type DeviceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceItemRequestBuilderDeleteOptions options for Delete
type DeviceItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceItemRequestBuilderGetOptions options for Get
type DeviceItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceItemRequestBuilderGetQueryParameters get entity from devices by key
type DeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceItemRequestBuilderPatchOptions options for Patch
type DeviceItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Deviceable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceItemRequestBuilder) CheckMemberGroups()(*i0a12eca11799833eefc611cf2bc5d081a6d4ca1f0077140fc3f7859ef5c7c862.CheckMemberGroupsRequestBuilder) {
    return i0a12eca11799833eefc611cf2bc5d081a6d4ca1f0077140fc3f7859ef5c7c862.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceItemRequestBuilder) CheckMemberObjects()(*ic0e516d9ef79746d4d17a660f00104ef0dac16cfa359979aa0dbf10dcce55f06.CheckMemberObjectsRequestBuilder) {
    return ic0e516d9ef79746d4d17a660f00104ef0dac16cfa359979aa0dbf10dcce55f06.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceItemRequestBuilder) Commands()(*i0ff7a5160416ceed56e16358c8af0f5a8032f984461e4e97dde1db62f7207fe4.CommandsRequestBuilder) {
    return i0ff7a5160416ceed56e16358c8af0f5a8032f984461e4e97dde1db62f7207fe4.NewCommandsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CommandsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.devices.item.commands.item collection
func (m *DeviceItemRequestBuilder) CommandsById(id string)(*ic49a80ae4774e504e1e936eaa7b3278aa408c882b3991bc5bc8969383c4d4d12.CommandItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["command_id"] = id
    }
    return ic49a80ae4774e504e1e936eaa7b3278aa408c882b3991bc5bc8969383c4d4d12.NewCommandItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceItemRequestBuilderInternal instantiates a new DeviceItemRequestBuilder and sets the default values.
func NewDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceItemRequestBuilder) {
    m := &DeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/devices/{device_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceItemRequestBuilder instantiates a new DeviceItemRequestBuilder and sets the default values.
func NewDeviceItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from devices
func (m *DeviceItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from devices by key
func (m *DeviceItemRequestBuilder) CreateGetRequestInformation(options *DeviceItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in devices
func (m *DeviceItemRequestBuilder) CreatePatchRequestInformation(options *DeviceItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from devices
func (m *DeviceItemRequestBuilder) Delete(options *DeviceItemRequestBuilderDeleteOptions)(error) {
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
func (m *DeviceItemRequestBuilder) Extensions()(*ie3938ac8c4c7db4930a30ba4c91d7ca7a7d9d6cac8a9274c997a3420e420b4dc.ExtensionsRequestBuilder) {
    return ie3938ac8c4c7db4930a30ba4c91d7ca7a7d9d6cac8a9274c997a3420e420b4dc.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.devices.item.extensions.item collection
func (m *DeviceItemRequestBuilder) ExtensionsById(id string)(*if0a4d871ad7cc2e496e78b2da0daa8ea59cf84cc6bada0cf4faf2f4fa43eb2f1.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return if0a4d871ad7cc2e496e78b2da0daa8ea59cf84cc6bada0cf4faf2f4fa43eb2f1.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get entity from devices by key
func (m *DeviceItemRequestBuilder) Get(options *DeviceItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Deviceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDeviceFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Deviceable), nil
}
func (m *DeviceItemRequestBuilder) GetMemberGroups()(*i552a6f26c390b8c7675f028d3a98d5ef0beebdb3ee31db85eea40825ede88db4.GetMemberGroupsRequestBuilder) {
    return i552a6f26c390b8c7675f028d3a98d5ef0beebdb3ee31db85eea40825ede88db4.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceItemRequestBuilder) GetMemberObjects()(*i0830d73c05d0cd2e685461a807491311731aedd53e2fe92794ce58ed2cf68cb3.GetMemberObjectsRequestBuilder) {
    return i0830d73c05d0cd2e685461a807491311731aedd53e2fe92794ce58ed2cf68cb3.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceItemRequestBuilder) MemberOf()(*i8cc492560989153387ef40098b2cfd2301510578bd7681f8a4c9af9aa911f0c1.MemberOfRequestBuilder) {
    return i8cc492560989153387ef40098b2cfd2301510578bd7681f8a4c9af9aa911f0c1.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.devices.item.memberOf.item collection
func (m *DeviceItemRequestBuilder) MemberOfById(id string)(*if0991c114c37c9c6d2fde813890c596fa6c31a4b084d49dd6eb0e8c33c7f8af6.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return if0991c114c37c9c6d2fde813890c596fa6c31a4b084d49dd6eb0e8c33c7f8af6.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in devices
func (m *DeviceItemRequestBuilder) Patch(options *DeviceItemRequestBuilderPatchOptions)(error) {
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
func (m *DeviceItemRequestBuilder) RegisteredOwners()(*iec69b6c64deeabcd1d010898cbfb5e8a0ca669911e712648a2f05592e7396043.RegisteredOwnersRequestBuilder) {
    return iec69b6c64deeabcd1d010898cbfb5e8a0ca669911e712648a2f05592e7396043.NewRegisteredOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredOwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.devices.item.registeredOwners.item collection
func (m *DeviceItemRequestBuilder) RegisteredOwnersById(id string)(*ic5b79b8426b0398a1d37364aa2476ccd846f17a8e2130477d3235a55b6728c2f.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return ic5b79b8426b0398a1d37364aa2476ccd846f17a8e2130477d3235a55b6728c2f.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceItemRequestBuilder) RegisteredUsers()(*id5a79299b989d24548d743c765056fa70097eae5165493f3dc78187dd251aa93.RegisteredUsersRequestBuilder) {
    return id5a79299b989d24548d743c765056fa70097eae5165493f3dc78187dd251aa93.NewRegisteredUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredUsersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.devices.item.registeredUsers.item collection
func (m *DeviceItemRequestBuilder) RegisteredUsersById(id string)(*id7b0a8873e3dc9a1070cd7d3a137f0ed2968cdf066a85a0810adb726fef41b0e.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return id7b0a8873e3dc9a1070cd7d3a137f0ed2968cdf066a85a0810adb726fef41b0e.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceItemRequestBuilder) Restore()(*ic2da6ff90c2f6b8ecb849d8a1362794682f6d7cfee5f9a38119f672e88e1fefc.RestoreRequestBuilder) {
    return ic2da6ff90c2f6b8ecb849d8a1362794682f6d7cfee5f9a38119f672e88e1fefc.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceItemRequestBuilder) TransitiveMemberOf()(*i08e90e59dc1f31b2ec2b796cb6007a6d5b03f5b99baeb73d4fb6c3c73e9807cc.TransitiveMemberOfRequestBuilder) {
    return i08e90e59dc1f31b2ec2b796cb6007a6d5b03f5b99baeb73d4fb6c3c73e9807cc.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.devices.item.transitiveMemberOf.item collection
func (m *DeviceItemRequestBuilder) TransitiveMemberOfById(id string)(*i17b2e41daecab12778ef8849a52a0050afea7ab215b7a063c8d145b656ff350c.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i17b2e41daecab12778ef8849a52a0050afea7ab215b7a063c8d145b656ff350c.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceItemRequestBuilder) UsageRights()(*if010828737e171d055e725c144a741fbeda7f194b68cc132db59f8a94634fc8b.UsageRightsRequestBuilder) {
    return if010828737e171d055e725c144a741fbeda7f194b68cc132db59f8a94634fc8b.NewUsageRightsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UsageRightsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.devices.item.usageRights.item collection
func (m *DeviceItemRequestBuilder) UsageRightsById(id string)(*i4e492465a29843eed16f5f1c1aed7913b386ec884cccbb7bfed560c3c6f30f75.UsageRightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["usageRight_id"] = id
    }
    return i4e492465a29843eed16f5f1c1aed7913b386ec884cccbb7bfed560c3c6f30f75.NewUsageRightItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}