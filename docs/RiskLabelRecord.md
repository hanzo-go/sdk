# RiskLabelRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is when the judged EVENT happened, RFC 3339 in UTC, truncated to the second. The filer supplies it, and it is what a maturity horizon measures from: this event&#39;s as-of is At plus the horizon. A resolve names it back exactly, to the second. | [optional] 
**By** | Pointer to **string** | By is the identity that asserted, stamped server-side at the write. | [optional] 
**Confidence** | Pointer to **float32** | Confidence is the filer&#39;s own confidence in [0,1] — 1 for a processor chargeback, less for an analyst&#39;s hunch. Zero is the ordinary value for a filer that stated none, and it means the weakest tie-break there is rather than \&quot;unknown\&quot;. It breaks a tie only WITHIN one precedence rank and can never lift a weak source above a strong one. | [optional] 
**Disposition** | Pointer to **string** | Disposition is what was concluded, from the closed set: &#x60;productive&#x60; — the event led somewhere, escalated, reported or charged back; &#x60;unproductive&#x60; — judged not suspicious; or the empty string for an explicit UNJUDGED, which is a real assertion (\&quot;we looked and could not say\&quot;) and not the absence of one. | [optional] 
**Evidence** | Pointer to **string** | Evidence is the pointer to the record this conclusion came from: a dispute id, a case id, a decision id. At most 512 bytes, required at the write, and opaque to this plane — stored and returned verbatim, never resolved. It is what an adverse action is defended with, which is why an assertion carrying none is refused at the door. | [optional] 
**Hold** | Pointer to **bool** | Hold is true while a litigation hold is on this record: retention will not dispose of it, at any age. False — and it is omitted then — leaves the record disposable once it is older than the boundary a sweep names. It is a fact about the RECORD and not about the world, so it is not folded into ID, no write path can set it, and the hold op is the one way it moves in either direction. | [optional] 
**Id** | Pointer to **string** | ID is the assertion&#39;s content digest — SHA-256 over every semantic field, rendered hex — computed server-side and never supplied. It is the key a redelivery collapses onto, and it is the id the hold op names. | [optional] 
**Kind** | Pointer to **string** | Kind is what the subject IS, from the closed set: account, agent, merchant, payout, person, session or transaction. With Subject and At it is the IDENTITY of the judged event — the triple a resolve names and the triple assertions are grouped by, so a typo in it would file a label against an event nobody asks about. | [optional] 
**Knowable** | Pointer to **string** | Knowable is when THIS PLANE could first have answered with the assertion: the later of Seen and the server clock at the write, derived server-side. It is the instant the leakage guard compares, so it is published beside the claim it was derived from — an answer whose rule nobody can see is one nobody can check. | [optional] 
**Seen** | Pointer to **string** | Seen is when the FILER said the assertion became knowable. It is provenance: it is recorded and published, and it decides nothing. | [optional] 
**Source** | Pointer to **string** | Source is WHO asserted, from the closed set: chargeoff, dispute, case, refund, review or sample. It is the primary term of the precedence rule — an unknown source has no rank and a conflict with it could not be resolved — so it is what decides which of two disagreeing assertions is in force. | [optional] 
**Subject** | Pointer to **string** | Subject is the entity that was judged, named in the TENANT&#39;S OWN namespace and at most 512 bytes. It is opaque here: stored, matched and returned verbatim, never dereferenced. It has no meaning outside this tenant — the record is the tenant&#39;s own file — so an id lifted from another tenant&#39;s response names nothing. | [optional] 
**Wrote** | Pointer to **string** | Wrote is the server clock at the write. It is the only time on the record the tenant did not supply, and it is what retention measures against. | [optional] 

## Methods

### NewRiskLabelRecord

`func NewRiskLabelRecord() *RiskLabelRecord`

NewRiskLabelRecord instantiates a new RiskLabelRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskLabelRecordWithDefaults

`func NewRiskLabelRecordWithDefaults() *RiskLabelRecord`

