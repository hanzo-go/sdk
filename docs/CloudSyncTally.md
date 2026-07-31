# CloudSyncTally

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Live** | Pointer to **int32** | Live is the number of vouchers newly posted to the live ledger. | [optional] 
**Sandbox** | Pointer to **int32** | Sandbox is the number newly posted to the sandbox ledger. | [optional] 

## Methods

### NewCloudSyncTally

`func NewCloudSyncTally() *CloudSyncTally`

NewCloudSyncTally instantiates a new CloudSyncTally object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSyncTallyWithDefaults

`func NewCloudSyncTallyWithDefaults() *CloudSyncTally`

NewCloudSyncTallyWithDefaults instantiates a new CloudSyncTally object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLive

`func (o *CloudSyncTally) GetLive() int32`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *CloudSyncTally) GetLiveOk() (*int32, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *CloudSyncTally) SetLive(v int32)`

SetLive sets Live field to given value.

### HasLive

`func (o *CloudSyncTally) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetSandbox

`func (o *CloudSyncTally) GetSandbox() int32`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *CloudSyncTally) GetSandboxOk() (*int32, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *CloudSyncTally) SetSandbox(v int32)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *CloudSyncTally) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


