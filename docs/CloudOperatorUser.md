# CloudOperatorUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Forbidden** | Pointer to **bool** |  | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 
**IsSuperAdmin** | Pointer to **bool** |  | [optional] 
**LastSignin** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudOperatorUser

`func NewCloudOperatorUser() *CloudOperatorUser`

NewCloudOperatorUser instantiates a new CloudOperatorUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudOperatorUserWithDefaults

`func NewCloudOperatorUserWithDefaults() *CloudOperatorUser`

NewCloudOperatorUserWithDefaults instantiates a new CloudOperatorUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CloudOperatorUser) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CloudOperatorUser) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CloudOperatorUser) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CloudOperatorUser) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloudOperatorUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudOperatorUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudOperatorUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudOperatorUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *CloudOperatorUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CloudOperatorUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CloudOperatorUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CloudOperatorUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetForbidden

`func (o *CloudOperatorUser) GetForbidden() bool`

GetForbidden returns the Forbidden field if non-nil, zero value otherwise.

### GetForbiddenOk

`func (o *CloudOperatorUser) GetForbiddenOk() (*bool, bool)`

GetForbiddenOk returns a tuple with the Forbidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidden

`func (o *CloudOperatorUser) SetForbidden(v bool)`

SetForbidden sets Forbidden field to given value.

### HasForbidden

`func (o *CloudOperatorUser) HasForbidden() bool`

HasForbidden returns a boolean if a field has been set.

### GetIsAdmin

`func (o *CloudOperatorUser) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *CloudOperatorUser) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *CloudOperatorUser) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *CloudOperatorUser) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetIsSuperAdmin

`func (o *CloudOperatorUser) GetIsSuperAdmin() bool`

GetIsSuperAdmin returns the IsSuperAdmin field if non-nil, zero value otherwise.

### GetIsSuperAdminOk

`func (o *CloudOperatorUser) GetIsSuperAdminOk() (*bool, bool)`

GetIsSuperAdminOk returns a tuple with the IsSuperAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperAdmin

`func (o *CloudOperatorUser) SetIsSuperAdmin(v bool)`

SetIsSuperAdmin sets IsSuperAdmin field to given value.

### HasIsSuperAdmin

`func (o *CloudOperatorUser) HasIsSuperAdmin() bool`

HasIsSuperAdmin returns a boolean if a field has been set.

### GetLastSignin

`func (o *CloudOperatorUser) GetLastSignin() string`

GetLastSignin returns the LastSignin field if non-nil, zero value otherwise.

### GetLastSigninOk

`func (o *CloudOperatorUser) GetLastSigninOk() (*string, bool)`

GetLastSigninOk returns a tuple with the LastSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSignin

`func (o *CloudOperatorUser) SetLastSignin(v string)`

SetLastSignin sets LastSignin field to given value.

### HasLastSignin

`func (o *CloudOperatorUser) HasLastSignin() bool`

HasLastSignin returns a boolean if a field has been set.

### GetName

`func (o *CloudOperatorUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudOperatorUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudOperatorUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudOperatorUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *CloudOperatorUser) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CloudOperatorUser) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CloudOperatorUser) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CloudOperatorUser) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTag

`func (o *CloudOperatorUser) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CloudOperatorUser) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CloudOperatorUser) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CloudOperatorUser) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


