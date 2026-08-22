# SeoCompetitorIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keywords** | Pointer to **[]string** | Keywords are the phrases. At least one; blanks are dropped. | [optional] 
**Language** | Pointer to **string** | Language is the ISO code. Defaults to \&quot;en\&quot;. | [optional] 
**Limit** | Pointer to **int32** | Limit is how many domains to return, 1 to 1000. Defaults to 100. | [optional] 
**Location** | Pointer to **int32** | Location is the market, as the upstream&#39;s numeric code. Defaults to 2840. | [optional] 

## Methods

### NewSeoCompetitorIn

`func NewSeoCompetitorIn() *SeoCompetitorIn`

NewSeoCompetitorIn instantiates a new SeoCompetitorIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoCompetitorInWithDefaults

`func NewSeoCompetitorInWithDefaults() *SeoCompetitorIn`

NewSeoCompetitorInWithDefaults instantiates a new SeoCompetitorIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeywords

`func (o *SeoCompetitorIn) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *SeoCompetitorIn) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *SeoCompetitorIn) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *SeoCompetitorIn) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetLanguage

`func (o *SeoCompetitorIn) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *SeoCompetitorIn) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *SeoCompetitorIn) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *SeoCompetitorIn) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLimit

`func (o *SeoCompetitorIn) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SeoCompetitorIn) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SeoCompetitorIn) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SeoCompetitorIn) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLocation

`func (o *SeoCompetitorIn) GetLocation() int32`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SeoCompetitorIn) GetLocationOk() (*int32, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SeoCompetitorIn) SetLocation(v int32)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SeoCompetitorIn) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


