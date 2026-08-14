# InvocationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invocations** | Pointer to [**[]InvocationView**](InvocationView.md) | Invocations is one row per past run, newest first. | [optional] 

## Methods

### NewInvocationList

`func NewInvocationList() *InvocationList`

NewInvocationList instantiates a new InvocationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvocationListWithDefaults

`func NewInvocationListWithDefaults() *InvocationList`

NewInvocationListWithDefaults instantiates a new InvocationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvocations

`func (o *InvocationList) GetInvocations() []InvocationView`

GetInvocations returns the Invocations field if non-nil, zero value otherwise.

### GetInvocationsOk

`func (o *InvocationList) GetInvocationsOk() (*[]InvocationView, bool)`

GetInvocationsOk returns a tuple with the Invocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocations

`func (o *InvocationList) SetInvocations(v []InvocationView)`

SetInvocations sets Invocations field to given value.

### HasInvocations

`func (o *InvocationList) HasInvocations() bool`

HasInvocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


