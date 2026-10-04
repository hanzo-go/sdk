# SeoSeoMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Competition** | Pointer to **float64** | Competition is how contested the advertising is, from 0 to 1. The upstream reports it as an index out of a hundred on one endpoint and as this fraction on another; it is the fraction here in both cases. | [optional] 
**Cpc** | Pointer to **float64** | CPC is the average cost of one advertising click, in USD. It is a reported statistic about somebody else&#39;s auction, not an amount this API moves. | [optional] 
**Difficulty** | Pointer to **int64** | Difficulty is how hard the first page is to reach organically, 0 to 100. Present on seoIdea, which measures it; absent on seoKeyword, which does not. | [optional] 
**Keyword** | Pointer to **string** | Keyword is the phrase. | [optional] 
**Level** | Pointer to **string** | Level is the same fact as a word: low, medium or high. | [optional] 
**Volume** | Pointer to **int64** | Volume is the average monthly searches. | [optional] 

## Methods

### NewSeoSeoMetric

`func NewSeoSeoMetric() *SeoSeoMetric`

NewSeoSeoMetric instantiates a new SeoSeoMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoSeoMetricWithDefaults

`func NewSeoSeoMetricWithDefaults() *SeoSeoMetric`

NewSeoSeoMetricWithDefaults instantiates a new SeoSeoMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompetition

`func (o *SeoSeoMetric) GetCompetition() float64`

GetCompetition returns the Competition field if non-nil, zero value otherwise.

### GetCompetitionOk

`func (o *SeoSeoMetric) GetCompetitionOk() (*float64, bool)`

GetCompetitionOk returns a tuple with the Competition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetition

`func (o *SeoSeoMetric) SetCompetition(v float64)`

SetCompetition sets Competition field to given value.

### HasCompetition

`func (o *SeoSeoMetric) HasCompetition() bool`

HasCompetition returns a boolean if a field has been set.

### GetCpc

`func (o *SeoSeoMetric) GetCpc() float64`

GetCpc returns the Cpc field if non-nil, zero value otherwise.

### GetCpcOk

`func (o *SeoSeoMetric) GetCpcOk() (*float64, bool)`

GetCpcOk returns a tuple with the Cpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpc

`func (o *SeoSeoMetric) SetCpc(v float64)`

SetCpc sets Cpc field to given value.

### HasCpc

`func (o *SeoSeoMetric) HasCpc() bool`

HasCpc returns a boolean if a field has been set.

### GetDifficulty

`func (o *SeoSeoMetric) GetDifficulty() int64`

GetDifficulty returns the Difficulty field if non-nil, zero value otherwise.

### GetDifficultyOk

`func (o *SeoSeoMetric) GetDifficultyOk() (*int64, bool)`

GetDifficultyOk returns a tuple with the Difficulty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifficulty

`func (o *SeoSeoMetric) SetDifficulty(v int64)`

SetDifficulty sets Difficulty field to given value.

### HasDifficulty

`func (o *SeoSeoMetric) HasDifficulty() bool`

HasDifficulty returns a boolean if a field has been set.

### GetKeyword

`func (o *SeoSeoMetric) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *SeoSeoMetric) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *SeoSeoMetric) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *SeoSeoMetric) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetLevel

`func (o *SeoSeoMetric) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *SeoSeoMetric) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *SeoSeoMetric) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *SeoSeoMetric) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetVolume

`func (o *SeoSeoMetric) GetVolume() int64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *SeoSeoMetric) GetVolumeOk() (*int64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *SeoSeoMetric) SetVolume(v int64)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *SeoSeoMetric) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


