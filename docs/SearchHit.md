# SearchHit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Corpus** | Pointer to **string** | Corpus is which store the document lives in: \&quot;kb\&quot; for a document either knowledge leg returned, \&quot;code\&quot; for a span out of one of the org&#39;s own repositories, \&quot;files\&quot; for a passage of one of the org&#39;s workspace files (ID is then the file&#39;s id and Title its citation). It is PROVENANCE — read it to say where a hit came from, not to branch on: the fused ranking is what decides order, and a caller that filters by corpus wants the backend&#39;s own endpoint instead. | [optional] 
**Doctype** | Pointer to **string** | DocType is the knowledge doctype: kb.page, kb.memory or kb.source from the semantic leg, and a lexical row&#39;s own doctype/type field otherwise. Absent when the row carried neither. | [optional] 
**Id** | Pointer to **string** | ID is the document&#39;s identity inside its corpus — the KB document name from the semantic leg, or a lexical row&#39;s own name/id/_id (falling back to \&quot;row-&lt;n&gt;\&quot; when the row carries none). It is unique with DocType, not alone: the pair is the key the two legs are fused on. | [optional] 
**Matched** | Pointer to [**[]SearchProvenance**](SearchProvenance.md) | Matched is one entry per leg that returned this document, with that leg&#39;s rank and native score. More than one entry means the legs AGREED, which is exactly why the hit outranks one a single leg found. Never empty on a returned hit. | [optional] 
**Passage** | Pointer to **string** | Passage is the matching passage, for a hit from the files corpus: the span of the file&#39;s section that matched, about 2000 bytes. Absent for every other corpus, whose hits name a document rather than a span of one. | [optional] 
**Project** | Pointer to **string** | Project is the project scope the document was indexed under. Absent for a document saved with none; Request.Project filters the semantic leg on it. | [optional] 
**Score** | Pointer to **float64** | Score is the FUSED score, not a relevance or a similarity: Reciprocal Rank Fusion sums 1/(60+rank) over each leg that returned the document, so it is bounded by roughly 1/61 per leg (about 0.033 for a document two legs put first) and hits are ordered by it, descending. Being built from ranks, it is comparable only WITHIN one response — never across queries, and never against a backend&#39;s own score, which stays in Matched. | [optional] 
**Title** | Pointer to **string** | Title is the document&#39;s display title — the indexed title from the semantic leg, the row&#39;s title (falling back to its name) from the lexical one. Absent when the document has none. | [optional] 
**Url** | Pointer to **string** | URL is where the document can be opened, carried from the indexed payload — the link back to the app a connector ingested it from. Absent for anything written in the product, which has no external address. | [optional] 

## Methods

### NewSearchHit

`func NewSearchHit() *SearchHit`

NewSearchHit instantiates a new SearchHit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchHitWithDefaults

`func NewSearchHitWithDefaults() *SearchHit`

NewSearchHitWithDefaults instantiates a new SearchHit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorpus

`func (o *SearchHit) GetCorpus() string`

GetCorpus returns the Corpus field if non-nil, zero value otherwise.

### GetCorpusOk

`func (o *SearchHit) GetCorpusOk() (*string, bool)`

GetCorpusOk returns a tuple with the Corpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorpus

`func (o *SearchHit) SetCorpus(v string)`

SetCorpus sets Corpus field to given value.

### HasCorpus

`func (o *SearchHit) HasCorpus() bool`

HasCorpus returns a boolean if a field has been set.

### GetDoctype

`func (o *SearchHit) GetDoctype() string`

GetDoctype returns the Doctype field if non-nil, zero value otherwise.

### GetDoctypeOk

`func (o *SearchHit) GetDoctypeOk() (*string, bool)`

GetDoctypeOk returns a tuple with the Doctype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoctype

`func (o *SearchHit) SetDoctype(v string)`

SetDoctype sets Doctype field to given value.

### HasDoctype

`func (o *SearchHit) HasDoctype() bool`

HasDoctype returns a boolean if a field has been set.

### GetId

`func (o *SearchHit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchHit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchHit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SearchHit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMatched

`func (o *SearchHit) GetMatched() []SearchProvenance`

GetMatched returns the Matched field if non-nil, zero value otherwise.

### GetMatchedOk

`func (o *SearchHit) GetMatchedOk() (*[]SearchProvenance, bool)`

GetMatchedOk returns a tuple with the Matched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatched

`func (o *SearchHit) SetMatched(v []SearchProvenance)`

SetMatched sets Matched field to given value.

### HasMatched

`func (o *SearchHit) HasMatched() bool`

HasMatched returns a boolean if a field has been set.

### GetPassage

`func (o *SearchHit) GetPassage() string`

GetPassage returns the Passage field if non-nil, zero value otherwise.

### GetPassageOk

`func (o *SearchHit) GetPassageOk() (*string, bool)`

GetPassageOk returns a tuple with the Passage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassage

`func (o *SearchHit) SetPassage(v string)`

SetPassage sets Passage field to given value.

### HasPassage

`func (o *SearchHit) HasPassage() bool`

HasPassage returns a boolean if a field has been set.

### GetProject

`func (o *SearchHit) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SearchHit) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SearchHit) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *SearchHit) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetScore

`func (o *SearchHit) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SearchHit) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SearchHit) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *SearchHit) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetTitle

`func (o *SearchHit) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SearchHit) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SearchHit) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SearchHit) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *SearchHit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SearchHit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SearchHit) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SearchHit) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


