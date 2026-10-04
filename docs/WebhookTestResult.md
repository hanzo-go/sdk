# WebhookTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delivered** | Pointer to **bool** | Delivered is whether the subscriber accepted the test POST. It is the whole answer: the send is synchronous and is not retried. | [optional] 
**DurationMs** | Pointer to **int64** | DurationMs is how long the single attempt took, in MILLISECONDS. | [optional] 
**Error** | Pointer to **string** | Error says what stopped it. Empty when delivered. | [optional] 
**HttpStatus** | Pointer to **int64** | HTTPStatus is what the subscriber answered, or 0 if it never answered. | [optional] 

## Methods

### NewWebhookTestResult

`func NewWebhookTestResult() *WebhookTestResult`

NewWebhookTestResult instantiates a new WebhookTestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookTestResultWithDefaults

`func NewWebhookTestResultWithDefaults() *WebhookTestResult`

NewWebhookTestResultWithDefaults instantiates a new WebhookTestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelivered

`func (o *WebhookTestResult) GetDelivered() bool`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *WebhookTestResult) GetDeliveredOk() (*bool, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *WebhookTestResult) SetDelivered(v bool)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *WebhookTestResult) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetDurationMs

`func (o *WebhookTestResult) GetDurationMs() int64`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *WebhookTestResult) GetDurationMsOk() (*int64, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *WebhookTestResult) SetDurationMs(v int64)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *WebhookTestResult) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetError

`func (o *WebhookTestResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WebhookTestResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WebhookTestResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *WebhookTestResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetHttpStatus

`func (o *WebhookTestResult) GetHttpStatus() int64`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *WebhookTestResult) GetHttpStatusOk() (*int64, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *WebhookTestResult) SetHttpStatus(v int64)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *WebhookTestResult) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


