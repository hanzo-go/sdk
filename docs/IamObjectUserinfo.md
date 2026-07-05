# IamObjectUserinfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**Aud** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EmailVerified** | Pointer to **bool** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**IsVerified** | Pointer to **bool** |  | [optional] 
**Iss** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Picture** | Pointer to **string** |  | [optional] 
**PreferredUsername** | Pointer to **string** |  | [optional] 
**RealName** | Pointer to **string** |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**Sub** | Pointer to **string** |  | [optional] 

## Methods

### NewIamObjectUserinfo

`func NewIamObjectUserinfo() *IamObjectUserinfo`

NewIamObjectUserinfo instantiates a new IamObjectUserinfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamObjectUserinfoWithDefaults

`func NewIamObjectUserinfoWithDefaults() *IamObjectUserinfo`

NewIamObjectUserinfoWithDefaults instantiates a new IamObjectUserinfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *IamObjectUserinfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IamObjectUserinfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IamObjectUserinfo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IamObjectUserinfo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAud

`func (o *IamObjectUserinfo) GetAud() string`

GetAud returns the Aud field if non-nil, zero value otherwise.

### GetAudOk

`func (o *IamObjectUserinfo) GetAudOk() (*string, bool)`

GetAudOk returns a tuple with the Aud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAud

`func (o *IamObjectUserinfo) SetAud(v string)`

SetAud sets Aud field to given value.

### HasAud

`func (o *IamObjectUserinfo) HasAud() bool`

HasAud returns a boolean if a field has been set.

### GetEmail

`func (o *IamObjectUserinfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IamObjectUserinfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IamObjectUserinfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IamObjectUserinfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailVerified

`func (o *IamObjectUserinfo) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *IamObjectUserinfo) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *IamObjectUserinfo) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *IamObjectUserinfo) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### GetGroups

`func (o *IamObjectUserinfo) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *IamObjectUserinfo) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *IamObjectUserinfo) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *IamObjectUserinfo) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetIsVerified

`func (o *IamObjectUserinfo) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *IamObjectUserinfo) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *IamObjectUserinfo) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *IamObjectUserinfo) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### GetIss

`func (o *IamObjectUserinfo) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *IamObjectUserinfo) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *IamObjectUserinfo) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *IamObjectUserinfo) HasIss() bool`

HasIss returns a boolean if a field has been set.

### GetName

`func (o *IamObjectUserinfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamObjectUserinfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamObjectUserinfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamObjectUserinfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *IamObjectUserinfo) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *IamObjectUserinfo) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *IamObjectUserinfo) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *IamObjectUserinfo) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPhone

`func (o *IamObjectUserinfo) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *IamObjectUserinfo) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *IamObjectUserinfo) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *IamObjectUserinfo) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPicture

`func (o *IamObjectUserinfo) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *IamObjectUserinfo) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *IamObjectUserinfo) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *IamObjectUserinfo) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetPreferredUsername

`func (o *IamObjectUserinfo) GetPreferredUsername() string`

GetPreferredUsername returns the PreferredUsername field if non-nil, zero value otherwise.

### GetPreferredUsernameOk

`func (o *IamObjectUserinfo) GetPreferredUsernameOk() (*string, bool)`

GetPreferredUsernameOk returns a tuple with the PreferredUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredUsername

`func (o *IamObjectUserinfo) SetPreferredUsername(v string)`

SetPreferredUsername sets PreferredUsername field to given value.

### HasPreferredUsername

`func (o *IamObjectUserinfo) HasPreferredUsername() bool`

HasPreferredUsername returns a boolean if a field has been set.

### GetRealName

`func (o *IamObjectUserinfo) GetRealName() string`

GetRealName returns the RealName field if non-nil, zero value otherwise.

### GetRealNameOk

`func (o *IamObjectUserinfo) GetRealNameOk() (*string, bool)`

GetRealNameOk returns a tuple with the RealName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealName

`func (o *IamObjectUserinfo) SetRealName(v string)`

SetRealName sets RealName field to given value.

### HasRealName

`func (o *IamObjectUserinfo) HasRealName() bool`

HasRealName returns a boolean if a field has been set.

### GetRoles

`func (o *IamObjectUserinfo) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IamObjectUserinfo) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IamObjectUserinfo) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IamObjectUserinfo) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSub

`func (o *IamObjectUserinfo) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *IamObjectUserinfo) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *IamObjectUserinfo) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *IamObjectUserinfo) HasSub() bool`

HasSub returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


