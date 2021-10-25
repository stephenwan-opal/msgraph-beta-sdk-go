package policies

import (
    i089acabe34d65ca8bb4a87f0ec051d51355817922b4d3ae2c686f574d377a384 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/featurerolloutpolicies"
    i0ea4dbcf49512a3d14a716de641b83949fd0f9b73ab17d8b05e57362070a224f "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/b2cauthenticationmethodspolicy"
    i154ced0e8b3656059650a4bd061d0aa6f15bc39d29746142cde157c94c8ce663 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/identitysecuritydefaultsenforcementpolicy"
    i23c105ad486b858fd7cc919b7625b7a6fc44a23cada9df239ade7ecbded23663 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/rolemanagementpolicyassignments"
    i2c474f2aded367ce2b8ec8be9a4c3fb07272675be1e73951735d60f2e6124711 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/authenticationflowspolicy"
    i37b836bda4608f9bed0d5aab3b4328d55f977ad2084b34eeeffb42e49f25237b "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/authenticationmethodspolicy"
    i3baef4935efa79b1ea53b0f8a4d12e8e33271d34f1b7fe4e5e3deefd698fb11d "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/adminconsentrequestpolicy"
    i4235b128855e4ef4e736e93e75306be2329042607217b14308bae073614f6b91 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/permissiongrantpolicies"
    i50186836a679beade2d866abd3fcf49adbbaf1d6161193ee717820c04941b95f "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/rolemanagementpolicies"
    i82a0672e3a16707a14cf887f3b2b473d9e08df769ebfa56b019ad2205ea8732d "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/claimsmappingpolicies"
    i8ab23dd53a7ba192025fbd1c53f22ecde4f5853cf55f14dadae7c3d54bfe6ecc "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/accessreviewpolicy"
    i8da786288f6b1e12f48f41a28df1f6b81f287a5228f3978a4411dc247e79361d "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/tokenissuancepolicies"
    i9b0ad1e2bc66101ab80e51c56c0530b8d739993816ab39caf1d4ce646d70b5f6 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/directoryroleaccessreviewpolicy"
    ia589aeb346820886345aeb5931d4c55613a827996ee7bc74d7b23ee8114c4fb2 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/conditionalaccesspolicies"
    ia91d9fa5e7c39ffd42d68d1e1b891366a26470bb743dc20e271a38370621a692 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/mobiledevicemanagementpolicies"
    ib06edf2243c232ea0ac4674c243fe1a2abb0ee0eed56c1a59ce081927673770e "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/activitybasedtimeoutpolicies"
    ib78e441cfc724e7ca2b7aa878fef387840fdd51d15409c2c4abf683297e6d1a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/tokenlifetimepolicies"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    idf9e8b9f9da041893e6b6cb24dd45425a8696e3da093aef925b748ba4b81e936 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/authorizationpolicy"
    iea55f9833b3f70a27d04fdb2fbb1fdc5b35798960f8ea227ae35055ce1b36486 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/homerealmdiscoverypolicies"
    ife15f5c3e6b78fb8bf35531d4fa132e7caa227e970656877474d6ac270ea6da2 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/mobileappmanagementpolicies"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i1e97552013b3bffa16b2d23c422d3b88e120ef118f30ef817760eb98fdf1c32e "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/tokenlifetimepolicies/item"
    i2e9acaa79a19f61068a524155872974e6a98a1692e31a3056613d7fadf7058aa "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/permissiongrantpolicies/item"
    i3718ba1ac7c6b6b07b29bbaa42dc50cc122da4f510e54af905f2c63f5fba5585 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/rolemanagementpolicyassignments/item"
    i3d72a8d05a1bc29e49b08e440b0f124e60a0a17cca39da307f0fc359f9ba8100 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/claimsmappingpolicies/item"
    i45b1afd0d7bc9158e4c62ccf7aeb399e7aef62431980207f6b6df372b62073b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/tokenissuancepolicies/item"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5632bf91798252c9bca8fa6eadd95db82f90a142d927c3c26e684ab6f2d90a28 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/featurerolloutpolicies/item"
    i987e92fec12c74acd600441b6cc0e2c18b324562d146333edb4272ddc51cc302 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/mobiledevicemanagementpolicies/item"
    ic2c067f4697fa94bc77e049a3cf157ed99ce46ed86b0a4e4b688d09062795945 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/homerealmdiscoverypolicies/item"
    ic80fa25ff11a44e63435e8a7ab079ccac6e26dd9dc32f6b7a64be90be2f58011 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/mobileappmanagementpolicies/item"
    ie2aa8a198ac4cb5a2a1216675a0ebc84eb1ccb50965ba209e1e87acfa29f6a87 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/activitybasedtimeoutpolicies/item"
    if5c99775547db21eee7a16da27e1a94c2a2912f7796bc11e4a67761a02798bdb "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/conditionalaccesspolicies/item"
    if796ee537abe837147e4bdf42134467182183733786146dccb3e3c9dd5e93b54 "github.com/microsoftgraph/msgraph-beta-sdk-go/policies/rolemanagementpolicies/item"
)

type PoliciesRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type PoliciesRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *PoliciesRequestBuilder) AccessReviewPolicy()(*i8ab23dd53a7ba192025fbd1c53f22ecde4f5853cf55f14dadae7c3d54bfe6ecc.AccessReviewPolicyRequestBuilder) {
    return i8ab23dd53a7ba192025fbd1c53f22ecde4f5853cf55f14dadae7c3d54bfe6ecc.NewAccessReviewPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) ActivityBasedTimeoutPolicies()(*ib06edf2243c232ea0ac4674c243fe1a2abb0ee0eed56c1a59ce081927673770e.ActivityBasedTimeoutPoliciesRequestBuilder) {
    return ib06edf2243c232ea0ac4674c243fe1a2abb0ee0eed56c1a59ce081927673770e.NewActivityBasedTimeoutPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) ActivityBasedTimeoutPoliciesById(id string)(*ie2aa8a198ac4cb5a2a1216675a0ebc84eb1ccb50965ba209e1e87acfa29f6a87.ActivityBasedTimeoutPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["activityBasedTimeoutPolicy_id"] = id
    }
    return ie2aa8a198ac4cb5a2a1216675a0ebc84eb1ccb50965ba209e1e87acfa29f6a87.NewActivityBasedTimeoutPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) AdminConsentRequestPolicy()(*i3baef4935efa79b1ea53b0f8a4d12e8e33271d34f1b7fe4e5e3deefd698fb11d.AdminConsentRequestPolicyRequestBuilder) {
    return i3baef4935efa79b1ea53b0f8a4d12e8e33271d34f1b7fe4e5e3deefd698fb11d.NewAdminConsentRequestPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) AuthenticationFlowsPolicy()(*i2c474f2aded367ce2b8ec8be9a4c3fb07272675be1e73951735d60f2e6124711.AuthenticationFlowsPolicyRequestBuilder) {
    return i2c474f2aded367ce2b8ec8be9a4c3fb07272675be1e73951735d60f2e6124711.NewAuthenticationFlowsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) AuthenticationMethodsPolicy()(*i37b836bda4608f9bed0d5aab3b4328d55f977ad2084b34eeeffb42e49f25237b.AuthenticationMethodsPolicyRequestBuilder) {
    return i37b836bda4608f9bed0d5aab3b4328d55f977ad2084b34eeeffb42e49f25237b.NewAuthenticationMethodsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) AuthorizationPolicy()(*idf9e8b9f9da041893e6b6cb24dd45425a8696e3da093aef925b748ba4b81e936.AuthorizationPolicyRequestBuilder) {
    return idf9e8b9f9da041893e6b6cb24dd45425a8696e3da093aef925b748ba4b81e936.NewAuthorizationPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) AuthorizationPolicyById(id string)(*idf9e8b9f9da041893e6b6cb24dd45425a8696e3da093aef925b748ba4b81e936.AuthorizationPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["authorizationPolicy_id"] = id
    }
    return idf9e8b9f9da041893e6b6cb24dd45425a8696e3da093aef925b748ba4b81e936.NewAuthorizationPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) B2cAuthenticationMethodsPolicy()(*i0ea4dbcf49512a3d14a716de641b83949fd0f9b73ab17d8b05e57362070a224f.B2cAuthenticationMethodsPolicyRequestBuilder) {
    return i0ea4dbcf49512a3d14a716de641b83949fd0f9b73ab17d8b05e57362070a224f.NewB2cAuthenticationMethodsPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) ClaimsMappingPolicies()(*i82a0672e3a16707a14cf887f3b2b473d9e08df769ebfa56b019ad2205ea8732d.ClaimsMappingPoliciesRequestBuilder) {
    return i82a0672e3a16707a14cf887f3b2b473d9e08df769ebfa56b019ad2205ea8732d.NewClaimsMappingPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) ClaimsMappingPoliciesById(id string)(*i3d72a8d05a1bc29e49b08e440b0f124e60a0a17cca39da307f0fc359f9ba8100.ClaimsMappingPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["claimsMappingPolicy_id"] = id
    }
    return i3d72a8d05a1bc29e49b08e440b0f124e60a0a17cca39da307f0fc359f9ba8100.NewClaimsMappingPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) ConditionalAccessPolicies()(*ia589aeb346820886345aeb5931d4c55613a827996ee7bc74d7b23ee8114c4fb2.ConditionalAccessPoliciesRequestBuilder) {
    return ia589aeb346820886345aeb5931d4c55613a827996ee7bc74d7b23ee8114c4fb2.NewConditionalAccessPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) ConditionalAccessPoliciesById(id string)(*if5c99775547db21eee7a16da27e1a94c2a2912f7796bc11e4a67761a02798bdb.ConditionalAccessPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["conditionalAccessPolicy_id"] = id
    }
    return if5c99775547db21eee7a16da27e1a94c2a2912f7796bc11e4a67761a02798bdb.NewConditionalAccessPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PoliciesRequestBuilder) {
    m := &PoliciesRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/policies{?select,expand}";
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
func NewPoliciesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *PoliciesRequestBuilder) CreateGetRequestInformation(q func (value *PoliciesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(PoliciesRequestBuilderGetQueryParameters)
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
func (m *PoliciesRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Policies, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *PoliciesRequestBuilder) DirectoryRoleAccessReviewPolicy()(*i9b0ad1e2bc66101ab80e51c56c0530b8d739993816ab39caf1d4ce646d70b5f6.DirectoryRoleAccessReviewPolicyRequestBuilder) {
    return i9b0ad1e2bc66101ab80e51c56c0530b8d739993816ab39caf1d4ce646d70b5f6.NewDirectoryRoleAccessReviewPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) FeatureRolloutPolicies()(*i089acabe34d65ca8bb4a87f0ec051d51355817922b4d3ae2c686f574d377a384.FeatureRolloutPoliciesRequestBuilder) {
    return i089acabe34d65ca8bb4a87f0ec051d51355817922b4d3ae2c686f574d377a384.NewFeatureRolloutPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) FeatureRolloutPoliciesById(id string)(*i5632bf91798252c9bca8fa6eadd95db82f90a142d927c3c26e684ab6f2d90a28.FeatureRolloutPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["featureRolloutPolicy_id"] = id
    }
    return i5632bf91798252c9bca8fa6eadd95db82f90a142d927c3c26e684ab6f2d90a28.NewFeatureRolloutPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) Get(q func (value *PoliciesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PolicyRoot, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPolicyRoot() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PolicyRoot), nil
}
func (m *PoliciesRequestBuilder) HomeRealmDiscoveryPolicies()(*iea55f9833b3f70a27d04fdb2fbb1fdc5b35798960f8ea227ae35055ce1b36486.HomeRealmDiscoveryPoliciesRequestBuilder) {
    return iea55f9833b3f70a27d04fdb2fbb1fdc5b35798960f8ea227ae35055ce1b36486.NewHomeRealmDiscoveryPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) HomeRealmDiscoveryPoliciesById(id string)(*ic2c067f4697fa94bc77e049a3cf157ed99ce46ed86b0a4e4b688d09062795945.HomeRealmDiscoveryPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["homeRealmDiscoveryPolicy_id"] = id
    }
    return ic2c067f4697fa94bc77e049a3cf157ed99ce46ed86b0a4e4b688d09062795945.NewHomeRealmDiscoveryPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) IdentitySecurityDefaultsEnforcementPolicy()(*i154ced0e8b3656059650a4bd061d0aa6f15bc39d29746142cde157c94c8ce663.IdentitySecurityDefaultsEnforcementPolicyRequestBuilder) {
    return i154ced0e8b3656059650a4bd061d0aa6f15bc39d29746142cde157c94c8ce663.NewIdentitySecurityDefaultsEnforcementPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) MobileAppManagementPolicies()(*ife15f5c3e6b78fb8bf35531d4fa132e7caa227e970656877474d6ac270ea6da2.MobileAppManagementPoliciesRequestBuilder) {
    return ife15f5c3e6b78fb8bf35531d4fa132e7caa227e970656877474d6ac270ea6da2.NewMobileAppManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) MobileAppManagementPoliciesById(id string)(*ic80fa25ff11a44e63435e8a7ab079ccac6e26dd9dc32f6b7a64be90be2f58011.MobilityManagementPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["mobilityManagementPolicy_id"] = id
    }
    return ic80fa25ff11a44e63435e8a7ab079ccac6e26dd9dc32f6b7a64be90be2f58011.NewMobilityManagementPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) MobileDeviceManagementPolicies()(*ia91d9fa5e7c39ffd42d68d1e1b891366a26470bb743dc20e271a38370621a692.MobileDeviceManagementPoliciesRequestBuilder) {
    return ia91d9fa5e7c39ffd42d68d1e1b891366a26470bb743dc20e271a38370621a692.NewMobileDeviceManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) MobileDeviceManagementPoliciesById(id string)(*i987e92fec12c74acd600441b6cc0e2c18b324562d146333edb4272ddc51cc302.MobilityManagementPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["mobilityManagementPolicy_id"] = id
    }
    return i987e92fec12c74acd600441b6cc0e2c18b324562d146333edb4272ddc51cc302.NewMobilityManagementPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Policies, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *PoliciesRequestBuilder) PermissionGrantPolicies()(*i4235b128855e4ef4e736e93e75306be2329042607217b14308bae073614f6b91.PermissionGrantPoliciesRequestBuilder) {
    return i4235b128855e4ef4e736e93e75306be2329042607217b14308bae073614f6b91.NewPermissionGrantPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) PermissionGrantPoliciesById(id string)(*i2e9acaa79a19f61068a524155872974e6a98a1692e31a3056613d7fadf7058aa.PermissionGrantPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["permissionGrantPolicy_id"] = id
    }
    return i2e9acaa79a19f61068a524155872974e6a98a1692e31a3056613d7fadf7058aa.NewPermissionGrantPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) RoleManagementPolicies()(*i50186836a679beade2d866abd3fcf49adbbaf1d6161193ee717820c04941b95f.RoleManagementPoliciesRequestBuilder) {
    return i50186836a679beade2d866abd3fcf49adbbaf1d6161193ee717820c04941b95f.NewRoleManagementPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) RoleManagementPoliciesById(id string)(*if796ee537abe837147e4bdf42134467182183733786146dccb3e3c9dd5e93b54.UnifiedRoleManagementPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["unifiedRoleManagementPolicy_id"] = id
    }
    return if796ee537abe837147e4bdf42134467182183733786146dccb3e3c9dd5e93b54.NewUnifiedRoleManagementPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) RoleManagementPolicyAssignments()(*i23c105ad486b858fd7cc919b7625b7a6fc44a23cada9df239ade7ecbded23663.RoleManagementPolicyAssignmentsRequestBuilder) {
    return i23c105ad486b858fd7cc919b7625b7a6fc44a23cada9df239ade7ecbded23663.NewRoleManagementPolicyAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) RoleManagementPolicyAssignmentsById(id string)(*i3718ba1ac7c6b6b07b29bbaa42dc50cc122da4f510e54af905f2c63f5fba5585.UnifiedRoleManagementPolicyAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["unifiedRoleManagementPolicyAssignment_id"] = id
    }
    return i3718ba1ac7c6b6b07b29bbaa42dc50cc122da4f510e54af905f2c63f5fba5585.NewUnifiedRoleManagementPolicyAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) TokenIssuancePolicies()(*i8da786288f6b1e12f48f41a28df1f6b81f287a5228f3978a4411dc247e79361d.TokenIssuancePoliciesRequestBuilder) {
    return i8da786288f6b1e12f48f41a28df1f6b81f287a5228f3978a4411dc247e79361d.NewTokenIssuancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) TokenIssuancePoliciesById(id string)(*i45b1afd0d7bc9158e4c62ccf7aeb399e7aef62431980207f6b6df372b62073b9.TokenIssuancePolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["tokenIssuancePolicy_id"] = id
    }
    return i45b1afd0d7bc9158e4c62ccf7aeb399e7aef62431980207f6b6df372b62073b9.NewTokenIssuancePolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) TokenLifetimePolicies()(*ib78e441cfc724e7ca2b7aa878fef387840fdd51d15409c2c4abf683297e6d1a0.TokenLifetimePoliciesRequestBuilder) {
    return ib78e441cfc724e7ca2b7aa878fef387840fdd51d15409c2c4abf683297e6d1a0.NewTokenLifetimePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PoliciesRequestBuilder) TokenLifetimePoliciesById(id string)(*i1e97552013b3bffa16b2d23c422d3b88e120ef118f30ef817760eb98fdf1c32e.TokenLifetimePolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["tokenLifetimePolicy_id"] = id
    }
    return i1e97552013b3bffa16b2d23c422d3b88e120ef118f30ef817760eb98fdf1c32e.NewTokenLifetimePolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}