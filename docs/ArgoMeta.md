# ArgoMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTimestamp** | Pointer to **string** | CreationTimestamp is when the source object was created, RFC 3339 to the second. Empty for a synthesized project. | [optional] 
**Labels** | Pointer to **map[string]string** | Labels are the labels this projection puts on the row, not the source object&#39;s full label set. An application projected from an App CR carries hanzo.ai/instance (its name), hanzo.ai/env (main, test or dev, from the namespace it was read from) and hanzo.ai/org when the CR declares a tenant. A Hanzo CD Application carries the CR&#39;s own labels verbatim. A project reflected from IAM carries hanzo.ai/org alone. | [optional] 
**Name** | Pointer to **string** | Name is the projected object&#39;s name: the App CR&#39;s metadata.name for an application, the CD Application&#39;s name for a CD row, and the IAM project name for a project. | [optional] 
**Namespace** | Pointer to **string** | Namespace is the namespace the source object was read from — the tenant or platform namespace for an App CR, CD&#39;s controller namespace for a CD row. Empty for a project synthesized here, which lives in no namespace. | [optional] 
**Uid** | Pointer to **string** | UID is the k8s metadata.uid of the source object, which is what the SPA keys a row on across refreshes. Empty for a synthesized project — there is no object to take one from. | [optional] 

## Methods

### NewArgoMeta

`func NewArgoMeta() *ArgoMeta`

NewArgoMeta instantiates a new ArgoMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoMetaWithDefaults

`func NewArgoMetaWithDefaults() *ArgoMeta`

NewArgoMetaWithDefaults instantiates a new ArgoMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTimestamp

`func (o *ArgoMeta) GetCreationTimestamp() string`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *ArgoMeta) GetCreationTimestampOk() (*string, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *ArgoMeta) SetCreationTimestamp(v string)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *ArgoMeta) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetLabels

`func (o *ArgoMeta) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ArgoMeta) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ArgoMeta) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ArgoMeta) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *ArgoMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArgoMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArgoMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArgoMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *ArgoMeta) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ArgoMeta) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ArgoMeta) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ArgoMeta) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetUid

`func (o *ArgoMeta) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ArgoMeta) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ArgoMeta) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ArgoMeta) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


