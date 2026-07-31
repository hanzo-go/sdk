# CloudAccView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basis** | Pointer to **string** | Basis is the qualification category: income, net_worth, professional_license, or entity. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is the unix second the record was created. | [optional] 
**EvidenceDocId** | Pointer to **string** | EvidenceDocID references an evidence document in the org&#39;s sealed data room. | [optional] 
**ExpiresAt** | Pointer to **int32** | ExpiresAt is the unix second a confirmation ages out; 0 means none. | [optional] 
**Id** | Pointer to **string** | ID is the accreditation record&#39;s opaque id. | [optional] 
**Method** | Pointer to **string** | Method is how the state was established: self_attested, third_party_letter, or provider_verified. | [optional] 
**Note** | Pointer to **string** | Note is a non-PII operator note. | [optional] 
**ReviewerSub** | Pointer to **string** | ReviewerSub is the org user who recorded a decision on this record. | [optional] 
**Status** | Pointer to **string** | Status is the tracked state: asserted, provider_verified, reviewer_confirmed, rejected, or expired. | [optional] 
**SubjectId** | Pointer to **string** | SubjectID is the opaque id of the subject the record is about. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is the unix second the record last changed. | [optional] 

## Methods

### NewCloudAccView

`func NewCloudAccView() *CloudAccView`

NewCloudAccView instantiates a new CloudAccView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAccViewWithDefaults

`func NewCloudAccViewWithDefaults() *CloudAccView`

NewCloudAccViewWithDefaults instantiates a new CloudAccView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasis

`func (o *CloudAccView) GetBasis() string`

GetBasis returns the Basis field if non-nil, zero value otherwise.

### GetBasisOk

`func (o *CloudAccView) GetBasisOk() (*string, bool)`

GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasis

`func (o *CloudAccView) SetBasis(v string)`

SetBasis sets Basis field to given value.

### HasBasis

`func (o *CloudAccView) HasBasis() bool`

HasBasis returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudAccView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudAccView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudAccView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudAccView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEvidenceDocId

`func (o *CloudAccView) GetEvidenceDocId() string`

GetEvidenceDocId returns the EvidenceDocId field if non-nil, zero value otherwise.

### GetEvidenceDocIdOk

`func (o *CloudAccView) GetEvidenceDocIdOk() (*string, bool)`

GetEvidenceDocIdOk returns a tuple with the EvidenceDocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceDocId

`func (o *CloudAccView) SetEvidenceDocId(v string)`

SetEvidenceDocId sets EvidenceDocId field to given value.

### HasEvidenceDocId

`func (o *CloudAccView) HasEvidenceDocId() bool`

HasEvidenceDocId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CloudAccView) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CloudAccView) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CloudAccView) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CloudAccView) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *CloudAccView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudAccView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudAccView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudAccView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethod

`func (o *CloudAccView) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CloudAccView) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CloudAccView) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *CloudAccView) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetNote

`func (o *CloudAccView) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CloudAccView) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CloudAccView) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CloudAccView) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetReviewerSub

`func (o *CloudAccView) GetReviewerSub() string`

GetReviewerSub returns the ReviewerSub field if non-nil, zero value otherwise.

### GetReviewerSubOk

`func (o *CloudAccView) GetReviewerSubOk() (*string, bool)`

GetReviewerSubOk returns a tuple with the ReviewerSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerSub

`func (o *CloudAccView) SetReviewerSub(v string)`

SetReviewerSub sets ReviewerSub field to given value.

### HasReviewerSub

`func (o *CloudAccView) HasReviewerSub() bool`

HasReviewerSub returns a boolean if a field has been set.

### GetStatus

`func (o *CloudAccView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAccView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAccView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAccView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubjectId

`func (o *CloudAccView) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *CloudAccView) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *CloudAccView) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *CloudAccView) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudAccView) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudAccView) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudAccView) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudAccView) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


