# MqAccountLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxConnections** | Pointer to **int32** | Maximum concurrent connections. | [optional] 
**MaxSubscriptions** | Pointer to **int32** | Maximum subscriptions. | [optional] 
**MaxPayload** | Pointer to **int32** | Maximum message payload size in bytes. | [optional] 
**MaxStreams** | Pointer to **int32** | Maximum JetStream streams. | [optional] 
**MaxConsumers** | Pointer to **int32** | Maximum JetStream consumers. | [optional] 
**MaxBytes** | Pointer to **int64** | Maximum total JetStream storage in bytes. | [optional] 

## Methods

### NewMqAccountLimits

`func NewMqAccountLimits() *MqAccountLimits`

NewMqAccountLimits instantiates a new MqAccountLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqAccountLimitsWithDefaults

`func NewMqAccountLimitsWithDefaults() *MqAccountLimits`

NewMqAccountLimitsWithDefaults instantiates a new MqAccountLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxConnections

`func (o *MqAccountLimits) GetMaxConnections() int32`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *MqAccountLimits) GetMaxConnectionsOk() (*int32, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *MqAccountLimits) SetMaxConnections(v int32)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *MqAccountLimits) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetMaxSubscriptions

`func (o *MqAccountLimits) GetMaxSubscriptions() int32`

GetMaxSubscriptions returns the MaxSubscriptions field if non-nil, zero value otherwise.

### GetMaxSubscriptionsOk

`func (o *MqAccountLimits) GetMaxSubscriptionsOk() (*int32, bool)`

GetMaxSubscriptionsOk returns a tuple with the MaxSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSubscriptions

`func (o *MqAccountLimits) SetMaxSubscriptions(v int32)`

SetMaxSubscriptions sets MaxSubscriptions field to given value.

### HasMaxSubscriptions

`func (o *MqAccountLimits) HasMaxSubscriptions() bool`

HasMaxSubscriptions returns a boolean if a field has been set.

### GetMaxPayload

`func (o *MqAccountLimits) GetMaxPayload() int32`

GetMaxPayload returns the MaxPayload field if non-nil, zero value otherwise.

### GetMaxPayloadOk

`func (o *MqAccountLimits) GetMaxPayloadOk() (*int32, bool)`

GetMaxPayloadOk returns a tuple with the MaxPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPayload

`func (o *MqAccountLimits) SetMaxPayload(v int32)`

SetMaxPayload sets MaxPayload field to given value.

### HasMaxPayload

`func (o *MqAccountLimits) HasMaxPayload() bool`

HasMaxPayload returns a boolean if a field has been set.

### GetMaxStreams

`func (o *MqAccountLimits) GetMaxStreams() int32`

GetMaxStreams returns the MaxStreams field if non-nil, zero value otherwise.

### GetMaxStreamsOk

`func (o *MqAccountLimits) GetMaxStreamsOk() (*int32, bool)`

GetMaxStreamsOk returns a tuple with the MaxStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStreams

`func (o *MqAccountLimits) SetMaxStreams(v int32)`

SetMaxStreams sets MaxStreams field to given value.

### HasMaxStreams

`func (o *MqAccountLimits) HasMaxStreams() bool`

HasMaxStreams returns a boolean if a field has been set.

### GetMaxConsumers

`func (o *MqAccountLimits) GetMaxConsumers() int32`

GetMaxConsumers returns the MaxConsumers field if non-nil, zero value otherwise.

### GetMaxConsumersOk

`func (o *MqAccountLimits) GetMaxConsumersOk() (*int32, bool)`

GetMaxConsumersOk returns a tuple with the MaxConsumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConsumers

`func (o *MqAccountLimits) SetMaxConsumers(v int32)`

SetMaxConsumers sets MaxConsumers field to given value.

### HasMaxConsumers

`func (o *MqAccountLimits) HasMaxConsumers() bool`

HasMaxConsumers returns a boolean if a field has been set.

### GetMaxBytes

`func (o *MqAccountLimits) GetMaxBytes() int64`

GetMaxBytes returns the MaxBytes field if non-nil, zero value otherwise.

### GetMaxBytesOk

`func (o *MqAccountLimits) GetMaxBytesOk() (*int64, bool)`

GetMaxBytesOk returns a tuple with the MaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBytes

`func (o *MqAccountLimits) SetMaxBytes(v int64)`

SetMaxBytes sets MaxBytes field to given value.

### HasMaxBytes

`func (o *MqAccountLimits) HasMaxBytes() bool`

HasMaxBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


