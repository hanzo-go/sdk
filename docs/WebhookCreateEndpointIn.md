# WebhookCreateEndpointIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description is a free-text label for the console. Optional, clipped to 1024 bytes. | [optional] 
**Events** | Pointer to **[]string** | Events are NATS subject patterns to subscribe to (e.g. \&quot;commerce.order.&gt;\&quot;). An empty or omitted list means EVERY event on the platform bus. Max 64 patterns, each max 256 bytes. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;active\&quot; or \&quot;disabled\&quot;. Empty defaults to active. A disabled endpoint receives no bus deliveries, but can still be exercised with POST /v1/webhook/{id}/test. | [optional] 
**Url** | Pointer to **string** | URL is the https:// address each matching event is POSTed to. Required, max 2048 bytes; http:// and every other scheme is refused, because a webhook carries signed event data and must not travel in the clear. | [optional] 

## Methods

### NewWebhookCreateEndpointIn

`func NewWebhookCreateEndpointIn() *WebhookCreateEndpointIn`

NewWebhookCreateEndpointIn instantiates a new WebhookCreateEndpointIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookCreateEndpointInWithDefaults

`func NewWebhookCreateEndpointInWithDefaults() *WebhookCreateEndpointIn`

NewWebhookCreateEndpointInWithDefaults instantiates a new WebhookCreateEndpointIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WebhookCreateEndpointIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookCreateEndpointIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookCreateEndpointIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookCreateEndpointIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEvents

`func (o *WebhookCreateEndpointIn) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookCreateEndpointIn) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookCreateEndpointIn) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *WebhookCreateEndpointIn) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetStatus

`func (o *WebhookCreateEndpointIn) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookCreateEndpointIn) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookCreateEndpointIn) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebhookCreateEndpointIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookCreateEndpointIn) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookCreateEndpointIn) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookCreateEndpointIn) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WebhookCreateEndpointIn) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


