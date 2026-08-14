# VectorStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCollections** | Pointer to **int32** | TotalCollections is how many collections the store holds. | [optional] 
**TotalStorageBytes** | Pointer to **int32** | TotalStorageBytes is the sum of every collection&#39;s on-disk size. | [optional] 
**TotalVectors** | Pointer to **int32** | TotalVectors is the sum of every collection&#39;s point count. | [optional] 

## Methods

### NewVectorStats

`func NewVectorStats() *VectorStats`

NewVectorStats instantiates a new VectorStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorStatsWithDefaults

`func NewVectorStatsWithDefaults() *VectorStats`

NewVectorStatsWithDefaults instantiates a new VectorStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCollections

`func (o *VectorStats) GetTotalCollections() int32`

GetTotalCollections returns the TotalCollections field if non-nil, zero value otherwise.

### GetTotalCollectionsOk

`func (o *VectorStats) GetTotalCollectionsOk() (*int32, bool)`

GetTotalCollectionsOk returns a tuple with the TotalCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCollections

`func (o *VectorStats) SetTotalCollections(v int32)`

SetTotalCollections sets TotalCollections field to given value.

### HasTotalCollections

`func (o *VectorStats) HasTotalCollections() bool`

HasTotalCollections returns a boolean if a field has been set.

### GetTotalStorageBytes

`func (o *VectorStats) GetTotalStorageBytes() int32`

GetTotalStorageBytes returns the TotalStorageBytes field if non-nil, zero value otherwise.

### GetTotalStorageBytesOk

`func (o *VectorStats) GetTotalStorageBytesOk() (*int32, bool)`

GetTotalStorageBytesOk returns a tuple with the TotalStorageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStorageBytes

`func (o *VectorStats) SetTotalStorageBytes(v int32)`

SetTotalStorageBytes sets TotalStorageBytes field to given value.

### HasTotalStorageBytes

`func (o *VectorStats) HasTotalStorageBytes() bool`

HasTotalStorageBytes returns a boolean if a field has been set.

### GetTotalVectors

`func (o *VectorStats) GetTotalVectors() int32`

GetTotalVectors returns the TotalVectors field if non-nil, zero value otherwise.

### GetTotalVectorsOk

`func (o *VectorStats) GetTotalVectorsOk() (*int32, bool)`

GetTotalVectorsOk returns a tuple with the TotalVectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVectors

`func (o *VectorStats) SetTotalVectors(v int32)`

SetTotalVectors sets TotalVectors field to given value.

### HasTotalVectors

`func (o *VectorStats) HasTotalVectors() bool`

HasTotalVectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


