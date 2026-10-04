# TaxProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**TaxAddress**](TaxAddress.md) | Address is lines 5 and 6. | [optional] 
**BusinessName** | Pointer to **string** | BusinessName is line 2: the business or disregarded entity name. | [optional] 
**Certification** | Pointer to [**TaxCertification**](TaxCertification.md) | Certification is where the Part II signature stands. | [optional] 
**Classification** | Pointer to **string** | Classification is line 3a. | [optional] 
**Consent** | Pointer to [**TaxConsent**](TaxConsent.md) | Consent is the electronic-delivery consent. Only the org itself reads it: how a payee takes delivery is not a payer&#39;s to know. | [optional] 
**Disclosure** | Pointer to **string** | Disclosure is the electronic-delivery disclosure the consent is given against. | [optional] 
**ExemptPayee** | Pointer to **string** | ExemptPayee is line 4&#39;s exempt payee code, 1–13. | [optional] 
**Expires** | Pointer to **int64** | Expires is when a certified W-8 stops being valid, unix seconds: the last day of the third calendar year after the year it was signed. A W-9 does not expire. | [optional] 
**Fatca** | Pointer to **string** | FATCA is line 4&#39;s FATCA exemption code, A–M. | [optional] 
**ForeignOwners** | Pointer to **bool** | ForeignOwners is line 3b: a flow-through entity with foreign partners, owners or beneficiaries. | [optional] 
**ForeignTin** | Pointer to **string** | ForeignTIN is a W-8&#39;s foreign tax identifying number, masked. The full number is answered only with the full TIN read. | [optional] 
**Form** | Pointer to **string** | Form is w9, w8ben or w8bene. | [optional] 
**Name** | Pointer to **string** | Name is line 1: the name on the income tax return, or of the foreign individual or organization that is the beneficial owner. | [optional] 
**Phone** | Pointer to **string** | Phone is the number a 1099 this org files as PAYER prints for it. | [optional] 
**Tin** | Pointer to **string** | TIN is Part I, masked to its last four digits. The full number is answered only by GET /v1/tax/w9/{id}/tin, to a payer this org granted it to. | [optional] 
**TinType** | Pointer to **string** | TINType is which Part I box: ssn (an SSN or ITIN) or ein. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when the profile last changed, unix seconds. | [optional] 
**Valid** | Pointer to **bool** | Valid is whether the form establishes what it certifies today — for a W-8, certified for its current version and not expired; for a W-9, a TIN on file. | [optional] 
**Version** | Pointer to **int64** | Version counts the W-9&#39;s revisions; a certification covers exactly one. | [optional] 
**W8** | Pointer to [**TaxW8**](TaxW8.md) | W8 is what a W-8BEN or W-8BEN-E certifies beyond name and address. | [optional] 

## Methods

### NewTaxProfile

`func NewTaxProfile() *TaxProfile`

NewTaxProfile instantiates a new TaxProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxProfileWithDefaults

`func NewTaxProfileWithDefaults() *TaxProfile`

NewTaxProfileWithDefaults instantiates a new TaxProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TaxProfile) GetAddress() TaxAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TaxProfile) GetAddressOk() (*TaxAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TaxProfile) SetAddress(v TaxAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *TaxProfile) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBusinessName

`func (o *TaxProfile) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *TaxProfile) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *TaxProfile) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *TaxProfile) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### GetCertification

`func (o *TaxProfile) GetCertification() TaxCertification`

GetCertification returns the Certification field if non-nil, zero value otherwise.

### GetCertificationOk

`func (o *TaxProfile) GetCertificationOk() (*TaxCertification, bool)`

GetCertificationOk returns a tuple with the Certification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertification

`func (o *TaxProfile) SetCertification(v TaxCertification)`

SetCertification sets Certification field to given value.

### HasCertification

`func (o *TaxProfile) HasCertification() bool`

HasCertification returns a boolean if a field has been set.

### GetClassification

`func (o *TaxProfile) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *TaxProfile) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *TaxProfile) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *TaxProfile) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetConsent

