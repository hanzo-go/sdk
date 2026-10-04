# TaxLedger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payees** | Pointer to [**[]TaxPayeeYear**](TaxPayeeYear.md) | Payees is every org the caller paid, by org id. | [optional] 
**Source** | Pointer to **string** | Source names where the year&#39;s numbers were read. | [optional] 
**ThresholdCents** | Pointer to **int64** | ThresholdCents is the §6041(a) threshold for the year. | [optional] 
**Year** | Pointer to **int64** | Year is the tax year. | [optional] 

## Methods

### NewTaxLedger

`func NewTaxLedger() *TaxLedger`

NewTaxLedger instantiates a new TaxLedger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxLedgerWithDefaults

`func NewTaxLedgerWithDefaults() *TaxLedger`

NewTaxLedgerWithDefaults instantiates a new TaxLedger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayees

`func (o *TaxLedger) GetPayees() []TaxPayeeYear`

GetPayees returns the Payees field if non-nil, zero value otherwise.

### GetPayeesOk

`func (o *TaxLedger) GetPayeesOk() (*[]TaxPayeeYear, bool)`

GetPayeesOk returns a tuple with the Payees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayees

`func (o *TaxLedger) SetPayees(v []TaxPayeeYear)`

SetPayees sets Payees field to given value.

### HasPayees

`func (o *TaxLedger) HasPayees() bool`

HasPayees returns a boolean if a field has been set.

### GetSource

`func (o *TaxLedger) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TaxLedger) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TaxLedger) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TaxLedger) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetThresholdCents

`func (o *TaxLedger) GetThresholdCents() int64`

GetThresholdCents returns the ThresholdCents field if non-nil, zero value otherwise.

### GetThresholdCentsOk

`func (o *TaxLedger) GetThresholdCentsOk() (*int64, bool)`

GetThresholdCentsOk returns a tuple with the ThresholdCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdCents

`func (o *TaxLedger) SetThresholdCents(v int64)`

SetThresholdCents sets ThresholdCents field to given value.

### HasThresholdCents

`func (o *TaxLedger) HasThresholdCents() bool`

HasThresholdCents returns a boolean if a field has been set.

### GetYear

`func (o *TaxLedger) GetYear() int64`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *TaxLedger) GetYearOk() (*int64, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *TaxLedger) SetYear(v int64)`

SetYear sets Year field to given value.

### HasYear

`func (o *TaxLedger) HasYear() bool`

HasYear returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


