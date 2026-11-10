# MarketingListStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bases** | Pointer to **map[string]int64** | Bases are the subscribers by the consent basis recorded for them, now. | [optional] 
**ConfirmationRate** | Pointer to **float64** | ConfirmationRate is the share of the range&#39;s requests confirmed in it, between 0 and 1; 0 with none requested. | [optional] 
**Counts** | Pointer to **map[string]int64** | Counts are the subscribers by status, now. | [optional] 
**Days** | Pointer to [**[]MarketingDay**](MarketingDay.md) | Days are the range, one entry a day, every day present. | [optional] 
**Growth** | Pointer to **int64** | Growth is the range&#39;s net change: subscribed, confirmed and imported, less unsubscribed, bounced, complained and cleaned. | [optional] 
**List** | Pointer to **string** | List is the list&#39;s id. | [optional] 
**Sources** | Pointer to **map[string]int64** | Sources are the subscribers by where they came from, now. | [optional] 
**UnsubscribeRate** | Pointer to **float64** | UnsubscribeRate is the range&#39;s unsubscribes over the subscribers there were to lose (subscribed now plus those who left in the range), between 0 and 1. | [optional] 

## Methods

### NewMarketingListStats

`func NewMarketingListStats() *MarketingListStats`

NewMarketingListStats instantiates a new MarketingListStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingListStatsWithDefaults

`func NewMarketingListStatsWithDefaults() *MarketingListStats`

NewMarketingListStatsWithDefaults instantiates a new MarketingListStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBases

`func (o *MarketingListStats) GetBases() map[string]int64`

GetBases returns the Bases field if non-nil, zero value otherwise.

### GetBasesOk

`func (o *MarketingListStats) GetBasesOk() (*map[string]int64, bool)`

GetBasesOk returns a tuple with the Bases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBases

`func (o *MarketingListStats) SetBases(v map[string]int64)`

SetBases sets Bases field to given value.

### HasBases

`func (o *MarketingListStats) HasBases() bool`

HasBases returns a boolean if a field has been set.

### GetConfirmationRate

`func (o *MarketingListStats) GetConfirmationRate() float64`

GetConfirmationRate returns the ConfirmationRate field if non-nil, zero value otherwise.

### GetConfirmationRateOk

`func (o *MarketingListStats) GetConfirmationRateOk() (*float64, bool)`

GetConfirmationRateOk returns a tuple with the ConfirmationRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationRate

`func (o *MarketingListStats) SetConfirmationRate(v float64)`

SetConfirmationRate sets ConfirmationRate field to given value.

### HasConfirmationRate

`func (o *MarketingListStats) HasConfirmationRate() bool`

HasConfirmationRate returns a boolean if a field has been set.

### GetCounts

`func (o *MarketingListStats) GetCounts() map[string]int64`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *MarketingListStats) GetCountsOk() (*map[string]int64, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *MarketingListStats) SetCounts(v map[string]int64)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *MarketingListStats) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetDays

`func (o *MarketingListStats) GetDays() []MarketingDay`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *MarketingListStats) GetDaysOk() (*[]MarketingDay, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *MarketingListStats) SetDays(v []MarketingDay)`

SetDays sets Days field to given value.

### HasDays

`func (o *MarketingListStats) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetGrowth

`func (o *MarketingListStats) GetGrowth() int64`

GetGrowth returns the Growth field if non-nil, zero value otherwise.

### GetGrowthOk

`func (o *MarketingListStats) GetGrowthOk() (*int64, bool)`

GetGrowthOk returns a tuple with the Growth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrowth

`func (o *MarketingListStats) SetGrowth(v int64)`

SetGrowth sets Growth field to given value.

### HasGrowth

`func (o *MarketingListStats) HasGrowth() bool`

HasGrowth returns a boolean if a field has been set.

### GetList

`func (o *MarketingListStats) GetList() string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *MarketingListStats) GetListOk() (*string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *MarketingListStats) SetList(v string)`

SetList sets List field to given value.

### HasList

`func (o *MarketingListStats) HasList() bool`

HasList returns a boolean if a field has been set.

### GetSources

`func (o *MarketingListStats) GetSources() map[string]int64`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *MarketingListStats) GetSourcesOk() (*map[string]int64, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *MarketingListStats) SetSources(v map[string]int64)`

SetSources sets Sources field to given value.

### HasSources

`func (o *MarketingListStats) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetUnsubscribeRate

`func (o *MarketingListStats) GetUnsubscribeRate() float64`

GetUnsubscribeRate returns the UnsubscribeRate field if non-nil, zero value otherwise.

### GetUnsubscribeRateOk

`func (o *MarketingListStats) GetUnsubscribeRateOk() (*float64, bool)`

GetUnsubscribeRateOk returns a tuple with the UnsubscribeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribeRate

`func (o *MarketingListStats) SetUnsubscribeRate(v float64)`

SetUnsubscribeRate sets UnsubscribeRate field to given value.

### HasUnsubscribeRate

`func (o *MarketingListStats) HasUnsubscribeRate() bool`

HasUnsubscribeRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


