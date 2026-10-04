# MarketingUnsubscribed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address is the recipient now opted out, normalized (lower-cased, trimmed) to the form the send gate matches on — so it can differ in case from the address the link carried. | [optional] 
**Channel** | Pointer to **string** | Channel is the ONE surface opted out of: email, sms, social, meta, google or tiktok. The other channels are untouched, and so is this address in every other org. | [optional] 
**Unsubscribed** | Pointer to **bool** | Unsubscribed is always true here: the opt-out is idempotent, so a second click on the same link confirms the same thing rather than reporting nothing changed. A refused token never reaches this shape — it is a 403. | [optional] 

## Methods

### NewMarketingUnsubscribed

`func NewMarketingUnsubscribed() *MarketingUnsubscribed`

NewMarketingUnsubscribed instantiates a new MarketingUnsubscribed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingUnsubscribedWithDefaults

`func NewMarketingUnsubscribedWithDefaults() *MarketingUnsubscribed`

NewMarketingUnsubscribedWithDefaults instantiates a new MarketingUnsubscribed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MarketingUnsubscribed) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MarketingUnsubscribed) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MarketingUnsubscribed) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MarketingUnsubscribed) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetChannel

`func (o *MarketingUnsubscribed) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *MarketingUnsubscribed) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *MarketingUnsubscribed) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *MarketingUnsubscribed) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *MarketingUnsubscribed) GetUnsubscribed() bool`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *MarketingUnsubscribed) GetUnsubscribedOk() (*bool, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *MarketingUnsubscribed) SetUnsubscribed(v bool)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *MarketingUnsubscribed) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


