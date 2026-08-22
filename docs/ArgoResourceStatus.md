# ArgoResourceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | Group is the object&#39;s API group: empty for the core group (Pod, Service, ConfigMap), otherwise apps, networking.k8s.io, autoscaling or policy — and hanzo.ai for the App CR itself. | [optional] 
**Health** | Pointer to [**ArgoHealth**](ArgoHealth.md) | Health is this object&#39;s own health, derived from its live state by the same rule the resource tree uses. | [optional] 
**Kind** | Pointer to **string** | Kind is the object kind — App, Deployment, ReplicaSet, Pod, Service, Ingress, HorizontalPodAutoscaler, PodDisruptionBudget, ConfigMap. Never Secret: the walk that produces these does not visit them. | [optional] 
**Name** | Pointer to **string** | Name is the object&#39;s metadata.name. | [optional] 
**Namespace** | Pointer to **string** | Namespace is the namespace the object was found in — the same one for every entry of an application, since the walk is confined to it. | [optional] 
**Status** | Pointer to **string** | Status is the APPLICATION&#39;s sync verdict repeated on every row, not a per-object one. The operator owns these children, so no child has a desired state of its own to compare against. | [optional] 
**Version** | Pointer to **string** | Version is the object&#39;s API version as the live object reports it: v1 for every kind here except the HorizontalPodAutoscaler, which is autoscaling/v2. | [optional] 

## Methods

### NewArgoResourceStatus

`func NewArgoResourceStatus() *ArgoResourceStatus`

NewArgoResourceStatus instantiates a new ArgoResourceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoResourceStatusWithDefaults

`func NewArgoResourceStatusWithDefaults() *ArgoResourceStatus`

NewArgoResourceStatusWithDefaults instantiates a new ArgoResourceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *ArgoResourceStatus) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ArgoResourceStatus) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ArgoResourceStatus) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ArgoResourceStatus) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHealth

`func (o *ArgoResourceStatus) GetHealth() ArgoHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ArgoResourceStatus) GetHealthOk() (*ArgoHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ArgoResourceStatus) SetHealth(v ArgoHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ArgoResourceStatus) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetKind

`func (o *ArgoResourceStatus) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ArgoResourceStatus) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ArgoResourceStatus) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ArgoResourceStatus) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ArgoResourceStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArgoResourceStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArgoResourceStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArgoResourceStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *ArgoResourceStatus) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ArgoResourceStatus) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ArgoResourceStatus) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ArgoResourceStatus) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetStatus

`func (o *ArgoResourceStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArgoResourceStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArgoResourceStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ArgoResourceStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *ArgoResourceStatus) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ArgoResourceStatus) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ArgoResourceStatus) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ArgoResourceStatus) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


