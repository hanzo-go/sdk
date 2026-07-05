# EngineRayClusterCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Namespace** | Pointer to **string** |  | [optional] [default to "default"]
**RayVersion** | Pointer to **string** |  | [optional] [default to "2.9.0"]
**Head** | Pointer to [**EngineRayClusterCreateHead**](EngineRayClusterCreateHead.md) |  | [optional] 
**Workers** | Pointer to [**[]EngineRayClusterCreateWorkersInner**](EngineRayClusterCreateWorkersInner.md) |  | [optional] 

## Methods

### NewEngineRayClusterCreate

`func NewEngineRayClusterCreate(name string, ) *EngineRayClusterCreate`

NewEngineRayClusterCreate instantiates a new EngineRayClusterCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineRayClusterCreateWithDefaults

`func NewEngineRayClusterCreateWithDefaults() *EngineRayClusterCreate`

NewEngineRayClusterCreateWithDefaults instantiates a new EngineRayClusterCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EngineRayClusterCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EngineRayClusterCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EngineRayClusterCreate) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *EngineRayClusterCreate) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *EngineRayClusterCreate) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *EngineRayClusterCreate) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *EngineRayClusterCreate) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetRayVersion

`func (o *EngineRayClusterCreate) GetRayVersion() string`

GetRayVersion returns the RayVersion field if non-nil, zero value otherwise.

### GetRayVersionOk

`func (o *EngineRayClusterCreate) GetRayVersionOk() (*string, bool)`

GetRayVersionOk returns a tuple with the RayVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRayVersion

`func (o *EngineRayClusterCreate) SetRayVersion(v string)`

SetRayVersion sets RayVersion field to given value.

### HasRayVersion

`func (o *EngineRayClusterCreate) HasRayVersion() bool`

HasRayVersion returns a boolean if a field has been set.

### GetHead

`func (o *EngineRayClusterCreate) GetHead() EngineRayClusterCreateHead`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *EngineRayClusterCreate) GetHeadOk() (*EngineRayClusterCreateHead, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *EngineRayClusterCreate) SetHead(v EngineRayClusterCreateHead)`

SetHead sets Head field to given value.

### HasHead

`func (o *EngineRayClusterCreate) HasHead() bool`

HasHead returns a boolean if a field has been set.

### GetWorkers

`func (o *EngineRayClusterCreate) GetWorkers() []EngineRayClusterCreateWorkersInner`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *EngineRayClusterCreate) GetWorkersOk() (*[]EngineRayClusterCreateWorkersInner, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *EngineRayClusterCreate) SetWorkers(v []EngineRayClusterCreateWorkersInner)`

SetWorkers sets Workers field to given value.

### HasWorkers

`func (o *EngineRayClusterCreate) HasWorkers() bool`

HasWorkers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


