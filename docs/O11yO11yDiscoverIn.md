# O11yO11yDiscoverIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregations** | Pointer to **[]string** | Aggregations are the measures to compute per group. Empty means a count. | [optional] 
**Filters** | Pointer to [**[]O11yO11yFilter**](O11yO11yFilter.md) | Filters narrow the scan; each is a field, an operator (eq, neq, like) and a value, and they combine with AND. | [optional] 
**GroupBy** | Pointer to **[]string** | GroupBy are the columns to group the rows by. | [optional] 
**Limit** | Pointer to **int64** | Limit caps how many rows come back. | [optional] 
**OrderBy** | Pointer to **string** | OrderBy is the column or aggregation to sort the rows on. | [optional] 
**OrderDir** | Pointer to **string** | OrderDir is asc or desc. | [optional] 
**Period** | Pointer to **string** | Period is the window to read, relative to now — 1h, 24h, 7d, 14d, 30d. | [optional] 
**Project** | **string** | Project is the project to read, as its id. Required. | 

## Methods

### NewO11yO11yDiscoverIn

`func NewO11yO11yDiscoverIn(project string, ) *O11yO11yDiscoverIn`

NewO11yO11yDiscoverIn instantiates a new O11yO11yDiscoverIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDiscoverInWithDefaults

`func NewO11yO11yDiscoverInWithDefaults() *O11yO11yDiscoverIn`

NewO11yO11yDiscoverInWithDefaults instantiates a new O11yO11yDiscoverIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregations

`func (o *O11yO11yDiscoverIn) GetAggregations() []string`

GetAggregations returns the Aggregations field if non-nil, zero value otherwise.

### GetAggregationsOk

`func (o *O11yO11yDiscoverIn) GetAggregationsOk() (*[]string, bool)`

GetAggregationsOk returns a tuple with the Aggregations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregations

`func (o *O11yO11yDiscoverIn) SetAggregations(v []string)`

SetAggregations sets Aggregations field to given value.

### HasAggregations

`func (o *O11yO11yDiscoverIn) HasAggregations() bool`

HasAggregations returns a boolean if a field has been set.

### GetFilters

`func (o *O11yO11yDiscoverIn) GetFilters() []O11yO11yFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *O11yO11yDiscoverIn) GetFiltersOk() (*[]O11yO11yFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *O11yO11yDiscoverIn) SetFilters(v []O11yO11yFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *O11yO11yDiscoverIn) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetGroupBy

`func (o *O11yO11yDiscoverIn) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *O11yO11yDiscoverIn) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *O11yO11yDiscoverIn) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *O11yO11yDiscoverIn) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetLimit

`func (o *O11yO11yDiscoverIn) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *O11yO11yDiscoverIn) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *O11yO11yDiscoverIn) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *O11yO11yDiscoverIn) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOrderBy

`func (o *O11yO11yDiscoverIn) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *O11yO11yDiscoverIn) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *O11yO11yDiscoverIn) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *O11yO11yDiscoverIn) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetOrderDir

`func (o *O11yO11yDiscoverIn) GetOrderDir() string`

GetOrderDir returns the OrderDir field if non-nil, zero value otherwise.

### GetOrderDirOk

`func (o *O11yO11yDiscoverIn) GetOrderDirOk() (*string, bool)`

GetOrderDirOk returns a tuple with the OrderDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDir

`func (o *O11yO11yDiscoverIn) SetOrderDir(v string)`

SetOrderDir sets OrderDir field to given value.

### HasOrderDir

`func (o *O11yO11yDiscoverIn) HasOrderDir() bool`

HasOrderDir returns a boolean if a field has been set.

### GetPeriod

`func (o *O11yO11yDiscoverIn) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *O11yO11yDiscoverIn) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *O11yO11yDiscoverIn) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *O11yO11yDiscoverIn) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetProject

`func (o *O11yO11yDiscoverIn) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *O11yO11yDiscoverIn) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *O11yO11yDiscoverIn) SetProject(v string)`

SetProject sets Project field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


