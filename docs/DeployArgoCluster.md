# DeployArgoCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionState** | Pointer to [**DeployArgoConnectionState**](DeployArgoConnectionState.md) | ConnectionState is whether the destination is reachable. | [optional] 
**Info** | Pointer to [**DeployArgoClusterInfo**](DeployArgoClusterInfo.md) | Info is the connection state again plus the count of applications targeting this destination. | [optional] 
**Name** | Pointer to **string** | Name is what the Destination column shows: \&quot;in-cluster\&quot; for this cluster, otherwise whatever spec.destination.name declares, falling back to the server URL when it declares none. | [optional] 
**Server** | Pointer to **string** | Server is the destination&#39;s API URL, and the key the list is deduplicated by. https://kubernetes.default.svc is this cluster. | [optional] 

## Methods

### NewDeployArgoCluster

`func NewDeployArgoCluster() *DeployArgoCluster`

NewDeployArgoCluster instantiates a new DeployArgoCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployArgoClusterWithDefaults

`func NewDeployArgoClusterWithDefaults() *DeployArgoCluster`

NewDeployArgoClusterWithDefaults instantiates a new DeployArgoCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionState

`func (o *DeployArgoCluster) GetConnectionState() DeployArgoConnectionState`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *DeployArgoCluster) GetConnectionStateOk() (*DeployArgoConnectionState, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *DeployArgoCluster) SetConnectionState(v DeployArgoConnectionState)`

SetConnectionState sets ConnectionState field to given value.

### HasConnectionState

`func (o *DeployArgoCluster) HasConnectionState() bool`

HasConnectionState returns a boolean if a field has been set.

### GetInfo

`func (o *DeployArgoCluster) GetInfo() DeployArgoClusterInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *DeployArgoCluster) GetInfoOk() (*DeployArgoClusterInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *DeployArgoCluster) SetInfo(v DeployArgoClusterInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *DeployArgoCluster) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetName

`func (o *DeployArgoCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeployArgoCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeployArgoCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeployArgoCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServer

`func (o *DeployArgoCluster) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *DeployArgoCluster) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *DeployArgoCluster) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *DeployArgoCluster) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


