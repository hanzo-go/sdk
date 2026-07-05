# SearchSimilarQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**SearchSimilarQueryId**](SearchSimilarQueryId.md) |  | 
**Embedder** | Pointer to **string** |  | [optional] 
**AttributesToRetrieve** | Pointer to **[]string** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] [default to 0]
**Limit** | Pointer to **int32** |  | [optional] [default to 20]
**Filter** | Pointer to [**SearchSearchQueryWithIndexFilter**](SearchSearchQueryWithIndexFilter.md) |  | [optional] 
**ShowRankingScore** | Pointer to **bool** |  | [optional] 
**ShowRankingScoreDetails** | Pointer to **bool** |  | [optional] 
**RankingScoreThreshold** | Pointer to **float32** |  | [optional] 
**RetrieveVectors** | Pointer to **bool** |  | [optional] 

## Methods

### NewSearchSimilarQuery

`func NewSearchSimilarQuery(id SearchSimilarQueryId, ) *SearchSimilarQuery`

NewSearchSimilarQuery instantiates a new SearchSimilarQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSimilarQueryWithDefaults

`func NewSearchSimilarQueryWithDefaults() *SearchSimilarQuery`

NewSearchSimilarQueryWithDefaults instantiates a new SearchSimilarQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchSimilarQuery) GetId() SearchSimilarQueryId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchSimilarQuery) GetIdOk() (*SearchSimilarQueryId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchSimilarQuery) SetId(v SearchSimilarQueryId)`

SetId sets Id field to given value.


### GetEmbedder

`func (o *SearchSimilarQuery) GetEmbedder() string`

GetEmbedder returns the Embedder field if non-nil, zero value otherwise.

### GetEmbedderOk

`func (o *SearchSimilarQuery) GetEmbedderOk() (*string, bool)`

GetEmbedderOk returns a tuple with the Embedder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedder

`func (o *SearchSimilarQuery) SetEmbedder(v string)`

SetEmbedder sets Embedder field to given value.

### HasEmbedder

`func (o *SearchSimilarQuery) HasEmbedder() bool`

HasEmbedder returns a boolean if a field has been set.

### GetAttributesToRetrieve

`func (o *SearchSimilarQuery) GetAttributesToRetrieve() []string`

GetAttributesToRetrieve returns the AttributesToRetrieve field if non-nil, zero value otherwise.

### GetAttributesToRetrieveOk

`func (o *SearchSimilarQuery) GetAttributesToRetrieveOk() (*[]string, bool)`

GetAttributesToRetrieveOk returns a tuple with the AttributesToRetrieve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesToRetrieve

`func (o *SearchSimilarQuery) SetAttributesToRetrieve(v []string)`

SetAttributesToRetrieve sets AttributesToRetrieve field to given value.

### HasAttributesToRetrieve

`func (o *SearchSimilarQuery) HasAttributesToRetrieve() bool`

HasAttributesToRetrieve returns a boolean if a field has been set.

### GetOffset

`func (o *SearchSimilarQuery) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchSimilarQuery) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchSimilarQuery) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SearchSimilarQuery) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *SearchSimilarQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchSimilarQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchSimilarQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchSimilarQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetFilter

`func (o *SearchSimilarQuery) GetFilter() SearchSearchQueryWithIndexFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SearchSimilarQuery) GetFilterOk() (*SearchSearchQueryWithIndexFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SearchSimilarQuery) SetFilter(v SearchSearchQueryWithIndexFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *SearchSimilarQuery) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetShowRankingScore

`func (o *SearchSimilarQuery) GetShowRankingScore() bool`

GetShowRankingScore returns the ShowRankingScore field if non-nil, zero value otherwise.

### GetShowRankingScoreOk

`func (o *SearchSimilarQuery) GetShowRankingScoreOk() (*bool, bool)`

GetShowRankingScoreOk returns a tuple with the ShowRankingScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowRankingScore

`func (o *SearchSimilarQuery) SetShowRankingScore(v bool)`

SetShowRankingScore sets ShowRankingScore field to given value.

### HasShowRankingScore

`func (o *SearchSimilarQuery) HasShowRankingScore() bool`

HasShowRankingScore returns a boolean if a field has been set.

### GetShowRankingScoreDetails

`func (o *SearchSimilarQuery) GetShowRankingScoreDetails() bool`

GetShowRankingScoreDetails returns the ShowRankingScoreDetails field if non-nil, zero value otherwise.

### GetShowRankingScoreDetailsOk

`func (o *SearchSimilarQuery) GetShowRankingScoreDetailsOk() (*bool, bool)`

GetShowRankingScoreDetailsOk returns a tuple with the ShowRankingScoreDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowRankingScoreDetails

`func (o *SearchSimilarQuery) SetShowRankingScoreDetails(v bool)`

SetShowRankingScoreDetails sets ShowRankingScoreDetails field to given value.

### HasShowRankingScoreDetails

`func (o *SearchSimilarQuery) HasShowRankingScoreDetails() bool`

HasShowRankingScoreDetails returns a boolean if a field has been set.

### GetRankingScoreThreshold

`func (o *SearchSimilarQuery) GetRankingScoreThreshold() float32`

GetRankingScoreThreshold returns the RankingScoreThreshold field if non-nil, zero value otherwise.

### GetRankingScoreThresholdOk

`func (o *SearchSimilarQuery) GetRankingScoreThresholdOk() (*float32, bool)`

GetRankingScoreThresholdOk returns a tuple with the RankingScoreThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankingScoreThreshold

`func (o *SearchSimilarQuery) SetRankingScoreThreshold(v float32)`

SetRankingScoreThreshold sets RankingScoreThreshold field to given value.

### HasRankingScoreThreshold

`func (o *SearchSimilarQuery) HasRankingScoreThreshold() bool`

HasRankingScoreThreshold returns a boolean if a field has been set.

### GetRetrieveVectors

`func (o *SearchSimilarQuery) GetRetrieveVectors() bool`

GetRetrieveVectors returns the RetrieveVectors field if non-nil, zero value otherwise.

### GetRetrieveVectorsOk

`func (o *SearchSimilarQuery) GetRetrieveVectorsOk() (*bool, bool)`

GetRetrieveVectorsOk returns a tuple with the RetrieveVectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrieveVectors

`func (o *SearchSimilarQuery) SetRetrieveVectors(v bool)`

SetRetrieveVectors sets RetrieveVectors field to given value.

### HasRetrieveVectors

`func (o *SearchSimilarQuery) HasRetrieveVectors() bool`

HasRetrieveVectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


