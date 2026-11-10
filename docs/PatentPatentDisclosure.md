# PatentPatentDisclosure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approvals** | Pointer to [**[]PatentPatentApproval**](PatentPatentApproval.md) | Approvals is DERIVED on read: the current decision of record from each person who has ruled on this disclosure. Never an input — a caller cannot approve its own document by writing a field in it. | [optional] 
**AuthoredBy** | Pointer to **string** | AuthoredBy is the validated principal who recorded this version. | [optional] 
**Body** | Pointer to **string** | Body is the document&#39;s text. Its digest is published beside it so a reader can tell a stored body from a re-sent one. | [optional] 
**BodySha256** | Pointer to **string** | BodySHA256 is the SERVER&#39;s digest of body. On a read it is the address of what is stored; a body is never hashed from what a reader sent. | [optional] 
**Canonical** | Pointer to **bool** | Canonical is DERIVED on read. | [optional] 
**Claims** | Pointer to **string** | Claims is the draft claim set, or a pointer to where it lives. | [optional] 
**ContentHash** | Pointer to **string** | ContentHash is the server&#39;s digest of the recorded content. | [optional] 
**FilingRef** | Pointer to **string** | FilingRef is a filing number a human recorded here after filing. Empty means nothing has been filed, and no code path fills it. | [optional] 
**Id** | Pointer to **string** | ID is the disclosure&#39;s stable id. Required. | [optional] 
**IdeaIds** | Pointer to **[]string** | IdeaIDs are the conception entries this disclosure is built from. The evidence half; a disclosure citing none is an assertion. | [optional] 
**ImplementationIds** | Pointer to **[]string** | ImplementationIDs are the artifacts it describes. | [optional] 
**Project** | Pointer to **string** | Project is the SERVER&#39;s scope value. | [optional] 
**RecordedAt** | Pointer to **string** | RecordedAt is the server&#39;s clock. | [optional] 
**Revision** | Pointer to **string** | Revision is original, corrected or retracted. | [optional] 
**Seq** | Pointer to **int64** | Seq is the append order. | [optional] 
**Status** | Pointer to **string** | Status is draft, reviewed, approved or filed. &#x60;approved&#x60; means the approvals below carry a decision for each required role, which is DERIVED on read, not a word a caller sets and gets believed. | [optional] 
**Summary** | Pointer to **string** | Summary is the invention in one paragraph, for the reader who decides whether to read the rest. | [optional] 
**Title** | Pointer to **string** | Title names the invention the way the document does. | [optional] 
**Visibility** | Pointer to **string** | Visibility is private, org or public; forced private on a write. | [optional] 

## Methods

### NewPatentPatentDisclosure

`func NewPatentPatentDisclosure() *PatentPatentDisclosure`

NewPatentPatentDisclosure instantiates a new PatentPatentDisclosure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentPatentDisclosureWithDefaults

`func NewPatentPatentDisclosureWithDefaults() *PatentPatentDisclosure`

NewPatentPatentDisclosureWithDefaults instantiates a new PatentPatentDisclosure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovals

`func (o *PatentPatentDisclosure) GetApprovals() []PatentPatentApproval`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *PatentPatentDisclosure) GetApprovalsOk() (*[]PatentPatentApproval, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *PatentPatentDisclosure) SetApprovals(v []PatentPatentApproval)`

SetApprovals sets Approvals field to given value.

### HasApprovals

`func (o *PatentPatentDisclosure) HasApprovals() bool`

HasApprovals returns a boolean if a field has been set.

### GetAuthoredBy

`func (o *PatentPatentDisclosure) GetAuthoredBy() string`

GetAuthoredBy returns the AuthoredBy field if non-nil, zero value otherwise.

### GetAuthoredByOk

`func (o *PatentPatentDisclosure) GetAuthoredByOk() (*string, bool)`

GetAuthoredByOk returns a tuple with the AuthoredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoredBy

`func (o *PatentPatentDisclosure) SetAuthoredBy(v string)`

SetAuthoredBy sets AuthoredBy field to given value.

### HasAuthoredBy

`func (o *PatentPatentDisclosure) HasAuthoredBy() bool`

HasAuthoredBy returns a boolean if a field has been set.

### GetBody

`func (o *PatentPatentDisclosure) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PatentPatentDisclosure) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PatentPatentDisclosure) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *PatentPatentDisclosure) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetBodySha256

`func (o *PatentPatentDisclosure) GetBodySha256() string`

GetBodySha256 returns the BodySha256 field if non-nil, zero value otherwise.

### GetBodySha256Ok

`func (o *PatentPatentDisclosure) GetBodySha256Ok() (*string, bool)`

GetBodySha256Ok returns a tuple with the BodySha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodySha256

`func (o *PatentPatentDisclosure) SetBodySha256(v string)`

SetBodySha256 sets BodySha256 field to given value.

### HasBodySha256

`func (o *PatentPatentDisclosure) HasBodySha256() bool`

HasBodySha256 returns a boolean if a field has been set.

### GetCanonical

`func (o *PatentPatentDisclosure) GetCanonical() bool`

GetCanonical returns the Canonical field if non-nil, zero value otherwise.

### GetCanonicalOk

