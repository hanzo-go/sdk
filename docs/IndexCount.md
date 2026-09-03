# IndexCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsIndexing** | Pointer to **bool** | IsIndexing is always false: writes are applied before their response, so there is never a background pass a caller could be waiting on. | [optional] 
**NumberOfDocuments** | Pointer to **int64** | NumberOfDocuments is how many documents this org holds in that index. | [optional] 

## Methods

### NewIndexCount

`func NewIndexCount() *IndexCount`

NewIndexCount instantiates a new IndexCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexCountWithDefaults

`func NewIndexCountWithDefaults() *IndexCount`

NewIndexCountWithDefaults instantiates a new IndexCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsIndexing

`func (o *IndexCount) GetIsIndexing() bool`

GetIsIndexing returns the IsIndexing field if non-nil, zero value otherwise.

### GetIsIndexingOk

`func (o *IndexCount) GetIsIndexingOk() (*bool, bool)`

GetIsIndexingOk returns a tuple with the IsIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIndexing

`func (o *IndexCount) SetIsIndexing(v bool)`

SetIsIndexing sets IsIndexing field to given value.

### HasIsIndexing

`func (o *IndexCount) HasIsIndexing() bool`

HasIsIndexing returns a boolean if a field has been set.

### GetNumberOfDocuments

`func (o *IndexCount) GetNumberOfDocuments() int64`

GetNumberOfDocuments returns the NumberOfDocuments field if non-nil, zero value otherwise.

### GetNumberOfDocumentsOk

`func (o *IndexCount) GetNumberOfDocumentsOk() (*int64, bool)`

GetNumberOfDocumentsOk returns a tuple with the NumberOfDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDocuments

`func (o *IndexCount) SetNumberOfDocuments(v int64)`

SetNumberOfDocuments sets NumberOfDocuments field to given value.

### HasNumberOfDocuments

`func (o *IndexCount) HasNumberOfDocuments() bool`

HasNumberOfDocuments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


