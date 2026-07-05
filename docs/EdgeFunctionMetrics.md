# EdgeFunctionMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionId** | Pointer to **string** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**Invocations** | Pointer to **int64** |  | [optional] 
**Errors** | Pointer to **int64** |  | [optional] 
**AvgDurationMs** | Pointer to **float32** |  | [optional] 
**P95DurationMs** | Pointer to **float32** |  | [optional] 
**P99DurationMs** | Pointer to **float32** |  | [optional] 
**DataInBytes** | Pointer to **int64** |  | [optional] 
**DataOutBytes** | Pointer to **int64** |  | [optional] 

## Methods

### NewEdgeFunctionMetrics

`func NewEdgeFunctionMetrics() *EdgeFunctionMetrics`

NewEdgeFunctionMetrics instantiates a new EdgeFunctionMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeFunctionMetricsWithDefaults

`func NewEdgeFunctionMetricsWithDefaults() *EdgeFunctionMetrics`

NewEdgeFunctionMetricsWithDefaults instantiates a new EdgeFunctionMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionId

`func (o *EdgeFunctionMetrics) GetFunctionId() string`

GetFunctionId returns the FunctionId field if non-nil, zero value otherwise.

### GetFunctionIdOk

`func (o *EdgeFunctionMetrics) GetFunctionIdOk() (*string, bool)`

GetFunctionIdOk returns a tuple with the FunctionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionId

`func (o *EdgeFunctionMetrics) SetFunctionId(v string)`

SetFunctionId sets FunctionId field to given value.

### HasFunctionId

`func (o *EdgeFunctionMetrics) HasFunctionId() bool`

HasFunctionId returns a boolean if a field has been set.

### GetPeriod

`func (o *EdgeFunctionMetrics) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *EdgeFunctionMetrics) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *EdgeFunctionMetrics) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *EdgeFunctionMetrics) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetInvocations

`func (o *EdgeFunctionMetrics) GetInvocations() int64`

GetInvocations returns the Invocations field if non-nil, zero value otherwise.

### GetInvocationsOk

`func (o *EdgeFunctionMetrics) GetInvocationsOk() (*int64, bool)`

GetInvocationsOk returns a tuple with the Invocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocations

`func (o *EdgeFunctionMetrics) SetInvocations(v int64)`

SetInvocations sets Invocations field to given value.

### HasInvocations

`func (o *EdgeFunctionMetrics) HasInvocations() bool`

HasInvocations returns a boolean if a field has been set.

### GetErrors

`func (o *EdgeFunctionMetrics) GetErrors() int64`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *EdgeFunctionMetrics) GetErrorsOk() (*int64, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *EdgeFunctionMetrics) SetErrors(v int64)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *EdgeFunctionMetrics) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetAvgDurationMs

`func (o *EdgeFunctionMetrics) GetAvgDurationMs() float32`

GetAvgDurationMs returns the AvgDurationMs field if non-nil, zero value otherwise.

### GetAvgDurationMsOk

`func (o *EdgeFunctionMetrics) GetAvgDurationMsOk() (*float32, bool)`

GetAvgDurationMsOk returns a tuple with the AvgDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDurationMs

`func (o *EdgeFunctionMetrics) SetAvgDurationMs(v float32)`

SetAvgDurationMs sets AvgDurationMs field to given value.

### HasAvgDurationMs

`func (o *EdgeFunctionMetrics) HasAvgDurationMs() bool`

HasAvgDurationMs returns a boolean if a field has been set.

### GetP95DurationMs

`func (o *EdgeFunctionMetrics) GetP95DurationMs() float32`

GetP95DurationMs returns the P95DurationMs field if non-nil, zero value otherwise.

### GetP95DurationMsOk

`func (o *EdgeFunctionMetrics) GetP95DurationMsOk() (*float32, bool)`

GetP95DurationMsOk returns a tuple with the P95DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP95DurationMs

`func (o *EdgeFunctionMetrics) SetP95DurationMs(v float32)`

SetP95DurationMs sets P95DurationMs field to given value.

### HasP95DurationMs

`func (o *EdgeFunctionMetrics) HasP95DurationMs() bool`

HasP95DurationMs returns a boolean if a field has been set.

### GetP99DurationMs

`func (o *EdgeFunctionMetrics) GetP99DurationMs() float32`

GetP99DurationMs returns the P99DurationMs field if non-nil, zero value otherwise.

### GetP99DurationMsOk

`func (o *EdgeFunctionMetrics) GetP99DurationMsOk() (*float32, bool)`

GetP99DurationMsOk returns a tuple with the P99DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP99DurationMs

`func (o *EdgeFunctionMetrics) SetP99DurationMs(v float32)`

SetP99DurationMs sets P99DurationMs field to given value.

### HasP99DurationMs

`func (o *EdgeFunctionMetrics) HasP99DurationMs() bool`

HasP99DurationMs returns a boolean if a field has been set.

### GetDataInBytes

`func (o *EdgeFunctionMetrics) GetDataInBytes() int64`

GetDataInBytes returns the DataInBytes field if non-nil, zero value otherwise.

### GetDataInBytesOk

`func (o *EdgeFunctionMetrics) GetDataInBytesOk() (*int64, bool)`

GetDataInBytesOk returns a tuple with the DataInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataInBytes

`func (o *EdgeFunctionMetrics) SetDataInBytes(v int64)`

SetDataInBytes sets DataInBytes field to given value.

### HasDataInBytes

`func (o *EdgeFunctionMetrics) HasDataInBytes() bool`

HasDataInBytes returns a boolean if a field has been set.

### GetDataOutBytes

`func (o *EdgeFunctionMetrics) GetDataOutBytes() int64`

GetDataOutBytes returns the DataOutBytes field if non-nil, zero value otherwise.

### GetDataOutBytesOk

`func (o *EdgeFunctionMetrics) GetDataOutBytesOk() (*int64, bool)`

GetDataOutBytesOk returns a tuple with the DataOutBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataOutBytes

`func (o *EdgeFunctionMetrics) SetDataOutBytes(v int64)`

SetDataOutBytes sets DataOutBytes field to given value.

### HasDataOutBytes

`func (o *EdgeFunctionMetrics) HasDataOutBytes() bool`

HasDataOutBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


