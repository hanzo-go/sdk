# IamProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Workspace** | Pointer to **string** |  | [optional] 

## Methods

### NewIamProject

`func NewIamProject() *IamProject`

NewIamProject instantiates a new IamProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamProjectWithDefaults

`func NewIamProjectWithDefaults() *IamProject`

NewIamProjectWithDefaults instantiates a new IamProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *IamProject) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamProject) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamProject) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamProject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamProject) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamProject) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamProject) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamProject) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDeleted

`func (o *IamProject) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IamProject) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IamProject) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *IamProject) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDescription

`func (o *IamProject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamProject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamProject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamProject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamProject) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamProject) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamProject) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamProject) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *IamProject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamProject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamProject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamProject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *IamProject) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *IamProject) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *IamProject) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *IamProject) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetMetadata

`func (o *IamProject) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamProject) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamProject) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamProject) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *IamProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *IamProject) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IamProject) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IamProject) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IamProject) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOwner

`func (o *IamProject) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamProject) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamProject) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamProject) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTags

`func (o *IamProject) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IamProject) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IamProject) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IamProject) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IamProject) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamProject) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamProject) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamProject) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWorkspace

`func (o *IamProject) GetWorkspace() string`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *IamProject) GetWorkspaceOk() (*string, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *IamProject) SetWorkspace(v string)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *IamProject) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


