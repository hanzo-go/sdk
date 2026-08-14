# O11yO11yFilterSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]O11yO11yFilterItem**](O11yO11yFilterItem.md) | Items are the conditions. | [optional] 
**Op** | Pointer to **string** | Operator combines the items — AND or OR. | [optional] 

## Methods

### NewO11yO11yFilterSet

`func NewO11yO11yFilterSet() *O11yO11yFilterSet`

NewO11yO11yFilterSet instantiates a new O11yO11yFilterSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yFilterSetWithDefaults

`func NewO11yO11yFilterSetWithDefaults() *O11yO11yFilterSet`

NewO11yO11yFilterSetWithDefaults instantiates a new O11yO11yFilterSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *O11yO11yFilterSet) GetItems() []O11yO11yFilterItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *O11yO11yFilterSet) GetItemsOk() (*[]O11yO11yFilterItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *O11yO11yFilterSet) SetItems(v []O11yO11yFilterItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *O11yO11yFilterSet) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOp

`func (o *O11yO11yFilterSet) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *O11yO11yFilterSet) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *O11yO11yFilterSet) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *O11yO11yFilterSet) HasOp() bool`

HasOp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


