# MqInfoOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jetstream** | Pointer to **bool** | JetStream is true when durable streams are enabled. | [optional] 
**MaxPayload** | Pointer to **int64** | MaxPayload is the broker&#39;s message-size ceiling in bytes. | [optional] 
**ServerId** | Pointer to **string** | Server is the broker&#39;s server id. | [optional] 
**ServerName** | Pointer to **string** | Name is the broker&#39;s server name. | [optional] 
**Streams** | Pointer to **int64** | Streams is the org&#39;s stream count. | [optional] 
**Version** | Pointer to **string** | Version is the broker&#39;s server version. | [optional] 

## Methods

### NewMqInfoOut

`func NewMqInfoOut() *MqInfoOut`

NewMqInfoOut instantiates a new MqInfoOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqInfoOutWithDefaults

`func NewMqInfoOutWithDefaults() *MqInfoOut`

NewMqInfoOutWithDefaults instantiates a new MqInfoOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJetstream

`func (o *MqInfoOut) GetJetstream() bool`

GetJetstream returns the Jetstream field if non-nil, zero value otherwise.

### GetJetstreamOk

`func (o *MqInfoOut) GetJetstreamOk() (*bool, bool)`

GetJetstreamOk returns a tuple with the Jetstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetstream

`func (o *MqInfoOut) SetJetstream(v bool)`

SetJetstream sets Jetstream field to given value.

### HasJetstream

`func (o *MqInfoOut) HasJetstream() bool`

HasJetstream returns a boolean if a field has been set.

### GetMaxPayload

`func (o *MqInfoOut) GetMaxPayload() int64`

GetMaxPayload returns the MaxPayload field if non-nil, zero value otherwise.

### GetMaxPayloadOk

`func (o *MqInfoOut) GetMaxPayloadOk() (*int64, bool)`

GetMaxPayloadOk returns a tuple with the MaxPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPayload

`func (o *MqInfoOut) SetMaxPayload(v int64)`

SetMaxPayload sets MaxPayload field to given value.

### HasMaxPayload

`func (o *MqInfoOut) HasMaxPayload() bool`

HasMaxPayload returns a boolean if a field has been set.

### GetServerId

`func (o *MqInfoOut) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *MqInfoOut) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *MqInfoOut) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *MqInfoOut) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *MqInfoOut) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *MqInfoOut) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *MqInfoOut) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *MqInfoOut) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetStreams

`func (o *MqInfoOut) GetStreams() int64`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *MqInfoOut) GetStreamsOk() (*int64, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *MqInfoOut) SetStreams(v int64)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *MqInfoOut) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetVersion

`func (o *MqInfoOut) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MqInfoOut) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MqInfoOut) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MqInfoOut) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


