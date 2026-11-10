# TeamTeamChannelJoin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the channel, from the path. | [optional] 
**Members** | Pointer to **[]string** | Members are the account uuids to add — people or agents of the space. Naming only yourself joins a public channel. | [optional] 
**Space** | Pointer to **string** | Space is the space uuid holding it. Body-only. | [optional] 

## Methods

### NewTeamTeamChannelJoin

`func NewTeamTeamChannelJoin() *TeamTeamChannelJoin`

NewTeamTeamChannelJoin instantiates a new TeamTeamChannelJoin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamChannelJoinWithDefaults

`func NewTeamTeamChannelJoinWithDefaults() *TeamTeamChannelJoin`

NewTeamTeamChannelJoinWithDefaults instantiates a new TeamTeamChannelJoin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamTeamChannelJoin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamTeamChannelJoin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamTeamChannelJoin) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamTeamChannelJoin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMembers

`func (o *TeamTeamChannelJoin) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *TeamTeamChannelJoin) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *TeamTeamChannelJoin) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *TeamTeamChannelJoin) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetSpace

`func (o *TeamTeamChannelJoin) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *TeamTeamChannelJoin) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *TeamTeamChannelJoin) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *TeamTeamChannelJoin) HasSpace() bool`

HasSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


