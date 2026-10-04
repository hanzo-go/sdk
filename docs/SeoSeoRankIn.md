# SeoSeoRankIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | Domain is the site, with or without a subdomain — \&quot;hanzo.ai\&quot;, \&quot;docs.hanzo.ai\&quot;. | [optional] 
**Language** | Pointer to **string** | Language is the ISO code. Defaults to \&quot;en\&quot;. | [optional] 
**Limit** | Pointer to **int64** | Limit is how many placements to return, 1 to 1000. Defaults to 100. | [optional] 
**Location** | Pointer to **int64** | Location is the market, as the upstream&#39;s numeric code. Defaults to 2840. | [optional] 

## Methods

### NewSeoSeoRankIn

`func NewSeoSeoRankIn() *SeoSeoRankIn`

NewSeoSeoRankIn instantiates a new SeoSeoRankIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoSeoRankInWithDefaults

`func NewSeoSeoRankInWithDefaults() *SeoSeoRankIn`

NewSeoSeoRankInWithDefaults instantiates a new SeoSeoRankIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *SeoSeoRankIn) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SeoSeoRankIn) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SeoSeoRankIn) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SeoSeoRankIn) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetLanguage

`func (o *SeoSeoRankIn) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *SeoSeoRankIn) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *SeoSeoRankIn) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *SeoSeoRankIn) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLimit

`func (o *SeoSeoRankIn) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SeoSeoRankIn) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SeoSeoRankIn) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SeoSeoRankIn) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLocation

`func (o *SeoSeoRankIn) GetLocation() int64`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SeoSeoRankIn) GetLocationOk() (*int64, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SeoSeoRankIn) SetLocation(v int64)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SeoSeoRankIn) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


