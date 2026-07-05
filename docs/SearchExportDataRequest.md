# SearchExportDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indexes** | Pointer to **[]string** |  | [optional] 
**SkipEmbeddings** | Pointer to **bool** |  | [optional] 

## Methods

### NewSearchExportDataRequest

`func NewSearchExportDataRequest() *SearchExportDataRequest`

NewSearchExportDataRequest instantiates a new SearchExportDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchExportDataRequestWithDefaults

`func NewSearchExportDataRequestWithDefaults() *SearchExportDataRequest`

NewSearchExportDataRequestWithDefaults instantiates a new SearchExportDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexes

`func (o *SearchExportDataRequest) GetIndexes() []string`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *SearchExportDataRequest) GetIndexesOk() (*[]string, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *SearchExportDataRequest) SetIndexes(v []string)`

SetIndexes sets Indexes field to given value.

### HasIndexes

`func (o *SearchExportDataRequest) HasIndexes() bool`

HasIndexes returns a boolean if a field has been set.

### GetSkipEmbeddings

`func (o *SearchExportDataRequest) GetSkipEmbeddings() bool`

GetSkipEmbeddings returns the SkipEmbeddings field if non-nil, zero value otherwise.

### GetSkipEmbeddingsOk

`func (o *SearchExportDataRequest) GetSkipEmbeddingsOk() (*bool, bool)`

GetSkipEmbeddingsOk returns a tuple with the SkipEmbeddings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipEmbeddings

`func (o *SearchExportDataRequest) SetSkipEmbeddings(v bool)`

SetSkipEmbeddings sets SkipEmbeddings field to given value.

### HasSkipEmbeddings

`func (o *SearchExportDataRequest) HasSkipEmbeddings() bool`

HasSkipEmbeddings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


