# EngineCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int32** |  | [optional] 
**TotalGpus** | Pointer to **int32** |  | [optional] 
**AvailableGpus** | Pointer to **int32** |  | [optional] 
**GpuTypes** | Pointer to **[]string** |  | [optional] 
**TotalMemoryGb** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEngineCluster

`func NewEngineCluster() *EngineCluster`

NewEngineCluster instantiates a new EngineCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineClusterWithDefaults

`func NewEngineClusterWithDefaults() *EngineCluster`

NewEngineClusterWithDefaults instantiates a new EngineCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EngineCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EngineCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EngineCluster) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EngineCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EngineCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EngineCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EngineCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EngineCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *EngineCluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EngineCluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EngineCluster) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EngineCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProvider

`func (o *EngineCluster) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EngineCluster) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EngineCluster) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EngineCluster) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRegion

`func (o *EngineCluster) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EngineCluster) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EngineCluster) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *EngineCluster) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetNodeCount

`func (o *EngineCluster) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *EngineCluster) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *EngineCluster) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *EngineCluster) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetTotalGpus

`func (o *EngineCluster) GetTotalGpus() int32`

GetTotalGpus returns the TotalGpus field if non-nil, zero value otherwise.

### GetTotalGpusOk

`func (o *EngineCluster) GetTotalGpusOk() (*int32, bool)`

GetTotalGpusOk returns a tuple with the TotalGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGpus

`func (o *EngineCluster) SetTotalGpus(v int32)`

SetTotalGpus sets TotalGpus field to given value.

### HasTotalGpus

`func (o *EngineCluster) HasTotalGpus() bool`

HasTotalGpus returns a boolean if a field has been set.

### GetAvailableGpus

`func (o *EngineCluster) GetAvailableGpus() int32`

GetAvailableGpus returns the AvailableGpus field if non-nil, zero value otherwise.

### GetAvailableGpusOk

`func (o *EngineCluster) GetAvailableGpusOk() (*int32, bool)`

GetAvailableGpusOk returns a tuple with the AvailableGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableGpus

`func (o *EngineCluster) SetAvailableGpus(v int32)`

SetAvailableGpus sets AvailableGpus field to given value.

### HasAvailableGpus

`func (o *EngineCluster) HasAvailableGpus() bool`

HasAvailableGpus returns a boolean if a field has been set.

### GetGpuTypes

`func (o *EngineCluster) GetGpuTypes() []string`

GetGpuTypes returns the GpuTypes field if non-nil, zero value otherwise.

### GetGpuTypesOk

`func (o *EngineCluster) GetGpuTypesOk() (*[]string, bool)`

GetGpuTypesOk returns a tuple with the GpuTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuTypes

`func (o *EngineCluster) SetGpuTypes(v []string)`

SetGpuTypes sets GpuTypes field to given value.

### HasGpuTypes

`func (o *EngineCluster) HasGpuTypes() bool`

HasGpuTypes returns a boolean if a field has been set.

### GetTotalMemoryGb

`func (o *EngineCluster) GetTotalMemoryGb() int32`

GetTotalMemoryGb returns the TotalMemoryGb field if non-nil, zero value otherwise.

### GetTotalMemoryGbOk

`func (o *EngineCluster) GetTotalMemoryGbOk() (*int32, bool)`

GetTotalMemoryGbOk returns a tuple with the TotalMemoryGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemoryGb

`func (o *EngineCluster) SetTotalMemoryGb(v int32)`

SetTotalMemoryGb sets TotalMemoryGb field to given value.

### HasTotalMemoryGb

`func (o *EngineCluster) HasTotalMemoryGb() bool`

HasTotalMemoryGb returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EngineCluster) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EngineCluster) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EngineCluster) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EngineCluster) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EngineCluster) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EngineCluster) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EngineCluster) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EngineCluster) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


