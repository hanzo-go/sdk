# SearchFederatedSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hits** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ProcessingTimeMs** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**EstimatedTotalHits** | Pointer to **int32** |  | [optional] 
**SemanticHitCount** | Pointer to **int32** |  | [optional] 
**FacetDistribution** | Pointer to **map[string]interface{}** |  | [optional] 
**FacetStats** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSearchFederatedSearchResult

`func NewSearchFederatedSearchResult() *SearchFederatedSearchResult`

NewSearchFederatedSearchResult instantiates a new SearchFederatedSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchFederatedSearchResultWithDefaults

`func NewSearchFederatedSearchResultWithDefaults() *SearchFederatedSearchResult`

NewSearchFederatedSearchResultWithDefaults instantiates a new SearchFederatedSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHits

`func (o *SearchFederatedSearchResult) GetHits() []map[string]interface{}`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *SearchFederatedSearchResult) GetHitsOk() (*[]map[string]interface{}, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *SearchFederatedSearchResult) SetHits(v []map[string]interface{})`

SetHits sets Hits field to given value.

### HasHits

`func (o *SearchFederatedSearchResult) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetProcessingTimeMs

`func (o *SearchFederatedSearchResult) GetProcessingTimeMs() int32`

GetProcessingTimeMs returns the ProcessingTimeMs field if non-nil, zero value otherwise.

### GetProcessingTimeMsOk

`func (o *SearchFederatedSearchResult) GetProcessingTimeMsOk() (*int32, bool)`

GetProcessingTimeMsOk returns a tuple with the ProcessingTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeMs

`func (o *SearchFederatedSearchResult) SetProcessingTimeMs(v int32)`

SetProcessingTimeMs sets ProcessingTimeMs field to given value.

### HasProcessingTimeMs

`func (o *SearchFederatedSearchResult) HasProcessingTimeMs() bool`

HasProcessingTimeMs returns a boolean if a field has been set.

### GetLimit

`func (o *SearchFederatedSearchResult) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchFederatedSearchResult) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchFederatedSearchResult) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchFederatedSearchResult) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *SearchFederatedSearchResult) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchFederatedSearchResult) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchFederatedSearchResult) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SearchFederatedSearchResult) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetEstimatedTotalHits

`func (o *SearchFederatedSearchResult) GetEstimatedTotalHits() int32`

GetEstimatedTotalHits returns the EstimatedTotalHits field if non-nil, zero value otherwise.

### GetEstimatedTotalHitsOk

`func (o *SearchFederatedSearchResult) GetEstimatedTotalHitsOk() (*int32, bool)`

GetEstimatedTotalHitsOk returns a tuple with the EstimatedTotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTotalHits

`func (o *SearchFederatedSearchResult) SetEstimatedTotalHits(v int32)`

SetEstimatedTotalHits sets EstimatedTotalHits field to given value.

### HasEstimatedTotalHits

`func (o *SearchFederatedSearchResult) HasEstimatedTotalHits() bool`

HasEstimatedTotalHits returns a boolean if a field has been set.

### GetSemanticHitCount

`func (o *SearchFederatedSearchResult) GetSemanticHitCount() int32`

GetSemanticHitCount returns the SemanticHitCount field if non-nil, zero value otherwise.

### GetSemanticHitCountOk

`func (o *SearchFederatedSearchResult) GetSemanticHitCountOk() (*int32, bool)`

GetSemanticHitCountOk returns a tuple with the SemanticHitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemanticHitCount

`func (o *SearchFederatedSearchResult) SetSemanticHitCount(v int32)`

SetSemanticHitCount sets SemanticHitCount field to given value.

### HasSemanticHitCount

`func (o *SearchFederatedSearchResult) HasSemanticHitCount() bool`

HasSemanticHitCount returns a boolean if a field has been set.

### GetFacetDistribution

`func (o *SearchFederatedSearchResult) GetFacetDistribution() map[string]interface{}`

GetFacetDistribution returns the FacetDistribution field if non-nil, zero value otherwise.

### GetFacetDistributionOk

`func (o *SearchFederatedSearchResult) GetFacetDistributionOk() (*map[string]interface{}, bool)`

GetFacetDistributionOk returns a tuple with the FacetDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetDistribution

`func (o *SearchFederatedSearchResult) SetFacetDistribution(v map[string]interface{})`

SetFacetDistribution sets FacetDistribution field to given value.

### HasFacetDistribution

`func (o *SearchFederatedSearchResult) HasFacetDistribution() bool`

HasFacetDistribution returns a boolean if a field has been set.

### GetFacetStats

`func (o *SearchFederatedSearchResult) GetFacetStats() map[string]interface{}`

GetFacetStats returns the FacetStats field if non-nil, zero value otherwise.

### GetFacetStatsOk

`func (o *SearchFederatedSearchResult) GetFacetStatsOk() (*map[string]interface{}, bool)`

GetFacetStatsOk returns a tuple with the FacetStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetStats

`func (o *SearchFederatedSearchResult) SetFacetStats(v map[string]interface{})`

SetFacetStats sets FacetStats field to given value.

### HasFacetStats

`func (o *SearchFederatedSearchResult) HasFacetStats() bool`

HasFacetStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


