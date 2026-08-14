# IamProjectsListOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | Pointer to [**[]IamProject**](IamProject.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewIamProjectsListOutput

`func NewIamProjectsListOutput() *IamProjectsListOutput`

NewIamProjectsListOutput instantiates a new IamProjectsListOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamProjectsListOutputWithDefaults

`func NewIamProjectsListOutputWithDefaults() *IamProjectsListOutput`

NewIamProjectsListOutputWithDefaults instantiates a new IamProjectsListOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjects

`func (o *IamProjectsListOutput) GetProjects() []IamProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *IamProjectsListOutput) GetProjectsOk() (*[]IamProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *IamProjectsListOutput) SetProjects(v []IamProject)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *IamProjectsListOutput) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetTotal

`func (o *IamProjectsListOutput) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IamProjectsListOutput) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IamProjectsListOutput) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *IamProjectsListOutput) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


