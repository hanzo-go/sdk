# SeoSeoIdeaOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to **string** | Cost is what this call cost, in USD, as an exact decimal string. | [optional] 
**Keywords** | Pointer to [**[]SeoSeoMetric**](SeoSeoMetric.md) | Keywords is the phrases found, each measured. | [optional] 
**Total** | Pointer to **int64** | Total is how many the upstream holds, which is usually more than Limit returned — it is what raising the limit would reach. | [optional] 

## Methods

### NewSeoSeoIdeaOut

`func NewSeoSeoIdeaOut() *SeoSeoIdeaOut`

NewSeoSeoIdeaOut instantiates a new SeoSeoIdeaOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoSeoIdeaOutWithDefaults

`func NewSeoSeoIdeaOutWithDefaults() *SeoSeoIdeaOut`

NewSeoSeoIdeaOutWithDefaults instantiates a new SeoSeoIdeaOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *SeoSeoIdeaOut) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *SeoSeoIdeaOut) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *SeoSeoIdeaOut) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *SeoSeoIdeaOut) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetKeywords

`func (o *SeoSeoIdeaOut) GetKeywords() []SeoSeoMetric`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *SeoSeoIdeaOut) GetKeywordsOk() (*[]SeoSeoMetric, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *SeoSeoIdeaOut) SetKeywords(v []SeoSeoMetric)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *SeoSeoIdeaOut) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetTotal

`func (o *SeoSeoIdeaOut) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SeoSeoIdeaOut) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SeoSeoIdeaOut) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SeoSeoIdeaOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


