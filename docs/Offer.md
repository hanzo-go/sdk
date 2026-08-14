# Offer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | whether it can be bought right now | [optional] 
**Currency** | Pointer to **string** | the currency both prices are in | [optional] 
**Domain** | Pointer to **string** | the name this quote prices | [optional] 
**Premium** | Pointer to **bool** | whether the registry prices it above the standard rate | [optional] 
**PriceCents** | Pointer to **int32** | sell (first-term registration) | [optional] 
**RenewalPriceCents** | Pointer to **int32** | sell (renewal) | [optional] 
**Tld** | Pointer to **string** | the top-level domain the name sits under | [optional] 

## Methods

### NewOffer

`func NewOffer() *Offer`

NewOffer instantiates a new Offer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferWithDefaults

`func NewOfferWithDefaults() *Offer`

NewOfferWithDefaults instantiates a new Offer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *Offer) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *Offer) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *Offer) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *Offer) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCurrency

`func (o *Offer) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Offer) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Offer) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Offer) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDomain

`func (o *Offer) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Offer) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Offer) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Offer) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPremium

`func (o *Offer) GetPremium() bool`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *Offer) GetPremiumOk() (*bool, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *Offer) SetPremium(v bool)`

SetPremium sets Premium field to given value.

### HasPremium

`func (o *Offer) HasPremium() bool`

HasPremium returns a boolean if a field has been set.

### GetPriceCents

`func (o *Offer) GetPriceCents() int32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *Offer) GetPriceCentsOk() (*int32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *Offer) SetPriceCents(v int32)`

SetPriceCents sets PriceCents field to given value.

### HasPriceCents

`func (o *Offer) HasPriceCents() bool`

HasPriceCents returns a boolean if a field has been set.

### GetRenewalPriceCents

`func (o *Offer) GetRenewalPriceCents() int32`

GetRenewalPriceCents returns the RenewalPriceCents field if non-nil, zero value otherwise.

### GetRenewalPriceCentsOk

`func (o *Offer) GetRenewalPriceCentsOk() (*int32, bool)`

GetRenewalPriceCentsOk returns a tuple with the RenewalPriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPriceCents

`func (o *Offer) SetRenewalPriceCents(v int32)`

SetRenewalPriceCents sets RenewalPriceCents field to given value.

### HasRenewalPriceCents

`func (o *Offer) HasRenewalPriceCents() bool`

HasRenewalPriceCents returns a boolean if a field has been set.

### GetTld

`func (o *Offer) GetTld() string`

GetTld returns the Tld field if non-nil, zero value otherwise.

### GetTldOk

`func (o *Offer) GetTldOk() (*string, bool)`

GetTldOk returns a tuple with the Tld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTld

`func (o *Offer) SetTld(v string)`

SetTld sets Tld field to given value.

### HasTld

`func (o *Offer) HasTld() bool`

HasTld returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


