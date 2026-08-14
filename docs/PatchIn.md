# PatchIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the repo to update, from the :name path segment. | [optional] 
**Public** | Pointer to **bool** | Public flips anonymous read access. Omit it and the request is refused — there is nothing else to update yet. | [optional] 

## Methods

### NewPatchIn

`func NewPatchIn() *PatchIn`

NewPatchIn instantiates a new PatchIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchInWithDefaults

`func NewPatchInWithDefaults() *PatchIn`

NewPatchInWithDefaults instantiates a new PatchIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublic

`func (o *PatchIn) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *PatchIn) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *PatchIn) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *PatchIn) HasPublic() bool`

HasPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


