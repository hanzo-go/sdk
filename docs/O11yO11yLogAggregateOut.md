# O11yO11yLogAggregateOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**map[string]O11yO11yLogAggregateBucket**](O11yO11yLogAggregateBucket.md) | Items are the buckets, keyed by bucket timestamp. | [optional] 

## Methods

### NewO11yO11yLogAggregateOut

`func NewO11yO11yLogAggregateOut() *O11yO11yLogAggregateOut`

NewO11yO11yLogAggregateOut instantiates a new O11yO11yLogAggregateOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLogAggregateOutWithDefaults

`func NewO11yO11yLogAggregateOutWithDefaults() *O11yO11yLogAggregateOut`

NewO11yO11yLogAggregateOutWithDefaults instantiates a new O11yO11yLogAggregateOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *O11yO11yLogAggregateOut) GetItems() map[string]O11yO11yLogAggregateBucket`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *O11yO11yLogAggregateOut) GetItemsOk() (*map[string]O11yO11yLogAggregateBucket, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *O11yO11yLogAggregateOut) SetItems(v map[string]O11yO11yLogAggregateBucket)`

SetItems sets Items field to given value.

### HasItems

`func (o *O11yO11yLogAggregateOut) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


