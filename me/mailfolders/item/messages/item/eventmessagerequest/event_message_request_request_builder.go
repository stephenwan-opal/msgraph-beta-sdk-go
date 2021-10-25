package eventmessagerequest

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i0899d26b0b0522b710df806f0454178eaf5c74fca4ec1761f9293365fb10a400 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/eventmessagerequest/accept"
    i1b1fb416c2a097fa7fb6d7f58517668193a1f68c1c86940328d84a96e121c128 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/eventmessagerequest/decline"
    i97f7c5bb3fb3d0e102fdbcc51d3c0673e7144cf63586517eaf7160cf5c7452e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/mailfolders/item/messages/item/eventmessagerequest/tentativelyaccept"
)

type EventMessageRequestRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func (m *EventMessageRequestRequestBuilder) Accept()(*i0899d26b0b0522b710df806f0454178eaf5c74fca4ec1761f9293365fb10a400.AcceptRequestBuilder) {
    return i0899d26b0b0522b710df806f0454178eaf5c74fca4ec1761f9293365fb10a400.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewEventMessageRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventMessageRequestRequestBuilder) {
    m := &EventMessageRequestRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/me/mailFolders/{mailFolder_id}/messages/{message_id}/microsoft.graph.eventMessageRequest";
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
func NewEventMessageRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventMessageRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventMessageRequestRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *EventMessageRequestRequestBuilder) Decline()(*i1b1fb416c2a097fa7fb6d7f58517668193a1f68c1c86940328d84a96e121c128.DeclineRequestBuilder) {
    return i1b1fb416c2a097fa7fb6d7f58517668193a1f68c1c86940328d84a96e121c128.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventMessageRequestRequestBuilder) TentativelyAccept()(*i97f7c5bb3fb3d0e102fdbcc51d3c0673e7144cf63586517eaf7160cf5c7452e8.TentativelyAcceptRequestBuilder) {
    return i97f7c5bb3fb3d0e102fdbcc51d3c0673e7144cf63586517eaf7160cf5c7452e8.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}