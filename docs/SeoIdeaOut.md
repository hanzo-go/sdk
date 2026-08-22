# SeoIdeaOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to **string** | Cost is what this call cost, in USD, as an exact decimal string. | [optional] 
**Keywords** | Pointer to [**[]SeoMetric**](SeoMetric.md) | Keywords is the phrases found, each measured. | [optional] 
**Total** | Pointer to **int32** | Total is how many the upstream holds, which is usually more than Limit returned — it is what raising the limit would reach. | [optional] 

## Methods

### NewSeoIdeaOut

`func NewSeoIdeaOut() *SeoIdeaOut`

NewSeoIdeaOut instantiates a new SeoIdeaOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoIdeaOutWithDefaults

`func NewSeoIdeaOutWithDefaults() *SeoIdeaOut`

NewSeoIdeaOutWithDefaults instantiates a new SeoIdeaOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *SeoIdeaOut) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *SeoIdeaOut) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *SeoIdeaOut) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *SeoIdeaOut) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetKeywords

`func (o *SeoIdeaOut) GetKeywords() []SeoMetric`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *SeoIdeaOut) GetKeywordsOk() (*[]SeoMetric, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *SeoIdeaOut) SetKeywords(v []SeoMetric)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *SeoIdeaOut) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetTotal

`func (o *SeoIdeaOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SeoIdeaOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SeoIdeaOut) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SeoIdeaOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


