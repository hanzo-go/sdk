# SeoMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Competition** | Pointer to **float32** | Competition is how contested the advertising is, from 0 to 1. The upstream reports it as an index out of a hundred on one endpoint and as this fraction on another; it is the fraction here in both cases. | [optional] 
**Cpc** | Pointer to **float32** | CPC is the average cost of one advertising click, in USD. It is a reported statistic about somebody else&#39;s auction, not an amount this API moves. | [optional] 
**Difficulty** | Pointer to **int32** | Difficulty is how hard the first page is to reach organically, 0 to 100. Present on seoIdea, which measures it; absent on seoKeyword, which does not. | [optional] 
**Keyword** | Pointer to **string** | Keyword is the phrase. | [optional] 
**Level** | Pointer to **string** | Level is the same fact as a word: low, medium or high. | [optional] 
**Volume** | Pointer to **int32** | Volume is the average monthly searches. | [optional] 

## Methods

### NewSeoMetric

`func NewSeoMetric() *SeoMetric`

NewSeoMetric instantiates a new SeoMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoMetricWithDefaults

`func NewSeoMetricWithDefaults() *SeoMetric`

NewSeoMetricWithDefaults instantiates a new SeoMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompetition

`func (o *SeoMetric) GetCompetition() float32`

GetCompetition returns the Competition field if non-nil, zero value otherwise.

### GetCompetitionOk

`func (o *SeoMetric) GetCompetitionOk() (*float32, bool)`

GetCompetitionOk returns a tuple with the Competition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetition

`func (o *SeoMetric) SetCompetition(v float32)`

SetCompetition sets Competition field to given value.

### HasCompetition

`func (o *SeoMetric) HasCompetition() bool`

HasCompetition returns a boolean if a field has been set.

### GetCpc

`func (o *SeoMetric) GetCpc() float32`

GetCpc returns the Cpc field if non-nil, zero value otherwise.

### GetCpcOk

`func (o *SeoMetric) GetCpcOk() (*float32, bool)`

GetCpcOk returns a tuple with the Cpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpc

`func (o *SeoMetric) SetCpc(v float32)`

SetCpc sets Cpc field to given value.

### HasCpc

`func (o *SeoMetric) HasCpc() bool`

HasCpc returns a boolean if a field has been set.

### GetDifficulty

`func (o *SeoMetric) GetDifficulty() int32`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *SeoMetric) GetDifficultyOk() (*int32, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *SeoMetric) SetDifficulty(v int32)`

SetDifficulty sets Difficulty field to given value.

### HasDifficulty

`func (o *SeoMetric) HasDifficulty() bool`

HasDifficulty returns a boolean if a field has been set.

### GetKeyword

`func (o *SeoMetric) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *SeoMetric) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *SeoMetric) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *SeoMetric) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetLevel

`func (o *SeoMetric) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *SeoMetric) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *SeoMetric) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *SeoMetric) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetVolume

`func (o *SeoMetric) GetVolume() int32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *SeoMetric) GetVolumeOk() (*int32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *SeoMetric) SetVolume(v int32)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *SeoMetric) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


