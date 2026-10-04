# MarketplacePrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **interface{}** |  | [optional] 
**Currency** | Pointer to **string** | Currency is the ISO 4217 code, e.g. \&quot;USD\&quot;. Empty means USD. | [optional] 
**Recipient** | Pointer to **string** | Recipient is the payout wallet ref the marketplace seller is paid at. | [optional] 

## Methods

### NewMarketplacePrice

`func NewMarketplacePrice() *MarketplacePrice`

NewMarketplacePrice instantiates a new MarketplacePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplacePriceWithDefaults

`func NewMarketplacePriceWithDefaults() *MarketplacePrice`

NewMarketplacePriceWithDefaults instantiates a new MarketplacePrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *MarketplacePrice) GetAmount() interface{}`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *MarketplacePrice) GetAmountOk() (*interface{}, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *MarketplacePrice) SetAmount(v interface{})`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *MarketplacePrice) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *MarketplacePrice) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *MarketplacePrice) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetCurrency

`func (o *MarketplacePrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MarketplacePrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MarketplacePrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MarketplacePrice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetRecipient

`func (o *MarketplacePrice) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *MarketplacePrice) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *MarketplacePrice) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *MarketplacePrice) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


