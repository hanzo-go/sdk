# SandboxQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the account the quota belongs to: the org, or a person in the signup org. | [optional] 
**Quota** | Pointer to **int64** | Quota is how many it may hold at once; -1 for no bound. | [optional] 
**Running** | Pointer to **int64** | Running is how many of its paid sandboxes hold a pod now. | [optional] 

## Methods

### NewSandboxQuota

`func NewSandboxQuota() *SandboxQuota`

NewSandboxQuota instantiates a new SandboxQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxQuotaWithDefaults

`func NewSandboxQuotaWithDefaults() *SandboxQuota`

NewSandboxQuotaWithDefaults instantiates a new SandboxQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *SandboxQuota) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SandboxQuota) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SandboxQuota) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SandboxQuota) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetQuota

`func (o *SandboxQuota) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *SandboxQuota) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *SandboxQuota) SetQuota(v int64)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *SandboxQuota) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetRunning

`func (o *SandboxQuota) GetRunning() int64`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *SandboxQuota) GetRunningOk() (*int64, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *SandboxQuota) SetRunning(v int64)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *SandboxQuota) HasRunning() bool`

HasRunning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


