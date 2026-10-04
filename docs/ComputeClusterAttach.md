# ComputeClusterAttach

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **bool** | Default marks this the org&#39;s default cluster for scheduling. | [optional] 
**Kubeconfig** | Pointer to **string** | Kubeconfig is the cluster&#39;s kubeconfig, verbatim. Required — a body without one is not an attach. | [optional] 
**Name** | Pointer to **string** | Name is the fleet-local name for the cluster; lower-cased, and the key the detach route addresses it by. Required. | [optional] 
**Provider** | Pointer to **string** | Provider is a free-form label for where the cluster runs (\&quot;gke\&quot;, \&quot;on-prem\&quot;); it is display only, not a routing key. | [optional] 

## Methods

### NewComputeClusterAttach

`func NewComputeClusterAttach() *ComputeClusterAttach`

NewComputeClusterAttach instantiates a new ComputeClusterAttach object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeClusterAttachWithDefaults

`func NewComputeClusterAttachWithDefaults() *ComputeClusterAttach`

NewComputeClusterAttachWithDefaults instantiates a new ComputeClusterAttach object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *ComputeClusterAttach) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ComputeClusterAttach) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ComputeClusterAttach) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ComputeClusterAttach) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetKubeconfig

`func (o *ComputeClusterAttach) GetKubeconfig() string`

GetKubeconfig returns the Kubeconfig field if non-nil, zero value otherwise.

### GetKubeconfigOk

`func (o *ComputeClusterAttach) GetKubeconfigOk() (*string, bool)`

GetKubeconfigOk returns a tuple with the Kubeconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeconfig

`func (o *ComputeClusterAttach) SetKubeconfig(v string)`

SetKubeconfig sets Kubeconfig field to given value.

### HasKubeconfig

`func (o *ComputeClusterAttach) HasKubeconfig() bool`

HasKubeconfig returns a boolean if a field has been set.

### GetName

`func (o *ComputeClusterAttach) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputeClusterAttach) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputeClusterAttach) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComputeClusterAttach) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *ComputeClusterAttach) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ComputeClusterAttach) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ComputeClusterAttach) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ComputeClusterAttach) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


