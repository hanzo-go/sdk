# O11yProcessListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int64** | epoch time in ms | [optional] 
**Filters** | Pointer to [**O11yFilterSet**](O11yFilterSet.md) |  | [optional] 
**GroupBy** | Pointer to [**[]O11yAttributeKey**](O11yAttributeKey.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**OrderBy** | Pointer to [**O11yOrderBy**](O11yOrderBy.md) |  | [optional] 
**Start** | Pointer to **int64** | epoch time in ms | [optional] 

## Methods

### NewO11yProcessListRequest

`func NewO11yProcessListRequest() *O11yProcessListRequest`

NewO11yProcessListRequest instantiates a new O11yProcessListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yProcessListRequestWithDefaults

`func NewO11yProcessListRequestWithDefaults() *O11yProcessListRequest`

NewO11yProcessListRequestWithDefaults instantiates a new O11yProcessListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *O11yProcessListRequest) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yProcessListRequest) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yProcessListRequest) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *O11yProcessListRequest) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFilters

`func (o *O11yProcessListRequest) GetFilters() O11yFilterSet`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *O11yProcessListRequest) GetFiltersOk() (*O11yFilterSet, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *O11yProcessListRequest) SetFilters(v O11yFilterSet)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *O11yProcessListRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetGroupBy

`func (o *O11yProcessListRequest) GetGroupBy() []O11yAttributeKey`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *O11yProcessListRequest) GetGroupByOk() (*[]O11yAttributeKey, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *O11yProcessListRequest) SetGroupBy(v []O11yAttributeKey)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *O11yProcessListRequest) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetLimit

`func (o *O11yProcessListRequest) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *O11yProcessListRequest) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *O11yProcessListRequest) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *O11yProcessListRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *O11yProcessListRequest) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *O11yProcessListRequest) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *O11yProcessListRequest) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *O11yProcessListRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrderBy

`func (o *O11yProcessListRequest) GetOrderBy() O11yOrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *O11yProcessListRequest) GetOrderByOk() (*O11yOrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *O11yProcessListRequest) SetOrderBy(v O11yOrderBy)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *O11yProcessListRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetStart

`func (o *O11yProcessListRequest) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yProcessListRequest) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yProcessListRequest) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *O11yProcessListRequest) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


