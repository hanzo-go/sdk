# TeamTeamDirectOpen

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to **[]string** | Members are the account uuids to talk to — people or agents of the space. The caller is always in the conversation and need not name themselves; an empty list is the caller&#39;s own notes-to-self conversation. | [optional] 
**Space** | Pointer to **string** | Space is the space uuid. Optional for a caller in exactly one space. | [optional] 

## Methods

### NewTeamTeamDirectOpen

`func NewTeamTeamDirectOpen() *TeamTeamDirectOpen`

NewTeamTeamDirectOpen instantiates a new TeamTeamDirectOpen object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamDirectOpenWithDefaults

`func NewTeamTeamDirectOpenWithDefaults() *TeamTeamDirectOpen`

NewTeamTeamDirectOpenWithDefaults instantiates a new TeamTeamDirectOpen object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *TeamTeamDirectOpen) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *TeamTeamDirectOpen) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *TeamTeamDirectOpen) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *TeamTeamDirectOpen) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetSpace

`func (o *TeamTeamDirectOpen) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *TeamTeamDirectOpen) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *TeamTeamDirectOpen) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *TeamTeamDirectOpen) HasSpace() bool`

HasSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


