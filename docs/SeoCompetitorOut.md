# SeoCompetitorOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Competitors** | Pointer to [**[]SeoDomain**](SeoDomain.md) | Competitors is one row per domain, strongest first. | [optional] 
**Cost** | Pointer to **string** | Cost is what this call cost, in USD, as an exact decimal string. | [optional] 
**Total** | Pointer to **int32** | Total is how many domains the upstream holds for these phrases. | [optional] 

## Methods

### NewSeoCompetitorOut

`func NewSeoCompetitorOut() *SeoCompetitorOut`

NewSeoCompetitorOut instantiates a new SeoCompetitorOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoCompetitorOutWithDefaults

`func NewSeoCompetitorOutWithDefaults() *SeoCompetitorOut`

NewSeoCompetitorOutWithDefaults instantiates a new SeoCompetitorOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompetitors

`func (o *SeoCompetitorOut) GetCompetitors() []SeoDomain`

GetCompetitors returns the Competitors field if non-nil, zero value otherwise.

### GetCompetitorsOk

`func (o *SeoCompetitorOut) GetCompetitorsOk() (*[]SeoDomain, bool)`

GetCompetitorsOk returns a tuple with the Competitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitors

`func (o *SeoCompetitorOut) SetCompetitors(v []SeoDomain)`

SetCompetitors sets Competitors field to given value.

### HasCompetitors

`func (o *SeoCompetitorOut) HasCompetitors() bool`

HasCompetitors returns a boolean if a field has been set.

### GetCost

`func (o *SeoCompetitorOut) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *SeoCompetitorOut) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *SeoCompetitorOut) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *SeoCompetitorOut) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetTotal

`func (o *SeoCompetitorOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SeoCompetitorOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SeoCompetitorOut) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SeoCompetitorOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


