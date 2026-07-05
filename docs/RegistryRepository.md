# RegistryRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** | Full repository name (project/repo) | [optional] 
**ArtifactCount** | Pointer to **int32** |  | [optional] 
**PullCount** | Pointer to **int32** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**UpdateTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRegistryRepository

`func NewRegistryRepository() *RegistryRepository`

NewRegistryRepository instantiates a new RegistryRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryRepositoryWithDefaults

`func NewRegistryRepositoryWithDefaults() *RegistryRepository`

NewRegistryRepositoryWithDefaults instantiates a new RegistryRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegistryRepository) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistryRepository) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistryRepository) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RegistryRepository) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RegistryRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegistryRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegistryRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegistryRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArtifactCount

`func (o *RegistryRepository) GetArtifactCount() int32`

GetArtifactCount returns the ArtifactCount field if non-nil, zero value otherwise.

### GetArtifactCountOk

`func (o *RegistryRepository) GetArtifactCountOk() (*int32, bool)`

GetArtifactCountOk returns a tuple with the ArtifactCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactCount

`func (o *RegistryRepository) SetArtifactCount(v int32)`

SetArtifactCount sets ArtifactCount field to given value.

### HasArtifactCount

`func (o *RegistryRepository) HasArtifactCount() bool`

HasArtifactCount returns a boolean if a field has been set.

### GetPullCount

`func (o *RegistryRepository) GetPullCount() int32`

GetPullCount returns the PullCount field if non-nil, zero value otherwise.

### GetPullCountOk

`func (o *RegistryRepository) GetPullCountOk() (*int32, bool)`

GetPullCountOk returns a tuple with the PullCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullCount

`func (o *RegistryRepository) SetPullCount(v int32)`

SetPullCount sets PullCount field to given value.

### HasPullCount

`func (o *RegistryRepository) HasPullCount() bool`

HasPullCount returns a boolean if a field has been set.

### GetCreationTime

`func (o *RegistryRepository) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RegistryRepository) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RegistryRepository) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *RegistryRepository) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *RegistryRepository) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *RegistryRepository) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *RegistryRepository) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *RegistryRepository) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


