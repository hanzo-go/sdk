# SocialSocialPostWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is the network to publish to: x, facebook, instagram, linkedin, tiktok, youtube or threads. Omitted means x.  Example: \&quot;x\&quot; | [optional] 
**Content** | Pointer to **string** | Content is the post&#39;s text. Required on every update, and bounded at 8192 characters.  Example: \&quot;Shipping today.\&quot; | [optional] 
**Media** | Pointer to **[]string** | Media is the post&#39;s attached media as URLs, at most 10. Omitting it CLEARS any stored media: this is a replacement, not a merge. | [optional] 
**ScheduleAt** | Pointer to **int64** | ScheduleAt is when to publish, as a unix timestamp in SECONDS. 0 means unscheduled. Moving it into the past here does NOT publish the post — that is the scheduler&#39;s to notice, or the publish operation&#39;s. | [optional] 
**Status** | Pointer to **string** | Status is the post&#39;s lifecycle state: draft, scheduled, published or failed. Omitting it RESETS the post to draft. | [optional] 

## Methods

### NewSocialSocialPostWrite

`func NewSocialSocialPostWrite() *SocialSocialPostWrite`

NewSocialSocialPostWrite instantiates a new SocialSocialPostWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialSocialPostWriteWithDefaults

`func NewSocialSocialPostWriteWithDefaults() *SocialSocialPostWrite`

NewSocialSocialPostWriteWithDefaults instantiates a new SocialSocialPostWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *SocialSocialPostWrite) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *SocialSocialPostWrite) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *SocialSocialPostWrite) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *SocialSocialPostWrite) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetContent

`func (o *SocialSocialPostWrite) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SocialSocialPostWrite) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SocialSocialPostWrite) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SocialSocialPostWrite) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetMedia

`func (o *SocialSocialPostWrite) GetMedia() []string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *SocialSocialPostWrite) GetMediaOk() (*[]string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *SocialSocialPostWrite) SetMedia(v []string)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *SocialSocialPostWrite) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetScheduleAt

`func (o *SocialSocialPostWrite) GetScheduleAt() int64`

GetScheduleAt returns the ScheduleAt field if non-nil, zero value otherwise.

### GetScheduleAtOk

`func (o *SocialSocialPostWrite) GetScheduleAtOk() (*int64, bool)`

GetScheduleAtOk returns a tuple with the ScheduleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleAt

`func (o *SocialSocialPostWrite) SetScheduleAt(v int64)`

SetScheduleAt sets ScheduleAt field to given value.

### HasScheduleAt

`func (o *SocialSocialPostWrite) HasScheduleAt() bool`

HasScheduleAt returns a boolean if a field has been set.

### GetStatus

`func (o *SocialSocialPostWrite) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SocialSocialPostWrite) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SocialSocialPostWrite) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SocialSocialPostWrite) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


