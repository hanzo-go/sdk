# DomainOffer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | whether it can be bought right now | [optional] 
**Currency** | Pointer to **string** | the currency both prices are in | [optional] 
**Domain** | Pointer to **string** | the name this quote prices | [optional] 
**Premium** | Pointer to **bool** | whether the registry prices it above the standard rate | [optional] 
**PriceCents** | Pointer to **int64** | sell (first-term registration) | [optional] 
**RenewalPriceCents** | Pointer to **int64** | sell (renewal) | [optional] 
**Tld** | Pointer to **string** | the top-level domain the name sits under | [optional] 

## Methods

### NewDomainOffer

`func NewDomainOffer() *DomainOffer`

NewDomainOffer instantiates a new DomainOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainOfferWithDefaults

`func NewDomainOfferWithDefaults() *DomainOffer`

NewDomainOfferWithDefaults instantiates a new DomainOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *DomainOffer) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *DomainOffer) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *DomainOffer) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *DomainOffer) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCurrency

`func (o *DomainOffer) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DomainOffer) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DomainOffer) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DomainOffer) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDomain

`func (o *DomainOffer) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainOffer) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainOffer) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DomainOffer) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetPremium

`func (o *DomainOffer) GetPremium() bool`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *DomainOffer) GetPremiumOk() (*bool, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *DomainOffer) SetPremium(v bool)`

SetPremium sets Premium field to given value.

### HasPremium

`func (o *DomainOffer) HasPremium() bool`

HasPremium returns a boolean if a field has been set.

### GetPriceCents

`func (o *DomainOffer) GetPriceCents() int64`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *DomainOffer) GetPriceCentsOk() (*int64, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *DomainOffer) SetPriceCents(v int64)`

SetPriceCents sets PriceCents field to given value.

### HasPriceCents

`func (o *DomainOffer) HasPriceCents() bool`

HasPriceCents returns a boolean if a field has been set.

### GetRenewalPriceCents

`func (o *DomainOffer) GetRenewalPriceCents() int64`

GetRenewalPriceCents returns the RenewalPriceCents field if non-nil, zero value otherwise.

### GetRenewalPriceCentsOk

`func (o *DomainOffer) GetRenewalPriceCentsOk() (*int64, bool)`

GetRenewalPriceCentsOk returns a tuple with the RenewalPriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalPriceCents

`func (o *DomainOffer) SetRenewalPriceCents(v int64)`

SetRenewalPriceCents sets RenewalPriceCents field to given value.

### HasRenewalPriceCents

`func (o *DomainOffer) HasRenewalPriceCents() bool`

HasRenewalPriceCents returns a boolean if a field has been set.

### GetTld

`func (o *DomainOffer) GetTld() string`

GetTld returns the Tld field if non-nil, zero value otherwise.

### GetTldOk

`func (o *DomainOffer) GetTldOk() (*string, bool)`

GetTldOk returns a tuple with the Tld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTld

`func (o *DomainOffer) SetTld(v string)`

SetTld sets Tld field to given value.

### HasTld

`func (o *DomainOffer) HasTld() bool`

HasTld returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


