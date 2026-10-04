# CampaignCampaignUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **string** |  | [optional] 
**Budget** | Pointer to **int64** |  | [optional] 
**Channels** | Pointer to [**[]CampaignChannelSpec**](CampaignChannelSpec.md) |  | [optional] 
**Content** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** | ID is the campaign to update, from the path. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ScheduleAt** | Pointer to **int64** |  | [optional] 

## Methods

### NewCampaignCampaignUpdate

`func NewCampaignCampaignUpdate() *CampaignCampaignUpdate`

NewCampaignCampaignUpdate instantiates a new CampaignCampaignUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignCampaignUpdateWithDefaults

`func NewCampaignCampaignUpdateWithDefaults() *CampaignCampaignUpdate`

NewCampaignCampaignUpdateWithDefaults instantiates a new CampaignCampaignUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *CampaignCampaignUpdate) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CampaignCampaignUpdate) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CampaignCampaignUpdate) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *CampaignCampaignUpdate) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetBudget

`func (o *CampaignCampaignUpdate) GetBudget() int64`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *CampaignCampaignUpdate) GetBudgetOk() (*int64, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *CampaignCampaignUpdate) SetBudget(v int64)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *CampaignCampaignUpdate) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetChannels

`func (o *CampaignCampaignUpdate) GetChannels() []CampaignChannelSpec`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *CampaignCampaignUpdate) GetChannelsOk() (*[]CampaignChannelSpec, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *CampaignCampaignUpdate) SetChannels(v []CampaignChannelSpec)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *CampaignCampaignUpdate) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetContent

`func (o *CampaignCampaignUpdate) GetContent() []string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CampaignCampaignUpdate) GetContentOk() (*[]string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CampaignCampaignUpdate) SetContent(v []string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CampaignCampaignUpdate) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetId

`func (o *CampaignCampaignUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignCampaignUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignCampaignUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CampaignCampaignUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CampaignCampaignUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignCampaignUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignCampaignUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignCampaignUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduleAt

`func (o *CampaignCampaignUpdate) GetScheduleAt() int64`

GetScheduleAt returns the ScheduleAt field if non-nil, zero value otherwise.

### GetScheduleAtOk

`func (o *CampaignCampaignUpdate) GetScheduleAtOk() (*int64, bool)`

GetScheduleAtOk returns a tuple with the ScheduleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleAt

`func (o *CampaignCampaignUpdate) SetScheduleAt(v int64)`

SetScheduleAt sets ScheduleAt field to given value.

### HasScheduleAt

`func (o *CampaignCampaignUpdate) HasScheduleAt() bool`

HasScheduleAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


