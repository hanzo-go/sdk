# TaxProfileIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**TaxAddress**](TaxAddress.md) | Address is W-9 lines 5 and 6, or a W-8&#39;s line 3 permanent residence address outside the United States (state and zip where the country uses them). Required: a statement owed on paper is mailed there. | [optional] 
**BusinessName** | Pointer to **string** | BusinessName is line 2, when it differs from line 1. | [optional] 
**Classification** | Pointer to **string** | Classification is W-9 line 3a: individual, c_corp, s_corp, partnership, trust_estate, llc_c, llc_s or llc_p. Required on a W-9; absent on a W-8. | [optional] 
**ElectronicConsent** | Pointer to **bool** | ElectronicConsent gives (true) or withdraws (false) consent to receive 1099s in the Hanzo inbox, against the disclosure every profile answers with. Omitted, the stored consent stands. | [optional] 
**ExemptPayee** | Pointer to **string** | ExemptPayee is the exempt payee code, 1–13. Entities only. | [optional] 
**Fatca** | Pointer to **string** | FATCA is the FATCA exemption code, A–M. Entities only. | [optional] 
**ForeignOwners** | Pointer to **bool** | ForeignOwners is line 3b; only a partnership, trust/estate or llc_p may set it. | [optional] 
**ForeignTin** | Pointer to **string** | ForeignTIN is a W-8&#39;s foreign tax identifying number (W-8BEN line 6a, W-8BEN-E line 9b), as the jurisdiction of residence issues it. Omitted afterwards, the stored number is kept; the empty string with w8.noForeignTin clears it. | [optional] 
**Form** | Pointer to **string** | Form is w9 (a U.S. person), w8ben (a foreign individual) or w8bene (a foreign entity). w9 when omitted. | [optional] 
**Name** | Pointer to **string** | Name is line 1: the name on the income tax return, or the foreign individual&#39;s or organization&#39;s name. Required. | [optional] 
**Phone** | Pointer to **string** | Phone is printed on the 1099s this org files as payer. | [optional] 
**Tin** | Pointer to **string** | TIN is W-9 Part I, or a W-8&#39;s U.S. TIN (W-8BEN line 5, W-8BEN-E line 8), digits with or without hyphens. Required on a W-9&#39;s first write; on a W-8, only for a treaty claim with no foreign TIN. Omitted afterwards, the stored number is kept, so an address change never re-sends it. | [optional] 
**TinType** | Pointer to **string** | TINType is ssn (an SSN or ITIN) or ein. An entity other than an individual states an EIN. | [optional] 
**W8** | Pointer to [**TaxW8**](TaxW8.md) | W8 is the rest of a W-8: the country of citizenship or incorporation, the chapter 3 and chapter 4 statuses of an entity, a date of birth, the treaty claim, and the signer&#39;s capacity. Required on a W-8; absent on a W-9. | [optional] 

## Methods

### NewTaxProfileIn

`func NewTaxProfileIn() *TaxProfileIn`

NewTaxProfileIn instantiates a new TaxProfileIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxProfileInWithDefaults

`func NewTaxProfileInWithDefaults() *TaxProfileIn`

NewTaxProfileInWithDefaults instantiates a new TaxProfileIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TaxProfileIn) GetAddress() TaxAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TaxProfileIn) GetAddressOk() (*TaxAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TaxProfileIn) SetAddress(v TaxAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TaxProfileIn) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBusinessName

`func (o *TaxProfileIn) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *TaxProfileIn) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *TaxProfileIn) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *TaxProfileIn) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### GetClassification

`func (o *TaxProfileIn) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *TaxProfileIn) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *TaxProfileIn) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *TaxProfileIn) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetElectronicConsent

`func (o *TaxProfileIn) GetElectronicConsent() bool`

GetElectronicConsent returns the ElectronicConsent field if non-nil, zero value otherwise.

### GetElectronicConsentOk

`func (o *TaxProfileIn) GetElectronicConsentOk() (*bool, bool)`

GetElectronicConsentOk returns a tuple with the ElectronicConsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectronicConsent

`func (o *TaxProfileIn) SetElectronicConsent(v bool)`

SetElectronicConsent sets ElectronicConsent field to given value.

### HasElectronicConsent

`func (o *TaxProfileIn) HasElectronicConsent() bool`

HasElectronicConsent returns a boolean if a field has been set.

### GetExemptPayee

