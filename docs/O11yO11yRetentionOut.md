# O11yO11yRetentionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColdStorageTtlDays** | Pointer to **int32** | ColdStorageTTLDays is how old data must be before it moves, in days. | [optional] 
**ColdStorageVolume** | Pointer to **string** | ColdStorageVolume names the volume aged data moves to. | [optional] 
**DefaultTtlDays** | Pointer to **int32** | DefaultTTLDays is the retention for data no rule matches, in days. | [optional] 
**ExpectedLogsMoveTtlDurationHrs** | Pointer to **int32** | ExpectedLogsMoveTTLHours is the pending logs cold-storage move TTL, in hours. | [optional] 
**ExpectedLogsTtlDurationHrs** | Pointer to **int32** | ExpectedLogsTTLHours is the pending logs TTL, in hours. | [optional] 
**Status** | Pointer to **string** | Status is the last TTL operation&#39;s state. | [optional] 
**TtlConditions** | Pointer to [**[]O11yO11yRetentionRule**](O11yO11yRetentionRule.md) | TTLConditions are the ordered per-label rules; the first match wins. | [optional] 
**Version** | Pointer to **string** | Version is the policy format version. | [optional] 

## Methods

### NewO11yO11yRetentionOut

`func NewO11yO11yRetentionOut() *O11yO11yRetentionOut`

NewO11yO11yRetentionOut instantiates a new O11yO11yRetentionOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yRetentionOutWithDefaults

`func NewO11yO11yRetentionOutWithDefaults() *O11yO11yRetentionOut`

NewO11yO11yRetentionOutWithDefaults instantiates a new O11yO11yRetentionOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColdStorageTtlDays

`func (o *O11yO11yRetentionOut) GetColdStorageTtlDays() int32`

GetColdStorageTtlDays returns the ColdStorageTtlDays field if non-nil, zero value otherwise.

### GetColdStorageTtlDaysOk

`func (o *O11yO11yRetentionOut) GetColdStorageTtlDaysOk() (*int32, bool)`

GetColdStorageTtlDaysOk returns a tuple with the ColdStorageTtlDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColdStorageTtlDays

`func (o *O11yO11yRetentionOut) SetColdStorageTtlDays(v int32)`

SetColdStorageTtlDays sets ColdStorageTtlDays field to given value.

### HasColdStorageTtlDays

`func (o *O11yO11yRetentionOut) HasColdStorageTtlDays() bool`

HasColdStorageTtlDays returns a boolean if a field has been set.

### GetColdStorageVolume

`func (o *O11yO11yRetentionOut) GetColdStorageVolume() string`

GetColdStorageVolume returns the ColdStorageVolume field if non-nil, zero value otherwise.

### GetColdStorageVolumeOk

`func (o *O11yO11yRetentionOut) GetColdStorageVolumeOk() (*string, bool)`

GetColdStorageVolumeOk returns a tuple with the ColdStorageVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColdStorageVolume

`func (o *O11yO11yRetentionOut) SetColdStorageVolume(v string)`

SetColdStorageVolume sets ColdStorageVolume field to given value.

### HasColdStorageVolume

`func (o *O11yO11yRetentionOut) HasColdStorageVolume() bool`

HasColdStorageVolume returns a boolean if a field has been set.

### GetDefaultTtlDays

`func (o *O11yO11yRetentionOut) GetDefaultTtlDays() int32`

GetDefaultTtlDays returns the DefaultTtlDays field if non-nil, zero value otherwise.

### GetDefaultTtlDaysOk

`func (o *O11yO11yRetentionOut) GetDefaultTtlDaysOk() (*int32, bool)`

GetDefaultTtlDaysOk returns a tuple with the DefaultTtlDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTtlDays

`func (o *O11yO11yRetentionOut) SetDefaultTtlDays(v int32)`

SetDefaultTtlDays sets DefaultTtlDays field to given value.

### HasDefaultTtlDays

`func (o *O11yO11yRetentionOut) HasDefaultTtlDays() bool`

HasDefaultTtlDays returns a boolean if a field has been set.

### GetExpectedLogsMoveTtlDurationHrs

`func (o *O11yO11yRetentionOut) GetExpectedLogsMoveTtlDurationHrs() int32`

GetExpectedLogsMoveTtlDurationHrs returns the ExpectedLogsMoveTtlDurationHrs field if non-nil, zero value otherwise.

### GetExpectedLogsMoveTtlDurationHrsOk

`func (o *O11yO11yRetentionOut) GetExpectedLogsMoveTtlDurationHrsOk() (*int32, bool)`

GetExpectedLogsMoveTtlDurationHrsOk returns a tuple with the ExpectedLogsMoveTtlDurationHrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedLogsMoveTtlDurationHrs

`func (o *O11yO11yRetentionOut) SetExpectedLogsMoveTtlDurationHrs(v int32)`

SetExpectedLogsMoveTtlDurationHrs sets ExpectedLogsMoveTtlDurationHrs field to given value.

### HasExpectedLogsMoveTtlDurationHrs

`func (o *O11yO11yRetentionOut) HasExpectedLogsMoveTtlDurationHrs() bool`

HasExpectedLogsMoveTtlDurationHrs returns a boolean if a field has been set.

### GetExpectedLogsTtlDurationHrs

`func (o *O11yO11yRetentionOut) GetExpectedLogsTtlDurationHrs() int32`

GetExpectedLogsTtlDurationHrs returns the ExpectedLogsTtlDurationHrs field if non-nil, zero value otherwise.

### GetExpectedLogsTtlDurationHrsOk

`func (o *O11yO11yRetentionOut) GetExpectedLogsTtlDurationHrsOk() (*int32, bool)`

GetExpectedLogsTtlDurationHrsOk returns a tuple with the ExpectedLogsTtlDurationHrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedLogsTtlDurationHrs

`func (o *O11yO11yRetentionOut) SetExpectedLogsTtlDurationHrs(v int32)`

SetExpectedLogsTtlDurationHrs sets ExpectedLogsTtlDurationHrs field to given value.

### HasExpectedLogsTtlDurationHrs

`func (o *O11yO11yRetentionOut) HasExpectedLogsTtlDurationHrs() bool`

HasExpectedLogsTtlDurationHrs returns a boolean if a field has been set.

### GetStatus

`func (o *O11yO11yRetentionOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11yO11yRetentionOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11yO11yRetentionOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11yO11yRetentionOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTtlConditions

`func (o *O11yO11yRetentionOut) GetTtlConditions() []O11yO11yRetentionRule`

GetTtlConditions returns the TtlConditions field if non-nil, zero value otherwise.

### GetTtlConditionsOk

`func (o *O11yO11yRetentionOut) GetTtlConditionsOk() (*[]O11yO11yRetentionRule, bool)`

GetTtlConditionsOk returns a tuple with the TtlConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlConditions

`func (o *O11yO11yRetentionOut) SetTtlConditions(v []O11yO11yRetentionRule)`

SetTtlConditions sets TtlConditions field to given value.

### HasTtlConditions

`func (o *O11yO11yRetentionOut) HasTtlConditions() bool`

HasTtlConditions returns a boolean if a field has been set.

### GetVersion

`func (o *O11yO11yRetentionOut) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *O11yO11yRetentionOut) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *O11yO11yRetentionOut) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *O11yO11yRetentionOut) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


