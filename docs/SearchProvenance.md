# SearchProvenance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backend** | Pointer to **string** | Backend is the leg that contributed this match: \&quot;index\&quot; (lexical), \&quot;vector\&quot; (semantic), \&quot;code\&quot; (the org&#39;s repositories) or \&quot;rerank\&quot; (the cross-encoder pass, whose Score is the relevance it assigned). It is the same name that leg reports itself under in Fusion.Backends, so a hit can be traced to a status. | [optional] 
**Rank** | Pointer to **int64** | Rank is this document&#39;s 1-based position in THAT leg&#39;s own result list, before fusion — 1 is the leg&#39;s best hit. It is the only input to the fused score: RRF adds 1/(60+rank) per leg, which is why a document two legs ranked second beats one a single leg ranked first. | [optional] 
**Score** | Pointer to **float64** | Score is the leg&#39;s NATIVE score, on that leg&#39;s own scale, reported for explanation and never used in ranking — the scales are incomparable (a cosine similarity against a term-match count), which is why fusion works on ranks. The vector leg reports the cosine similarity of the document&#39;s best-matching passage; the lexical leg exposes no per-row score and reports 0, meaning \&quot;unscored\&quot;, not \&quot;scored zero\&quot;. | [optional] 

## Methods

### NewSearchProvenance

`func NewSearchProvenance() *SearchProvenance`

NewSearchProvenance instantiates a new SearchProvenance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchProvenanceWithDefaults

`func NewSearchProvenanceWithDefaults() *SearchProvenance`

NewSearchProvenanceWithDefaults instantiates a new SearchProvenance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackend

`func (o *SearchProvenance) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *SearchProvenance) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *SearchProvenance) SetBackend(v string)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *SearchProvenance) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetRank

`func (o *SearchProvenance) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *SearchProvenance) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *SearchProvenance) SetRank(v int64)`

SetRank sets Rank field to given value.

### HasRank

`func (o *SearchProvenance) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetScore

`func (o *SearchProvenance) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SearchProvenance) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SearchProvenance) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *SearchProvenance) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


