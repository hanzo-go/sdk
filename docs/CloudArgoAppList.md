# CloudArgoAppList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]CloudArgoApp**](CloudArgoApp.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**CloudArgoListMeta**](CloudArgoListMeta.md) |  | [optional] 

## Methods

### NewCloudArgoAppList

`func NewCloudArgoAppList() *CloudArgoAppList`

NewCloudArgoAppList instantiates a new CloudArgoAppList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudArgoAppListWithDefaults

`func NewCloudArgoAppListWithDefaults() *CloudArgoAppList`

NewCloudArgoAppListWithDefaults instantiates a new CloudArgoAppList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *CloudArgoAppList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *CloudArgoAppList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *CloudArgoAppList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *CloudArgoAppList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *CloudArgoAppList) GetItems() []CloudArgoApp`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CloudArgoAppList) GetItemsOk() (*[]CloudArgoApp, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CloudArgoAppList) SetItems(v []CloudArgoApp)`

SetItems sets Items field to given value.

### HasItems

`func (o *CloudArgoAppList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *CloudArgoAppList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CloudArgoAppList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CloudArgoAppList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *CloudArgoAppList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *CloudArgoAppList) GetMetadata() CloudArgoListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CloudArgoAppList) GetMetadataOk() (*CloudArgoListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CloudArgoAppList) SetMetadata(v CloudArgoListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CloudArgoAppList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


