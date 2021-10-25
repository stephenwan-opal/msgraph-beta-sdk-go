package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1a6fc94f2f407ef3ebbe51616e637c92682d9e0ba95d63e328274ae0ca5e5585 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/series"
    i5514cc60a132eba6d691ecfedecb029f9216a5dbeb19e42cdfa613c6966b7985 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/setposition"
    i6c767d0f0aeafa0422ac258cb03f2f748ab5ff0f24e33eb53db0105c62751f12 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/axes"
    i6f7c0fdf59d27b7f6e35df125c5e04485948ede897cc39b55525df5269deee99 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/imagewithwidth"
    i75490ea495a0803fdb95819346a6400eda5818426d9d2c360fd61560bbe3de50 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/imagewithwidthwithheightwithfittingmode"
    i8718ad311741532f01f0e102a7ec6ee973acf27cc6b7486c1057be248645b83d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/image"
    i946d011c71201d495be613c93f0a2cbc94ba16fda085a01bb70bf395f669c694 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/worksheet"
    ib426fdddb3005ff263ea40872b2ca2aa5bd42da8a15c273661d77a8623802ca2 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/setdata"
    ibd0c451689424d53c5dd6e330bc9ec2dc6c43431343e4fef77249490a2a0a8bf "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/format"
    ic7e9dcd5b5ec78229b5118ea614cfb86dee03f2154cd15b2b6b4c975d5b3c516 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/title"
    if9003a39f57a4011c21bf7517e3383b25de9c1c503b81f5f739649b5504c0037 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/legend"
    if9d41657bb26ba2fefbba18ac0d4a84726a97f2555201028777d697a01f1e493 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/imagewithwidthwithheight"
    ifd66e8b22d6f091a1e8f1971387da931b57b37d5067f454ebf033a330227893d "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/datalabels"
    i7d3d41b3c182984a4875637200cb5f4c401fec09228cccd72c9da2b28993aea9 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/workbook/tables/item/worksheet/charts/item/series/item"
)

type WorkbookChartRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type WorkbookChartRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *WorkbookChartRequestBuilder) Axes()(*i6c767d0f0aeafa0422ac258cb03f2f748ab5ff0f24e33eb53db0105c62751f12.AxesRequestBuilder) {
    return i6c767d0f0aeafa0422ac258cb03f2f748ab5ff0f24e33eb53db0105c62751f12.NewAxesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewWorkbookChartRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookChartRequestBuilder) {
    m := &WorkbookChartRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/workbooks/{driveItem_id}/workbook/tables/{workbookTable_id}/worksheet/charts/{workbookChart_id}{?select,expand}";
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
func NewWorkbookChartRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WorkbookChartRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkbookChartRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WorkbookChartRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WorkbookChartRequestBuilder) CreateGetRequestInformation(q func (value *WorkbookChartRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(WorkbookChartRequestBuilderGetQueryParameters)
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
func (m *WorkbookChartRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChart, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *WorkbookChartRequestBuilder) DataLabels()(*ifd66e8b22d6f091a1e8f1971387da931b57b37d5067f454ebf033a330227893d.DataLabelsRequestBuilder) {
    return ifd66e8b22d6f091a1e8f1971387da931b57b37d5067f454ebf033a330227893d.NewDataLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *WorkbookChartRequestBuilder) Format()(*ibd0c451689424d53c5dd6e330bc9ec2dc6c43431343e4fef77249490a2a0a8bf.FormatRequestBuilder) {
    return ibd0c451689424d53c5dd6e330bc9ec2dc6c43431343e4fef77249490a2a0a8bf.NewFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) Get(q func (value *WorkbookChartRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChart, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewWorkbookChart() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChart), nil
}
func (m *WorkbookChartRequestBuilder) Image()(*i8718ad311741532f01f0e102a7ec6ee973acf27cc6b7486c1057be248645b83d.ImageRequestBuilder) {
    return i8718ad311741532f01f0e102a7ec6ee973acf27cc6b7486c1057be248645b83d.NewImageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) ImageWithWidth(width *int32)(*i6f7c0fdf59d27b7f6e35df125c5e04485948ede897cc39b55525df5269deee99.ImageWithWidthRequestBuilder) {
    return i6f7c0fdf59d27b7f6e35df125c5e04485948ede897cc39b55525df5269deee99.NewImageWithWidthRequestBuilderInternal(m.pathParameters, m.requestAdapter, width);
}
func (m *WorkbookChartRequestBuilder) ImageWithWidthWithHeight(width *int32, height *int32)(*if9d41657bb26ba2fefbba18ac0d4a84726a97f2555201028777d697a01f1e493.ImageWithWidthWithHeightRequestBuilder) {
    return if9d41657bb26ba2fefbba18ac0d4a84726a97f2555201028777d697a01f1e493.NewImageWithWidthWithHeightRequestBuilderInternal(m.pathParameters, m.requestAdapter, width, height);
}
func (m *WorkbookChartRequestBuilder) ImageWithWidthWithHeightWithFittingMode(width *int32, height *int32, fittingMode *string)(*i75490ea495a0803fdb95819346a6400eda5818426d9d2c360fd61560bbe3de50.ImageWithWidthWithHeightWithFittingModeRequestBuilder) {
    return i75490ea495a0803fdb95819346a6400eda5818426d9d2c360fd61560bbe3de50.NewImageWithWidthWithHeightWithFittingModeRequestBuilderInternal(m.pathParameters, m.requestAdapter, width, height, fittingMode);
}
func (m *WorkbookChartRequestBuilder) Legend()(*if9003a39f57a4011c21bf7517e3383b25de9c1c503b81f5f739649b5504c0037.LegendRequestBuilder) {
    return if9003a39f57a4011c21bf7517e3383b25de9c1c503b81f5f739649b5504c0037.NewLegendRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.WorkbookChart, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *WorkbookChartRequestBuilder) Series()(*i1a6fc94f2f407ef3ebbe51616e637c92682d9e0ba95d63e328274ae0ca5e5585.SeriesRequestBuilder) {
    return i1a6fc94f2f407ef3ebbe51616e637c92682d9e0ba95d63e328274ae0ca5e5585.NewSeriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) SeriesById(id string)(*i7d3d41b3c182984a4875637200cb5f4c401fec09228cccd72c9da2b28993aea9.WorkbookChartSeriesRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["workbookChartSeries_id"] = id
    }
    return i7d3d41b3c182984a4875637200cb5f4c401fec09228cccd72c9da2b28993aea9.NewWorkbookChartSeriesRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) SetData()(*ib426fdddb3005ff263ea40872b2ca2aa5bd42da8a15c273661d77a8623802ca2.SetDataRequestBuilder) {
    return ib426fdddb3005ff263ea40872b2ca2aa5bd42da8a15c273661d77a8623802ca2.NewSetDataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) SetPosition()(*i5514cc60a132eba6d691ecfedecb029f9216a5dbeb19e42cdfa613c6966b7985.SetPositionRequestBuilder) {
    return i5514cc60a132eba6d691ecfedecb029f9216a5dbeb19e42cdfa613c6966b7985.NewSetPositionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) Title()(*ic7e9dcd5b5ec78229b5118ea614cfb86dee03f2154cd15b2b6b4c975d5b3c516.TitleRequestBuilder) {
    return ic7e9dcd5b5ec78229b5118ea614cfb86dee03f2154cd15b2b6b4c975d5b3c516.NewTitleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WorkbookChartRequestBuilder) Worksheet()(*i946d011c71201d495be613c93f0a2cbc94ba16fda085a01bb70bf395f669c694.WorksheetRequestBuilder) {
    return i946d011c71201d495be613c93f0a2cbc94ba16fda085a01bb70bf395f669c694.NewWorksheetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}