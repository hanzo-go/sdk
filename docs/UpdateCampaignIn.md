# UpdateCampaignIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the provider ad-account this campaign runs on (Meta act_&lt;id&gt;). Optional. | [optional] 
**Budget** | Pointer to **int64** | Budget is the campaign budget in MINOR units (cents). Negative values clamp to 0. | [optional] 
**Name** | Pointer to **string** | Name is the campaign&#39;s display label. Required; trimmed and bounded to 1024 bytes. | [optional] 
**Objective** | Pointer to **string** | Objective is the campaign goal as the provider names it. Optional, bounded to 1024 bytes. | [optional] 
**Platform** | Pointer to **string** | Platform is the ad network: meta, google, tiktok or x. Empty defaults to meta. | [optional] 
**Spend** | Pointer to **int64** | Spend is the amount spent so far in MINOR units (cents). Negative values clamp to 0. | [optional] 
**Status** | Pointer to **string** | Status is the lifecycle state: draft, active, paused or completed. Empty defaults to draft. | [optional] 

## Methods

### NewUpdateCampaignIn

`func NewUpdateCampaignIn() *UpdateCampaignIn`

NewUpdateCampaignIn instantiates a new UpdateCampaignIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCampaignInWithDefaults

`func NewUpdateCampaignInWithDefaults() *UpdateCampaignIn`

NewUpdateCampaignInWithDefaults instantiates a new UpdateCampaignIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *UpdateCampaignIn) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateCampaignIn) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateCampaignIn) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateCampaignIn) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetBudget

`func (o *UpdateCampaignIn) GetBudget() int64`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *UpdateCampaignIn) GetBudgetOk() (*int64, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *UpdateCampaignIn) SetBudget(v int64)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *UpdateCampaignIn) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetName

`func (o *UpdateCampaignIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCampaignIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCampaignIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCampaignIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjective

`func (o *UpdateCampaignIn) GetObjective() string`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *UpdateCampaignIn) GetObjectiveOk() (*string, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *UpdateCampaignIn) SetObjective(v string)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *UpdateCampaignIn) HasObjective() bool`

HasObjective returns a boolean if a field has been set.

### GetPlatform

`func (o *UpdateCampaignIn) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UpdateCampaignIn) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UpdateCampaignIn) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *UpdateCampaignIn) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSpend

`func (o *UpdateCampaignIn) GetSpend() int64`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *UpdateCampaignIn) GetSpendOk() (*int64, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *UpdateCampaignIn) SetSpend(v int64)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *UpdateCampaignIn) HasSpend() bool`

HasSpend returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateCampaignIn) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateCampaignIn) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateCampaignIn) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateCampaignIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


