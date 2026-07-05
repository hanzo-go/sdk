# StreamBrokerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **int32** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**PubsubUrl** | Pointer to **string** | Connected PubSub server URL | [optional] 
**Uptime** | Pointer to **string** |  | [optional] 
**Topics** | Pointer to **int32** |  | [optional] 
**Config** | Pointer to [**StreamBrokerInfoConfig**](StreamBrokerInfoConfig.md) |  | [optional] 

## Methods

### NewStreamBrokerInfo

`func NewStreamBrokerInfo() *StreamBrokerInfo`

NewStreamBrokerInfo instantiates a new StreamBrokerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamBrokerInfoWithDefaults

`func NewStreamBrokerInfoWithDefaults() *StreamBrokerInfo`

NewStreamBrokerInfoWithDefaults instantiates a new StreamBrokerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *StreamBrokerInfo) GetNodeId() int32`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *StreamBrokerInfo) GetNodeIdOk() (*int32, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *StreamBrokerInfo) SetNodeId(v int32)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *StreamBrokerInfo) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetHost

`func (o *StreamBrokerInfo) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *StreamBrokerInfo) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *StreamBrokerInfo) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *StreamBrokerInfo) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *StreamBrokerInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *StreamBrokerInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *StreamBrokerInfo) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *StreamBrokerInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetVersion

`func (o *StreamBrokerInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StreamBrokerInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StreamBrokerInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StreamBrokerInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetPubsubUrl

`func (o *StreamBrokerInfo) GetPubsubUrl() string`

GetPubsubUrl returns the PubsubUrl field if non-nil, zero value otherwise.

### GetPubsubUrlOk

`func (o *StreamBrokerInfo) GetPubsubUrlOk() (*string, bool)`

GetPubsubUrlOk returns a tuple with the PubsubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubsubUrl

`func (o *StreamBrokerInfo) SetPubsubUrl(v string)`

SetPubsubUrl sets PubsubUrl field to given value.

### HasPubsubUrl

`func (o *StreamBrokerInfo) HasPubsubUrl() bool`

HasPubsubUrl returns a boolean if a field has been set.

### GetUptime

`func (o *StreamBrokerInfo) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *StreamBrokerInfo) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *StreamBrokerInfo) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *StreamBrokerInfo) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetTopics

`func (o *StreamBrokerInfo) GetTopics() int32`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *StreamBrokerInfo) GetTopicsOk() (*int32, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *StreamBrokerInfo) SetTopics(v int32)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *StreamBrokerInfo) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetConfig

`func (o *StreamBrokerInfo) GetConfig() StreamBrokerInfoConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *StreamBrokerInfo) GetConfigOk() (*StreamBrokerInfoConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *StreamBrokerInfo) SetConfig(v StreamBrokerInfoConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *StreamBrokerInfo) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


