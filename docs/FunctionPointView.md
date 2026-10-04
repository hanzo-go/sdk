# FunctionPointView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**T** | Pointer to **string** | the bucket&#39;s start, RFC3339 (UTC) | [optional] 
**V** | Pointer to **int64** | how many invocations fell in it — a real count, never interpolated | [optional] 

## Methods

### NewFunctionPointView

`func NewFunctionPointView() *FunctionPointView`

NewFunctionPointView instantiates a new FunctionPointView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionPointViewWithDefaults

`func NewFunctionPointViewWithDefaults() *FunctionPointView`

NewFunctionPointViewWithDefaults instantiates a new FunctionPointView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetT

`func (o *FunctionPointView) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *FunctionPointView) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *FunctionPointView) SetT(v string)`

SetT sets T field to given value.

### HasT

`func (o *FunctionPointView) HasT() bool`

HasT returns a boolean if a field has been set.

### GetV

`func (o *FunctionPointView) GetV() int64`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *FunctionPointView) GetVOk() (*int64, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *FunctionPointView) SetV(v int64)`

SetV sets V field to given value.

### HasV

`func (o *FunctionPointView) HasV() bool`

HasV returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


