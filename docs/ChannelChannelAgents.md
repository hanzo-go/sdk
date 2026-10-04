# ChannelChannelAgents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is the transport these bindings are for. | [optional] 
**Default** | Pointer to **string** | Default is the agent that answers any room without a binding of its own; \&quot;hanzo\&quot; when the org has never set one. | [optional] 
**Rooms** | Pointer to **map[string]string** | Rooms maps a platform room id to the agent that answers there. | [optional] 

## Methods

### NewChannelChannelAgents

`func NewChannelChannelAgents() *ChannelChannelAgents`

NewChannelChannelAgents instantiates a new ChannelChannelAgents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelChannelAgentsWithDefaults

`func NewChannelChannelAgentsWithDefaults() *ChannelChannelAgents`

NewChannelChannelAgentsWithDefaults instantiates a new ChannelChannelAgents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ChannelChannelAgents) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ChannelChannelAgents) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ChannelChannelAgents) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ChannelChannelAgents) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDefault

`func (o *ChannelChannelAgents) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ChannelChannelAgents) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ChannelChannelAgents) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ChannelChannelAgents) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetRooms

`func (o *ChannelChannelAgents) GetRooms() map[string]string`

GetRooms returns the Rooms field if non-nil, zero value otherwise.

### GetRoomsOk

`func (o *ChannelChannelAgents) GetRoomsOk() (*map[string]string, bool)`

GetRoomsOk returns a tuple with the Rooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooms

`func (o *ChannelChannelAgents) SetRooms(v map[string]string)`

SetRooms sets Rooms field to given value.

### HasRooms

`func (o *ChannelChannelAgents) HasRooms() bool`

HasRooms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


