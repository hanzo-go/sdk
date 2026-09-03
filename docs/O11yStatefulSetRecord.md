# O11yStatefulSetRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPods** | Pointer to **int64** |  | [optional] 
**DesiredPods** | Pointer to **int64** |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**PodCountsByPhase** | Pointer to [**O11yPodCountsByPhase**](O11yPodCountsByPhase.md) |  | [optional] 
**StatefulSetCPU** | Pointer to **float64** |  | [optional] 
**StatefulSetCPULimit** | Pointer to **float64** |  | [optional] 
**StatefulSetCPURequest** | Pointer to **float64** |  | [optional] 
**StatefulSetMemory** | Pointer to **float64** |  | [optional] 
**StatefulSetMemoryLimit** | Pointer to **float64** |  | [optional] 
**StatefulSetMemoryRequest** | Pointer to **float64** |  | [optional] 
**StatefulSetName** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yStatefulSetRecord

`func NewO11yStatefulSetRecord() *O11yStatefulSetRecord`

NewO11yStatefulSetRecord instantiates a new O11yStatefulSetRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yStatefulSetRecordWithDefaults

`func NewO11yStatefulSetRecordWithDefaults() *O11yStatefulSetRecord`

NewO11yStatefulSetRecordWithDefaults instantiates a new O11yStatefulSetRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPods

`func (o *O11yStatefulSetRecord) GetCurrentPods() int64`

GetCurrentPods returns the CurrentPods field if non-nil, zero value otherwise.

### GetCurrentPodsOk

`func (o *O11yStatefulSetRecord) GetCurrentPodsOk() (*int64, bool)`

GetCurrentPodsOk returns a tuple with the CurrentPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPods

`func (o *O11yStatefulSetRecord) SetCurrentPods(v int64)`

SetCurrentPods sets CurrentPods field to given value.

### HasCurrentPods

`func (o *O11yStatefulSetRecord) HasCurrentPods() bool`

HasCurrentPods returns a boolean if a field has been set.

### GetDesiredPods

`func (o *O11yStatefulSetRecord) GetDesiredPods() int64`

GetDesiredPods returns the DesiredPods field if non-nil, zero value otherwise.

### GetDesiredPodsOk

`func (o *O11yStatefulSetRecord) GetDesiredPodsOk() (*int64, bool)`

GetDesiredPodsOk returns a tuple with the DesiredPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredPods

`func (o *O11yStatefulSetRecord) SetDesiredPods(v int64)`

SetDesiredPods sets DesiredPods field to given value.

### HasDesiredPods

`func (o *O11yStatefulSetRecord) HasDesiredPods() bool`

HasDesiredPods returns a boolean if a field has been set.

### GetMeta

