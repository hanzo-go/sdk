# O11yO11yQueueFilterSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]O11yO11yQueueFilterRule**](O11yO11yQueueFilterRule.md) | Items are the predicates. | [optional] 
**Op** | Pointer to **string** | Op combines the items: AND or OR. | [optional] 

## Methods

### NewO11yO11yQueueFilterSet

`func NewO11yO11yQueueFilterSet() *O11yO11yQueueFilterSet`

NewO11yO11yQueueFilterSet instantiates a new O11yO11yQueueFilterSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yQueueFilterSetWithDefaults

`func NewO11yO11yQueueFilterSetWithDefaults() *O11yO11yQueueFilterSet`

NewO11yO11yQueueFilterSetWithDefaults instantiates a new O11yO11yQueueFilterSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *O11yO11yQueueFilterSet) GetItems() []O11yO11yQueueFilterRule`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *O11yO11yQueueFilterSet) GetItemsOk() (*[]O11yO11yQueueFilterRule, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *O11yO11yQueueFilterSet) SetItems(v []O11yO11yQueueFilterRule)`

SetItems sets Items field to given value.

### HasItems

`func (o *O11yO11yQueueFilterSet) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOp

`func (o *O11yO11yQueueFilterSet) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *O11yO11yQueueFilterSet) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *O11yO11yQueueFilterSet) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *O11yO11yQueueFilterSet) HasOp() bool`

HasOp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


