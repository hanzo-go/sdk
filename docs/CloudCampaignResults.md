# CloudCampaignResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbTest** | Pointer to **interface{}** |  | [optional] 
**Available** | Pointer to **bool** |  | [optional] 
**Cac** | Pointer to **float32** |  | [optional] 
**CampaignId** | Pointer to **string** |  | [optional] 
**Channels** | Pointer to [**[]CloudChannelMetric**](CloudChannelMetric.md) |  | [optional] 
**Clicks** | Pointer to **int32** |  | [optional] 
**Conversions** | Pointer to **int32** |  | [optional] 
**Ctr** | Pointer to **float32** |  | [optional] 
**Cvr** | Pointer to **float32** |  | [optional] 
**End** | Pointer to **string** |  | [optional] 
**Impressions** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Range** | Pointer to **string** |  | [optional] 
**Revenue** | Pointer to **float32** |  | [optional] 
**Roas** | Pointer to **float32** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**SpendCents** | Pointer to **int32** |  | [optional] 
**Start** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Visitors** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudCampaignResults

`func NewCloudCampaignResults() *CloudCampaignResults`

NewCloudCampaignResults instantiates a new CloudCampaignResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCampaignResultsWithDefaults

`func NewCloudCampaignResultsWithDefaults() *CloudCampaignResults`

NewCloudCampaignResultsWithDefaults instantiates a new CloudCampaignResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbTest

`func (o *CloudCampaignResults) GetAbTest() interface{}`

GetAbTest returns the AbTest field if non-nil, zero value otherwise.

### GetAbTestOk

`func (o *CloudCampaignResults) GetAbTestOk() (*interface{}, bool)`

GetAbTestOk returns a tuple with the AbTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbTest

`func (o *CloudCampaignResults) SetAbTest(v interface{})`

SetAbTest sets AbTest field to given value.

### HasAbTest

`func (o *CloudCampaignResults) HasAbTest() bool`

HasAbTest returns a boolean if a field has been set.

### SetAbTestNil

`func (o *CloudCampaignResults) SetAbTestNil(b bool)`

 SetAbTestNil sets the value for AbTest to be an explicit nil

### UnsetAbTest
`func (o *CloudCampaignResults) UnsetAbTest()`

UnsetAbTest ensures that no value is present for AbTest, not even an explicit nil
### GetAvailable

`func (o *CloudCampaignResults) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CloudCampaignResults) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CloudCampaignResults) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *CloudCampaignResults) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCac

`func (o *CloudCampaignResults) GetCac() float32`

GetCac returns the Cac field if non-nil, zero value otherwise.

### GetCacOk

`func (o *CloudCampaignResults) GetCacOk() (*float32, bool)`

GetCacOk returns a tuple with the Cac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCac

`func (o *CloudCampaignResults) SetCac(v float32)`

SetCac sets Cac field to given value.

### HasCac

`func (o *CloudCampaignResults) HasCac() bool`

HasCac returns a boolean if a field has been set.

### GetCampaignId

`func (o *CloudCampaignResults) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CloudCampaignResults) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CloudCampaignResults) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *CloudCampaignResults) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetChannels

`func (o *CloudCampaignResults) GetChannels() []CloudChannelMetric`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *CloudCampaignResults) GetChannelsOk() (*[]CloudChannelMetric, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *CloudCampaignResults) SetChannels(v []CloudChannelMetric)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *CloudCampaignResults) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetClicks

