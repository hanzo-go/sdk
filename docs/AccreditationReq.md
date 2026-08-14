# AccreditationReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basis** | Pointer to **string** | Basis is the qualification category: income, net_worth, professional_license, or entity. | [optional] 
**EvidenceDocId** | Pointer to **string** | EvidenceDocID references an evidence document in the org&#39;s sealed data room. | [optional] 
**ExpiresAt** | Pointer to **int32** | ExpiresAt is the unix second a confirmation ages out; 0 means none. | [optional] 
**Method** | Pointer to **string** | Method is how the state was established: self_attested, third_party_letter, or provider_verified. | [optional] 
**Note** | Pointer to **string** | Note is a non-PII operator note. | [optional] 
**Status** | Pointer to **string** | Status may only be \&quot;asserted\&quot; (empty reads as asserted); every confirmed, rejected or expired state is recorded via the decision endpoint. | [optional] 
**SubjectId** | Pointer to **string** | SubjectID names the subject this record is about; it must exist within the org. | [optional] 

## Methods

### NewAccreditationReq

`func NewAccreditationReq() *AccreditationReq`

NewAccreditationReq instantiates a new AccreditationReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccreditationReqWithDefaults

`func NewAccreditationReqWithDefaults() *AccreditationReq`

NewAccreditationReqWithDefaults instantiates a new AccreditationReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasis

`func (o *AccreditationReq) GetBasis() string`

GetBasis returns the Basis field if non-nil, zero value otherwise.

### GetBasisOk

`func (o *AccreditationReq) GetBasisOk() (*string, bool)`

GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasis

`func (o *AccreditationReq) SetBasis(v string)`

SetBasis sets Basis field to given value.

### HasBasis

`func (o *AccreditationReq) HasBasis() bool`

HasBasis returns a boolean if a field has been set.

### GetEvidenceDocId

`func (o *AccreditationReq) GetEvidenceDocId() string`

GetEvidenceDocId returns the EvidenceDocId field if non-nil, zero value otherwise.

### GetEvidenceDocIdOk

`func (o *AccreditationReq) GetEvidenceDocIdOk() (*string, bool)`

GetEvidenceDocIdOk returns a tuple with the EvidenceDocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceDocId

`func (o *AccreditationReq) SetEvidenceDocId(v string)`

SetEvidenceDocId sets EvidenceDocId field to given value.

### HasEvidenceDocId

`func (o *AccreditationReq) HasEvidenceDocId() bool`

HasEvidenceDocId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AccreditationReq) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AccreditationReq) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AccreditationReq) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AccreditationReq) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetMethod

`func (o *AccreditationReq) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AccreditationReq) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AccreditationReq) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *AccreditationReq) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetNote

`func (o *AccreditationReq) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *AccreditationReq) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *AccreditationReq) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *AccreditationReq) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetStatus

`func (o *AccreditationReq) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccreditationReq) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccreditationReq) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccreditationReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubjectId

`func (o *AccreditationReq) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *AccreditationReq) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *AccreditationReq) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *AccreditationReq) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


