# InfoOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jetstream** | Pointer to **bool** | JetStream is true when durable streams are enabled. | [optional] 
**MaxPayload** | Pointer to **int32** | MaxPayload is the broker&#39;s message-size ceiling in bytes. | [optional] 
**ServerId** | Pointer to **string** | Server is the broker&#39;s server id. | [optional] 
**ServerName** | Pointer to **string** | Name is the broker&#39;s server name. | [optional] 
**Streams** | Pointer to **int32** | Streams is the org&#39;s stream count. | [optional] 
**Version** | Pointer to **string** | Version is the broker&#39;s server version. | [optional] 

## Methods

### NewInfoOut

`func NewInfoOut() *InfoOut`

NewInfoOut instantiates a new InfoOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoOutWithDefaults

`func NewInfoOutWithDefaults() *InfoOut`

NewInfoOutWithDefaults instantiates a new InfoOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJetstream

`func (o *InfoOut) GetJetstream() bool`

GetJetstream returns the Jetstream field if non-nil, zero value otherwise.

### GetJetstreamOk

`func (o *InfoOut) GetJetstreamOk() (*bool, bool)`

GetJetstreamOk returns a tuple with the Jetstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetstream

`func (o *InfoOut) SetJetstream(v bool)`

SetJetstream sets Jetstream field to given value.

### HasJetstream

`func (o *InfoOut) HasJetstream() bool`

HasJetstream returns a boolean if a field has been set.

### GetMaxPayload

`func (o *InfoOut) GetMaxPayload() int32`

GetMaxPayload returns the MaxPayload field if non-nil, zero value otherwise.

### GetMaxPayloadOk

`func (o *InfoOut) GetMaxPayloadOk() (*int32, bool)`

GetMaxPayloadOk returns a tuple with the MaxPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPayload

`func (o *InfoOut) SetMaxPayload(v int32)`

SetMaxPayload sets MaxPayload field to given value.

### HasMaxPayload

`func (o *InfoOut) HasMaxPayload() bool`

HasMaxPayload returns a boolean if a field has been set.

### GetServerId

`func (o *InfoOut) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *InfoOut) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *InfoOut) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *InfoOut) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *InfoOut) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *InfoOut) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *InfoOut) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *InfoOut) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetStreams

`func (o *InfoOut) GetStreams() int32`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *InfoOut) GetStreamsOk() (*int32, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *InfoOut) SetStreams(v int32)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *InfoOut) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetVersion

`func (o *InfoOut) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InfoOut) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InfoOut) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InfoOut) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