`func (o *PatentPatentDisclosure) GetCanonicalOk() (*bool, bool)`

GetCanonicalOk returns a tuple with the Canonical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonical

`func (o *PatentPatentDisclosure) SetCanonical(v bool)`

SetCanonical sets Canonical field to given value.

### HasCanonical

`func (o *PatentPatentDisclosure) HasCanonical() bool`

HasCanonical returns a boolean if a field has been set.

### GetClaims

`func (o *PatentPatentDisclosure) GetClaims() string`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *PatentPatentDisclosure) GetClaimsOk() (*string, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *PatentPatentDisclosure) SetClaims(v string)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *PatentPatentDisclosure) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetContentHash

`func (o *PatentPatentDisclosure) GetContentHash() string`

GetContentHash returns the ContentHash field if non-nil, zero value otherwise.

### GetContentHashOk

`func (o *PatentPatentDisclosure) GetContentHashOk() (*string, bool)`

GetContentHashOk returns a tuple with the ContentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHash

`func (o *PatentPatentDisclosure) SetContentHash(v string)`

SetContentHash sets ContentHash field to given value.

### HasContentHash

`func (o *PatentPatentDisclosure) HasContentHash() bool`

HasContentHash returns a boolean if a field has been set.

### GetFilingRef

`func (o *PatentPatentDisclosure) GetFilingRef() string`

GetFilingRef returns the FilingRef field if non-nil, zero value otherwise.

### GetFilingRefOk

`func (o *PatentPatentDisclosure) GetFilingRefOk() (*string, bool)`

GetFilingRefOk returns a tuple with the FilingRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilingRef

`func (o *PatentPatentDisclosure) SetFilingRef(v string)`

SetFilingRef sets FilingRef field to given value.

### HasFilingRef

`func (o *PatentPatentDisclosure) HasFilingRef() bool`

HasFilingRef returns a boolean if a field has been set.

### GetId

`func (o *PatentPatentDisclosure) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatentPatentDisclosure) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatentPatentDisclosure) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatentPatentDisclosure) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdeaIds

`func (o *PatentPatentDisclosure) GetIdeaIds() []string`

GetIdeaIds returns the IdeaIds field if non-nil, zero value otherwise.

### GetIdeaIdsOk

`func (o *PatentPatentDisclosure) GetIdeaIdsOk() (*[]string, bool)`

GetIdeaIdsOk returns a tuple with the IdeaIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdeaIds

`func (o *PatentPatentDisclosure) SetIdeaIds(v []string)`

SetIdeaIds sets IdeaIds field to given value.

### HasIdeaIds

`func (o *PatentPatentDisclosure) HasIdeaIds() bool`

HasIdeaIds returns a boolean if a field has been set.

### GetImplementationIds

`func (o *PatentPatentDisclosure) GetImplementationIds() []string`

GetImplementationIds returns the ImplementationIds field if non-nil, zero value otherwise.

### GetImplementationIdsOk

`func (o *PatentPatentDisclosure) GetImplementationIdsOk() (*[]string, bool)`

GetImplementationIdsOk returns a tuple with the ImplementationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementationIds

`func (o *PatentPatentDisclosure) SetImplementationIds(v []string)`

SetImplementationIds sets ImplementationIds field to given value.

### HasImplementationIds

`func (o *PatentPatentDisclosure) HasImplementationIds() bool`

HasImplementationIds returns a boolean if a field has been set.

### GetProject

`func (o *PatentPatentDisclosure) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PatentPatentDisclosure) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PatentPatentDisclosure) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *PatentPatentDisclosure) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRecordedAt

`func (o *PatentPatentDisclosure) GetRecordedAt() string`

GetRecordedAt returns the RecordedAt field if non-nil, zero value otherwise.

### GetRecordedAtOk

`func (o *PatentPatentDisclosure) GetRecordedAtOk() (*string, bool)`

GetRecordedAtOk returns a tuple with the RecordedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordedAt

`func (o *PatentPatentDisclosure) SetRecordedAt(v string)`

SetRecordedAt sets RecordedAt field to given value.

### HasRecordedAt

`func (o *PatentPatentDisclosure) HasRecordedAt() bool`

HasRecordedAt returns a boolean if a field has been set.

### GetRevision

`func (o *PatentPatentDisclosure) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *PatentPatentDisclosure) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *PatentPatentDisclosure) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *PatentPatentDisclosure) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSeq

`func (o *PatentPatentDisclosure) GetSeq() int64`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *PatentPatentDisclosure) GetSeqOk() (*int64, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *PatentPatentDisclosure) SetSeq(v int64)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *PatentPatentDisclosure) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetStatus

`func (o *PatentPatentDisclosure) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatentPatentDisclosure) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatentPatentDisclosure) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatentPatentDisclosure) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSummary

`func (o *PatentPatentDisclosure) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PatentPatentDisclosure) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PatentPatentDisclosure) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *PatentPatentDisclosure) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *PatentPatentDisclosure) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PatentPatentDisclosure) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PatentPatentDisclosure) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PatentPatentDisclosure) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVisibility

`func (o *PatentPatentDisclosure) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *PatentPatentDisclosure) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *PatentPatentDisclosure) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *PatentPatentDisclosure) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


