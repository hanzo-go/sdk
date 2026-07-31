# CloudArgoProjectList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]CloudArgoProject**](CloudArgoProject.md) |  | [optional] 
**Metadata** | Pointer to [**CloudArgoListMeta**](CloudArgoListMeta.md) |  | [optional] 

## Methods

### NewCloudArgoProjectList

`func NewCloudArgoProjectList() *CloudArgoProjectList`

NewCloudArgoProjectList instantiates a new CloudArgoProjectList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudArgoProjectListWithDefaults

`func NewCloudArgoProjectListWithDefaults() *CloudArgoProjectList`

NewCloudArgoProjectListWithDefaults instantiates a new CloudArgoProjectList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CloudArgoProjectList) GetItems() []CloudArgoProject`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CloudArgoProjectList) GetItemsOk() (*[]CloudArgoProject, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CloudArgoProjectList) SetItems(v []CloudArgoProject)`

SetItems sets Items field to given value.

### HasItems

`func (o *CloudArgoProjectList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMetadata

`func (o *CloudArgoProjectList) GetMetadata() CloudArgoListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CloudArgoProjectList) GetMetadataOk() (*CloudArgoListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CloudArgoProjectList) SetMetadata(v CloudArgoListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CloudArgoProjectList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


