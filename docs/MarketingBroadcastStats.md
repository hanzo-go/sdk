# MarketingBroadcastStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **int64** | Audience is how many subscribers it selected when it started. | [optional] 
**Bounced** | Pointer to **int64** | Bounced counts messages the provider bounced. | [optional] 
**Broadcast** | Pointer to **string** | Broadcast is the broadcast&#39;s id. | [optional] 
**ClickRate** | Pointer to **float64** | ClickRate is clicked over delivered. | [optional] 
**Clicked** | Pointer to **int64** | Clicked counts messages with at least one click. | [optional] 
**Complained** | Pointer to **int64** | Complained counts messages reported as spam. | [optional] 
**Delivered** | Pointer to **int64** | Delivered, Opened, Clicked, Bounced, Complained and Unsubscribed count messages the email plane reported each of, once per message. | [optional] 
**Failed** | Pointer to **int64** | Failed counts messages that did not go out. | [optional] 
**OpenRate** | Pointer to **float64** | OpenRate and ClickRate are opened and clicked over delivered (over sent before any delivery is reported), between 0 and 1. | [optional] 
**Opened** | Pointer to **int64** | Opened counts messages opened at least once. | [optional] 
**Refused** | Pointer to **int64** | Refused counts recipients the suppression list held. | [optional] 
**Sent** | Pointer to **int64** | Sent counts messages the email capability took. | [optional] 
**Status** | Pointer to **string** | Status is the broadcast&#39;s status. | [optional] 
**Unsubscribed** | Pointer to **int64** | Unsubscribed counts messages whose reader unsubscribed. | [optional] 

## Methods

### NewMarketingBroadcastStats

`func NewMarketingBroadcastStats() *MarketingBroadcastStats`

NewMarketingBroadcastStats instantiates a new MarketingBroadcastStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingBroadcastStatsWithDefaults

`func NewMarketingBroadcastStatsWithDefaults() *MarketingBroadcastStats`

NewMarketingBroadcastStatsWithDefaults instantiates a new MarketingBroadcastStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *MarketingBroadcastStats) GetAudience() int64`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *MarketingBroadcastStats) GetAudienceOk() (*int64, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *MarketingBroadcastStats) SetAudience(v int64)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *MarketingBroadcastStats) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetBounced

`func (o *MarketingBroadcastStats) GetBounced() int64`

GetBounced returns the Bounced field if non-nil, zero value otherwise.

### GetBouncedOk

`func (o *MarketingBroadcastStats) GetBouncedOk() (*int64, bool)`

GetBouncedOk returns a tuple with the Bounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounced

`func (o *MarketingBroadcastStats) SetBounced(v int64)`

SetBounced sets Bounced field to given value.

### HasBounced

`func (o *MarketingBroadcastStats) HasBounced() bool`

HasBounced returns a boolean if a field has been set.

### GetBroadcast

`func (o *MarketingBroadcastStats) GetBroadcast() string`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *MarketingBroadcastStats) GetBroadcastOk() (*string, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *MarketingBroadcastStats) SetBroadcast(v string)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *MarketingBroadcastStats) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.

### GetClickRate

`func (o *MarketingBroadcastStats) GetClickRate() float64`

GetClickRate returns the ClickRate field if non-nil, zero value otherwise.

### GetClickRateOk

`func (o *MarketingBroadcastStats) GetClickRateOk() (*float64, bool)`

GetClickRateOk returns a tuple with the ClickRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickRate

`func (o *MarketingBroadcastStats) SetClickRate(v float64)`

SetClickRate sets ClickRate field to given value.

### HasClickRate

`func (o *MarketingBroadcastStats) HasClickRate() bool`

HasClickRate returns a boolean if a field has been set.

### GetClicked

`func (o *MarketingBroadcastStats) GetClicked() int64`

GetClicked returns the Clicked field if non-nil, zero value otherwise.

### GetClickedOk

`func (o *MarketingBroadcastStats) GetClickedOk() (*int64, bool)`

GetClickedOk returns a tuple with the Clicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicked

`func (o *MarketingBroadcastStats) SetClicked(v int64)`

SetClicked sets Clicked field to given value.

### HasClicked

`func (o *MarketingBroadcastStats) HasClicked() bool`

HasClicked returns a boolean if a field has been set.

### GetComplained

`func (o *MarketingBroadcastStats) GetComplained() int64`

GetComplained returns the Complained field if non-nil, zero value otherwise.

### GetComplainedOk

`func (o *MarketingBroadcastStats) GetComplainedOk() (*int64, bool)`

GetComplainedOk returns a tuple with the Complained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplained

`func (o *MarketingBroadcastStats) SetComplained(v int64)`

SetComplained sets Complained field to given value.

### HasComplained

`func (o *MarketingBroadcastStats) HasComplained() bool`

HasComplained returns a boolean if a field has been set.

### GetDelivered

`func (o *MarketingBroadcastStats) GetDelivered() int64`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *MarketingBroadcastStats) GetDeliveredOk() (*int64, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *MarketingBroadcastStats) SetDelivered(v int64)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *MarketingBroadcastStats) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetFailed

`func (o *MarketingBroadcastStats) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *MarketingBroadcastStats) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *MarketingBroadcastStats) SetFailed(v int64)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *MarketingBroadcastStats) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetOpenRate

`func (o *MarketingBroadcastStats) GetOpenRate() float64`

GetOpenRate returns the OpenRate field if non-nil, zero value otherwise.

### GetOpenRateOk

`func (o *MarketingBroadcastStats) GetOpenRateOk() (*float64, bool)`

GetOpenRateOk returns a tuple with the OpenRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenRate

`func (o *MarketingBroadcastStats) SetOpenRate(v float64)`

SetOpenRate sets OpenRate field to given value.

### HasOpenRate

`func (o *MarketingBroadcastStats) HasOpenRate() bool`

HasOpenRate returns a boolean if a field has been set.

### GetOpened

`func (o *MarketingBroadcastStats) GetOpened() int64`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *MarketingBroadcastStats) GetOpenedOk() (*int64, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *MarketingBroadcastStats) SetOpened(v int64)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *MarketingBroadcastStats) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetRefused

`func (o *MarketingBroadcastStats) GetRefused() int64`

GetRefused returns the Refused field if non-nil, zero value otherwise.

### GetRefusedOk

`func (o *MarketingBroadcastStats) GetRefusedOk() (*int64, bool)`

GetRefusedOk returns a tuple with the Refused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefused

`func (o *MarketingBroadcastStats) SetRefused(v int64)`

SetRefused sets Refused field to given value.

### HasRefused

`func (o *MarketingBroadcastStats) HasRefused() bool`

HasRefused returns a boolean if a field has been set.

### GetSent

`func (o *MarketingBroadcastStats) GetSent() int64`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *MarketingBroadcastStats) GetSentOk() (*int64, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *MarketingBroadcastStats) SetSent(v int64)`

SetSent sets Sent field to given value.

### HasSent

`func (o *MarketingBroadcastStats) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetStatus

`func (o *MarketingBroadcastStats) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketingBroadcastStats) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketingBroadcastStats) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarketingBroadcastStats) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *MarketingBroadcastStats) GetUnsubscribed() int64`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *MarketingBroadcastStats) GetUnsubscribedOk() (*int64, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *MarketingBroadcastStats) SetUnsubscribed(v int64)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *MarketingBroadcastStats) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


