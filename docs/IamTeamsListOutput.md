# IamTeamsListOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Teams** | Pointer to [**[]IamTeam**](IamTeam.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewIamTeamsListOutput

`func NewIamTeamsListOutput() *IamTeamsListOutput`

NewIamTeamsListOutput instantiates a new IamTeamsListOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamTeamsListOutputWithDefaults

`func NewIamTeamsListOutputWithDefaults() *IamTeamsListOutput`

NewIamTeamsListOutputWithDefaults instantiates a new IamTeamsListOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeams

`func (o *IamTeamsListOutput) GetTeams() []IamTeam`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *IamTeamsListOutput) GetTeamsOk() (*[]IamTeam, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *IamTeamsListOutput) SetTeams(v []IamTeam)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *IamTeamsListOutput) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetTotal

`func (o *IamTeamsListOutput) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IamTeamsListOutput) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IamTeamsListOutput) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *IamTeamsListOutput) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


