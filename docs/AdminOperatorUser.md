# AdminOperatorUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 
**IsGlobalAdmin** | Pointer to **bool** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**LastSignin** | Pointer to **string** |  | [optional] 
**Forbidden** | Pointer to **bool** |  | [optional] 

## Methods

### NewAdminOperatorUser

`func NewAdminOperatorUser() *AdminOperatorUser`

NewAdminOperatorUser instantiates a new AdminOperatorUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminOperatorUserWithDefaults

`func NewAdminOperatorUserWithDefaults() *AdminOperatorUser`

NewAdminOperatorUserWithDefaults instantiates a new AdminOperatorUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *AdminOperatorUser) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AdminOperatorUser) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AdminOperatorUser) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AdminOperatorUser) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetName

`func (o *AdminOperatorUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminOperatorUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminOperatorUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdminOperatorUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *AdminOperatorUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AdminOperatorUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AdminOperatorUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AdminOperatorUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDisplayName

`func (o *AdminOperatorUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AdminOperatorUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AdminOperatorUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AdminOperatorUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsAdmin

`func (o *AdminOperatorUser) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *AdminOperatorUser) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *AdminOperatorUser) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *AdminOperatorUser) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetIsGlobalAdmin

`func (o *AdminOperatorUser) GetIsGlobalAdmin() bool`

GetIsGlobalAdmin returns the IsGlobalAdmin field if non-nil, zero value otherwise.

### GetIsGlobalAdminOk

`func (o *AdminOperatorUser) GetIsGlobalAdminOk() (*bool, bool)`

GetIsGlobalAdminOk returns a tuple with the IsGlobalAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalAdmin

`func (o *AdminOperatorUser) SetIsGlobalAdmin(v bool)`

SetIsGlobalAdmin sets IsGlobalAdmin field to given value.

### HasIsGlobalAdmin

`func (o *AdminOperatorUser) HasIsGlobalAdmin() bool`

HasIsGlobalAdmin returns a boolean if a field has been set.

### GetTag

`func (o *AdminOperatorUser) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *AdminOperatorUser) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *AdminOperatorUser) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *AdminOperatorUser) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetCreated

`func (o *AdminOperatorUser) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AdminOperatorUser) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AdminOperatorUser) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AdminOperatorUser) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastSignin

`func (o *AdminOperatorUser) GetLastSignin() string`

GetLastSignin returns the LastSignin field if non-nil, zero value otherwise.

### GetLastSigninOk

`func (o *AdminOperatorUser) GetLastSigninOk() (*string, bool)`

GetLastSigninOk returns a tuple with the LastSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSignin

`func (o *AdminOperatorUser) SetLastSignin(v string)`

SetLastSignin sets LastSignin field to given value.

### HasLastSignin

`func (o *AdminOperatorUser) HasLastSignin() bool`

HasLastSignin returns a boolean if a field has been set.

### GetForbidden

`func (o *AdminOperatorUser) GetForbidden() bool`

GetForbidden returns the Forbidden field if non-nil, zero value otherwise.

### GetForbiddenOk

`func (o *AdminOperatorUser) GetForbiddenOk() (*bool, bool)`

GetForbiddenOk returns a tuple with the Forbidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidden

`func (o *AdminOperatorUser) SetForbidden(v bool)`

SetForbidden sets Forbidden field to given value.

### HasForbidden

`func (o *AdminOperatorUser) HasForbidden() bool`

HasForbidden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


