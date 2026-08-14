# Formation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlreadyIncorporated** | Pointer to **bool** | AlreadyIncorporated declares an org that already has a legal entity, which takes the import path (structure → import → company) instead of forming one. | [optional] 
**CapTableImported** | Pointer to **bool** | CapTableImported reports whether the existing company&#39;s cap table has been imported onto the canonical cap table. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is the unix second the formation was opened. | [optional] 
**DocumentIds** | Pointer to **[]string** | DocumentIDs are the data room ids of the GENERATED formation documents. | [optional] 
**EsignRef** | Pointer to **string** | EsignRef is the e-signature provider&#39;s reference for the signature request. | [optional] 
**Filing** | Pointer to [**Filing**](Filing.md) | Filing is the state-of-incorporation filing record, once documents exist. | [optional] 
**Founders** | Pointer to [**[]Founder**](Founder.md) | Founders is every founding stakeholder, with its equity split and KYC state. | [optional] 
**Genesis** | Pointer to [**Genesis**](Genesis.md) | Genesis is the cap-table equity genesis, once recorded. | [optional] 
**Imported** | Pointer to **bool** | Imported reports whether the existing company&#39;s corporate documents have been ingested into the org&#39;s data room. | [optional] 
**ImportedDocs** | Pointer to **[]string** | ImportedDocs are the data room ids of the documents ingested from Drive. | [optional] 
**Jurisdiction** | Pointer to **string** | Jurisdiction is the state of formation: DE or WY. | [optional] 
**Name** | Pointer to **string** | Name is the company name the entity is being formed under. | [optional] 
**Org** | Pointer to **string** | Org is the owning org — the tenant key, and the reason there is exactly one formation per org. | [optional] 
**Paid** | Pointer to **bool** | Paid reports whether the one-time formation fee has been charged. | [optional] 
**PaymentRef** | Pointer to **string** | PaymentRef is the billing reference recorded for the charged formation fee on the org&#39;s own ledger. | [optional] 
**Signed** | Pointer to **bool** | Signed reports whether the formation documents have come back signed — the e-signature provider&#39;s answer, which a real provider&#39;s webhook drives. | [optional] 
**Stage** | Pointer to **string** | Stage is the machine&#39;s current state: structure, founders, payment, documents, esign or genesis on the formation path, import on the skip path, and company at the terminal. | [optional] 
**Structure** | Pointer to **string** | Structure is the legal entity being formed: c-corp, llc or dao-llc. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is the unix second of the most recent write to the formation. | [optional] 

## Methods

### NewFormation

`func NewFormation() *Formation`

NewFormation instantiates a new Formation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormationWithDefaults

`func NewFormationWithDefaults() *Formation`

NewFormationWithDefaults instantiates a new Formation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlreadyIncorporated

`func (o *Formation) GetAlreadyIncorporated() bool`

GetAlreadyIncorporated returns the AlreadyIncorporated field if non-nil, zero value otherwise.

### GetAlreadyIncorporatedOk

`func (o *Formation) GetAlreadyIncorporatedOk() (*bool, bool)`

GetAlreadyIncorporatedOk returns a tuple with the AlreadyIncorporated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyIncorporated

`func (o *Formation) SetAlreadyIncorporated(v bool)`

SetAlreadyIncorporated sets AlreadyIncorporated field to given value.

### HasAlreadyIncorporated

`func (o *Formation) HasAlreadyIncorporated() bool`

HasAlreadyIncorporated returns a boolean if a field has been set.

### GetCapTableImported

`func (o *Formation) GetCapTableImported() bool`

GetCapTableImported returns the CapTableImported field if non-nil, zero value otherwise.

### GetCapTableImportedOk

`func (o *Formation) GetCapTableImportedOk() (*bool, bool)`

GetCapTableImportedOk returns a tuple with the CapTableImported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapTableImported

`func (o *Formation) SetCapTableImported(v bool)`

SetCapTableImported sets CapTableImported field to given value.

### HasCapTableImported

`func (o *Formation) HasCapTableImported() bool`

HasCapTableImported returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Formation) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Formation) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Formation) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Formation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDocumentIds

`func (o *Formation) GetDocumentIds() []string`

GetDocumentIds returns the DocumentIds field if non-nil, zero value otherwise.

### GetDocumentIdsOk

`func (o *Formation) GetDocumentIdsOk() (*[]string, bool)`

GetDocumentIdsOk returns a tuple with the DocumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIds

`func (o *Formation) SetDocumentIds(v []string)`

SetDocumentIds sets DocumentIds field to given value.

### HasDocumentIds

`func (o *Formation) HasDocumentIds() bool`

HasDocumentIds returns a boolean if a field has been set.

### GetEsignRef

`func (o *Formation) GetEsignRef() string`

GetEsignRef returns the EsignRef field if non-nil, zero value otherwise.

### GetEsignRefOk

`func (o *Formation) GetEsignRefOk() (*string, bool)`

GetEsignRefOk returns a tuple with the EsignRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsignRef

