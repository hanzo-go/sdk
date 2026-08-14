# O11yO11yLogFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]O11yO11yLogFilterItem**](O11yO11yLogFilterItem.md) | Items are the predicates. | [optional] 
**Op** | Pointer to **string** | Op combines the items: AND or OR. | [optional] 

## Methods

### NewO11yO11yLogFilter

`func NewO11yO11yLogFilter() *O11yO11yLogFilter`

NewO11yO11yLogFilter instantiates a new O11yO11yLogFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLogFilterWithDefaults

`func NewO11yO11yLogFilterWithDefaults() *O11yO11yLogFilter`

NewO11yO11yLogFilterWithDefaults instantiates a new O11yO11yLogFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *O11yO11yLogFilter) GetItems() []O11yO11yLogFilterItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *O11yO11yLogFilter) GetItemsOk() (*[]O11yO11yLogFilterItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *O11yO11yLogFilter) SetItems(v []O11yO11yLogFilterItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *O11yO11yLogFilter) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOp

`func (o *O11yO11yLogFilter) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *O11yO11yLogFilter) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *O11yO11yLogFilter) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *O11yO11yLogFilter) HasOp() bool`

HasOp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


