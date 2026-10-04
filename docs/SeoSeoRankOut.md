# SeoSeoRankOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to **string** | Cost is what this call cost, in USD, as an exact decimal string. | [optional] 
**Rankings** | Pointer to [**[]SeoSeoRanking**](SeoSeoRanking.md) | Rankings is one row per phrase the domain places for. | [optional] 
**Total** | Pointer to **int64** | Total is how many placements the upstream holds for this domain. | [optional] 

## Methods

### NewSeoSeoRankOut

`func NewSeoSeoRankOut() *SeoSeoRankOut`

NewSeoSeoRankOut instantiates a new SeoSeoRankOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoSeoRankOutWithDefaults

`func NewSeoSeoRankOutWithDefaults() *SeoSeoRankOut`

NewSeoSeoRankOutWithDefaults instantiates a new SeoSeoRankOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *SeoSeoRankOut) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *SeoSeoRankOut) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *SeoSeoRankOut) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *SeoSeoRankOut) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetRankings

`func (o *SeoSeoRankOut) GetRankings() []SeoSeoRanking`

GetRankings returns the Rankings field if non-nil, zero value otherwise.

### GetRankingsOk

`func (o *SeoSeoRankOut) GetRankingsOk() (*[]SeoSeoRanking, bool)`

GetRankingsOk returns a tuple with the Rankings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankings

`func (o *SeoSeoRankOut) SetRankings(v []SeoSeoRanking)`

SetRankings sets Rankings field to given value.

### HasRankings

`func (o *SeoSeoRankOut) HasRankings() bool`

HasRankings returns a boolean if a field has been set.

### GetTotal

`func (o *SeoSeoRankOut) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SeoSeoRankOut) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SeoSeoRankOut) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SeoSeoRankOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


