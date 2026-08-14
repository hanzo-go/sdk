# RedeemInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is the promo code from the path. | [optional] 
**Instrument** | Pointer to **string** | Instrument identifies the payment method. It is the anti-farming key: one redemption per instrument, fleet-wide, and it is REQUIRED — an absent instrument is refused, never waved through. | [optional] 

## Methods

### NewRedeemInput

`func NewRedeemInput() *RedeemInput`

NewRedeemInput instantiates a new RedeemInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedeemInputWithDefaults

`func NewRedeemInputWithDefaults() *RedeemInput`

NewRedeemInputWithDefaults instantiates a new RedeemInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *RedeemInput) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RedeemInput) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RedeemInput) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *RedeemInput) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetInstrument

`func (o *RedeemInput) GetInstrument() string`

GetInstrument returns the Instrument field if non-nil, zero value otherwise.

### GetInstrumentOk

`func (o *RedeemInput) GetInstrumentOk() (*string, bool)`

GetInstrumentOk returns a tuple with the Instrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrument

`func (o *RedeemInput) SetInstrument(v string)`

SetInstrument sets Instrument field to given value.

### HasInstrument

`func (o *RedeemInput) HasInstrument() bool`

HasInstrument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


