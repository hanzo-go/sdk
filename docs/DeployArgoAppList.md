# DeployArgoAppList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion is the constant \&quot;argoproj.io/v1alpha1\&quot;. | [optional] 
**Items** | Pointer to [**[]DeployArgoApp**](DeployArgoApp.md) | Items is one entry per operator App CR the caller may see — its own org&#39;s, or every platform namespace&#39;s for a SuperAdmin — followed, for a SuperAdmin only, by every Hanzo CD Application in the cluster. Empty (never null) rather than absent when the caller owns nothing. | [optional] 
**Kind** | Pointer to **string** | Kind is the constant \&quot;ApplicationList\&quot;. | [optional] 
**Metadata** | Pointer to [**DeployArgoListMeta**](DeployArgoListMeta.md) | Metadata is the list envelope the SPA expects; it carries no resume point. | [optional] 

## Methods

### NewDeployArgoAppList

`func NewDeployArgoAppList() *DeployArgoAppList`

NewDeployArgoAppList instantiates a new DeployArgoAppList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployArgoAppListWithDefaults

`func NewDeployArgoAppListWithDefaults() *DeployArgoAppList`

NewDeployArgoAppListWithDefaults instantiates a new DeployArgoAppList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *DeployArgoAppList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *DeployArgoAppList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *DeployArgoAppList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *DeployArgoAppList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *DeployArgoAppList) GetItems() []DeployArgoApp`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DeployArgoAppList) GetItemsOk() (*[]DeployArgoApp, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DeployArgoAppList) SetItems(v []DeployArgoApp)`

SetItems sets Items field to given value.

### HasItems

`func (o *DeployArgoAppList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *DeployArgoAppList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DeployArgoAppList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DeployArgoAppList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DeployArgoAppList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *DeployArgoAppList) GetMetadata() DeployArgoListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeployArgoAppList) GetMetadataOk() (*DeployArgoListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeployArgoAppList) SetMetadata(v DeployArgoListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeployArgoAppList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


