# ArgoAppList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion is the constant \&quot;argoproj.io/v1alpha1\&quot;. | [optional] 
**Items** | Pointer to [**[]ArgoApp**](ArgoApp.md) | Items is one entry per operator App CR the caller may see — its own org&#39;s, or every platform namespace&#39;s for a SuperAdmin — followed, for a SuperAdmin only, by every Hanzo CD Application in the cluster. Empty (never null) rather than absent when the caller owns nothing. | [optional] 
**Kind** | Pointer to **string** | Kind is the constant \&quot;ApplicationList\&quot;. | [optional] 
**Metadata** | Pointer to [**ArgoListMeta**](ArgoListMeta.md) | Metadata is the list envelope the SPA expects; it carries no resume point. | [optional] 

## Methods

### NewArgoAppList

`func NewArgoAppList() *ArgoAppList`

NewArgoAppList instantiates a new ArgoAppList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgoAppListWithDefaults

`func NewArgoAppListWithDefaults() *ArgoAppList`

NewArgoAppListWithDefaults instantiates a new ArgoAppList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ArgoAppList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ArgoAppList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ArgoAppList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ArgoAppList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *ArgoAppList) GetItems() []ArgoApp`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ArgoAppList) GetItemsOk() (*[]ArgoApp, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ArgoAppList) SetItems(v []ArgoApp)`

SetItems sets Items field to given value.

### HasItems

`func (o *ArgoAppList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *ArgoAppList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ArgoAppList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ArgoAppList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ArgoAppList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *ArgoAppList) GetMetadata() ArgoListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ArgoAppList) GetMetadataOk() (*ArgoListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ArgoAppList) SetMetadata(v ArgoListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ArgoAppList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


