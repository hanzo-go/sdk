# CampaignCampaignResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbTest** | Pointer to **interface{}** |  | [optional] 
**Available** | Pointer to **bool** | Available is false when the analytics warehouse is not connected or the query failed: the funnel below is then zero because nothing could be read, not because nothing happened. Spend and Channels are still real — they come from the connectors, not the warehouse. | [optional] 
**Cac** | Pointer to **float64** | CAC is customer acquisition cost: spend DOLLARS per conversion, rounded to cents. 0 when nothing converted — that is \&quot;not yet computable\&quot;, not \&quot;free\&quot;. | [optional] 
**CampaignId** | Pointer to **string** | CampaignID is the campaign these results are for, echoed from the request. | [optional] 
**Channels** | Pointer to [**[]CampaignChannelMetric**](CampaignChannelMetric.md) | Channels is the per-channel spend breakdown that SpendCents sums, one row per channel on the campaign including the ones that never launched. | [optional] 
**Clicks** | Pointer to **int64** | Clicks is the campaign&#39;s click events over the window. | [optional] 
**Conversions** | Pointer to **int64** | Conversions is the terminal funnel events attributed to the campaign — orders completed, signups completed, explicit conversion events. | [optional] 
**Ctr** | Pointer to **float64** | CTR is clicks per impression, a fraction rounded to 4 places (0.0123 &#x3D; 1.23%), not a percentage. 0 when there were no impressions to divide by. | [optional] 
**Cvr** | Pointer to **float64** | CVR is conversions per click, a fraction rounded to 4 places. 0 when there were no clicks. | [optional] 
**End** | Pointer to **string** | End is the window&#39;s end, RFC3339 UTC — the read&#39;s own clock unless an explicit pair was given. The window is a LOOKBACK, not the campaign&#39;s own lifetime. | [optional] 
**Impressions** | Pointer to **int64** | Impressions is how many times the campaign&#39;s creatives were shown, counted from its utm_campaign-tagged impression events. | [optional] 
**Name** | Pointer to **string** | Name is the campaign&#39;s display name at read time, so a result can be labelled without a second fetch. | [optional] 
**Range** | Pointer to **string** | Range is the window actually used: 24h, 7d, 30d, 90d, or \&quot;custom\&quot; when an explicit start/end pair was honored. An unparseable or absent range reads 30d, so this is the value to trust, not the one that was sent. | [optional] 
**Revenue** | Pointer to **float64** | Revenue is the summed revenue attribute of the campaign&#39;s events, in whole CURRENCY UNITS (dollars) — the one money value here that is not in cents. | [optional] 
**Roas** | Pointer to **float64** | ROAS is return on ad spend: revenue per spend DOLLAR, rounded to 2 places (2.5 &#x3D; $2.50 back per $1). 0 when nothing was spent. | [optional] 
**Source** | Pointer to **string** | Source names the analytics table the funnel was read from, so an operator can see exactly what was counted. Set even when Available is false. | [optional] 
**SpendCents** | Pointer to **int64** | SpendCents is the campaign&#39;s total spend in CENTS: the sum of what each live channel&#39;s provider reports. A channel whose spend could not be read contributes 0 and says so on its own row. | [optional] 
**Start** | Pointer to **string** | Start is the window&#39;s inclusive start, RFC3339 UTC. | [optional] 
**Status** | Pointer to **string** | Status is the campaign&#39;s lifecycle state at read time — draft, live, paused, completed or failed. A draft has never run, so its funnel is legitimately zero. | [optional] 
**Visitors** | Pointer to **int64** | Visitors is how many distinct people the campaign reached, counted by event identity across ALL its events in the window — not a subset of Impressions, so it can exceed them for a campaign whose provider reports clicks but not views. | [optional] 

## Methods

### NewCampaignCampaignResults

`func NewCampaignCampaignResults() *CampaignCampaignResults`

NewCampaignCampaignResults instantiates a new CampaignCampaignResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignCampaignResultsWithDefaults

`func NewCampaignCampaignResultsWithDefaults() *CampaignCampaignResults`

NewCampaignCampaignResultsWithDefaults instantiates a new CampaignCampaignResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbTest

`func (o *CampaignCampaignResults) GetAbTest() interface{}`

GetAbTest returns the AbTest field if non-nil, zero value otherwise.

### GetAbTestOk

`func (o *CampaignCampaignResults) GetAbTestOk() (*interface{}, bool)`

GetAbTestOk returns a tuple with the AbTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbTest

`func (o *CampaignCampaignResults) SetAbTest(v interface{})`

SetAbTest sets AbTest field to given value.

### HasAbTest

`func (o *CampaignCampaignResults) HasAbTest() bool`

HasAbTest returns a boolean if a field has been set.

### SetAbTestNil

`func (o *CampaignCampaignResults) SetAbTestNil(b bool)`

 SetAbTestNil sets the value for AbTest to be an explicit nil

### UnsetAbTest
`func (o *CampaignCampaignResults) UnsetAbTest()`

UnsetAbTest ensures that no value is present for AbTest, not even an explicit nil
### GetAvailable

`func (o *CampaignCampaignResults) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CampaignCampaignResults) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CampaignCampaignResults) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *CampaignCampaignResults) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCac

`func (o *CampaignCampaignResults) GetCac() float64`

GetCac returns the Cac field if non-nil, zero value otherwise.

### GetCacOk

`func (o *CampaignCampaignResults) GetCacOk() (*float64, bool)`

GetCacOk returns a tuple with the Cac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCac

`func (o *CampaignCampaignResults) SetCac(v float64)`

SetCac sets Cac field to given value.

### HasCac

`func (o *CampaignCampaignResults) HasCac() bool`

HasCac returns a boolean if a field has been set.

### GetCampaignId

`func (o *CampaignCampaignResults) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CampaignCampaignResults) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CampaignCampaignResults) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *CampaignCampaignResults) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetChannels

