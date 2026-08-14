# PromoStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Promo** | Pointer to [**Promo**](Promo.md) |  | [optional] 
**Redeemed** | Pointer to **int32** | Redeemed is how many orgs have taken it, Remaining how many are left under the fleet-wide cap. | [optional] 
**Remaining** | Pointer to **int32** |  | [optional] 

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

`func (o *PromoStatus) GetRedeemed() int32`

GetRedeemed returns the Redeemed field if non-nil, zero value otherwise.

### GetRedeemedOk

`func (o *PromoStatus) GetRedeemedOk() (*int32, bool)`

GetRedeemedOk returns a tuple with the Redeemed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemed

`func (o *PromoStatus) SetRedeemed(v int32)`

SetRedeemed sets Redeemed field to given value.

### HasRedeemed

`func (o *PromoStatus) HasRedeemed() bool`

HasRedeemed returns a boolean if a field has been set.

### GetRemaining

`func (o *PromoStatus) GetRemaining() int32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *PromoStatus) GetRemainingOk() (*int32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *PromoStatus) SetRemaining(v int32)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *PromoStatus) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


