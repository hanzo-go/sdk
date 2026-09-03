# IndexStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseSize** | Pointer to **int64** | DatabaseSize is the org&#39;s total document count across its indexes. It is a count, not bytes: the store is shared by every tenant, so a byte figure would either be the whole file (another tenant&#39;s size) or a fiction. | [optional] 
**Indexes** | Pointer to [**map[string]IndexCount**](IndexCount.md) | Indexes maps each index uid to its own count. | [optional] 

## Methods

### NewIndexStats

`func NewIndexStats() *IndexStats`

NewIndexStats instantiates a new IndexStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexStatsWithDefaults

`func NewIndexStatsWithDefaults() *IndexStats`

NewIndexStatsWithDefaults instantiates a new IndexStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseSize

`func (o *IndexStats) GetDatabaseSize() int64`

GetDatabaseSize returns the DatabaseSize field if non-nil, zero value otherwise.

### GetDatabaseSizeOk

`func (o *IndexStats) GetDatabaseSizeOk() (*int64, bool)`

GetDatabaseSizeOk returns a tuple with the DatabaseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseSize

`func (o *IndexStats) SetDatabaseSize(v int64)`

SetDatabaseSize sets DatabaseSize field to given value.

### HasDatabaseSize

`func (o *IndexStats) HasDatabaseSize() bool`

HasDatabaseSize returns a boolean if a field has been set.

### GetIndexes

`func (o *IndexStats) GetIndexes() map[string]IndexCount`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *IndexStats) GetIndexesOk() (*map[string]IndexCount, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *IndexStats) SetIndexes(v map[string]IndexCount)`

SetIndexes sets Indexes field to given value.

### HasIndexes

`func (o *IndexStats) HasIndexes() bool`

HasIndexes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


