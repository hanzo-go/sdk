# Campaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | Pointer to **int64** | Budget is what the campaign is allowed to cost, in USD cents. A negative value is clamped to 0; nothing enforces the ceiling here. | [optional] 
**Channel** | Pointer to **string** | Channel is the delivery surface: email, sms, social, meta, google or tiktok. Empty means email. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is unix seconds when the campaign was registered. Server-assigned and never rewritten — an update leaves it as it was. | [optional] 
**Id** | Pointer to **string** | ID is the server-assigned campaign id (\&quot;camp_\&quot; + 128 random bits). | [optional] 
**Name** | Pointer to **string** | Name is the campaign&#39;s label. Required, trimmed, capped at 1024 bytes. | [optional] 
**Objective** | Pointer to **string** | Objective is the free-text goal (\&quot;signups\&quot;), capped at 1024 bytes. | [optional] 
**ScheduledAt** | Pointer to **int64** | ScheduledAt is the unix send time; 0 means unscheduled. Setting it on a campaign with no explicit status makes that status \&quot;scheduled\&quot;. | [optional] 
**Spend** | Pointer to **int64** | Spend is what the campaign has cost so far, in USD cents, clamped to &gt;&#x3D; 0. The CALLER owns it: no send, ad buy or invoice moves it, so it changes only when create or update carries a new value. It is summed across the org&#39;s campaigns into GET /v1/marketing/summary. | [optional] 
**Status** | Pointer to **string** | Status is the lifecycle: draft, scheduled, active, paused or completed. Empty means draft. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is unix seconds of the last write. Server-assigned on create and on every update or schedule change, and the campaign list is ordered by it, newest first. | [optional] 

## Methods

### NewCampaign

`func NewCampaign() *Campaign`

NewCampaign instantiates a new Campaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignWithDefaults

`func NewCampaignWithDefaults() *Campaign`

NewCampaignWithDefaults instantiates a new Campaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudget

`func (o *Campaign) GetBudget() int64`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *Campaign) GetBudgetOk() (*int64, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *Campaign) SetBudget(v int64)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *Campaign) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetChannel

`func (o *Campaign) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Campaign) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Campaign) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *Campaign) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Campaign) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Campaign) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Campaign) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Campaign) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *Campaign) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Campaign) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Campaign) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Campaign) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Campaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Campaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Campaign) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Campaign) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjective

`func (o *Campaign) GetObjective() string`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *Campaign) GetObjectiveOk() (*string, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *Campaign) SetObjective(v string)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *Campaign) HasObjective() bool`

HasObjective returns a boolean if a field has been set.

### GetScheduledAt

`func (o *Campaign) GetScheduledAt() int64`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *Campaign) GetScheduledAtOk() (*int64, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *Campaign) SetScheduledAt(v int64)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *Campaign) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetSpend

`func (o *Campaign) GetSpend() int64`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *Campaign) GetSpendOk() (*int64, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *Campaign) SetSpend(v int64)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *Campaign) HasSpend() bool`

HasSpend returns a boolean if a field has been set.

### GetStatus

`func (o *Campaign) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Campaign) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Campaign) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Campaign) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Campaign) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Campaign) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Campaign) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Campaign) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


