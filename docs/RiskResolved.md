# RiskResolved

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsOf** | Pointer to **string** | AsOf is the instant this answer was true at: the event time plus the horizon. Nothing seen after it was visible to this resolution. | [optional] 
**At** | Pointer to **string** | At is the event&#39;s instant, RFC 3339, echoed. It is what the horizon is measured from, so At plus the horizon is AsOf. | [optional] 
**By** | Pointer to **string** | By is the identity that filed the WINNING assertion, &#x60;&lt;home org&gt;/&lt;user&gt;&#x60;, stamped server-side from the validated principal at the write and never taken from a body — an attribution the caller chose is not attribution. It is the winner&#39;s alone; every losing assertion keeps its own and is returned whole in Conflicts. | [optional] 
**Confidence** | Pointer to **float64** | Confidence is the winning assertion&#39;s own confidence in [0,1], zero when its filer stated none. It is reported because it is a term of the rule that picked the winner, and it is the weakest term but one: it breaks a tie inside one rank and never lifts a weak source above a strong one. | [optional] 
**Conflicts** | Pointer to [**[]RiskLabelRecord**](RiskLabelRecord.md) | Conflicts is every other visible assertion, strongest first, whole. They are kept and returned rather than dropped, so an adverse action can show that the plane knew of a contrary claim and say why it lost. They are horizon-filtered exactly like the winner: an assertion that was not knowable yet cannot even be named here, because naming it would leak its existence into a past decision. | [optional] 
**Contested** | Pointer to **bool** | Contested is true when a visible assertion claimed a DIFFERENT disposition. Two sources agreeing is corroboration, not conflict. | [optional] 
**Disposition** | Pointer to **string** | Disposition is the claim IN FORCE at AsOf: productive, unproductive, or the empty string for an explicit unjudged. It is the winning assertion&#39;s own claim, never a vote or an average — an average of two adjudications is a third claim nobody made. A matured event nobody judged is not answered here at all; it is counted in Unlabelled, because manufacturing a negative there is how a fraud model comes to describe the incumbent block list. | [optional] 
**Evidence** | Pointer to **string** | Evidence is the winning assertion&#39;s pointer to the record behind it — the dispute, case or decision id it was filed with, opaque and verbatim. It travels with the answer so an adverse action can name what judged the subject without a second read. | [optional] 
**Id** | Pointer to **string** | ID is the winning assertion&#39;s content digest, so this answer traces to the exact record it came from — and that record can be placed under litigation hold by naming this id. | [optional] 
**Kind** | Pointer to **string** | Kind is the judged entity&#39;s type, echoed from the event that was named. With Subject and At it is how a caller joins this answer back onto the training row or the decision it asked about. | [optional] 
**Source** | Pointer to **string** | Source is who filed the winning assertion, and it is the PRIMARY term of the rule that picked it. Sources rank by adjudication weight — chargeoff, dispute, case, refund, review, sample, strongest first — and only inside one rank do the tie-breaks run, in order: the assertion that became KNOWABLE latest, then the higher confidence, then the lower id. The vocabulary op publishes that order from the same declaration the resolver reads, so a caller holding a contested answer can reproduce it. | [optional] 
**Subject** | Pointer to **string** | Subject is the entity id, echoed from the event that was named — the tenant&#39;s own key, returned verbatim. | [optional] 

## Methods

### NewRiskResolved

`func NewRiskResolved() *RiskResolved`

NewRiskResolved instantiates a new RiskResolved object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskResolvedWithDefaults

`func NewRiskResolvedWithDefaults() *RiskResolved`

NewRiskResolvedWithDefaults instantiates a new RiskResolved object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsOf

`func (o *RiskResolved) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *RiskResolved) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *RiskResolved) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *RiskResolved) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetAt

`func (o *RiskResolved) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *RiskResolved) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *RiskResolved) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *RiskResolved) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBy

`func (o *RiskResolved) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *RiskResolved) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *RiskResolved) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *RiskResolved) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetConfidence

`func (o *RiskResolved) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *RiskResolved) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *RiskResolved) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *RiskResolved) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetConflicts

`func (o *RiskResolved) GetConflicts() []RiskLabelRecord`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *RiskResolved) GetConflictsOk() (*[]RiskLabelRecord, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *RiskResolved) SetConflicts(v []RiskLabelRecord)`

SetConflicts sets Conflicts field to given value.

### HasConflicts

`func (o *RiskResolved) HasConflicts() bool`

HasConflicts returns a boolean if a field has been set.

### GetContested

`func (o *RiskResolved) GetContested() bool`

GetContested returns the Contested field if non-nil, zero value otherwise.

### GetContestedOk

`func (o *RiskResolved) GetContestedOk() (*bool, bool)`

GetContestedOk returns a tuple with the Contested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContested

`func (o *RiskResolved) SetContested(v bool)`

SetContested sets Contested field to given value.

### HasContested

`func (o *RiskResolved) HasContested() bool`

HasContested returns a boolean if a field has been set.

### GetDisposition

`func (o *RiskResolved) GetDisposition() string`

GetDisposition returns the Disposition field if non-nil, zero value otherwise.

### GetDispositionOk

`func (o *RiskResolved) GetDispositionOk() (*string, bool)`

GetDispositionOk returns a tuple with the Disposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposition

`func (o *RiskResolved) SetDisposition(v string)`

SetDisposition sets Disposition field to given value.

### HasDisposition

`func (o *RiskResolved) HasDisposition() bool`

HasDisposition returns a boolean if a field has been set.

### GetEvidence

`func (o *RiskResolved) GetEvidence() string`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *RiskResolved) GetEvidenceOk() (*string, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *RiskResolved) SetEvidence(v string)`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *RiskResolved) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.

### GetId

`func (o *RiskResolved) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskResolved) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskResolved) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskResolved) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *RiskResolved) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RiskResolved) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RiskResolved) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RiskResolved) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSource

`func (o *RiskResolved) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RiskResolved) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RiskResolved) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RiskResolved) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSubject

`func (o *RiskResolved) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RiskResolved) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RiskResolved) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *RiskResolved) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


