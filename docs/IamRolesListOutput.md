# IamRolesListOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to [**[]IamRole**](IamRole.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewIamRolesListOutput

`func NewIamRolesListOutput() *IamRolesListOutput`

NewIamRolesListOutput instantiates a new IamRolesListOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamRolesListOutputWithDefaults

`func NewIamRolesListOutputWithDefaults() *IamRolesListOutput`

NewIamRolesListOutputWithDefaults instantiates a new IamRolesListOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *IamRolesListOutput) GetRoles() []IamRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *IamRolesListOutput) GetRolesOk() (*[]IamRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *IamRolesListOutput) SetRoles(v []IamRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *IamRolesListOutput) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetTotal

`func (o *IamRolesListOutput) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IamRolesListOutput) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IamRolesListOutput) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *IamRolesListOutput) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