NewRiskLabelRecordWithDefaults instantiates a new RiskLabelRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *RiskLabelRecord) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *RiskLabelRecord) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *RiskLabelRecord) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *RiskLabelRecord) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBy

`func (o *RiskLabelRecord) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *RiskLabelRecord) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *RiskLabelRecord) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *RiskLabelRecord) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetConfidence

`func (o *RiskLabelRecord) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *RiskLabelRecord) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *RiskLabelRecord) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *RiskLabelRecord) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetDisposition

`func (o *RiskLabelRecord) GetDisposition() string`

GetDisposition returns the Disposition field if non-nil, zero value otherwise.

### GetDispositionOk

`func (o *RiskLabelRecord) GetDispositionOk() (*string, bool)`

GetDispositionOk returns a tuple with the Disposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposition

`func (o *RiskLabelRecord) SetDisposition(v string)`

SetDisposition sets Disposition field to given value.

### HasDisposition

`func (o *RiskLabelRecord) HasDisposition() bool`

HasDisposition returns a boolean if a field has been set.

### GetEvidence

`func (o *RiskLabelRecord) GetEvidence() string`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *RiskLabelRecord) GetEvidenceOk() (*string, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *RiskLabelRecord) SetEvidence(v string)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *RiskLabelRecord) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.

### GetHold

`func (o *RiskLabelRecord) GetHold() bool`

GetHold returns the Hold field if non-nil, zero value otherwise.

### GetHoldOk

`func (o *RiskLabelRecord) GetHoldOk() (*bool, bool)`

GetHoldOk returns a tuple with the Hold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHold

`func (o *RiskLabelRecord) SetHold(v bool)`

SetHold sets Hold field to given value.

### HasHold

`func (o *RiskLabelRecord) HasHold() bool`

HasHold returns a boolean if a field has been set.

### GetId

`func (o *RiskLabelRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskLabelRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskLabelRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskLabelRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *RiskLabelRecord) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RiskLabelRecord) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RiskLabelRecord) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RiskLabelRecord) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetKnowable

`func (o *RiskLabelRecord) GetKnowable() string`

GetKnowable returns the Knowable field if non-nil, zero value otherwise.

### GetKnowableOk

`func (o *RiskLabelRecord) GetKnowableOk() (*string, bool)`

GetKnowableOk returns a tuple with the Knowable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowable

`func (o *RiskLabelRecord) SetKnowable(v string)`

SetKnowable sets Knowable field to given value.

### HasKnowable

`func (o *RiskLabelRecord) HasKnowable() bool`

HasKnowable returns a boolean if a field has been set.

### GetSeen

`func (o *RiskLabelRecord) GetSeen() string`

GetSeen returns the Seen field if non-nil, zero value otherwise.

### GetSeenOk

`func (o *RiskLabelRecord) GetSeenOk() (*string, bool)`

GetSeenOk returns a tuple with the Seen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeen

`func (o *RiskLabelRecord) SetSeen(v string)`

SetSeen sets Seen field to given value.

### HasSeen

`func (o *RiskLabelRecord) HasSeen() bool`

HasSeen returns a boolean if a field has been set.

### GetSource

`func (o *RiskLabelRecord) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RiskLabelRecord) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RiskLabelRecord) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RiskLabelRecord) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSubject

`func (o *RiskLabelRecord) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RiskLabelRecord) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RiskLabelRecord) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *RiskLabelRecord) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetWrote

`func (o *RiskLabelRecord) GetWrote() string`

GetWrote returns the Wrote field if non-nil, zero value otherwise.

### GetWroteOk

`func (o *RiskLabelRecord) GetWroteOk() (*string, bool)`

GetWroteOk returns a tuple with the Wrote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrote

`func (o *RiskLabelRecord) SetWrote(v string)`

SetWrote sets Wrote field to given value.

### HasWrote

`func (o *RiskLabelRecord) HasWrote() bool`

HasWrote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


