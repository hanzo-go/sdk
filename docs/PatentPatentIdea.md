# PatentPatentIdea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alternatives** | Pointer to [**[]PatentPatentRejectedAlternative**](PatentPatentRejectedAlternative.md) | Alternatives are the routes considered and rejected, each with its reason. | [optional] 
**AuthoredBy** | Pointer to **string** | AuthoredBy is the validated principal who recorded this version — an IAM subject, never a body field. It is the §115 answer, and a model has none. | [optional] 
**Canonical** | Pointer to **bool** | Canonical is DERIVED on read: a version is canonical when it is the latest-appended non-retracted version of its id. | [optional] 
**ContentHash** | Pointer to **string** | ContentHash is the SERVER&#39;s digest of the recorded content and this version&#39;s identity. Computed on a write; read back on a read. | [optional] 
**Id** | Pointer to **string** | ID is the entry&#39;s stable id, the key its versions accrue under: lowercase, terse, and named for the idea rather than the file — latency-aware-request-routing. Required on a write. | [optional] 
**Kind** | Pointer to **string** | Kind is what kind of conception this is: problem, alternative, architecture, mechanism, decision or rejection. Required on a write. | [optional] 
**Notes** | Pointer to **string** | Notes is anything a later reader needs that the fields above cannot carry. | [optional] 
**PriorArt** | Pointer to **string** | PriorArt is what the author knows already existed, named as the author knows it. It is a statement of the inventor&#39;s awareness, not a search result, and it is not a substitute for one. | [optional] 
**Problem** | Pointer to **string** | Problem is the technical problem being solved, and why the existing approaches fail at it. | [optional] 
**Project** | Pointer to **string** | Project is the sub-scope inside the org this entry is filed under. On a write it is the SERVER&#39;s value and a project in the body is ignored. | [optional] 
**RecordedAt** | Pointer to **string** | RecordedAt is when this store accepted the version, RFC3339, from the server clock. | [optional] 
**Revision** | Pointer to **string** | Revision is the producer&#39;s declared kind of this version: original, corrected or retracted. A retracted latest version WITHDRAWS the id; every earlier version stays retained. | [optional] 
**Seq** | Pointer to **int64** | Seq is this version&#39;s place in the store&#39;s append order. The ordering authority; never a client value. | [optional] 
**Signature** | Pointer to [**PatentPatentSignature**](PatentPatentSignature.md) | PatentSignature is an optional detached signature over this entry&#39;s digest. | [optional] 
**Statement** | Pointer to **string** | Statement is the operative idea in the author&#39;s own prose — definite enough that a reader could implement it from this alone. Required, and the field an examiner would actually read. | [optional] 
**Title** | Pointer to **string** | Title is the one-line name of the entry. | [optional] 
**Visibility** | Pointer to **string** | Visibility is private, org or public. A write forces private. | [optional] 

## Methods

### NewPatentPatentIdea

`func NewPatentPatentIdea() *PatentPatentIdea`

NewPatentPatentIdea instantiates a new PatentPatentIdea object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentPatentIdeaWithDefaults

`func NewPatentPatentIdeaWithDefaults() *PatentPatentIdea`

NewPatentPatentIdeaWithDefaults instantiates a new PatentPatentIdea object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternatives

`func (o *PatentPatentIdea) GetAlternatives() []PatentPatentRejectedAlternative`

GetAlternatives returns the Alternatives field if non-nil, zero value otherwise.

### GetAlternativesOk

`func (o *PatentPatentIdea) GetAlternativesOk() (*[]PatentPatentRejectedAlternative, bool)`

GetAlternativesOk returns a tuple with the Alternatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatives

`func (o *PatentPatentIdea) SetAlternatives(v []PatentPatentRejectedAlternative)`

SetAlternatives sets Alternatives field to given value.

### HasAlternatives

`func (o *PatentPatentIdea) HasAlternatives() bool`

HasAlternatives returns a boolean if a field has been set.

### GetAuthoredBy

`func (o *PatentPatentIdea) GetAuthoredBy() string`

GetAuthoredBy returns the AuthoredBy field if non-nil, zero value otherwise.

### GetAuthoredByOk

`func (o *PatentPatentIdea) GetAuthoredByOk() (*string, bool)`

GetAuthoredByOk returns a tuple with the AuthoredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoredBy

`func (o *PatentPatentIdea) SetAuthoredBy(v string)`

