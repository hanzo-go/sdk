# WebhookDeliveryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | Pointer to **int64** | Attempt is which try this row is, starting at 1. The ladder waits 1s, then 5s, then 25s before the next one. | [optional] 
**Created** | Pointer to **string** | Created is when the attempt was made, RFC3339 in UTC. | [optional] 
**Delivery** | Pointer to **string** | DeliveryID groups the attempts for ONE event to ONE endpoint. Rows sharing it are the same delivery being retried, not separate events. | [optional] 
**DurationMs** | Pointer to **int64** | DurationMs is how long this attempt took end to end, in MILLISECONDS. | [optional] 
**Endpoint** | Pointer to **string** | EndpointID is which subscriber this attempt was for. | [optional] 
**Error** | Pointer to **string** | Error says what went wrong on a non-ok attempt. Empty on success. | [optional] 
**HttpStatus** | Pointer to **int64** | HTTPStatus is what the subscriber answered. ZERO means it never answered — a refused connection, a DNS failure or a timeout — which is why a zero here is not a 200. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot; when the subscriber accepted it, \&quot;retrying\&quot; while a further attempt will follow, and \&quot;failed\&quot; when none will. Exactly one row of a delivery is terminal. | [optional] 
**Subject** | Pointer to **string** | Subject is the event that was delivered (\&quot;commerce.order.created\&quot;). A manual test send carries \&quot;webhook.test\&quot;. | [optional] 

## Methods

### NewWebhookDeliveryRow

`func NewWebhookDeliveryRow() *WebhookDeliveryRow`

NewWebhookDeliveryRow instantiates a new WebhookDeliveryRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDeliveryRowWithDefaults

`func NewWebhookDeliveryRowWithDefaults() *WebhookDeliveryRow`

NewWebhookDeliveryRowWithDefaults instantiates a new WebhookDeliveryRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempt

`func (o *WebhookDeliveryRow) GetAttempt() int64`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *WebhookDeliveryRow) GetAttemptOk() (*int64, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *WebhookDeliveryRow) SetAttempt(v int64)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *WebhookDeliveryRow) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetCreated

`func (o *WebhookDeliveryRow) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WebhookDeliveryRow) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WebhookDeliveryRow) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *WebhookDeliveryRow) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDelivery

`func (o *WebhookDeliveryRow) GetDelivery() string`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *WebhookDeliveryRow) GetDeliveryOk() (*string, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *WebhookDeliveryRow) SetDelivery(v string)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *WebhookDeliveryRow) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.

### GetDurationMs

`func (o *WebhookDeliveryRow) GetDurationMs() int64`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *WebhookDeliveryRow) GetDurationMsOk() (*int64, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *WebhookDeliveryRow) SetDurationMs(v int64)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *WebhookDeliveryRow) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetEndpoint

`func (o *WebhookDeliveryRow) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *WebhookDeliveryRow) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *WebhookDeliveryRow) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *WebhookDeliveryRow) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetError

`func (o *WebhookDeliveryRow) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WebhookDeliveryRow) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WebhookDeliveryRow) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *WebhookDeliveryRow) HasError() bool`

HasError returns a boolean if a field has been set.

### GetHttpStatus

`func (o *WebhookDeliveryRow) GetHttpStatus() int64`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *WebhookDeliveryRow) GetHttpStatusOk() (*int64, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *WebhookDeliveryRow) SetHttpStatus(v int64)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *WebhookDeliveryRow) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.

### GetStatus

`func (o *WebhookDeliveryRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookDeliveryRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookDeliveryRow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebhookDeliveryRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubject

`func (o *WebhookDeliveryRow) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *WebhookDeliveryRow) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *WebhookDeliveryRow) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *WebhookDeliveryRow) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


