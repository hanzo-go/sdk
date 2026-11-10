# TransportTransportAgents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to **map[string]string** | Channels maps a platform channel id to the agent that answers there. | [optional] 
**Default** | Pointer to **string** | Default is the agent that answers any channel without a binding of its own; \&quot;hanzo\&quot; when the org has never set one. | [optional] 
**Transport** | Pointer to **string** | Transport is the transport these bindings are for. | [optional] 

## Methods

### NewTransportTransportAgents

`func NewTransportTransportAgents() *TransportTransportAgents`

NewTransportTransportAgents instantiates a new TransportTransportAgents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportTransportAgentsWithDefaults

`func NewTransportTransportAgentsWithDefaults() *TransportTransportAgents`

NewTransportTransportAgentsWithDefaults instantiates a new TransportTransportAgents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *TransportTransportAgents) GetChannels() map[string]string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *TransportTransportAgents) GetChannelsOk() (*map[string]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *TransportTransportAgents) SetChannels(v map[string]string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *TransportTransportAgents) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetDefault

`func (o *TransportTransportAgents) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *TransportTransportAgents) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *TransportTransportAgents) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *TransportTransportAgents) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetTransport

`func (o *TransportTransportAgents) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *TransportTransportAgents) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *TransportTransportAgents) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *TransportTransportAgents) HasTransport() bool`

HasTransport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


