# SeoSeoRanking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keyword** | Pointer to **string** | Keyword is the phrase searched. | [optional] 
**Position** | Pointer to **int64** | Position is the absolute rank on the results page, counting every element — so it is what a person scrolling actually passes, not the organic-only rank. | [optional] 
**Title** | Pointer to **string** | Title is that result&#39;s headline. | [optional] 
**Traffic** | Pointer to **float64** | Traffic is the estimated monthly visits this placement earns. | [optional] 
**Url** | Pointer to **string** | URL is the page of the target that placed. | [optional] 
**Volume** | Pointer to **int64** | Volume is the phrase&#39;s average monthly searches. | [optional] 

## Methods

### NewSeoSeoRanking

`func NewSeoSeoRanking() *SeoSeoRanking`

NewSeoSeoRanking instantiates a new SeoSeoRanking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoSeoRankingWithDefaults

`func NewSeoSeoRankingWithDefaults() *SeoSeoRanking`

NewSeoSeoRankingWithDefaults instantiates a new SeoSeoRanking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyword

`func (o *SeoSeoRanking) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *SeoSeoRanking) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *SeoSeoRanking) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *SeoSeoRanking) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetPosition

`func (o *SeoSeoRanking) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *SeoSeoRanking) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *SeoSeoRanking) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *SeoSeoRanking) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetTitle

`func (o *SeoSeoRanking) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SeoSeoRanking) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SeoSeoRanking) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SeoSeoRanking) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTraffic

`func (o *SeoSeoRanking) GetTraffic() float64`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *SeoSeoRanking) GetTrafficOk() (*float64, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *SeoSeoRanking) SetTraffic(v float64)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *SeoSeoRanking) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.

### GetUrl

`func (o *SeoSeoRanking) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SeoSeoRanking) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SeoSeoRanking) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SeoSeoRanking) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVolume

`func (o *SeoSeoRanking) GetVolume() int64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *SeoSeoRanking) GetVolumeOk() (*int64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *SeoSeoRanking) SetVolume(v int64)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *SeoSeoRanking) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


