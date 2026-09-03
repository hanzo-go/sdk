# O11yPostableNodes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int64** |  | [optional] 
**Filter** | Pointer to [**O11yFilter**](O11yFilter.md) |  | [optional] 
**GroupBy** | Pointer to [**[]O11yGroupByKey**](O11yGroupByKey.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**OrderBy** | Pointer to [**O11yQuerybuildertypesv5OrderBy**](O11yQuerybuildertypesv5OrderBy.md) |  | [optional] 
**Start** | Pointer to **int64** |  | [optional] 

## Methods

### NewO11yPostableNodes

`func NewO11yPostableNodes() *O11yPostableNodes`

NewO11yPostableNodes instantiates a new O11yPostableNodes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yPostableNodesWithDefaults

`func NewO11yPostableNodesWithDefaults() *O11yPostableNodes`

NewO11yPostableNodesWithDefaults instantiates a new O11yPostableNodes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *O11yPostableNodes) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yPostableNodes) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yPostableNodes) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *O11yPostableNodes) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFilter

`func (o *O11yPostableNodes) GetFilter() O11yFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *O11yPostableNodes) GetFilterOk() (*O11yFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *O11yPostableNodes) SetFilter(v O11yFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *O11yPostableNodes) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetGroupBy

`func (o *O11yPostableNodes) GetGroupBy() []O11yGroupByKey`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *O11yPostableNodes) GetGroupByOk() (*[]O11yGroupByKey, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *O11yPostableNodes) SetGroupBy(v []O11yGroupByKey)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *O11yPostableNodes) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetLimit

`func (o *O11yPostableNodes) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *O11yPostableNodes) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *O11yPostableNodes) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *O11yPostableNodes) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *O11yPostableNodes) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *O11yPostableNodes) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *O11yPostableNodes) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *O11yPostableNodes) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrderBy

`func (o *O11yPostableNodes) GetOrderBy() O11yQuerybuildertypesv5OrderBy`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *O11yPostableNodes) GetOrderByOk() (*O11yQuerybuildertypesv5OrderBy, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *O11yPostableNodes) SetOrderBy(v O11yQuerybuildertypesv5OrderBy)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *O11yPostableNodes) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetStart

`func (o *O11yPostableNodes) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yPostableNodes) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yPostableNodes) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *O11yPostableNodes) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


