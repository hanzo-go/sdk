# SearchSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hits** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**EstimatedTotalHits** | Pointer to **int32** |  | [optional] 
**TotalHits** | Pointer to **int32** |  | [optional] 
**TotalPages** | Pointer to **int32** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**HitsPerPage** | Pointer to **int32** |  | [optional] 
**ProcessingTimeMs** | Pointer to **int32** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**FacetDistribution** | Pointer to **map[string]map[string]int32** |  | [optional] 
**FacetStats** | Pointer to [**map[string]SearchSearchResultFacetStatsValue**](SearchSearchResultFacetStatsValue.md) |  | [optional] 
**SemanticHitCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewSearchSearchResult

`func NewSearchSearchResult() *SearchSearchResult`

NewSearchSearchResult instantiates a new SearchSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSearchResultWithDefaults

`func NewSearchSearchResultWithDefaults() *SearchSearchResult`

NewSearchSearchResultWithDefaults instantiates a new SearchSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHits

`func (o *SearchSearchResult) GetHits() []map[string]interface{}`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *SearchSearchResult) GetHitsOk() (*[]map[string]interface{}, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *SearchSearchResult) SetHits(v []map[string]interface{})`

SetHits sets Hits field to given value.

### HasHits

`func (o *SearchSearchResult) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetOffset

`func (o *SearchSearchResult) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchSearchResult) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchSearchResult) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SearchSearchResult) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *SearchSearchResult) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchSearchResult) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchSearchResult) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchSearchResult) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetEstimatedTotalHits

`func (o *SearchSearchResult) GetEstimatedTotalHits() int32`

GetEstimatedTotalHits returns the EstimatedTotalHits field if non-nil, zero value otherwise.

### GetEstimatedTotalHitsOk

`func (o *SearchSearchResult) GetEstimatedTotalHitsOk() (*int32, bool)`

GetEstimatedTotalHitsOk returns a tuple with the EstimatedTotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTotalHits

`func (o *SearchSearchResult) SetEstimatedTotalHits(v int32)`

SetEstimatedTotalHits sets EstimatedTotalHits field to given value.

### HasEstimatedTotalHits

`func (o *SearchSearchResult) HasEstimatedTotalHits() bool`

HasEstimatedTotalHits returns a boolean if a field has been set.

### GetTotalHits

`func (o *SearchSearchResult) GetTotalHits() int32`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *SearchSearchResult) GetTotalHitsOk() (*int32, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *SearchSearchResult) SetTotalHits(v int32)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *SearchSearchResult) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.

### GetTotalPages

`func (o *SearchSearchResult) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *SearchSearchResult) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *SearchSearchResult) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *SearchSearchResult) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetPage

`func (o *SearchSearchResult) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SearchSearchResult) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SearchSearchResult) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *SearchSearchResult) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetHitsPerPage

`func (o *SearchSearchResult) GetHitsPerPage() int32`

GetHitsPerPage returns the HitsPerPage field if non-nil, zero value otherwise.

### GetHitsPerPageOk

`func (o *SearchSearchResult) GetHitsPerPageOk() (*int32, bool)`

GetHitsPerPageOk returns a tuple with the HitsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitsPerPage

`func (o *SearchSearchResult) SetHitsPerPage(v int32)`

SetHitsPerPage sets HitsPerPage field to given value.

### HasHitsPerPage

`func (o *SearchSearchResult) HasHitsPerPage() bool`

HasHitsPerPage returns a boolean if a field has been set.

### GetProcessingTimeMs

`func (o *SearchSearchResult) GetProcessingTimeMs() int32`

GetProcessingTimeMs returns the ProcessingTimeMs field if non-nil, zero value otherwise.

### GetProcessingTimeMsOk

`func (o *SearchSearchResult) GetProcessingTimeMsOk() (*int32, bool)`

GetProcessingTimeMsOk returns a tuple with the ProcessingTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeMs

`func (o *SearchSearchResult) SetProcessingTimeMs(v int32)`

SetProcessingTimeMs sets ProcessingTimeMs field to given value.

### HasProcessingTimeMs

`func (o *SearchSearchResult) HasProcessingTimeMs() bool`

HasProcessingTimeMs returns a boolean if a field has been set.

### GetQuery

`func (o *SearchSearchResult) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchSearchResult) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchSearchResult) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SearchSearchResult) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetFacetDistribution

`func (o *SearchSearchResult) GetFacetDistribution() map[string]map[string]int32`

GetFacetDistribution returns the FacetDistribution field if non-nil, zero value otherwise.

### GetFacetDistributionOk

`func (o *SearchSearchResult) GetFacetDistributionOk() (*map[string]map[string]int32, bool)`

GetFacetDistributionOk returns a tuple with the FacetDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetDistribution

`func (o *SearchSearchResult) SetFacetDistribution(v map[string]map[string]int32)`

SetFacetDistribution sets FacetDistribution field to given value.

### HasFacetDistribution

`func (o *SearchSearchResult) HasFacetDistribution() bool`

HasFacetDistribution returns a boolean if a field has been set.

### GetFacetStats

`func (o *SearchSearchResult) GetFacetStats() map[string]SearchSearchResultFacetStatsValue`

GetFacetStats returns the FacetStats field if non-nil, zero value otherwise.

### GetFacetStatsOk

`func (o *SearchSearchResult) GetFacetStatsOk() (*map[string]SearchSearchResultFacetStatsValue, bool)`

GetFacetStatsOk returns a tuple with the FacetStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetStats

`func (o *SearchSearchResult) SetFacetStats(v map[string]SearchSearchResultFacetStatsValue)`

SetFacetStats sets FacetStats field to given value.

### HasFacetStats

`func (o *SearchSearchResult) HasFacetStats() bool`

HasFacetStats returns a boolean if a field has been set.

### GetSemanticHitCount

`func (o *SearchSearchResult) GetSemanticHitCount() int32`

GetSemanticHitCount returns the SemanticHitCount field if non-nil, zero value otherwise.

### GetSemanticHitCountOk

`func (o *SearchSearchResult) GetSemanticHitCountOk() (*int32, bool)`

GetSemanticHitCountOk returns a tuple with the SemanticHitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemanticHitCount

`func (o *SearchSearchResult) SetSemanticHitCount(v int32)`

SetSemanticHitCount sets SemanticHitCount field to given value.

### HasSemanticHitCount

`func (o *SearchSearchResult) HasSemanticHitCount() bool`

HasSemanticHitCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


