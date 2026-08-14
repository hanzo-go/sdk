# TestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delivered** | Pointer to **bool** |  | [optional] 
**DurationMs** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**HttpStatus** | Pointer to **int32** |  | [optional] 

## Methods

### NewTestResult

`func NewTestResult() *TestResult`

NewTestResult instantiates a new TestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultWithDefaults

`func NewTestResultWithDefaults() *TestResult`

NewTestResultWithDefaults instantiates a new TestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelivered

`func (o *TestResult) GetDelivered() bool`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *TestResult) GetDeliveredOk() (*bool, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *TestResult) SetDelivered(v bool)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *TestResult) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetDurationMs

`func (o *TestResult) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *TestResult) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *TestResult) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *TestResult) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetError

`func (o *TestResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TestResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TestResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *TestResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetHttpStatus

`func (o *TestResult) GetHttpStatus() int32`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *TestResult) GetHttpStatusOk() (*int32, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *TestResult) SetHttpStatus(v int32)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *TestResult) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


