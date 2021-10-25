package iosvppapp

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i01615e369fddd26fbf9f2a9762b1d4ac304ddda1d4948d54db1d56bf795e2221 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/devicestatuses/item/app/iosvppapp/revokedevicelicense"
    i246c46f95b555db6d8964aa1e90be55f7be5eac10e1764d3c3f5441f722d171f "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/devicestatuses/item/app/iosvppapp/revokealllicenses"
    iea6b669b897a4c035a59b5f017a44d3620d7cabef826776dd92c368838079fae "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/devicestatuses/item/app/iosvppapp/revokeuserlicense"
)

type IosVppAppRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func NewIosVppAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IosVppAppRequestBuilder) {
    m := &IosVppAppRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceAppManagement/mobileApps/{mobileApp_id}/deviceStatuses/{mobileAppInstallStatus_id}/app/microsoft.graph.iosVppApp";
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
func NewIosVppAppRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IosVppAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIosVppAppRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *IosVppAppRequestBuilder) RevokeAllLicenses()(*i246c46f95b555db6d8964aa1e90be55f7be5eac10e1764d3c3f5441f722d171f.RevokeAllLicensesRequestBuilder) {
    return i246c46f95b555db6d8964aa1e90be55f7be5eac10e1764d3c3f5441f722d171f.NewRevokeAllLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IosVppAppRequestBuilder) RevokeDeviceLicense()(*i01615e369fddd26fbf9f2a9762b1d4ac304ddda1d4948d54db1d56bf795e2221.RevokeDeviceLicenseRequestBuilder) {
    return i01615e369fddd26fbf9f2a9762b1d4ac304ddda1d4948d54db1d56bf795e2221.NewRevokeDeviceLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IosVppAppRequestBuilder) RevokeUserLicense()(*iea6b669b897a4c035a59b5f017a44d3620d7cabef826776dd92c368838079fae.RevokeUserLicenseRequestBuilder) {
    return iea6b669b897a4c035a59b5f017a44d3620d7cabef826776dd92c368838079fae.NewRevokeUserLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}