# SearchMultiSearch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]SearchMultiSearchResultsResultsInner**](SearchMultiSearchResultsResultsInner.md) |  | [optional] 
**Hits** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ProcessingTimeMs** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**EstimatedTotalHits** | Pointer to **int32** |  | [optional] 
**SemanticHitCount** | Pointer to **int32** |  | [optional] 
**FacetDistribution** | Pointer to **map[string]interface{}** |  | [optional] 
**FacetStats** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSearchMultiSearch200Response

`func NewSearchMultiSearch200Response() *SearchMultiSearch200Response`

NewSearchMultiSearch200Response instantiates a new SearchMultiSearch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchMultiSearch200ResponseWithDefaults

`func NewSearchMultiSearch200ResponseWithDefaults() *SearchMultiSearch200Response`

NewSearchMultiSearch200ResponseWithDefaults instantiates a new SearchMultiSearch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *SearchMultiSearch200Response) GetResults() []SearchMultiSearchResultsResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchMultiSearch200Response) GetResultsOk() (*[]SearchMultiSearchResultsResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchMultiSearch200Response) SetResults(v []SearchMultiSearchResultsResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *SearchMultiSearch200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetHits

`func (o *SearchMultiSearch200Response) GetHits() []map[string]interface{}`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *SearchMultiSearch200Response) GetHitsOk() (*[]map[string]interface{}, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *SearchMultiSearch200Response) SetHits(v []map[string]interface{})`

SetHits sets Hits field to given value.

### HasHits

`func (o *SearchMultiSearch200Response) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetProcessingTimeMs

`func (o *SearchMultiSearch200Response) GetProcessingTimeMs() int32`

GetProcessingTimeMs returns the ProcessingTimeMs field if non-nil, zero value otherwise.

### GetProcessingTimeMsOk

`func (o *SearchMultiSearch200Response) GetProcessingTimeMsOk() (*int32, bool)`

GetProcessingTimeMsOk returns a tuple with the ProcessingTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeMs

`func (o *SearchMultiSearch200Response) SetProcessingTimeMs(v int32)`

SetProcessingTimeMs sets ProcessingTimeMs field to given value.

### HasProcessingTimeMs

`func (o *SearchMultiSearch200Response) HasProcessingTimeMs() bool`

HasProcessingTimeMs returns a boolean if a field has been set.

### GetLimit

`func (o *SearchMultiSearch200Response) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchMultiSearch200Response) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchMultiSearch200Response) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchMultiSearch200Response) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *SearchMultiSearch200Response) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchMultiSearch200Response) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchMultiSearch200Response) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SearchMultiSearch200Response) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetEstimatedTotalHits

`func (o *SearchMultiSearch200Response) GetEstimatedTotalHits() int32`

GetEstimatedTotalHits returns the EstimatedTotalHits field if non-nil, zero value otherwise.

### GetEstimatedTotalHitsOk

`func (o *SearchMultiSearch200Response) GetEstimatedTotalHitsOk() (*int32, bool)`

GetEstimatedTotalHitsOk returns a tuple with the EstimatedTotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTotalHits

`func (o *SearchMultiSearch200Response) SetEstimatedTotalHits(v int32)`

SetEstimatedTotalHits sets EstimatedTotalHits field to given value.

### HasEstimatedTotalHits

`func (o *SearchMultiSearch200Response) HasEstimatedTotalHits() bool`

HasEstimatedTotalHits returns a boolean if a field has been set.

### GetSemanticHitCount

`func (o *SearchMultiSearch200Response) GetSemanticHitCount() int32`

GetSemanticHitCount returns the SemanticHitCount field if non-nil, zero value otherwise.

### GetSemanticHitCountOk

`func (o *SearchMultiSearch200Response) GetSemanticHitCountOk() (*int32, bool)`

GetSemanticHitCountOk returns a tuple with the SemanticHitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemanticHitCount

`func (o *SearchMultiSearch200Response) SetSemanticHitCount(v int32)`

SetSemanticHitCount sets SemanticHitCount field to given value.

### HasSemanticHitCount

`func (o *SearchMultiSearch200Response) HasSemanticHitCount() bool`

HasSemanticHitCount returns a boolean if a field has been set.

### GetFacetDistribution

`func (o *SearchMultiSearch200Response) GetFacetDistribution() map[string]interface{}`

GetFacetDistribution returns the FacetDistribution field if non-nil, zero value otherwise.

### GetFacetDistributionOk

`func (o *SearchMultiSearch200Response) GetFacetDistributionOk() (*map[string]interface{}, bool)`

GetFacetDistributionOk returns a tuple with the FacetDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetDistribution

`func (o *SearchMultiSearch200Response) SetFacetDistribution(v map[string]interface{})`

SetFacetDistribution sets FacetDistribution field to given value.

### HasFacetDistribution

`func (o *SearchMultiSearch200Response) HasFacetDistribution() bool`

HasFacetDistribution returns a boolean if a field has been set.

### GetFacetStats

`func (o *SearchMultiSearch200Response) GetFacetStats() map[string]interface{}`

GetFacetStats returns the FacetStats field if non-nil, zero value otherwise.

### GetFacetStatsOk

`func (o *SearchMultiSearch200Response) GetFacetStatsOk() (*map[string]interface{}, bool)`

GetFacetStatsOk returns a tuple with the FacetStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetStats

`func (o *SearchMultiSearch200Response) SetFacetStats(v map[string]interface{})`

SetFacetStats sets FacetStats field to given value.

### HasFacetStats

`func (o *SearchMultiSearch200Response) HasFacetStats() bool`

HasFacetStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


