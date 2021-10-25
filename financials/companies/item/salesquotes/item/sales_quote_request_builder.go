package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1a8296a2f13ef3c5ac1362a227b917a25604e9d63082e5b5ecddc5ba5a0cb2ca "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/currency"
    i72ac360010a42d020b08044ba3e52ca3d2232ffd23edf081ecd8320e34798c78 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/paymentterm"
    i753033a2c9ce81017b74127e71bff7620415ec40c37877a3a9156a8a870a74e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/customer"
    i80936828602bda4666c4eced48786bc2c1c175251ae1bdf1e0ef93378a7b543b "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/makeinvoice"
    if48d4a5b807de32dba7b56d098cd88b658ed510c5f97e6920c0a3329c6c77003 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/send"
    iff1ecc8ba14343acb5c8a33ac4d9ba85826c36dacacf53c8b8b8b49f6b7d6cf7 "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/shipmentmethod"
    iff3796c5efae8359f95151c24f08603c4e641c88c6f20e5cee406c319745af8a "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/salesquotelines"
    id3852f54376c295ec4d02955c83253b69237fa6ba86eb2cdb39dbbf4b1afc4ce "github.com/microsoftgraph/msgraph-beta-sdk-go/financials/companies/item/salesquotes/item/salesquotelines/item"
)

type SalesQuoteRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type SalesQuoteRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func NewSalesQuoteRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SalesQuoteRequestBuilder) {
    m := &SalesQuoteRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/financials/companies/{company_id}/salesQuotes/{salesQuote_id}{?select,expand}";
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
func NewSalesQuoteRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SalesQuoteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSalesQuoteRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *SalesQuoteRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SalesQuoteRequestBuilder) CreateGetRequestInformation(q func (value *SalesQuoteRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(SalesQuoteRequestBuilderGetQueryParameters)
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
func (m *SalesQuoteRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesQuote, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *SalesQuoteRequestBuilder) Currency()(*i1a8296a2f13ef3c5ac1362a227b917a25604e9d63082e5b5ecddc5ba5a0cb2ca.CurrencyRequestBuilder) {
    return i1a8296a2f13ef3c5ac1362a227b917a25604e9d63082e5b5ecddc5ba5a0cb2ca.NewCurrencyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesQuoteRequestBuilder) Customer()(*i753033a2c9ce81017b74127e71bff7620415ec40c37877a3a9156a8a870a74e5.CustomerRequestBuilder) {
    return i753033a2c9ce81017b74127e71bff7620415ec40c37877a3a9156a8a870a74e5.NewCustomerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesQuoteRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *SalesQuoteRequestBuilder) Get(q func (value *SalesQuoteRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesQuote, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewSalesQuote() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesQuote), nil
}
func (m *SalesQuoteRequestBuilder) MakeInvoice()(*i80936828602bda4666c4eced48786bc2c1c175251ae1bdf1e0ef93378a7b543b.MakeInvoiceRequestBuilder) {
    return i80936828602bda4666c4eced48786bc2c1c175251ae1bdf1e0ef93378a7b543b.NewMakeInvoiceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesQuoteRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SalesQuote, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *SalesQuoteRequestBuilder) PaymentTerm()(*i72ac360010a42d020b08044ba3e52ca3d2232ffd23edf081ecd8320e34798c78.PaymentTermRequestBuilder) {
    return i72ac360010a42d020b08044ba3e52ca3d2232ffd23edf081ecd8320e34798c78.NewPaymentTermRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesQuoteRequestBuilder) SalesQuoteLines()(*iff3796c5efae8359f95151c24f08603c4e641c88c6f20e5cee406c319745af8a.SalesQuoteLinesRequestBuilder) {
    return iff3796c5efae8359f95151c24f08603c4e641c88c6f20e5cee406c319745af8a.NewSalesQuoteLinesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesQuoteRequestBuilder) SalesQuoteLinesById(id string)(*id3852f54376c295ec4d02955c83253b69237fa6ba86eb2cdb39dbbf4b1afc4ce.SalesQuoteLineRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["salesQuoteLine_id"] = id
    }
    return id3852f54376c295ec4d02955c83253b69237fa6ba86eb2cdb39dbbf4b1afc4ce.NewSalesQuoteLineRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SalesQuoteRequestBuilder) Send()(*if48d4a5b807de32dba7b56d098cd88b658ed510c5f97e6920c0a3329c6c77003.SendRequestBuilder) {
    return if48d4a5b807de32dba7b56d098cd88b658ed510c5f97e6920c0a3329c6c77003.NewSendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SalesQuoteRequestBuilder) ShipmentMethod()(*iff1ecc8ba14343acb5c8a33ac4d9ba85826c36dacacf53c8b8b8b49f6b7d6cf7.ShipmentMethodRequestBuilder) {
    return iff1ecc8ba14343acb5c8a33ac4d9ba85826c36dacacf53c8b8b8b49f6b7d6cf7.NewShipmentMethodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}