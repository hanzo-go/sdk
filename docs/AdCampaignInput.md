# AdCampaignInput

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

### NewAdCampaignInput

`func NewAdCampaignInput() *AdCampaignInput`

NewAdCampaignInput instantiates a new AdCampaignInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdCampaignInputWithDefaults

`func NewAdCampaignInputWithDefaults() *AdCampaignInput`

NewAdCampaignInputWithDefaults instantiates a new AdCampaignInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AdCampaignInput) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AdCampaignInput) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AdCampaignInput) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AdCampaignInput) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetBudget

`func (o *AdCampaignInput) GetBudget() int64`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *AdCampaignInput) GetBudgetOk() (*int64, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *AdCampaignInput) SetBudget(v int64)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *AdCampaignInput) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetName

`func (o *AdCampaignInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdCampaignInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdCampaignInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdCampaignInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjective

`func (o *AdCampaignInput) GetObjective() string`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *AdCampaignInput) GetObjectiveOk() (*string, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *AdCampaignInput) SetObjective(v string)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *AdCampaignInput) HasObjective() bool`

HasObjective returns a boolean if a field has been set.

### GetPlatform

`func (o *AdCampaignInput) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AdCampaignInput) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AdCampaignInput) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AdCampaignInput) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSpend

`func (o *AdCampaignInput) GetSpend() int64`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *AdCampaignInput) GetSpendOk() (*int64, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *AdCampaignInput) SetSpend(v int64)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *AdCampaignInput) HasSpend() bool`

HasSpend returns a boolean if a field has been set.

### GetStatus

`func (o *AdCampaignInput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdCampaignInput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdCampaignInput) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdCampaignInput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


