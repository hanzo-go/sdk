# IamControllersCheckPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | **string** |  | 
**Name** | **string** |  | 
**Password** | **string** |  | 
**Ldap** | Pointer to **string** | Non-empty toggles LDAP-password verification mode | [optional] 

## Methods

### NewIamControllersCheckPasswordRequest

`func NewIamControllersCheckPasswordRequest(owner string, name string, password string, ) *IamControllersCheckPasswordRequest`

NewIamControllersCheckPasswordRequest instantiates a new IamControllersCheckPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamControllersCheckPasswordRequestWithDefaults

`func NewIamControllersCheckPasswordRequestWithDefaults() *IamControllersCheckPasswordRequest`

NewIamControllersCheckPasswordRequestWithDefaults instantiates a new IamControllersCheckPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *IamControllersCheckPasswordRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *IamControllersCheckPasswordRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *IamControllersCheckPasswordRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetName

`func (o *IamControllersCheckPasswordRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamControllersCheckPasswordRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamControllersCheckPasswordRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *IamControllersCheckPasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IamControllersCheckPasswordRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IamControllersCheckPasswordRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetLdap

`func (o *IamControllersCheckPasswordRequest) GetLdap() string`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *IamControllersCheckPasswordRequest) GetLdapOk() (*string, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *IamControllersCheckPasswordRequest) SetLdap(v string)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *IamControllersCheckPasswordRequest) HasLdap() bool`

HasLdap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


