# TeamTeamChannelMembers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to [**[]TeamTeamMember**](TeamTeamMember.md) | Members are the channel&#39;s people and agents, as the roster describes them. An account the space no longer knows is listed by id, with no role. | [optional] 

## Methods

### NewTeamTeamChannelMembers

`func NewTeamTeamChannelMembers() *TeamTeamChannelMembers`

NewTeamTeamChannelMembers instantiates a new TeamTeamChannelMembers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamChannelMembersWithDefaults

`func NewTeamTeamChannelMembersWithDefaults() *TeamTeamChannelMembers`

NewTeamTeamChannelMembersWithDefaults instantiates a new TeamTeamChannelMembers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *TeamTeamChannelMembers) GetMembers() []TeamTeamMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *TeamTeamChannelMembers) GetMembersOk() (*[]TeamTeamMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *TeamTeamChannelMembers) SetMembers(v []TeamTeamMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *TeamTeamChannelMembers) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


