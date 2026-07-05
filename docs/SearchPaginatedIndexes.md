# SearchPaginatedIndexes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]SearchIndexView**](SearchIndexView.md) |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewSearchPaginatedIndexes

`func NewSearchPaginatedIndexes() *SearchPaginatedIndexes`

NewSearchPaginatedIndexes instantiates a new SearchPaginatedIndexes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchPaginatedIndexesWithDefaults

`func NewSearchPaginatedIndexesWithDefaults() *SearchPaginatedIndexes`

NewSearchPaginatedIndexesWithDefaults instantiates a new SearchPaginatedIndexes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *SearchPaginatedIndexes) GetResults() []SearchIndexView`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchPaginatedIndexes) GetResultsOk() (*[]SearchIndexView, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchPaginatedIndexes) SetResults(v []SearchIndexView)`

SetResults sets Results field to given value.

### HasResults

`func (o *SearchPaginatedIndexes) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetOffset

`func (o *SearchPaginatedIndexes) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchPaginatedIndexes) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchPaginatedIndexes) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SearchPaginatedIndexes) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *SearchPaginatedIndexes) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchPaginatedIndexes) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchPaginatedIndexes) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchPaginatedIndexes) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotal

`func (o *SearchPaginatedIndexes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SearchPaginatedIndexes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SearchPaginatedIndexes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SearchPaginatedIndexes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


