# O11yDaemonSetRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentNodes** | Pointer to **int32** |  | [optional] 
**DaemonSetCPU** | Pointer to **float32** |  | [optional] 
**DaemonSetCPULimit** | Pointer to **float32** |  | [optional] 
**DaemonSetCPURequest** | Pointer to **float32** |  | [optional] 
**DaemonSetMemory** | Pointer to **float32** |  | [optional] 
**DaemonSetMemoryLimit** | Pointer to **float32** |  | [optional] 
**DaemonSetMemoryRequest** | Pointer to **float32** |  | [optional] 
**DaemonSetName** | Pointer to **string** |  | [optional] 
**DesiredNodes** | Pointer to **int32** |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**PodCountsByPhase** | Pointer to [**O11yPodCountsByPhase**](O11yPodCountsByPhase.md) |  | [optional] 

## Methods

### NewO11yDaemonSetRecord

`func NewO11yDaemonSetRecord() *O11yDaemonSetRecord`

NewO11yDaemonSetRecord instantiates a new O11yDaemonSetRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yDaemonSetRecordWithDefaults

`func NewO11yDaemonSetRecordWithDefaults() *O11yDaemonSetRecord`

NewO11yDaemonSetRecordWithDefaults instantiates a new O11yDaemonSetRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentNodes

`func (o *O11yDaemonSetRecord) GetCurrentNodes() int32`

GetCurrentNodes returns the CurrentNodes field if non-nil, zero value otherwise.

### GetCurrentNodesOk

`func (o *O11yDaemonSetRecord) GetCurrentNodesOk() (*int32, bool)`

GetCurrentNodesOk returns a tuple with the CurrentNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNodes

`func (o *O11yDaemonSetRecord) SetCurrentNodes(v int32)`

SetCurrentNodes sets CurrentNodes field to given value.

### HasCurrentNodes

`func (o *O11yDaemonSetRecord) HasCurrentNodes() bool`

HasCurrentNodes returns a boolean if a field has been set.

### GetDaemonSetCPU

`func (o *O11yDaemonSetRecord) GetDaemonSetCPU() float32`

GetDaemonSetCPU returns the DaemonSetCPU field if non-nil, zero value otherwise.

### GetDaemonSetCPUOk

`func (o *O11yDaemonSetRecord) GetDaemonSetCPUOk() (*float32, bool)`

GetDaemonSetCPUOk returns a tuple with the DaemonSetCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonSetCPU

`func (o *O11yDaemonSetRecord) SetDaemonSetCPU(v float32)`

SetDaemonSetCPU sets DaemonSetCPU field to given value.

### HasDaemonSetCPU

`func (o *O11yDaemonSetRecord) HasDaemonSetCPU() bool`

HasDaemonSetCPU returns a boolean if a field has been set.

### GetDaemonSetCPULimit

`func (o *O11yDaemonSetRecord) GetDaemonSetCPULimit() float32`

GetDaemonSetCPULimit returns the DaemonSetCPULimit field if non-nil, zero value otherwise.

### GetDaemonSetCPULimitOk

`func (o *O11yDaemonSetRecord) GetDaemonSetCPULimitOk() (*float32, bool)`

GetDaemonSetCPULimitOk returns a tuple with the DaemonSetCPULimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonSetCPULimit

`func (o *O11yDaemonSetRecord) SetDaemonSetCPULimit(v float32)`

SetDaemonSetCPULimit sets DaemonSetCPULimit field to given value.

### HasDaemonSetCPULimit

`func (o *O11yDaemonSetRecord) HasDaemonSetCPULimit() bool`

HasDaemonSetCPULimit returns a boolean if a field has been set.

### GetDaemonSetCPURequest

`func (o *O11yDaemonSetRecord) GetDaemonSetCPURequest() float32`

GetDaemonSetCPURequest returns the DaemonSetCPURequest field if non-nil, zero value otherwise.

### GetDaemonSetCPURequestOk

`func (o *O11yDaemonSetRecord) GetDaemonSetCPURequestOk() (*float32, bool)`

GetDaemonSetCPURequestOk returns a tuple with the DaemonSetCPURequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonSetCPURequest

`func (o *O11yDaemonSetRecord) SetDaemonSetCPURequest(v float32)`

SetDaemonSetCPURequest sets DaemonSetCPURequest field to given value.

### HasDaemonSetCPURequest

`func (o *O11yDaemonSetRecord) HasDaemonSetCPURequest() bool`

HasDaemonSetCPURequest returns a boolean if a field has been set.

### GetDaemonSetMemory

`func (o *O11yDaemonSetRecord) GetDaemonSetMemory() float32`

GetDaemonSetMemory returns the DaemonSetMemory field if non-nil, zero value otherwise.

### GetDaemonSetMemoryOk

