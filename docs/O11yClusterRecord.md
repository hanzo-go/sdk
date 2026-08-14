# O11yClusterRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterCPU** | Pointer to **float32** |  | [optional] 
**ClusterCPUAllocatable** | Pointer to **float32** |  | [optional] 
**ClusterMemory** | Pointer to **float32** |  | [optional] 
**ClusterMemoryAllocatable** | Pointer to **float32** |  | [optional] 
**ClusterName** | Pointer to **string** | TODO(nikhilmantri0902): once the underlying attr key is migrated to k8s.cluster.uid (see ClusterNameAttrKey), surface ClusterUID alongside (or replace) ClusterName. | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**NodeCountsByReadiness** | Pointer to [**O11yNodeCountsByReadiness**](O11yNodeCountsByReadiness.md) |  | [optional] 
**PodCountsByPhase** | Pointer to [**O11yPodCountsByPhase**](O11yPodCountsByPhase.md) |  | [optional] 

## Methods

### NewO11yClusterRecord

`func NewO11yClusterRecord() *O11yClusterRecord`

NewO11yClusterRecord instantiates a new O11yClusterRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yClusterRecordWithDefaults

`func NewO11yClusterRecordWithDefaults() *O11yClusterRecord`

NewO11yClusterRecordWithDefaults instantiates a new O11yClusterRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterCPU

`func (o *O11yClusterRecord) GetClusterCPU() float32`

GetClusterCPU returns the ClusterCPU field if non-nil, zero value otherwise.

### GetClusterCPUOk

`func (o *O11yClusterRecord) GetClusterCPUOk() (*float32, bool)`

GetClusterCPUOk returns a tuple with the ClusterCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCPU

`func (o *O11yClusterRecord) SetClusterCPU(v float32)`

SetClusterCPU sets ClusterCPU field to given value.

### HasClusterCPU

`func (o *O11yClusterRecord) HasClusterCPU() bool`

HasClusterCPU returns a boolean if a field has been set.

### GetClusterCPUAllocatable

`func (o *O11yClusterRecord) GetClusterCPUAllocatable() float32`

GetClusterCPUAllocatable returns the ClusterCPUAllocatable field if non-nil, zero value otherwise.

### GetClusterCPUAllocatableOk

`func (o *O11yClusterRecord) GetClusterCPUAllocatableOk() (*float32, bool)`

GetClusterCPUAllocatableOk returns a tuple with the ClusterCPUAllocatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCPUAllocatable

`func (o *O11yClusterRecord) SetClusterCPUAllocatable(v float32)`

SetClusterCPUAllocatable sets ClusterCPUAllocatable field to given value.

### HasClusterCPUAllocatable

`func (o *O11yClusterRecord) HasClusterCPUAllocatable() bool`

HasClusterCPUAllocatable returns a boolean if a field has been set.

### GetClusterMemory

`func (o *O11yClusterRecord) GetClusterMemory() float32`

GetClusterMemory returns the ClusterMemory field if non-nil, zero value otherwise.

### GetClusterMemoryOk

`func (o *O11yClusterRecord) GetClusterMemoryOk() (*float32, bool)`

GetClusterMemoryOk returns a tuple with the ClusterMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMemory

`func (o *O11yClusterRecord) SetClusterMemory(v float32)`

SetClusterMemory sets ClusterMemory field to given value.

### HasClusterMemory

`func (o *O11yClusterRecord) HasClusterMemory() bool`

HasClusterMemory returns a boolean if a field has been set.

### GetClusterMemoryAllocatable

`func (o *O11yClusterRecord) GetClusterMemoryAllocatable() float32`

GetClusterMemoryAllocatable returns the ClusterMemoryAllocatable field if non-nil, zero value otherwise.

### GetClusterMemoryAllocatableOk

`func (o *O11yClusterRecord) GetClusterMemoryAllocatableOk() (*float32, bool)`

GetClusterMemoryAllocatableOk returns a tuple with the ClusterMemoryAllocatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMemoryAllocatable

`func (o *O11yClusterRecord) SetClusterMemoryAllocatable(v float32)`

SetClusterMemoryAllocatable sets ClusterMemoryAllocatable field to given value.

### HasClusterMemoryAllocatable

`func (o *O11yClusterRecord) HasClusterMemoryAllocatable() bool`

HasClusterMemoryAllocatable returns a boolean if a field has been set.

### GetClusterName

`func (o *O11yClusterRecord) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *O11yClusterRecord) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *O11yClusterRecord) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *O11yClusterRecord) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetMeta

`func (o *O11yClusterRecord) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yClusterRecord) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yClusterRecord) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yClusterRecord) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetNodeCountsByReadiness

`func (o *O11yClusterRecord) GetNodeCountsByReadiness() O11yNodeCountsByReadiness`

GetNodeCountsByReadiness returns the NodeCountsByReadiness field if non-nil, zero value otherwise.

### GetNodeCountsByReadinessOk

`func (o *O11yClusterRecord) GetNodeCountsByReadinessOk() (*O11yNodeCountsByReadiness, bool)`

GetNodeCountsByReadinessOk returns a tuple with the NodeCountsByReadiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCountsByReadiness

`func (o *O11yClusterRecord) SetNodeCountsByReadiness(v O11yNodeCountsByReadiness)`

SetNodeCountsByReadiness sets NodeCountsByReadiness field to given value.

### HasNodeCountsByReadiness

`func (o *O11yClusterRecord) HasNodeCountsByReadiness() bool`

HasNodeCountsByReadiness returns a boolean if a field has been set.

### GetPodCountsByPhase

`func (o *O11yClusterRecord) GetPodCountsByPhase() O11yPodCountsByPhase`

GetPodCountsByPhase returns the PodCountsByPhase field if non-nil, zero value otherwise.

### GetPodCountsByPhaseOk

`func (o *O11yClusterRecord) GetPodCountsByPhaseOk() (*O11yPodCountsByPhase, bool)`

GetPodCountsByPhaseOk returns a tuple with the PodCountsByPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCountsByPhase

`func (o *O11yClusterRecord) SetPodCountsByPhase(v O11yPodCountsByPhase)`

SetPodCountsByPhase sets PodCountsByPhase field to given value.

### HasPodCountsByPhase

`func (o *O11yClusterRecord) HasPodCountsByPhase() bool`

HasPodCountsByPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


