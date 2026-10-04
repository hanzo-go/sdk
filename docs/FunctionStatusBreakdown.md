# FunctionStatusBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **int64** | invocations that ran and failed | [optional] 
**Success** | Pointer to **int64** | invocations whose code ran and wrote nothing to stderr | [optional] 
**Timeout** | Pointer to **int64** | invocations that hit their configured deadline | [optional] 

## Methods

### NewFunctionStatusBreakdown

`func NewFunctionStatusBreakdown() *FunctionStatusBreakdown`

NewFunctionStatusBreakdown instantiates a new FunctionStatusBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionStatusBreakdownWithDefaults

`func NewFunctionStatusBreakdownWithDefaults() *FunctionStatusBreakdown`

NewFunctionStatusBreakdownWithDefaults instantiates a new FunctionStatusBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *FunctionStatusBreakdown) GetError() int64`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FunctionStatusBreakdown) GetErrorOk() (*int64, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FunctionStatusBreakdown) SetError(v int64)`

SetError sets Error field to given value.

### HasError

`func (o *FunctionStatusBreakdown) HasError() bool`

HasError returns a boolean if a field has been set.

### GetSuccess

`func (o *FunctionStatusBreakdown) GetSuccess() int64`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *FunctionStatusBreakdown) GetSuccessOk() (*int64, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *FunctionStatusBreakdown) SetSuccess(v int64)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *FunctionStatusBreakdown) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetTimeout

`func (o *FunctionStatusBreakdown) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *FunctionStatusBreakdown) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *FunctionStatusBreakdown) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *FunctionStatusBreakdown) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


