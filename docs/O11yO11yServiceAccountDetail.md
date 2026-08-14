# O11yO11yServiceAccountDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | CreatedAt is when the account was created. | [optional] 
**Email** | Pointer to **string** | Email is the address the account authenticates as. | [optional] 
**Id** | Pointer to **string** | ID is the service account id. | [optional] 
**Name** | Pointer to **string** | Name is the account&#39;s name. | [optional] 
**OrgId** | Pointer to **string** | OrgID is the org the account belongs to. | [optional] 
**ServiceAccountRoles** | Pointer to [**[]O11yO11yServiceAccountRole**](O11yO11yServiceAccountRole.md) | ServiceAccountRoles are the account&#39;s role assignments. | [optional] 
**Status** | Pointer to **string** | Status is active or deleted. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | UpdatedAt is when it last changed. | [optional] 

## Methods

### NewO11yO11yServiceAccountDetail

`func NewO11yO11yServiceAccountDetail() *O11yO11yServiceAccountDetail`

NewO11yO11yServiceAccountDetail instantiates a new O11yO11yServiceAccountDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yServiceAccountDetailWithDefaults

`func NewO11yO11yServiceAccountDetailWithDefaults() *O11yO11yServiceAccountDetail`

NewO11yO11yServiceAccountDetailWithDefaults instantiates a new O11yO11yServiceAccountDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *O11yO11yServiceAccountDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yO11yServiceAccountDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yO11yServiceAccountDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yO11yServiceAccountDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *O11yO11yServiceAccountDetail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *O11yO11yServiceAccountDetail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *O11yO11yServiceAccountDetail) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *O11yO11yServiceAccountDetail) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yServiceAccountDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yServiceAccountDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yServiceAccountDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yServiceAccountDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yServiceAccountDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yServiceAccountDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yServiceAccountDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yServiceAccountDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *O11yO11yServiceAccountDetail) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yO11yServiceAccountDetail) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yO11yServiceAccountDetail) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yO11yServiceAccountDetail) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetServiceAccountRoles

`func (o *O11yO11yServiceAccountDetail) GetServiceAccountRoles() []O11yO11yServiceAccountRole`

GetServiceAccountRoles returns the ServiceAccountRoles field if non-nil, zero value otherwise.

### GetServiceAccountRolesOk

`func (o *O11yO11yServiceAccountDetail) GetServiceAccountRolesOk() (*[]O11yO11yServiceAccountRole, bool)`

GetServiceAccountRolesOk returns a tuple with the ServiceAccountRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountRoles

`func (o *O11yO11yServiceAccountDetail) SetServiceAccountRoles(v []O11yO11yServiceAccountRole)`

SetServiceAccountRoles sets ServiceAccountRoles field to given value.

### HasServiceAccountRoles

`func (o *O11yO11yServiceAccountDetail) HasServiceAccountRoles() bool`

HasServiceAccountRoles returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yServiceAccountDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yServiceAccountDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yServiceAccountDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yServiceAccountDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *O11yO11yServiceAccountDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yO11yServiceAccountDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yO11yServiceAccountDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yO11yServiceAccountDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


