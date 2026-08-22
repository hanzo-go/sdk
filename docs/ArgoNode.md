# ArgoNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is the object&#39;s creationTimestamp, RFC 3339 UTC to the second. Absent when the object carries none. | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Health** | Pointer to [**ArgoHealth**](ArgoHealth.md) | Health is the node&#39;s own derived health. Always present on a node of this tree; a kind with no health signal of its own reports Healthy, since a ConfigMap existing IS its healthy state. | [optional] 
**Images** | Pointer to **[]string** | Images are the container images running on this node. Always absent — the tag travels as the \&quot;Image Tag\&quot; chip in Info instead, which is where the SPA reads it on a node. | [optional] 
**Info** | Pointer to [**[]ArgoInfoItem**](ArgoInfoItem.md) | Info are the chips shown on the node. At most one: the image tag — the RUNNING tag on a Deployment, ReplicaSet or Pod, and the DECLARED tag on the App CR at the root. Absent on a node that carries no image at all. | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**ParentRefs** | Pointer to [**[]ArgoResourceRef**](ArgoResourceRef.md) | ParentRefs are the node&#39;s edges UPWARD, which is how the SPA draws the DAG from this flat list. Exactly one entry where present: a depth-1 object points at the App CR, a ReplicaSet at its Deployment, a Pod at its ReplicaSet (or at the Deployment whose selector matches it, when the ReplicaSet is gone). Absent on the root. | [optional] 
**ResourceVersion** | Pointer to **string** | ResourceVersion is the k8s version a watch would resume from. Always empty: the tree is rebuilt from live reads on every request, including on every frame of the SSE stream, so there is no revision to resume from. | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewArgoNode

`func NewArgoNode() *ArgoNode`

NewArgoNode instantiates a new ArgoNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoNodeWithDefaults

`func NewArgoNodeWithDefaults() *ArgoNode`

NewArgoNodeWithDefaults instantiates a new ArgoNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ArgoNode) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArgoNode) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArgoNode) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArgoNode) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGroup

`func (o *ArgoNode) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ArgoNode) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ArgoNode) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ArgoNode) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHealth

`func (o *ArgoNode) GetHealth() ArgoHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ArgoNode) GetHealthOk() (*ArgoHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ArgoNode) SetHealth(v ArgoHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ArgoNode) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetImages

`func (o *ArgoNode) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ArgoNode) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ArgoNode) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *ArgoNode) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetInfo

`func (o *ArgoNode) GetInfo() []ArgoInfoItem`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ArgoNode) GetInfoOk() (*[]ArgoInfoItem, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ArgoNode) SetInfo(v []ArgoInfoItem)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ArgoNode) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetKind

`func (o *ArgoNode) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ArgoNode) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ArgoNode) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ArgoNode) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ArgoNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArgoNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArgoNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArgoNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *ArgoNode) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ArgoNode) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ArgoNode) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ArgoNode) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetParentRefs

`func (o *ArgoNode) GetParentRefs() []ArgoResourceRef`

GetParentRefs returns the ParentRefs field if non-nil, zero value otherwise.

### GetParentRefsOk

`func (o *ArgoNode) GetParentRefsOk() (*[]ArgoResourceRef, bool)`

GetParentRefsOk returns a tuple with the ParentRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRefs

`func (o *ArgoNode) SetParentRefs(v []ArgoResourceRef)`

SetParentRefs sets ParentRefs field to given value.

### HasParentRefs

`func (o *ArgoNode) HasParentRefs() bool`

HasParentRefs returns a boolean if a field has been set.

### GetResourceVersion

`func (o *ArgoNode) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *ArgoNode) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *ArgoNode) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *ArgoNode) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetUid

`func (o *ArgoNode) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ArgoNode) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ArgoNode) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ArgoNode) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetVersion

`func (o *ArgoNode) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ArgoNode) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ArgoNode) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ArgoNode) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


