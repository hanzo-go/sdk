# Match

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backend** | Pointer to **string** | Backend is the leg that contributed this match: \&quot;index\&quot; (lexical), \&quot;vector\&quot; (semantic) or \&quot;code\&quot; (the org&#39;s repositories). It is the same name that leg reports itself under in Response.Backends, so a hit can be traced to a status. | [optional] 
**Rank** | Pointer to **int32** | Rank is this document&#39;s 1-based position in THAT leg&#39;s own result list, before fusion — 1 is the leg&#39;s best hit. It is the only input to the fused score: RRF adds 1/(60+rank) per leg, which is why a document two legs ranked second beats one a single leg ranked first. | [optional] 
**Score** | Pointer to **float32** | Score is the leg&#39;s NATIVE score, on that leg&#39;s own scale, reported for explanation and never used in ranking — the scales are incomparable (a cosine similarity against a term-match count), which is why fusion works on ranks. The vector leg reports Qdrant&#39;s cosine similarity; the lexical leg exposes no per-row score and reports 0, meaning \&quot;unscored\&quot;, not \&quot;scored zero\&quot;. | [optional] 

## Methods

### NewMatch

`func NewMatch() *Match`

NewMatch instantiates a new Match object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchWithDefaults

`func NewMatchWithDefaults() *Match`

NewMatchWithDefaults instantiates a new Match object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackend

`func (o *Match) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *Match) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *Match) SetBackend(v string)`

SetBackend sets Backend field to given value.

### HasBackend

`func (o *Match) HasBackend() bool`

HasBackend returns a boolean if a field has been set.

### GetRank

`func (o *Match) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *Match) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *Match) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *Match) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetScore

`func (o *Match) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Match) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Match) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *Match) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


