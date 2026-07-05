# SearchSearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Q** | Pointer to **string** | Search query text | [optional] 
**Offset** | Pointer to **int32** |  | [optional] [default to 0]
**Limit** | Pointer to **int32** |  | [optional] [default to 20]
**Page** | Pointer to **int32** |  | [optional] 
**HitsPerPage** | Pointer to **int32** |  | [optional] 
**AttributesToRetrieve** | Pointer to **[]string** |  | [optional] 
**AttributesToHighlight** | Pointer to **[]string** |  | [optional] 
**AttributesToCrop** | Pointer to **[]string** |  | [optional] 
**CropLength** | Pointer to **int32** |  | [optional] [default to 10]
**CropMarker** | Pointer to **string** |  | [optional] [default to "..."]
**HighlightPreTag** | Pointer to **string** |  | [optional] [default to "<em>"]
**HighlightPostTag** | Pointer to **string** |  | [optional] [default to "</em>"]
**Filter** | Pointer to [**SearchSearchQueryFilter**](SearchSearchQueryFilter.md) |  | [optional] 
**Sort** | Pointer to **[]string** |  | [optional] 
**Facets** | Pointer to **[]string** |  | [optional] 
**ShowMatchesPosition** | Pointer to **bool** |  | [optional] 
**ShowRankingScore** | Pointer to **bool** |  | [optional] 
**ShowRankingScoreDetails** | Pointer to **bool** |  | [optional] 
**MatchingStrategy** | Pointer to **string** |  | [optional] 
**RankingScoreThreshold** | Pointer to **float32** |  | [optional] 
**Distinct** | Pointer to **string** |  | [optional] 
**Vector** | Pointer to **[]float32** |  | [optional] 
**Hybrid** | Pointer to [**SearchSearchQueryHybrid**](SearchSearchQueryHybrid.md) |  | [optional] 
**RetrieveVectors** | Pointer to **bool** |  | [optional] 
**Locales** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSearchSearchQuery

`func NewSearchSearchQuery() *SearchSearchQuery`

NewSearchSearchQuery instantiates a new SearchSearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSearchQueryWithDefaults

`func NewSearchSearchQueryWithDefaults() *SearchSearchQuery`

NewSearchSearchQueryWithDefaults instantiates a new SearchSearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQ

`func (o *SearchSearchQuery) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *SearchSearchQuery) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *SearchSearchQuery) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *SearchSearchQuery) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetOffset

`func (o *SearchSearchQuery) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchSearchQuery) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchSearchQuery) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SearchSearchQuery) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *SearchSearchQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchSearchQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchSearchQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchSearchQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetPage

