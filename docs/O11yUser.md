# O11yUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **interface{}** |  | [optional] 
**IsRoot** | Pointer to **bool** |  | [optional] 
**OrgId** | Pointer to **interface{}** |  | [optional] 
**Status** | Pointer to **interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewO11yUser

`func NewO11yUser() *O11yUser`

NewO11yUser instantiates a new O11yUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yUserWithDefaults

`func NewO11yUserWithDefaults() *O11yUser`

NewO11yUserWithDefaults instantiates a new O11yUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *O11yUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *O11yUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *O11yUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *O11yUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *O11yUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *O11yUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *O11yUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *O11yUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *O11yUser) GetEmail() interface{}`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *O11yUser) GetEmailOk() (*interface{}, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *O11yUser) SetEmail(v interface{})`

SetEmail sets Email field to given value.

### HasEmail

`func (o *O11yUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *O11yUser) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *O11yUser) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetId

`func (o *O11yUser) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yUser) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yUser) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *O11yUser) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *O11yUser) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *O11yUser) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsRoot

`func (o *O11yUser) GetIsRoot() bool`

GetIsRoot returns the IsRoot field if non-nil, zero value otherwise.

### GetIsRootOk

`func (o *O11yUser) GetIsRootOk() (*bool, bool)`

GetIsRootOk returns a tuple with the IsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoot

`func (o *O11yUser) SetIsRoot(v bool)`

SetIsRoot sets IsRoot field to given value.

### HasIsRoot

`func (o *O11yUser) HasIsRoot() bool`

HasIsRoot returns a boolean if a field has been set.

### GetOrgId

`func (o *O11yUser) GetOrgId() interface{}`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *O11yUser) GetOrgIdOk() (*interface{}, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *O11yUser) SetOrgId(v interface{})`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *O11yUser) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### SetOrgIdNil

`func (o *O11yUser) SetOrgIdNil(b bool)`

 SetOrgIdNil sets the value for OrgId to be an explicit nil

### UnsetOrgId
`func (o *O11yUser) UnsetOrgId()`

UnsetOrgId ensures that no value is present for OrgId, not even an explicit nil
### GetStatus

`func (o *O11yUser) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yUser) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yUser) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yUser) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *O11yUser) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *O11yUser) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUpdatedAt

`func (o *O11yUser) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *O11yUser) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *O11yUser) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *O11yUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


