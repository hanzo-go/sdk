# CloudVectorStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCollections** | Pointer to **int32** | TotalCollections is how many collections the store holds. | [optional] 
**TotalStorageBytes** | Pointer to **int32** | TotalStorageBytes is the sum of every collection&#39;s on-disk size. | [optional] 
**TotalVectors** | Pointer to **int32** | TotalVectors is the sum of every collection&#39;s point count. | [optional] 

## Methods

### NewCloudVectorStats

`func NewCloudVectorStats() *CloudVectorStats`

NewCloudVectorStats instantiates a new CloudVectorStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVectorStatsWithDefaults

`func NewCloudVectorStatsWithDefaults() *CloudVectorStats`

NewCloudVectorStatsWithDefaults instantiates a new CloudVectorStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCollections

`func (o *CloudVectorStats) GetTotalCollections() int32`

GetTotalCollections returns the TotalCollections field if non-nil, zero value otherwise.

### GetTotalCollectionsOk

`func (o *CloudVectorStats) GetTotalCollectionsOk() (*int32, bool)`

GetTotalCollectionsOk returns a tuple with the TotalCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCollections

`func (o *CloudVectorStats) SetTotalCollections(v int32)`

SetTotalCollections sets TotalCollections field to given value.

### HasTotalCollections

`func (o *CloudVectorStats) HasTotalCollections() bool`

HasTotalCollections returns a boolean if a field has been set.

### GetTotalStorageBytes

`func (o *CloudVectorStats) GetTotalStorageBytes() int32`

GetTotalStorageBytes returns the TotalStorageBytes field if non-nil, zero value otherwise.

### GetTotalStorageBytesOk

`func (o *CloudVectorStats) GetTotalStorageBytesOk() (*int32, bool)`

GetTotalStorageBytesOk returns a tuple with the TotalStorageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStorageBytes

`func (o *CloudVectorStats) SetTotalStorageBytes(v int32)`

SetTotalStorageBytes sets TotalStorageBytes field to given value.

### HasTotalStorageBytes

`func (o *CloudVectorStats) HasTotalStorageBytes() bool`

HasTotalStorageBytes returns a boolean if a field has been set.

### GetTotalVectors

`func (o *CloudVectorStats) GetTotalVectors() int32`

GetTotalVectors returns the TotalVectors field if non-nil, zero value otherwise.

### GetTotalVectorsOk

`func (o *CloudVectorStats) GetTotalVectorsOk() (*int32, bool)`

GetTotalVectorsOk returns a tuple with the TotalVectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVectors

`func (o *CloudVectorStats) SetTotalVectors(v int32)`

SetTotalVectors sets TotalVectors field to given value.

### HasTotalVectors

`func (o *CloudVectorStats) HasTotalVectors() bool`

HasTotalVectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


