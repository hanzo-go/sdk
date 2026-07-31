# CloudRedemption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the promo redeemed. | [optional] 
**CreditCents** | Pointer to **int32** | CreditCents is the discount value credited to the org&#39;s wallet — the promo is realized as a NON-CASH credit, not a subscription coupon. | [optional] 
**CreditEntryId** | Pointer to **string** | CreditEntryID is the finance ledger entry that credit landed in. | [optional] 
**Plan** | Pointer to **string** | Plan and Seats are what was redeemed against. | [optional] 
**RedeemedAt** | Pointer to **int32** | RedeemedAt is unix seconds. | [optional] 
**Seats** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudRedemption

`func NewCloudRedemption() *CloudRedemption`

NewCloudRedemption instantiates a new CloudRedemption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRedemptionWithDefaults

`func NewCloudRedemptionWithDefaults() *CloudRedemption`

NewCloudRedemptionWithDefaults instantiates a new CloudRedemption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CloudRedemption) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CloudRedemption) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CloudRedemption) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CloudRedemption) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreditCents

`func (o *CloudRedemption) GetCreditCents() int32`

GetCreditCents returns the CreditCents field if non-nil, zero value otherwise.

### GetCreditCentsOk

`func (o *CloudRedemption) GetCreditCentsOk() (*int32, bool)`

GetCreditCentsOk returns a tuple with the CreditCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCents

`func (o *CloudRedemption) SetCreditCents(v int32)`

SetCreditCents sets CreditCents field to given value.

### HasCreditCents

`func (o *CloudRedemption) HasCreditCents() bool`

HasCreditCents returns a boolean if a field has been set.

### GetCreditEntryId

`func (o *CloudRedemption) GetCreditEntryId() string`

GetCreditEntryId returns the CreditEntryId field if non-nil, zero value otherwise.

### GetCreditEntryIdOk

`func (o *CloudRedemption) GetCreditEntryIdOk() (*string, bool)`

GetCreditEntryIdOk returns a tuple with the CreditEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditEntryId

`func (o *CloudRedemption) SetCreditEntryId(v string)`

SetCreditEntryId sets CreditEntryId field to given value.

### HasCreditEntryId

`func (o *CloudRedemption) HasCreditEntryId() bool`

HasCreditEntryId returns a boolean if a field has been set.

### GetPlan

`func (o *CloudRedemption) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CloudRedemption) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CloudRedemption) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CloudRedemption) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetRedeemedAt

`func (o *CloudRedemption) GetRedeemedAt() int32`

GetRedeemedAt returns the RedeemedAt field if non-nil, zero value otherwise.

### GetRedeemedAtOk

`func (o *CloudRedemption) GetRedeemedAtOk() (*int32, bool)`

GetRedeemedAtOk returns a tuple with the RedeemedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemedAt

`func (o *CloudRedemption) SetRedeemedAt(v int32)`

SetRedeemedAt sets RedeemedAt field to given value.

### HasRedeemedAt

`func (o *CloudRedemption) HasRedeemedAt() bool`

HasRedeemedAt returns a boolean if a field has been set.

### GetSeats

`func (o *CloudRedemption) GetSeats() int32`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *CloudRedemption) GetSeatsOk() (*int32, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *CloudRedemption) SetSeats(v int32)`

SetSeats sets Seats field to given value.

### HasSeats

`func (o *CloudRedemption) HasSeats() bool`

HasSeats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


