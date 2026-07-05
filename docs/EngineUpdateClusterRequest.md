# EngineUpdateClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Kubeconfig** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewEngineUpdateClusterRequest

`func NewEngineUpdateClusterRequest() *EngineUpdateClusterRequest`

NewEngineUpdateClusterRequest instantiates a new EngineUpdateClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineUpdateClusterRequestWithDefaults

`func NewEngineUpdateClusterRequestWithDefaults() *EngineUpdateClusterRequest`

NewEngineUpdateClusterRequestWithDefaults instantiates a new EngineUpdateClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EngineUpdateClusterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EngineUpdateClusterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EngineUpdateClusterRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EngineUpdateClusterRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKubeconfig

`func (o *EngineUpdateClusterRequest) GetKubeconfig() string`

GetKubeconfig returns the Kubeconfig field if non-nil, zero value otherwise.

### GetKubeconfigOk

`func (o *EngineUpdateClusterRequest) GetKubeconfigOk() (*string, bool)`

GetKubeconfigOk returns a tuple with the Kubeconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeconfig

`func (o *EngineUpdateClusterRequest) SetKubeconfig(v string)`

SetKubeconfig sets Kubeconfig field to given value.

### HasKubeconfig

`func (o *EngineUpdateClusterRequest) HasKubeconfig() bool`

HasKubeconfig returns a boolean if a field has been set.

### GetLabels

`func (o *EngineUpdateClusterRequest) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EngineUpdateClusterRequest) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EngineUpdateClusterRequest) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EngineUpdateClusterRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


