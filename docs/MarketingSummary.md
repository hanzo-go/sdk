# MarketingSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **int64** | Active is how many of them sit in the \&quot;active\&quot; state exactly. A scheduled or paused campaign counts in Campaigns and not here, so Active is never a share of anything but the whole. | [optional] 
**Budget** | Pointer to **int64** | Budget is every campaign&#39;s budget summed, in USD cents. | [optional] 
**Campaigns** | Pointer to **int64** | Campaigns is how many campaigns the org has, counted in every lifecycle state including draft and completed. | [optional] 
**Spend** | Pointer to **int64** | Spend is every campaign&#39;s spend summed, in USD cents. It adds up figures callers wrote on the campaigns themselves — this app meters no delivery against a budget — so it is a reported total, not an observed one. | [optional] 

## Methods

### NewMarketingSummary

`func NewMarketingSummary() *MarketingSummary`

NewMarketingSummary instantiates a new MarketingSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingSummaryWithDefaults

`func NewMarketingSummaryWithDefaults() *MarketingSummary`

NewMarketingSummaryWithDefaults instantiates a new MarketingSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *MarketingSummary) GetActive() int64`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *MarketingSummary) GetActiveOk() (*int64, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *MarketingSummary) SetActive(v int64)`

SetActive sets Active field to given value.

### HasActive

`func (o *MarketingSummary) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBudget

`func (o *MarketingSummary) GetBudget() int64`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *MarketingSummary) GetBudgetOk() (*int64, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *MarketingSummary) SetBudget(v int64)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *MarketingSummary) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetCampaigns

`func (o *MarketingSummary) GetCampaigns() int64`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *MarketingSummary) GetCampaignsOk() (*int64, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *MarketingSummary) SetCampaigns(v int64)`

SetCampaigns sets Campaigns field to given value.

### HasCampaigns

`func (o *MarketingSummary) HasCampaigns() bool`

HasCampaigns returns a boolean if a field has been set.

### GetSpend

`func (o *MarketingSummary) GetSpend() int64`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *MarketingSummary) GetSpendOk() (*int64, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *MarketingSummary) SetSpend(v int64)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *MarketingSummary) HasSpend() bool`

HasSpend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


