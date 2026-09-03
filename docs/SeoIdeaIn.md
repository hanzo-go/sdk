# SeoIdeaIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keywords** | Pointer to **[]string** | Keywords are the seeds. At least one; blanks are dropped. | [optional] 
**Language** | Pointer to **string** | Language is the ISO code. Defaults to \&quot;en\&quot;. | [optional] 
**Limit** | Pointer to **int64** | Limit is how many phrases to return, 1 to 1000. Defaults to 100. It is what this call is priced on, because the upstream charges per row. | [optional] 
**Location** | Pointer to **int64** | Location is the market, as the upstream&#39;s numeric code. Defaults to 2840. | [optional] 

## Methods

### NewSeoIdeaIn

`func NewSeoIdeaIn() *SeoIdeaIn`

NewSeoIdeaIn instantiates a new SeoIdeaIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoIdeaInWithDefaults

`func NewSeoIdeaInWithDefaults() *SeoIdeaIn`

NewSeoIdeaInWithDefaults instantiates a new SeoIdeaIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeywords

`func (o *SeoIdeaIn) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *SeoIdeaIn) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *SeoIdeaIn) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *SeoIdeaIn) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetLanguage

`func (o *SeoIdeaIn) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *SeoIdeaIn) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *SeoIdeaIn) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *SeoIdeaIn) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLimit

`func (o *SeoIdeaIn) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SeoIdeaIn) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SeoIdeaIn) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SeoIdeaIn) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLocation

`func (o *SeoIdeaIn) GetLocation() int64`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SeoIdeaIn) GetLocationOk() (*int64, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SeoIdeaIn) SetLocation(v int64)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SeoIdeaIn) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


