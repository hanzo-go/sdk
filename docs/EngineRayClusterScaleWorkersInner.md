# EngineRayClusterScaleWorkersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | **string** |  | 
**Replicas** | **int32** |  | 
**MinReplicas** | Pointer to **int32** |  | [optional] 
**MaxReplicas** | Pointer to **int32** |  | [optional] 

## Methods

### NewEngineRayClusterScaleWorkersInner

`func NewEngineRayClusterScaleWorkersInner(groupName string, replicas int32, ) *EngineRayClusterScaleWorkersInner`

NewEngineRayClusterScaleWorkersInner instantiates a new EngineRayClusterScaleWorkersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineRayClusterScaleWorkersInnerWithDefaults

`func NewEngineRayClusterScaleWorkersInnerWithDefaults() *EngineRayClusterScaleWorkersInner`

NewEngineRayClusterScaleWorkersInnerWithDefaults instantiates a new EngineRayClusterScaleWorkersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *EngineRayClusterScaleWorkersInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *EngineRayClusterScaleWorkersInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *EngineRayClusterScaleWorkersInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetReplicas

`func (o *EngineRayClusterScaleWorkersInner) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *EngineRayClusterScaleWorkersInner) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *EngineRayClusterScaleWorkersInner) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetMinReplicas

`func (o *EngineRayClusterScaleWorkersInner) GetMinReplicas() int32`

GetMinReplicas returns the MinReplicas field if non-nil, zero value otherwise.

### GetMinReplicasOk

`func (o *EngineRayClusterScaleWorkersInner) GetMinReplicasOk() (*int32, bool)`

GetMinReplicasOk returns a tuple with the MinReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReplicas

`func (o *EngineRayClusterScaleWorkersInner) SetMinReplicas(v int32)`

SetMinReplicas sets MinReplicas field to given value.

### HasMinReplicas

`func (o *EngineRayClusterScaleWorkersInner) HasMinReplicas() bool`

HasMinReplicas returns a boolean if a field has been set.

### GetMaxReplicas

`func (o *EngineRayClusterScaleWorkersInner) GetMaxReplicas() int32`

GetMaxReplicas returns the MaxReplicas field if non-nil, zero value otherwise.

### GetMaxReplicasOk

`func (o *EngineRayClusterScaleWorkersInner) GetMaxReplicasOk() (*int32, bool)`

GetMaxReplicasOk returns a tuple with the MaxReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReplicas

`func (o *EngineRayClusterScaleWorkersInner) SetMaxReplicas(v int32)`

SetMaxReplicas sets MaxReplicas field to given value.

### HasMaxReplicas

`func (o *EngineRayClusterScaleWorkersInner) HasMaxReplicas() bool`

HasMaxReplicas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