`func (o *CampaignCampaignResults) GetChannels() []CampaignChannelMetric`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *CampaignCampaignResults) GetChannelsOk() (*[]CampaignChannelMetric, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *CampaignCampaignResults) SetChannels(v []CampaignChannelMetric)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *CampaignCampaignResults) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetClicks

`func (o *CampaignCampaignResults) GetClicks() int64`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *CampaignCampaignResults) GetClicksOk() (*int64, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *CampaignCampaignResults) SetClicks(v int64)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *CampaignCampaignResults) HasClicks() bool`

HasClicks returns a boolean if a field has been set.

### GetConversions

`func (o *CampaignCampaignResults) GetConversions() int64`

GetConversions returns the Conversions field if non-nil, zero value otherwise.

### GetConversionsOk

`func (o *CampaignCampaignResults) GetConversionsOk() (*int64, bool)`

GetConversionsOk returns a tuple with the Conversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversions

`func (o *CampaignCampaignResults) SetConversions(v int64)`

SetConversions sets Conversions field to given value.

### HasConversions

`func (o *CampaignCampaignResults) HasConversions() bool`

HasConversions returns a boolean if a field has been set.

### GetCtr

`func (o *CampaignCampaignResults) GetCtr() float64`

GetCtr returns the Ctr field if non-nil, zero value otherwise.

### GetCtrOk

`func (o *CampaignCampaignResults) GetCtrOk() (*float64, bool)`

GetCtrOk returns a tuple with the Ctr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtr

`func (o *CampaignCampaignResults) SetCtr(v float64)`

SetCtr sets Ctr field to given value.

### HasCtr

`func (o *CampaignCampaignResults) HasCtr() bool`

HasCtr returns a boolean if a field has been set.

### GetCvr

`func (o *CampaignCampaignResults) GetCvr() float64`

GetCvr returns the Cvr field if non-nil, zero value otherwise.

### GetCvrOk

`func (o *CampaignCampaignResults) GetCvrOk() (*float64, bool)`

GetCvrOk returns a tuple with the Cvr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvr

`func (o *CampaignCampaignResults) SetCvr(v float64)`

SetCvr sets Cvr field to given value.

### HasCvr

`func (o *CampaignCampaignResults) HasCvr() bool`

HasCvr returns a boolean if a field has been set.

### GetEnd

`func (o *CampaignCampaignResults) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CampaignCampaignResults) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CampaignCampaignResults) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CampaignCampaignResults) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetImpressions

