# IndexIndexStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseSize** | Pointer to **int64** | DatabaseSize is the org&#39;s total document count across its indexes. It is a count, not bytes: the store is shared by every tenant, so a byte figure would either be the whole file (another tenant&#39;s size) or a fiction. | [optional] 
**Indexes** | Pointer to [**map[string]IndexIndexCount**](IndexIndexCount.md) | Indexes maps each index uid to its own count. | [optional] 

## Methods

### NewIndexIndexStats

`func NewIndexIndexStats() *IndexIndexStats`

NewIndexIndexStats instantiates a new IndexIndexStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexIndexStatsWithDefaults

`func NewIndexIndexStatsWithDefaults() *IndexIndexStats`

NewIndexIndexStatsWithDefaults instantiates a new IndexIndexStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseSize

`func (o *IndexIndexStats) GetDatabaseSize() int64`

GetDatabaseSize returns the DatabaseSize field if non-nil, zero value otherwise.

### GetDatabaseSizeOk

`func (o *IndexIndexStats) GetDatabaseSizeOk() (*int64, bool)`

GetDatabaseSizeOk returns a tuple with the DatabaseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseSize

`func (o *IndexIndexStats) SetDatabaseSize(v int64)`

SetDatabaseSize sets DatabaseSize field to given value.

### HasDatabaseSize

`func (o *IndexIndexStats) HasDatabaseSize() bool`

HasDatabaseSize returns a boolean if a field has been set.

### GetIndexes

`func (o *IndexIndexStats) GetIndexes() map[string]IndexIndexCount`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *IndexIndexStats) GetIndexesOk() (*map[string]IndexIndexCount, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *IndexIndexStats) SetIndexes(v map[string]IndexIndexCount)`

SetIndexes sets Indexes field to given value.

### HasIndexes

`func (o *IndexIndexStats) HasIndexes() bool`

HasIndexes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


