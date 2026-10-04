# TaxTreaty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Article** | Pointer to **string** | Article is the treaty article and paragraph claimed, e.g. \&quot;12(2)\&quot;. | [optional] 
**Conditions** | Pointer to **string** | Conditions is the explanation the form asks for: the conditions of the article the beneficial owner meets. | [optional] 
**Country** | Pointer to **string** | Country is the treaty country the beneficial owner is resident in. | [optional] 
**Income** | Pointer to **string** | Income is the type of income the claim covers: services, rents, royalties or other. | [optional] 
**Lob** | Pointer to **string** | LOB is W-8BEN-E line 14b: the treaty&#39;s limitation on benefits provision the entity meets. Required on a W-8BEN-E treaty claim; a W-8BEN has no line 14b. | [optional] 
**RateBps** | Pointer to **int64** | RateBps is the claimed withholding rate in basis points: 0 is exempt, 1000 is 10%. | [optional] 

## Methods

### NewTaxTreaty

`func NewTaxTreaty() *TaxTreaty`

NewTaxTreaty instantiates a new TaxTreaty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxTreatyWithDefaults

`func NewTaxTreatyWithDefaults() *TaxTreaty`

NewTaxTreatyWithDefaults instantiates a new TaxTreaty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticle

`func (o *TaxTreaty) GetArticle() string`

GetArticle returns the Article field if non-nil, zero value otherwise.

### GetArticleOk

`func (o *TaxTreaty) GetArticleOk() (*string, bool)`

GetArticleOk returns a tuple with the Article field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticle

`func (o *TaxTreaty) SetArticle(v string)`

SetArticle sets Article field to given value.

### HasArticle

`func (o *TaxTreaty) HasArticle() bool`

HasArticle returns a boolean if a field has been set.

### GetConditions

`func (o *TaxTreaty) GetConditions() string`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *TaxTreaty) GetConditionsOk() (*string, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *TaxTreaty) SetConditions(v string)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *TaxTreaty) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCountry

`func (o *TaxTreaty) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TaxTreaty) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TaxTreaty) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *TaxTreaty) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetIncome

`func (o *TaxTreaty) GetIncome() string`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *TaxTreaty) GetIncomeOk() (*string, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *TaxTreaty) SetIncome(v string)`

SetIncome sets Income field to given value.

### HasIncome

`func (o *TaxTreaty) HasIncome() bool`

HasIncome returns a boolean if a field has been set.

### GetLob

`func (o *TaxTreaty) GetLob() string`

GetLob returns the Lob field if non-nil, zero value otherwise.

### GetLobOk

`func (o *TaxTreaty) GetLobOk() (*string, bool)`

GetLobOk returns a tuple with the Lob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLob

`func (o *TaxTreaty) SetLob(v string)`

SetLob sets Lob field to given value.

### HasLob

`func (o *TaxTreaty) HasLob() bool`

HasLob returns a boolean if a field has been set.

### GetRateBps

`func (o *TaxTreaty) GetRateBps() int64`

GetRateBps returns the RateBps field if non-nil, zero value otherwise.

### GetRateBpsOk

`func (o *TaxTreaty) GetRateBpsOk() (*int64, bool)`

GetRateBpsOk returns a tuple with the RateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateBps

`func (o *TaxTreaty) SetRateBps(v int64)`

SetRateBps sets RateBps field to given value.

### HasRateBps

`func (o *TaxTreaty) HasRateBps() bool`

HasRateBps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


