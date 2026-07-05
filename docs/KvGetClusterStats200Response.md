# KvGetClusterStats200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectedClients** | Pointer to **int32** |  | [optional] 
**UsedMemoryMb** | Pointer to **float32** |  | [optional] 
**UsedMemoryPeakMb** | Pointer to **float32** |  | [optional] 
**TotalCommandsProcessed** | Pointer to **int64** |  | [optional] 
**KeyspaceHits** | Pointer to **int64** |  | [optional] 
**KeyspaceMisses** | Pointer to **int64** |  | [optional] 
**HitRate** | Pointer to **float32** |  | [optional] 
**OpsPerSec** | Pointer to **int32** |  | [optional] 
**TotalKeys** | Pointer to **int32** |  | [optional] 
**ExpiredKeys** | Pointer to **int32** |  | [optional] 
**EvictedKeys** | Pointer to **int32** |  | [optional] 

## Methods

### NewKvGetClusterStats200Response

`func NewKvGetClusterStats200Response() *KvGetClusterStats200Response`

NewKvGetClusterStats200Response instantiates a new KvGetClusterStats200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvGetClusterStats200ResponseWithDefaults

`func NewKvGetClusterStats200ResponseWithDefaults() *KvGetClusterStats200Response`

NewKvGetClusterStats200ResponseWithDefaults instantiates a new KvGetClusterStats200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectedClients

`func (o *KvGetClusterStats200Response) GetConnectedClients() int32`

GetConnectedClients returns the ConnectedClients field if non-nil, zero value otherwise.

### GetConnectedClientsOk

`func (o *KvGetClusterStats200Response) GetConnectedClientsOk() (*int32, bool)`

GetConnectedClientsOk returns a tuple with the ConnectedClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedClients

`func (o *KvGetClusterStats200Response) SetConnectedClients(v int32)`

SetConnectedClients sets ConnectedClients field to given value.

### HasConnectedClients

`func (o *KvGetClusterStats200Response) HasConnectedClients() bool`

HasConnectedClients returns a boolean if a field has been set.

### GetUsedMemoryMb

`func (o *KvGetClusterStats200Response) GetUsedMemoryMb() float32`

GetUsedMemoryMb returns the UsedMemoryMb field if non-nil, zero value otherwise.

### GetUsedMemoryMbOk

`func (o *KvGetClusterStats200Response) GetUsedMemoryMbOk() (*float32, bool)`

GetUsedMemoryMbOk returns a tuple with the UsedMemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemoryMb

`func (o *KvGetClusterStats200Response) SetUsedMemoryMb(v float32)`

SetUsedMemoryMb sets UsedMemoryMb field to given value.

### HasUsedMemoryMb

`func (o *KvGetClusterStats200Response) HasUsedMemoryMb() bool`

HasUsedMemoryMb returns a boolean if a field has been set.

### GetUsedMemoryPeakMb

`func (o *KvGetClusterStats200Response) GetUsedMemoryPeakMb() float32`

GetUsedMemoryPeakMb returns the UsedMemoryPeakMb field if non-nil, zero value otherwise.

### GetUsedMemoryPeakMbOk

`func (o *KvGetClusterStats200Response) GetUsedMemoryPeakMbOk() (*float32, bool)`

GetUsedMemoryPeakMbOk returns a tuple with the UsedMemoryPeakMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemoryPeakMb

`func (o *KvGetClusterStats200Response) SetUsedMemoryPeakMb(v float32)`

SetUsedMemoryPeakMb sets UsedMemoryPeakMb field to given value.

### HasUsedMemoryPeakMb

`func (o *KvGetClusterStats200Response) HasUsedMemoryPeakMb() bool`

HasUsedMemoryPeakMb returns a boolean if a field has been set.

### GetTotalCommandsProcessed

`func (o *KvGetClusterStats200Response) GetTotalCommandsProcessed() int64`

GetTotalCommandsProcessed returns the TotalCommandsProcessed field if non-nil, zero value otherwise.

### GetTotalCommandsProcessedOk

`func (o *KvGetClusterStats200Response) GetTotalCommandsProcessedOk() (*int64, bool)`

GetTotalCommandsProcessedOk returns a tuple with the TotalCommandsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCommandsProcessed

`func (o *KvGetClusterStats200Response) SetTotalCommandsProcessed(v int64)`

SetTotalCommandsProcessed sets TotalCommandsProcessed field to given value.

### HasTotalCommandsProcessed

`func (o *KvGetClusterStats200Response) HasTotalCommandsProcessed() bool`

HasTotalCommandsProcessed returns a boolean if a field has been set.

### GetKeyspaceHits

`func (o *KvGetClusterStats200Response) GetKeyspaceHits() int64`

GetKeyspaceHits returns the KeyspaceHits field if non-nil, zero value otherwise.

