package drive

import (
    i192e552808d72c666fa89ea5e18d213d8060951b798ab506f46a9f7ffdf83e2a "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items"
    i1cde53b5097fe4f2cb7d9aa6a1d6263ca7dc45f359a394f6451d31a82baa66f1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/searchwithq"
    i2f3fa4e2e4547c94b77043edcae61480a38269272b0e62403050580a8364f5c9 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/sharedwithme"
    i3dba79c08f5528ca1f738cd0ec4d5c0b89913bea342b964174e72bc175dec1d4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/recent"
    i4e0de778e26a386528f38fa1e34ccca7e9004fb8a097e4795502e4f91e7913ba "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/following"
    i8195f82296fd049d853cdd9b0d96130049f0e63bf935143dc8aed2b4066352f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/special"
    i89fcd6e566145dd6fd7ea2956e5aefd9030e1a74e3028ab21e1cac08c5d17ead "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/activities"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ie75e0ac9d6da32720a3dd716caabe531bea74aa98e60d982dabe2088718e1e8d "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/root"
    iee45bfc2e2e1ccdf052d488c498a03424a7325bcbac25a672e0946a3251ff078 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list"
    iee7675fea869eb6434c7377c309bb1287126402b996d180fde09843483ab5b0d "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i50d8648f64c4421d77d91aa220edfff9d5ba824309e53cbcc3025b0592285332 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/following/item"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5d9fbdca07c6e20b1cc99de5efe021bbc1fb6e6b9aec7b3e289fa7d91f62974d "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/special/item"
    i75b317198811cf16709f6ce43e8700a26a56733bc41233a94a23920a7de396f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/activities/item"
    i7ba1848b91a2a586b2d0cf4042a70732acf64fd1f43dec2814281ec5e078edc1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/items/item"
    id087d332b06d6a2e3560c5cbdd124c635b06366297ad2046147358062b8678af "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/bundles/item"
)

type DriveRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type DriveRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *DriveRequestBuilder) Activities()(*i89fcd6e566145dd6fd7ea2956e5aefd9030e1a74e3028ab21e1cac08c5d17ead.ActivitiesRequestBuilder) {
    return i89fcd6e566145dd6fd7ea2956e5aefd9030e1a74e3028ab21e1cac08c5d17ead.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) ActivitiesById(id string)(*i75b317198811cf16709f6ce43e8700a26a56733bc41233a94a23920a7de396f5.ItemActivityOLDRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i75b317198811cf16709f6ce43e8700a26a56733bc41233a94a23920a7de396f5.NewItemActivityOLDRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveRequestBuilder) Bundles()(*iee7675fea869eb6434c7377c309bb1287126402b996d180fde09843483ab5b0d.BundlesRequestBuilder) {
    return iee7675fea869eb6434c7377c309bb1287126402b996d180fde09843483ab5b0d.NewBundlesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) BundlesById(id string)(*id087d332b06d6a2e3560c5cbdd124c635b06366297ad2046147358062b8678af.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return id087d332b06d6a2e3560c5cbdd124c635b06366297ad2046147358062b8678af.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewDriveRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveRequestBuilder) {
    m := &DriveRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/drive{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewDriveRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DriveRequestBuilder) CreateGetRequestInformation(q func (value *DriveRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(DriveRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DriveRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Drive, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *DriveRequestBuilder) Following()(*i4e0de778e26a386528f38fa1e34ccca7e9004fb8a097e4795502e4f91e7913ba.FollowingRequestBuilder) {
    return i4e0de778e26a386528f38fa1e34ccca7e9004fb8a097e4795502e4f91e7913ba.NewFollowingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) FollowingById(id string)(*i50d8648f64c4421d77d91aa220edfff9d5ba824309e53cbcc3025b0592285332.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i50d8648f64c4421d77d91aa220edfff9d5ba824309e53cbcc3025b0592285332.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveRequestBuilder) Get(q func (value *DriveRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Drive, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDrive() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Drive), nil
}
func (m *DriveRequestBuilder) Items()(*i192e552808d72c666fa89ea5e18d213d8060951b798ab506f46a9f7ffdf83e2a.ItemsRequestBuilder) {
    return i192e552808d72c666fa89ea5e18d213d8060951b798ab506f46a9f7ffdf83e2a.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) ItemsById(id string)(*i7ba1848b91a2a586b2d0cf4042a70732acf64fd1f43dec2814281ec5e078edc1.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i7ba1848b91a2a586b2d0cf4042a70732acf64fd1f43dec2814281ec5e078edc1.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveRequestBuilder) List()(*iee45bfc2e2e1ccdf052d488c498a03424a7325bcbac25a672e0946a3251ff078.ListRequestBuilder) {
    return iee45bfc2e2e1ccdf052d488c498a03424a7325bcbac25a672e0946a3251ff078.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Drive, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *DriveRequestBuilder) Recent()(*i3dba79c08f5528ca1f738cd0ec4d5c0b89913bea342b964174e72bc175dec1d4.RecentRequestBuilder) {
    return i3dba79c08f5528ca1f738cd0ec4d5c0b89913bea342b964174e72bc175dec1d4.NewRecentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) Root()(*ie75e0ac9d6da32720a3dd716caabe531bea74aa98e60d982dabe2088718e1e8d.RootRequestBuilder) {
    return ie75e0ac9d6da32720a3dd716caabe531bea74aa98e60d982dabe2088718e1e8d.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) SearchWithQ(q *string)(*i1cde53b5097fe4f2cb7d9aa6a1d6263ca7dc45f359a394f6451d31a82baa66f1.SearchWithQRequestBuilder) {
    return i1cde53b5097fe4f2cb7d9aa6a1d6263ca7dc45f359a394f6451d31a82baa66f1.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
func (m *DriveRequestBuilder) SharedWithMe()(*i2f3fa4e2e4547c94b77043edcae61480a38269272b0e62403050580a8364f5c9.SharedWithMeRequestBuilder) {
    return i2f3fa4e2e4547c94b77043edcae61480a38269272b0e62403050580a8364f5c9.NewSharedWithMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) Special()(*i8195f82296fd049d853cdd9b0d96130049f0e63bf935143dc8aed2b4066352f6.SpecialRequestBuilder) {
    return i8195f82296fd049d853cdd9b0d96130049f0e63bf935143dc8aed2b4066352f6.NewSpecialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) SpecialById(id string)(*i5d9fbdca07c6e20b1cc99de5efe021bbc1fb6e6b9aec7b3e289fa7d91f62974d.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i5d9fbdca07c6e20b1cc99de5efe021bbc1fb6e6b9aec7b3e289fa7d91f62974d.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}