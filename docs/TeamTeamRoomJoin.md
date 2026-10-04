# TeamTeamRoomJoin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the room, from the path. | [optional] 
**Members** | Pointer to **[]string** | Members are the account uuids to add — people or agents of the space. Naming only yourself joins a public channel. | [optional] 
**Space** | Pointer to **string** | Space is the space uuid holding it. Body-only. | [optional] 

## Methods

### NewTeamTeamRoomJoin

`func NewTeamTeamRoomJoin() *TeamTeamRoomJoin`

NewTeamTeamRoomJoin instantiates a new TeamTeamRoomJoin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamRoomJoinWithDefaults

`func NewTeamTeamRoomJoinWithDefaults() *TeamTeamRoomJoin`

NewTeamTeamRoomJoinWithDefaults instantiates a new TeamTeamRoomJoin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamTeamRoomJoin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamTeamRoomJoin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamTeamRoomJoin) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TeamTeamRoomJoin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMembers

`func (o *TeamTeamRoomJoin) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *TeamTeamRoomJoin) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *TeamTeamRoomJoin) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *TeamTeamRoomJoin) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetSpace

`func (o *TeamTeamRoomJoin) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *TeamTeamRoomJoin) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *TeamTeamRoomJoin) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *TeamTeamRoomJoin) HasSpace() bool`

HasSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


