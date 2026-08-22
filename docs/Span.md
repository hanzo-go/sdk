# Span

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndLine** | Pointer to **int32** | EndLine is the last line of the span, inclusive. It equals Line for a one-line span rather than being zero or absent. | [optional] 
**File** | Pointer to **string** | File is the path inside the repo, relative to its root and never absolute. | [optional] 
**Kind** | Pointer to **string** | Kind is what the indexer decided this chunk IS — \&quot;func\&quot;, \&quot;method\&quot;, \&quot;type\&quot;, \&quot;struct\&quot;, \&quot;interface\&quot;, \&quot;var\&quot;, \&quot;const\&quot;, or \&quot;block\&quot; for a run of code that declares nothing. Absent when the chunker could not classify it. | [optional] 
**Line** | Pointer to **int32** | Line is where the span starts, 1-based, as an editor counts. | [optional] 
**Repo** | Pointer to **string** | Repo is the indexed repository the span was found in, as it was indexed (\&quot;owner/name\&quot;). A search may be scoped to one repo or run across all of them, so this is how a caller tells the results apart. | [optional] 
**Role** | Pointer to **string** | context: match | definition | caller | [optional] 
**Score** | Pointer to **float32** | Score ranks this span against the OTHERS IN THE SAME RESPONSE and means nothing across responses or between tiers: the hybrid tier&#39;s number is a reciprocal-rank fusion sum (Σ 1/(60+rank), so tenths at best), the symbol tier&#39;s is a descending position count, and the text and semantic tiers pass through bm25 and cosine. Compare within a list; never threshold on it. | [optional] 
**Snippet** | Pointer to **string** | Snippet is the code itself: a bounded excerpt on /search, the whole chunk on /context — which is why the same type serves both and why a /context span is the one an agent pastes into its window. | [optional] 
**Symbol** | Pointer to **string** | Symbol is the declared name, when the span declares one. Absent on a block. | [optional] 
**Tier** | Pointer to **string** | Tier is which retrieval produced the span: \&quot;hybrid\&quot; (the default — all three fused), \&quot;text\&quot; (trigram/FTS), \&quot;regex\&quot;, \&quot;semantic\&quot; (vector), or \&quot;symbol\&quot;. It is what explains a Score, so the two travel together. | [optional] 

## Methods

### NewSpan

`func NewSpan() *Span`

NewSpan instantiates a new Span object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpanWithDefaults

`func NewSpanWithDefaults() *Span`

NewSpanWithDefaults instantiates a new Span object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndLine

`func (o *Span) GetEndLine() int32`

GetEndLine returns the EndLine field if non-nil, zero value otherwise.

### GetEndLineOk

`func (o *Span) GetEndLineOk() (*int32, bool)`

GetEndLineOk returns a tuple with the EndLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLine

`func (o *Span) SetEndLine(v int32)`

SetEndLine sets EndLine field to given value.

### HasEndLine

`func (o *Span) HasEndLine() bool`

HasEndLine returns a boolean if a field has been set.

### GetFile

`func (o *Span) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *Span) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *Span) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *Span) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetKind

`func (o *Span) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Span) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Span) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Span) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLine

`func (o *Span) GetLine() int32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *Span) GetLineOk() (*int32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *Span) SetLine(v int32)`

SetLine sets Line field to given value.

### HasLine

`func (o *Span) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetRepo

`func (o *Span) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *Span) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *Span) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *Span) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRole

`func (o *Span) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Span) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Span) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Span) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetScore

`func (o *Span) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Span) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Span) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *Span) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSnippet

`func (o *Span) GetSnippet() string`

GetSnippet returns the Snippet field if non-nil, zero value otherwise.

### GetSnippetOk

`func (o *Span) GetSnippetOk() (*string, bool)`

GetSnippetOk returns a tuple with the Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippet

`func (o *Span) SetSnippet(v string)`

SetSnippet sets Snippet field to given value.

### HasSnippet

`func (o *Span) HasSnippet() bool`

HasSnippet returns a boolean if a field has been set.

### GetSymbol

`func (o *Span) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Span) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Span) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Span) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTier

`func (o *Span) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *Span) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *Span) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *Span) HasTier() bool`

HasTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