`func (o *Formation) SetEsignRef(v string)`

SetEsignRef sets EsignRef field to given value.

### HasEsignRef

`func (o *Formation) HasEsignRef() bool`

HasEsignRef returns a boolean if a field has been set.

### GetFiling

`func (o *Formation) GetFiling() Filing`

GetFiling returns the Filing field if non-nil, zero value otherwise.

### GetFilingOk

`func (o *Formation) GetFilingOk() (*Filing, bool)`

GetFilingOk returns a tuple with the Filing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiling

`func (o *Formation) SetFiling(v Filing)`

SetFiling sets Filing field to given value.

### HasFiling

`func (o *Formation) HasFiling() bool`

HasFiling returns a boolean if a field has been set.

### GetFounders

`func (o *Formation) GetFounders() []Founder`

GetFounders returns the Founders field if non-nil, zero value otherwise.

### GetFoundersOk

`func (o *Formation) GetFoundersOk() (*[]Founder, bool)`

GetFoundersOk returns a tuple with the Founders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFounders

`func (o *Formation) SetFounders(v []Founder)`

SetFounders sets Founders field to given value.

### HasFounders

`func (o *Formation) HasFounders() bool`

HasFounders returns a boolean if a field has been set.

### GetGenesis

`func (o *Formation) GetGenesis() Genesis`

GetGenesis returns the Genesis field if non-nil, zero value otherwise.

### GetGenesisOk

`func (o *Formation) GetGenesisOk() (*Genesis, bool)`

GetGenesisOk returns a tuple with the Genesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenesis

`func (o *Formation) SetGenesis(v Genesis)`

SetGenesis sets Genesis field to given value.

### HasGenesis

`func (o *Formation) HasGenesis() bool`

HasGenesis returns a boolean if a field has been set.

### GetImported

`func (o *Formation) GetImported() bool`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *Formation) GetImportedOk() (*bool, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *Formation) SetImported(v bool)`

SetImported sets Imported field to given value.

### HasImported

`func (o *Formation) HasImported() bool`

HasImported returns a boolean if a field has been set.

### GetImportedDocs

`func (o *Formation) GetImportedDocs() []string`

GetImportedDocs returns the ImportedDocs field if non-nil, zero value otherwise.

### GetImportedDocsOk

`func (o *Formation) GetImportedDocsOk() (*[]string, bool)`

GetImportedDocsOk returns a tuple with the ImportedDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedDocs

`func (o *Formation) SetImportedDocs(v []string)`

SetImportedDocs sets ImportedDocs field to given value.

### HasImportedDocs

`func (o *Formation) HasImportedDocs() bool`

HasImportedDocs returns a boolean if a field has been set.

### GetJurisdiction

`func (o *Formation) GetJurisdiction() string`

GetJurisdiction returns the Jurisdiction field if non-nil, zero value otherwise.

### GetJurisdictionOk

`func (o *Formation) GetJurisdictionOk() (*string, bool)`

GetJurisdictionOk returns a tuple with the Jurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJurisdiction

`func (o *Formation) SetJurisdiction(v string)`

SetJurisdiction sets Jurisdiction field to given value.

### HasJurisdiction

`func (o *Formation) HasJurisdiction() bool`

HasJurisdiction returns a boolean if a field has been set.

### GetName

`func (o *Formation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Formation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Formation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Formation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *Formation) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *Formation) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *Formation) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *Formation) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPaid

`func (o *Formation) GetPaid() bool`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *Formation) GetPaidOk() (*bool, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *Formation) SetPaid(v bool)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *Formation) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetPaymentRef

`func (o *Formation) GetPaymentRef() string`

GetPaymentRef returns the PaymentRef field if non-nil, zero value otherwise.

### GetPaymentRefOk

`func (o *Formation) GetPaymentRefOk() (*string, bool)`

GetPaymentRefOk returns a tuple with the PaymentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRef

`func (o *Formation) SetPaymentRef(v string)`

SetPaymentRef sets PaymentRef field to given value.

### HasPaymentRef

`func (o *Formation) HasPaymentRef() bool`

HasPaymentRef returns a boolean if a field has been set.

### GetSigned

`func (o *Formation) GetSigned() bool`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *Formation) GetSignedOk() (*bool, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *Formation) SetSigned(v bool)`

SetSigned sets Signed field to given value.

### HasSigned

`func (o *Formation) HasSigned() bool`

HasSigned returns a boolean if a field has been set.

### GetStage

`func (o *Formation) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *Formation) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *Formation) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *Formation) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetStructure

`func (o *Formation) GetStructure() string`

GetStructure returns the Structure field if non-nil, zero value otherwise.

### GetStructureOk

`func (o *Formation) GetStructureOk() (*string, bool)`

GetStructureOk returns a tuple with the Structure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructure

`func (o *Formation) SetStructure(v string)`

SetStructure sets Structure field to given value.

### HasStructure

`func (o *Formation) HasStructure() bool`

HasStructure returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Formation) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Formation) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Formation) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Formation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


