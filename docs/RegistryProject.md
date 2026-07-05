# RegistryProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Public** | Pointer to **bool** | Whether the project is publicly accessible | [optional] 
**OwnerId** | Pointer to **int32** |  | [optional] 
**RepoCount** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to [**RegistryProjectMetadata**](RegistryProjectMetadata.md) |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**UpdateTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRegistryProject

`func NewRegistryProject() *RegistryProject`

NewRegistryProject instantiates a new RegistryProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryProjectWithDefaults

`func NewRegistryProjectWithDefaults() *RegistryProject`

NewRegistryProjectWithDefaults instantiates a new RegistryProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *RegistryProject) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RegistryProject) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RegistryProject) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *RegistryProject) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetName

`func (o *RegistryProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegistryProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegistryProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegistryProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublic

`func (o *RegistryProject) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *RegistryProject) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *RegistryProject) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *RegistryProject) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetOwnerId

`func (o *RegistryProject) GetOwnerId() int32`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *RegistryProject) GetOwnerIdOk() (*int32, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *RegistryProject) SetOwnerId(v int32)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *RegistryProject) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetRepoCount

`func (o *RegistryProject) GetRepoCount() int32`

GetRepoCount returns the RepoCount field if non-nil, zero value otherwise.

### GetRepoCountOk

`func (o *RegistryProject) GetRepoCountOk() (*int32, bool)`

GetRepoCountOk returns a tuple with the RepoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoCount

`func (o *RegistryProject) SetRepoCount(v int32)`

SetRepoCount sets RepoCount field to given value.

### HasRepoCount

`func (o *RegistryProject) HasRepoCount() bool`

HasRepoCount returns a boolean if a field has been set.

### GetMetadata

`func (o *RegistryProject) GetMetadata() RegistryProjectMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RegistryProject) GetMetadataOk() (*RegistryProjectMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RegistryProject) SetMetadata(v RegistryProjectMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RegistryProject) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCreationTime

`func (o *RegistryProject) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RegistryProject) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RegistryProject) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *RegistryProject) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *RegistryProject) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *RegistryProject) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *RegistryProject) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *RegistryProject) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


