# ChannelAgents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is the transport these bindings are for. | [optional] 
**Default** | Pointer to **string** | Default is the agent that answers any room without a binding of its own; \&quot;hanzo\&quot; when the org has never set one. | [optional] 
**Rooms** | Pointer to **map[string]string** | Rooms maps a platform room id to the agent that answers there. | [optional] 

## Methods

### NewChannelAgents

`func NewChannelAgents() *ChannelAgents`

NewChannelAgents instantiates a new ChannelAgents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelAgentsWithDefaults

`func NewChannelAgentsWithDefaults() *ChannelAgents`

NewChannelAgentsWithDefaults instantiates a new ChannelAgents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ChannelAgents) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ChannelAgents) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ChannelAgents) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ChannelAgents) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDefault

`func (o *ChannelAgents) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ChannelAgents) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ChannelAgents) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ChannelAgents) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetRooms

`func (o *ChannelAgents) GetRooms() map[string]string`

GetRooms returns the Rooms field if non-nil, zero value otherwise.

### GetRoomsOk

`func (o *ChannelAgents) GetRoomsOk() (*map[string]string, bool)`

GetRoomsOk returns a tuple with the Rooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooms

`func (o *ChannelAgents) SetRooms(v map[string]string)`

SetRooms sets Rooms field to given value.

### HasRooms

`func (o *ChannelAgents) HasRooms() bool`

HasRooms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


