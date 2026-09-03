# SyncTally

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Live** | Pointer to **int64** | Live is the number of vouchers newly posted to the live ledger. | [optional] 
**Sandbox** | Pointer to **int64** | Sandbox is the number newly posted to the sandbox ledger. | [optional] 

## Methods

### NewSyncTally

`func NewSyncTally() *SyncTally`

NewSyncTally instantiates a new SyncTally object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncTallyWithDefaults

`func NewSyncTallyWithDefaults() *SyncTally`

NewSyncTallyWithDefaults instantiates a new SyncTally object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLive

`func (o *SyncTally) GetLive() int64`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *SyncTally) GetLiveOk() (*int64, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *SyncTally) SetLive(v int64)`

SetLive sets Live field to given value.

### HasLive

`func (o *SyncTally) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetSandbox

`func (o *SyncTally) GetSandbox() int64`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *SyncTally) GetSandboxOk() (*int64, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *SyncTally) SetSandbox(v int64)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *SyncTally) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


