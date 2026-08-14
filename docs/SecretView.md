# SecretView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MountedBy** | Pointer to **string** | MountedBy is a function that mounts it. | [optional] 
**Name** | Pointer to **string** | Name is the environment variable the function mounts. | [optional] 
**Namespace** | Pointer to **string** | Namespace is the group the mounting function belongs to. | [optional] 

## Methods

### NewSecretView

`func NewSecretView() *SecretView`

NewSecretView instantiates a new SecretView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretViewWithDefaults

`func NewSecretViewWithDefaults() *SecretView`

NewSecretViewWithDefaults instantiates a new SecretView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountedBy

`func (o *SecretView) GetMountedBy() string`

GetMountedBy returns the MountedBy field if non-nil, zero value otherwise.

### GetMountedByOk

`func (o *SecretView) GetMountedByOk() (*string, bool)`

GetMountedByOk returns a tuple with the MountedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountedBy

`func (o *SecretView) SetMountedBy(v string)`

SetMountedBy sets MountedBy field to given value.

### HasMountedBy

`func (o *SecretView) HasMountedBy() bool`

HasMountedBy returns a boolean if a field has been set.

### GetName

`func (o *SecretView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecretView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecretView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecretView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *SecretView) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SecretView) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SecretView) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SecretView) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


