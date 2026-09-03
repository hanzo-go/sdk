# O11yO11yRetentionSetIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColdStorageDurationDays** | Pointer to **int64** | ColdStorageDurationDays is how old data must be before it moves, in days. | [optional] 
**ColdStorageVolume** | Pointer to **string** | ColdStorageVolume names the volume aged data moves to, when set. | [optional] 
**DefaultTTLDays** | Pointer to **int64** | DefaultTTLDays is the retention for data no rule matches, in days. | [optional] 
**TtlConditions** | Pointer to [**[]O11yO11yRetentionRule**](O11yO11yRetentionRule.md) | TTLConditions are ordered per-label rules; the first matching rule wins. | [optional] 
**Type** | Pointer to **string** | Type is the signal the policy applies to — traces, metrics or logs. | [optional] 

## Methods

### NewO11yO11yRetentionSetIn

`func NewO11yO11yRetentionSetIn() *O11yO11yRetentionSetIn`

NewO11yO11yRetentionSetIn instantiates a new O11yO11yRetentionSetIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yRetentionSetInWithDefaults

`func NewO11yO11yRetentionSetInWithDefaults() *O11yO11yRetentionSetIn`

NewO11yO11yRetentionSetInWithDefaults instantiates a new O11yO11yRetentionSetIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColdStorageDurationDays

`func (o *O11yO11yRetentionSetIn) GetColdStorageDurationDays() int64`

GetColdStorageDurationDays returns the ColdStorageDurationDays field if non-nil, zero value otherwise.

### GetColdStorageDurationDaysOk

`func (o *O11yO11yRetentionSetIn) GetColdStorageDurationDaysOk() (*int64, bool)`

GetColdStorageDurationDaysOk returns a tuple with the ColdStorageDurationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColdStorageDurationDays

`func (o *O11yO11yRetentionSetIn) SetColdStorageDurationDays(v int64)`

SetColdStorageDurationDays sets ColdStorageDurationDays field to given value.

### HasColdStorageDurationDays

`func (o *O11yO11yRetentionSetIn) HasColdStorageDurationDays() bool`

HasColdStorageDurationDays returns a boolean if a field has been set.

### GetColdStorageVolume

`func (o *O11yO11yRetentionSetIn) GetColdStorageVolume() string`

GetColdStorageVolume returns the ColdStorageVolume field if non-nil, zero value otherwise.

### GetColdStorageVolumeOk

`func (o *O11yO11yRetentionSetIn) GetColdStorageVolumeOk() (*string, bool)`

GetColdStorageVolumeOk returns a tuple with the ColdStorageVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColdStorageVolume

`func (o *O11yO11yRetentionSetIn) SetColdStorageVolume(v string)`

SetColdStorageVolume sets ColdStorageVolume field to given value.

### HasColdStorageVolume

`func (o *O11yO11yRetentionSetIn) HasColdStorageVolume() bool`

HasColdStorageVolume returns a boolean if a field has been set.

### GetDefaultTTLDays

`func (o *O11yO11yRetentionSetIn) GetDefaultTTLDays() int64`

GetDefaultTTLDays returns the DefaultTTLDays field if non-nil, zero value otherwise.

### GetDefaultTTLDaysOk

`func (o *O11yO11yRetentionSetIn) GetDefaultTTLDaysOk() (*int64, bool)`

GetDefaultTTLDaysOk returns a tuple with the DefaultTTLDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTTLDays

`func (o *O11yO11yRetentionSetIn) SetDefaultTTLDays(v int64)`

SetDefaultTTLDays sets DefaultTTLDays field to given value.

### HasDefaultTTLDays

`func (o *O11yO11yRetentionSetIn) HasDefaultTTLDays() bool`

HasDefaultTTLDays returns a boolean if a field has been set.

### GetTtlConditions

`func (o *O11yO11yRetentionSetIn) GetTtlConditions() []O11yO11yRetentionRule`

GetTtlConditions returns the TtlConditions field if non-nil, zero value otherwise.

### GetTtlConditionsOk

`func (o *O11yO11yRetentionSetIn) GetTtlConditionsOk() (*[]O11yO11yRetentionRule, bool)`

GetTtlConditionsOk returns a tuple with the TtlConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlConditions

`func (o *O11yO11yRetentionSetIn) SetTtlConditions(v []O11yO11yRetentionRule)`

SetTtlConditions sets TtlConditions field to given value.

### HasTtlConditions

`func (o *O11yO11yRetentionSetIn) HasTtlConditions() bool`

HasTtlConditions returns a boolean if a field has been set.

### GetType

`func (o *O11yO11yRetentionSetIn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yO11yRetentionSetIn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yO11yRetentionSetIn) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yO11yRetentionSetIn) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


