# MqAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Account ID. | [optional] 
**OrgId** | Pointer to **string** | Hanzo IAM organization ID. | [optional] 
**Name** | Pointer to **string** | Account display name. | [optional] 
**Connections** | Pointer to **int32** | Number of active connections. | [optional] 
**Subscriptions** | Pointer to **int32** | Number of active subscriptions. | [optional] 
**DataIn** | Pointer to **int64** | Total bytes received. | [optional] 
**DataOut** | Pointer to **int64** | Total bytes sent. | [optional] 
**SlowConsumers** | Pointer to **int32** | Number of slow consumers. | [optional] 
**Streams** | Pointer to **int32** | Number of JetStream streams. | [optional] 
**Consumers** | Pointer to **int32** | Number of JetStream consumers. | [optional] 
**Limits** | Pointer to [**MqAccountLimits**](MqAccountLimits.md) |  | [optional] 

## Methods

### NewMqAccount

`func NewMqAccount() *MqAccount`

NewMqAccount instantiates a new MqAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqAccountWithDefaults

`func NewMqAccountWithDefaults() *MqAccount`

NewMqAccountWithDefaults instantiates a new MqAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MqAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MqAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MqAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MqAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *MqAccount) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *MqAccount) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *MqAccount) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *MqAccount) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetName

`func (o *MqAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MqAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MqAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MqAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConnections

`func (o *MqAccount) GetConnections() int32`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *MqAccount) GetConnectionsOk() (*int32, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *MqAccount) SetConnections(v int32)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *MqAccount) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetSubscriptions

`func (o *MqAccount) GetSubscriptions() int32`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *MqAccount) GetSubscriptionsOk() (*int32, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *MqAccount) SetSubscriptions(v int32)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *MqAccount) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetDataIn

`func (o *MqAccount) GetDataIn() int64`

GetDataIn returns the DataIn field if non-nil, zero value otherwise.

### GetDataInOk

`func (o *MqAccount) GetDataInOk() (*int64, bool)`

GetDataInOk returns a tuple with the DataIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIn

`func (o *MqAccount) SetDataIn(v int64)`

SetDataIn sets DataIn field to given value.

### HasDataIn

`func (o *MqAccount) HasDataIn() bool`

HasDataIn returns a boolean if a field has been set.

### GetDataOut

`func (o *MqAccount) GetDataOut() int64`

GetDataOut returns the DataOut field if non-nil, zero value otherwise.

### GetDataOutOk

`func (o *MqAccount) GetDataOutOk() (*int64, bool)`

GetDataOutOk returns a tuple with the DataOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataOut

`func (o *MqAccount) SetDataOut(v int64)`

SetDataOut sets DataOut field to given value.

### HasDataOut

`func (o *MqAccount) HasDataOut() bool`

HasDataOut returns a boolean if a field has been set.

### GetSlowConsumers

`func (o *MqAccount) GetSlowConsumers() int32`

GetSlowConsumers returns the SlowConsumers field if non-nil, zero value otherwise.

### GetSlowConsumersOk

`func (o *MqAccount) GetSlowConsumersOk() (*int32, bool)`

GetSlowConsumersOk returns a tuple with the SlowConsumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowConsumers

`func (o *MqAccount) SetSlowConsumers(v int32)`

SetSlowConsumers sets SlowConsumers field to given value.

### HasSlowConsumers

`func (o *MqAccount) HasSlowConsumers() bool`

HasSlowConsumers returns a boolean if a field has been set.

### GetStreams

`func (o *MqAccount) GetStreams() int32`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *MqAccount) GetStreamsOk() (*int32, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *MqAccount) SetStreams(v int32)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *MqAccount) HasStreams() bool`

HasStreams returns a boolean if a field has been set.

### GetConsumers

`func (o *MqAccount) GetConsumers() int32`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *MqAccount) GetConsumersOk() (*int32, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *MqAccount) SetConsumers(v int32)`

SetConsumers sets Consumers field to given value.

### HasConsumers

`func (o *MqAccount) HasConsumers() bool`

HasConsumers returns a boolean if a field has been set.

### GetLimits

`func (o *MqAccount) GetLimits() MqAccountLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *MqAccount) GetLimitsOk() (*MqAccountLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *MqAccount) SetLimits(v MqAccountLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *MqAccount) HasLimits() bool`

HasLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


