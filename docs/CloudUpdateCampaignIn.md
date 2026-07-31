# CloudUpdateCampaignIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the provider ad-account this campaign runs on (Meta act_&lt;id&gt;). Optional. | [optional] 
**Budget** | Pointer to **int32** | Budget is the campaign budget in MINOR units (cents). Negative values clamp to 0. | [optional] 
**Name** | Pointer to **string** | Name is the campaign&#39;s display label. Required; trimmed and bounded to 1024 bytes. | [optional] 
**Objective** | Pointer to **string** | Objective is the campaign goal as the provider names it. Optional, bounded to 1024 bytes. | [optional] 
**Platform** | Pointer to **string** | Platform is the ad network: meta, google, tiktok or x. Empty defaults to meta. | [optional] 
**Spend** | Pointer to **int32** | Spend is the amount spent so far in MINOR units (cents). Negative values clamp to 0. | [optional] 
**Status** | Pointer to **string** | Status is the lifecycle state: draft, active, paused or completed. Empty defaults to draft. | [optional] 

## Methods

### NewCloudUpdateCampaignIn

`func NewCloudUpdateCampaignIn() *CloudUpdateCampaignIn`

NewCloudUpdateCampaignIn instantiates a new CloudUpdateCampaignIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudUpdateCampaignInWithDefaults

`func NewCloudUpdateCampaignInWithDefaults() *CloudUpdateCampaignIn`

NewCloudUpdateCampaignInWithDefaults instantiates a new CloudUpdateCampaignIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudUpdateCampaignIn) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudUpdateCampaignIn) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudUpdateCampaignIn) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudUpdateCampaignIn) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetBudget

`func (o *CloudUpdateCampaignIn) GetBudget() int32`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *CloudUpdateCampaignIn) GetBudgetOk() (*int32, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *CloudUpdateCampaignIn) SetBudget(v int32)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *CloudUpdateCampaignIn) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetName

`func (o *CloudUpdateCampaignIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudUpdateCampaignIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudUpdateCampaignIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudUpdateCampaignIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjective

`func (o *CloudUpdateCampaignIn) GetObjective() string`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *CloudUpdateCampaignIn) GetObjectiveOk() (*string, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *CloudUpdateCampaignIn) SetObjective(v string)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *CloudUpdateCampaignIn) HasObjective() bool`

HasObjective returns a boolean if a field has been set.

### GetPlatform

`func (o *CloudUpdateCampaignIn) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CloudUpdateCampaignIn) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CloudUpdateCampaignIn) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CloudUpdateCampaignIn) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSpend

`func (o *CloudUpdateCampaignIn) GetSpend() int32`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *CloudUpdateCampaignIn) GetSpendOk() (*int32, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *CloudUpdateCampaignIn) SetSpend(v int32)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *CloudUpdateCampaignIn) HasSpend() bool`

HasSpend returns a boolean if a field has been set.

### GetStatus

`func (o *CloudUpdateCampaignIn) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudUpdateCampaignIn) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudUpdateCampaignIn) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudUpdateCampaignIn) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


