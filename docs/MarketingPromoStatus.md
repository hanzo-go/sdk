# MarketingPromoStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Promo** | Pointer to [**MarketingPromo**](MarketingPromo.md) | Promo is the offer itself. It is fleet-wide, identical for every org — only the two counters beside it move. | [optional] 
**Redeemed** | Pointer to **int64** | Redeemed is how many orgs have taken it, Remaining how many are left under the fleet-wide cap. | [optional] 
**Remaining** | Pointer to **int64** | Remaining is MaxRedemptions minus Redeemed, floored at 0. At 0 the next redeem is declined, and a quote reports ineligible rather than pricing an offer that cannot be taken. | [optional] 

## Methods

### NewMarketingPromoStatus

`func NewMarketingPromoStatus() *MarketingPromoStatus`

NewMarketingPromoStatus instantiates a new MarketingPromoStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingPromoStatusWithDefaults

`func NewMarketingPromoStatusWithDefaults() *MarketingPromoStatus`

NewMarketingPromoStatusWithDefaults instantiates a new MarketingPromoStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromo

`func (o *MarketingPromoStatus) GetPromo() MarketingPromo`

GetPromo returns the Promo field if non-nil, zero value otherwise.

### GetPromoOk

`func (o *MarketingPromoStatus) GetPromoOk() (*MarketingPromo, bool)`

GetPromoOk returns a tuple with the Promo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromo

`func (o *MarketingPromoStatus) SetPromo(v MarketingPromo)`

SetPromo sets Promo field to given value.

### HasPromo

`func (o *MarketingPromoStatus) HasPromo() bool`

HasPromo returns a boolean if a field has been set.

### GetRedeemed

`func (o *MarketingPromoStatus) GetRedeemed() int64`

GetRedeemed returns the Redeemed field if non-nil, zero value otherwise.

### GetRedeemedOk

`func (o *MarketingPromoStatus) GetRedeemedOk() (*int64, bool)`

GetRedeemedOk returns a tuple with the Redeemed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemed

`func (o *MarketingPromoStatus) SetRedeemed(v int64)`

SetRedeemed sets Redeemed field to given value.

### HasRedeemed

`func (o *MarketingPromoStatus) HasRedeemed() bool`

HasRedeemed returns a boolean if a field has been set.

### GetRemaining

`func (o *MarketingPromoStatus) GetRemaining() int64`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *MarketingPromoStatus) GetRemainingOk() (*int64, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *MarketingPromoStatus) SetRemaining(v int64)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *MarketingPromoStatus) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


