# DeployArgoResourceRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Group is the object&#39;s API group: empty for the core group (Pod, Service, ConfigMap), otherwise apps, networking.k8s.io, autoscaling or policy — and hanzo.ai for the App CR at the root. | [optional] 
**Kind** | Pointer to **string** | Kind is the object kind. The root is the App CR; below it come Deployment, ReplicaSet, Pod, Service, Ingress, HorizontalPodAutoscaler, PodDisruptionBudget and ConfigMap. Never Secret — the walk does not visit them, so no materialized environment can reach the tree. | [optional] 
**Name** | Pointer to **string** | Name is the object&#39;s metadata.name. | [optional] 
**Namespace** | Pointer to **string** | Namespace is the namespace the walk ran in, the same for every node of one tree. | [optional] 
**Uid** | Pointer to **string** | UID is the object&#39;s metadata.uid. Absent on a PARENT reference, which addresses its target by kind and name rather than by identity. | [optional] 
**Version** | Pointer to **string** | Version is the object&#39;s API version as the live object reports it: v1 for every kind the walk reaches except the HorizontalPodAutoscaler, which is autoscaling/v2. | [optional] 

## Methods

### NewDeployArgoResourceRef

`func NewDeployArgoResourceRef() *DeployArgoResourceRef`

NewDeployArgoResourceRef instantiates a new DeployArgoResourceRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployArgoResourceRefWithDefaults

`func NewDeployArgoResourceRefWithDefaults() *DeployArgoResourceRef`

NewDeployArgoResourceRefWithDefaults instantiates a new DeployArgoResourceRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *DeployArgoResourceRef) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DeployArgoResourceRef) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DeployArgoResourceRef) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DeployArgoResourceRef) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetKind

`func (o *DeployArgoResourceRef) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DeployArgoResourceRef) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DeployArgoResourceRef) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DeployArgoResourceRef) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *DeployArgoResourceRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeployArgoResourceRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeployArgoResourceRef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeployArgoResourceRef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *DeployArgoResourceRef) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *DeployArgoResourceRef) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *DeployArgoResourceRef) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *DeployArgoResourceRef) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetUid

`func (o *DeployArgoResourceRef) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *DeployArgoResourceRef) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *DeployArgoResourceRef) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *DeployArgoResourceRef) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetVersion

`func (o *DeployArgoResourceRef) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeployArgoResourceRef) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeployArgoResourceRef) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeployArgoResourceRef) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


