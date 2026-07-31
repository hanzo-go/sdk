# CloudCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**IdlePVCs** | Pointer to **int32** |  | [optional] 
**MonthlyCents** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodePools** | Pointer to **int32** |  | [optional] 
**Nodes** | Pointer to **int32** |  | [optional] 
**Pods** | Pointer to **int32** |  | [optional] 
**Pools** | Pointer to [**[]CloudNodePool**](CloudNodePool.md) |  | [optional] 
**Pvcs** | Pointer to **int32** |  | [optional] 
**Pvs** | Pointer to **int32** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**ScanError** | Pointer to **string** |  | [optional] 
**Scanned** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudCluster

`func NewCloudCluster() *CloudCluster`

NewCloudCluster instantiates a new CloudCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudClusterWithDefaults

`func NewCloudClusterWithDefaults() *CloudCluster`

NewCloudClusterWithDefaults instantiates a new CloudCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCluster) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdlePVCs

`func (o *CloudCluster) GetIdlePVCs() int32`

GetIdlePVCs returns the IdlePVCs field if non-nil, zero value otherwise.

### GetIdlePVCsOk

`func (o *CloudCluster) GetIdlePVCsOk() (*int32, bool)`

GetIdlePVCsOk returns a tuple with the IdlePVCs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdlePVCs

`func (o *CloudCluster) SetIdlePVCs(v int32)`

SetIdlePVCs sets IdlePVCs field to given value.

### HasIdlePVCs

`func (o *CloudCluster) HasIdlePVCs() bool`

HasIdlePVCs returns a boolean if a field has been set.

### GetMonthlyCents

`func (o *CloudCluster) GetMonthlyCents() int32`

GetMonthlyCents returns the MonthlyCents field if non-nil, zero value otherwise.

### GetMonthlyCentsOk

`func (o *CloudCluster) GetMonthlyCentsOk() (*int32, bool)`

GetMonthlyCentsOk returns a tuple with the MonthlyCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCents

`func (o *CloudCluster) SetMonthlyCents(v int32)`

SetMonthlyCents sets MonthlyCents field to given value.

### HasMonthlyCents

`func (o *CloudCluster) HasMonthlyCents() bool`

HasMonthlyCents returns a boolean if a field has been set.

### GetName

`func (o *CloudCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodePools

`func (o *CloudCluster) GetNodePools() int32`

GetNodePools returns the NodePools field if non-nil, zero value otherwise.

### GetNodePoolsOk

`func (o *CloudCluster) GetNodePoolsOk() (*int32, bool)`

GetNodePoolsOk returns a tuple with the NodePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePools

`func (o *CloudCluster) SetNodePools(v int32)`

SetNodePools sets NodePools field to given value.

### HasNodePools

`func (o *CloudCluster) HasNodePools() bool`

HasNodePools returns a boolean if a field has been set.

### GetNodes

`func (o *CloudCluster) GetNodes() int32`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *CloudCluster) GetNodesOk() (*int32, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *CloudCluster) SetNodes(v int32)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *CloudCluster) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetPods

`func (o *CloudCluster) GetPods() int32`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *CloudCluster) GetPodsOk() (*int32, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *CloudCluster) SetPods(v int32)`

SetPods sets Pods field to given value.

### HasPods

`func (o *CloudCluster) HasPods() bool`

HasPods returns a boolean if a field has been set.

### GetPools

`func (o *CloudCluster) GetPools() []CloudNodePool`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *CloudCluster) GetPoolsOk() (*[]CloudNodePool, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *CloudCluster) SetPools(v []CloudNodePool)`

SetPools sets Pools field to given value.

### HasPools

`func (o *CloudCluster) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetPvcs

`func (o *CloudCluster) GetPvcs() int32`

GetPvcs returns the Pvcs field if non-nil, zero value otherwise.

### GetPvcsOk

`func (o *CloudCluster) GetPvcsOk() (*int32, bool)`

GetPvcsOk returns a tuple with the Pvcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvcs

`func (o *CloudCluster) SetPvcs(v int32)`

SetPvcs sets Pvcs field to given value.

### HasPvcs

`func (o *CloudCluster) HasPvcs() bool`

HasPvcs returns a boolean if a field has been set.

### GetPvs

`func (o *CloudCluster) GetPvs() int32`

GetPvs returns the Pvs field if non-nil, zero value otherwise.

### GetPvsOk

`func (o *CloudCluster) GetPvsOk() (*int32, bool)`

GetPvsOk returns a tuple with the Pvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvs

`func (o *CloudCluster) SetPvs(v int32)`

SetPvs sets Pvs field to given value.

### HasPvs

`func (o *CloudCluster) HasPvs() bool`

HasPvs returns a boolean if a field has been set.

### GetRegion

`func (o *CloudCluster) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudCluster) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudCluster) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudCluster) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetScanError

`func (o *CloudCluster) GetScanError() string`

GetScanError returns the ScanError field if non-nil, zero value otherwise.

### GetScanErrorOk

`func (o *CloudCluster) GetScanErrorOk() (*string, bool)`

GetScanErrorOk returns a tuple with the ScanError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanError

`func (o *CloudCluster) SetScanError(v string)`

SetScanError sets ScanError field to given value.

### HasScanError

`func (o *CloudCluster) HasScanError() bool`

HasScanError returns a boolean if a field has been set.

### GetScanned

`func (o *CloudCluster) GetScanned() bool`

GetScanned returns the Scanned field if non-nil, zero value otherwise.

### GetScannedOk

`func (o *CloudCluster) GetScannedOk() (*bool, bool)`

GetScannedOk returns a tuple with the Scanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanned

`func (o *CloudCluster) SetScanned(v bool)`

SetScanned sets Scanned field to given value.

### HasScanned

`func (o *CloudCluster) HasScanned() bool`

HasScanned returns a boolean if a field has been set.

### GetStatus

`func (o *CloudCluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudCluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudCluster) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *CloudCluster) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CloudCluster) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CloudCluster) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CloudCluster) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


