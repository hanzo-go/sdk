# CloudRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudRole

`func NewCloudRole() *CloudRole`

NewCloudRole instantiates a new CloudRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRoleWithDefaults

`func NewCloudRoleWithDefaults() *CloudRole`

NewCloudRoleWithDefaults instantiates a new CloudRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *CloudRole) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CloudRole) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CloudRole) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *CloudRole) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUser

`func (o *CloudRole) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CloudRole) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CloudRole) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CloudRole) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


