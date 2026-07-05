# SearchPaginatedKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]SearchKeyView**](SearchKeyView.md) |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewSearchPaginatedKeys

`func NewSearchPaginatedKeys() *SearchPaginatedKeys`

NewSearchPaginatedKeys instantiates a new SearchPaginatedKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchPaginatedKeysWithDefaults

`func NewSearchPaginatedKeysWithDefaults() *SearchPaginatedKeys`

NewSearchPaginatedKeysWithDefaults instantiates a new SearchPaginatedKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *SearchPaginatedKeys) GetResults() []SearchKeyView`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchPaginatedKeys) GetResultsOk() (*[]SearchKeyView, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchPaginatedKeys) SetResults(v []SearchKeyView)`

SetResults sets Results field to given value.

### HasResults

`func (o *SearchPaginatedKeys) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetOffset

`func (o *SearchPaginatedKeys) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchPaginatedKeys) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchPaginatedKeys) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SearchPaginatedKeys) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *SearchPaginatedKeys) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchPaginatedKeys) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchPaginatedKeys) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchPaginatedKeys) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotal

`func (o *SearchPaginatedKeys) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SearchPaginatedKeys) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SearchPaginatedKeys) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SearchPaginatedKeys) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


