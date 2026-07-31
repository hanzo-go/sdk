# CloudRedeemInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the promo code from the path. | [optional] 
**Instrument** | Pointer to **string** | Instrument identifies the payment method. It is the anti-farming key: one redemption per instrument, fleet-wide. | [optional] 
**Plan** | Pointer to **string** | Plan is the plan being redeemed against: pro, max or team. | [optional] 
**Seats** | Pointer to **int32** | Seats is the Team seat count; 0 means 1. Seats beyond the promo&#39;s teamSeatCap bill at list. | [optional] 

## Methods

### NewCloudRedeemInput

`func NewCloudRedeemInput() *CloudRedeemInput`

NewCloudRedeemInput instantiates a new CloudRedeemInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRedeemInputWithDefaults

`func NewCloudRedeemInputWithDefaults() *CloudRedeemInput`

NewCloudRedeemInputWithDefaults instantiates a new CloudRedeemInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CloudRedeemInput) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CloudRedeemInput) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CloudRedeemInput) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CloudRedeemInput) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetInstrument

`func (o *CloudRedeemInput) GetInstrument() string`

GetInstrument returns the Instrument field if non-nil, zero value otherwise.

### GetInstrumentOk

`func (o *CloudRedeemInput) GetInstrumentOk() (*string, bool)`

GetInstrumentOk returns a tuple with the Instrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrument

`func (o *CloudRedeemInput) SetInstrument(v string)`

SetInstrument sets Instrument field to given value.

### HasInstrument

`func (o *CloudRedeemInput) HasInstrument() bool`

HasInstrument returns a boolean if a field has been set.

### GetPlan

`func (o *CloudRedeemInput) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CloudRedeemInput) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CloudRedeemInput) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *CloudRedeemInput) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSeats

`func (o *CloudRedeemInput) GetSeats() int32`

GetSeats returns the Seats field if non-nil, zero value otherwise.

### GetSeatsOk

`func (o *CloudRedeemInput) GetSeatsOk() (*int32, bool)`

GetSeatsOk returns a tuple with the Seats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeats

`func (o *CloudRedeemInput) SetSeats(v int32)`

SetSeats sets Seats field to given value.

### HasSeats

`func (o *CloudRedeemInput) HasSeats() bool`

HasSeats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


