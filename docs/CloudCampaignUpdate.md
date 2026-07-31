# CloudCampaignUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **string** |  | [optional] 
**Budget** | Pointer to **int32** |  | [optional] 
**Channels** | Pointer to [**[]CloudChannelSpec**](CloudChannelSpec.md) |  | [optional] 
**Content** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** | ID is the campaign to update, from the path. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ScheduleAt** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudCampaignUpdate

`func NewCloudCampaignUpdate() *CloudCampaignUpdate`

NewCloudCampaignUpdate instantiates a new CloudCampaignUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCampaignUpdateWithDefaults

`func NewCloudCampaignUpdateWithDefaults() *CloudCampaignUpdate`

NewCloudCampaignUpdateWithDefaults instantiates a new CloudCampaignUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *CloudCampaignUpdate) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CloudCampaignUpdate) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CloudCampaignUpdate) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *CloudCampaignUpdate) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetBudget

`func (o *CloudCampaignUpdate) GetBudget() int32`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *CloudCampaignUpdate) GetBudgetOk() (*int32, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *CloudCampaignUpdate) SetBudget(v int32)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *CloudCampaignUpdate) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetChannels

`func (o *CloudCampaignUpdate) GetChannels() []CloudChannelSpec`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *CloudCampaignUpdate) GetChannelsOk() (*[]CloudChannelSpec, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *CloudCampaignUpdate) SetChannels(v []CloudChannelSpec)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *CloudCampaignUpdate) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetContent

`func (o *CloudCampaignUpdate) GetContent() []string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CloudCampaignUpdate) GetContentOk() (*[]string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CloudCampaignUpdate) SetContent(v []string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CloudCampaignUpdate) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetId

`func (o *CloudCampaignUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCampaignUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCampaignUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCampaignUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudCampaignUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCampaignUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCampaignUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCampaignUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduleAt

`func (o *CloudCampaignUpdate) GetScheduleAt() int32`

GetScheduleAt returns the ScheduleAt field if non-nil, zero value otherwise.

### GetScheduleAtOk

`func (o *CloudCampaignUpdate) GetScheduleAtOk() (*int32, bool)`

GetScheduleAtOk returns a tuple with the ScheduleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleAt

`func (o *CloudCampaignUpdate) SetScheduleAt(v int32)`

SetScheduleAt sets ScheduleAt field to given value.

### HasScheduleAt

`func (o *CloudCampaignUpdate) HasScheduleAt() bool`

HasScheduleAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


