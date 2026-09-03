# CampaignRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **string** | Audience is an opaque reference to the segment this campaign targets. It is stored and echoed but not yet handed to the executors — a channel targets through the provider account it runs under — so it is documentation for now. Absent when never set. | [optional] 
**Budget** | Pointer to **int64** | Budget is the campaign&#39;s total budget in CENTS, handed to each executor as the budget for its channel. 0 means none was set. | [optional] 
**Channels** | Pointer to [**[]ChannelSpec**](ChannelSpec.md) | Channels are the fan-out targets, at most one per kind and at most 12, each carrying its own post-launch state. Empty means nothing to launch, which is what makes a launch of this campaign a 400. | [optional] 
**Content** | Pointer to **[]string** | Content is the ordered creative set, at most 32, empty entries dropped. Content[0] is the creative that runs; the rest are A/B variants a wired experiment can assign per launch. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is when the campaign was created, in unix seconds. Server-set. | [optional] 
**Id** | Pointer to **string** | ID is the campaign&#39;s server-minted handle — \&quot;cmp_\&quot; and 128 random bits — and the id every other campaign call is addressed by. Never read off the wire: a create that sends one has it ignored. | [optional] 
**Name** | Pointer to **string** | Name is the campaign&#39;s display name. Required on write, trimmed, and capped at 2048 characters. | [optional] 
**ScheduleAt** | Pointer to **int64** | ScheduleAt is when the campaign should run, in unix seconds. 0 (absent) means launch immediately. It is passed to each executor; nothing in this service wakes up to launch it for you. | [optional] 
**Status** | Pointer to **string** | Status is the lifecycle state, server-owned and never accepted from a caller. Four values actually occur: draft (inert and fully mutable — nothing is sent and no budget is committed), live, paused and failed. After a fan-out live means AT LEAST ONE channel launched — read the channel rows for the rest — and failed means none did. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is the last write in unix seconds — an edit, a launch or a pause. Server-set on every save. | [optional] 

## Methods

### NewCampaignRecord

`func NewCampaignRecord() *CampaignRecord`

NewCampaignRecord instantiates a new CampaignRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignRecordWithDefaults

`func NewCampaignRecordWithDefaults() *CampaignRecord`

NewCampaignRecordWithDefaults instantiates a new CampaignRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *CampaignRecord) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CampaignRecord) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CampaignRecord) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *CampaignRecord) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetBudget

`func (o *CampaignRecord) GetBudget() int64`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *CampaignRecord) GetBudgetOk() (*int64, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *CampaignRecord) SetBudget(v int64)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *CampaignRecord) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetChannels

`func (o *CampaignRecord) GetChannels() []ChannelSpec`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *CampaignRecord) GetChannelsOk() (*[]ChannelSpec, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *CampaignRecord) SetChannels(v []ChannelSpec)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *CampaignRecord) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetContent

`func (o *CampaignRecord) GetContent() []string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CampaignRecord) GetContentOk() (*[]string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CampaignRecord) SetContent(v []string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CampaignRecord) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CampaignRecord) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CampaignRecord) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CampaignRecord) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CampaignRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CampaignRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CampaignRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CampaignRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduleAt

`func (o *CampaignRecord) GetScheduleAt() int64`

GetScheduleAt returns the ScheduleAt field if non-nil, zero value otherwise.

### GetScheduleAtOk

`func (o *CampaignRecord) GetScheduleAtOk() (*int64, bool)`

GetScheduleAtOk returns a tuple with the ScheduleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleAt

`func (o *CampaignRecord) SetScheduleAt(v int64)`

SetScheduleAt sets ScheduleAt field to given value.

### HasScheduleAt

`func (o *CampaignRecord) HasScheduleAt() bool`

HasScheduleAt returns a boolean if a field has been set.

### GetStatus

`func (o *CampaignRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CampaignRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CampaignRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CampaignRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CampaignRecord) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CampaignRecord) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CampaignRecord) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CampaignRecord) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


