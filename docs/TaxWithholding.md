# TaxWithholding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DueCents** | Pointer to **int64** | DueCents is what 24% of the reported boxes comes to — what should have been withheld. Nothing on Hanzo&#39;s rails withholds, so box 4 reports zero, and the payer is liable for the amount it did not withhold (IRC §3403). | [optional] 
**Rate** | Pointer to **int64** | Rate is the backup withholding rate, in percent: 24. | [optional] 
**Reason** | Pointer to **string** | Reason says why, and which rule — or, when nothing is required, what the payer should still know. | [optional] 
**Required** | Pointer to **bool** | Required is true when the payer holds no TIN for the payee — the IRC §3406(a)(1)(A) trigger these forms can know. An uncertified W-9 is not one for these payments, and says so in Reason without requiring anything. | [optional] 

## Methods

### NewTaxWithholding

`func NewTaxWithholding() *TaxWithholding`

NewTaxWithholding instantiates a new TaxWithholding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxWithholdingWithDefaults

`func NewTaxWithholdingWithDefaults() *TaxWithholding`

NewTaxWithholdingWithDefaults instantiates a new TaxWithholding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDueCents

`func (o *TaxWithholding) GetDueCents() int64`

GetDueCents returns the DueCents field if non-nil, zero value otherwise.

### GetDueCentsOk

`func (o *TaxWithholding) GetDueCentsOk() (*int64, bool)`

GetDueCentsOk returns a tuple with the DueCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueCents

`func (o *TaxWithholding) SetDueCents(v int64)`

SetDueCents sets DueCents field to given value.

### HasDueCents

`func (o *TaxWithholding) HasDueCents() bool`

HasDueCents returns a boolean if a field has been set.

### GetRate

`func (o *TaxWithholding) GetRate() int64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *TaxWithholding) GetRateOk() (*int64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *TaxWithholding) SetRate(v int64)`

SetRate sets Rate field to given value.

### HasRate

`func (o *TaxWithholding) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetReason

`func (o *TaxWithholding) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TaxWithholding) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TaxWithholding) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TaxWithholding) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRequired

`func (o *TaxWithholding) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *TaxWithholding) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *TaxWithholding) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *TaxWithholding) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