`func (o *CampaignCampaignResults) GetImpressions() int64`

GetImpressions returns the Impressions field if non-nil, zero value otherwise.

### GetImpressionsOk

`func (o *CampaignCampaignResults) GetImpressionsOk() (*int64, bool)`

GetImpressionsOk returns a tuple with the Impressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressions

`func (o *CampaignCampaignResults) SetImpressions(v int64)`

SetImpressions sets Impressions field to given value.

### HasImpressions

`func (o *CampaignCampaignResults) HasImpressions() bool`

HasImpressions returns a boolean if a field has been set.

### GetName

`func (o *CampaignCampaignResults) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignCampaignResults) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignCampaignResults) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignCampaignResults) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRange

`func (o *CampaignCampaignResults) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CampaignCampaignResults) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CampaignCampaignResults) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *CampaignCampaignResults) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRevenue

`func (o *CampaignCampaignResults) GetRevenue() float64`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *CampaignCampaignResults) GetRevenueOk() (*float64, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *CampaignCampaignResults) SetRevenue(v float64)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *CampaignCampaignResults) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetRoas

`func (o *CampaignCampaignResults) GetRoas() float64`

GetRoas returns the Roas field if non-nil, zero value otherwise.

### GetRoasOk

`func (o *CampaignCampaignResults) GetRoasOk() (*float64, bool)`

GetRoasOk returns a tuple with the Roas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoas

`func (o *CampaignCampaignResults) SetRoas(v float64)`

SetRoas sets Roas field to given value.

### HasRoas

`func (o *CampaignCampaignResults) HasRoas() bool`

HasRoas returns a boolean if a field has been set.

### GetSource

`func (o *CampaignCampaignResults) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CampaignCampaignResults) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CampaignCampaignResults) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CampaignCampaignResults) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSpendCents

`func (o *CampaignCampaignResults) GetSpendCents() int64`

GetSpendCents returns the SpendCents field if non-nil, zero value otherwise.

### GetSpendCentsOk

`func (o *CampaignCampaignResults) GetSpendCentsOk() (*int64, bool)`

GetSpendCentsOk returns a tuple with the SpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendCents

`func (o *CampaignCampaignResults) SetSpendCents(v int64)`

SetSpendCents sets SpendCents field to given value.

### HasSpendCents

`func (o *CampaignCampaignResults) HasSpendCents() bool`

HasSpendCents returns a boolean if a field has been set.

### GetStart

`func (o *CampaignCampaignResults) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CampaignCampaignResults) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CampaignCampaignResults) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *CampaignCampaignResults) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStatus

`func (o *CampaignCampaignResults) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CampaignCampaignResults) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CampaignCampaignResults) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CampaignCampaignResults) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVisitors

`func (o *CampaignCampaignResults) GetVisitors() int64`

GetVisitors returns the Visitors field if non-nil, zero value otherwise.

### GetVisitorsOk

`func (o *CampaignCampaignResults) GetVisitorsOk() (*int64, bool)`

GetVisitorsOk returns a tuple with the Visitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitors

`func (o *CampaignCampaignResults) SetVisitors(v int64)`

SetVisitors sets Visitors field to given value.

### HasVisitors

`func (o *CampaignCampaignResults) HasVisitors() bool`

HasVisitors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


