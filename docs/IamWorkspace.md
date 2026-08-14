# IamWorkspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** |  | [optional] 
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

## Methods

### NewIamWorkspace

`func NewIamWorkspace() *IamWorkspace`

NewIamWorkspace instantiates a new IamWorkspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamWorkspaceWithDefaults

`func NewIamWorkspaceWithDefaults() *IamWorkspace`

NewIamWorkspaceWithDefaults instantiates a new IamWorkspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *IamWorkspace) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *IamWorkspace) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *IamWorkspace) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *IamWorkspace) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IamWorkspace) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IamWorkspace) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IamWorkspace) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IamWorkspace) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedTime

`func (o *IamWorkspace) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *IamWorkspace) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *IamWorkspace) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *IamWorkspace) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDeleted

`func (o *IamWorkspace) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *IamWorkspace) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *IamWorkspace) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *IamWorkspace) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetDescription

`func (o *IamWorkspace) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IamWorkspace) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IamWorkspace) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IamWorkspace) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *IamWorkspace) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IamWorkspace) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IamWorkspace) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IamWorkspace) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *IamWorkspace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IamWorkspace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IamWorkspace) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IamWorkspace) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *IamWorkspace) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *IamWorkspace) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *IamWorkspace) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *IamWorkspace) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetMetadata

`func (o *IamWorkspace) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *IamWorkspace) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *IamWorkspace) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *IamWorkspace) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *IamWorkspace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamWorkspace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamWorkspace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamWorkspace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *IamWorkspace) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *IamWorkspace) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *IamWorkspace) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *IamWorkspace) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOwner

`func (o *IamWorkspace) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamWorkspace) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamWorkspace) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *IamWorkspace) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTags

`func (o *IamWorkspace) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IamWorkspace) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IamWorkspace) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IamWorkspace) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IamWorkspace) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IamWorkspace) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IamWorkspace) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IamWorkspace) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


