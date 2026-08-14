# O11yNodeRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to **interface{}** |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**NodeCPU** | Pointer to **float32** |  | [optional] 
**NodeCPUAllocatable** | Pointer to **float32** |  | [optional] 
**NodeCountsByReadiness** | Pointer to [**O11yNodeCountsByReadiness**](O11yNodeCountsByReadiness.md) |  | [optional] 
**NodeMemory** | Pointer to **float32** |  | [optional] 
**NodeMemoryAllocatable** | Pointer to **float32** |  | [optional] 
**NodeName** | Pointer to **string** |  | [optional] 
**PodCountsByPhase** | Pointer to [**O11yPodCountsByPhase**](O11yPodCountsByPhase.md) |  | [optional] 

## Methods

### NewO11yNodeRecord

`func NewO11yNodeRecord() *O11yNodeRecord`

NewO11yNodeRecord instantiates a new O11yNodeRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yNodeRecordWithDefaults

`func NewO11yNodeRecordWithDefaults() *O11yNodeRecord`

NewO11yNodeRecordWithDefaults instantiates a new O11yNodeRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *O11yNodeRecord) GetCondition() interface{}`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *O11yNodeRecord) GetConditionOk() (*interface{}, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *O11yNodeRecord) SetCondition(v interface{})`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *O11yNodeRecord) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *O11yNodeRecord) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *O11yNodeRecord) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetMeta

`func (o *O11yNodeRecord) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yNodeRecord) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yNodeRecord) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yNodeRecord) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetNodeCPU

`func (o *O11yNodeRecord) GetNodeCPU() float32`

GetNodeCPU returns the NodeCPU field if non-nil, zero value otherwise.

### GetNodeCPUOk

`func (o *O11yNodeRecord) GetNodeCPUOk() (*float32, bool)`

GetNodeCPUOk returns a tuple with the NodeCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCPU

`func (o *O11yNodeRecord) SetNodeCPU(v float32)`

SetNodeCPU sets NodeCPU field to given value.

### HasNodeCPU

`func (o *O11yNodeRecord) HasNodeCPU() bool`

HasNodeCPU returns a boolean if a field has been set.

### GetNodeCPUAllocatable

`func (o *O11yNodeRecord) GetNodeCPUAllocatable() float32`

GetNodeCPUAllocatable returns the NodeCPUAllocatable field if non-nil, zero value otherwise.

### GetNodeCPUAllocatableOk

`func (o *O11yNodeRecord) GetNodeCPUAllocatableOk() (*float32, bool)`

GetNodeCPUAllocatableOk returns a tuple with the NodeCPUAllocatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCPUAllocatable

`func (o *O11yNodeRecord) SetNodeCPUAllocatable(v float32)`

SetNodeCPUAllocatable sets NodeCPUAllocatable field to given value.

### HasNodeCPUAllocatable

`func (o *O11yNodeRecord) HasNodeCPUAllocatable() bool`

HasNodeCPUAllocatable returns a boolean if a field has been set.

### GetNodeCountsByReadiness

`func (o *O11yNodeRecord) GetNodeCountsByReadiness() O11yNodeCountsByReadiness`

GetNodeCountsByReadiness returns the NodeCountsByReadiness field if non-nil, zero value otherwise.

### GetNodeCountsByReadinessOk

`func (o *O11yNodeRecord) GetNodeCountsByReadinessOk() (*O11yNodeCountsByReadiness, bool)`

GetNodeCountsByReadinessOk returns a tuple with the NodeCountsByReadiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCountsByReadiness

`func (o *O11yNodeRecord) SetNodeCountsByReadiness(v O11yNodeCountsByReadiness)`

SetNodeCountsByReadiness sets NodeCountsByReadiness field to given value.

### HasNodeCountsByReadiness

`func (o *O11yNodeRecord) HasNodeCountsByReadiness() bool`

HasNodeCountsByReadiness returns a boolean if a field has been set.

### GetNodeMemory

`func (o *O11yNodeRecord) GetNodeMemory() float32`

GetNodeMemory returns the NodeMemory field if non-nil, zero value otherwise.

### GetNodeMemoryOk

`func (o *O11yNodeRecord) GetNodeMemoryOk() (*float32, bool)`

GetNodeMemoryOk returns a tuple with the NodeMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeMemory

`func (o *O11yNodeRecord) SetNodeMemory(v float32)`

SetNodeMemory sets NodeMemory field to given value.

### HasNodeMemory

`func (o *O11yNodeRecord) HasNodeMemory() bool`

HasNodeMemory returns a boolean if a field has been set.

### GetNodeMemoryAllocatable

`func (o *O11yNodeRecord) GetNodeMemoryAllocatable() float32`

GetNodeMemoryAllocatable returns the NodeMemoryAllocatable field if non-nil, zero value otherwise.

### GetNodeMemoryAllocatableOk

`func (o *O11yNodeRecord) GetNodeMemoryAllocatableOk() (*float32, bool)`

GetNodeMemoryAllocatableOk returns a tuple with the NodeMemoryAllocatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeMemoryAllocatable

`func (o *O11yNodeRecord) SetNodeMemoryAllocatable(v float32)`

SetNodeMemoryAllocatable sets NodeMemoryAllocatable field to given value.

### HasNodeMemoryAllocatable

`func (o *O11yNodeRecord) HasNodeMemoryAllocatable() bool`

HasNodeMemoryAllocatable returns a boolean if a field has been set.

### GetNodeName

`func (o *O11yNodeRecord) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *O11yNodeRecord) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *O11yNodeRecord) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *O11yNodeRecord) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetPodCountsByPhase

`func (o *O11yNodeRecord) GetPodCountsByPhase() O11yPodCountsByPhase`

GetPodCountsByPhase returns the PodCountsByPhase field if non-nil, zero value otherwise.

### GetPodCountsByPhaseOk

`func (o *O11yNodeRecord) GetPodCountsByPhaseOk() (*O11yPodCountsByPhase, bool)`

GetPodCountsByPhaseOk returns a tuple with the PodCountsByPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCountsByPhase

`func (o *O11yNodeRecord) SetPodCountsByPhase(v O11yPodCountsByPhase)`

SetPodCountsByPhase sets PodCountsByPhase field to given value.

### HasPodCountsByPhase

`func (o *O11yNodeRecord) HasPodCountsByPhase() bool`

HasPodCountsByPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


