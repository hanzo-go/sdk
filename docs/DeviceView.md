# DeviceView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]LinkView**](LinkView.md) | Accounts is every account the caller has signed in on this machine. | [optional] 
**ActiveSessions** | Pointer to **int64** | ActiveSessions is how many agent sessions the caller currently has running on this machine; 0 where the agent plane is not mounted. | [optional] 
**Host** | Pointer to **string** | Host is the machine&#39;s hostname label, from its most-recently-seen account. | [optional] 
**LastSeen** | Pointer to **string** | LastSeen is when any account on this machine last reported, RFC 3339 UTC. | [optional] 
**Machine** | Pointer to **string** | Machine is the stable machine identifier. | [optional] 
**Os** | Pointer to **string** | OS is the machine&#39;s operating system label. | [optional] 

## Methods

### NewDeviceView

`func NewDeviceView() *DeviceView`

NewDeviceView instantiates a new DeviceView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceViewWithDefaults

`func NewDeviceViewWithDefaults() *DeviceView`

NewDeviceViewWithDefaults instantiates a new DeviceView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *DeviceView) GetAccounts() []LinkView`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *DeviceView) GetAccountsOk() (*[]LinkView, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *DeviceView) SetAccounts(v []LinkView)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *DeviceView) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetActiveSessions

`func (o *DeviceView) GetActiveSessions() int64`

GetActiveSessions returns the ActiveSessions field if non-nil, zero value otherwise.

### GetActiveSessionsOk

`func (o *DeviceView) GetActiveSessionsOk() (*int64, bool)`

GetActiveSessionsOk returns a tuple with the ActiveSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSessions

`func (o *DeviceView) SetActiveSessions(v int64)`

SetActiveSessions sets ActiveSessions field to given value.

### HasActiveSessions

`func (o *DeviceView) HasActiveSessions() bool`

HasActiveSessions returns a boolean if a field has been set.

### GetHost

`func (o *DeviceView) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DeviceView) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DeviceView) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DeviceView) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetLastSeen

`func (o *DeviceView) GetLastSeen() string`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *DeviceView) GetLastSeenOk() (*string, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *DeviceView) SetLastSeen(v string)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *DeviceView) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMachine

`func (o *DeviceView) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *DeviceView) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *DeviceView) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *DeviceView) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetOs

`func (o *DeviceView) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *DeviceView) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *DeviceView) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *DeviceView) HasOs() bool`

HasOs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


