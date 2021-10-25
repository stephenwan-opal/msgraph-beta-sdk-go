package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1928a13dc54c73d5c9d9c1da4d9b70aeed20716bc8d9c77536a95f1b2bc0a66c "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/resources"
    i2d5d8db0d4884b3dbd8ece33930e5eb637d4a721b2a8a2c09e8aca02014b6af2 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/return_escpaped"
    i43ba3ee4b1fe518a6d013e2f0b889058897248ca13cf8fd56f3a4e3f6dbcf726 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/outcomes"
    i63c795cc15fa99fd56c532c984b57d606b0ce50ae0f0462a8f05f2c97403505a "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/submittedresources"
    i768987a34c53877b80bd9b75ccfde8cec3078a7b79600a3541157f92a02bd7ae "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/submit"
    iad0e46857ffe4aca71c347ec356f1f060d1acbfe420d89db1d61dd5f39ae6479 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/setupresourcesfolder"
    ie5fd4dbd2c6ed61da731d9780ed8f1ec6a8a6533c73bce16ca0f6e33a6b8cf4d "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/unsubmit"
    i6d2213e16b46c89802a19dd596bafab2a2de02beb8b6102d680539d7de269026 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/resources/item"
    i96ff390efa3e1cf5f1c7b76015acbcac008ccfe65211d104227d3786192347c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/outcomes/item"
    ic22ac8d3bfbbeab72b02b700408b7d279cdca486676844964c112e35bc79e358 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/classes/item/assignments/item/submissions/item/submittedresources/item"
)

type EducationSubmissionRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type EducationSubmissionRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func NewEducationSubmissionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationSubmissionRequestBuilder) {
    m := &EducationSubmissionRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/education/classes/{educationClass_id}/assignments/{educationAssignment_id}/submissions/{educationSubmission_id}{?select,expand}";
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
func NewEducationSubmissionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationSubmissionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationSubmissionRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *EducationSubmissionRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EducationSubmissionRequestBuilder) CreateGetRequestInformation(q func (value *EducationSubmissionRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(EducationSubmissionRequestBuilderGetQueryParameters)
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
func (m *EducationSubmissionRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSubmission, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EducationSubmissionRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EducationSubmissionRequestBuilder) Get(q func (value *EducationSubmissionRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSubmission, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEducationSubmission() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSubmission), nil
}
func (m *EducationSubmissionRequestBuilder) Outcomes()(*i43ba3ee4b1fe518a6d013e2f0b889058897248ca13cf8fd56f3a4e3f6dbcf726.OutcomesRequestBuilder) {
    return i43ba3ee4b1fe518a6d013e2f0b889058897248ca13cf8fd56f3a4e3f6dbcf726.NewOutcomesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) OutcomesById(id string)(*i96ff390efa3e1cf5f1c7b76015acbcac008ccfe65211d104227d3786192347c4.EducationOutcomeRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationOutcome_id"] = id
    }
    return i96ff390efa3e1cf5f1c7b76015acbcac008ccfe65211d104227d3786192347c4.NewEducationOutcomeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSubmission, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EducationSubmissionRequestBuilder) Resources()(*i1928a13dc54c73d5c9d9c1da4d9b70aeed20716bc8d9c77536a95f1b2bc0a66c.ResourcesRequestBuilder) {
    return i1928a13dc54c73d5c9d9c1da4d9b70aeed20716bc8d9c77536a95f1b2bc0a66c.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) ResourcesById(id string)(*i6d2213e16b46c89802a19dd596bafab2a2de02beb8b6102d680539d7de269026.EducationSubmissionResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationSubmissionResource_id"] = id
    }
    return i6d2213e16b46c89802a19dd596bafab2a2de02beb8b6102d680539d7de269026.NewEducationSubmissionResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Return_escpaped()(*i2d5d8db0d4884b3dbd8ece33930e5eb637d4a721b2a8a2c09e8aca02014b6af2.ReturnRequestBuilder) {
    return i2d5d8db0d4884b3dbd8ece33930e5eb637d4a721b2a8a2c09e8aca02014b6af2.NewReturnRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) SetUpResourcesFolder()(*iad0e46857ffe4aca71c347ec356f1f060d1acbfe420d89db1d61dd5f39ae6479.SetUpResourcesFolderRequestBuilder) {
    return iad0e46857ffe4aca71c347ec356f1f060d1acbfe420d89db1d61dd5f39ae6479.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Submit()(*i768987a34c53877b80bd9b75ccfde8cec3078a7b79600a3541157f92a02bd7ae.SubmitRequestBuilder) {
    return i768987a34c53877b80bd9b75ccfde8cec3078a7b79600a3541157f92a02bd7ae.NewSubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) SubmittedResources()(*i63c795cc15fa99fd56c532c984b57d606b0ce50ae0f0462a8f05f2c97403505a.SubmittedResourcesRequestBuilder) {
    return i63c795cc15fa99fd56c532c984b57d606b0ce50ae0f0462a8f05f2c97403505a.NewSubmittedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) SubmittedResourcesById(id string)(*ic22ac8d3bfbbeab72b02b700408b7d279cdca486676844964c112e35bc79e358.EducationSubmissionResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["educationSubmissionResource_id"] = id
    }
    return ic22ac8d3bfbbeab72b02b700408b7d279cdca486676844964c112e35bc79e358.NewEducationSubmissionResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationSubmissionRequestBuilder) Unsubmit()(*ie5fd4dbd2c6ed61da731d9780ed8f1ec6a8a6533c73bce16ca0f6e33a6b8cf4d.UnsubmitRequestBuilder) {
    return ie5fd4dbd2c6ed61da731d9780ed8f1ec6a8a6533c73bce16ca0f6e33a6b8cf4d.NewUnsubmitRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}