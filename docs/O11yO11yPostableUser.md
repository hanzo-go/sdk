# O11yO11yPostableUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | DisplayName is the new member&#39;s display name. | [optional] 
**Email** | Pointer to **string** | Email is the new member&#39;s address. Required. | [optional] 
**FrontendBaseUrl** | Pointer to **string** | FrontendBaseUrl is the console origin the invite link is built on. | [optional] 
**UserRoles** | Pointer to [**[]O11yO11yRoleID**](O11yO11yRoleID.md) | UserRoles are the roles the member starts with, each by id. | [optional] 

## Methods

### NewO11yO11yPostableUser

`func NewO11yO11yPostableUser() *O11yO11yPostableUser`

NewO11yO11yPostableUser instantiates a new O11yO11yPostableUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yPostableUserWithDefaults

`func NewO11yO11yPostableUserWithDefaults() *O11yO11yPostableUser`

NewO11yO11yPostableUserWithDefaults instantiates a new O11yO11yPostableUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *O11yO11yPostableUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *O11yO11yPostableUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *O11yO11yPostableUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *O11yO11yPostableUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *O11yO11yPostableUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *O11yO11yPostableUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *O11yO11yPostableUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *O11yO11yPostableUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFrontendBaseUrl

`func (o *O11yO11yPostableUser) GetFrontendBaseUrl() string`

GetFrontendBaseUrl returns the FrontendBaseUrl field if non-nil, zero value otherwise.

### GetFrontendBaseUrlOk

`func (o *O11yO11yPostableUser) GetFrontendBaseUrlOk() (*string, bool)`

GetFrontendBaseUrlOk returns a tuple with the FrontendBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontendBaseUrl

`func (o *O11yO11yPostableUser) SetFrontendBaseUrl(v string)`

SetFrontendBaseUrl sets FrontendBaseUrl field to given value.

### HasFrontendBaseUrl

`func (o *O11yO11yPostableUser) HasFrontendBaseUrl() bool`

HasFrontendBaseUrl returns a boolean if a field has been set.

### GetUserRoles

`func (o *O11yO11yPostableUser) GetUserRoles() []O11yO11yRoleID`

GetUserRoles returns the UserRoles field if non-nil, zero value otherwise.

### GetUserRolesOk

`func (o *O11yO11yPostableUser) GetUserRolesOk() (*[]O11yO11yRoleID, bool)`

GetUserRolesOk returns a tuple with the UserRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRoles

`func (o *O11yO11yPostableUser) SetUserRoles(v []O11yO11yRoleID)`

SetUserRoles sets UserRoles field to given value.

### HasUserRoles

`func (o *O11yO11yPostableUser) HasUserRoles() bool`

HasUserRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