`func (o *TaxProfileIn) GetExemptPayee() string`

GetExemptPayee returns the ExemptPayee field if non-nil, zero value otherwise.

### GetExemptPayeeOk

`func (o *TaxProfileIn) GetExemptPayeeOk() (*string, bool)`

GetExemptPayeeOk returns a tuple with the ExemptPayee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptPayee

`func (o *TaxProfileIn) SetExemptPayee(v string)`

SetExemptPayee sets ExemptPayee field to given value.

### HasExemptPayee

`func (o *TaxProfileIn) HasExemptPayee() bool`

HasExemptPayee returns a boolean if a field has been set.

### GetFatca

`func (o *TaxProfileIn) GetFatca() string`

GetFatca returns the Fatca field if non-nil, zero value otherwise.

### GetFatcaOk

`func (o *TaxProfileIn) GetFatcaOk() (*string, bool)`

GetFatcaOk returns a tuple with the Fatca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFatca

`func (o *TaxProfileIn) SetFatca(v string)`

SetFatca sets Fatca field to given value.

### HasFatca

`func (o *TaxProfileIn) HasFatca() bool`

HasFatca returns a boolean if a field has been set.

### GetForeignOwners

`func (o *TaxProfileIn) GetForeignOwners() bool`

GetForeignOwners returns the ForeignOwners field if non-nil, zero value otherwise.

### GetForeignOwnersOk

`func (o *TaxProfileIn) GetForeignOwnersOk() (*bool, bool)`

GetForeignOwnersOk returns a tuple with the ForeignOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignOwners

`func (o *TaxProfileIn) SetForeignOwners(v bool)`

SetForeignOwners sets ForeignOwners field to given value.

### HasForeignOwners

`func (o *TaxProfileIn) HasForeignOwners() bool`

HasForeignOwners returns a boolean if a field has been set.

### GetForeignTin

`func (o *TaxProfileIn) GetForeignTin() string`

GetForeignTin returns the ForeignTin field if non-nil, zero value otherwise.

### GetForeignTinOk

`func (o *TaxProfileIn) GetForeignTinOk() (*string, bool)`

GetForeignTinOk returns a tuple with the ForeignTin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignTin

`func (o *TaxProfileIn) SetForeignTin(v string)`

SetForeignTin sets ForeignTin field to given value.

### HasForeignTin

`func (o *TaxProfileIn) HasForeignTin() bool`

HasForeignTin returns a boolean if a field has been set.

### GetForm

`func (o *TaxProfileIn) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *TaxProfileIn) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *TaxProfileIn) SetForm(v string)`

SetForm sets Form field to given value.

### HasForm

`func (o *TaxProfileIn) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetName

`func (o *TaxProfileIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxProfileIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxProfileIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxProfileIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *TaxProfileIn) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *TaxProfileIn) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *TaxProfileIn) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *TaxProfileIn) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetTin

`func (o *TaxProfileIn) GetTin() string`

GetTin returns the Tin field if non-nil, zero value otherwise.

### GetTinOk

`func (o *TaxProfileIn) GetTinOk() (*string, bool)`

GetTinOk returns a tuple with the Tin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTin

`func (o *TaxProfileIn) SetTin(v string)`

SetTin sets Tin field to given value.

### HasTin

`func (o *TaxProfileIn) HasTin() bool`

HasTin returns a boolean if a field has been set.

### GetTinType

`func (o *TaxProfileIn) GetTinType() string`

GetTinType returns the TinType field if non-nil, zero value otherwise.

### GetTinTypeOk

`func (o *TaxProfileIn) GetTinTypeOk() (*string, bool)`

GetTinTypeOk returns a tuple with the TinType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTinType

`func (o *TaxProfileIn) SetTinType(v string)`

SetTinType sets TinType field to given value.

### HasTinType

`func (o *TaxProfileIn) HasTinType() bool`

HasTinType returns a boolean if a field has been set.

### GetW8

`func (o *TaxProfileIn) GetW8() TaxW8`

GetW8 returns the W8 field if non-nil, zero value otherwise.

### GetW8Ok

`func (o *TaxProfileIn) GetW8Ok() (*TaxW8, bool)`

GetW8Ok returns a tuple with the W8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW8

`func (o *TaxProfileIn) SetW8(v TaxW8)`

SetW8 sets W8 field to given value.

### HasW8

`func (o *TaxProfileIn) HasW8() bool`

HasW8 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


