# O11yO11yMetricStatsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int64** | End is the end of the window as a Unix timestamp in milliseconds. Required. | 
**Filter** | Pointer to [**O11yO11yMetricFilter**](O11yO11yMetricFilter.md) | Filter narrows the metrics counted. | [optional] 
**Limit** | **int64** | Limit caps how many metrics come back, between 1 and 5000. Required. | 
**Offset** | Pointer to **int64** | Offset is how many metrics to skip, for paging. | [optional] 
**OrderBy** | Pointer to [**O11yO11yMetricOrder**](O11yO11yMetricOrder.md) | OrderBy sorts the page, by samples or timeseries. | [optional] 
**Start** | **int64** | Start is the start of the window as a Unix timestamp in milliseconds. Required. | 

## Methods

### NewO11yO11yMetricStatsIn

`func NewO11yO11yMetricStatsIn(end int64, limit int64, start int64, ) *O11yO11yMetricStatsIn`

NewO11yO11yMetricStatsIn instantiates a new O11yO11yMetricStatsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yMetricStatsInWithDefaults

`func NewO11yO11yMetricStatsInWithDefaults() *O11yO11yMetricStatsIn`

NewO11yO11yMetricStatsInWithDefaults instantiates a new O11yO11yMetricStatsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *O11yO11yMetricStatsIn) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yO11yMetricStatsIn) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yO11yMetricStatsIn) SetEnd(v int64)`

SetEnd sets End field to given value.


### GetFilter

`func (o *O11yO11yMetricStatsIn) GetFilter() O11yO11yMetricFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *O11yO11yMetricStatsIn) GetFilterOk() (*O11yO11yMetricFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *O11yO11yMetricStatsIn) SetFilter(v O11yO11yMetricFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *O11yO11yMetricStatsIn) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetLimit

`func (o *O11yO11yMetricStatsIn) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *O11yO11yMetricStatsIn) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *O11yO11yMetricStatsIn) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *O11yO11yMetricStatsIn) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *O11yO11yMetricStatsIn) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *O11yO11yMetricStatsIn) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *O11yO11yMetricStatsIn) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrderBy

`func (o *O11yO11yMetricStatsIn) GetOrderBy() O11yO11yMetricOrder`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *O11yO11yMetricStatsIn) GetOrderByOk() (*O11yO11yMetricOrder, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *O11yO11yMetricStatsIn) SetOrderBy(v O11yO11yMetricOrder)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *O11yO11yMetricStatsIn) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetStart

`func (o *O11yO11yMetricStatsIn) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yO11yMetricStatsIn) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yO11yMetricStatsIn) SetStart(v int64)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


