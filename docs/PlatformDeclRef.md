# PlatformDeclRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the declaration&#39;s name, the file&#39;s basename. | [optional] 
**Org** | Pointer to **string** | Org is the values directory the declaration lives in. | [optional] 

## Methods

### NewPlatformDeclRef

`func NewPlatformDeclRef() *PlatformDeclRef`

NewPlatformDeclRef instantiates a new PlatformDeclRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformDeclRefWithDefaults

`func NewPlatformDeclRefWithDefaults() *PlatformDeclRef`

NewPlatformDeclRefWithDefaults instantiates a new PlatformDeclRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlatformDeclRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformDeclRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformDeclRef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PlatformDeclRef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *PlatformDeclRef) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PlatformDeclRef) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PlatformDeclRef) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PlatformDeclRef) HasOrg() bool`

HasOrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


