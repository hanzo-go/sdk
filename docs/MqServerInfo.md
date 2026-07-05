# MqServerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **string** | Unique server identifier. | [optional] 
**ServerName** | Pointer to **string** | Server name. | [optional] 
**Version** | Pointer to **string** | NATS server version. | [optional] 
**Go** | Pointer to **string** | Go runtime version. | [optional] 
**Host** | Pointer to **string** | Server host. | [optional] 
**Port** | Pointer to **int32** | Client port. | [optional] 
**MaxPayload** | Pointer to **int32** | Maximum payload size in bytes. | [optional] 
**Jetstream** | Pointer to **bool** | Whether JetStream is enabled. | [optional] 
**Cluster** | Pointer to [**MqServerInfoCluster**](MqServerInfoCluster.md) |  | [optional] 
**Uptime** | Pointer to **string** | Server uptime. | [optional] 
**TotalConnections** | Pointer to **int32** | Total connections served since start. | [optional] 

## Methods

### NewMqServerInfo

`func NewMqServerInfo() *MqServerInfo`

NewMqServerInfo instantiates a new MqServerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqServerInfoWithDefaults

`func NewMqServerInfoWithDefaults() *MqServerInfo`

NewMqServerInfoWithDefaults instantiates a new MqServerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *MqServerInfo) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *MqServerInfo) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *MqServerInfo) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *MqServerInfo) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *MqServerInfo) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *MqServerInfo) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *MqServerInfo) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *MqServerInfo) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetVersion

`func (o *MqServerInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MqServerInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MqServerInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MqServerInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetGo

`func (o *MqServerInfo) GetGo() string`

GetGo returns the Go field if non-nil, zero value otherwise.

### GetGoOk

`func (o *MqServerInfo) GetGoOk() (*string, bool)`

GetGoOk returns a tuple with the Go field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGo

`func (o *MqServerInfo) SetGo(v string)`

SetGo sets Go field to given value.

### HasGo

`func (o *MqServerInfo) HasGo() bool`

HasGo returns a boolean if a field has been set.

### GetHost

`func (o *MqServerInfo) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MqServerInfo) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MqServerInfo) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *MqServerInfo) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *MqServerInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MqServerInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MqServerInfo) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *MqServerInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetMaxPayload

`func (o *MqServerInfo) GetMaxPayload() int32`

GetMaxPayload returns the MaxPayload field if non-nil, zero value otherwise.

### GetMaxPayloadOk

`func (o *MqServerInfo) GetMaxPayloadOk() (*int32, bool)`

GetMaxPayloadOk returns a tuple with the MaxPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPayload

`func (o *MqServerInfo) SetMaxPayload(v int32)`

SetMaxPayload sets MaxPayload field to given value.

### HasMaxPayload

`func (o *MqServerInfo) HasMaxPayload() bool`

HasMaxPayload returns a boolean if a field has been set.

### GetJetstream

`func (o *MqServerInfo) GetJetstream() bool`

GetJetstream returns the Jetstream field if non-nil, zero value otherwise.

### GetJetstreamOk

`func (o *MqServerInfo) GetJetstreamOk() (*bool, bool)`

GetJetstreamOk returns a tuple with the Jetstream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetstream

`func (o *MqServerInfo) SetJetstream(v bool)`

SetJetstream sets Jetstream field to given value.

### HasJetstream

`func (o *MqServerInfo) HasJetstream() bool`

HasJetstream returns a boolean if a field has been set.

### GetCluster

`func (o *MqServerInfo) GetCluster() MqServerInfoCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *MqServerInfo) GetClusterOk() (*MqServerInfoCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *MqServerInfo) SetCluster(v MqServerInfoCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *MqServerInfo) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetUptime

`func (o *MqServerInfo) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *MqServerInfo) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *MqServerInfo) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *MqServerInfo) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetTotalConnections

`func (o *MqServerInfo) GetTotalConnections() int32`

GetTotalConnections returns the TotalConnections field if non-nil, zero value otherwise.

### GetTotalConnectionsOk

`func (o *MqServerInfo) GetTotalConnectionsOk() (*int32, bool)`

GetTotalConnectionsOk returns a tuple with the TotalConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalConnections

`func (o *MqServerInfo) SetTotalConnections(v int32)`

SetTotalConnections sets TotalConnections field to given value.

### HasTotalConnections

`func (o *MqServerInfo) HasTotalConnections() bool`

HasTotalConnections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


