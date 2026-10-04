# TeamTeamDirect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **bool** | Created is true when this call opened the conversation (201) and false when it already existed (200). | [optional] 
**Room** | Pointer to [**TeamTeamRoom**](TeamTeamRoom.md) | Room is the direct message, as the room listing answers it. | [optional] 

## Methods

### NewTeamTeamDirect

`func NewTeamTeamDirect() *TeamTeamDirect`

NewTeamTeamDirect instantiates a new TeamTeamDirect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamTeamDirectWithDefaults

`func NewTeamTeamDirectWithDefaults() *TeamTeamDirect`

NewTeamTeamDirectWithDefaults instantiates a new TeamTeamDirect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *TeamTeamDirect) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TeamTeamDirect) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TeamTeamDirect) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *TeamTeamDirect) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetRoom

`func (o *TeamTeamDirect) GetRoom() TeamTeamRoom`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *TeamTeamDirect) GetRoomOk() (*TeamTeamRoom, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *TeamTeamDirect) SetRoom(v TeamTeamRoom)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *TeamTeamDirect) HasRoom() bool`

HasRoom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


