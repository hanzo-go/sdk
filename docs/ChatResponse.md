# ChatResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Funnel** | Pointer to [**Funnel**](Funnel.md) | Funnel is the org&#39;s trailing-window traffic → signups → orders. | [optional] 
**Reply** | Pointer to **string** | Reply is the coach&#39;s answer, grounded only in the quests and funnel below. When no AI plane is reachable it is the deterministic reply naming the top real quest — never silence, never invention. | [optional] 
**Suggestions** | Pointer to [**[]Suggestion**](Suggestion.md) | Suggestions are the current candidate quests, ranked best-first. | [optional] 

## Methods

### NewChatResponse

`func NewChatResponse() *ChatResponse`

NewChatResponse instantiates a new ChatResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatResponseWithDefaults

`func NewChatResponseWithDefaults() *ChatResponse`

NewChatResponseWithDefaults instantiates a new ChatResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunnel

`func (o *ChatResponse) GetFunnel() Funnel`

GetFunnel returns the Funnel field if non-nil, zero value otherwise.

### GetFunnelOk

`func (o *ChatResponse) GetFunnelOk() (*Funnel, bool)`

GetFunnelOk returns a tuple with the Funnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnel

`func (o *ChatResponse) SetFunnel(v Funnel)`

SetFunnel sets Funnel field to given value.

### HasFunnel

`func (o *ChatResponse) HasFunnel() bool`

HasFunnel returns a boolean if a field has been set.

### GetReply

`func (o *ChatResponse) GetReply() string`

GetReply returns the Reply field if non-nil, zero value otherwise.

### GetReplyOk

`func (o *ChatResponse) GetReplyOk() (*string, bool)`

GetReplyOk returns a tuple with the Reply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReply

`func (o *ChatResponse) SetReply(v string)`

SetReply sets Reply field to given value.

### HasReply

`func (o *ChatResponse) HasReply() bool`

HasReply returns a boolean if a field has been set.

### GetSuggestions

`func (o *ChatResponse) GetSuggestions() []Suggestion`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *ChatResponse) GetSuggestionsOk() (*[]Suggestion, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *ChatResponse) SetSuggestions(v []Suggestion)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *ChatResponse) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