`func (o *O11yDaemonSetRecord) GetDaemonSetMemoryOk() (*float32, bool)`

GetDaemonSetMemoryOk returns a tuple with the DaemonSetMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonSetMemory

`func (o *O11yDaemonSetRecord) SetDaemonSetMemory(v float32)`

SetDaemonSetMemory sets DaemonSetMemory field to given value.

### HasDaemonSetMemory

`func (o *O11yDaemonSetRecord) HasDaemonSetMemory() bool`

HasDaemonSetMemory returns a boolean if a field has been set.

### GetDaemonSetMemoryLimit

`func (o *O11yDaemonSetRecord) GetDaemonSetMemoryLimit() float32`

GetDaemonSetMemoryLimit returns the DaemonSetMemoryLimit field if non-nil, zero value otherwise.

### GetDaemonSetMemoryLimitOk

`func (o *O11yDaemonSetRecord) GetDaemonSetMemoryLimitOk() (*float32, bool)`

GetDaemonSetMemoryLimitOk returns a tuple with the DaemonSetMemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonSetMemoryLimit

`func (o *O11yDaemonSetRecord) SetDaemonSetMemoryLimit(v float32)`

SetDaemonSetMemoryLimit sets DaemonSetMemoryLimit field to given value.

### HasDaemonSetMemoryLimit

`func (o *O11yDaemonSetRecord) HasDaemonSetMemoryLimit() bool`

HasDaemonSetMemoryLimit returns a boolean if a field has been set.

### GetDaemonSetMemoryRequest

`func (o *O11yDaemonSetRecord) GetDaemonSetMemoryRequest() float32`

GetDaemonSetMemoryRequest returns the DaemonSetMemoryRequest field if non-nil, zero value otherwise.

### GetDaemonSetMemoryRequestOk

`func (o *O11yDaemonSetRecord) GetDaemonSetMemoryRequestOk() (*float32, bool)`

GetDaemonSetMemoryRequestOk returns a tuple with the DaemonSetMemoryRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonSetMemoryRequest

`func (o *O11yDaemonSetRecord) SetDaemonSetMemoryRequest(v float32)`

SetDaemonSetMemoryRequest sets DaemonSetMemoryRequest field to given value.

### HasDaemonSetMemoryRequest

`func (o *O11yDaemonSetRecord) HasDaemonSetMemoryRequest() bool`

HasDaemonSetMemoryRequest returns a boolean if a field has been set.

### GetDaemonSetName

`func (o *O11yDaemonSetRecord) GetDaemonSetName() string`

GetDaemonSetName returns the DaemonSetName field if non-nil, zero value otherwise.

### GetDaemonSetNameOk

`func (o *O11yDaemonSetRecord) GetDaemonSetNameOk() (*string, bool)`

GetDaemonSetNameOk returns a tuple with the DaemonSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonSetName

`func (o *O11yDaemonSetRecord) SetDaemonSetName(v string)`

SetDaemonSetName sets DaemonSetName field to given value.

### HasDaemonSetName

`func (o *O11yDaemonSetRecord) HasDaemonSetName() bool`

HasDaemonSetName returns a boolean if a field has been set.

### GetDesiredNodes

`func (o *O11yDaemonSetRecord) GetDesiredNodes() int32`

GetDesiredNodes returns the DesiredNodes field if non-nil, zero value otherwise.

### GetDesiredNodesOk

`func (o *O11yDaemonSetRecord) GetDesiredNodesOk() (*int32, bool)`

GetDesiredNodesOk returns a tuple with the DesiredNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredNodes

`func (o *O11yDaemonSetRecord) SetDesiredNodes(v int32)`

SetDesiredNodes sets DesiredNodes field to given value.

### HasDesiredNodes

`func (o *O11yDaemonSetRecord) HasDesiredNodes() bool`

HasDesiredNodes returns a boolean if a field has been set.

### GetMeta

`func (o *O11yDaemonSetRecord) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yDaemonSetRecord) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yDaemonSetRecord) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yDaemonSetRecord) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPodCountsByPhase

`func (o *O11yDaemonSetRecord) GetPodCountsByPhase() O11yPodCountsByPhase`

GetPodCountsByPhase returns the PodCountsByPhase field if non-nil, zero value otherwise.

### GetPodCountsByPhaseOk

`func (o *O11yDaemonSetRecord) GetPodCountsByPhaseOk() (*O11yPodCountsByPhase, bool)`

GetPodCountsByPhaseOk returns a tuple with the PodCountsByPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCountsByPhase

`func (o *O11yDaemonSetRecord) SetPodCountsByPhase(v O11yPodCountsByPhase)`

SetPodCountsByPhase sets PodCountsByPhase field to given value.

### HasPodCountsByPhase

`func (o *O11yDaemonSetRecord) HasPodCountsByPhase() bool`

HasPodCountsByPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


