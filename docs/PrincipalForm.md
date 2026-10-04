# PrincipalForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certified** | Pointer to **bool** | Certified is whether the signature covers the current facts. | [optional] 
**Chapter3** | Pointer to **string** | Chapter3 and Chapter4 are a W-8BEN-E&#39;s statuses — to the org&#39;s admins. | [optional] 
**Chapter4** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** | Country is US for a W-9, or the country of citizenship or incorporation. | [optional] 
**Expires** | Pointer to **int64** | Expires is when a W-8 stops being valid, unix seconds. | [optional] 
**ForeignTin** | Pointer to **string** |  | [optional] 
**Form** | Pointer to **string** | Form is w9, w8ben or w8bene. | [optional] 
**Residence** | Pointer to **string** | Residence is the country of the permanent residence address. | [optional] 
**Tin** | Pointer to **string** | TIN and ForeignTIN are masked — answered to the org&#39;s admins. | [optional] 
**Treaty** | Pointer to [**PrincipalTreaty**](PrincipalTreaty.md) | Treaty is a W-8&#39;s treaty claim — to the org&#39;s admins. | [optional] 
**UsPerson** | Pointer to **bool** | USPerson is true for a W-9. | [optional] 
**Valid** | Pointer to **bool** | Valid is whether it establishes what it certifies today. | [optional] 

## Methods

### NewPrincipalForm

`func NewPrincipalForm() *PrincipalForm`

NewPrincipalForm instantiates a new PrincipalForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalFormWithDefaults

`func NewPrincipalFormWithDefaults() *PrincipalForm`

NewPrincipalFormWithDefaults instantiates a new PrincipalForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertified

`func (o *PrincipalForm) GetCertified() bool`

GetCertified returns the Certified field if non-nil, zero value otherwise.

### GetCertifiedOk

`func (o *PrincipalForm) GetCertifiedOk() (*bool, bool)`

GetCertifiedOk returns a tuple with the Certified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertified

`func (o *PrincipalForm) SetCertified(v bool)`

SetCertified sets Certified field to given value.

### HasCertified

`func (o *PrincipalForm) HasCertified() bool`

HasCertified returns a boolean if a field has been set.

### GetChapter3

`func (o *PrincipalForm) GetChapter3() string`

GetChapter3 returns the Chapter3 field if non-nil, zero value otherwise.

### GetChapter3Ok

`func (o *PrincipalForm) GetChapter3Ok() (*string, bool)`

GetChapter3Ok returns a tuple with the Chapter3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapter3

`func (o *PrincipalForm) SetChapter3(v string)`

SetChapter3 sets Chapter3 field to given value.

### HasChapter3

`func (o *PrincipalForm) HasChapter3() bool`

HasChapter3 returns a boolean if a field has been set.

### GetChapter4

`func (o *PrincipalForm) GetChapter4() string`

GetChapter4 returns the Chapter4 field if non-nil, zero value otherwise.

### GetChapter4Ok

`func (o *PrincipalForm) GetChapter4Ok() (*string, bool)`

GetChapter4Ok returns a tuple with the Chapter4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapter4

`func (o *PrincipalForm) SetChapter4(v string)`

SetChapter4 sets Chapter4 field to given value.

### HasChapter4

`func (o *PrincipalForm) HasChapter4() bool`

HasChapter4 returns a boolean if a field has been set.

### GetCountry

`func (o *PrincipalForm) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PrincipalForm) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PrincipalForm) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PrincipalForm) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetExpires

`func (o *PrincipalForm) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *PrincipalForm) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *PrincipalForm) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *PrincipalForm) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetForeignTin

`func (o *PrincipalForm) GetForeignTin() string`

GetForeignTin returns the ForeignTin field if non-nil, zero value otherwise.

### GetForeignTinOk

`func (o *PrincipalForm) GetForeignTinOk() (*string, bool)`

GetForeignTinOk returns a tuple with the ForeignTin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignTin

`func (o *PrincipalForm) SetForeignTin(v string)`

SetForeignTin sets ForeignTin field to given value.

### HasForeignTin

`func (o *PrincipalForm) HasForeignTin() bool`

HasForeignTin returns a boolean if a field has been set.

### GetForm

`func (o *PrincipalForm) GetForm() string`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *PrincipalForm) GetFormOk() (*string, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *PrincipalForm) SetForm(v string)`

SetForm sets Form field to given value.

### HasForm

`func (o *PrincipalForm) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetResidence

`func (o *PrincipalForm) GetResidence() string`

GetResidence returns the Residence field if non-nil, zero value otherwise.

### GetResidenceOk

`func (o *PrincipalForm) GetResidenceOk() (*string, bool)`

GetResidenceOk returns a tuple with the Residence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidence

`func (o *PrincipalForm) SetResidence(v string)`

SetResidence sets Residence field to given value.

### HasResidence

`func (o *PrincipalForm) HasResidence() bool`

HasResidence returns a boolean if a field has been set.

### GetTin

`func (o *PrincipalForm) GetTin() string`

GetTin returns the Tin field if non-nil, zero value otherwise.

### GetTinOk

`func (o *PrincipalForm) GetTinOk() (*string, bool)`

GetTinOk returns a tuple with the Tin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTin

`func (o *PrincipalForm) SetTin(v string)`

SetTin sets Tin field to given value.

### HasTin

`func (o *PrincipalForm) HasTin() bool`

HasTin returns a boolean if a field has been set.

### GetTreaty

`func (o *PrincipalForm) GetTreaty() PrincipalTreaty`

GetTreaty returns the Treaty field if non-nil, zero value otherwise.

### GetTreatyOk

`func (o *PrincipalForm) GetTreatyOk() (*PrincipalTreaty, bool)`

GetTreatyOk returns a tuple with the Treaty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreaty

`func (o *PrincipalForm) SetTreaty(v PrincipalTreaty)`

SetTreaty sets Treaty field to given value.

### HasTreaty

`func (o *PrincipalForm) HasTreaty() bool`

HasTreaty returns a boolean if a field has been set.

### GetUsPerson

`func (o *PrincipalForm) GetUsPerson() bool`

GetUsPerson returns the UsPerson field if non-nil, zero value otherwise.

### GetUsPersonOk

`func (o *PrincipalForm) GetUsPersonOk() (*bool, bool)`

GetUsPersonOk returns a tuple with the UsPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsPerson

`func (o *PrincipalForm) SetUsPerson(v bool)`

SetUsPerson sets UsPerson field to given value.

### HasUsPerson

`func (o *PrincipalForm) HasUsPerson() bool`

HasUsPerson returns a boolean if a field has been set.

### GetValid

`func (o *PrincipalForm) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *PrincipalForm) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *PrincipalForm) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *PrincipalForm) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


