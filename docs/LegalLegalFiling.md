# LegalLegalFiling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is when the filing was recorded, in unix seconds. | [optional] 
**DocumentIds** | Pointer to **[]string** | DocumentIDs are the generated documents this filing covers. At least one, and each must already belong to Org — a filing can never reach across tenants. | [optional] 
**Id** | Pointer to **string** | ID is the filing&#39;s server-minted handle, \&quot;filing_\&quot;-prefixed. It is OURS, not the state&#39;s — a state file number arrives only from a real filing partner. | [optional] 
**Jurisdiction** | Pointer to **string** | Jurisdiction is the state or agency the filing is for, as the caller wrote it (\&quot;DE\&quot;). Free text, not validated against a list. Absent when the caller sent none. | [optional] 
**Note** | Pointer to **string** | Note is the provider&#39;s explanation of Status, in prose meant for a person. On the manual default it says the documents were generated for signature and the org files them through its registered agent. Absent when the provider offered nothing. | [optional] 
**Org** | Pointer to **string** | Org owns the filing. Every query is scoped to it, so another tenant&#39;s filing is indistinguishable from a missing one. | [optional] 
**Provider** | Pointer to **string** | Provider is the filing partner that took the submission, or \&quot;manual\&quot; when no partner is wired on this deployment. | [optional] 
**Status** | Pointer to **string** | Status is where the filing stands: manual (no partner wired — file through your registered agent), submitted (a partner accepted it), filed (the state accepted it) or rejected. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when it last changed, in unix seconds. Equal to CreatedAt until a provider reports movement. | [optional] 

## Methods

### NewLegalLegalFiling

`func NewLegalLegalFiling() *LegalLegalFiling`

NewLegalLegalFiling instantiates a new LegalLegalFiling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalLegalFilingWithDefaults

`func NewLegalLegalFilingWithDefaults() *LegalLegalFiling`

NewLegalLegalFilingWithDefaults instantiates a new LegalLegalFiling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LegalLegalFiling) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LegalLegalFiling) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LegalLegalFiling) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LegalLegalFiling) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDocumentIds

`func (o *LegalLegalFiling) GetDocumentIds() []string`

GetDocumentIds returns the DocumentIds field if non-nil, zero value otherwise.

### GetDocumentIdsOk

`func (o *LegalLegalFiling) GetDocumentIdsOk() (*[]string, bool)`

GetDocumentIdsOk returns a tuple with the DocumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIds

`func (o *LegalLegalFiling) SetDocumentIds(v []string)`

SetDocumentIds sets DocumentIds field to given value.

### HasDocumentIds

`func (o *LegalLegalFiling) HasDocumentIds() bool`

HasDocumentIds returns a boolean if a field has been set.

### GetId

`func (o *LegalLegalFiling) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LegalLegalFiling) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LegalLegalFiling) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LegalLegalFiling) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJurisdiction

`func (o *LegalLegalFiling) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *LegalLegalFiling) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *LegalLegalFiling) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *LegalLegalFiling) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.

### GetNote

`func (o *LegalLegalFiling) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *LegalLegalFiling) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *LegalLegalFiling) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *LegalLegalFiling) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOrg

`func (o *LegalLegalFiling) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *LegalLegalFiling) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *LegalLegalFiling) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *LegalLegalFiling) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetProvider

`func (o *LegalLegalFiling) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *LegalLegalFiling) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *LegalLegalFiling) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *LegalLegalFiling) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *LegalLegalFiling) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LegalLegalFiling) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LegalLegalFiling) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LegalLegalFiling) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LegalLegalFiling) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LegalLegalFiling) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LegalLegalFiling) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LegalLegalFiling) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


