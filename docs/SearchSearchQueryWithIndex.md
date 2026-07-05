# SearchSearchQueryWithIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexUid** | **string** |  | 
**Q** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Filter** | Pointer to [**SearchSearchQueryWithIndexFilter**](SearchSearchQueryWithIndexFilter.md) |  | [optional] 
**Sort** | Pointer to **[]string** |  | [optional] 
**Facets** | Pointer to **[]string** |  | [optional] 
**AttributesToRetrieve** | Pointer to **[]string** |  | [optional] 
**ShowMatchesPosition** | Pointer to **bool** |  | [optional] 
**ShowRankingScore** | Pointer to **bool** |  | [optional] 
**MatchingStrategy** | Pointer to **string** |  | [optional] 
**FederationOptions** | Pointer to [**SearchSearchQueryWithIndexFederationOptions**](SearchSearchQueryWithIndexFederationOptions.md) |  | [optional] 

## Methods

### NewSearchSearchQueryWithIndex

`func NewSearchSearchQueryWithIndex(indexUid string, ) *SearchSearchQueryWithIndex`

NewSearchSearchQueryWithIndex instantiates a new SearchSearchQueryWithIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSearchQueryWithIndexWithDefaults

`func NewSearchSearchQueryWithIndexWithDefaults() *SearchSearchQueryWithIndex`

NewSearchSearchQueryWithIndexWithDefaults instantiates a new SearchSearchQueryWithIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexUid

`func (o *SearchSearchQueryWithIndex) GetIndexUid() string`

GetIndexUid returns the IndexUid field if non-nil, zero value otherwise.

### GetIndexUidOk

`func (o *SearchSearchQueryWithIndex) GetIndexUidOk() (*string, bool)`

GetIndexUidOk returns a tuple with the IndexUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexUid

`func (o *SearchSearchQueryWithIndex) SetIndexUid(v string)`

SetIndexUid sets IndexUid field to given value.


### GetQ

`func (o *SearchSearchQueryWithIndex) GetQ() string`

GetQ returns the Q field if non-nil, zero value otherwise.

### GetQOk

`func (o *SearchSearchQueryWithIndex) GetQOk() (*string, bool)`

GetQOk returns a tuple with the Q field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQ

`func (o *SearchSearchQueryWithIndex) SetQ(v string)`

SetQ sets Q field to given value.

### HasQ

`func (o *SearchSearchQueryWithIndex) HasQ() bool`

HasQ returns a boolean if a field has been set.

### GetOffset

`func (o *SearchSearchQueryWithIndex) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchSearchQueryWithIndex) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchSearchQueryWithIndex) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SearchSearchQueryWithIndex) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *SearchSearchQueryWithIndex) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchSearchQueryWithIndex) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchSearchQueryWithIndex) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchSearchQueryWithIndex) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetFilter

`func (o *SearchSearchQueryWithIndex) GetFilter() SearchSearchQueryWithIndexFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SearchSearchQueryWithIndex) GetFilterOk() (*SearchSearchQueryWithIndexFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SearchSearchQueryWithIndex) SetFilter(v SearchSearchQueryWithIndexFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *SearchSearchQueryWithIndex) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *SearchSearchQueryWithIndex) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *SearchSearchQueryWithIndex) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *SearchSearchQueryWithIndex) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *SearchSearchQueryWithIndex) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetFacets

`func (o *SearchSearchQueryWithIndex) GetFacets() []string`

GetFacets returns the Facets field if non-nil, zero value otherwise.

### GetFacetsOk

`func (o *SearchSearchQueryWithIndex) GetFacetsOk() (*[]string, bool)`

GetFacetsOk returns a tuple with the Facets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacets

`func (o *SearchSearchQueryWithIndex) SetFacets(v []string)`

SetFacets sets Facets field to given value.

### HasFacets

`func (o *SearchSearchQueryWithIndex) HasFacets() bool`

HasFacets returns a boolean if a field has been set.

### GetAttributesToRetrieve

