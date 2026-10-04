# SeoSeoKeywordIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keywords** | Pointer to **[]string** | Keywords are the phrases. At least one; blanks are dropped. | [optional] 
**Language** | Pointer to **string** | Language is the ISO code. Defaults to \&quot;en\&quot;. | [optional] 
**Location** | Pointer to **int64** | Location is the market, as the upstream&#39;s numeric code. Defaults to 2840, the United States. | [optional] 

## Methods

### NewSeoSeoKeywordIn

`func NewSeoSeoKeywordIn() *SeoSeoKeywordIn`

NewSeoSeoKeywordIn instantiates a new SeoSeoKeywordIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoSeoKeywordInWithDefaults

`func NewSeoSeoKeywordInWithDefaults() *SeoSeoKeywordIn`

NewSeoSeoKeywordInWithDefaults instantiates a new SeoSeoKeywordIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeywords

`func (o *SeoSeoKeywordIn) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *SeoSeoKeywordIn) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *SeoSeoKeywordIn) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *SeoSeoKeywordIn) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetLanguage

`func (o *SeoSeoKeywordIn) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *SeoSeoKeywordIn) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *SeoSeoKeywordIn) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *SeoSeoKeywordIn) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLocation

`func (o *SeoSeoKeywordIn) GetLocation() int64`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SeoSeoKeywordIn) GetLocationOk() (*int64, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SeoSeoKeywordIn) SetLocation(v int64)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SeoSeoKeywordIn) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


