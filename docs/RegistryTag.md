# RegistryTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ArtifactId** | Pointer to **int32** |  | [optional] 
**PushTime** | Pointer to **time.Time** |  | [optional] 
**Immutable** | Pointer to **bool** |  | [optional] 

## Methods

### NewRegistryTag

`func NewRegistryTag() *RegistryTag`

NewRegistryTag instantiates a new RegistryTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryTagWithDefaults

`func NewRegistryTagWithDefaults() *RegistryTag`

NewRegistryTagWithDefaults instantiates a new RegistryTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegistryTag) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistryTag) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistryTag) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RegistryTag) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RegistryTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegistryTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegistryTag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegistryTag) HasName() bool`

HasName returns a boolean if a field has been set.

### GetArtifactId

`func (o *RegistryTag) GetArtifactId() int32`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *RegistryTag) GetArtifactIdOk() (*int32, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *RegistryTag) SetArtifactId(v int32)`

SetArtifactId sets ArtifactId field to given value.

### HasArtifactId

`func (o *RegistryTag) HasArtifactId() bool`

HasArtifactId returns a boolean if a field has been set.

### GetPushTime

`func (o *RegistryTag) GetPushTime() time.Time`

GetPushTime returns the PushTime field if non-nil, zero value otherwise.

### GetPushTimeOk

`func (o *RegistryTag) GetPushTimeOk() (*time.Time, bool)`

GetPushTimeOk returns a tuple with the PushTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushTime

`func (o *RegistryTag) SetPushTime(v time.Time)`

SetPushTime sets PushTime field to given value.

### HasPushTime

`func (o *RegistryTag) HasPushTime() bool`

HasPushTime returns a boolean if a field has been set.

### GetImmutable

`func (o *RegistryTag) GetImmutable() bool`

GetImmutable returns the Immutable field if non-nil, zero value otherwise.

### GetImmutableOk

`func (o *RegistryTag) GetImmutableOk() (*bool, bool)`

GetImmutableOk returns a tuple with the Immutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutable

`func (o *RegistryTag) SetImmutable(v bool)`

SetImmutable sets Immutable field to given value.

### HasImmutable

`func (o *RegistryTag) HasImmutable() bool`

HasImmutable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


