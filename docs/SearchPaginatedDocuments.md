# SearchPaginatedDocuments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewSearchPaginatedDocuments

`func NewSearchPaginatedDocuments() *SearchPaginatedDocuments`

NewSearchPaginatedDocuments instantiates a new SearchPaginatedDocuments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchPaginatedDocumentsWithDefaults

`func NewSearchPaginatedDocumentsWithDefaults() *SearchPaginatedDocuments`

NewSearchPaginatedDocumentsWithDefaults instantiates a new SearchPaginatedDocuments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *SearchPaginatedDocuments) GetResults() []map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchPaginatedDocuments) GetResultsOk() (*[]map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchPaginatedDocuments) SetResults(v []map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *SearchPaginatedDocuments) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetOffset

`func (o *SearchPaginatedDocuments) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchPaginatedDocuments) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchPaginatedDocuments) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SearchPaginatedDocuments) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *SearchPaginatedDocuments) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchPaginatedDocuments) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchPaginatedDocuments) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchPaginatedDocuments) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotal

`func (o *SearchPaginatedDocuments) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SearchPaginatedDocuments) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SearchPaginatedDocuments) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SearchPaginatedDocuments) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