`func (o *SearchSearchQuery) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SearchSearchQuery) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SearchSearchQuery) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *SearchSearchQuery) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetHitsPerPage

`func (o *SearchSearchQuery) GetHitsPerPage() int32`

GetHitsPerPage returns the HitsPerPage field if non-nil, zero value otherwise.

### GetHitsPerPageOk

`func (o *SearchSearchQuery) GetHitsPerPageOk() (*int32, bool)`

GetHitsPerPageOk returns a tuple with the HitsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHitsPerPage

`func (o *SearchSearchQuery) SetHitsPerPage(v int32)`

SetHitsPerPage sets HitsPerPage field to given value.

### HasHitsPerPage

`func (o *SearchSearchQuery) HasHitsPerPage() bool`

HasHitsPerPage returns a boolean if a field has been set.

### GetAttributesToRetrieve

`func (o *SearchSearchQuery) GetAttributesToRetrieve() []string`

GetAttributesToRetrieve returns the AttributesToRetrieve field if non-nil, zero value otherwise.

### GetAttributesToRetrieveOk

`func (o *SearchSearchQuery) GetAttributesToRetrieveOk() (*[]string, bool)`

GetAttributesToRetrieveOk returns a tuple with the AttributesToRetrieve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesToRetrieve

`func (o *SearchSearchQuery) SetAttributesToRetrieve(v []string)`

SetAttributesToRetrieve sets AttributesToRetrieve field to given value.

### HasAttributesToRetrieve

`func (o *SearchSearchQuery) HasAttributesToRetrieve() bool`

HasAttributesToRetrieve returns a boolean if a field has been set.

### GetAttributesToHighlight

`func (o *SearchSearchQuery) GetAttributesToHighlight() []string`

GetAttributesToHighlight returns the AttributesToHighlight field if non-nil, zero value otherwise.

### GetAttributesToHighlightOk

`func (o *SearchSearchQuery) GetAttributesToHighlightOk() (*[]string, bool)`

GetAttributesToHighlightOk returns a tuple with the AttributesToHighlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesToHighlight

`func (o *SearchSearchQuery) SetAttributesToHighlight(v []string)`

SetAttributesToHighlight sets AttributesToHighlight field to given value.

### HasAttributesToHighlight

`func (o *SearchSearchQuery) HasAttributesToHighlight() bool`

HasAttributesToHighlight returns a boolean if a field has been set.

### GetAttributesToCrop

`func (o *SearchSearchQuery) GetAttributesToCrop() []string`

GetAttributesToCrop returns the AttributesToCrop field if non-nil, zero value otherwise.

### GetAttributesToCropOk

`func (o *SearchSearchQuery) GetAttributesToCropOk() (*[]string, bool)`

GetAttributesToCropOk returns a tuple with the AttributesToCrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesToCrop

`func (o *SearchSearchQuery) SetAttributesToCrop(v []string)`

SetAttributesToCrop sets AttributesToCrop field to given value.

### HasAttributesToCrop

`func (o *SearchSearchQuery) HasAttributesToCrop() bool`

HasAttributesToCrop returns a boolean if a field has been set.

### GetCropLength

`func (o *SearchSearchQuery) GetCropLength() int32`

GetCropLength returns the CropLength field if non-nil, zero value otherwise.

### GetCropLengthOk

`func (o *SearchSearchQuery) GetCropLengthOk() (*int32, bool)`

GetCropLengthOk returns a tuple with the CropLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCropLength

`func (o *SearchSearchQuery) SetCropLength(v int32)`

SetCropLength sets CropLength field to given value.

### HasCropLength

`func (o *SearchSearchQuery) HasCropLength() bool`

HasCropLength returns a boolean if a field has been set.

### GetCropMarker

`func (o *SearchSearchQuery) GetCropMarker() string`

GetCropMarker returns the CropMarker field if non-nil, zero value otherwise.

### GetCropMarkerOk

`func (o *SearchSearchQuery) GetCropMarkerOk() (*string, bool)`

GetCropMarkerOk returns a tuple with the CropMarker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCropMarker

`func (o *SearchSearchQuery) SetCropMarker(v string)`

SetCropMarker sets CropMarker field to given value.

### HasCropMarker

`func (o *SearchSearchQuery) HasCropMarker() bool`

HasCropMarker returns a boolean if a field has been set.

### GetHighlightPreTag

`func (o *SearchSearchQuery) GetHighlightPreTag() string`

GetHighlightPreTag returns the HighlightPreTag field if non-nil, zero value otherwise.

### GetHighlightPreTagOk

`func (o *SearchSearchQuery) GetHighlightPreTagOk() (*string, bool)`

GetHighlightPreTagOk returns a tuple with the HighlightPreTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightPreTag

`func (o *SearchSearchQuery) SetHighlightPreTag(v string)`

SetHighlightPreTag sets HighlightPreTag field to given value.

### HasHighlightPreTag

`func (o *SearchSearchQuery) HasHighlightPreTag() bool`

HasHighlightPreTag returns a boolean if a field has been set.

### GetHighlightPostTag

`func (o *SearchSearchQuery) GetHighlightPostTag() string`

GetHighlightPostTag returns the HighlightPostTag field if non-nil, zero value otherwise.

### GetHighlightPostTagOk

`func (o *SearchSearchQuery) GetHighlightPostTagOk() (*string, bool)`

GetHighlightPostTagOk returns a tuple with the HighlightPostTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighlightPostTag

`func (o *SearchSearchQuery) SetHighlightPostTag(v string)`

SetHighlightPostTag sets HighlightPostTag field to given value.

### HasHighlightPostTag

`func (o *SearchSearchQuery) HasHighlightPostTag() bool`

HasHighlightPostTag returns a boolean if a field has been set.

### GetFilter

`func (o *SearchSearchQuery) GetFilter() SearchSearchQueryFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SearchSearchQuery) GetFilterOk() (*SearchSearchQueryFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SearchSearchQuery) SetFilter(v SearchSearchQueryFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *SearchSearchQuery) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *SearchSearchQuery) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *SearchSearchQuery) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *SearchSearchQuery) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *SearchSearchQuery) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetFacets

