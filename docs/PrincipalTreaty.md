# PrincipalTreaty

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

### NewPrincipalTreaty

`func NewPrincipalTreaty() *PrincipalTreaty`

NewPrincipalTreaty instantiates a new PrincipalTreaty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalTreatyWithDefaults

`func NewPrincipalTreatyWithDefaults() *PrincipalTreaty`

NewPrincipalTreatyWithDefaults instantiates a new PrincipalTreaty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArticle

`func (o *PrincipalTreaty) GetArticle() string`

GetArticle returns the Article field if non-nil, zero value otherwise.

### GetArticleOk

`func (o *PrincipalTreaty) GetArticleOk() (*string, bool)`

GetArticleOk returns a tuple with the Article field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticle

`func (o *PrincipalTreaty) SetArticle(v string)`

SetArticle sets Article field to given value.

### HasArticle

`func (o *PrincipalTreaty) HasArticle() bool`

HasArticle returns a boolean if a field has been set.

### GetConditions

`func (o *PrincipalTreaty) GetConditions() string`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PrincipalTreaty) GetConditionsOk() (*string, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PrincipalTreaty) SetConditions(v string)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *PrincipalTreaty) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCountry

`func (o *PrincipalTreaty) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PrincipalTreaty) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PrincipalTreaty) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PrincipalTreaty) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetIncome

`func (o *PrincipalTreaty) GetIncome() string`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *PrincipalTreaty) GetIncomeOk() (*string, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *PrincipalTreaty) SetIncome(v string)`

SetIncome sets Income field to given value.

### HasIncome

`func (o *PrincipalTreaty) HasIncome() bool`

HasIncome returns a boolean if a field has been set.

### GetLob

`func (o *PrincipalTreaty) GetLob() string`

GetLob returns the Lob field if non-nil, zero value otherwise.

### GetLobOk

`func (o *PrincipalTreaty) GetLobOk() (*string, bool)`

GetLobOk returns a tuple with the Lob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLob

`func (o *PrincipalTreaty) SetLob(v string)`

SetLob sets Lob field to given value.

### HasLob

`func (o *PrincipalTreaty) HasLob() bool`

HasLob returns a boolean if a field has been set.

### GetRateBps

`func (o *PrincipalTreaty) GetRateBps() int64`

GetRateBps returns the RateBps field if non-nil, zero value otherwise.

### GetRateBpsOk

`func (o *PrincipalTreaty) GetRateBpsOk() (*int64, bool)`

GetRateBpsOk returns a tuple with the RateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateBps

`func (o *PrincipalTreaty) SetRateBps(v int64)`

SetRateBps sets RateBps field to given value.

### HasRateBps

`func (o *PrincipalTreaty) HasRateBps() bool`

HasRateBps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


