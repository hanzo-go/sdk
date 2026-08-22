# SeoKeywordOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to **string** | Cost is what this call cost, in USD, as an exact decimal string. It is the upstream&#39;s own number and it is what was debited. | [optional] 
**Keywords** | Pointer to [**[]SeoMetric**](SeoMetric.md) | Keywords is one measurement per phrase, in the order the upstream answered. | [optional] 

## Methods

### NewSeoKeywordOut

`func NewSeoKeywordOut() *SeoKeywordOut`

NewSeoKeywordOut instantiates a new SeoKeywordOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoKeywordOutWithDefaults

`func NewSeoKeywordOutWithDefaults() *SeoKeywordOut`

NewSeoKeywordOutWithDefaults instantiates a new SeoKeywordOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *SeoKeywordOut) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *SeoKeywordOut) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *SeoKeywordOut) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *SeoKeywordOut) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetKeywords

`func (o *SeoKeywordOut) GetKeywords() []SeoMetric`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *SeoKeywordOut) GetKeywordsOk() (*[]SeoMetric, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *SeoKeywordOut) SetKeywords(v []SeoMetric)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *SeoKeywordOut) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


