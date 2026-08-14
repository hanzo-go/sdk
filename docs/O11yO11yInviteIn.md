# O11yO11yInviteIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email is the address the invitation goes to. | [optional] 
**FrontendBaseUrl** | Pointer to **string** | FrontendBaseUrl is the console origin the invite link is built on. | [optional] 
**Name** | Pointer to **string** | Name is the invitee&#39;s display name. | [optional] 
**Role** | Pointer to **string** | Role is the role they will hold on accepting — ADMIN, EDITOR or VIEWER. | [optional] 

## Methods

### NewO11yO11yInviteIn

`func NewO11yO11yInviteIn() *O11yO11yInviteIn`

NewO11yO11yInviteIn instantiates a new O11yO11yInviteIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yInviteInWithDefaults

`func NewO11yO11yInviteInWithDefaults() *O11yO11yInviteIn`

NewO11yO11yInviteInWithDefaults instantiates a new O11yO11yInviteIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *O11yO11yInviteIn) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *O11yO11yInviteIn) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *O11yO11yInviteIn) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *O11yO11yInviteIn) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFrontendBaseUrl

`func (o *O11yO11yInviteIn) GetFrontendBaseUrl() string`

GetFrontendBaseUrl returns the FrontendBaseUrl field if non-nil, zero value otherwise.

### GetFrontendBaseUrlOk

`func (o *O11yO11yInviteIn) GetFrontendBaseUrlOk() (*string, bool)`

GetFrontendBaseUrlOk returns a tuple with the FrontendBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontendBaseUrl

`func (o *O11yO11yInviteIn) SetFrontendBaseUrl(v string)`

SetFrontendBaseUrl sets FrontendBaseUrl field to given value.

### HasFrontendBaseUrl

`func (o *O11yO11yInviteIn) HasFrontendBaseUrl() bool`

HasFrontendBaseUrl returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yInviteIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yInviteIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yInviteIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yInviteIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRole

`func (o *O11yO11yInviteIn) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *O11yO11yInviteIn) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *O11yO11yInviteIn) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *O11yO11yInviteIn) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


