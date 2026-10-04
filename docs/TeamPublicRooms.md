# TeamPublicRooms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rooms** | Pointer to [**[]TeamListed**](TeamListed.md) | Rooms is every published room the query matched, newest-written first. | [optional] 

## Methods

### NewTeamPublicRooms

`func NewTeamPublicRooms() *TeamPublicRooms`

NewTeamPublicRooms instantiates a new TeamPublicRooms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamPublicRoomsWithDefaults

`func NewTeamPublicRoomsWithDefaults() *TeamPublicRooms`

NewTeamPublicRoomsWithDefaults instantiates a new TeamPublicRooms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRooms

`func (o *TeamPublicRooms) GetRooms() []TeamListed`

GetRooms returns the Rooms field if non-nil, zero value otherwise.

### GetRoomsOk

`func (o *TeamPublicRooms) GetRoomsOk() (*[]TeamListed, bool)`

GetRoomsOk returns a tuple with the Rooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooms

`func (o *TeamPublicRooms) SetRooms(v []TeamListed)`

SetRooms sets Rooms field to given value.

### HasRooms

`func (o *TeamPublicRooms) HasRooms() bool`

HasRooms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


