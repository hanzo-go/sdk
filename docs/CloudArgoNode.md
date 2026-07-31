# CloudArgoNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Health** | Pointer to [**CloudArgoHealth**](CloudArgoHealth.md) |  | [optional] 
**Images** | Pointer to **[]string** |  | [optional] 
**Info** | Pointer to [**[]CloudArgoInfoItem**](CloudArgoInfoItem.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**ParentRefs** | Pointer to [**[]CloudArgoResourceRef**](CloudArgoResourceRef.md) |  | [optional] 
**ResourceVersion** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudArgoNode

`func NewCloudArgoNode() *CloudArgoNode`

NewCloudArgoNode instantiates a new CloudArgoNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudArgoNodeWithDefaults

`func NewCloudArgoNodeWithDefaults() *CloudArgoNode`

NewCloudArgoNodeWithDefaults instantiates a new CloudArgoNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CloudArgoNode) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudArgoNode) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudArgoNode) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudArgoNode) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetGroup

`func (o *CloudArgoNode) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CloudArgoNode) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CloudArgoNode) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CloudArgoNode) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHealth

`func (o *CloudArgoNode) GetHealth() CloudArgoHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *CloudArgoNode) GetHealthOk() (*CloudArgoHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *CloudArgoNode) SetHealth(v CloudArgoHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *CloudArgoNode) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetImages

`func (o *CloudArgoNode) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CloudArgoNode) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CloudArgoNode) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *CloudArgoNode) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetInfo

`func (o *CloudArgoNode) GetInfo() []CloudArgoInfoItem`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CloudArgoNode) GetInfoOk() (*[]CloudArgoInfoItem, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CloudArgoNode) SetInfo(v []CloudArgoInfoItem)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CloudArgoNode) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetKind

`func (o *CloudArgoNode) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudArgoNode) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudArgoNode) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudArgoNode) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *CloudArgoNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudArgoNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudArgoNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudArgoNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *CloudArgoNode) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CloudArgoNode) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CloudArgoNode) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CloudArgoNode) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetParentRefs

`func (o *CloudArgoNode) GetParentRefs() []CloudArgoResourceRef`

GetParentRefs returns the ParentRefs field if non-nil, zero value otherwise.

### GetParentRefsOk

`func (o *CloudArgoNode) GetParentRefsOk() (*[]CloudArgoResourceRef, bool)`

GetParentRefsOk returns a tuple with the ParentRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRefs

`func (o *CloudArgoNode) SetParentRefs(v []CloudArgoResourceRef)`

SetParentRefs sets ParentRefs field to given value.

### HasParentRefs

`func (o *CloudArgoNode) HasParentRefs() bool`

HasParentRefs returns a boolean if a field has been set.

### GetResourceVersion

`func (o *CloudArgoNode) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *CloudArgoNode) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *CloudArgoNode) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *CloudArgoNode) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetUid

`func (o *CloudArgoNode) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *CloudArgoNode) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *CloudArgoNode) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *CloudArgoNode) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetVersion

`func (o *CloudArgoNode) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CloudArgoNode) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CloudArgoNode) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CloudArgoNode) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


