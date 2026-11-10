# TransportTransportAgentsPut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to **map[string]string** | Channels binds platform channel ids to agents; channels not named are left alone. | [optional] 
**Default** | Pointer to **string** | Default sets the agent for channels with no binding of their own; \&quot;hanzo\&quot; restores the built-in. Empty or absent leaves it unchanged. | [optional] 
**Transport** | Pointer to **string** | Transport is the transport to edit. Required; an unknown value is a 404. | [optional] 
**Unbind** | Pointer to **[]string** | Unbind removes the bindings of these channels, so they fall back to Default. | [optional] 

## Methods

### NewTransportTransportAgentsPut

`func NewTransportTransportAgentsPut() *TransportTransportAgentsPut`

NewTransportTransportAgentsPut instantiates a new TransportTransportAgentsPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportTransportAgentsPutWithDefaults

`func NewTransportTransportAgentsPutWithDefaults() *TransportTransportAgentsPut`

NewTransportTransportAgentsPutWithDefaults instantiates a new TransportTransportAgentsPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *TransportTransportAgentsPut) GetChannels() map[string]string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *TransportTransportAgentsPut) GetChannelsOk() (*map[string]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *TransportTransportAgentsPut) SetChannels(v map[string]string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *TransportTransportAgentsPut) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetDefault

`func (o *TransportTransportAgentsPut) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *TransportTransportAgentsPut) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *TransportTransportAgentsPut) SetDefault(v string)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *TransportTransportAgentsPut) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetTransport

`func (o *TransportTransportAgentsPut) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *TransportTransportAgentsPut) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *TransportTransportAgentsPut) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *TransportTransportAgentsPut) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetUnbind

`func (o *TransportTransportAgentsPut) GetUnbind() []string`

GetUnbind returns the Unbind field if non-nil, zero value otherwise.

### GetUnbindOk

`func (o *TransportTransportAgentsPut) GetUnbindOk() (*[]string, bool)`

GetUnbindOk returns a tuple with the Unbind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbind

`func (o *TransportTransportAgentsPut) SetUnbind(v []string)`

SetUnbind sets Unbind field to given value.

### HasUnbind

`func (o *TransportTransportAgentsPut) HasUnbind() bool`

HasUnbind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


