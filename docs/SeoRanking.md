# SeoRanking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keyword** | Pointer to **string** | Keyword is the phrase searched. | [optional] 
**Position** | Pointer to **int32** | Position is the absolute rank on the results page, counting every element — so it is what a person scrolling actually passes, not the organic-only rank. | [optional] 
**Title** | Pointer to **string** | Title is that result&#39;s headline. | [optional] 
**Traffic** | Pointer to **float32** | Traffic is the estimated monthly visits this placement earns. | [optional] 
**Url** | Pointer to **string** | URL is the page of the target that placed. | [optional] 
**Volume** | Pointer to **int32** | Volume is the phrase&#39;s average monthly searches. | [optional] 

## Methods

### NewSeoRanking

`func NewSeoRanking() *SeoRanking`

NewSeoRanking instantiates a new SeoRanking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoRankingWithDefaults

`func NewSeoRankingWithDefaults() *SeoRanking`

NewSeoRankingWithDefaults instantiates a new SeoRanking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyword

`func (o *SeoRanking) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *SeoRanking) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *SeoRanking) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *SeoRanking) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetPosition

`func (o *SeoRanking) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *SeoRanking) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *SeoRanking) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *SeoRanking) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetTitle

`func (o *SeoRanking) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SeoRanking) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SeoRanking) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SeoRanking) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTraffic

`func (o *SeoRanking) GetTraffic() float32`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *SeoRanking) GetTrafficOk() (*float32, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *SeoRanking) SetTraffic(v float32)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *SeoRanking) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.

### GetUrl

`func (o *SeoRanking) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SeoRanking) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SeoRanking) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SeoRanking) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVolume

`func (o *SeoRanking) GetVolume() int32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *SeoRanking) GetVolumeOk() (*int32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *SeoRanking) SetVolume(v int32)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *SeoRanking) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


