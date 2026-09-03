# CampaignUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **string** |  | [optional] 
**Budget** | Pointer to **int64** |  | [optional] 
**Channels** | Pointer to [**[]ChannelSpec**](ChannelSpec.md) |  | [optional] 
**Content** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** | ID is the campaign to update, from the path. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ScheduleAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewCampaignUpdate

`func NewCampaignUpdate() *CampaignUpdate`

NewCampaignUpdate instantiates a new CampaignUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignUpdateWithDefaults

`func NewCampaignUpdateWithDefaults() *CampaignUpdate`

NewCampaignUpdateWithDefaults instantiates a new CampaignUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *CampaignUpdate) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CampaignUpdate) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CampaignUpdate) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *CampaignUpdate) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetBudget

`func (o *CampaignUpdate) GetBudget() int64`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *CampaignUpdate) GetBudgetOk() (*int64, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *CampaignUpdate) SetBudget(v int64)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *CampaignUpdate) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetChannels

`func (o *CampaignUpdate) GetChannels() []ChannelSpec`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *CampaignUpdate) GetChannelsOk() (*[]ChannelSpec, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *CampaignUpdate) SetChannels(v []ChannelSpec)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *CampaignUpdate) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetContent

`func (o *CampaignUpdate) GetContent() []string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CampaignUpdate) GetContentOk() (*[]string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CampaignUpdate) SetContent(v []string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CampaignUpdate) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetId

`func (o *CampaignUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CampaignUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CampaignUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduleAt

`func (o *CampaignUpdate) GetScheduleAt() int64`

GetScheduleAt returns the ScheduleAt field if non-nil, zero value otherwise.

### GetScheduleAtOk

`func (o *CampaignUpdate) GetScheduleAtOk() (*int64, bool)`

GetScheduleAtOk returns a tuple with the ScheduleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleAt

`func (o *CampaignUpdate) SetScheduleAt(v int64)`

SetScheduleAt sets ScheduleAt field to given value.

### HasScheduleAt

`func (o *CampaignUpdate) HasScheduleAt() bool`

HasScheduleAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


