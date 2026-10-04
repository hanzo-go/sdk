# ComplianceAccView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basis** | Pointer to **string** | Basis is the qualification category: income, net_worth, professional_license, or entity. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is the unix second the record was created. | [optional] 
**EvidenceDocId** | Pointer to **string** | EvidenceDocID references an evidence document in the org&#39;s sealed data room. | [optional] 
**ExpiresAt** | Pointer to **int64** | ExpiresAt is the unix second a confirmation ages out; 0 means none. | [optional] 
**Id** | Pointer to **string** | ID is the accreditation record&#39;s opaque id. | [optional] 
**Method** | Pointer to **string** | Method is how the state was established: self_attested, third_party_letter, or provider_verified. | [optional] 
**Note** | Pointer to **string** | Note is a non-PII operator note. | [optional] 
**ReviewerSub** | Pointer to **string** | ReviewerSub is the org user who recorded a decision on this record. | [optional] 
**Status** | Pointer to **string** | Status is the tracked state: asserted, provider_verified, reviewer_confirmed, rejected, or expired. | [optional] 
**SubjectId** | Pointer to **string** | SubjectID is the opaque id of the subject the record is about. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is the unix second the record last changed. | [optional] 

## Methods

### NewComplianceAccView

`func NewComplianceAccView() *ComplianceAccView`

NewComplianceAccView instantiates a new ComplianceAccView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceAccViewWithDefaults

`func NewComplianceAccViewWithDefaults() *ComplianceAccView`

NewComplianceAccViewWithDefaults instantiates a new ComplianceAccView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasis

`func (o *ComplianceAccView) GetBasis() string`

GetBasis returns the Basis field if non-nil, zero value otherwise.

### GetBasisOk

`func (o *ComplianceAccView) GetBasisOk() (*string, bool)`

GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasis

`func (o *ComplianceAccView) SetBasis(v string)`

SetBasis sets Basis field to given value.

### HasBasis

`func (o *ComplianceAccView) HasBasis() bool`

HasBasis returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ComplianceAccView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ComplianceAccView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ComplianceAccView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ComplianceAccView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEvidenceDocId

`func (o *ComplianceAccView) GetEvidenceDocId() string`

GetEvidenceDocId returns the EvidenceDocId field if non-nil, zero value otherwise.

### GetEvidenceDocIdOk

`func (o *ComplianceAccView) GetEvidenceDocIdOk() (*string, bool)`

GetEvidenceDocIdOk returns a tuple with the EvidenceDocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceDocId

`func (o *ComplianceAccView) SetEvidenceDocId(v string)`

SetEvidenceDocId sets EvidenceDocId field to given value.

### HasEvidenceDocId

`func (o *ComplianceAccView) HasEvidenceDocId() bool`

HasEvidenceDocId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ComplianceAccView) GetExpiresAt() int64`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ComplianceAccView) GetExpiresAtOk() (*int64, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ComplianceAccView) SetExpiresAt(v int64)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ComplianceAccView) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *ComplianceAccView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComplianceAccView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComplianceAccView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComplianceAccView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethod

`func (o *ComplianceAccView) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ComplianceAccView) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ComplianceAccView) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ComplianceAccView) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetNote

`func (o *ComplianceAccView) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ComplianceAccView) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ComplianceAccView) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ComplianceAccView) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetReviewerSub

`func (o *ComplianceAccView) GetReviewerSub() string`

GetReviewerSub returns the ReviewerSub field if non-nil, zero value otherwise.

### GetReviewerSubOk

`func (o *ComplianceAccView) GetReviewerSubOk() (*string, bool)`

GetReviewerSubOk returns a tuple with the ReviewerSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerSub

`func (o *ComplianceAccView) SetReviewerSub(v string)`

SetReviewerSub sets ReviewerSub field to given value.

### HasReviewerSub

`func (o *ComplianceAccView) HasReviewerSub() bool`

HasReviewerSub returns a boolean if a field has been set.

### GetStatus

`func (o *ComplianceAccView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComplianceAccView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComplianceAccView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ComplianceAccView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubjectId

`func (o *ComplianceAccView) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *ComplianceAccView) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *ComplianceAccView) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *ComplianceAccView) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ComplianceAccView) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ComplianceAccView) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ComplianceAccView) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ComplianceAccView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


