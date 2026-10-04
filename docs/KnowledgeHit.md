# KnowledgeHit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Doctype** | Pointer to **string** | DocType is which kind of knowledge matched, by address: kb.page (a wiki page), kb.memory (a unit of agent memory) or kb.source (a document a connector ingested). Those three are the whole indexed set, and searchIn.DocTypes filters on them. | [optional] 
**Name** | Pointer to **string** | Name is the document&#39;s name in the framework store — the id to read or open it with. Unique per (org, doctype), so it identifies the document with DocType and not alone. | [optional] 
**Project** | Pointer to **string** | Project is the project scope the document was saved under. Absent for a document saved with none, which is also why a project-scoped query cannot reach it. | [optional] 
**Provider** | Pointer to **string** | Provider is the connector that ingested the document — github, slack, google or notion. Absent for a page or memory written in the product, which came from no connector. | [optional] 
**Score** | Pointer to **float64** | Score is the cosine similarity between the query&#39;s embedding and the document&#39;s BEST-matching passage, from -1 to 1, higher being closer. A document is cut into passages of about 2000 bytes, each embedded on its own, so a fact deep in a long page scores as high as one in its first paragraph. 0 for a document only the lexical leg found, which has no similarity to report. There is no absolute cutoff: what counts as a good score moves with the query and the embedding model, so compare scores within one response and not across queries. | [optional] 
**Title** | Pointer to **string** | Title is the document&#39;s title. Empty for a document saved without one; it is a label to show, never the id (that is Name). | [optional] 
**Url** | Pointer to **string** | URL is the document&#39;s link back into the app it was ingested from. Absent when the document carries none, which is the normal case for pages and memories. | [optional] 

## Methods

### NewKnowledgeHit

`func NewKnowledgeHit() *KnowledgeHit`

NewKnowledgeHit instantiates a new KnowledgeHit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnowledgeHitWithDefaults

`func NewKnowledgeHitWithDefaults() *KnowledgeHit`

NewKnowledgeHitWithDefaults instantiates a new KnowledgeHit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoctype

`func (o *KnowledgeHit) GetDoctype() string`

GetDoctype returns the Doctype field if non-nil, zero value otherwise.

### GetDoctypeOk

`func (o *KnowledgeHit) GetDoctypeOk() (*string, bool)`

GetDoctypeOk returns a tuple with the Doctype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoctype

`func (o *KnowledgeHit) SetDoctype(v string)`

SetDoctype sets Doctype field to given value.

### HasDoctype

`func (o *KnowledgeHit) HasDoctype() bool`

HasDoctype returns a boolean if a field has been set.

### GetName

`func (o *KnowledgeHit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KnowledgeHit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KnowledgeHit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KnowledgeHit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *KnowledgeHit) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *KnowledgeHit) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *KnowledgeHit) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *KnowledgeHit) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProvider

`func (o *KnowledgeHit) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *KnowledgeHit) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *KnowledgeHit) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *KnowledgeHit) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetScore

`func (o *KnowledgeHit) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *KnowledgeHit) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *KnowledgeHit) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *KnowledgeHit) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetTitle

`func (o *KnowledgeHit) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *KnowledgeHit) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *KnowledgeHit) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *KnowledgeHit) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *KnowledgeHit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *KnowledgeHit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *KnowledgeHit) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *KnowledgeHit) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


