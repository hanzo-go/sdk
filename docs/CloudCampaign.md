# CloudCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | Pointer to **int32** | Budget and Spend are minor units (USD cents), clamped to &gt;&#x3D; 0. | [optional] 
**Channel** | Pointer to **string** | Channel is the delivery surface: email, sms, social, meta, google or tiktok. Empty means email. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt and UpdatedAt are unix seconds, both server-assigned. | [optional] 
**Id** | Pointer to **string** | ID is the server-assigned campaign id (\&quot;camp_\&quot; + 128 random bits). | [optional] 
**Name** | Pointer to **string** | Name is the campaign&#39;s label. Required, trimmed, capped at 1024 bytes. | [optional] 
**Objective** | Pointer to **string** | Objective is the free-text goal (\&quot;signups\&quot;), capped at 1024 bytes. | [optional] 
**ScheduledAt** | Pointer to **int32** | ScheduledAt is the unix send time; 0 means unscheduled. Setting it on a campaign with no explicit status makes that status \&quot;scheduled\&quot;. | [optional] 
**Spend** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** | Status is the lifecycle: draft, scheduled, active, paused or completed. Empty means draft. | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudCampaign

`func NewCloudCampaign() *CloudCampaign`

NewCloudCampaign instantiates a new CloudCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCampaignWithDefaults

`func NewCloudCampaignWithDefaults() *CloudCampaign`

NewCloudCampaignWithDefaults instantiates a new CloudCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudget

`func (o *CloudCampaign) GetBudget() int32`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *CloudCampaign) GetBudgetOk() (*int32, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *CloudCampaign) SetBudget(v int32)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *CloudCampaign) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetChannel

`func (o *CloudCampaign) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CloudCampaign) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CloudCampaign) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CloudCampaign) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudCampaign) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudCampaign) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudCampaign) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudCampaign) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CloudCampaign) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCampaign) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCampaign) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCampaign) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudCampaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCampaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCampaign) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCampaign) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjective

`func (o *CloudCampaign) GetObjective() string`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *CloudCampaign) GetObjectiveOk() (*string, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *CloudCampaign) SetObjective(v string)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *CloudCampaign) HasObjective() bool`

HasObjective returns a boolean if a field has been set.

### GetScheduledAt

`func (o *CloudCampaign) GetScheduledAt() int32`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *CloudCampaign) GetScheduledAtOk() (*int32, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *CloudCampaign) SetScheduledAt(v int32)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *CloudCampaign) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetSpend

`func (o *CloudCampaign) GetSpend() int32`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *CloudCampaign) GetSpendOk() (*int32, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *CloudCampaign) SetSpend(v int32)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *CloudCampaign) HasSpend() bool`

HasSpend returns a boolean if a field has been set.

### GetStatus

`func (o *CloudCampaign) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudCampaign) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudCampaign) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudCampaign) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudCampaign) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudCampaign) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudCampaign) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudCampaign) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


