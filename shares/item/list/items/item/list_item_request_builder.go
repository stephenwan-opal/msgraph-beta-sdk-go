package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0df30f15a37bda4c75a746f1d3a78aeb6d7fedcda8fa91a25630e37a46bb63d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item/createlink"
    i3d8467aa52be83bf5e27b9b74f21a76f1a5c28b249229ca2c40ba746dc31f6cb "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item/analytics"
    i426197eeda24a2ef666da9df0ddff559ab7003700bb784ad8da8009e9521a011 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item/driveitem"
    i5245c95090e9c6e9cde87b4772a00a7124828683c592734b734a62638f6f0df8 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i5ce7a932ecf70db030b8364f7699aa3822b2ace9797a93c63743a24c7b9c40e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item/activities"
    i6a33e6735c94a4eec0295b3339373e3d3d94a57d512c4c1297f7ec69a5bcace5 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item/fields"
    i7e01e46513a14884280655da10ff1bf036a7401d1f1dd3ee9dbcbe02ffe6757e "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item/versions"
    i201a560e0b52f390181eee26b5bb385daeff5fecb17cf5d3d44e3139031471e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item/activities/item"
    iec1083d88bf75a5c5d9ec0cd2e11b051c52addbc2efb1c26db7c0e9b60aab3d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item/versions/item"
)

type ListItemRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type ListItemRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *ListItemRequestBuilder) Activities()(*i5ce7a932ecf70db030b8364f7699aa3822b2ace9797a93c63743a24c7b9c40e2.ActivitiesRequestBuilder) {
    return i5ce7a932ecf70db030b8364f7699aa3822b2ace9797a93c63743a24c7b9c40e2.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) ActivitiesById(id string)(*i201a560e0b52f390181eee26b5bb385daeff5fecb17cf5d3d44e3139031471e5.ItemActivityOLDRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i201a560e0b52f390181eee26b5bb385daeff5fecb17cf5d3d44e3139031471e5.NewItemActivityOLDRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Analytics()(*i3d8467aa52be83bf5e27b9b74f21a76f1a5c28b249229ca2c40ba746dc31f6cb.AnalyticsRequestBuilder) {
    return i3d8467aa52be83bf5e27b9b74f21a76f1a5c28b249229ca2c40ba746dc31f6cb.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/shares/{sharedDriveItem_id}/list/items/{listItem_id}{?select,expand}";
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
func NewListItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ListItemRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
func (m *ListItemRequestBuilder) CreateGetRequestInformation(q func (value *ListItemRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ListItemRequestBuilderGetQueryParameters)
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
func (m *ListItemRequestBuilder) CreateLink()(*i0df30f15a37bda4c75a746f1d3a78aeb6d7fedcda8fa91a25630e37a46bb63d3.CreateLinkRequestBuilder) {
    return i0df30f15a37bda4c75a746f1d3a78aeb6d7fedcda8fa91a25630e37a46bb63d3.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItem, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ListItemRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemRequestBuilder) DriveItem()(*i426197eeda24a2ef666da9df0ddff559ab7003700bb784ad8da8009e9521a011.DriveItemRequestBuilder) {
    return i426197eeda24a2ef666da9df0ddff559ab7003700bb784ad8da8009e9521a011.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Fields()(*i6a33e6735c94a4eec0295b3339373e3d3d94a57d512c4c1297f7ec69a5bcace5.FieldsRequestBuilder) {
    return i6a33e6735c94a4eec0295b3339373e3d3d94a57d512c4c1297f7ec69a5bcace5.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Get(q func (value *ListItemRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItem, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewListItem() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItem), nil
}
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, endDateTime *string, interval *string)(*i5245c95090e9c6e9cde87b4772a00a7124828683c592734b734a62638f6f0df8.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i5245c95090e9c6e9cde87b4772a00a7124828683c592734b734a62638f6f0df8.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime, interval);
}
func (m *ListItemRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItem, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *ListItemRequestBuilder) Versions()(*i7e01e46513a14884280655da10ff1bf036a7401d1f1dd3ee9dbcbe02ffe6757e.VersionsRequestBuilder) {
    return i7e01e46513a14884280655da10ff1bf036a7401d1f1dd3ee9dbcbe02ffe6757e.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) VersionsById(id string)(*iec1083d88bf75a5c5d9ec0cd2e11b051c52addbc2efb1c26db7c0e9b60aab3d7.ListItemVersionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return iec1083d88bf75a5c5d9ec0cd2e11b051c52addbc2efb1c26db7c0e9b60aab3d7.NewListItemVersionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}