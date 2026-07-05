# SearchIndexStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfDocuments** | Pointer to **int32** |  | [optional] 
**RawDocumentDbSize** | Pointer to **int64** |  | [optional] 
**MaxDocumentSize** | Pointer to **int64** |  | [optional] 
**AvgDocumentSize** | Pointer to **int64** |  | [optional] 
**IsIndexing** | Pointer to **bool** |  | [optional] 
**FieldDistribution** | Pointer to **map[string]int32** |  | [optional] 

## Methods

### NewSearchIndexStats

`func NewSearchIndexStats() *SearchIndexStats`

NewSearchIndexStats instantiates a new SearchIndexStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIndexStatsWithDefaults

`func NewSearchIndexStatsWithDefaults() *SearchIndexStats`

NewSearchIndexStatsWithDefaults instantiates a new SearchIndexStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfDocuments

`func (o *SearchIndexStats) GetNumberOfDocuments() int32`

GetNumberOfDocuments returns the NumberOfDocuments field if non-nil, zero value otherwise.

### GetNumberOfDocumentsOk

`func (o *SearchIndexStats) GetNumberOfDocumentsOk() (*int32, bool)`

GetNumberOfDocumentsOk returns a tuple with the NumberOfDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDocuments

`func (o *SearchIndexStats) SetNumberOfDocuments(v int32)`

SetNumberOfDocuments sets NumberOfDocuments field to given value.

### HasNumberOfDocuments

`func (o *SearchIndexStats) HasNumberOfDocuments() bool`

HasNumberOfDocuments returns a boolean if a field has been set.

### GetRawDocumentDbSize

`func (o *SearchIndexStats) GetRawDocumentDbSize() int64`

GetRawDocumentDbSize returns the RawDocumentDbSize field if non-nil, zero value otherwise.

### GetRawDocumentDbSizeOk

`func (o *SearchIndexStats) GetRawDocumentDbSizeOk() (*int64, bool)`

GetRawDocumentDbSizeOk returns a tuple with the RawDocumentDbSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawDocumentDbSize

`func (o *SearchIndexStats) SetRawDocumentDbSize(v int64)`

SetRawDocumentDbSize sets RawDocumentDbSize field to given value.

### HasRawDocumentDbSize

`func (o *SearchIndexStats) HasRawDocumentDbSize() bool`

HasRawDocumentDbSize returns a boolean if a field has been set.

### GetMaxDocumentSize

`func (o *SearchIndexStats) GetMaxDocumentSize() int64`

GetMaxDocumentSize returns the MaxDocumentSize field if non-nil, zero value otherwise.

### GetMaxDocumentSizeOk

`func (o *SearchIndexStats) GetMaxDocumentSizeOk() (*int64, bool)`

GetMaxDocumentSizeOk returns a tuple with the MaxDocumentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDocumentSize

`func (o *SearchIndexStats) SetMaxDocumentSize(v int64)`

SetMaxDocumentSize sets MaxDocumentSize field to given value.

### HasMaxDocumentSize

`func (o *SearchIndexStats) HasMaxDocumentSize() bool`

HasMaxDocumentSize returns a boolean if a field has been set.

### GetAvgDocumentSize

`func (o *SearchIndexStats) GetAvgDocumentSize() int64`

GetAvgDocumentSize returns the AvgDocumentSize field if non-nil, zero value otherwise.

### GetAvgDocumentSizeOk

`func (o *SearchIndexStats) GetAvgDocumentSizeOk() (*int64, bool)`

GetAvgDocumentSizeOk returns a tuple with the AvgDocumentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDocumentSize

`func (o *SearchIndexStats) SetAvgDocumentSize(v int64)`

SetAvgDocumentSize sets AvgDocumentSize field to given value.

### HasAvgDocumentSize

`func (o *SearchIndexStats) HasAvgDocumentSize() bool`

HasAvgDocumentSize returns a boolean if a field has been set.

### GetIsIndexing

`func (o *SearchIndexStats) GetIsIndexing() bool`

GetIsIndexing returns the IsIndexing field if non-nil, zero value otherwise.

### GetIsIndexingOk

`func (o *SearchIndexStats) GetIsIndexingOk() (*bool, bool)`

GetIsIndexingOk returns a tuple with the IsIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIndexing

`func (o *SearchIndexStats) SetIsIndexing(v bool)`

SetIsIndexing sets IsIndexing field to given value.

### HasIsIndexing

`func (o *SearchIndexStats) HasIsIndexing() bool`

HasIsIndexing returns a boolean if a field has been set.

### GetFieldDistribution

`func (o *SearchIndexStats) GetFieldDistribution() map[string]int32`

GetFieldDistribution returns the FieldDistribution field if non-nil, zero value otherwise.

### GetFieldDistributionOk

`func (o *SearchIndexStats) GetFieldDistributionOk() (*map[string]int32, bool)`

GetFieldDistributionOk returns a tuple with the FieldDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldDistribution

`func (o *SearchIndexStats) SetFieldDistribution(v map[string]int32)`

SetFieldDistribution sets FieldDistribution field to given value.

### HasFieldDistribution

`func (o *SearchIndexStats) HasFieldDistribution() bool`

HasFieldDistribution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


