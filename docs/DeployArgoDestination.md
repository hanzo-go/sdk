# DeployArgoDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | ArgoCD allows a destination by cluster name; omitted for the in-cluster projection. | [optional] 
**Namespace** | Pointer to **string** | Namespace is where in that cluster the workload lands. \&quot;*\&quot; on a project&#39;s destination fence means any namespace. | [optional] 
**Server** | Pointer to **string** | Server is the cluster API URL the application reconciles into. Everything this plane projects lands in the cluster it runs in, so it is https://kubernetes.default.svc — except on a project&#39;s destination fence, where \&quot;*\&quot; means any cluster. | [optional] 

## Methods

### NewDeployArgoDestination

`func NewDeployArgoDestination() *DeployArgoDestination`

NewDeployArgoDestination instantiates a new DeployArgoDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployArgoDestinationWithDefaults

`func NewDeployArgoDestinationWithDefaults() *DeployArgoDestination`

NewDeployArgoDestinationWithDefaults instantiates a new DeployArgoDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeployArgoDestination) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeployArgoDestination) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeployArgoDestination) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeployArgoDestination) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *DeployArgoDestination) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DeployArgoDestination) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DeployArgoDestination) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *DeployArgoDestination) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetServer

`func (o *DeployArgoDestination) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *DeployArgoDestination) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *DeployArgoDestination) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *DeployArgoDestination) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


