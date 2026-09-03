# CampaignWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **string** | Audience is the segment or audience selector this campaign targets. | [optional] 
**Budget** | Pointer to **int64** | Budget is the campaign&#39;s total budget in CENTS. Negative reads as 0. | [optional] 
**Channels** | Pointer to [**[]ChannelSpec**](ChannelSpec.md) | Channels are the fan-out targets, at most one per kind (paid, organic, email) and at most 12. A channel&#39;s status and provider id are server-owned: whatever the caller sends for them is replaced with \&quot;pending\&quot;. | [optional] 
**Content** | Pointer to **[]string** | Content is the ordered creative set. Content[0] is the active creative and the rest are A/B variants; at most 32, empty entries dropped. | [optional] 
**Name** | Pointer to **string** | Name is the campaign&#39;s display name. Required; trimmed and capped at 2048 characters. | [optional] 
**ScheduleAt** | Pointer to **int64** | ScheduleAt is when the campaign should run, in unix seconds. Negative reads as 0 (immediately). | [optional] 

## Methods

### NewCampaignWrite

`func NewCampaignWrite() *CampaignWrite`

NewCampaignWrite instantiates a new CampaignWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignWriteWithDefaults

`func NewCampaignWriteWithDefaults() *CampaignWrite`

NewCampaignWriteWithDefaults instantiates a new CampaignWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *CampaignWrite) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CampaignWrite) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CampaignWrite) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *CampaignWrite) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetBudget

`func (o *CampaignWrite) GetBudget() int64`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *CampaignWrite) GetBudgetOk() (*int64, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *CampaignWrite) SetBudget(v int64)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *CampaignWrite) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetChannels

`func (o *CampaignWrite) GetChannels() []ChannelSpec`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *CampaignWrite) GetChannelsOk() (*[]ChannelSpec, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *CampaignWrite) SetChannels(v []ChannelSpec)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *CampaignWrite) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetContent

`func (o *CampaignWrite) GetContent() []string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CampaignWrite) GetContentOk() (*[]string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CampaignWrite) SetContent(v []string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CampaignWrite) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetName

`func (o *CampaignWrite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignWrite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignWrite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignWrite) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduleAt

`func (o *CampaignWrite) GetScheduleAt() int64`

GetScheduleAt returns the ScheduleAt field if non-nil, zero value otherwise.

### GetScheduleAtOk

`func (o *CampaignWrite) GetScheduleAtOk() (*int64, bool)`

GetScheduleAtOk returns a tuple with the ScheduleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleAt

`func (o *CampaignWrite) SetScheduleAt(v int64)`

SetScheduleAt sets ScheduleAt field to given value.

### HasScheduleAt

`func (o *CampaignWrite) HasScheduleAt() bool`

HasScheduleAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


