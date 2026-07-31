# CloudCampaignRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **string** |  | [optional] 
**Budget** | Pointer to **int32** |  | [optional] 
**Channels** | Pointer to [**[]CloudChannelSpec**](CloudChannelSpec.md) |  | [optional] 
**Content** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ScheduleAt** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudCampaignRecord

`func NewCloudCampaignRecord() *CloudCampaignRecord`

NewCloudCampaignRecord instantiates a new CloudCampaignRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCampaignRecordWithDefaults

`func NewCloudCampaignRecordWithDefaults() *CloudCampaignRecord`

NewCloudCampaignRecordWithDefaults instantiates a new CloudCampaignRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *CloudCampaignRecord) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CloudCampaignRecord) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CloudCampaignRecord) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *CloudCampaignRecord) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetBudget

`func (o *CloudCampaignRecord) GetBudget() int32`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *CloudCampaignRecord) GetBudgetOk() (*int32, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *CloudCampaignRecord) SetBudget(v int32)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *CloudCampaignRecord) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetChannels

`func (o *CloudCampaignRecord) GetChannels() []CloudChannelSpec`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *CloudCampaignRecord) GetChannelsOk() (*[]CloudChannelSpec, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *CloudCampaignRecord) SetChannels(v []CloudChannelSpec)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *CloudCampaignRecord) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetContent

`func (o *CloudCampaignRecord) GetContent() []string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CloudCampaignRecord) GetContentOk() (*[]string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CloudCampaignRecord) SetContent(v []string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CloudCampaignRecord) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudCampaignRecord) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudCampaignRecord) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudCampaignRecord) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudCampaignRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CloudCampaignRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCampaignRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCampaignRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCampaignRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CloudCampaignRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudCampaignRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudCampaignRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudCampaignRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduleAt

`func (o *CloudCampaignRecord) GetScheduleAt() int32`

GetScheduleAt returns the ScheduleAt field if non-nil, zero value otherwise.

### GetScheduleAtOk

`func (o *CloudCampaignRecord) GetScheduleAtOk() (*int32, bool)`

GetScheduleAtOk returns a tuple with the ScheduleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleAt

`func (o *CloudCampaignRecord) SetScheduleAt(v int32)`

SetScheduleAt sets ScheduleAt field to given value.

### HasScheduleAt

`func (o *CloudCampaignRecord) HasScheduleAt() bool`

HasScheduleAt returns a boolean if a field has been set.

### GetStatus

`func (o *CloudCampaignRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudCampaignRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudCampaignRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudCampaignRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudCampaignRecord) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudCampaignRecord) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudCampaignRecord) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudCampaignRecord) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


