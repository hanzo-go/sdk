# TargetList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | Pointer to [**[]TargetView**](TargetView.md) | Targets is every target registered to the caller&#39;s org. | [optional] 

## Methods

### NewTargetList

`func NewTargetList() *TargetList`

NewTargetList instantiates a new TargetList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetListWithDefaults

`func NewTargetListWithDefaults() *TargetList`

NewTargetListWithDefaults instantiates a new TargetList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *TargetList) GetTargets() []TargetView`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *TargetList) GetTargetsOk() (*[]TargetView, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *TargetList) SetTargets(v []TargetView)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *TargetList) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