`func (o *SearchSearchQuery) GetFacets() []string`

GetFacets returns the Facets field if non-nil, zero value otherwise.

### GetFacetsOk

`func (o *SearchSearchQuery) GetFacetsOk() (*[]string, bool)`

GetFacetsOk returns a tuple with the Facets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacets

`func (o *SearchSearchQuery) SetFacets(v []string)`

SetFacets sets Facets field to given value.

### HasFacets

`func (o *SearchSearchQuery) HasFacets() bool`

HasFacets returns a boolean if a field has been set.

### GetShowMatchesPosition

`func (o *SearchSearchQuery) GetShowMatchesPosition() bool`

GetShowMatchesPosition returns the ShowMatchesPosition field if non-nil, zero value otherwise.

### GetShowMatchesPositionOk

`func (o *SearchSearchQuery) GetShowMatchesPositionOk() (*bool, bool)`

GetShowMatchesPositionOk returns a tuple with the ShowMatchesPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowMatchesPosition

`func (o *SearchSearchQuery) SetShowMatchesPosition(v bool)`

SetShowMatchesPosition sets ShowMatchesPosition field to given value.

### HasShowMatchesPosition

`func (o *SearchSearchQuery) HasShowMatchesPosition() bool`

HasShowMatchesPosition returns a boolean if a field has been set.

### GetShowRankingScore

`func (o *SearchSearchQuery) GetShowRankingScore() bool`

GetShowRankingScore returns the ShowRankingScore field if non-nil, zero value otherwise.

### GetShowRankingScoreOk

`func (o *SearchSearchQuery) GetShowRankingScoreOk() (*bool, bool)`

GetShowRankingScoreOk returns a tuple with the ShowRankingScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowRankingScore

`func (o *SearchSearchQuery) SetShowRankingScore(v bool)`

SetShowRankingScore sets ShowRankingScore field to given value.

### HasShowRankingScore

`func (o *SearchSearchQuery) HasShowRankingScore() bool`

HasShowRankingScore returns a boolean if a field has been set.

### GetShowRankingScoreDetails

`func (o *SearchSearchQuery) GetShowRankingScoreDetails() bool`

GetShowRankingScoreDetails returns the ShowRankingScoreDetails field if non-nil, zero value otherwise.

### GetShowRankingScoreDetailsOk

`func (o *SearchSearchQuery) GetShowRankingScoreDetailsOk() (*bool, bool)`

GetShowRankingScoreDetailsOk returns a tuple with the ShowRankingScoreDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowRankingScoreDetails

`func (o *SearchSearchQuery) SetShowRankingScoreDetails(v bool)`

SetShowRankingScoreDetails sets ShowRankingScoreDetails field to given value.

### HasShowRankingScoreDetails

`func (o *SearchSearchQuery) HasShowRankingScoreDetails() bool`

HasShowRankingScoreDetails returns a boolean if a field has been set.

### GetMatchingStrategy

`func (o *SearchSearchQuery) GetMatchingStrategy() string`

GetMatchingStrategy returns the MatchingStrategy field if non-nil, zero value otherwise.

### GetMatchingStrategyOk

`func (o *SearchSearchQuery) GetMatchingStrategyOk() (*string, bool)`

GetMatchingStrategyOk returns a tuple with the MatchingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingStrategy

`func (o *SearchSearchQuery) SetMatchingStrategy(v string)`

