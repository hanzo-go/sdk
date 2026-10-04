# MarketplaceEarnings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByRail** | Pointer to **map[string]interface{}** | ByRail sums them by the rail that moved them. | [optional] 
**Currency** | Pointer to **string** | Currency is USD. | [optional] 
**Gross** | Pointer to **interface{}** |  | [optional] 
**Partial** | Pointer to **bool** | Partial is true when the year held more than one read sums. | [optional] 
**Payments** | Pointer to **int64** | Payments is how many. | [optional] 
**Year** | Pointer to **int64** | Year is the calendar year (UTC). | [optional] 

## Methods

### NewMarketplaceEarnings

`func NewMarketplaceEarnings() *MarketplaceEarnings`

NewMarketplaceEarnings instantiates a new MarketplaceEarnings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceEarningsWithDefaults

`func NewMarketplaceEarningsWithDefaults() *MarketplaceEarnings`

NewMarketplaceEarningsWithDefaults instantiates a new MarketplaceEarnings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByRail

`func (o *MarketplaceEarnings) GetByRail() map[string]interface{}`

GetByRail returns the ByRail field if non-nil, zero value otherwise.

### GetByRailOk

`func (o *MarketplaceEarnings) GetByRailOk() (*map[string]interface{}, bool)`

GetByRailOk returns a tuple with the ByRail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByRail

`func (o *MarketplaceEarnings) SetByRail(v map[string]interface{})`

SetByRail sets ByRail field to given value.

### HasByRail

`func (o *MarketplaceEarnings) HasByRail() bool`

HasByRail returns a boolean if a field has been set.

### GetCurrency

`func (o *MarketplaceEarnings) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MarketplaceEarnings) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MarketplaceEarnings) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MarketplaceEarnings) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetGross

`func (o *MarketplaceEarnings) GetGross() interface{}`

GetGross returns the Gross field if non-nil, zero value otherwise.

### GetGrossOk

`func (o *MarketplaceEarnings) GetGrossOk() (*interface{}, bool)`

GetGrossOk returns a tuple with the Gross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGross

`func (o *MarketplaceEarnings) SetGross(v interface{})`

SetGross sets Gross field to given value.

### HasGross

`func (o *MarketplaceEarnings) HasGross() bool`

HasGross returns a boolean if a field has been set.

### SetGrossNil

`func (o *MarketplaceEarnings) SetGrossNil(b bool)`

 SetGrossNil sets the value for Gross to be an explicit nil

### UnsetGross
`func (o *MarketplaceEarnings) UnsetGross()`

UnsetGross ensures that no value is present for Gross, not even an explicit nil
### GetPartial

`func (o *MarketplaceEarnings) GetPartial() bool`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *MarketplaceEarnings) GetPartialOk() (*bool, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *MarketplaceEarnings) SetPartial(v bool)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *MarketplaceEarnings) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetPayments

`func (o *MarketplaceEarnings) GetPayments() int64`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *MarketplaceEarnings) GetPaymentsOk() (*int64, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *MarketplaceEarnings) SetPayments(v int64)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *MarketplaceEarnings) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetYear

`func (o *MarketplaceEarnings) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *MarketplaceEarnings) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *MarketplaceEarnings) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *MarketplaceEarnings) HasYear() bool`

HasYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


