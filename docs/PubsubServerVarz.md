# PubsubServerVarz

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **string** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Uptime** | Pointer to **string** |  | [optional] 
**Mem** | Pointer to **int32** |  | [optional] 
**Cores** | Pointer to **int32** |  | [optional] 
**Connections** | Pointer to **int32** |  | [optional] 
**Subscriptions** | Pointer to **int32** |  | [optional] 
**SlowConsumers** | Pointer to **int32** |  | [optional] 
**InMsgs** | Pointer to **int32** |  | [optional] 
**OutMsgs** | Pointer to **int32** |  | [optional] 
**InBytes** | Pointer to **int32** |  | [optional] 
**OutBytes** | Pointer to **int32** |  | [optional] 

## Methods

### NewPubsubServerVarz

`func NewPubsubServerVarz() *PubsubServerVarz`

NewPubsubServerVarz instantiates a new PubsubServerVarz object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubsubServerVarzWithDefaults

`func NewPubsubServerVarzWithDefaults() *PubsubServerVarz`

NewPubsubServerVarzWithDefaults instantiates a new PubsubServerVarz object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *PubsubServerVarz) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *PubsubServerVarz) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *PubsubServerVarz) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *PubsubServerVarz) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *PubsubServerVarz) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *PubsubServerVarz) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *PubsubServerVarz) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *PubsubServerVarz) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetVersion

`func (o *PubsubServerVarz) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PubsubServerVarz) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PubsubServerVarz) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PubsubServerVarz) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUptime

`func (o *PubsubServerVarz) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *PubsubServerVarz) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *PubsubServerVarz) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *PubsubServerVarz) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetMem

`func (o *PubsubServerVarz) GetMem() int32`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *PubsubServerVarz) GetMemOk() (*int32, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *PubsubServerVarz) SetMem(v int32)`

SetMem sets Mem field to given value.

### HasMem

`func (o *PubsubServerVarz) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetCores

`func (o *PubsubServerVarz) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *PubsubServerVarz) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *PubsubServerVarz) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *PubsubServerVarz) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetConnections

`func (o *PubsubServerVarz) GetConnections() int32`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *PubsubServerVarz) GetConnectionsOk() (*int32, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *PubsubServerVarz) SetConnections(v int32)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *PubsubServerVarz) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetSubscriptions

`func (o *PubsubServerVarz) GetSubscriptions() int32`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *PubsubServerVarz) GetSubscriptionsOk() (*int32, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *PubsubServerVarz) SetSubscriptions(v int32)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *PubsubServerVarz) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetSlowConsumers

`func (o *PubsubServerVarz) GetSlowConsumers() int32`

GetSlowConsumers returns the SlowConsumers field if non-nil, zero value otherwise.

### GetSlowConsumersOk

`func (o *PubsubServerVarz) GetSlowConsumersOk() (*int32, bool)`

GetSlowConsumersOk returns a tuple with the SlowConsumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowConsumers

`func (o *PubsubServerVarz) SetSlowConsumers(v int32)`

SetSlowConsumers sets SlowConsumers field to given value.

### HasSlowConsumers

`func (o *PubsubServerVarz) HasSlowConsumers() bool`

HasSlowConsumers returns a boolean if a field has been set.

### GetInMsgs

`func (o *PubsubServerVarz) GetInMsgs() int32`

GetInMsgs returns the InMsgs field if non-nil, zero value otherwise.

### GetInMsgsOk

`func (o *PubsubServerVarz) GetInMsgsOk() (*int32, bool)`

GetInMsgsOk returns a tuple with the InMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMsgs

`func (o *PubsubServerVarz) SetInMsgs(v int32)`

SetInMsgs sets InMsgs field to given value.

### HasInMsgs

`func (o *PubsubServerVarz) HasInMsgs() bool`

HasInMsgs returns a boolean if a field has been set.

### GetOutMsgs

`func (o *PubsubServerVarz) GetOutMsgs() int32`

GetOutMsgs returns the OutMsgs field if non-nil, zero value otherwise.

### GetOutMsgsOk

`func (o *PubsubServerVarz) GetOutMsgsOk() (*int32, bool)`

GetOutMsgsOk returns a tuple with the OutMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutMsgs

`func (o *PubsubServerVarz) SetOutMsgs(v int32)`

SetOutMsgs sets OutMsgs field to given value.

### HasOutMsgs

`func (o *PubsubServerVarz) HasOutMsgs() bool`

HasOutMsgs returns a boolean if a field has been set.

### GetInBytes

`func (o *PubsubServerVarz) GetInBytes() int32`

GetInBytes returns the InBytes field if non-nil, zero value otherwise.

### GetInBytesOk

`func (o *PubsubServerVarz) GetInBytesOk() (*int32, bool)`

GetInBytesOk returns a tuple with the InBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInBytes

`func (o *PubsubServerVarz) SetInBytes(v int32)`

SetInBytes sets InBytes field to given value.

### HasInBytes

`func (o *PubsubServerVarz) HasInBytes() bool`

HasInBytes returns a boolean if a field has been set.

### GetOutBytes

`func (o *PubsubServerVarz) GetOutBytes() int32`

GetOutBytes returns the OutBytes field if non-nil, zero value otherwise.

### GetOutBytesOk

`func (o *PubsubServerVarz) GetOutBytesOk() (*int32, bool)`

GetOutBytesOk returns a tuple with the OutBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutBytes

`func (o *PubsubServerVarz) SetOutBytes(v int32)`

SetOutBytes sets OutBytes field to given value.

### HasOutBytes

`func (o *PubsubServerVarz) HasOutBytes() bool`

HasOutBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


