# Quote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeCents** | Pointer to **int64** | ChargeCents is what month one costs after the discount, in USD cents, totalled over the seats quoted. On team that is a multiple of the seat count, so it is not ListCents minus DiscountCents. | [optional] 
**Code** | Pointer to **string** | Code is the promo that was priced, as stored. | [optional] 
**DiscountCents** | Pointer to **int64** | DiscountCents is what the promo takes off month one, in USD cents. The promo rate reaches at most TeamSeatCap seats; seats past the cap bill at full list and add nothing here. It is arithmetic only — quoting credits nothing, counts nothing and reserves nothing. | [optional] 
**Eligible** | Pointer to **bool** | Eligible says whether a redeem would be accepted right now; Reason says why not when it would not. | [optional] 
**ListCents** | Pointer to **int64** | ListCents is the undiscounted month price in USD cents: PER SEAT on team, the whole month on pro and max, 0 for a plan with no list price. | [optional] 
**Plan** | Pointer to **string** | Plan is the tier priced, lower-cased and trimmed: pro, max or team. Unlike a redemption&#39;s plan this one comes from the REQUEST — quoting has no side effects, so it will happily price a plan the caller does not hold. | [optional] 
**Reason** | Pointer to **string** | Reason is why Eligible is false, drawn from: \&quot;promo redemption is closed\&quot; (the subsystem is off, which is how it ships), \&quot;promo redemption cap reached\&quot;, \&quot;promo is not active\&quot;, \&quot;plan is free or unknown; nothing to discount\&quot;, \&quot;promo does not cover plan &lt;plan&gt;\&quot;. Absent when Eligible is true. | [optional] 
**Remaining** | Pointer to **int64** | Remaining is how many redemptions are left under the fleet-wide cap. | [optional] 
**Seats** | Pointer to **int64** | Seats is the seat count priced; a request of 0 or less was read as 1. It only bites on team, the one per-seat plan — pro and max are single-seat and ignore it. | [optional] 

## Methods

### NewQuote

`func NewQuote() *Quote`

NewQuote instantiates a new Quote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteWithDefaults

`func NewQuoteWithDefaults() *Quote`

NewQuoteWithDefaults instantiates a new Quote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeCents

`func (o *Quote) GetChargeCents() int64`

GetChargeCents returns the ChargeCents field if non-nil, zero value otherwise.

### GetChargeCentsOk

`func (o *Quote) GetChargeCentsOk() (*int64, bool)`

GetChargeCentsOk returns a tuple with the ChargeCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeCents

`func (o *Quote) SetChargeCents(v int64)`

SetChargeCents sets ChargeCents field to given value.

### HasChargeCents

`func (o *Quote) HasChargeCents() bool`

HasChargeCents returns a boolean if a field has been set.

### GetCode

`func (o *Quote) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Quote) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Quote) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Quote) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDiscountCents

`func (o *Quote) GetDiscountCents() int64`

GetDiscountCents returns the DiscountCents field if non-nil, zero value otherwise.

### GetDiscountCentsOk

`func (o *Quote) GetDiscountCentsOk() (*int64, bool)`

GetDiscountCentsOk returns a tuple with the DiscountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCents

`func (o *Quote) SetDiscountCents(v int64)`

SetDiscountCents sets DiscountCents field to given value.

### HasDiscountCents

`func (o *Quote) HasDiscountCents() bool`

HasDiscountCents returns a boolean if a field has been set.

### GetEligible

`func (o *Quote) GetEligible() bool`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *Quote) GetEligibleOk() (*bool, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *Quote) SetEligible(v bool)`

SetEligible sets Eligible field to given value.

### HasEligible

`func (o *Quote) HasEligible() bool`

HasEligible returns a boolean if a field has been set.

### GetListCents

`func (o *Quote) GetListCents() int64`

GetListCents returns the ListCents field if non-nil, zero value otherwise.

### GetListCentsOk

`func (o *Quote) GetListCentsOk() (*int64, bool)`

GetListCentsOk returns a tuple with the ListCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListCents

`func (o *Quote) SetListCents(v int64)`

SetListCents sets ListCents field to given value.

### HasListCents

`func (o *Quote) HasListCents() bool`

HasListCents returns a boolean if a field has been set.

### GetPlan

`func (o *Quote) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Quote) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Quote) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *Quote) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetReason

`func (o *Quote) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Quote) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Quote) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Quote) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRemaining

`func (o *Quote) GetRemaining() int64`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *Quote) GetRemainingOk() (*int64, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *Quote) SetRemaining(v int64)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *Quote) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.

### GetSeats

`func (o *Quote) GetSeats() int64`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *Quote) GetSeatsOk() (*int64, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *Quote) SetSeats(v int64)`

SetSeats sets Seats field to given value.

### HasSeats

`func (o *Quote) HasSeats() bool`

HasSeats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


