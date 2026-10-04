# FunctionInvocationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invocations** | Pointer to [**[]FunctionInvocationView**](FunctionInvocationView.md) | Invocations is one row per past run, newest first. | [optional] 

## Methods

### NewFunctionInvocationList

`func NewFunctionInvocationList() *FunctionInvocationList`

NewFunctionInvocationList instantiates a new FunctionInvocationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionInvocationListWithDefaults

`func NewFunctionInvocationListWithDefaults() *FunctionInvocationList`

NewFunctionInvocationListWithDefaults instantiates a new FunctionInvocationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvocations

`func (o *FunctionInvocationList) GetInvocations() []FunctionInvocationView`

GetInvocations returns the Invocations field if non-nil, zero value otherwise.

### GetInvocationsOk

`func (o *FunctionInvocationList) GetInvocationsOk() (*[]FunctionInvocationView, bool)`

GetInvocationsOk returns a tuple with the Invocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocations

`func (o *FunctionInvocationList) SetInvocations(v []FunctionInvocationView)`

SetInvocations sets Invocations field to given value.

### HasInvocations

`func (o *FunctionInvocationList) HasInvocations() bool`

HasInvocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


