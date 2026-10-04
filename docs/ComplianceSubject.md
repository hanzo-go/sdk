# ComplianceSubject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is when the subject was first recorded, Unix SECONDS. | [optional] 
**Email** | Pointer to **string** | Email is the party&#39;s address, when the org supplied one. It is PII: sealed at rest, returned only to the owning org, and never copied into a check record. | [optional] 
**Id** | Pointer to **string** | ID is the opaque handle every other record uses to point at this party. It is the only reference that leaves this type, which is what keeps the PII in one place: a check, an accreditation and an audit row all carry the id and none of them carry the name. | [optional] 
**Kind** | Pointer to **string** | Kind is what is being verified: \&quot;individual\&quot; (a natural person, so KYC) or \&quot;business\&quot; (a legal entity, so KYB). It decides which provider flow runs. | [optional] 
**Name** | Pointer to **string** | Name is the party&#39;s name, under the same PII rule as Email. For a business it is the legal entity name rather than a trading name, since that is what a provider verifies against. | [optional] 
**Org** | Pointer to **string** | Org is the tenant that is doing the verifying — the party who must answer for this record, not the party being verified. A subject is returned only to it. | [optional] 
**Ref** | Pointer to **string** | Ref is the org&#39;s OWN identifier for this party, carried so a caller can match a subject back to their system without keeping a second mapping. Opaque here: nothing in this plane parses or enforces it. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when the subject&#39;s own fields last changed, Unix seconds. A check moving to a new status does not touch it — that history lives on the check. | [optional] 

## Methods

### NewComplianceSubject

`func NewComplianceSubject() *ComplianceSubject`

NewComplianceSubject instantiates a new ComplianceSubject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComplianceSubjectWithDefaults

`func NewComplianceSubjectWithDefaults() *ComplianceSubject`

NewComplianceSubjectWithDefaults instantiates a new ComplianceSubject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ComplianceSubject) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ComplianceSubject) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ComplianceSubject) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ComplianceSubject) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *ComplianceSubject) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ComplianceSubject) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ComplianceSubject) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ComplianceSubject) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *ComplianceSubject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComplianceSubject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComplianceSubject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComplianceSubject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *ComplianceSubject) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ComplianceSubject) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ComplianceSubject) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ComplianceSubject) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ComplianceSubject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComplianceSubject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComplianceSubject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComplianceSubject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrg

`func (o *ComplianceSubject) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *ComplianceSubject) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *ComplianceSubject) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *ComplianceSubject) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetRef

`func (o *ComplianceSubject) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ComplianceSubject) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ComplianceSubject) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ComplianceSubject) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ComplianceSubject) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ComplianceSubject) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ComplianceSubject) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ComplianceSubject) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


