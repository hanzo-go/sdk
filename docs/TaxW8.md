# TaxW8

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Birth** | Pointer to **string** | Birth is W-8BEN line 8, the individual&#39;s date of birth, YYYY-MM-DD. | [optional] 
**Capacity** | Pointer to **string** | Capacity is the capacity in which the signer signs for the beneficial owner — \&quot;Director\&quot;, \&quot;Authorized officer\&quot;. Required on a W-8BEN-E; on a W-8BEN, empty when the beneficial owner signs. | [optional] 
**Chapter3** | Pointer to **string** | Chapter3 is W-8BEN-E line 4. A W-8BEN&#39;s beneficial owner is an individual and states none. | [optional] 
**Chapter4** | Pointer to **string** | Chapter4 is W-8BEN-E line 5, the FATCA status. | [optional] 
**Country** | Pointer to **string** | Country is line 2: the country of citizenship (W-8BEN) or of incorporation or organization (W-8BEN-E), ISO 3166-1 alpha-2. | [optional] 
**Giin** | Pointer to **string** | GIIN is W-8BEN-E line 9a, when the chapter 4 status carries one. | [optional] 
**NoForeignTin** | Pointer to **bool** | NoForeignTIN is W-8BEN line 6b / W-8BEN-E line 9b&#39;s alternative: the jurisdiction of residence does not require or issue a foreign TIN. | [optional] 
**Treaty** | Pointer to [**TaxTreaty**](TaxTreaty.md) | Treaty is the claim of treaty benefits (W-8BEN Part II lines 9-10, W-8BEN-E Part III lines 14-15), when one is made. | [optional] 

## Methods

### NewTaxW8

`func NewTaxW8() *TaxW8`

NewTaxW8 instantiates a new TaxW8 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxW8WithDefaults

`func NewTaxW8WithDefaults() *TaxW8`

NewTaxW8WithDefaults instantiates a new TaxW8 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBirth

`func (o *TaxW8) GetBirth() string`

GetBirth returns the Birth field if non-nil, zero value otherwise.

### GetBirthOk

`func (o *TaxW8) GetBirthOk() (*string, bool)`

GetBirthOk returns a tuple with the Birth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirth

`func (o *TaxW8) SetBirth(v string)`

SetBirth sets Birth field to given value.

### HasBirth

`func (o *TaxW8) HasBirth() bool`

HasBirth returns a boolean if a field has been set.

### GetCapacity

`func (o *TaxW8) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *TaxW8) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *TaxW8) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *TaxW8) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetChapter3

`func (o *TaxW8) GetChapter3() string`

GetChapter3 returns the Chapter3 field if non-nil, zero value otherwise.

### GetChapter3Ok

`func (o *TaxW8) GetChapter3Ok() (*string, bool)`

GetChapter3Ok returns a tuple with the Chapter3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapter3

`func (o *TaxW8) SetChapter3(v string)`

SetChapter3 sets Chapter3 field to given value.

### HasChapter3

`func (o *TaxW8) HasChapter3() bool`

HasChapter3 returns a boolean if a field has been set.

### GetChapter4

`func (o *TaxW8) GetChapter4() string`

GetChapter4 returns the Chapter4 field if non-nil, zero value otherwise.

### GetChapter4Ok

`func (o *TaxW8) GetChapter4Ok() (*string, bool)`

GetChapter4Ok returns a tuple with the Chapter4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapter4

`func (o *TaxW8) SetChapter4(v string)`

SetChapter4 sets Chapter4 field to given value.

### HasChapter4

`func (o *TaxW8) HasChapter4() bool`

HasChapter4 returns a boolean if a field has been set.

### GetCountry

`func (o *TaxW8) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TaxW8) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TaxW8) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *TaxW8) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetGiin

`func (o *TaxW8) GetGiin() string`

GetGiin returns the Giin field if non-nil, zero value otherwise.

### GetGiinOk

`func (o *TaxW8) GetGiinOk() (*string, bool)`

GetGiinOk returns a tuple with the Giin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiin

`func (o *TaxW8) SetGiin(v string)`

SetGiin sets Giin field to given value.

### HasGiin

`func (o *TaxW8) HasGiin() bool`

HasGiin returns a boolean if a field has been set.

### GetNoForeignTin

`func (o *TaxW8) GetNoForeignTin() bool`

GetNoForeignTin returns the NoForeignTin field if non-nil, zero value otherwise.

### GetNoForeignTinOk

`func (o *TaxW8) GetNoForeignTinOk() (*bool, bool)`

GetNoForeignTinOk returns a tuple with the NoForeignTin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoForeignTin

`func (o *TaxW8) SetNoForeignTin(v bool)`

SetNoForeignTin sets NoForeignTin field to given value.

### HasNoForeignTin

`func (o *TaxW8) HasNoForeignTin() bool`

HasNoForeignTin returns a boolean if a field has been set.

### GetTreaty

`func (o *TaxW8) GetTreaty() TaxTreaty`

GetTreaty returns the Treaty field if non-nil, zero value otherwise.

### GetTreatyOk

`func (o *TaxW8) GetTreatyOk() (*TaxTreaty, bool)`

GetTreatyOk returns a tuple with the Treaty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreaty

`func (o *TaxW8) SetTreaty(v TaxTreaty)`

SetTreaty sets Treaty field to given value.

### HasTreaty

`func (o *TaxW8) HasTreaty() bool`

HasTreaty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


