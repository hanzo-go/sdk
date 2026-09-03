# RedeemResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlreadyRedeemed** | Pointer to **bool** | AlreadyRedeemed is true when this org had already taken the promo and the call was an idempotent replay. | [optional] 
**ChargeCents** | Pointer to **int64** | ChargeCents is what month one costs after the discount, DiscountCents the discount that produced it. Both are quoted figures against the org&#39;s derived plan — NOTHING WAS CREDITED and no wallet moved. | [optional] 
**DiscountCents** | Pointer to **int64** | DiscountCents is the discount claimed for month one, in USD cents, at the single-seat floor. It is the same figure recorded on the Redemption, and it is evidence an admin may later grant against — not a balance. | [optional] 
**Redemption** | Pointer to [**Redemption**](Redemption.md) | Redemption is the row that was recorded — the org&#39;s claim on this promo, with the server-derived plan and seat count. On a replay it is the ORIGINAL row, so its redeemedAt is when the org first took the promo, not now. | [optional] 

## Methods

### NewRedeemResult

`func NewRedeemResult() *RedeemResult`

NewRedeemResult instantiates a new RedeemResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedeemResultWithDefaults

`func NewRedeemResultWithDefaults() *RedeemResult`

NewRedeemResultWithDefaults instantiates a new RedeemResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlreadyRedeemed

`func (o *RedeemResult) GetAlreadyRedeemed() bool`

GetAlreadyRedeemed returns the AlreadyRedeemed field if non-nil, zero value otherwise.

### GetAlreadyRedeemedOk

`func (o *RedeemResult) GetAlreadyRedeemedOk() (*bool, bool)`

GetAlreadyRedeemedOk returns a tuple with the AlreadyRedeemed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyRedeemed

`func (o *RedeemResult) SetAlreadyRedeemed(v bool)`

SetAlreadyRedeemed sets AlreadyRedeemed field to given value.

### HasAlreadyRedeemed

`func (o *RedeemResult) HasAlreadyRedeemed() bool`

HasAlreadyRedeemed returns a boolean if a field has been set.

### GetChargeCents

`func (o *RedeemResult) GetChargeCents() int64`

GetChargeCents returns the ChargeCents field if non-nil, zero value otherwise.

### GetChargeCentsOk

`func (o *RedeemResult) GetChargeCentsOk() (*int64, bool)`

GetChargeCentsOk returns a tuple with the ChargeCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeCents

`func (o *RedeemResult) SetChargeCents(v int64)`

SetChargeCents sets ChargeCents field to given value.

### HasChargeCents

`func (o *RedeemResult) HasChargeCents() bool`

HasChargeCents returns a boolean if a field has been set.

### GetDiscountCents

`func (o *RedeemResult) GetDiscountCents() int64`

GetDiscountCents returns the DiscountCents field if non-nil, zero value otherwise.

### GetDiscountCentsOk

`func (o *RedeemResult) GetDiscountCentsOk() (*int64, bool)`

GetDiscountCentsOk returns a tuple with the DiscountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCents

`func (o *RedeemResult) SetDiscountCents(v int64)`

SetDiscountCents sets DiscountCents field to given value.

### HasDiscountCents

`func (o *RedeemResult) HasDiscountCents() bool`

HasDiscountCents returns a boolean if a field has been set.

### GetRedemption

`func (o *RedeemResult) GetRedemption() Redemption`

GetRedemption returns the Redemption field if non-nil, zero value otherwise.

### GetRedemptionOk

`func (o *RedeemResult) GetRedemptionOk() (*Redemption, bool)`

GetRedemptionOk returns a tuple with the Redemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemption

`func (o *RedeemResult) SetRedemption(v Redemption)`

SetRedemption sets Redemption field to given value.

### HasRedemption

`func (o *RedeemResult) HasRedemption() bool`

HasRedemption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


