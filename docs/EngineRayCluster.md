# EngineRayCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Head** | Pointer to [**EngineRayClusterHead**](EngineRayClusterHead.md) |  | [optional] 
**Workers** | Pointer to [**[]EngineRayClusterWorkersInner**](EngineRayClusterWorkersInner.md) |  | [optional] 
**RayVersion** | Pointer to **string** |  | [optional] 
**DashboardUrl** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEngineRayCluster

`func NewEngineRayCluster() *EngineRayCluster`

NewEngineRayCluster instantiates a new EngineRayCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineRayClusterWithDefaults

`func NewEngineRayClusterWithDefaults() *EngineRayCluster`

NewEngineRayClusterWithDefaults instantiates a new EngineRayCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EngineRayCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EngineRayCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EngineRayCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EngineRayCluster) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *EngineRayCluster) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *EngineRayCluster) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *EngineRayCluster) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *EngineRayCluster) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetStatus

`func (o *EngineRayCluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EngineRayCluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EngineRayCluster) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EngineRayCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHead

`func (o *EngineRayCluster) GetHead() EngineRayClusterHead`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *EngineRayCluster) GetHeadOk() (*EngineRayClusterHead, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *EngineRayCluster) SetHead(v EngineRayClusterHead)`

SetHead sets Head field to given value.

### HasHead

`func (o *EngineRayCluster) HasHead() bool`

HasHead returns a boolean if a field has been set.

### GetWorkers

`func (o *EngineRayCluster) GetWorkers() []EngineRayClusterWorkersInner`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *EngineRayCluster) GetWorkersOk() (*[]EngineRayClusterWorkersInner, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *EngineRayCluster) SetWorkers(v []EngineRayClusterWorkersInner)`

SetWorkers sets Workers field to given value.

### HasWorkers

`func (o *EngineRayCluster) HasWorkers() bool`

HasWorkers returns a boolean if a field has been set.

### GetRayVersion

`func (o *EngineRayCluster) GetRayVersion() string`

GetRayVersion returns the RayVersion field if non-nil, zero value otherwise.

### GetRayVersionOk

`func (o *EngineRayCluster) GetRayVersionOk() (*string, bool)`

GetRayVersionOk returns a tuple with the RayVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRayVersion

`func (o *EngineRayCluster) SetRayVersion(v string)`

SetRayVersion sets RayVersion field to given value.

### HasRayVersion

`func (o *EngineRayCluster) HasRayVersion() bool`

HasRayVersion returns a boolean if a field has been set.

### GetDashboardUrl

`func (o *EngineRayCluster) GetDashboardUrl() string`

GetDashboardUrl returns the DashboardUrl field if non-nil, zero value otherwise.

### GetDashboardUrlOk

`func (o *EngineRayCluster) GetDashboardUrlOk() (*string, bool)`

GetDashboardUrlOk returns a tuple with the DashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardUrl

`func (o *EngineRayCluster) SetDashboardUrl(v string)`

SetDashboardUrl sets DashboardUrl field to given value.

### HasDashboardUrl

`func (o *EngineRayCluster) HasDashboardUrl() bool`

HasDashboardUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EngineRayCluster) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EngineRayCluster) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EngineRayCluster) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EngineRayCluster) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


