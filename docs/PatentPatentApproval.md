# PatentPatentApproval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthoredBy** | Pointer to **string** | AuthoredBy is the validated principal who decided. The same rule as everywhere here: it comes from the credential, never the body, so one person cannot record a decision on another&#39;s behalf. | [optional] 
**Canonical** | Pointer to **bool** | Canonical is DERIVED on read: this author&#39;s latest word on this disclosure. | [optional] 
**Comment** | Pointer to **string** | Comment is the reviewer&#39;s note — the part an inventor reads. | [optional] 
**ContentHash** | Pointer to **string** | ContentHash is the server&#39;s digest of the recorded content. | [optional] 
**Decision** | Pointer to **string** | Decision is approved, rejected or needs-work. A needs-work is dated and retained exactly like a sign-off, because it IS one. | [optional] 
**DisclosureId** | Pointer to **string** | DisclosureID is the document this decides on. Required. | [optional] 
**Id** | Pointer to **string** | ID is the server&#39;s handle for this decision, derived from the disclosure and its author. Read-only: a caller cannot name someone else&#39;s approval. | [optional] 
**Project** | Pointer to **string** | Project is the SERVER&#39;s scope value. | [optional] 
**RecordedAt** | Pointer to **string** | RecordedAt is the server&#39;s clock. | [optional] 
**Revision** | Pointer to **string** | Revision is original, corrected or retracted. | [optional] 
**Role** | Pointer to **string** | Role is the capacity the person acted in: author, cto, reviewer, attorney or counsel. Required, because \&quot;approved\&quot; by nobody is not a fact. | [optional] 
**Seq** | Pointer to **int64** | Seq is the append order. | [optional] 

## Methods

### NewPatentPatentApproval

`func NewPatentPatentApproval() *PatentPatentApproval`

NewPatentPatentApproval instantiates a new PatentPatentApproval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentPatentApprovalWithDefaults

`func NewPatentPatentApprovalWithDefaults() *PatentPatentApproval`

NewPatentPatentApprovalWithDefaults instantiates a new PatentPatentApproval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthoredBy

`func (o *PatentPatentApproval) GetAuthoredBy() string`

GetAuthoredBy returns the AuthoredBy field if non-nil, zero value otherwise.

### GetAuthoredByOk

`func (o *PatentPatentApproval) GetAuthoredByOk() (*string, bool)`

GetAuthoredByOk returns a tuple with the AuthoredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoredBy

`func (o *PatentPatentApproval) SetAuthoredBy(v string)`

SetAuthoredBy sets AuthoredBy field to given value.

### HasAuthoredBy

`func (o *PatentPatentApproval) HasAuthoredBy() bool`

HasAuthoredBy returns a boolean if a field has been set.

### GetCanonical

`func (o *PatentPatentApproval) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *PatentPatentApproval) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *PatentPatentApproval) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.

### HasCanonical

`func (o *PatentPatentApproval) HasCanonical() bool`

HasCanonical returns a boolean if a field has been set.

### GetComment

`func (o *PatentPatentApproval) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PatentPatentApproval) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PatentPatentApproval) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PatentPatentApproval) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContentHash

`func (o *PatentPatentApproval) GetContentHash() string`

GetContentHash returns the ContentHash field if non-nil, zero value otherwise.

### GetContentHashOk

`func (o *PatentPatentApproval) GetContentHashOk() (*string, bool)`

GetContentHashOk returns a tuple with the ContentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHash

`func (o *PatentPatentApproval) SetContentHash(v string)`

SetContentHash sets ContentHash field to given value.

### HasContentHash

`func (o *PatentPatentApproval) HasContentHash() bool`

HasContentHash returns a boolean if a field has been set.

### GetDecision

`func (o *PatentPatentApproval) GetDecision() string`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *PatentPatentApproval) GetDecisionOk() (*string, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *PatentPatentApproval) SetDecision(v string)`

SetDecision sets Decision field to given value.

### HasDecision

`func (o *PatentPatentApproval) HasDecision() bool`

HasDecision returns a boolean if a field has been set.

### GetDisclosureId

`func (o *PatentPatentApproval) GetDisclosureId() string`

GetDisclosureId returns the DisclosureId field if non-nil, zero value otherwise.

### GetDisclosureIdOk

`func (o *PatentPatentApproval) GetDisclosureIdOk() (*string, bool)`

GetDisclosureIdOk returns a tuple with the DisclosureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosureId

`func (o *PatentPatentApproval) SetDisclosureId(v string)`

SetDisclosureId sets DisclosureId field to given value.

### HasDisclosureId

`func (o *PatentPatentApproval) HasDisclosureId() bool`

HasDisclosureId returns a boolean if a field has been set.

### GetId

`func (o *PatentPatentApproval) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatentPatentApproval) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatentPatentApproval) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatentPatentApproval) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProject

`func (o *PatentPatentApproval) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PatentPatentApproval) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PatentPatentApproval) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PatentPatentApproval) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRecordedAt

`func (o *PatentPatentApproval) GetRecordedAt() string`

GetRecordedAt returns the RecordedAt field if non-nil, zero value otherwise.

### GetRecordedAtOk

`func (o *PatentPatentApproval) GetRecordedAtOk() (*string, bool)`

GetRecordedAtOk returns a tuple with the RecordedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedAt

`func (o *PatentPatentApproval) SetRecordedAt(v string)`

SetRecordedAt sets RecordedAt field to given value.

### HasRecordedAt

`func (o *PatentPatentApproval) HasRecordedAt() bool`

HasRecordedAt returns a boolean if a field has been set.

### GetRevision

`func (o *PatentPatentApproval) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *PatentPatentApproval) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *PatentPatentApproval) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *PatentPatentApproval) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRole

`func (o *PatentPatentApproval) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatentPatentApproval) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatentPatentApproval) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatentPatentApproval) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSeq

`func (o *PatentPatentApproval) GetSeq() int64`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *PatentPatentApproval) GetSeqOk() (*int64, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *PatentPatentApproval) SetSeq(v int64)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *PatentPatentApproval) HasSeq() bool`

HasSeq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


