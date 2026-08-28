# DeploymentDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | Pointer to [**[]ContainerDetail**](ContainerDetail.md) |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ReadyReplicas** | Pointer to **int32** |  | [optional] 
**Replicas** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewDeploymentDetail

`func NewDeploymentDetail() *DeploymentDetail`

NewDeploymentDetail instantiates a new DeploymentDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentDetailWithDefaults

`func NewDeploymentDetailWithDefaults() *DeploymentDetail`

NewDeploymentDetailWithDefaults instantiates a new DeploymentDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *DeploymentDetail) GetContainers() []ContainerDetail`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *DeploymentDetail) GetContainersOk() (*[]ContainerDetail, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *DeploymentDetail) SetContainers(v []ContainerDetail)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *DeploymentDetail) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetCreatedTime

`func (o *DeploymentDetail) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DeploymentDetail) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DeploymentDetail) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *DeploymentDetail) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetName

`func (o *DeploymentDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReadyReplicas

`func (o *DeploymentDetail) GetReadyReplicas() int32`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *DeploymentDetail) GetReadyReplicasOk() (*int32, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *DeploymentDetail) SetReadyReplicas(v int32)`

SetReadyReplicas sets ReadyReplicas field to given value.

### HasReadyReplicas

`func (o *DeploymentDetail) HasReadyReplicas() bool`

HasReadyReplicas returns a boolean if a field has been set.

### GetReplicas

`func (o *DeploymentDetail) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *DeploymentDetail) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *DeploymentDetail) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *DeploymentDetail) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