SetMatchingStrategy sets MatchingStrategy field to given value.

### HasMatchingStrategy

`func (o *SearchSearchQuery) HasMatchingStrategy() bool`

HasMatchingStrategy returns a boolean if a field has been set.

### GetRankingScoreThreshold

`func (o *SearchSearchQuery) GetRankingScoreThreshold() float32`

GetRankingScoreThreshold returns the RankingScoreThreshold field if non-nil, zero value otherwise.

### GetRankingScoreThresholdOk

`func (o *SearchSearchQuery) GetRankingScoreThresholdOk() (*float32, bool)`

GetRankingScoreThresholdOk returns a tuple with the RankingScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankingScoreThreshold

`func (o *SearchSearchQuery) SetRankingScoreThreshold(v float32)`

SetRankingScoreThreshold sets RankingScoreThreshold field to given value.

### HasRankingScoreThreshold

`func (o *SearchSearchQuery) HasRankingScoreThreshold() bool`

HasRankingScoreThreshold returns a boolean if a field has been set.

### GetDistinct

`func (o *SearchSearchQuery) GetDistinct() string`

GetDistinct returns the Distinct field if non-nil, zero value otherwise.

### GetDistinctOk

`func (o *SearchSearchQuery) GetDistinctOk() (*string, bool)`

GetDistinctOk returns a tuple with the Distinct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinct

`func (o *SearchSearchQuery) SetDistinct(v string)`

SetDistinct sets Distinct field to given value.

### HasDistinct

`func (o *SearchSearchQuery) HasDistinct() bool`

HasDistinct returns a boolean if a field has been set.

### GetVector

`func (o *SearchSearchQuery) GetVector() []float32`

GetVector returns the Vector field if non-nil, zero value otherwise.

### GetVectorOk

`func (o *SearchSearchQuery) GetVectorOk() (*[]float32, bool)`

GetVectorOk returns a tuple with the Vector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVector

`func (o *SearchSearchQuery) SetVector(v []float32)`

SetVector sets Vector field to given value.

### HasVector

`func (o *SearchSearchQuery) HasVector() bool`

HasVector returns a boolean if a field has been set.

### GetHybrid

`func (o *SearchSearchQuery) GetHybrid() SearchSearchQueryHybrid`

GetHybrid returns the Hybrid field if non-nil, zero value otherwise.

### GetHybridOk

`func (o *SearchSearchQuery) GetHybridOk() (*SearchSearchQueryHybrid, bool)`

GetHybridOk returns a tuple with the Hybrid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHybrid

`func (o *SearchSearchQuery) SetHybrid(v SearchSearchQueryHybrid)`

SetHybrid sets Hybrid field to given value.

### HasHybrid

`func (o *SearchSearchQuery) HasHybrid() bool`

HasHybrid returns a boolean if a field has been set.

### GetRetrieveVectors

`func (o *SearchSearchQuery) GetRetrieveVectors() bool`

GetRetrieveVectors returns the RetrieveVectors field if non-nil, zero value otherwise.

### GetRetrieveVectorsOk

`func (o *SearchSearchQuery) GetRetrieveVectorsOk() (*bool, bool)`

GetRetrieveVectorsOk returns a tuple with the RetrieveVectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrieveVectors

`func (o *SearchSearchQuery) SetRetrieveVectors(v bool)`

SetRetrieveVectors sets RetrieveVectors field to given value.

### HasRetrieveVectors

`func (o *SearchSearchQuery) HasRetrieveVectors() bool`

HasRetrieveVectors returns a boolean if a field has been set.

### GetLocales

`func (o *SearchSearchQuery) GetLocales() []string`

GetLocales returns the Locales field if non-nil, zero value otherwise.

### GetLocalesOk

`func (o *SearchSearchQuery) GetLocalesOk() (*[]string, bool)`

GetLocalesOk returns a tuple with the Locales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocales

`func (o *SearchSearchQuery) SetLocales(v []string)`

SetLocales sets Locales field to given value.

### HasLocales

`func (o *SearchSearchQuery) HasLocales() bool`

HasLocales returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


