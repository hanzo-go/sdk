# ChannelAgentsPut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is the transport to edit. Required; an unknown value is a 404. | [optional] 
**Default** | Pointer to **string** | Default sets the agent for rooms with no binding of their own; \&quot;hanzo\&quot; restores the built-in. Empty or absent leaves it unchanged. | [optional] 
**Rooms** | Pointer to **map[string]string** | Rooms binds platform room ids to agents; rooms not named are left alone. | [optional] 
**Unbind** | Pointer to **[]string** | Unbind removes the bindings of these rooms, so they fall back to Default. | [optional] 

## Methods

### NewChannelAgentsPut

`func NewChannelAgentsPut() *ChannelAgentsPut`

NewChannelAgentsPut instantiates a new ChannelAgentsPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelAgentsPutWithDefaults

`func NewChannelAgentsPutWithDefaults() *ChannelAgentsPut`

NewChannelAgentsPutWithDefaults instantiates a new ChannelAgentsPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ChannelAgentsPut) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ChannelAgentsPut) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ChannelAgentsPut) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ChannelAgentsPut) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDefault

`func (o *ChannelAgentsPut) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ChannelAgentsPut) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ChannelAgentsPut) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ChannelAgentsPut) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetRooms

`func (o *ChannelAgentsPut) GetRooms() map[string]string`

GetRooms returns the Rooms field if non-nil, zero value otherwise.

### GetRoomsOk

`func (o *ChannelAgentsPut) GetRoomsOk() (*map[string]string, bool)`

GetRoomsOk returns a tuple with the Rooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooms

`func (o *ChannelAgentsPut) SetRooms(v map[string]string)`

SetRooms sets Rooms field to given value.

### HasRooms

`func (o *ChannelAgentsPut) HasRooms() bool`

HasRooms returns a boolean if a field has been set.

### GetUnbind

`func (o *ChannelAgentsPut) GetUnbind() []string`

GetUnbind returns the Unbind field if non-nil, zero value otherwise.

### GetUnbindOk

`func (o *ChannelAgentsPut) GetUnbindOk() (*[]string, bool)`

GetUnbindOk returns a tuple with the Unbind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbind

`func (o *ChannelAgentsPut) SetUnbind(v []string)`

SetUnbind sets Unbind field to given value.

### HasUnbind

`func (o *ChannelAgentsPut) HasUnbind() bool`

HasUnbind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


