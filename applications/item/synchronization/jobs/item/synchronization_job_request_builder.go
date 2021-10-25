package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i33749878be357a8a4e7ab7549dfb7ef00a61ae4756862d17fdbe0399ec25776c "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/stop"
    i3f4cdc08d50a52eb36625eb92c9b78c15ed747a07ec2270b3b04bdbeea41d91f "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/provisionondemand"
    i47d716a54a179680138af986e459c792376d7f4c6514002b3e70daf5a1a76fef "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/restart"
    i8400fe465eee24495a8d2b54b414e58dba9e8d403a34045a85f9ef8eb6613f0f "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/pause"
    i9e52a20ec6431f72acddd1559db2e4f0fd21df996b830a7113908ef411982967 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/schema"
    idaec5822ce9a73474ab924b2776d55984f7bb1b0ccd24f511556b91451979237 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/start"
    ifd48d70ab4c59f1f11d77d1818ba8659d5894f3575962284a88e156d2365ee90 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/synchronization/jobs/item/validatecredentials"
)

type SynchronizationJobRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type SynchronizationJobRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func NewSynchronizationJobRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SynchronizationJobRequestBuilder) {
    m := &SynchronizationJobRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/applications/{application_id}/synchronization/jobs/{synchronizationJob_id}{?select,expand}";
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
func NewSynchronizationJobRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SynchronizationJobRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSynchronizationJobRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *SynchronizationJobRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SynchronizationJobRequestBuilder) CreateGetRequestInformation(q func (value *SynchronizationJobRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(SynchronizationJobRequestBuilderGetQueryParameters)
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
func (m *SynchronizationJobRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationJob, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SynchronizationJobRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *SynchronizationJobRequestBuilder) Get(q func (value *SynchronizationJobRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationJob, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSynchronizationJob() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationJob), nil
}
func (m *SynchronizationJobRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SynchronizationJob, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *SynchronizationJobRequestBuilder) Pause()(*i8400fe465eee24495a8d2b54b414e58dba9e8d403a34045a85f9ef8eb6613f0f.PauseRequestBuilder) {
    return i8400fe465eee24495a8d2b54b414e58dba9e8d403a34045a85f9ef8eb6613f0f.NewPauseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationJobRequestBuilder) ProvisionOnDemand()(*i3f4cdc08d50a52eb36625eb92c9b78c15ed747a07ec2270b3b04bdbeea41d91f.ProvisionOnDemandRequestBuilder) {
    return i3f4cdc08d50a52eb36625eb92c9b78c15ed747a07ec2270b3b04bdbeea41d91f.NewProvisionOnDemandRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationJobRequestBuilder) Restart()(*i47d716a54a179680138af986e459c792376d7f4c6514002b3e70daf5a1a76fef.RestartRequestBuilder) {
    return i47d716a54a179680138af986e459c792376d7f4c6514002b3e70daf5a1a76fef.NewRestartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationJobRequestBuilder) Schema()(*i9e52a20ec6431f72acddd1559db2e4f0fd21df996b830a7113908ef411982967.SchemaRequestBuilder) {
    return i9e52a20ec6431f72acddd1559db2e4f0fd21df996b830a7113908ef411982967.NewSchemaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationJobRequestBuilder) Start()(*idaec5822ce9a73474ab924b2776d55984f7bb1b0ccd24f511556b91451979237.StartRequestBuilder) {
    return idaec5822ce9a73474ab924b2776d55984f7bb1b0ccd24f511556b91451979237.NewStartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationJobRequestBuilder) Stop()(*i33749878be357a8a4e7ab7549dfb7ef00a61ae4756862d17fdbe0399ec25776c.StopRequestBuilder) {
    return i33749878be357a8a4e7ab7549dfb7ef00a61ae4756862d17fdbe0399ec25776c.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SynchronizationJobRequestBuilder) ValidateCredentials()(*ifd48d70ab4c59f1f11d77d1818ba8659d5894f3575962284a88e156d2365ee90.ValidateCredentialsRequestBuilder) {
    return ifd48d70ab4c59f1f11d77d1818ba8659d5894f3575962284a88e156d2365ee90.NewValidateCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}