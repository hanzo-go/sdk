# MqServerInfoCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Cluster name. | [optional] 
**Leader** | Pointer to **string** | Current Raft leader server name. | [optional] 
**Replicas** | Pointer to [**[]MqServerInfoClusterReplicasInner**](MqServerInfoClusterReplicasInner.md) |  | [optional] 

## Methods

### NewMqServerInfoCluster

`func NewMqServerInfoCluster() *MqServerInfoCluster`

NewMqServerInfoCluster instantiates a new MqServerInfoCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqServerInfoClusterWithDefaults

`func NewMqServerInfoClusterWithDefaults() *MqServerInfoCluster`

NewMqServerInfoClusterWithDefaults instantiates a new MqServerInfoCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MqServerInfoCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MqServerInfoCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MqServerInfoCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MqServerInfoCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLeader

`func (o *MqServerInfoCluster) GetLeader() string`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *MqServerInfoCluster) GetLeaderOk() (*string, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *MqServerInfoCluster) SetLeader(v string)`

SetLeader sets Leader field to given value.

### HasLeader

`func (o *MqServerInfoCluster) HasLeader() bool`

HasLeader returns a boolean if a field has been set.

### GetReplicas

`func (o *MqServerInfoCluster) GetReplicas() []MqServerInfoClusterReplicasInner`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *MqServerInfoCluster) GetReplicasOk() (*[]MqServerInfoClusterReplicasInner, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *MqServerInfoCluster) SetReplicas(v []MqServerInfoClusterReplicasInner)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *MqServerInfoCluster) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


