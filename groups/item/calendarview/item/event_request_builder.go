package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i1bcb02c01f74f11e3aa24ae414b4c02fbc829ed1cb109d46b654b3389426216e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/multivalueextendedproperties"
    i31a3b37c9dbd3939c543d5983439186da29622278bf71c6087acbd816f5c9239 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/dismissreminder"
    i3b113a779a2e1c26df62a5fc2779e15b76ed3e1090178152bfb73a05102fd7ea "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/extensions"
    i42ac3ded08c709492e9cc73bc81cf7f4f1449cd1dfae36ce915fb987d3483119 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/decline"
    i478bc546e67483a8a03a15808e207d1d211bbbfe1f2859bbad91fd68f0e9d11a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/forward"
    i4c47294298a9edb565053aaf021f3f8a03f8dd2db6441a9594f3c81e71a84ca2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/snoozereminder"
    i68776d8a0abd76d8055c38aaa88c6dfe00f5914bab2339676865bca7101b4c45 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances"
    i68a37fe1f53b4e518739b68cfbe9c468a06abeb302d9309775d0fb5864c0f17c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/calendar"
    i7c71d80390df615835b393e48dd37596b7a2b2840beddb58f10ed47068f5db1a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/tentativelyaccept"
    i7e83e7737d083b5a247d4b212e9838745f73ee0044cef7495db101a4d27d0473 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/cancel"
    i990077bc9239ae4abf654300201198e18c2649d863293f51f3f9660e7dd48615 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/singlevalueextendedproperties"
    ibeac8ac69135ca43836cbc7e86a3296df23382a77085df4916e1d8db90d1d1b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences"
    id7b295f54e3c066e9aee590c7e1cf7df8f692457387f19fc085cb1eaf48eebe5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/accept"
    ide5e10ee2a6673c0fe49d9fdac7898a497e239b5500fc38a3178a52ff622bf83 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/attachments"
    i0aff26eaed51b1fe4dcefaf7fdc9082ff77390914e54b94471bbabd9a4585874 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item"
    i0fdb512d350c6ee756976fb95fad12a002ad0d7b0d3fcf13590e89792367c733 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/multivalueextendedproperties/item"
    i53332243af1a23b575f7314ef56a041e8aad0a06d071d02238c1b61aeb8cefd9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item"
    iab96f1b678646231464bd3b7fb938638c604d6f548e5c414d63590125b8591b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/singlevalueextendedproperties/item"
    ic351771b8e3695ea47a713295cb3355c9bcc6d5ad38b45b39b527beb6475782b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/attachments/item"
    ie035c10affe323f0959d54cad2fac3bd37b1a2e6c88bcaece5284fe0542447db "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/extensions/item"
)

type EventRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type EventRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    EndDateTime *string;
    Expand []string;
    Select_escpaped []string;
    StartDateTime *string;
}
func (m *EventRequestBuilder) Accept()(*id7b295f54e3c066e9aee590c7e1cf7df8f692457387f19fc085cb1eaf48eebe5.AcceptRequestBuilder) {
    return id7b295f54e3c066e9aee590c7e1cf7df8f692457387f19fc085cb1eaf48eebe5.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*ide5e10ee2a6673c0fe49d9fdac7898a497e239b5500fc38a3178a52ff622bf83.AttachmentsRequestBuilder) {
    return ide5e10ee2a6673c0fe49d9fdac7898a497e239b5500fc38a3178a52ff622bf83.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) AttachmentsById(id string)(*ic351771b8e3695ea47a713295cb3355c9bcc6d5ad38b45b39b527beb6475782b.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ic351771b8e3695ea47a713295cb3355c9bcc6d5ad38b45b39b527beb6475782b.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i68a37fe1f53b4e518739b68cfbe9c468a06abeb302d9309775d0fb5864c0f17c.CalendarRequestBuilder) {
    return i68a37fe1f53b4e518739b68cfbe9c468a06abeb302d9309775d0fb5864c0f17c.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i7e83e7737d083b5a247d4b212e9838745f73ee0044cef7495db101a4d27d0473.CancelRequestBuilder) {
    return i7e83e7737d083b5a247d4b212e9838745f73ee0044cef7495db101a4d27d0473.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/groups/{group_id}/calendarView/{event_id}{?startDateTime,endDateTime,select,expand}";
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
func NewEventRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *EventRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EventRequestBuilder) CreateGetRequestInformation(q func (value *EventRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(EventRequestBuilderGetQueryParameters)
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
func (m *EventRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EventRequestBuilder) Decline()(*i42ac3ded08c709492e9cc73bc81cf7f4f1449cd1dfae36ce915fb987d3483119.DeclineRequestBuilder) {
    return i42ac3ded08c709492e9cc73bc81cf7f4f1449cd1dfae36ce915fb987d3483119.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EventRequestBuilder) DismissReminder()(*i31a3b37c9dbd3939c543d5983439186da29622278bf71c6087acbd816f5c9239.DismissReminderRequestBuilder) {
    return i31a3b37c9dbd3939c543d5983439186da29622278bf71c6087acbd816f5c9239.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrences()(*ibeac8ac69135ca43836cbc7e86a3296df23382a77085df4916e1d8db90d1d1b2.ExceptionOccurrencesRequestBuilder) {
    return ibeac8ac69135ca43836cbc7e86a3296df23382a77085df4916e1d8db90d1d1b2.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrencesById(id string)(*i53332243af1a23b575f7314ef56a041e8aad0a06d071d02238c1b61aeb8cefd9.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i53332243af1a23b575f7314ef56a041e8aad0a06d071d02238c1b61aeb8cefd9.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*i3b113a779a2e1c26df62a5fc2779e15b76ed3e1090178152bfb73a05102fd7ea.ExtensionsRequestBuilder) {
    return i3b113a779a2e1c26df62a5fc2779e15b76ed3e1090178152bfb73a05102fd7ea.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExtensionsById(id string)(*ie035c10affe323f0959d54cad2fac3bd37b1a2e6c88bcaece5284fe0542447db.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ie035c10affe323f0959d54cad2fac3bd37b1a2e6c88bcaece5284fe0542447db.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i478bc546e67483a8a03a15808e207d1d211bbbfe1f2859bbad91fd68f0e9d11a.ForwardRequestBuilder) {
    return i478bc546e67483a8a03a15808e207d1d211bbbfe1f2859bbad91fd68f0e9d11a.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Get(q func (value *EventRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEvent() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event), nil
}
func (m *EventRequestBuilder) Instances()(*i68776d8a0abd76d8055c38aaa88c6dfe00f5914bab2339676865bca7101b4c45.InstancesRequestBuilder) {
    return i68776d8a0abd76d8055c38aaa88c6dfe00f5914bab2339676865bca7101b4c45.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) InstancesById(id string)(*i0aff26eaed51b1fe4dcefaf7fdc9082ff77390914e54b94471bbabd9a4585874.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i0aff26eaed51b1fe4dcefaf7fdc9082ff77390914e54b94471bbabd9a4585874.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*i1bcb02c01f74f11e3aa24ae414b4c02fbc829ed1cb109d46b654b3389426216e.MultiValueExtendedPropertiesRequestBuilder) {
    return i1bcb02c01f74f11e3aa24ae414b4c02fbc829ed1cb109d46b654b3389426216e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i0fdb512d350c6ee756976fb95fad12a002ad0d7b0d3fcf13590e89792367c733.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i0fdb512d350c6ee756976fb95fad12a002ad0d7b0d3fcf13590e89792367c733.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*i990077bc9239ae4abf654300201198e18c2649d863293f51f3f9660e7dd48615.SingleValueExtendedPropertiesRequestBuilder) {
    return i990077bc9239ae4abf654300201198e18c2649d863293f51f3f9660e7dd48615.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*iab96f1b678646231464bd3b7fb938638c604d6f548e5c414d63590125b8591b1.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return iab96f1b678646231464bd3b7fb938638c604d6f548e5c414d63590125b8591b1.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i4c47294298a9edb565053aaf021f3f8a03f8dd2db6441a9594f3c81e71a84ca2.SnoozeReminderRequestBuilder) {
    return i4c47294298a9edb565053aaf021f3f8a03f8dd2db6441a9594f3c81e71a84ca2.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i7c71d80390df615835b393e48dd37596b7a2b2840beddb58f10ed47068f5db1a.TentativelyAcceptRequestBuilder) {
    return i7c71d80390df615835b393e48dd37596b7a2b2840beddb58f10ed47068f5db1a.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}