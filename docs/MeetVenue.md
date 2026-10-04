# MeetVenue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the media room to join: the value POST /v1/meet/getToken takes as roomName, and the value the media server keys participants on. | [optional] 
**Ready** | Pointer to **bool** | Ready reports that this deployment can mint a join token for this room. It is false on a deployment holding no media-server key, where Name is still correct — the name is a property of the room and the key is a property of the deployment, so a caller learns the room&#39;s identity either way and learns not to offer a join button. | [optional] 
**Ws** | Pointer to **string** | WS is where the media plane is — the address a client opens its own browser-to-server connection to. Empty when this deployment has not been told where its media server lives, which is reported rather than refused: a surface can say a call is unavailable without a second request. | [optional] 

## Methods

### NewMeetVenue

`func NewMeetVenue() *MeetVenue`

NewMeetVenue instantiates a new MeetVenue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetVenueWithDefaults

`func NewMeetVenueWithDefaults() *MeetVenue`

NewMeetVenueWithDefaults instantiates a new MeetVenue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MeetVenue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MeetVenue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MeetVenue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MeetVenue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReady

`func (o *MeetVenue) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *MeetVenue) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *MeetVenue) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *MeetVenue) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetWs

`func (o *MeetVenue) GetWs() string`

GetWs returns the Ws field if non-nil, zero value otherwise.

### GetWsOk

`func (o *MeetVenue) GetWsOk() (*string, bool)`

GetWsOk returns a tuple with the Ws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWs

`func (o *MeetVenue) SetWs(v string)`

SetWs sets Ws field to given value.

### HasWs

`func (o *MeetVenue) HasWs() bool`

HasWs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


