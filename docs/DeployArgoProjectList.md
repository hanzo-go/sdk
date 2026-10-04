# DeployArgoProjectList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]DeployArgoProject**](DeployArgoProject.md) | Items is the projects visible to the caller — its own organization&#39;s, or every organization&#39;s for a SuperAdmin. A project named \&quot;default\&quot; is always present and is prepended when IAM does not carry one, because that is what an application with no project label groups under. | [optional] 
**Metadata** | Pointer to [**DeployArgoListMeta**](DeployArgoListMeta.md) | Metadata is the list envelope the SPA expects; it carries no resume point. | [optional] 

## Methods

### NewDeployArgoProjectList

`func NewDeployArgoProjectList() *DeployArgoProjectList`

NewDeployArgoProjectList instantiates a new DeployArgoProjectList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployArgoProjectListWithDefaults

`func NewDeployArgoProjectListWithDefaults() *DeployArgoProjectList`

NewDeployArgoProjectListWithDefaults instantiates a new DeployArgoProjectList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *DeployArgoProjectList) GetItems() []DeployArgoProject`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DeployArgoProjectList) GetItemsOk() (*[]DeployArgoProject, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DeployArgoProjectList) SetItems(v []DeployArgoProject)`

SetItems sets Items field to given value.

### HasItems

`func (o *DeployArgoProjectList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMetadata

`func (o *DeployArgoProjectList) GetMetadata() DeployArgoListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeployArgoProjectList) GetMetadataOk() (*DeployArgoListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeployArgoProjectList) SetMetadata(v DeployArgoListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeployArgoProjectList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


