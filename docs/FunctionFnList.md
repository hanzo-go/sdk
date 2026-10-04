# FunctionFnList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Functions** | Pointer to [**[]FunctionFunctionView**](FunctionFunctionView.md) | Functions is one row per published function. | [optional] 

## Methods

### NewFunctionFnList

`func NewFunctionFnList() *FunctionFnList`

NewFunctionFnList instantiates a new FunctionFnList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionFnListWithDefaults

`func NewFunctionFnListWithDefaults() *FunctionFnList`

NewFunctionFnListWithDefaults instantiates a new FunctionFnList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctions

`func (o *FunctionFnList) GetFunctions() []FunctionFunctionView`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *FunctionFnList) GetFunctionsOk() (*[]FunctionFunctionView, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *FunctionFnList) SetFunctions(v []FunctionFunctionView)`

SetFunctions sets Functions field to given value.

### HasFunctions

`func (o *FunctionFnList) HasFunctions() bool`

HasFunctions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


