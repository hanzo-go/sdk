# CloudRedeemResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlreadyRedeemed** | Pointer to **bool** | AlreadyRedeemed is true when this org had already taken the promo and the call was an idempotent replay — nothing was credited a second time. | [optional] 
**ChargeCents** | Pointer to **int32** | ChargeCents is what month one costs after the discount, DiscountCents the credit that produced it. | [optional] 
**DiscountCents** | Pointer to **int32** |  | [optional] 
**Redemption** | Pointer to [**CloudRedemption**](CloudRedemption.md) |  | [optional] 

## Methods

### NewCloudRedeemResult

`func NewCloudRedeemResult() *CloudRedeemResult`

NewCloudRedeemResult instantiates a new CloudRedeemResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRedeemResultWithDefaults

`func NewCloudRedeemResultWithDefaults() *CloudRedeemResult`

NewCloudRedeemResultWithDefaults instantiates a new CloudRedeemResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlreadyRedeemed

`func (o *CloudRedeemResult) GetAlreadyRedeemed() bool`

GetAlreadyRedeemed returns the AlreadyRedeemed field if non-nil, zero value otherwise.

### GetAlreadyRedeemedOk

`func (o *CloudRedeemResult) GetAlreadyRedeemedOk() (*bool, bool)`

GetAlreadyRedeemedOk returns a tuple with the AlreadyRedeemed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyRedeemed

`func (o *CloudRedeemResult) SetAlreadyRedeemed(v bool)`

SetAlreadyRedeemed sets AlreadyRedeemed field to given value.

### HasAlreadyRedeemed

`func (o *CloudRedeemResult) HasAlreadyRedeemed() bool`

HasAlreadyRedeemed returns a boolean if a field has been set.

### GetChargeCents

`func (o *CloudRedeemResult) GetChargeCents() int32`

GetChargeCents returns the ChargeCents field if non-nil, zero value otherwise.

### GetChargeCentsOk

`func (o *CloudRedeemResult) GetChargeCentsOk() (*int32, bool)`

GetChargeCentsOk returns a tuple with the ChargeCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeCents

`func (o *CloudRedeemResult) SetChargeCents(v int32)`

SetChargeCents sets ChargeCents field to given value.

### HasChargeCents

`func (o *CloudRedeemResult) HasChargeCents() bool`

HasChargeCents returns a boolean if a field has been set.

### GetDiscountCents

`func (o *CloudRedeemResult) GetDiscountCents() int32`

GetDiscountCents returns the DiscountCents field if non-nil, zero value otherwise.

### GetDiscountCentsOk

`func (o *CloudRedeemResult) GetDiscountCentsOk() (*int32, bool)`

GetDiscountCentsOk returns a tuple with the DiscountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCents

`func (o *CloudRedeemResult) SetDiscountCents(v int32)`

SetDiscountCents sets DiscountCents field to given value.

### HasDiscountCents

`func (o *CloudRedeemResult) HasDiscountCents() bool`

HasDiscountCents returns a boolean if a field has been set.

### GetRedemption

`func (o *CloudRedeemResult) GetRedemption() CloudRedemption`

GetRedemption returns the Redemption field if non-nil, zero value otherwise.

### GetRedemptionOk

`func (o *CloudRedeemResult) GetRedemptionOk() (*CloudRedemption, bool)`

GetRedemptionOk returns a tuple with the Redemption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemption

`func (o *CloudRedeemResult) SetRedemption(v CloudRedemption)`

SetRedemption sets Redemption field to given value.

### HasRedemption

`func (o *CloudRedeemResult) HasRedemption() bool`

HasRedemption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


