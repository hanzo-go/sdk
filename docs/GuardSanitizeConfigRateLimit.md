# GuardSanitizeConfigRateLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**RequestsPerMinute** | Pointer to **int32** |  | [optional] [default to 60]
**TokensPerMinute** | Pointer to **int32** |  | [optional] [default to 100000]
**BurstSize** | Pointer to **int32** |  | [optional] [default to 10]

## Methods

### NewGuardSanitizeConfigRateLimit

`func NewGuardSanitizeConfigRateLimit() *GuardSanitizeConfigRateLimit`

NewGuardSanitizeConfigRateLimit instantiates a new GuardSanitizeConfigRateLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardSanitizeConfigRateLimitWithDefaults

`func NewGuardSanitizeConfigRateLimitWithDefaults() *GuardSanitizeConfigRateLimit`

NewGuardSanitizeConfigRateLimitWithDefaults instantiates a new GuardSanitizeConfigRateLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GuardSanitizeConfigRateLimit) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GuardSanitizeConfigRateLimit) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GuardSanitizeConfigRateLimit) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GuardSanitizeConfigRateLimit) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRequestsPerMinute

`func (o *GuardSanitizeConfigRateLimit) GetRequestsPerMinute() int32`

GetRequestsPerMinute returns the RequestsPerMinute field if non-nil, zero value otherwise.

### GetRequestsPerMinuteOk

`func (o *GuardSanitizeConfigRateLimit) GetRequestsPerMinuteOk() (*int32, bool)`

GetRequestsPerMinuteOk returns a tuple with the RequestsPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsPerMinute

`func (o *GuardSanitizeConfigRateLimit) SetRequestsPerMinute(v int32)`

SetRequestsPerMinute sets RequestsPerMinute field to given value.

### HasRequestsPerMinute

`func (o *GuardSanitizeConfigRateLimit) HasRequestsPerMinute() bool`

HasRequestsPerMinute returns a boolean if a field has been set.

### GetTokensPerMinute

`func (o *GuardSanitizeConfigRateLimit) GetTokensPerMinute() int32`

GetTokensPerMinute returns the TokensPerMinute field if non-nil, zero value otherwise.

### GetTokensPerMinuteOk

`func (o *GuardSanitizeConfigRateLimit) GetTokensPerMinuteOk() (*int32, bool)`

GetTokensPerMinuteOk returns a tuple with the TokensPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensPerMinute

`func (o *GuardSanitizeConfigRateLimit) SetTokensPerMinute(v int32)`

SetTokensPerMinute sets TokensPerMinute field to given value.

### HasTokensPerMinute

`func (o *GuardSanitizeConfigRateLimit) HasTokensPerMinute() bool`

HasTokensPerMinute returns a boolean if a field has been set.

### GetBurstSize

`func (o *GuardSanitizeConfigRateLimit) GetBurstSize() int32`

GetBurstSize returns the BurstSize field if non-nil, zero value otherwise.

### GetBurstSizeOk

`func (o *GuardSanitizeConfigRateLimit) GetBurstSizeOk() (*int32, bool)`

GetBurstSizeOk returns a tuple with the BurstSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstSize

`func (o *GuardSanitizeConfigRateLimit) SetBurstSize(v int32)`

SetBurstSize sets BurstSize field to given value.

### HasBurstSize

`func (o *GuardSanitizeConfigRateLimit) HasBurstSize() bool`

HasBurstSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


