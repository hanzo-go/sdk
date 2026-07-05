# EngineClusterCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Provider** | **string** |  | 
**Region** | Pointer to **string** |  | [optional] 
**Kubeconfig** | Pointer to **string** | Base64-encoded kubeconfig for the cluster | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewEngineClusterCreate

`func NewEngineClusterCreate(name string, provider string, ) *EngineClusterCreate`

NewEngineClusterCreate instantiates a new EngineClusterCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineClusterCreateWithDefaults

`func NewEngineClusterCreateWithDefaults() *EngineClusterCreate`

NewEngineClusterCreateWithDefaults instantiates a new EngineClusterCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EngineClusterCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EngineClusterCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EngineClusterCreate) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *EngineClusterCreate) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EngineClusterCreate) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EngineClusterCreate) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetRegion

`func (o *EngineClusterCreate) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EngineClusterCreate) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EngineClusterCreate) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *EngineClusterCreate) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetKubeconfig

`func (o *EngineClusterCreate) GetKubeconfig() string`

GetKubeconfig returns the Kubeconfig field if non-nil, zero value otherwise.

### GetKubeconfigOk

`func (o *EngineClusterCreate) GetKubeconfigOk() (*string, bool)`

GetKubeconfigOk returns a tuple with the Kubeconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeconfig

`func (o *EngineClusterCreate) SetKubeconfig(v string)`

SetKubeconfig sets Kubeconfig field to given value.

### HasKubeconfig

`func (o *EngineClusterCreate) HasKubeconfig() bool`

HasKubeconfig returns a boolean if a field has been set.

### GetLabels

`func (o *EngineClusterCreate) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EngineClusterCreate) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EngineClusterCreate) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EngineClusterCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