`func (o *TaxProfile) GetConsent() TaxConsent`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *TaxProfile) GetConsentOk() (*TaxConsent, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *TaxProfile) SetConsent(v TaxConsent)`

SetConsent sets Consent field to given value.

### HasConsent

`func (o *TaxProfile) HasConsent() bool`

HasConsent returns a boolean if a field has been set.

### GetDisclosure

`func (o *TaxProfile) GetDisclosure() string`

GetDisclosure returns the Disclosure field if non-nil, zero value otherwise.

### GetDisclosureOk

`func (o *TaxProfile) GetDisclosureOk() (*string, bool)`

GetDisclosureOk returns a tuple with the Disclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosure

`func (o *TaxProfile) SetDisclosure(v string)`

SetDisclosure sets Disclosure field to given value.

### HasDisclosure

`func (o *TaxProfile) HasDisclosure() bool`

HasDisclosure returns a boolean if a field has been set.

### GetExemptPayee

`func (o *TaxProfile) GetExemptPayee() string`

GetExemptPayee returns the ExemptPayee field if non-nil, zero value otherwise.

### GetExemptPayeeOk

`func (o *TaxProfile) GetExemptPayeeOk() (*string, bool)`

GetExemptPayeeOk returns a tuple with the ExemptPayee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExemptPayee

`func (o *TaxProfile) SetExemptPayee(v string)`

SetExemptPayee sets ExemptPayee field to given value.

### HasExemptPayee

`func (o *TaxProfile) HasExemptPayee() bool`

HasExemptPayee returns a boolean if a field has been set.

### GetExpires

`func (o *TaxProfile) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *TaxProfile) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *TaxProfile) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *TaxProfile) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetFatca

`func (o *TaxProfile) GetFatca() string`

GetFatca returns the Fatca field if non-nil, zero value otherwise.

### GetFatcaOk

`func (o *TaxProfile) GetFatcaOk() (*string, bool)`

GetFatcaOk returns a tuple with the Fatca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFatca

`func (o *TaxProfile) SetFatca(v string)`

SetFatca sets Fatca field to given value.

### HasFatca

`func (o *TaxProfile) HasFatca() bool`

HasFatca returns a boolean if a field has been set.

### GetForeignOwners

`func (o *TaxProfile) GetForeignOwners() bool`

GetForeignOwners returns the ForeignOwners field if non-nil, zero value otherwise.

### GetForeignOwnersOk

`func (o *TaxProfile) GetForeignOwnersOk() (*bool, bool)`

GetForeignOwnersOk returns a tuple with the ForeignOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignOwners

`func (o *TaxProfile) SetForeignOwners(v bool)`

SetForeignOwners sets ForeignOwners field to given value.

### HasForeignOwners

`func (o *TaxProfile) HasForeignOwners() bool`

HasForeignOwners returns a boolean if a field has been set.

### GetForeignTin

`func (o *TaxProfile) GetForeignTin() string`

GetForeignTin returns the ForeignTin field if non-nil, zero value otherwise.

### GetForeignTinOk

`func (o *TaxProfile) GetForeignTinOk() (*string, bool)`

GetForeignTinOk returns a tuple with the ForeignTin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignTin

`func (o *TaxProfile) SetForeignTin(v string)`

SetForeignTin sets ForeignTin field to given value.

### HasForeignTin

`func (o *TaxProfile) HasForeignTin() bool`

HasForeignTin returns a boolean if a field has been set.

### GetForm

`func (o *TaxProfile) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *TaxProfile) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *TaxProfile) SetForm(v string)`

SetForm sets Form field to given value.

### HasForm

`func (o *TaxProfile) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetName

`func (o *TaxProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *TaxProfile) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *TaxProfile) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *TaxProfile) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *TaxProfile) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetTin

`func (o *TaxProfile) GetTin() string`

GetTin returns the Tin field if non-nil, zero value otherwise.

### GetTinOk

`func (o *TaxProfile) GetTinOk() (*string, bool)`

GetTinOk returns a tuple with the Tin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTin

`func (o *TaxProfile) SetTin(v string)`

SetTin sets Tin field to given value.

### HasTin

`func (o *TaxProfile) HasTin() bool`

HasTin returns a boolean if a field has been set.

### GetTinType

`func (o *TaxProfile) GetTinType() string`

GetTinType returns the TinType field if non-nil, zero value otherwise.

### GetTinTypeOk

`func (o *TaxProfile) GetTinTypeOk() (*string, bool)`

GetTinTypeOk returns a tuple with the TinType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTinType

`func (o *TaxProfile) SetTinType(v string)`

SetTinType sets TinType field to given value.

### HasTinType

`func (o *TaxProfile) HasTinType() bool`

HasTinType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TaxProfile) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaxProfile) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaxProfile) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TaxProfile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValid

`func (o *TaxProfile) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *TaxProfile) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *TaxProfile) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *TaxProfile) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetVersion

`func (o *TaxProfile) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TaxProfile) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TaxProfile) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TaxProfile) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetW8

`func (o *TaxProfile) GetW8() TaxW8`

GetW8 returns the W8 field if non-nil, zero value otherwise.

### GetW8Ok

`func (o *TaxProfile) GetW8Ok() (*TaxW8, bool)`

GetW8Ok returns a tuple with the W8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW8

`func (o *TaxProfile) SetW8(v TaxW8)`

SetW8 sets W8 field to given value.

### HasW8

`func (o *TaxProfile) HasW8() bool`

HasW8 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


