package print

import (
    i0d270c48a37eef745b51dd39fcd6e6c396a49d6e354a0a22c4b38a373b1f8649 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/shares"
    i2a6efa35f1da64a13760e57fe6078241cde1fc0716fe6cf741b1913916af2a06 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printershares"
    i32ff24c6864832ad0f41a0c6482a3f153c54ab6eec76bf30bcdc4a10beb817fc "github.com/microsoftgraph/msgraph-beta-sdk-go/print/services"
    i4c4b7dddb590dec724ed827d1d5c27e935dc2a75d5fc9b1ddb47f1cc82d2f70a "github.com/microsoftgraph/msgraph-beta-sdk-go/print/taskdefinitions"
    i70c2e2fc495120f72326108a5af86cde4684b679d130a86acdd4046390fee262 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports"
    i844ad4158981a129d54aa4a317c17dc8b56ba989d7b7cfa3ffe6aebbb129c245 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/operations"
    i92d1f50b33fa740ede98125a2225a30e704277949de6d5ba06e1e4f209889a99 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ie0ee21308c5a5876efdcb565ab54dddf497480052514698aa5df8076282afb54 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/connectors"
    i041fa93b1b8c4c6acaa44fa25a4bab5b60619191120dc3943102fdec799fca39 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/shares/item"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i20c82299a583bf997be8654a60766afa3b225e0e8b54ab661cb734c46c2add9b "github.com/microsoftgraph/msgraph-beta-sdk-go/print/operations/item"
    i21ad7a37923a403c8172690e362e15a8a0298e4a82eecbedee9e691ff99bdde3 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printers/item"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5ec86a089dc848e3bf34cce2d116379624ef91e1c58d8a32e664e6fb5397c958 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/taskdefinitions/item"
    i6ac2f8e4588dd2304d6a10565c050a97a9cab528db29bf610191796b5477b498 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/printershares/item"
    i9508344ddcc1d2a6a0ed36e7aeb07d60730f841e1e8d4674d8ae4d356d4932a9 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/services/item"
    id93d2fbcff2a13af1a3388fed5faf01351e12efec3ce632d2152e37f5ad29d99 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/connectors/item"
)

type PrintRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type PrintRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *PrintRequestBuilder) Connectors()(*ie0ee21308c5a5876efdcb565ab54dddf497480052514698aa5df8076282afb54.ConnectorsRequestBuilder) {
    return ie0ee21308c5a5876efdcb565ab54dddf497480052514698aa5df8076282afb54.NewConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintRequestBuilder) ConnectorsById(id string)(*id93d2fbcff2a13af1a3388fed5faf01351e12efec3ce632d2152e37f5ad29d99.PrintConnectorRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["printConnector_id"] = id
    }
    return id93d2fbcff2a13af1a3388fed5faf01351e12efec3ce632d2152e37f5ad29d99.NewPrintConnectorRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewPrintRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintRequestBuilder) {
    m := &PrintRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/print{?select,expand}";
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
func NewPrintRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrintRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrintRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *PrintRequestBuilder) CreateGetRequestInformation(q func (value *PrintRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(PrintRequestBuilderGetQueryParameters)
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
func (m *PrintRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Print, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *PrintRequestBuilder) Get(q func (value *PrintRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Print, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPrint() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Print), nil
}
func (m *PrintRequestBuilder) Operations()(*i844ad4158981a129d54aa4a317c17dc8b56ba989d7b7cfa3ffe6aebbb129c245.OperationsRequestBuilder) {
    return i844ad4158981a129d54aa4a317c17dc8b56ba989d7b7cfa3ffe6aebbb129c245.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintRequestBuilder) OperationsById(id string)(*i20c82299a583bf997be8654a60766afa3b225e0e8b54ab661cb734c46c2add9b.PrintOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["printOperation_id"] = id
    }
    return i20c82299a583bf997be8654a60766afa3b225e0e8b54ab661cb734c46c2add9b.NewPrintOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrintRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Print, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *PrintRequestBuilder) Printers()(*i92d1f50b33fa740ede98125a2225a30e704277949de6d5ba06e1e4f209889a99.PrintersRequestBuilder) {
    return i92d1f50b33fa740ede98125a2225a30e704277949de6d5ba06e1e4f209889a99.NewPrintersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintRequestBuilder) PrintersById(id string)(*i21ad7a37923a403c8172690e362e15a8a0298e4a82eecbedee9e691ff99bdde3.PrinterRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["printer_id"] = id
    }
    return i21ad7a37923a403c8172690e362e15a8a0298e4a82eecbedee9e691ff99bdde3.NewPrinterRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrintRequestBuilder) PrinterShares()(*i2a6efa35f1da64a13760e57fe6078241cde1fc0716fe6cf741b1913916af2a06.PrinterSharesRequestBuilder) {
    return i2a6efa35f1da64a13760e57fe6078241cde1fc0716fe6cf741b1913916af2a06.NewPrinterSharesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintRequestBuilder) PrinterSharesById(id string)(*i6ac2f8e4588dd2304d6a10565c050a97a9cab528db29bf610191796b5477b498.PrinterShareRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["printerShare_id"] = id
    }
    return i6ac2f8e4588dd2304d6a10565c050a97a9cab528db29bf610191796b5477b498.NewPrinterShareRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrintRequestBuilder) Reports()(*i70c2e2fc495120f72326108a5af86cde4684b679d130a86acdd4046390fee262.ReportsRequestBuilder) {
    return i70c2e2fc495120f72326108a5af86cde4684b679d130a86acdd4046390fee262.NewReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintRequestBuilder) Services()(*i32ff24c6864832ad0f41a0c6482a3f153c54ab6eec76bf30bcdc4a10beb817fc.ServicesRequestBuilder) {
    return i32ff24c6864832ad0f41a0c6482a3f153c54ab6eec76bf30bcdc4a10beb817fc.NewServicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintRequestBuilder) ServicesById(id string)(*i9508344ddcc1d2a6a0ed36e7aeb07d60730f841e1e8d4674d8ae4d356d4932a9.PrintServiceRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["printService_id"] = id
    }
    return i9508344ddcc1d2a6a0ed36e7aeb07d60730f841e1e8d4674d8ae4d356d4932a9.NewPrintServiceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrintRequestBuilder) Shares()(*i0d270c48a37eef745b51dd39fcd6e6c396a49d6e354a0a22c4b38a373b1f8649.SharesRequestBuilder) {
    return i0d270c48a37eef745b51dd39fcd6e6c396a49d6e354a0a22c4b38a373b1f8649.NewSharesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintRequestBuilder) SharesById(id string)(*i041fa93b1b8c4c6acaa44fa25a4bab5b60619191120dc3943102fdec799fca39.PrinterShareRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["printerShare_id"] = id
    }
    return i041fa93b1b8c4c6acaa44fa25a4bab5b60619191120dc3943102fdec799fca39.NewPrinterShareRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrintRequestBuilder) TaskDefinitions()(*i4c4b7dddb590dec724ed827d1d5c27e935dc2a75d5fc9b1ddb47f1cc82d2f70a.TaskDefinitionsRequestBuilder) {
    return i4c4b7dddb590dec724ed827d1d5c27e935dc2a75d5fc9b1ddb47f1cc82d2f70a.NewTaskDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PrintRequestBuilder) TaskDefinitionsById(id string)(*i5ec86a089dc848e3bf34cce2d116379624ef91e1c58d8a32e664e6fb5397c958.PrintTaskDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["printTaskDefinition_id"] = id
    }
    return i5ec86a089dc848e3bf34cce2d116379624ef91e1c58d8a32e664e6fb5397c958.NewPrintTaskDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}