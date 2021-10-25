package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i01322f6c445fc8cfdd645f8a58a2e9dc2aae2538661d3565a27837d1aa1a3bf3 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item/analytics"
    i0d9fa4e07c80bd530f75b82be901e491731fc66da0aea17d127c28e0b96cdd3f "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item/createlink"
    i70f1efce412deb28482396740c683ea9f4b411c123301ce44883bab61b141a65 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item/versions"
    i71f902ff67b40c4be62d504c1d2c0f525058bba7aa6c03f0ad5fea4098c14d1f "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item/driveitem"
    i9237c91883a88a40d038ddbc99b525375d2ffafd6ef05bfcdf2ce7a59c7b967d "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item/activities"
    ia0af2bad6bfa2334cda9b36762c7754513aad438f8dad54fe36023b80b1086e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item/fields"
    ia8cb9d287c711158fa129209470a23d78a03e7501d70219257f78200123e41e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i2b138b210334df4f8b6453d94f09db4128ad42a5036fede3e837fa332b533270 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item/versions/item"
    ieddbfd259f1fbc5f85a15691baa1cba1c9fc88a0b59e0d161ce0c97e208a77bc "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item/activities/item"
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
func (m *ListItemRequestBuilder) Activities()(*i9237c91883a88a40d038ddbc99b525375d2ffafd6ef05bfcdf2ce7a59c7b967d.ActivitiesRequestBuilder) {
    return i9237c91883a88a40d038ddbc99b525375d2ffafd6ef05bfcdf2ce7a59c7b967d.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) ActivitiesById(id string)(*ieddbfd259f1fbc5f85a15691baa1cba1c9fc88a0b59e0d161ce0c97e208a77bc.ItemActivityOLDRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return ieddbfd259f1fbc5f85a15691baa1cba1c9fc88a0b59e0d161ce0c97e208a77bc.NewItemActivityOLDRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Analytics()(*i01322f6c445fc8cfdd645f8a58a2e9dc2aae2538661d3565a27837d1aa1a3bf3.AnalyticsRequestBuilder) {
    return i01322f6c445fc8cfdd645f8a58a2e9dc2aae2538661d3565a27837d1aa1a3bf3.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/sites/{site_id}/lists/{list_id}/items/{listItem_id}{?select,expand}";
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
func (m *ListItemRequestBuilder) CreateLink()(*i0d9fa4e07c80bd530f75b82be901e491731fc66da0aea17d127c28e0b96cdd3f.CreateLinkRequestBuilder) {
    return i0d9fa4e07c80bd530f75b82be901e491731fc66da0aea17d127c28e0b96cdd3f.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ListItemRequestBuilder) DriveItem()(*i71f902ff67b40c4be62d504c1d2c0f525058bba7aa6c03f0ad5fea4098c14d1f.DriveItemRequestBuilder) {
    return i71f902ff67b40c4be62d504c1d2c0f525058bba7aa6c03f0ad5fea4098c14d1f.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Fields()(*ia0af2bad6bfa2334cda9b36762c7754513aad438f8dad54fe36023b80b1086e4.FieldsRequestBuilder) {
    return ia0af2bad6bfa2334cda9b36762c7754513aad438f8dad54fe36023b80b1086e4.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, endDateTime *string, interval *string)(*ia8cb9d287c711158fa129209470a23d78a03e7501d70219257f78200123e41e5.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return ia8cb9d287c711158fa129209470a23d78a03e7501d70219257f78200123e41e5.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime, interval);
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
func (m *ListItemRequestBuilder) Versions()(*i70f1efce412deb28482396740c683ea9f4b411c123301ce44883bab61b141a65.VersionsRequestBuilder) {
    return i70f1efce412deb28482396740c683ea9f4b411c123301ce44883bab61b141a65.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) VersionsById(id string)(*i2b138b210334df4f8b6453d94f09db4128ad42a5036fede3e837fa332b533270.ListItemVersionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return i2b138b210334df4f8b6453d94f09db4128ad42a5036fede3e837fa332b533270.NewListItemVersionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}