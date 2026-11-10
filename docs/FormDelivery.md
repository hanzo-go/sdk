# FormDelivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | Pointer to **int64** | Attempts is how many posts were made. | [optional] 
**Code** | Pointer to **int64** | Code is the last HTTP status the address answered; 0 when it did not answer. | [optional] 
**Created** | Pointer to **string** | Created is when the first attempt was made, RFC 3339. | [optional] 
**Error** | Pointer to **string** | Error is why the last attempt failed. | [optional] 
**Id** | Pointer to **string** | ID is the delivery&#39;s id; the body carries it as event_id. | [optional] 
**Next** | Pointer to **string** | Next is when the next attempt is due, RFC 3339; empty when none is. | [optional] 
**Response** | Pointer to **string** | Response is the response&#39;s id. | [optional] 
**Status** | Pointer to **string** | Status is delivered, retrying or failed. | [optional] 
**Updated** | Pointer to **string** | Updated is when the last was, RFC 3339. | [optional] 
**Webhook** | Pointer to **string** | Webhook is the webhook&#39;s id. | [optional] 

## Methods

### NewFormDelivery

`func NewFormDelivery() *FormDelivery`

NewFormDelivery instantiates a new FormDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormDeliveryWithDefaults

`func NewFormDeliveryWithDefaults() *FormDelivery`

NewFormDeliveryWithDefaults instantiates a new FormDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *FormDelivery) GetAttempts() int64`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *FormDelivery) GetAttemptsOk() (*int64, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *FormDelivery) SetAttempts(v int64)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *FormDelivery) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetCode

`func (o *FormDelivery) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FormDelivery) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FormDelivery) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *FormDelivery) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreated

`func (o *FormDelivery) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FormDelivery) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FormDelivery) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FormDelivery) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetError

`func (o *FormDelivery) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FormDelivery) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FormDelivery) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *FormDelivery) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *FormDelivery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormDelivery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormDelivery) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormDelivery) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNext

`func (o *FormDelivery) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *FormDelivery) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *FormDelivery) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *FormDelivery) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetResponse

`func (o *FormDelivery) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *FormDelivery) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *FormDelivery) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *FormDelivery) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetStatus

`func (o *FormDelivery) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FormDelivery) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FormDelivery) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FormDelivery) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdated

`func (o *FormDelivery) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FormDelivery) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FormDelivery) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FormDelivery) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetWebhook

`func (o *FormDelivery) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *FormDelivery) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *FormDelivery) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *FormDelivery) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