`func (o *O11yStatefulSetRecord) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yStatefulSetRecord) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yStatefulSetRecord) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yStatefulSetRecord) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPodCountsByPhase

`func (o *O11yStatefulSetRecord) GetPodCountsByPhase() O11yPodCountsByPhase`

GetPodCountsByPhase returns the PodCountsByPhase field if non-nil, zero value otherwise.

### GetPodCountsByPhaseOk

`func (o *O11yStatefulSetRecord) GetPodCountsByPhaseOk() (*O11yPodCountsByPhase, bool)`

GetPodCountsByPhaseOk returns a tuple with the PodCountsByPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCountsByPhase

`func (o *O11yStatefulSetRecord) SetPodCountsByPhase(v O11yPodCountsByPhase)`

SetPodCountsByPhase sets PodCountsByPhase field to given value.

### HasPodCountsByPhase

`func (o *O11yStatefulSetRecord) HasPodCountsByPhase() bool`

HasPodCountsByPhase returns a boolean if a field has been set.

### GetStatefulSetCPU

`func (o *O11yStatefulSetRecord) GetStatefulSetCPU() float64`

GetStatefulSetCPU returns the StatefulSetCPU field if non-nil, zero value otherwise.

### GetStatefulSetCPUOk

`func (o *O11yStatefulSetRecord) GetStatefulSetCPUOk() (*float64, bool)`

GetStatefulSetCPUOk returns a tuple with the StatefulSetCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatefulSetCPU

`func (o *O11yStatefulSetRecord) SetStatefulSetCPU(v float64)`

SetStatefulSetCPU sets StatefulSetCPU field to given value.

### HasStatefulSetCPU

`func (o *O11yStatefulSetRecord) HasStatefulSetCPU() bool`

HasStatefulSetCPU returns a boolean if a field has been set.

### GetStatefulSetCPULimit

`func (o *O11yStatefulSetRecord) GetStatefulSetCPULimit() float64`

GetStatefulSetCPULimit returns the StatefulSetCPULimit field if non-nil, zero value otherwise.

### GetStatefulSetCPULimitOk

`func (o *O11yStatefulSetRecord) GetStatefulSetCPULimitOk() (*float64, bool)`

GetStatefulSetCPULimitOk returns a tuple with the StatefulSetCPULimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatefulSetCPULimit

`func (o *O11yStatefulSetRecord) SetStatefulSetCPULimit(v float64)`

SetStatefulSetCPULimit sets StatefulSetCPULimit field to given value.

### HasStatefulSetCPULimit

`func (o *O11yStatefulSetRecord) HasStatefulSetCPULimit() bool`

HasStatefulSetCPULimit returns a boolean if a field has been set.

### GetStatefulSetCPURequest

`func (o *O11yStatefulSetRecord) GetStatefulSetCPURequest() float64`

GetStatefulSetCPURequest returns the StatefulSetCPURequest field if non-nil, zero value otherwise.

### GetStatefulSetCPURequestOk

`func (o *O11yStatefulSetRecord) GetStatefulSetCPURequestOk() (*float64, bool)`

GetStatefulSetCPURequestOk returns a tuple with the StatefulSetCPURequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatefulSetCPURequest

`func (o *O11yStatefulSetRecord) SetStatefulSetCPURequest(v float64)`

SetStatefulSetCPURequest sets StatefulSetCPURequest field to given value.

### HasStatefulSetCPURequest

`func (o *O11yStatefulSetRecord) HasStatefulSetCPURequest() bool`

HasStatefulSetCPURequest returns a boolean if a field has been set.

### GetStatefulSetMemory

`func (o *O11yStatefulSetRecord) GetStatefulSetMemory() float64`

GetStatefulSetMemory returns the StatefulSetMemory field if non-nil, zero value otherwise.

### GetStatefulSetMemoryOk

`func (o *O11yStatefulSetRecord) GetStatefulSetMemoryOk() (*float64, bool)`

GetStatefulSetMemoryOk returns a tuple with the StatefulSetMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatefulSetMemory

`func (o *O11yStatefulSetRecord) SetStatefulSetMemory(v float64)`

SetStatefulSetMemory sets StatefulSetMemory field to given value.

### HasStatefulSetMemory

`func (o *O11yStatefulSetRecord) HasStatefulSetMemory() bool`

HasStatefulSetMemory returns a boolean if a field has been set.

### GetStatefulSetMemoryLimit

`func (o *O11yStatefulSetRecord) GetStatefulSetMemoryLimit() float64`

GetStatefulSetMemoryLimit returns the StatefulSetMemoryLimit field if non-nil, zero value otherwise.

### GetStatefulSetMemoryLimitOk

`func (o *O11yStatefulSetRecord) GetStatefulSetMemoryLimitOk() (*float64, bool)`

GetStatefulSetMemoryLimitOk returns a tuple with the StatefulSetMemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatefulSetMemoryLimit

`func (o *O11yStatefulSetRecord) SetStatefulSetMemoryLimit(v float64)`

SetStatefulSetMemoryLimit sets StatefulSetMemoryLimit field to given value.

### HasStatefulSetMemoryLimit

`func (o *O11yStatefulSetRecord) HasStatefulSetMemoryLimit() bool`

HasStatefulSetMemoryLimit returns a boolean if a field has been set.

### GetStatefulSetMemoryRequest

`func (o *O11yStatefulSetRecord) GetStatefulSetMemoryRequest() float64`

GetStatefulSetMemoryRequest returns the StatefulSetMemoryRequest field if non-nil, zero value otherwise.

### GetStatefulSetMemoryRequestOk

`func (o *O11yStatefulSetRecord) GetStatefulSetMemoryRequestOk() (*float64, bool)`

GetStatefulSetMemoryRequestOk returns a tuple with the StatefulSetMemoryRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatefulSetMemoryRequest

`func (o *O11yStatefulSetRecord) SetStatefulSetMemoryRequest(v float64)`

SetStatefulSetMemoryRequest sets StatefulSetMemoryRequest field to given value.

### HasStatefulSetMemoryRequest

`func (o *O11yStatefulSetRecord) HasStatefulSetMemoryRequest() bool`

HasStatefulSetMemoryRequest returns a boolean if a field has been set.

### GetStatefulSetName

`func (o *O11yStatefulSetRecord) GetStatefulSetName() string`

GetStatefulSetName returns the StatefulSetName field if non-nil, zero value otherwise.

### GetStatefulSetNameOk

`func (o *O11yStatefulSetRecord) GetStatefulSetNameOk() (*string, bool)`

GetStatefulSetNameOk returns a tuple with the StatefulSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatefulSetName

`func (o *O11yStatefulSetRecord) SetStatefulSetName(v string)`

SetStatefulSetName sets StatefulSetName field to given value.

### HasStatefulSetName

`func (o *O11yStatefulSetRecord) HasStatefulSetName() bool`

HasStatefulSetName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


