# EmailStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Campaign** | Pointer to **string** | Campaign is the one campaign counted, when the request named one. | [optional] 
**Campaigns** | Pointer to [**[]EmailPart**](EmailPart.md) | Campaigns is the range by campaign, the most sent first, at most 200; empty when the request named a campaign. | [optional] 
**Interval** | Pointer to **string** | Interval is day or hour: the step of Series. | [optional] 
**Provider** | Pointer to **string** | Provider is the one rail counted, when the request named one. | [optional] 
**Providers** | Pointer to [**[]EmailPart**](EmailPart.md) | Providers is the range by provider. | [optional] 
**Series** | Pointer to [**[]EmailPoint**](EmailPoint.md) | Series is every day or hour of the range in order, empty ones included. | [optional] 
**Since** | Pointer to **time.Time** | Since is the start of the range, UTC. | [optional] 
**Totals** | Pointer to [**EmailTally**](EmailTally.md) | Totals is the whole range. | [optional] 
**Until** | Pointer to **time.Time** | Until is the end of the range, UTC, not included. | [optional] 

## Methods

### NewEmailStats

`func NewEmailStats() *EmailStats`

NewEmailStats instantiates a new EmailStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailStatsWithDefaults

`func NewEmailStatsWithDefaults() *EmailStats`

NewEmailStatsWithDefaults instantiates a new EmailStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaign

`func (o *EmailStats) GetCampaign() string`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *EmailStats) GetCampaignOk() (*string, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *EmailStats) SetCampaign(v string)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *EmailStats) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### GetCampaigns

`func (o *EmailStats) GetCampaigns() []EmailPart`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *EmailStats) GetCampaignsOk() (*[]EmailPart, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *EmailStats) SetCampaigns(v []EmailPart)`

SetCampaigns sets Campaigns field to given value.

### HasCampaigns

`func (o *EmailStats) HasCampaigns() bool`

HasCampaigns returns a boolean if a field has been set.

### GetInterval

`func (o *EmailStats) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *EmailStats) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *EmailStats) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *EmailStats) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetProvider

`func (o *EmailStats) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *EmailStats) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *EmailStats) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *EmailStats) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProviders

`func (o *EmailStats) GetProviders() []EmailPart`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *EmailStats) GetProvidersOk() (*[]EmailPart, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *EmailStats) SetProviders(v []EmailPart)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *EmailStats) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetSeries

`func (o *EmailStats) GetSeries() []EmailPoint`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *EmailStats) GetSeriesOk() (*[]EmailPoint, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *EmailStats) SetSeries(v []EmailPoint)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *EmailStats) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetSince

`func (o *EmailStats) GetSince() time.Time`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *EmailStats) GetSinceOk() (*time.Time, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *EmailStats) SetSince(v time.Time)`

SetSince sets Since field to given value.

### HasSince

`func (o *EmailStats) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetTotals

`func (o *EmailStats) GetTotals() EmailTally`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *EmailStats) GetTotalsOk() (*EmailTally, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *EmailStats) SetTotals(v EmailTally)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *EmailStats) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### GetUntil

`func (o *EmailStats) GetUntil() time.Time`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *EmailStats) GetUntilOk() (*time.Time, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *EmailStats) SetUntil(v time.Time)`

SetUntil sets Until field to given value.

### HasUntil

`func (o *EmailStats) HasUntil() bool`

HasUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


