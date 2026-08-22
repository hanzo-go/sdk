# SeoRankOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | Pointer to **string** | Cost is what this call cost, in USD, as an exact decimal string. | [optional] 
**Rankings** | Pointer to [**[]SeoRanking**](SeoRanking.md) | Rankings is one row per phrase the domain places for. | [optional] 
**Total** | Pointer to **int32** | Total is how many placements the upstream holds for this domain. | [optional] 

## Methods

### NewSeoRankOut

`func NewSeoRankOut() *SeoRankOut`

NewSeoRankOut instantiates a new SeoRankOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoRankOutWithDefaults

`func NewSeoRankOutWithDefaults() *SeoRankOut`

NewSeoRankOutWithDefaults instantiates a new SeoRankOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *SeoRankOut) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *SeoRankOut) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *SeoRankOut) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *SeoRankOut) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetRankings

`func (o *SeoRankOut) GetRankings() []SeoRanking`

GetRankings returns the Rankings field if non-nil, zero value otherwise.

### GetRankingsOk

`func (o *SeoRankOut) GetRankingsOk() (*[]SeoRanking, bool)`

GetRankingsOk returns a tuple with the Rankings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankings

`func (o *SeoRankOut) SetRankings(v []SeoRanking)`

SetRankings sets Rankings field to given value.

### HasRankings

`func (o *SeoRankOut) HasRankings() bool`

HasRankings returns a boolean if a field has been set.

### GetTotal

`func (o *SeoRankOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SeoRankOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SeoRankOut) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SeoRankOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


