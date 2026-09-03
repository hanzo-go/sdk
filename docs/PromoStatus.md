# PromoStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Promo** | Pointer to [**Promo**](Promo.md) | Promo is the offer itself. It is fleet-wide, identical for every org — only the two counters beside it move. | [optional] 
**Redeemed** | Pointer to **int64** | Redeemed is how many orgs have taken it, Remaining how many are left under the fleet-wide cap. | [optional] 
**Remaining** | Pointer to **int64** | Remaining is MaxRedemptions minus Redeemed, floored at 0. At 0 the next redeem is declined, and a quote reports ineligible rather than pricing an offer that cannot be taken. | [optional] 

## Methods

### NewPromoStatus

`func NewPromoStatus() *PromoStatus`

NewPromoStatus instantiates a new PromoStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromoStatusWithDefaults

`func NewPromoStatusWithDefaults() *PromoStatus`

NewPromoStatusWithDefaults instantiates a new PromoStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromo

`func (o *PromoStatus) GetPromo() Promo`

GetPromo returns the Promo field if non-nil, zero value otherwise.

### GetPromoOk

`func (o *PromoStatus) GetPromoOk() (*Promo, bool)`

GetPromoOk returns a tuple with the Promo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromo

`func (o *PromoStatus) SetPromo(v Promo)`

SetPromo sets Promo field to given value.

### HasPromo

`func (o *PromoStatus) HasPromo() bool`

HasPromo returns a boolean if a field has been set.

### GetRedeemed

`func (o *PromoStatus) GetRedeemed() int64`

GetRedeemed returns the Redeemed field if non-nil, zero value otherwise.

### GetRedeemedOk

`func (o *PromoStatus) GetRedeemedOk() (*int64, bool)`

GetRedeemedOk returns a tuple with the Redeemed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemed

`func (o *PromoStatus) SetRedeemed(v int64)`

SetRedeemed sets Redeemed field to given value.

### HasRedeemed

`func (o *PromoStatus) HasRedeemed() bool`

HasRedeemed returns a boolean if a field has been set.

### GetRemaining

`func (o *PromoStatus) GetRemaining() int64`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *PromoStatus) GetRemainingOk() (*int64, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *PromoStatus) SetRemaining(v int64)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *PromoStatus) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


