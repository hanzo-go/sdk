# ChannelChannelAgentsPut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is the transport to edit. Required; an unknown value is a 404. | [optional] 
**Default** | Pointer to **string** | Default sets the agent for rooms with no binding of their own; \&quot;hanzo\&quot; restores the built-in. Empty or absent leaves it unchanged. | [optional] 
**Rooms** | Pointer to **map[string]string** | Rooms binds platform room ids to agents; rooms not named are left alone. | [optional] 
**Unbind** | Pointer to **[]string** | Unbind removes the bindings of these rooms, so they fall back to Default. | [optional] 

## Methods

### NewChannelChannelAgentsPut

`func NewChannelChannelAgentsPut() *ChannelChannelAgentsPut`

NewChannelChannelAgentsPut instantiates a new ChannelChannelAgentsPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelChannelAgentsPutWithDefaults

`func NewChannelChannelAgentsPutWithDefaults() *ChannelChannelAgentsPut`

NewChannelChannelAgentsPutWithDefaults instantiates a new ChannelChannelAgentsPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ChannelChannelAgentsPut) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ChannelChannelAgentsPut) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ChannelChannelAgentsPut) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ChannelChannelAgentsPut) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDefault

`func (o *ChannelChannelAgentsPut) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ChannelChannelAgentsPut) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ChannelChannelAgentsPut) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ChannelChannelAgentsPut) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetRooms

`func (o *ChannelChannelAgentsPut) GetRooms() map[string]string`

GetRooms returns the Rooms field if non-nil, zero value otherwise.

### GetRoomsOk

`func (o *ChannelChannelAgentsPut) GetRoomsOk() (*map[string]string, bool)`

GetRoomsOk returns a tuple with the Rooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooms

`func (o *ChannelChannelAgentsPut) SetRooms(v map[string]string)`

SetRooms sets Rooms field to given value.

### HasRooms

`func (o *ChannelChannelAgentsPut) HasRooms() bool`

HasRooms returns a boolean if a field has been set.

### GetUnbind

`func (o *ChannelChannelAgentsPut) GetUnbind() []string`

GetUnbind returns the Unbind field if non-nil, zero value otherwise.

### GetUnbindOk

`func (o *ChannelChannelAgentsPut) GetUnbindOk() (*[]string, bool)`

GetUnbindOk returns a tuple with the Unbind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbind

`func (o *ChannelChannelAgentsPut) SetUnbind(v []string)`

SetUnbind sets Unbind field to given value.

### HasUnbind

`func (o *ChannelChannelAgentsPut) HasUnbind() bool`

HasUnbind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


