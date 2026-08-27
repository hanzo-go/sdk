# TeamRooms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rooms** | Pointer to [**[]TeamRoom**](TeamRoom.md) | Rooms is every room of every workspace the caller&#39;s org owns, each with the work facet it carries. | [optional] 

## Methods

### NewTeamRooms

`func NewTeamRooms() *TeamRooms`

NewTeamRooms instantiates a new TeamRooms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamRoomsWithDefaults

`func NewTeamRoomsWithDefaults() *TeamRooms`

NewTeamRoomsWithDefaults instantiates a new TeamRooms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRooms

`func (o *TeamRooms) GetRooms() []TeamRoom`

GetRooms returns the Rooms field if non-nil, zero value otherwise.

### GetRoomsOk

`func (o *TeamRooms) GetRoomsOk() (*[]TeamRoom, bool)`

GetRoomsOk returns a tuple with the Rooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooms

`func (o *TeamRooms) SetRooms(v []TeamRoom)`

SetRooms sets Rooms field to given value.

### HasRooms

`func (o *TeamRooms) HasRooms() bool`

HasRooms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


