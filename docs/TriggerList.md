# TriggerList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Triggers** | Pointer to [**[]TriggerView**](TriggerView.md) | Triggers is one row per function, describing how it is reached. | [optional] 

## Methods

### NewTriggerList

`func NewTriggerList() *TriggerList`

NewTriggerList instantiates a new TriggerList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerListWithDefaults

`func NewTriggerListWithDefaults() *TriggerList`

NewTriggerListWithDefaults instantiates a new TriggerList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggers

`func (o *TriggerList) GetTriggers() []TriggerView`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *TriggerList) GetTriggersOk() (*[]TriggerView, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *TriggerList) SetTriggers(v []TriggerView)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *TriggerList) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


