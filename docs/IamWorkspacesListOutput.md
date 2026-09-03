# IamWorkspacesListOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**Workspaces** | Pointer to [**[]IamWorkspace**](IamWorkspace.md) |  | [optional] 

## Methods

### NewIamWorkspacesListOutput

`func NewIamWorkspacesListOutput() *IamWorkspacesListOutput`

NewIamWorkspacesListOutput instantiates a new IamWorkspacesListOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamWorkspacesListOutputWithDefaults

`func NewIamWorkspacesListOutputWithDefaults() *IamWorkspacesListOutput`

NewIamWorkspacesListOutputWithDefaults instantiates a new IamWorkspacesListOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *IamWorkspacesListOutput) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IamWorkspacesListOutput) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IamWorkspacesListOutput) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *IamWorkspacesListOutput) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetWorkspaces

`func (o *IamWorkspacesListOutput) GetWorkspaces() []IamWorkspace`

GetWorkspaces returns the Workspaces field if non-nil, zero value otherwise.

### GetWorkspacesOk

`func (o *IamWorkspacesListOutput) GetWorkspacesOk() (*[]IamWorkspace, bool)`

GetWorkspacesOk returns a tuple with the Workspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaces

`func (o *IamWorkspacesListOutput) SetWorkspaces(v []IamWorkspace)`

SetWorkspaces sets Workspaces field to given value.

### HasWorkspaces

`func (o *IamWorkspacesListOutput) HasWorkspaces() bool`

HasWorkspaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


