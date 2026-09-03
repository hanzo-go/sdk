# Redemption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the promo redeemed. | [optional] 
**DiscountCents** | Pointer to **int64** | DiscountCents is the month-one discount this redemption CLAIMS, in USD cents. It is a recorded figure, NOT a balance: nothing was credited and no wallet moved. An admin granting against this claim is what would make it money, and that decision happens on the admin surface, not here. | [optional] 
**Plan** | Pointer to **string** | Plan is the tier redeemed against: pro, max or team. It is DERIVED from the org&#39;s live ACTIVE/TRIALING subscription, never read from the request, so it is what the org actually holds rather than what it claimed. | [optional] 
**RedeemedAt** | Pointer to **int64** | RedeemedAt is unix seconds. | [optional] 
**Seats** | Pointer to **int64** | Seats is the seat count the claim was priced at, and it is ALWAYS 1. No server-side authority on this surface answers \&quot;how many seats\&quot;, and the caller&#39;s own number is exactly the input that once inflated these claims, so a redemption records the single-seat floor and an admin resolves the real count against subscription data at grant time. | [optional] 

## Methods

### NewRedemption

`func NewRedemption() *Redemption`

NewRedemption instantiates a new Redemption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedemptionWithDefaults

`func NewRedemptionWithDefaults() *Redemption`

NewRedemptionWithDefaults instantiates a new Redemption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Redemption) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Redemption) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Redemption) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Redemption) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDiscountCents

`func (o *Redemption) GetDiscountCents() int64`

GetDiscountCents returns the DiscountCents field if non-nil, zero value otherwise.

### GetDiscountCentsOk

`func (o *Redemption) GetDiscountCentsOk() (*int64, bool)`

GetDiscountCentsOk returns a tuple with the DiscountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCents

`func (o *Redemption) SetDiscountCents(v int64)`

SetDiscountCents sets DiscountCents field to given value.

### HasDiscountCents

`func (o *Redemption) HasDiscountCents() bool`

HasDiscountCents returns a boolean if a field has been set.

### GetPlan

`func (o *Redemption) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Redemption) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Redemption) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *Redemption) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetRedeemedAt

`func (o *Redemption) GetRedeemedAt() int64`

GetRedeemedAt returns the RedeemedAt field if non-nil, zero value otherwise.

### GetRedeemedAtOk

`func (o *Redemption) GetRedeemedAtOk() (*int64, bool)`

GetRedeemedAtOk returns a tuple with the RedeemedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemedAt

`func (o *Redemption) SetRedeemedAt(v int64)`

SetRedeemedAt sets RedeemedAt field to given value.

### HasRedeemedAt

`func (o *Redemption) HasRedeemedAt() bool`

HasRedeemedAt returns a boolean if a field has been set.

### GetSeats

`func (o *Redemption) GetSeats() int64`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *Redemption) GetSeatsOk() (*int64, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *Redemption) SetSeats(v int64)`

SetSeats sets Seats field to given value.

### HasSeats

`func (o *Redemption) HasSeats() bool`

HasSeats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


