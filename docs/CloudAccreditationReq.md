# CloudAccreditationReq

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

### NewCloudAccreditationReq

`func NewCloudAccreditationReq() *CloudAccreditationReq`

NewCloudAccreditationReq instantiates a new CloudAccreditationReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAccreditationReqWithDefaults

`func NewCloudAccreditationReqWithDefaults() *CloudAccreditationReq`

NewCloudAccreditationReqWithDefaults instantiates a new CloudAccreditationReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasis

`func (o *CloudAccreditationReq) GetBasis() string`

GetBasis returns the Basis field if non-nil, zero value otherwise.

### GetBasisOk

`func (o *CloudAccreditationReq) GetBasisOk() (*string, bool)`

GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasis

`func (o *CloudAccreditationReq) SetBasis(v string)`

SetBasis sets Basis field to given value.

### HasBasis

`func (o *CloudAccreditationReq) HasBasis() bool`

HasBasis returns a boolean if a field has been set.

### GetEvidenceDocId

`func (o *CloudAccreditationReq) GetEvidenceDocId() string`

GetEvidenceDocId returns the EvidenceDocId field if non-nil, zero value otherwise.

### GetEvidenceDocIdOk

`func (o *CloudAccreditationReq) GetEvidenceDocIdOk() (*string, bool)`

GetEvidenceDocIdOk returns a tuple with the EvidenceDocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceDocId

`func (o *CloudAccreditationReq) SetEvidenceDocId(v string)`

SetEvidenceDocId sets EvidenceDocId field to given value.

### HasEvidenceDocId

`func (o *CloudAccreditationReq) HasEvidenceDocId() bool`

HasEvidenceDocId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CloudAccreditationReq) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CloudAccreditationReq) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CloudAccreditationReq) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CloudAccreditationReq) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetMethod

`func (o *CloudAccreditationReq) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *CloudAccreditationReq) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *CloudAccreditationReq) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *CloudAccreditationReq) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetNote

`func (o *CloudAccreditationReq) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CloudAccreditationReq) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CloudAccreditationReq) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CloudAccreditationReq) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetStatus

`func (o *CloudAccreditationReq) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAccreditationReq) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAccreditationReq) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAccreditationReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubjectId

`func (o *CloudAccreditationReq) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *CloudAccreditationReq) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *CloudAccreditationReq) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *CloudAccreditationReq) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


