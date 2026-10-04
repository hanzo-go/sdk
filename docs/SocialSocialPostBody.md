# SocialSocialPostBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Channel is the network to publish to: x, facebook, instagram, linkedin, tiktok, youtube or threads. Omitted means x.  Example: \&quot;x\&quot; | [optional] 
**Content** | Pointer to **string** | Content is the post&#39;s text. Required — an empty body is a 400 — and bounded at 8192 characters, comfortably above every network&#39;s own limit.  Example: \&quot;Shipping today.\&quot; | [optional] 
**Media** | Pointer to **[]string** | Media is the post&#39;s attached media as URLs, at most 10, each bounded at 1024 characters. Blank entries are dropped. Omitting it CLEARS any stored media. | [optional] 
**ScheduleAt** | Pointer to **int64** | ScheduleAt is when to publish, as a unix timestamp in SECONDS. 0 means unscheduled. A negative value is clamped to 0. It only matters for a post whose status is scheduled. | [optional] 
**Status** | Pointer to **string** | Status is the post&#39;s lifecycle state: draft, scheduled, published or failed. Omitted means draft. The transient publishing claim is never settable here — accepting it from a request would let a caller wedge or replay the guard that stops two publishers double-posting the same row. | [optional] 

## Methods

### NewSocialSocialPostBody

`func NewSocialSocialPostBody() *SocialSocialPostBody`

NewSocialSocialPostBody instantiates a new SocialSocialPostBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialSocialPostBodyWithDefaults

`func NewSocialSocialPostBodyWithDefaults() *SocialSocialPostBody`

NewSocialSocialPostBodyWithDefaults instantiates a new SocialSocialPostBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *SocialSocialPostBody) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *SocialSocialPostBody) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *SocialSocialPostBody) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *SocialSocialPostBody) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetContent

`func (o *SocialSocialPostBody) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SocialSocialPostBody) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SocialSocialPostBody) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SocialSocialPostBody) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetMedia

`func (o *SocialSocialPostBody) GetMedia() []string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *SocialSocialPostBody) GetMediaOk() (*[]string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *SocialSocialPostBody) SetMedia(v []string)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *SocialSocialPostBody) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetScheduleAt

`func (o *SocialSocialPostBody) GetScheduleAt() int64`

GetScheduleAt returns the ScheduleAt field if non-nil, zero value otherwise.

### GetScheduleAtOk

`func (o *SocialSocialPostBody) GetScheduleAtOk() (*int64, bool)`

GetScheduleAtOk returns a tuple with the ScheduleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleAt

`func (o *SocialSocialPostBody) SetScheduleAt(v int64)`

SetScheduleAt sets ScheduleAt field to given value.

### HasScheduleAt

`func (o *SocialSocialPostBody) HasScheduleAt() bool`

HasScheduleAt returns a boolean if a field has been set.

### GetStatus

`func (o *SocialSocialPostBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SocialSocialPostBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SocialSocialPostBody) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SocialSocialPostBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


