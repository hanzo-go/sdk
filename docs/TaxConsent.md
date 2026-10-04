# TaxConsent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **int64** | At is when the consent was given or withdrawn, unix seconds. | [optional] 
**By** | Pointer to **string** | By is the IAM user who gave or withdrew it. | [optional] 
**Electronic** | Pointer to **bool** | Electronic is true while the consent stands. Withdrawing it sends every statement furnished afterwards on paper. | [optional] 

## Methods

### NewTaxConsent

`func NewTaxConsent() *TaxConsent`

NewTaxConsent instantiates a new TaxConsent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxConsentWithDefaults

`func NewTaxConsentWithDefaults() *TaxConsent`

NewTaxConsentWithDefaults instantiates a new TaxConsent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *TaxConsent) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *TaxConsent) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *TaxConsent) SetAt(v int64)`

SetAt sets At field to given value.

### HasAt

`func (o *TaxConsent) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBy

`func (o *TaxConsent) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *TaxConsent) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *TaxConsent) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *TaxConsent) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetElectronic

`func (o *TaxConsent) GetElectronic() bool`

GetElectronic returns the Electronic field if non-nil, zero value otherwise.

### GetElectronicOk

`func (o *TaxConsent) GetElectronicOk() (*bool, bool)`

GetElectronicOk returns a tuple with the Electronic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectronic

`func (o *TaxConsent) SetElectronic(v bool)`

SetElectronic sets Electronic field to given value.

### HasElectronic

`func (o *TaxConsent) HasElectronic() bool`

HasElectronic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


