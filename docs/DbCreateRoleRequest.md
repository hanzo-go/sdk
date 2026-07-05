# DbCreateRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | [**DbRoleCreate**](DbRoleCreate.md) |  | 

## Methods

### NewDbCreateRoleRequest

`func NewDbCreateRoleRequest(role DbRoleCreate, ) *DbCreateRoleRequest`

NewDbCreateRoleRequest instantiates a new DbCreateRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbCreateRoleRequestWithDefaults

`func NewDbCreateRoleRequestWithDefaults() *DbCreateRoleRequest`

NewDbCreateRoleRequestWithDefaults instantiates a new DbCreateRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *DbCreateRoleRequest) GetRole() DbRoleCreate`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DbCreateRoleRequest) GetRoleOk() (*DbRoleCreate, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DbCreateRoleRequest) SetRole(v DbRoleCreate)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


