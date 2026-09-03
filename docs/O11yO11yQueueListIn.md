# O11yO11yQueueListIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int64** | End is the window&#39;s end, epoch nanoseconds. | [optional] 
**Filters** | Pointer to [**O11yO11yQueueFilterSet**](O11yO11yQueueFilterSet.md) | Filters narrow the rows by span attribute; null means all rows. | [optional] 
**Limit** | Pointer to **int64** | Limit caps how many rows come back. | [optional] 
**Start** | Pointer to **int64** | Start is the window&#39;s start, epoch nanoseconds. | [optional] 

## Methods

### NewO11yO11yQueueListIn

`func NewO11yO11yQueueListIn() *O11yO11yQueueListIn`

NewO11yO11yQueueListIn instantiates a new O11yO11yQueueListIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yQueueListInWithDefaults

`func NewO11yO11yQueueListInWithDefaults() *O11yO11yQueueListIn`

NewO11yO11yQueueListInWithDefaults instantiates a new O11yO11yQueueListIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *O11yO11yQueueListIn) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yO11yQueueListIn) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yO11yQueueListIn) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *O11yO11yQueueListIn) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFilters

`func (o *O11yO11yQueueListIn) GetFilters() O11yO11yQueueFilterSet`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *O11yO11yQueueListIn) GetFiltersOk() (*O11yO11yQueueFilterSet, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *O11yO11yQueueListIn) SetFilters(v O11yO11yQueueFilterSet)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *O11yO11yQueueListIn) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLimit

`func (o *O11yO11yQueueListIn) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *O11yO11yQueueListIn) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *O11yO11yQueueListIn) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *O11yO11yQueueListIn) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetStart

`func (o *O11yO11yQueueListIn) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yO11yQueueListIn) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yO11yQueueListIn) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *O11yO11yQueueListIn) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


