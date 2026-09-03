# O11yDeploymentRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailablePods** | Pointer to **int64** |  | [optional] 
**DeploymentCPU** | Pointer to **float64** |  | [optional] 
**DeploymentCPULimit** | Pointer to **float64** |  | [optional] 
**DeploymentCPURequest** | Pointer to **float64** |  | [optional] 
**DeploymentMemory** | Pointer to **float64** |  | [optional] 
**DeploymentMemoryLimit** | Pointer to **float64** |  | [optional] 
**DeploymentMemoryRequest** | Pointer to **float64** |  | [optional] 
**DeploymentName** | Pointer to **string** |  | [optional] 
**DesiredPods** | Pointer to **int64** |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**PodCountsByPhase** | Pointer to [**O11yPodCountsByPhase**](O11yPodCountsByPhase.md) |  | [optional] 

## Methods

### NewO11yDeploymentRecord

`func NewO11yDeploymentRecord() *O11yDeploymentRecord`

NewO11yDeploymentRecord instantiates a new O11yDeploymentRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yDeploymentRecordWithDefaults

`func NewO11yDeploymentRecordWithDefaults() *O11yDeploymentRecord`

NewO11yDeploymentRecordWithDefaults instantiates a new O11yDeploymentRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailablePods

`func (o *O11yDeploymentRecord) GetAvailablePods() int64`

GetAvailablePods returns the AvailablePods field if non-nil, zero value otherwise.

### GetAvailablePodsOk

`func (o *O11yDeploymentRecord) GetAvailablePodsOk() (*int64, bool)`

GetAvailablePodsOk returns a tuple with the AvailablePods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePods

`func (o *O11yDeploymentRecord) SetAvailablePods(v int64)`

SetAvailablePods sets AvailablePods field to given value.

### HasAvailablePods

`func (o *O11yDeploymentRecord) HasAvailablePods() bool`

HasAvailablePods returns a boolean if a field has been set.

### GetDeploymentCPU

`func (o *O11yDeploymentRecord) GetDeploymentCPU() float64`

GetDeploymentCPU returns the DeploymentCPU field if non-nil, zero value otherwise.

### GetDeploymentCPUOk

`func (o *O11yDeploymentRecord) GetDeploymentCPUOk() (*float64, bool)`

GetDeploymentCPUOk returns a tuple with the DeploymentCPU field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCPU

`func (o *O11yDeploymentRecord) SetDeploymentCPU(v float64)`

SetDeploymentCPU sets DeploymentCPU field to given value.

### HasDeploymentCPU

`func (o *O11yDeploymentRecord) HasDeploymentCPU() bool`

HasDeploymentCPU returns a boolean if a field has been set.

### GetDeploymentCPULimit

`func (o *O11yDeploymentRecord) GetDeploymentCPULimit() float64`

GetDeploymentCPULimit returns the DeploymentCPULimit field if non-nil, zero value otherwise.

### GetDeploymentCPULimitOk

`func (o *O11yDeploymentRecord) GetDeploymentCPULimitOk() (*float64, bool)`

GetDeploymentCPULimitOk returns a tuple with the DeploymentCPULimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCPULimit

`func (o *O11yDeploymentRecord) SetDeploymentCPULimit(v float64)`

SetDeploymentCPULimit sets DeploymentCPULimit field to given value.

### HasDeploymentCPULimit

`func (o *O11yDeploymentRecord) HasDeploymentCPULimit() bool`

HasDeploymentCPULimit returns a boolean if a field has been set.

### GetDeploymentCPURequest

`func (o *O11yDeploymentRecord) GetDeploymentCPURequest() float64`

GetDeploymentCPURequest returns the DeploymentCPURequest field if non-nil, zero value otherwise.

### GetDeploymentCPURequestOk

`func (o *O11yDeploymentRecord) GetDeploymentCPURequestOk() (*float64, bool)`

GetDeploymentCPURequestOk returns a tuple with the DeploymentCPURequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCPURequest

`func (o *O11yDeploymentRecord) SetDeploymentCPURequest(v float64)`

SetDeploymentCPURequest sets DeploymentCPURequest field to given value.

### HasDeploymentCPURequest

`func (o *O11yDeploymentRecord) HasDeploymentCPURequest() bool`

HasDeploymentCPURequest returns a boolean if a field has been set.

### GetDeploymentMemory

`func (o *O11yDeploymentRecord) GetDeploymentMemory() float64`

GetDeploymentMemory returns the DeploymentMemory field if non-nil, zero value otherwise.

### GetDeploymentMemoryOk

