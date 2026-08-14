# PushRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**PushPusher**](PushPusher.md) |  | [optional] 

## Methods

### NewPushRepository

`func NewPushRepository() *PushRepository`

NewPushRepository instantiates a new PushRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushRepositoryWithDefaults

`func NewPushRepositoryWithDefaults() *PushRepository`

NewPushRepositoryWithDefaults instantiates a new PushRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PushRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PushRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PushRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PushRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *PushRepository) GetOwner() PushPusher`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PushRepository) GetOwnerOk() (*PushPusher, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PushRepository) SetOwner(v PushPusher)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PushRepository) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