`func (o *SearchSearchQueryWithIndex) GetAttributesToRetrieve() []string`

GetAttributesToRetrieve returns the AttributesToRetrieve field if non-nil, zero value otherwise.

### GetAttributesToRetrieveOk

`func (o *SearchSearchQueryWithIndex) GetAttributesToRetrieveOk() (*[]string, bool)`

GetAttributesToRetrieveOk returns a tuple with the AttributesToRetrieve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesToRetrieve

`func (o *SearchSearchQueryWithIndex) SetAttributesToRetrieve(v []string)`

SetAttributesToRetrieve sets AttributesToRetrieve field to given value.

### HasAttributesToRetrieve

`func (o *SearchSearchQueryWithIndex) HasAttributesToRetrieve() bool`

HasAttributesToRetrieve returns a boolean if a field has been set.

### GetShowMatchesPosition

`func (o *SearchSearchQueryWithIndex) GetShowMatchesPosition() bool`

GetShowMatchesPosition returns the ShowMatchesPosition field if non-nil, zero value otherwise.

### GetShowMatchesPositionOk

`func (o *SearchSearchQueryWithIndex) GetShowMatchesPositionOk() (*bool, bool)`

GetShowMatchesPositionOk returns a tuple with the ShowMatchesPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowMatchesPosition

`func (o *SearchSearchQueryWithIndex) SetShowMatchesPosition(v bool)`

SetShowMatchesPosition sets ShowMatchesPosition field to given value.

### HasShowMatchesPosition

`func (o *SearchSearchQueryWithIndex) HasShowMatchesPosition() bool`

HasShowMatchesPosition returns a boolean if a field has been set.

### GetShowRankingScore

`func (o *SearchSearchQueryWithIndex) GetShowRankingScore() bool`

GetShowRankingScore returns the ShowRankingScore field if non-nil, zero value otherwise.

### GetShowRankingScoreOk

`func (o *SearchSearchQueryWithIndex) GetShowRankingScoreOk() (*bool, bool)`

GetShowRankingScoreOk returns a tuple with the ShowRankingScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowRankingScore

`func (o *SearchSearchQueryWithIndex) SetShowRankingScore(v bool)`

SetShowRankingScore sets ShowRankingScore field to given value.

### HasShowRankingScore

`func (o *SearchSearchQueryWithIndex) HasShowRankingScore() bool`

HasShowRankingScore returns a boolean if a field has been set.

### GetMatchingStrategy

`func (o *SearchSearchQueryWithIndex) GetMatchingStrategy() string`

GetMatchingStrategy returns the MatchingStrategy field if non-nil, zero value otherwise.

### GetMatchingStrategyOk

`func (o *SearchSearchQueryWithIndex) GetMatchingStrategyOk() (*string, bool)`

GetMatchingStrategyOk returns a tuple with the MatchingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingStrategy

`func (o *SearchSearchQueryWithIndex) SetMatchingStrategy(v string)`

SetMatchingStrategy sets MatchingStrategy field to given value.

### HasMatchingStrategy

`func (o *SearchSearchQueryWithIndex) HasMatchingStrategy() bool`

HasMatchingStrategy returns a boolean if a field has been set.

### GetFederationOptions

`func (o *SearchSearchQueryWithIndex) GetFederationOptions() SearchSearchQueryWithIndexFederationOptions`

GetFederationOptions returns the FederationOptions field if non-nil, zero value otherwise.

### GetFederationOptionsOk

`func (o *SearchSearchQueryWithIndex) GetFederationOptionsOk() (*SearchSearchQueryWithIndexFederationOptions, bool)`

GetFederationOptionsOk returns a tuple with the FederationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationOptions

`func (o *SearchSearchQueryWithIndex) SetFederationOptions(v SearchSearchQueryWithIndexFederationOptions)`

SetFederationOptions sets FederationOptions field to given value.

### HasFederationOptions

`func (o *SearchSearchQueryWithIndex) HasFederationOptions() bool`

HasFederationOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


