# RegistryUpdateProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Public** | Pointer to **bool** |  | [optional] 
**StorageLimit** | Pointer to **int64** |  | [optional] 
**Metadata** | Pointer to [**RegistryUpdateProjectRequestMetadata**](RegistryUpdateProjectRequestMetadata.md) |  | [optional] 

## Methods

### NewRegistryUpdateProjectRequest

`func NewRegistryUpdateProjectRequest() *RegistryUpdateProjectRequest`

NewRegistryUpdateProjectRequest instantiates a new RegistryUpdateProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryUpdateProjectRequestWithDefaults

`func NewRegistryUpdateProjectRequestWithDefaults() *RegistryUpdateProjectRequest`

NewRegistryUpdateProjectRequestWithDefaults instantiates a new RegistryUpdateProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublic

`func (o *RegistryUpdateProjectRequest) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *RegistryUpdateProjectRequest) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *RegistryUpdateProjectRequest) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *RegistryUpdateProjectRequest) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetStorageLimit

`func (o *RegistryUpdateProjectRequest) GetStorageLimit() int64`

GetStorageLimit returns the StorageLimit field if non-nil, zero value otherwise.

### GetStorageLimitOk

`func (o *RegistryUpdateProjectRequest) GetStorageLimitOk() (*int64, bool)`

GetStorageLimitOk returns a tuple with the StorageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLimit

`func (o *RegistryUpdateProjectRequest) SetStorageLimit(v int64)`

SetStorageLimit sets StorageLimit field to given value.

### HasStorageLimit

`func (o *RegistryUpdateProjectRequest) HasStorageLimit() bool`

HasStorageLimit returns a boolean if a field has been set.

### GetMetadata

`func (o *RegistryUpdateProjectRequest) GetMetadata() RegistryUpdateProjectRequestMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RegistryUpdateProjectRequest) GetMetadataOk() (*RegistryUpdateProjectRequestMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RegistryUpdateProjectRequest) SetMetadata(v RegistryUpdateProjectRequestMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RegistryUpdateProjectRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


