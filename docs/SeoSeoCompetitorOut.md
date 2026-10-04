# SeoSeoCompetitorOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Competitors** | Pointer to [**[]SeoSeoDomain**](SeoSeoDomain.md) | Competitors is one row per domain, strongest first. | [optional] 
**Cost** | Pointer to **string** | Cost is what this call cost, in USD, as an exact decimal string. | [optional] 
**Total** | Pointer to **int64** | Total is how many domains the upstream holds for these phrases. | [optional] 

## Methods

### NewSeoSeoCompetitorOut

`func NewSeoSeoCompetitorOut() *SeoSeoCompetitorOut`

NewSeoSeoCompetitorOut instantiates a new SeoSeoCompetitorOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeoSeoCompetitorOutWithDefaults

`func NewSeoSeoCompetitorOutWithDefaults() *SeoSeoCompetitorOut`

NewSeoSeoCompetitorOutWithDefaults instantiates a new SeoSeoCompetitorOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompetitors

`func (o *SeoSeoCompetitorOut) GetCompetitors() []SeoSeoDomain`

GetCompetitors returns the Competitors field if non-nil, zero value otherwise.

### GetCompetitorsOk

`func (o *SeoSeoCompetitorOut) GetCompetitorsOk() (*[]SeoSeoDomain, bool)`

GetCompetitorsOk returns a tuple with the Competitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitors

`func (o *SeoSeoCompetitorOut) SetCompetitors(v []SeoSeoDomain)`

SetCompetitors sets Competitors field to given value.

### HasCompetitors

`func (o *SeoSeoCompetitorOut) HasCompetitors() bool`

HasCompetitors returns a boolean if a field has been set.

### GetCost

`func (o *SeoSeoCompetitorOut) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *SeoSeoCompetitorOut) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *SeoSeoCompetitorOut) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *SeoSeoCompetitorOut) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetTotal

`func (o *SeoSeoCompetitorOut) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SeoSeoCompetitorOut) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SeoSeoCompetitorOut) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SeoSeoCompetitorOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


