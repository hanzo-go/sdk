# RiskLabelFact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is when the judged event happened, RFC 3339. | [optional] 
**Confidence** | Pointer to **float32** | Confidence in [0,1]. A processor chargeback is 1; an analyst&#39;s hunch is not. It breaks a tie WITHIN a precedence rank and can never lift a weak source above a strong one — otherwise every caller would send 1.  A litigation hold is NOT a field here. It is a fact about the record and not about the world, so it is not part of what was asserted, it is not in the content digest, and it has its own op — which is also the only way one can be released. Carried here it was silently a no-op on any record that already existed: the digest was the same, the insert was ignored, and the caller was told &#x60;duplicate&#x60; while the hold it asked for was never placed. | [optional] 
**Disposition** | Pointer to **string** | Disposition is productive, unproductive, or empty for an explicit unjudged — the AML engine&#39;s own vocabulary, verbatim. | [optional] 
**Evidence** | Pointer to **string** | Evidence points at the record this conclusion came from: a dispute id, a case id, a decision id. Required, because a label with no evidence cannot be defended when the adverse action it fed is challenged. | [optional] 
**Kind** | Pointer to **string** | Kind is what the subject is: account, agent, merchant, payout, person, session or transaction. Closed, because a typo in an open field would shard a tenant&#39;s labels into a partition nothing reads and nothing would say so. | [optional] 
**Seen** | Pointer to **string** | Seen is when this assertion became KNOWABLE, RFC 3339. It is required and it is not At: a chargeback lands 30 to 120 days after the transaction it judges, and a training set joined on At alone knows the future. Everything this plane does to prevent leakage is computed from Seen. | [optional] 
**Source** | Pointer to **string** | Source is who asserted: chargeoff, dispute, case, refund, review or sample. It is the primary term of the precedence rule, so it is closed — an unknown source has no rank and a conflict with it could not be resolved. | [optional] 
**Subject** | Pointer to **string** | Subject identifies the thing being judged, in the tenant&#39;s own namespace. | [optional] 

## Methods

### NewRiskLabelFact

`func NewRiskLabelFact() *RiskLabelFact`

NewRiskLabelFact instantiates a new RiskLabelFact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskLabelFactWithDefaults

`func NewRiskLabelFactWithDefaults() *RiskLabelFact`

NewRiskLabelFactWithDefaults instantiates a new RiskLabelFact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *RiskLabelFact) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *RiskLabelFact) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *RiskLabelFact) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *RiskLabelFact) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetConfidence

`func (o *RiskLabelFact) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *RiskLabelFact) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *RiskLabelFact) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *RiskLabelFact) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetDisposition

`func (o *RiskLabelFact) GetDisposition() string`

GetDisposition returns the Disposition field if non-nil, zero value otherwise.

### GetDispositionOk

`func (o *RiskLabelFact) GetDispositionOk() (*string, bool)`

GetDispositionOk returns a tuple with the Disposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposition

`func (o *RiskLabelFact) SetDisposition(v string)`

SetDisposition sets Disposition field to given value.

### HasDisposition

`func (o *RiskLabelFact) HasDisposition() bool`

HasDisposition returns a boolean if a field has been set.

### GetEvidence

`func (o *RiskLabelFact) GetEvidence() string`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *RiskLabelFact) GetEvidenceOk() (*string, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *RiskLabelFact) SetEvidence(v string)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *RiskLabelFact) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.

### GetKind

`func (o *RiskLabelFact) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RiskLabelFact) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RiskLabelFact) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RiskLabelFact) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSeen

`func (o *RiskLabelFact) GetSeen() string`

GetSeen returns the Seen field if non-nil, zero value otherwise.

### GetSeenOk

`func (o *RiskLabelFact) GetSeenOk() (*string, bool)`

GetSeenOk returns a tuple with the Seen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeen

`func (o *RiskLabelFact) SetSeen(v string)`

SetSeen sets Seen field to given value.

### HasSeen

`func (o *RiskLabelFact) HasSeen() bool`

HasSeen returns a boolean if a field has been set.

### GetSource

`func (o *RiskLabelFact) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RiskLabelFact) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RiskLabelFact) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RiskLabelFact) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSubject

`func (o *RiskLabelFact) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RiskLabelFact) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RiskLabelFact) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *RiskLabelFact) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


