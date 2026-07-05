# SearchStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseSize** | Pointer to **int64** |  | [optional] 
**UsedDatabaseSize** | Pointer to **int64** |  | [optional] 
**LastUpdate** | Pointer to **time.Time** |  | [optional] 
**Indexes** | Pointer to [**map[string]SearchIndexStats**](SearchIndexStats.md) |  | [optional] 

## Methods

### NewSearchStats

`func NewSearchStats() *SearchStats`

NewSearchStats instantiates a new SearchStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchStatsWithDefaults

`func NewSearchStatsWithDefaults() *SearchStats`

NewSearchStatsWithDefaults instantiates a new SearchStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseSize

`func (o *SearchStats) GetDatabaseSize() int64`

GetDatabaseSize returns the DatabaseSize field if non-nil, zero value otherwise.

### GetDatabaseSizeOk

`func (o *SearchStats) GetDatabaseSizeOk() (*int64, bool)`

GetDatabaseSizeOk returns a tuple with the DatabaseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseSize

`func (o *SearchStats) SetDatabaseSize(v int64)`

SetDatabaseSize sets DatabaseSize field to given value.

### HasDatabaseSize

`func (o *SearchStats) HasDatabaseSize() bool`

HasDatabaseSize returns a boolean if a field has been set.

### GetUsedDatabaseSize

`func (o *SearchStats) GetUsedDatabaseSize() int64`

GetUsedDatabaseSize returns the UsedDatabaseSize field if non-nil, zero value otherwise.

### GetUsedDatabaseSizeOk

`func (o *SearchStats) GetUsedDatabaseSizeOk() (*int64, bool)`

GetUsedDatabaseSizeOk returns a tuple with the UsedDatabaseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedDatabaseSize

`func (o *SearchStats) SetUsedDatabaseSize(v int64)`

SetUsedDatabaseSize sets UsedDatabaseSize field to given value.

### HasUsedDatabaseSize

`func (o *SearchStats) HasUsedDatabaseSize() bool`

HasUsedDatabaseSize returns a boolean if a field has been set.

### GetLastUpdate

`func (o *SearchStats) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *SearchStats) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *SearchStats) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *SearchStats) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetIndexes

`func (o *SearchStats) GetIndexes() map[string]SearchIndexStats`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *SearchStats) GetIndexesOk() (*map[string]SearchIndexStats, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *SearchStats) SetIndexes(v map[string]SearchIndexStats)`

SetIndexes sets Indexes field to given value.

### HasIndexes

`func (o *SearchStats) HasIndexes() bool`

HasIndexes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


