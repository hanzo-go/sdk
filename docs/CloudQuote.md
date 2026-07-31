# CloudQuote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeCents** | Pointer to **int32** |  | [optional] 
**Code** | Pointer to **string** | Code, Plan and Seats echo what was quoted. | [optional] 
**DiscountCents** | Pointer to **int32** |  | [optional] 
**Eligible** | Pointer to **bool** | Eligible says whether a redeem would be accepted right now; Reason says why not when it would not. | [optional] 
**ListCents** | Pointer to **int32** | ListCents is the undiscounted month price, ChargeCents what would be charged, DiscountCents the difference — all in USD cents. | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Remaining** | Pointer to **int32** | Remaining is how many redemptions are left under the fleet-wide cap. | [optional] 
**Seats** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudQuote

`func NewCloudQuote() *CloudQuote`

NewCloudQuote instantiates a new CloudQuote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudQuoteWithDefaults

`func NewCloudQuoteWithDefaults() *CloudQuote`

NewCloudQuoteWithDefaults instantiates a new CloudQuote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeCents

`func (o *CloudQuote) GetChargeCents() int32`

GetChargeCents returns the ChargeCents field if non-nil, zero value otherwise.

### GetChargeCentsOk

`func (o *CloudQuote) GetChargeCentsOk() (*int32, bool)`

GetChargeCentsOk returns a tuple with the ChargeCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeCents

`func (o *CloudQuote) SetChargeCents(v int32)`

SetChargeCents sets ChargeCents field to given value.

### HasChargeCents

`func (o *CloudQuote) HasChargeCents() bool`

HasChargeCents returns a boolean if a field has been set.

### GetCode

`func (o *CloudQuote) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CloudQuote) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CloudQuote) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CloudQuote) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDiscountCents

`func (o *CloudQuote) GetDiscountCents() int32`

GetDiscountCents returns the DiscountCents field if non-nil, zero value otherwise.

### GetDiscountCentsOk

`func (o *CloudQuote) GetDiscountCentsOk() (*int32, bool)`

GetDiscountCentsOk returns a tuple with the DiscountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCents

`func (o *CloudQuote) SetDiscountCents(v int32)`

SetDiscountCents sets DiscountCents field to given value.

### HasDiscountCents

`func (o *CloudQuote) HasDiscountCents() bool`

HasDiscountCents returns a boolean if a field has been set.

### GetEligible

`func (o *CloudQuote) GetEligible() bool`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *CloudQuote) GetEligibleOk() (*bool, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *CloudQuote) SetEligible(v bool)`

SetEligible sets Eligible field to given value.

### HasEligible

`func (o *CloudQuote) HasEligible() bool`

HasEligible returns a boolean if a field has been set.

### GetListCents

`func (o *CloudQuote) GetListCents() int32`

GetListCents returns the ListCents field if non-nil, zero value otherwise.

### GetListCentsOk

`func (o *CloudQuote) GetListCentsOk() (*int32, bool)`

GetListCentsOk returns a tuple with the ListCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListCents

`func (o *CloudQuote) SetListCents(v int32)`

SetListCents sets ListCents field to given value.

### HasListCents

`func (o *CloudQuote) HasListCents() bool`

HasListCents returns a boolean if a field has been set.

### GetPlan

`func (o *CloudQuote) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CloudQuote) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CloudQuote) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CloudQuote) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetReason

`func (o *CloudQuote) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CloudQuote) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CloudQuote) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CloudQuote) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRemaining

`func (o *CloudQuote) GetRemaining() int32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *CloudQuote) GetRemainingOk() (*int32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *CloudQuote) SetRemaining(v int32)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *CloudQuote) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.

### GetSeats

`func (o *CloudQuote) GetSeats() int32`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *CloudQuote) GetSeatsOk() (*int32, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *CloudQuote) SetSeats(v int32)`

SetSeats sets Seats field to given value.

### HasSeats

`func (o *CloudQuote) HasSeats() bool`

HasSeats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