`func (o *CloudCampaignResults) GetClicks() int32`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *CloudCampaignResults) GetClicksOk() (*int32, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *CloudCampaignResults) SetClicks(v int32)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *CloudCampaignResults) HasClicks() bool`

HasClicks returns a boolean if a field has been set.

### GetConversions

`func (o *CloudCampaignResults) GetConversions() int32`

GetConversions returns the Conversions field if non-nil, zero value otherwise.

### GetConversionsOk

`func (o *CloudCampaignResults) GetConversionsOk() (*int32, bool)`

GetConversionsOk returns a tuple with the Conversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversions

`func (o *CloudCampaignResults) SetConversions(v int32)`

SetConversions sets Conversions field to given value.

### HasConversions

`func (o *CloudCampaignResults) HasConversions() bool`

HasConversions returns a boolean if a field has been set.

### GetCtr

`func (o *CloudCampaignResults) GetCtr() float32`

GetCtr returns the Ctr field if non-nil, zero value otherwise.

### GetCtrOk

`func (o *CloudCampaignResults) GetCtrOk() (*float32, bool)`

GetCtrOk returns a tuple with the Ctr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtr

`func (o *CloudCampaignResults) SetCtr(v float32)`

SetCtr sets Ctr field to given value.

### HasCtr

`func (o *CloudCampaignResults) HasCtr() bool`

HasCtr returns a boolean if a field has been set.

### GetCvr

`func (o *CloudCampaignResults) GetCvr() float32`

GetCvr returns the Cvr field if non-nil, zero value otherwise.

### GetCvrOk

`func (o *CloudCampaignResults) GetCvrOk() (*float32, bool)`

GetCvrOk returns a tuple with the Cvr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvr

`func (o *CloudCampaignResults) SetCvr(v float32)`

SetCvr sets Cvr field to given value.

### HasCvr

`func (o *CloudCampaignResults) HasCvr() bool`

HasCvr returns a boolean if a field has been set.

### GetEnd

`func (o *CloudCampaignResults) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CloudCampaignResults) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CloudCampaignResults) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CloudCampaignResults) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetImpressions

`func (o *CloudCampaignResults) GetImpressions() int32`

GetImpressions returns the Impressions field if non-nil, zero value otherwise.

### GetImpressionsOk

`func (o *CloudCampaignResults) GetImpressionsOk() (*int32, bool)`

GetImpressionsOk returns a tuple with the Impressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressions

`func (o *CloudCampaignResults) SetImpressions(v int32)`

SetImpressions sets Impressions field to given value.

### HasImpressions

`func (o *CloudCampaignResults) HasImpressions() bool`

HasImpressions returns a boolean if a field has been set.

### GetName

`func (o *CloudCampaignResults) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCampaignResults) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCampaignResults) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCampaignResults) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRange

`func (o *CloudCampaignResults) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CloudCampaignResults) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CloudCampaignResults) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *CloudCampaignResults) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRevenue

`func (o *CloudCampaignResults) GetRevenue() float32`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *CloudCampaignResults) GetRevenueOk() (*float32, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *CloudCampaignResults) SetRevenue(v float32)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *CloudCampaignResults) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetRoas

`func (o *CloudCampaignResults) GetRoas() float32`

GetRoas returns the Roas field if non-nil, zero value otherwise.

### GetRoasOk

`func (o *CloudCampaignResults) GetRoasOk() (*float32, bool)`

GetRoasOk returns a tuple with the Roas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoas

`func (o *CloudCampaignResults) SetRoas(v float32)`

SetRoas sets Roas field to given value.

### HasRoas

`func (o *CloudCampaignResults) HasRoas() bool`

HasRoas returns a boolean if a field has been set.

### GetSource

`func (o *CloudCampaignResults) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudCampaignResults) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudCampaignResults) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudCampaignResults) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSpendCents

`func (o *CloudCampaignResults) GetSpendCents() int32`

GetSpendCents returns the SpendCents field if non-nil, zero value otherwise.

### GetSpendCentsOk

`func (o *CloudCampaignResults) GetSpendCentsOk() (*int32, bool)`

GetSpendCentsOk returns a tuple with the SpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendCents

`func (o *CloudCampaignResults) SetSpendCents(v int32)`

SetSpendCents sets SpendCents field to given value.

### HasSpendCents

`func (o *CloudCampaignResults) HasSpendCents() bool`

HasSpendCents returns a boolean if a field has been set.

### GetStart

`func (o *CloudCampaignResults) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CloudCampaignResults) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CloudCampaignResults) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *CloudCampaignResults) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStatus

`func (o *CloudCampaignResults) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudCampaignResults) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudCampaignResults) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudCampaignResults) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVisitors

`func (o *CloudCampaignResults) GetVisitors() int32`

GetVisitors returns the Visitors field if non-nil, zero value otherwise.

### GetVisitorsOk

`func (o *CloudCampaignResults) GetVisitorsOk() (*int32, bool)`

GetVisitorsOk returns a tuple with the Visitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitors

`func (o *CloudCampaignResults) SetVisitors(v int32)`

SetVisitors sets Visitors field to given value.

### HasVisitors

`func (o *CloudCampaignResults) HasVisitors() bool`

HasVisitors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


