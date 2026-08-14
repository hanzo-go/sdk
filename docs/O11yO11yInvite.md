# O11yO11yInvite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when it was created. | [optional] 
**Email** | Pointer to **string** | Email is the address it was sent to. | [optional] 
**Id** | Pointer to **string** | ID is the invitation&#39;s id. | [optional] 
**InviteLink** | Pointer to **string** | InviteLink is the full link mailed to them. | [optional] 
**Name** | Pointer to **string** | Name is the invitee&#39;s display name. | [optional] 
**OrgId** | Pointer to **string** | OrgID is the org they are invited into. | [optional] 
**Role** | Pointer to **string** | Role is the role the invitee will hold. | [optional] 
**Token** | Pointer to **string** | Token is the secret that redeems it. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when it last changed. | [optional] 

## Methods

### NewO11yO11yInvite

`func NewO11yO11yInvite() *O11yO11yInvite`

NewO11yO11yInvite instantiates a new O11yO11yInvite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yInviteWithDefaults

`func NewO11yO11yInviteWithDefaults() *O11yO11yInvite`

NewO11yO11yInviteWithDefaults instantiates a new O11yO11yInvite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *O11yO11yInvite) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yInvite) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yInvite) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yInvite) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *O11yO11yInvite) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *O11yO11yInvite) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *O11yO11yInvite) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *O11yO11yInvite) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yInvite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yInvite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yInvite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yInvite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInviteLink

`func (o *O11yO11yInvite) GetInviteLink() string`

GetInviteLink returns the InviteLink field if non-nil, zero value otherwise.

### GetInviteLinkOk

`func (o *O11yO11yInvite) GetInviteLinkOk() (*string, bool)`

GetInviteLinkOk returns a tuple with the InviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteLink

`func (o *O11yO11yInvite) SetInviteLink(v string)`

SetInviteLink sets InviteLink field to given value.

### HasInviteLink

`func (o *O11yO11yInvite) HasInviteLink() bool`

HasInviteLink returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yInvite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yInvite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yInvite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yInvite) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *O11yO11yInvite) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yO11yInvite) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yO11yInvite) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yO11yInvite) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetRole

`func (o *O11yO11yInvite) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *O11yO11yInvite) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *O11yO11yInvite) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *O11yO11yInvite) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetToken

`func (o *O11yO11yInvite) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *O11yO11yInvite) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *O11yO11yInvite) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *O11yO11yInvite) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yInvite) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yInvite) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yInvite) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yInvite) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


