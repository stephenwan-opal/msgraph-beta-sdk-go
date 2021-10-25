package me

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i066e6c0b7a628617371533e62ce3ebdd3bb41009a8fbfe0051aebeee5a091763 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments"
    i1197007b396c8142142d808832d472fa4a0ca1796dcda8056a9c759ced2a4005 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/schools"
    i1b5284c067c2acfcb9ec377599aa90647d0e8c253bdeac21d56033eae704633e "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/rubrics"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5f1f4bf3151bbb3488a322f91d88a1e3277a2a2e390a895db3a6a5a0dafd3bd0 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/user"
    i5f4f009fc61fc7227f7443b57e61741a07864a630e39eba9bf7e097c583915e3 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/taughtclasses"
    ia9901f126435d74fc86b3bf99ec52d925f16943c99de2d04f1f3f010f4c84fbc "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/classes"
    i719722077dd2c46836e8fc87facd02a34341429e3162704b105c41c28deda2b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/rubrics/item"
    ieba30bd3da2220b37fa7fcf6200b1473597f54cc0733c3310394e58e26492403 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/me/assignments/item"
)

type MeRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type MeRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *MeRequestBuilder) Assignments()(*i066e6c0b7a628617371533e62ce3ebdd3bb41009a8fbfe0051aebeee5a091763.AssignmentsRequestBuilder) {
    return i066e6c0b7a628617371533e62ce3ebdd3bb41009a8fbfe0051aebeee5a091763.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) AssignmentsById(id string)(*ieba30bd3da2220b37fa7fcf6200b1473597f54cc0733c3310394e58e26492403.EducationAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationAssignment_id"] = id
    }
    return ieba30bd3da2220b37fa7fcf6200b1473597f54cc0733c3310394e58e26492403.NewEducationAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) Classes()(*ia9901f126435d74fc86b3bf99ec52d925f16943c99de2d04f1f3f010f4c84fbc.ClassesRequestBuilder) {
    return ia9901f126435d74fc86b3bf99ec52d925f16943c99de2d04f1f3f010f4c84fbc.NewClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewMeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MeRequestBuilder) {
    m := &MeRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/education/me{?select,expand}";
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
func NewMeRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MeRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *MeRequestBuilder) CreateGetRequestInformation(q func (value *MeRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(MeRequestBuilderGetQueryParameters)
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
func (m *MeRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationUser, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *MeRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *MeRequestBuilder) Get(q func (value *MeRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationUser, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEducationUser() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationUser), nil
}
func (m *MeRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationUser, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *MeRequestBuilder) Rubrics()(*i1b5284c067c2acfcb9ec377599aa90647d0e8c253bdeac21d56033eae704633e.RubricsRequestBuilder) {
    return i1b5284c067c2acfcb9ec377599aa90647d0e8c253bdeac21d56033eae704633e.NewRubricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) RubricsById(id string)(*i719722077dd2c46836e8fc87facd02a34341429e3162704b105c41c28deda2b9.EducationRubricRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationRubric_id"] = id
    }
    return i719722077dd2c46836e8fc87facd02a34341429e3162704b105c41c28deda2b9.NewEducationRubricRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) Schools()(*i1197007b396c8142142d808832d472fa4a0ca1796dcda8056a9c759ced2a4005.SchoolsRequestBuilder) {
    return i1197007b396c8142142d808832d472fa4a0ca1796dcda8056a9c759ced2a4005.NewSchoolsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) TaughtClasses()(*i5f4f009fc61fc7227f7443b57e61741a07864a630e39eba9bf7e097c583915e3.TaughtClassesRequestBuilder) {
    return i5f4f009fc61fc7227f7443b57e61741a07864a630e39eba9bf7e097c583915e3.NewTaughtClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *MeRequestBuilder) User()(*i5f1f4bf3151bbb3488a322f91d88a1e3277a2a2e390a895db3a6a5a0dafd3bd0.UserRequestBuilder) {
    return i5f1f4bf3151bbb3488a322f91d88a1e3277a2a2e390a895db3a6a5a0dafd3bd0.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}