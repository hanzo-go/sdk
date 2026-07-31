# CloudPatchIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the repo to update, from the :name path segment. | [optional] 
**Public** | Pointer to **bool** | Public flips anonymous read access. Omit it and the request is refused — there is nothing else to update yet. | [optional] 

## Methods

### NewCloudPatchIn

`func NewCloudPatchIn() *CloudPatchIn`

NewCloudPatchIn instantiates a new CloudPatchIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPatchInWithDefaults

`func NewCloudPatchInWithDefaults() *CloudPatchIn`

NewCloudPatchInWithDefaults instantiates a new CloudPatchIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloudPatchIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudPatchIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudPatchIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudPatchIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublic

`func (o *CloudPatchIn) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *CloudPatchIn) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *CloudPatchIn) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *CloudPatchIn) HasPublic() bool`

HasPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


