# RegistryProjectCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectName** | **string** |  | 
**Public** | Pointer to **bool** |  | [optional] [default to false]
**StorageLimit** | Pointer to **int64** | Storage quota in bytes (-1 for unlimited) | [optional] 
**Metadata** | Pointer to [**RegistryProjectCreateMetadata**](RegistryProjectCreateMetadata.md) |  | [optional] 

## Methods

### NewRegistryProjectCreate

`func NewRegistryProjectCreate(projectName string, ) *RegistryProjectCreate`

NewRegistryProjectCreate instantiates a new RegistryProjectCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryProjectCreateWithDefaults

`func NewRegistryProjectCreateWithDefaults() *RegistryProjectCreate`

NewRegistryProjectCreateWithDefaults instantiates a new RegistryProjectCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectName

`func (o *RegistryProjectCreate) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *RegistryProjectCreate) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *RegistryProjectCreate) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetPublic

`func (o *RegistryProjectCreate) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *RegistryProjectCreate) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *RegistryProjectCreate) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *RegistryProjectCreate) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetStorageLimit

`func (o *RegistryProjectCreate) GetStorageLimit() int64`

GetStorageLimit returns the StorageLimit field if non-nil, zero value otherwise.

### GetStorageLimitOk

`func (o *RegistryProjectCreate) GetStorageLimitOk() (*int64, bool)`

GetStorageLimitOk returns a tuple with the StorageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLimit

`func (o *RegistryProjectCreate) SetStorageLimit(v int64)`

SetStorageLimit sets StorageLimit field to given value.

### HasStorageLimit

`func (o *RegistryProjectCreate) HasStorageLimit() bool`

HasStorageLimit returns a boolean if a field has been set.

### GetMetadata

`func (o *RegistryProjectCreate) GetMetadata() RegistryProjectCreateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RegistryProjectCreate) GetMetadataOk() (*RegistryProjectCreateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RegistryProjectCreate) SetMetadata(v RegistryProjectCreateMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RegistryProjectCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


