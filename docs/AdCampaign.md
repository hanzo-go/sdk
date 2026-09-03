# AdCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the provider ad-account the campaign runs under, in Meta&#39;s act_&lt;id&gt; form. Empty until the org supplies one or a launch resolves it. | [optional] 
**Budget** | Pointer to **int64** | Budget is the campaign&#39;s authorized spend in MINOR units (cents). Negative clamps to 0. It is the org&#39;s stored plan: a Meta launch creates the campaign object only, and the delivering budget lives on the ad set. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the campaign was first stored, in unix seconds. It never changes, including across a full-replace update. | [optional] 
**ExternalId** | Pointer to **string** | ExternalID is the ad network&#39;s own campaign id, written by a successful launch and by nothing else — an update never touches it. Empty means this campaign has never reached its network. | [optional] 
**Id** | Pointer to **string** | ID is the campaign&#39;s server-minted handle, \&quot;camp_\&quot; + 32 hex. A create body cannot choose it, and it is the id every other route addresses. | [optional] 
**Name** | Pointer to **string** | Name is the campaign&#39;s display label, and the name Meta creates the campaign object under at launch. Required; trimmed and bounded to 1024 bytes. | [optional] 
**Objective** | Pointer to **string** | Objective is the campaign goal spelled as the provider names it (\&quot;conversions\&quot;, \&quot;OUTCOME_TRAFFIC\&quot;), passed through to the network verbatim at launch — Meta defaults an empty one to OUTCOME_TRAFFIC. Free text, bounded to 1024 bytes; no vocabulary is enforced here. | [optional] 
**Platform** | Pointer to **string** | Platform is the ad network: meta, google, tiktok or x, and nothing else — a write naming another is 400. Empty stores as meta. Only meta executes today; launching any of the other three is 501. | [optional] 
**Spend** | Pointer to **int64** | Spend is spend-to-date in MINOR units (cents), as last written through create or update. Negative clamps to 0. It is NOT read back from the network — that is a separate insights call — so 0 means nothing was recorded here, not that nothing was spent. | [optional] 
**Status** | Pointer to **string** | Status is the lifecycle: draft, active, paused or completed, and nothing else — a write naming another is 400. Empty stores as draft; a successful launch sets active. It records what this deployment did, not what the ad network currently reports. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when the row was last written, in unix seconds — set by create, update and launch. Listings are ordered by it, newest first. | [optional] 

## Methods

### NewAdCampaign

`func NewAdCampaign() *AdCampaign`

NewAdCampaign instantiates a new AdCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdCampaignWithDefaults

`func NewAdCampaignWithDefaults() *AdCampaign`

NewAdCampaignWithDefaults instantiates a new AdCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AdCampaign) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AdCampaign) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AdCampaign) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AdCampaign) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetBudget

`func (o *AdCampaign) GetBudget() int64`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *AdCampaign) GetBudgetOk() (*int64, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *AdCampaign) SetBudget(v int64)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *AdCampaign) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AdCampaign) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdCampaign) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdCampaign) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdCampaign) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExternalId

`func (o *AdCampaign) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AdCampaign) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AdCampaign) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AdCampaign) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *AdCampaign) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdCampaign) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdCampaign) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdCampaign) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AdCampaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdCampaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdCampaign) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdCampaign) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjective

`func (o *AdCampaign) GetObjective() string`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *AdCampaign) GetObjectiveOk() (*string, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *AdCampaign) SetObjective(v string)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *AdCampaign) HasObjective() bool`

HasObjective returns a boolean if a field has been set.

### GetPlatform

`func (o *AdCampaign) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AdCampaign) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AdCampaign) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AdCampaign) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSpend

`func (o *AdCampaign) GetSpend() int64`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *AdCampaign) GetSpendOk() (*int64, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *AdCampaign) SetSpend(v int64)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *AdCampaign) HasSpend() bool`

HasSpend returns a boolean if a field has been set.

### GetStatus

`func (o *AdCampaign) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdCampaign) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdCampaign) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdCampaign) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdCampaign) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdCampaign) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdCampaign) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdCampaign) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


