# PublicRooms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rooms** | Pointer to [**[]Listed**](Listed.md) | Rooms is every published room the query matched, newest-written first. | [optional] 

## Methods

### NewPublicRooms

`func NewPublicRooms() *PublicRooms`

NewPublicRooms instantiates a new PublicRooms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicRoomsWithDefaults

`func NewPublicRoomsWithDefaults() *PublicRooms`

NewPublicRoomsWithDefaults instantiates a new PublicRooms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRooms

`func (o *PublicRooms) GetRooms() []Listed`

GetRooms returns the Rooms field if non-nil, zero value otherwise.

### GetRoomsOk

`func (o *PublicRooms) GetRoomsOk() (*[]Listed, bool)`

GetRoomsOk returns a tuple with the Rooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooms

`func (o *PublicRooms) SetRooms(v []Listed)`

SetRooms sets Rooms field to given value.

### HasRooms

`func (o *PublicRooms) HasRooms() bool`

HasRooms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


