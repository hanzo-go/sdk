# MarketingDay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bounced** | Pointer to **int64** | Bounced counts addresses the provider bounced. | [optional] 
**Cleaned** | Pointer to **int64** | Cleaned counts addresses removed as unmailable. | [optional] 
**Complained** | Pointer to **int64** | Complained counts addresses that reported a mail as spam. | [optional] 
**Confirmed** | Pointer to **int64** | Confirmed counts subscriptions confirmed by their address&#39;s owner. | [optional] 
**Day** | Pointer to **string** | Day is the date, YYYY-MM-DD (UTC). | [optional] 
**Imported** | Pointer to **int64** | Imported counts rows imports added. | [optional] 
**Requested** | Pointer to **int64** | Requested counts subscriptions asked for and waiting to be confirmed. | [optional] 
**Subscribed** | Pointer to **int64** | Subscribed counts subscriptions that began without a confirmation (single opt-in, the API, the preference page). | [optional] 
**Unsubscribed** | Pointer to **int64** | Unsubscribed counts subscribers who left. | [optional] 

## Methods

### NewMarketingDay

`func NewMarketingDay() *MarketingDay`

NewMarketingDay instantiates a new MarketingDay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingDayWithDefaults

`func NewMarketingDayWithDefaults() *MarketingDay`

NewMarketingDayWithDefaults instantiates a new MarketingDay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBounced

`func (o *MarketingDay) GetBounced() int64`

GetBounced returns the Bounced field if non-nil, zero value otherwise.

### GetBouncedOk

`func (o *MarketingDay) GetBouncedOk() (*int64, bool)`

GetBouncedOk returns a tuple with the Bounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounced

`func (o *MarketingDay) SetBounced(v int64)`

SetBounced sets Bounced field to given value.

### HasBounced

`func (o *MarketingDay) HasBounced() bool`

HasBounced returns a boolean if a field has been set.

### GetCleaned

`func (o *MarketingDay) GetCleaned() int64`

GetCleaned returns the Cleaned field if non-nil, zero value otherwise.

### GetCleanedOk

`func (o *MarketingDay) GetCleanedOk() (*int64, bool)`

GetCleanedOk returns a tuple with the Cleaned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleaned

`func (o *MarketingDay) SetCleaned(v int64)`

SetCleaned sets Cleaned field to given value.

### HasCleaned

`func (o *MarketingDay) HasCleaned() bool`

HasCleaned returns a boolean if a field has been set.

### GetComplained

`func (o *MarketingDay) GetComplained() int64`

GetComplained returns the Complained field if non-nil, zero value otherwise.

### GetComplainedOk

`func (o *MarketingDay) GetComplainedOk() (*int64, bool)`

GetComplainedOk returns a tuple with the Complained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplained

`func (o *MarketingDay) SetComplained(v int64)`

SetComplained sets Complained field to given value.

### HasComplained

`func (o *MarketingDay) HasComplained() bool`

HasComplained returns a boolean if a field has been set.

### GetConfirmed

`func (o *MarketingDay) GetConfirmed() int64`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *MarketingDay) GetConfirmedOk() (*int64, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *MarketingDay) SetConfirmed(v int64)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *MarketingDay) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetDay

`func (o *MarketingDay) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *MarketingDay) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *MarketingDay) SetDay(v string)`

SetDay sets Day field to given value.

### HasDay

`func (o *MarketingDay) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetImported

`func (o *MarketingDay) GetImported() int64`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *MarketingDay) GetImportedOk() (*int64, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *MarketingDay) SetImported(v int64)`

SetImported sets Imported field to given value.

### HasImported

`func (o *MarketingDay) HasImported() bool`

HasImported returns a boolean if a field has been set.

### GetRequested

`func (o *MarketingDay) GetRequested() int64`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *MarketingDay) GetRequestedOk() (*int64, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *MarketingDay) SetRequested(v int64)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *MarketingDay) HasRequested() bool`

HasRequested returns a boolean if a field has been set.

### GetSubscribed

`func (o *MarketingDay) GetSubscribed() int64`

GetSubscribed returns the Subscribed field if non-nil, zero value otherwise.

### GetSubscribedOk

`func (o *MarketingDay) GetSubscribedOk() (*int64, bool)`

GetSubscribedOk returns a tuple with the Subscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribed

`func (o *MarketingDay) SetSubscribed(v int64)`

SetSubscribed sets Subscribed field to given value.

### HasSubscribed

`func (o *MarketingDay) HasSubscribed() bool`

HasSubscribed returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *MarketingDay) GetUnsubscribed() int64`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *MarketingDay) GetUnsubscribedOk() (*int64, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *MarketingDay) SetUnsubscribed(v int64)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *MarketingDay) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


