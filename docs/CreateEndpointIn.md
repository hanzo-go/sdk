# CreateEndpointIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description is a free-text label for the console. Optional, clipped to 1024 bytes. | [optional] 
**Events** | Pointer to **[]string** | Events are NATS subject patterns to subscribe to (e.g. \&quot;commerce.order.&gt;\&quot;). An empty or omitted list means EVERY event on the platform bus. Max 64 patterns, each max 256 bytes. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;active\&quot; or \&quot;disabled\&quot;. Empty defaults to active. A disabled endpoint receives no bus deliveries, but can still be exercised with POST /v1/webhooks/{id}/test. | [optional] 
**Url** | Pointer to **string** | URL is the https:// address each matching event is POSTed to. Required, max 2048 bytes; http:// and every other scheme is refused, because a webhook carries signed event data and must not travel in the clear. | [optional] 

## Methods

### NewCreateEndpointIn

`func NewCreateEndpointIn() *CreateEndpointIn`

NewCreateEndpointIn instantiates a new CreateEndpointIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEndpointInWithDefaults

`func NewCreateEndpointInWithDefaults() *CreateEndpointIn`

NewCreateEndpointInWithDefaults instantiates a new CreateEndpointIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateEndpointIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateEndpointIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateEndpointIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateEndpointIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEvents

`func (o *CreateEndpointIn) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CreateEndpointIn) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CreateEndpointIn) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CreateEndpointIn) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetStatus

`func (o *CreateEndpointIn) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateEndpointIn) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateEndpointIn) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateEndpointIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *CreateEndpointIn) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateEndpointIn) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateEndpointIn) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateEndpointIn) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


