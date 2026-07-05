# PaasDOKSCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int32** |  | [optional] 
**NodePools** | Pointer to [**[]PaasNodePool**](PaasNodePool.md) |  | [optional] 
**Ha** | Pointer to **bool** |  | [optional] 
**Endpoint** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPaasDOKSCluster

`func NewPaasDOKSCluster() *PaasDOKSCluster`

NewPaasDOKSCluster instantiates a new PaasDOKSCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaasDOKSClusterWithDefaults

`func NewPaasDOKSClusterWithDefaults() *PaasDOKSCluster`

NewPaasDOKSClusterWithDefaults instantiates a new PaasDOKSCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaasDOKSCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaasDOKSCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaasDOKSCluster) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaasDOKSCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PaasDOKSCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaasDOKSCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaasDOKSCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaasDOKSCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *PaasDOKSCluster) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PaasDOKSCluster) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PaasDOKSCluster) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PaasDOKSCluster) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetVersion

`func (o *PaasDOKSCluster) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PaasDOKSCluster) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PaasDOKSCluster) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PaasDOKSCluster) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStatus

`func (o *PaasDOKSCluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaasDOKSCluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaasDOKSCluster) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaasDOKSCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNodeCount

`func (o *PaasDOKSCluster) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *PaasDOKSCluster) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *PaasDOKSCluster) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *PaasDOKSCluster) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetNodePools

`func (o *PaasDOKSCluster) GetNodePools() []PaasNodePool`

GetNodePools returns the NodePools field if non-nil, zero value otherwise.

### GetNodePoolsOk

`func (o *PaasDOKSCluster) GetNodePoolsOk() (*[]PaasNodePool, bool)`

GetNodePoolsOk returns a tuple with the NodePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePools

`func (o *PaasDOKSCluster) SetNodePools(v []PaasNodePool)`

SetNodePools sets NodePools field to given value.

### HasNodePools

`func (o *PaasDOKSCluster) HasNodePools() bool`

HasNodePools returns a boolean if a field has been set.

### GetHa

`func (o *PaasDOKSCluster) GetHa() bool`

GetHa returns the Ha field if non-nil, zero value otherwise.

### GetHaOk

`func (o *PaasDOKSCluster) GetHaOk() (*bool, bool)`

GetHaOk returns a tuple with the Ha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHa

`func (o *PaasDOKSCluster) SetHa(v bool)`

SetHa sets Ha field to given value.

### HasHa

`func (o *PaasDOKSCluster) HasHa() bool`

HasHa returns a boolean if a field has been set.

### GetEndpoint

`func (o *PaasDOKSCluster) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *PaasDOKSCluster) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *PaasDOKSCluster) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *PaasDOKSCluster) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PaasDOKSCluster) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PaasDOKSCluster) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PaasDOKSCluster) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PaasDOKSCluster) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