SetAuthoredBy sets AuthoredBy field to given value.

### HasAuthoredBy

`func (o *PatentPatentIdea) HasAuthoredBy() bool`

HasAuthoredBy returns a boolean if a field has been set.

### GetCanonical

`func (o *PatentPatentIdea) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *PatentPatentIdea) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *PatentPatentIdea) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.

### HasCanonical

`func (o *PatentPatentIdea) HasCanonical() bool`

HasCanonical returns a boolean if a field has been set.

### GetContentHash

`func (o *PatentPatentIdea) GetContentHash() string`

GetContentHash returns the ContentHash field if non-nil, zero value otherwise.

### GetContentHashOk

`func (o *PatentPatentIdea) GetContentHashOk() (*string, bool)`

GetContentHashOk returns a tuple with the ContentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHash

`func (o *PatentPatentIdea) SetContentHash(v string)`

SetContentHash sets ContentHash field to given value.

### HasContentHash

`func (o *PatentPatentIdea) HasContentHash() bool`

HasContentHash returns a boolean if a field has been set.

### GetId

`func (o *PatentPatentIdea) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatentPatentIdea) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatentPatentIdea) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatentPatentIdea) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *PatentPatentIdea) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PatentPatentIdea) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PatentPatentIdea) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PatentPatentIdea) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetNotes

`func (o *PatentPatentIdea) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PatentPatentIdea) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PatentPatentIdea) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PatentPatentIdea) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPriorArt

`func (o *PatentPatentIdea) GetPriorArt() string`

GetPriorArt returns the PriorArt field if non-nil, zero value otherwise.

### GetPriorArtOk

`func (o *PatentPatentIdea) GetPriorArtOk() (*string, bool)`

GetPriorArtOk returns a tuple with the PriorArt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorArt

`func (o *PatentPatentIdea) SetPriorArt(v string)`

SetPriorArt sets PriorArt field to given value.

### HasPriorArt

`func (o *PatentPatentIdea) HasPriorArt() bool`

HasPriorArt returns a boolean if a field has been set.

### GetProblem

`func (o *PatentPatentIdea) GetProblem() string`

GetProblem returns the Problem field if non-nil, zero value otherwise.

### GetProblemOk

`func (o *PatentPatentIdea) GetProblemOk() (*string, bool)`

GetProblemOk returns a tuple with the Problem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblem

`func (o *PatentPatentIdea) SetProblem(v string)`

SetProblem sets Problem field to given value.

### HasProblem

`func (o *PatentPatentIdea) HasProblem() bool`

HasProblem returns a boolean if a field has been set.

### GetProject

`func (o *PatentPatentIdea) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PatentPatentIdea) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PatentPatentIdea) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PatentPatentIdea) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRecordedAt

`func (o *PatentPatentIdea) GetRecordedAt() string`

GetRecordedAt returns the RecordedAt field if non-nil, zero value otherwise.

### GetRecordedAtOk

`func (o *PatentPatentIdea) GetRecordedAtOk() (*string, bool)`

GetRecordedAtOk returns a tuple with the RecordedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedAt

`func (o *PatentPatentIdea) SetRecordedAt(v string)`

SetRecordedAt sets RecordedAt field to given value.

### HasRecordedAt

`func (o *PatentPatentIdea) HasRecordedAt() bool`

HasRecordedAt returns a boolean if a field has been set.

### GetRevision

`func (o *PatentPatentIdea) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *PatentPatentIdea) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *PatentPatentIdea) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *PatentPatentIdea) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSeq

`func (o *PatentPatentIdea) GetSeq() int64`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *PatentPatentIdea) GetSeqOk() (*int64, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *PatentPatentIdea) SetSeq(v int64)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *PatentPatentIdea) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetSignature

`func (o *PatentPatentIdea) GetSignature() PatentPatentSignature`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *PatentPatentIdea) GetSignatureOk() (*PatentPatentSignature, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *PatentPatentIdea) SetSignature(v PatentPatentSignature)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *PatentPatentIdea) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetStatement

`func (o *PatentPatentIdea) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *PatentPatentIdea) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *PatentPatentIdea) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *PatentPatentIdea) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetTitle

`func (o *PatentPatentIdea) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PatentPatentIdea) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PatentPatentIdea) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PatentPatentIdea) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVisibility

`func (o *PatentPatentIdea) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *PatentPatentIdea) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *PatentPatentIdea) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *PatentPatentIdea) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