`func (o *O11yDeploymentRecord) GetDeploymentMemoryOk() (*float64, bool)`

GetDeploymentMemoryOk returns a tuple with the DeploymentMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentMemory

`func (o *O11yDeploymentRecord) SetDeploymentMemory(v float64)`

SetDeploymentMemory sets DeploymentMemory field to given value.

### HasDeploymentMemory

`func (o *O11yDeploymentRecord) HasDeploymentMemory() bool`

HasDeploymentMemory returns a boolean if a field has been set.

### GetDeploymentMemoryLimit

`func (o *O11yDeploymentRecord) GetDeploymentMemoryLimit() float64`

GetDeploymentMemoryLimit returns the DeploymentMemoryLimit field if non-nil, zero value otherwise.

### GetDeploymentMemoryLimitOk

`func (o *O11yDeploymentRecord) GetDeploymentMemoryLimitOk() (*float64, bool)`

GetDeploymentMemoryLimitOk returns a tuple with the DeploymentMemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentMemoryLimit

`func (o *O11yDeploymentRecord) SetDeploymentMemoryLimit(v float64)`

SetDeploymentMemoryLimit sets DeploymentMemoryLimit field to given value.

### HasDeploymentMemoryLimit

`func (o *O11yDeploymentRecord) HasDeploymentMemoryLimit() bool`

HasDeploymentMemoryLimit returns a boolean if a field has been set.

### GetDeploymentMemoryRequest

`func (o *O11yDeploymentRecord) GetDeploymentMemoryRequest() float64`

GetDeploymentMemoryRequest returns the DeploymentMemoryRequest field if non-nil, zero value otherwise.

### GetDeploymentMemoryRequestOk

`func (o *O11yDeploymentRecord) GetDeploymentMemoryRequestOk() (*float64, bool)`

GetDeploymentMemoryRequestOk returns a tuple with the DeploymentMemoryRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentMemoryRequest

`func (o *O11yDeploymentRecord) SetDeploymentMemoryRequest(v float64)`

SetDeploymentMemoryRequest sets DeploymentMemoryRequest field to given value.

### HasDeploymentMemoryRequest

`func (o *O11yDeploymentRecord) HasDeploymentMemoryRequest() bool`

HasDeploymentMemoryRequest returns a boolean if a field has been set.

### GetDeploymentName

`func (o *O11yDeploymentRecord) GetDeploymentName() string`

GetDeploymentName returns the DeploymentName field if non-nil, zero value otherwise.

### GetDeploymentNameOk

`func (o *O11yDeploymentRecord) GetDeploymentNameOk() (*string, bool)`

GetDeploymentNameOk returns a tuple with the DeploymentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentName

`func (o *O11yDeploymentRecord) SetDeploymentName(v string)`

SetDeploymentName sets DeploymentName field to given value.

### HasDeploymentName

`func (o *O11yDeploymentRecord) HasDeploymentName() bool`

HasDeploymentName returns a boolean if a field has been set.

### GetDesiredPods

`func (o *O11yDeploymentRecord) GetDesiredPods() int64`

GetDesiredPods returns the DesiredPods field if non-nil, zero value otherwise.

### GetDesiredPodsOk

`func (o *O11yDeploymentRecord) GetDesiredPodsOk() (*int64, bool)`

GetDesiredPodsOk returns a tuple with the DesiredPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredPods

`func (o *O11yDeploymentRecord) SetDesiredPods(v int64)`

SetDesiredPods sets DesiredPods field to given value.

### HasDesiredPods

`func (o *O11yDeploymentRecord) HasDesiredPods() bool`

HasDesiredPods returns a boolean if a field has been set.

### GetMeta

`func (o *O11yDeploymentRecord) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *O11yDeploymentRecord) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *O11yDeploymentRecord) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *O11yDeploymentRecord) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPodCountsByPhase

`func (o *O11yDeploymentRecord) GetPodCountsByPhase() O11yPodCountsByPhase`

GetPodCountsByPhase returns the PodCountsByPhase field if non-nil, zero value otherwise.

### GetPodCountsByPhaseOk

`func (o *O11yDeploymentRecord) GetPodCountsByPhaseOk() (*O11yPodCountsByPhase, bool)`

GetPodCountsByPhaseOk returns a tuple with the PodCountsByPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCountsByPhase

`func (o *O11yDeploymentRecord) SetPodCountsByPhase(v O11yPodCountsByPhase)`

SetPodCountsByPhase sets PodCountsByPhase field to given value.

### HasPodCountsByPhase

`func (o *O11yDeploymentRecord) HasPodCountsByPhase() bool`

HasPodCountsByPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


