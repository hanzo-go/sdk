# O11yJobListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int32** | epoch time in ms | [optional] 
**Filters** | Pointer to [**O11yFilterSet**](O11yFilterSet.md) |  | [optional] 
**GroupBy** | Pointer to [**[]O11yAttributeKey**](O11yAttributeKey.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**OrderBy** | Pointer to [**O11yOrderBy**](O11yOrderBy.md) |  | [optional] 
**Start** | Pointer to **int32** | epoch time in ms | [optional] 

## Methods

### NewO11yJobListRequest

`func NewO11yJobListRequest() *O11yJobListRequest`

NewO11yJobListRequest instantiates a new O11yJobListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yJobListRequestWithDefaults

`func NewO11yJobListRequestWithDefaults() *O11yJobListRequest`

NewO11yJobListRequestWithDefaults instantiates a new O11yJobListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *O11yJobListRequest) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yJobListRequest) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yJobListRequest) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *O11yJobListRequest) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFilters

`func (o *O11yJobListRequest) GetFilters() O11yFilterSet`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *O11yJobListRequest) GetFiltersOk() (*O11yFilterSet, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *O11yJobListRequest) SetFilters(v O11yFilterSet)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *O11yJobListRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetGroupBy

`func (o *O11yJobListRequest) GetGroupBy() []O11yAttributeKey`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *O11yJobListRequest) GetGroupByOk() (*[]O11yAttributeKey, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *O11yJobListRequest) SetGroupBy(v []O11yAttributeKey)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *O11yJobListRequest) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetLimit

`func (o *O11yJobListRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *O11yJobListRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *O11yJobListRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *O11yJobListRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *O11yJobListRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *O11yJobListRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *O11yJobListRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *O11yJobListRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrderBy

`func (o *O11yJobListRequest) GetOrderBy() O11yOrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *O11yJobListRequest) GetOrderByOk() (*O11yOrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *O11yJobListRequest) SetOrderBy(v O11yOrderBy)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *O11yJobListRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetStart

`func (o *O11yJobListRequest) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yJobListRequest) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yJobListRequest) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *O11yJobListRequest) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


