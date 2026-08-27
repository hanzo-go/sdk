# RoomWork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Open** | Pointer to **int32** | Open is how many items are still work: everything whose status does not end it. It is the number a channel header shows. | [optional] 
**Room** | Pointer to **string** | Room is the room these counts are for, echoed back as it was resolved. | [optional] 
**Status** | Pointer to **map[string]int32** | Status is the count per board column, carrying EVERY column this surface knows — an empty column reads 0 rather than being absent, so a caller can render the board without inventing the vocabulary. The keys are the same closed set every other operation here validates against. | [optional] 
**Total** | Pointer to **int32** | Total is every item bound to this room, settled ones included, so Total minus Open is what the room has finished. | [optional] 
**Updated** | Pointer to **int32** | Updated is when anything in this room&#39;s work last moved, in unix seconds. ABSENT when the room has no work at all: zero would read as the epoch, and a room nobody has filed anything in has no last activity rather than an infinitely old one. Total is 0 in exactly that case. | [optional] 

## Methods

### NewRoomWork

`func NewRoomWork() *RoomWork`

NewRoomWork instantiates a new RoomWork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomWorkWithDefaults

`func NewRoomWorkWithDefaults() *RoomWork`

NewRoomWorkWithDefaults instantiates a new RoomWork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpen

`func (o *RoomWork) GetOpen() int32`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *RoomWork) GetOpenOk() (*int32, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *RoomWork) SetOpen(v int32)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *RoomWork) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetRoom

`func (o *RoomWork) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *RoomWork) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *RoomWork) SetRoom(v string)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *RoomWork) HasRoom() bool`

HasRoom returns a boolean if a field has been set.

### GetStatus

`func (o *RoomWork) GetStatus() map[string]int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RoomWork) GetStatusOk() (*map[string]int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RoomWork) SetStatus(v map[string]int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RoomWork) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotal

`func (o *RoomWork) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *RoomWork) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *RoomWork) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *RoomWork) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUpdated

`func (o *RoomWork) GetUpdated() int32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *RoomWork) GetUpdatedOk() (*int32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *RoomWork) SetUpdated(v int32)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *RoomWork) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


