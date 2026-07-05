# MqConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Connection ID. | [optional] 
**Name** | Pointer to **string** | Client-supplied connection name. | [optional] 
**Ip** | Pointer to **string** | Client IP address. | [optional] 
**Port** | Pointer to **int32** | Client port. | [optional] 
**Lang** | Pointer to **string** | Client library language. | [optional] 
**Version** | Pointer to **string** | Client library version. | [optional] 
**Subscriptions** | Pointer to **int32** | Number of subscriptions on this connection. | [optional] 
**InMsgs** | Pointer to **int64** | Messages received. | [optional] 
**OutMsgs** | Pointer to **int64** | Messages sent. | [optional] 
**InBytes** | Pointer to **int64** | Bytes received. | [optional] 
**OutBytes** | Pointer to **int64** | Bytes sent. | [optional] 
**Uptime** | Pointer to **string** | Connection uptime. | [optional] 
**Idle** | Pointer to **string** | Time since last activity. | [optional] 
**Started** | Pointer to **time.Time** | Connection start time. | [optional] 
**TlsVersion** | Pointer to **string** | TLS version (if TLS is active). | [optional] 

## Methods

### NewMqConnection

`func NewMqConnection() *MqConnection`

NewMqConnection instantiates a new MqConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqConnectionWithDefaults

`func NewMqConnectionWithDefaults() *MqConnection`

NewMqConnectionWithDefaults instantiates a new MqConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MqConnection) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MqConnection) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MqConnection) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MqConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MqConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MqConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MqConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MqConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIp

`func (o *MqConnection) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *MqConnection) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *MqConnection) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *MqConnection) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPort

`func (o *MqConnection) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MqConnection) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MqConnection) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *MqConnection) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetLang

`func (o *MqConnection) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *MqConnection) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *MqConnection) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *MqConnection) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetVersion

`func (o *MqConnection) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MqConnection) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MqConnection) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MqConnection) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSubscriptions

`func (o *MqConnection) GetSubscriptions() int32`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *MqConnection) GetSubscriptionsOk() (*int32, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *MqConnection) SetSubscriptions(v int32)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *MqConnection) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetInMsgs

`func (o *MqConnection) GetInMsgs() int64`

GetInMsgs returns the InMsgs field if non-nil, zero value otherwise.

### GetInMsgsOk

`func (o *MqConnection) GetInMsgsOk() (*int64, bool)`

GetInMsgsOk returns a tuple with the InMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMsgs

`func (o *MqConnection) SetInMsgs(v int64)`

SetInMsgs sets InMsgs field to given value.

### HasInMsgs

`func (o *MqConnection) HasInMsgs() bool`

HasInMsgs returns a boolean if a field has been set.

### GetOutMsgs

`func (o *MqConnection) GetOutMsgs() int64`

GetOutMsgs returns the OutMsgs field if non-nil, zero value otherwise.

### GetOutMsgsOk

`func (o *MqConnection) GetOutMsgsOk() (*int64, bool)`

GetOutMsgsOk returns a tuple with the OutMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutMsgs

`func (o *MqConnection) SetOutMsgs(v int64)`

SetOutMsgs sets OutMsgs field to given value.

### HasOutMsgs

`func (o *MqConnection) HasOutMsgs() bool`

HasOutMsgs returns a boolean if a field has been set.

### GetInBytes

`func (o *MqConnection) GetInBytes() int64`

GetInBytes returns the InBytes field if non-nil, zero value otherwise.

### GetInBytesOk

`func (o *MqConnection) GetInBytesOk() (*int64, bool)`

GetInBytesOk returns a tuple with the InBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInBytes

`func (o *MqConnection) SetInBytes(v int64)`

SetInBytes sets InBytes field to given value.

### HasInBytes

`func (o *MqConnection) HasInBytes() bool`

HasInBytes returns a boolean if a field has been set.

### GetOutBytes

`func (o *MqConnection) GetOutBytes() int64`

GetOutBytes returns the OutBytes field if non-nil, zero value otherwise.

### GetOutBytesOk

`func (o *MqConnection) GetOutBytesOk() (*int64, bool)`

GetOutBytesOk returns a tuple with the OutBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutBytes

`func (o *MqConnection) SetOutBytes(v int64)`

SetOutBytes sets OutBytes field to given value.

### HasOutBytes

`func (o *MqConnection) HasOutBytes() bool`

HasOutBytes returns a boolean if a field has been set.

### GetUptime

`func (o *MqConnection) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *MqConnection) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *MqConnection) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *MqConnection) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetIdle

`func (o *MqConnection) GetIdle() string`

GetIdle returns the Idle field if non-nil, zero value otherwise.

### GetIdleOk

`func (o *MqConnection) GetIdleOk() (*string, bool)`

GetIdleOk returns a tuple with the Idle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdle

`func (o *MqConnection) SetIdle(v string)`

SetIdle sets Idle field to given value.

### HasIdle

`func (o *MqConnection) HasIdle() bool`

HasIdle returns a boolean if a field has been set.

### GetStarted

`func (o *MqConnection) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *MqConnection) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *MqConnection) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *MqConnection) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetTlsVersion

`func (o *MqConnection) GetTlsVersion() string`

GetTlsVersion returns the TlsVersion field if non-nil, zero value otherwise.

### GetTlsVersionOk

`func (o *MqConnection) GetTlsVersionOk() (*string, bool)`

GetTlsVersionOk returns a tuple with the TlsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVersion

`func (o *MqConnection) SetTlsVersion(v string)`

SetTlsVersion sets TlsVersion field to given value.

### HasTlsVersion

`func (o *MqConnection) HasTlsVersion() bool`

HasTlsVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


