# IndexDocuments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int64** | Limit is how many documents this page could hold. | [optional] 
**Offset** | Pointer to **int64** | Offset is where this page starts. | [optional] 
**Results** | Pointer to **[]interface{}** | Results are the documents themselves, exactly as they were stored. | [optional] 
**Total** | Pointer to **int64** | Total is how many documents the index holds altogether. | [optional] 

## Methods

### NewIndexDocuments

`func NewIndexDocuments() *IndexDocuments`

NewIndexDocuments instantiates a new IndexDocuments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexDocumentsWithDefaults

`func NewIndexDocumentsWithDefaults() *IndexDocuments`

NewIndexDocumentsWithDefaults instantiates a new IndexDocuments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *IndexDocuments) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *IndexDocuments) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *IndexDocuments) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *IndexDocuments) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *IndexDocuments) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *IndexDocuments) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *IndexDocuments) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *IndexDocuments) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetResults

`func (o *IndexDocuments) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *IndexDocuments) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *IndexDocuments) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *IndexDocuments) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotal

`func (o *IndexDocuments) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IndexDocuments) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IndexDocuments) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *IndexDocuments) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


