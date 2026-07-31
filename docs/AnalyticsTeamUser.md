# AnalyticsTeamUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**TeamId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**User** | Pointer to [**AnalyticsUser**](AnalyticsUser.md) |  | [optional] 

## Methods

### NewAnalyticsTeamUser

`func NewAnalyticsTeamUser() *AnalyticsTeamUser`

NewAnalyticsTeamUser instantiates a new AnalyticsTeamUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsTeamUserWithDefaults

`func NewAnalyticsTeamUserWithDefaults() *AnalyticsTeamUser`

NewAnalyticsTeamUserWithDefaults instantiates a new AnalyticsTeamUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AnalyticsTeamUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnalyticsTeamUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnalyticsTeamUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AnalyticsTeamUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTeamId

`func (o *AnalyticsTeamUser) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *AnalyticsTeamUser) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *AnalyticsTeamUser) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *AnalyticsTeamUser) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetUserId

`func (o *AnalyticsTeamUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AnalyticsTeamUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AnalyticsTeamUser) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AnalyticsTeamUser) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetRole

`func (o *AnalyticsTeamUser) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AnalyticsTeamUser) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AnalyticsTeamUser) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AnalyticsTeamUser) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AnalyticsTeamUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AnalyticsTeamUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AnalyticsTeamUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AnalyticsTeamUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AnalyticsTeamUser) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AnalyticsTeamUser) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AnalyticsTeamUser) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AnalyticsTeamUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *AnalyticsTeamUser) GetUser() AnalyticsUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AnalyticsTeamUser) GetUserOk() (*AnalyticsUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AnalyticsTeamUser) SetUser(v AnalyticsUser)`

SetUser sets User field to given value.

### HasUser

`func (o *AnalyticsTeamUser) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


