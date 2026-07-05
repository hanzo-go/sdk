# GuardSanitizeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pii** | Pointer to [**GuardSanitizeConfigPii**](GuardSanitizeConfigPii.md) |  | [optional] 
**Injection** | Pointer to [**GuardSanitizeConfigInjection**](GuardSanitizeConfigInjection.md) |  | [optional] 
**ContentFilter** | Pointer to [**GuardSanitizeConfigContentFilter**](GuardSanitizeConfigContentFilter.md) |  | [optional] 
**RateLimit** | Pointer to [**GuardSanitizeConfigRateLimit**](GuardSanitizeConfigRateLimit.md) |  | [optional] 

## Methods

### NewGuardSanitizeConfig

`func NewGuardSanitizeConfig() *GuardSanitizeConfig`

NewGuardSanitizeConfig instantiates a new GuardSanitizeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardSanitizeConfigWithDefaults

`func NewGuardSanitizeConfigWithDefaults() *GuardSanitizeConfig`

NewGuardSanitizeConfigWithDefaults instantiates a new GuardSanitizeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPii

`func (o *GuardSanitizeConfig) GetPii() GuardSanitizeConfigPii`

GetPii returns the Pii field if non-nil, zero value otherwise.

### GetPiiOk

`func (o *GuardSanitizeConfig) GetPiiOk() (*GuardSanitizeConfigPii, bool)`

GetPiiOk returns a tuple with the Pii field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPii

`func (o *GuardSanitizeConfig) SetPii(v GuardSanitizeConfigPii)`

SetPii sets Pii field to given value.

### HasPii

`func (o *GuardSanitizeConfig) HasPii() bool`

HasPii returns a boolean if a field has been set.

### GetInjection

`func (o *GuardSanitizeConfig) GetInjection() GuardSanitizeConfigInjection`

GetInjection returns the Injection field if non-nil, zero value otherwise.

### GetInjectionOk

`func (o *GuardSanitizeConfig) GetInjectionOk() (*GuardSanitizeConfigInjection, bool)`

GetInjectionOk returns a tuple with the Injection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInjection

`func (o *GuardSanitizeConfig) SetInjection(v GuardSanitizeConfigInjection)`

SetInjection sets Injection field to given value.

### HasInjection

`func (o *GuardSanitizeConfig) HasInjection() bool`

HasInjection returns a boolean if a field has been set.

### GetContentFilter

`func (o *GuardSanitizeConfig) GetContentFilter() GuardSanitizeConfigContentFilter`

GetContentFilter returns the ContentFilter field if non-nil, zero value otherwise.

### GetContentFilterOk

`func (o *GuardSanitizeConfig) GetContentFilterOk() (*GuardSanitizeConfigContentFilter, bool)`

GetContentFilterOk returns a tuple with the ContentFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFilter

`func (o *GuardSanitizeConfig) SetContentFilter(v GuardSanitizeConfigContentFilter)`

SetContentFilter sets ContentFilter field to given value.

### HasContentFilter

`func (o *GuardSanitizeConfig) HasContentFilter() bool`

HasContentFilter returns a boolean if a field has been set.

### GetRateLimit

`func (o *GuardSanitizeConfig) GetRateLimit() GuardSanitizeConfigRateLimit`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *GuardSanitizeConfig) GetRateLimitOk() (*GuardSanitizeConfigRateLimit, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *GuardSanitizeConfig) SetRateLimit(v GuardSanitizeConfigRateLimit)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *GuardSanitizeConfig) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


