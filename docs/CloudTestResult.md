# CloudTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delivered** | Pointer to **bool** |  | [optional] 
**DurationMs** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**HttpStatus** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudTestResult

`func NewCloudTestResult() *CloudTestResult`

NewCloudTestResult instantiates a new CloudTestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTestResultWithDefaults

`func NewCloudTestResultWithDefaults() *CloudTestResult`

NewCloudTestResultWithDefaults instantiates a new CloudTestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelivered

`func (o *CloudTestResult) GetDelivered() bool`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *CloudTestResult) GetDeliveredOk() (*bool, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *CloudTestResult) SetDelivered(v bool)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *CloudTestResult) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetDurationMs

`func (o *CloudTestResult) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *CloudTestResult) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *CloudTestResult) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *CloudTestResult) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetError

`func (o *CloudTestResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CloudTestResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CloudTestResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CloudTestResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetHttpStatus

`func (o *CloudTestResult) GetHttpStatus() int32`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *CloudTestResult) GetHttpStatusOk() (*int32, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *CloudTestResult) SetHttpStatus(v int32)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *CloudTestResult) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


