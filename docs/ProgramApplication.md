# ProgramApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to **string** | Company is the applicant&#39;s company name. | [optional] 
**CompanyId** | Pointer to **string** | CompanyID is the CRM Company minted for this lead at intake, so the startup also appears in the org&#39;s standard CRM tabs. Empty when that best-effort projection did not run. | [optional] 
**ContactId** | Pointer to **string** | ContactID is the CRM Contact minted for this lead at intake. Empty when that best-effort projection did not run. | [optional] 
**ContactName** | Pointer to **string** | ContactName is the person who applied. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is the unix second the application arrived. Server-owned. | [optional] 
**Email** | Pointer to **string** | Email is the applicant&#39;s email — half of the (email, company) key a resubmission refreshes instead of duplicating. | [optional] 
**Events** | Pointer to [**[]StageEvent**](StageEvent.md) | Events is the append-only stage-transition log, oldest first. | [optional] 
**Id** | Pointer to **string** | ID is the server-minted application id (\&quot;appl_\&quot; + 128 random bits). | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata is the FULL submitted form, every field, including the arrays the promoted columns above do not carry (tier1Investors, useCases) and the deterministic tier1Matched list. | [optional] 
**Reason** | Pointer to **string** | Reason is why the application was rejected, required to reject. Empty otherwise. | [optional] 
**Role** | Pointer to **string** | Role is the applicant&#39;s role at their company. | [optional] 
**Screen** | Pointer to [**ScreenResult**](ScreenResult.md) | Screen is the AI screen. It runs after intake, so a freshly created application carries a \&quot;pending\&quot; screen. | [optional] 
**Stage** | Pointer to **string** | Stage is the pipeline stage: applied, screened, qualified, credits-offered, onboarded or rejected. Server-owned — it starts at \&quot;applied\&quot; and moves only through the transition machine. | [optional] 
**Tier1** | Pointer to **bool** | Tier1 is whether the applicant is tier-1 backed, derived deterministically at intake from the submitted fund list — independent of the AI screen. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is the unix second of the last write. Server-owned. | [optional] 
**Website** | Pointer to **string** | Website is the applicant&#39;s website as submitted. | [optional] 

## Methods

### NewProgramApplication

`func NewProgramApplication() *ProgramApplication`

NewProgramApplication instantiates a new ProgramApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramApplicationWithDefaults

`func NewProgramApplicationWithDefaults() *ProgramApplication`

NewProgramApplicationWithDefaults instantiates a new ProgramApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *ProgramApplication) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ProgramApplication) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ProgramApplication) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ProgramApplication) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCompanyId

`func (o *ProgramApplication) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *ProgramApplication) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *ProgramApplication) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *ProgramApplication) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetContactId

`func (o *ProgramApplication) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *ProgramApplication) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *ProgramApplication) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *ProgramApplication) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### GetContactName

`func (o *ProgramApplication) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *ProgramApplication) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *ProgramApplication) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *ProgramApplication) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProgramApplication) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProgramApplication) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProgramApplication) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProgramApplication) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *ProgramApplication) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ProgramApplication) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ProgramApplication) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ProgramApplication) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEvents

`func (o *ProgramApplication) GetEvents() []StageEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ProgramApplication) GetEventsOk() (*[]StageEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ProgramApplication) SetEvents(v []StageEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ProgramApplication) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetId

`func (o *ProgramApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProgramApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProgramApplication) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProgramApplication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ProgramApplication) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProgramApplication) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProgramApplication) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ProgramApplication) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReason

`func (o *ProgramApplication) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ProgramApplication) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ProgramApplication) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ProgramApplication) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRole

`func (o *ProgramApplication) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ProgramApplication) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ProgramApplication) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ProgramApplication) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetScreen

`func (o *ProgramApplication) GetScreen() ScreenResult`

GetScreen returns the Screen field if non-nil, zero value otherwise.

### GetScreenOk

`func (o *ProgramApplication) GetScreenOk() (*ScreenResult, bool)`

GetScreenOk returns a tuple with the Screen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreen

`func (o *ProgramApplication) SetScreen(v ScreenResult)`

SetScreen sets Screen field to given value.

### HasScreen

`func (o *ProgramApplication) HasScreen() bool`

HasScreen returns a boolean if a field has been set.

### GetStage

`func (o *ProgramApplication) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *ProgramApplication) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *ProgramApplication) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *ProgramApplication) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetTier1

`func (o *ProgramApplication) GetTier1() bool`

GetTier1 returns the Tier1 field if non-nil, zero value otherwise.

### GetTier1Ok

`func (o *ProgramApplication) GetTier1Ok() (*bool, bool)`

GetTier1Ok returns a tuple with the Tier1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier1

`func (o *ProgramApplication) SetTier1(v bool)`

SetTier1 sets Tier1 field to given value.

### HasTier1

`func (o *ProgramApplication) HasTier1() bool`

HasTier1 returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProgramApplication) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProgramApplication) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProgramApplication) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProgramApplication) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWebsite

`func (o *ProgramApplication) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *ProgramApplication) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *ProgramApplication) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *ProgramApplication) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


