# SocialPostWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is the network to publish to: x, facebook, instagram, linkedin, tiktok, youtube or threads. Omitted means x.  Example: \&quot;x\&quot; | [optional] 
**Content** | Pointer to **string** | Content is the post&#39;s text. Required on every update, and bounded at 8192 characters.  Example: \&quot;Shipping today.\&quot; | [optional] 
**Media** | Pointer to **[]string** | Media is the post&#39;s attached media as URLs, at most 10. Omitting it CLEARS any stored media: this is a replacement, not a merge. | [optional] 
**ScheduleAt** | Pointer to **int32** | ScheduleAt is when to publish, as a unix timestamp in SECONDS. 0 means unscheduled. Moving it into the past here does NOT publish the post — that is the scheduler&#39;s to notice, or the publish operation&#39;s. | [optional] 
**Status** | Pointer to **string** | Status is the post&#39;s lifecycle state: draft, scheduled, published or failed. Omitting it RESETS the post to draft. | [optional] 

## Methods

### NewSocialPostWrite

`func NewSocialPostWrite() *SocialPostWrite`

NewSocialPostWrite instantiates a new SocialPostWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialPostWriteWithDefaults

`func NewSocialPostWriteWithDefaults() *SocialPostWrite`

NewSocialPostWriteWithDefaults instantiates a new SocialPostWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *SocialPostWrite) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *SocialPostWrite) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *SocialPostWrite) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *SocialPostWrite) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetContent

`func (o *SocialPostWrite) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SocialPostWrite) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SocialPostWrite) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SocialPostWrite) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetMedia

`func (o *SocialPostWrite) GetMedia() []string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *SocialPostWrite) GetMediaOk() (*[]string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *SocialPostWrite) SetMedia(v []string)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *SocialPostWrite) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetScheduleAt

`func (o *SocialPostWrite) GetScheduleAt() int32`

GetScheduleAt returns the ScheduleAt field if non-nil, zero value otherwise.

### GetScheduleAtOk

`func (o *SocialPostWrite) GetScheduleAtOk() (*int32, bool)`

GetScheduleAtOk returns a tuple with the ScheduleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleAt

`func (o *SocialPostWrite) SetScheduleAt(v int32)`

SetScheduleAt sets ScheduleAt field to given value.

### HasScheduleAt

`func (o *SocialPostWrite) HasScheduleAt() bool`

HasScheduleAt returns a boolean if a field has been set.

### GetStatus

`func (o *SocialPostWrite) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SocialPostWrite) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SocialPostWrite) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SocialPostWrite) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