### GetKeyspaceHitsOk

`func (o *KvGetClusterStats200Response) GetKeyspaceHitsOk() (*int64, bool)`

GetKeyspaceHitsOk returns a tuple with the KeyspaceHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspaceHits

`func (o *KvGetClusterStats200Response) SetKeyspaceHits(v int64)`

SetKeyspaceHits sets KeyspaceHits field to given value.

### HasKeyspaceHits

`func (o *KvGetClusterStats200Response) HasKeyspaceHits() bool`

HasKeyspaceHits returns a boolean if a field has been set.

### GetKeyspaceMisses

`func (o *KvGetClusterStats200Response) GetKeyspaceMisses() int64`

GetKeyspaceMisses returns the KeyspaceMisses field if non-nil, zero value otherwise.

### GetKeyspaceMissesOk

`func (o *KvGetClusterStats200Response) GetKeyspaceMissesOk() (*int64, bool)`

GetKeyspaceMissesOk returns a tuple with the KeyspaceMisses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyspaceMisses

`func (o *KvGetClusterStats200Response) SetKeyspaceMisses(v int64)`

SetKeyspaceMisses sets KeyspaceMisses field to given value.

### HasKeyspaceMisses

`func (o *KvGetClusterStats200Response) HasKeyspaceMisses() bool`

HasKeyspaceMisses returns a boolean if a field has been set.

### GetHitRate

`func (o *KvGetClusterStats200Response) GetHitRate() float32`

GetHitRate returns the HitRate field if non-nil, zero value otherwise.

### GetHitRateOk

`func (o *KvGetClusterStats200Response) GetHitRateOk() (*float32, bool)`

GetHitRateOk returns a tuple with the HitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitRate

`func (o *KvGetClusterStats200Response) SetHitRate(v float32)`

SetHitRate sets HitRate field to given value.

### HasHitRate

`func (o *KvGetClusterStats200Response) HasHitRate() bool`

HasHitRate returns a boolean if a field has been set.

### GetOpsPerSec

`func (o *KvGetClusterStats200Response) GetOpsPerSec() int32`

GetOpsPerSec returns the OpsPerSec field if non-nil, zero value otherwise.

### GetOpsPerSecOk

`func (o *KvGetClusterStats200Response) GetOpsPerSecOk() (*int32, bool)`

GetOpsPerSecOk returns a tuple with the OpsPerSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpsPerSec

`func (o *KvGetClusterStats200Response) SetOpsPerSec(v int32)`

SetOpsPerSec sets OpsPerSec field to given value.

### HasOpsPerSec

`func (o *KvGetClusterStats200Response) HasOpsPerSec() bool`

HasOpsPerSec returns a boolean if a field has been set.

### GetTotalKeys

`func (o *KvGetClusterStats200Response) GetTotalKeys() int32`

GetTotalKeys returns the TotalKeys field if non-nil, zero value otherwise.

### GetTotalKeysOk

`func (o *KvGetClusterStats200Response) GetTotalKeysOk() (*int32, bool)`

GetTotalKeysOk returns a tuple with the TotalKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalKeys

`func (o *KvGetClusterStats200Response) SetTotalKeys(v int32)`

SetTotalKeys sets TotalKeys field to given value.

### HasTotalKeys

`func (o *KvGetClusterStats200Response) HasTotalKeys() bool`

HasTotalKeys returns a boolean if a field has been set.

### GetExpiredKeys

`func (o *KvGetClusterStats200Response) GetExpiredKeys() int32`

GetExpiredKeys returns the ExpiredKeys field if non-nil, zero value otherwise.

### GetExpiredKeysOk

`func (o *KvGetClusterStats200Response) GetExpiredKeysOk() (*int32, bool)`

GetExpiredKeysOk returns a tuple with the ExpiredKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredKeys

`func (o *KvGetClusterStats200Response) SetExpiredKeys(v int32)`

SetExpiredKeys sets ExpiredKeys field to given value.

### HasExpiredKeys

`func (o *KvGetClusterStats200Response) HasExpiredKeys() bool`

HasExpiredKeys returns a boolean if a field has been set.

### GetEvictedKeys

`func (o *KvGetClusterStats200Response) GetEvictedKeys() int32`

GetEvictedKeys returns the EvictedKeys field if non-nil, zero value otherwise.

### GetEvictedKeysOk

`func (o *KvGetClusterStats200Response) GetEvictedKeysOk() (*int32, bool)`

GetEvictedKeysOk returns a tuple with the EvictedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvictedKeys

`func (o *KvGetClusterStats200Response) SetEvictedKeys(v int32)`

SetEvictedKeys sets EvictedKeys field to given value.

### HasEvictedKeys

`func (o *KvGetClusterStats200Response) HasEvictedKeys() bool`

HasEvictedKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


