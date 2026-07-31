# CloudCalendarPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Body is the post text. Required. | [optional] 
**Channel** | Pointer to **string** | Channel is the target network: x, facebook, instagram, linkedin, tiktok, youtube or threads. Required — a post must name where it goes. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt and UpdatedAt are unix seconds, both server-assigned. | [optional] 
**Error** | Pointer to **string** | Error is the exact reason the last publish attempt failed — the honest record behind a \&quot;failed\&quot; status, never a faked success. | [optional] 
**Id** | Pointer to **string** | ID is the server-assigned post id (\&quot;cal_\&quot; + 128 random bits). | [optional] 
**PublishedAt** | Pointer to **int32** | PublishedAt is when the publish succeeded; 0 until it does. | [optional] 
**ScheduledAt** | Pointer to **int32** | ScheduledAt is the unix publish time; 0 leaves the post a draft, and any value makes it \&quot;scheduled\&quot; for the durable sweep to pick up. | [optional] 
**Status** | Pointer to **string** | Status is draft, scheduled, published, failed or canceled. Server-owned. | [optional] 
**Title** | Pointer to **string** | Title is the post&#39;s internal label, capped at 1024 bytes. | [optional] 
**UpdatedAt** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudCalendarPost

`func NewCloudCalendarPost() *CloudCalendarPost`

NewCloudCalendarPost instantiates a new CloudCalendarPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCalendarPostWithDefaults

`func NewCloudCalendarPostWithDefaults() *CloudCalendarPost`

NewCloudCalendarPostWithDefaults instantiates a new CloudCalendarPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *CloudCalendarPost) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CloudCalendarPost) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CloudCalendarPost) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *CloudCalendarPost) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetChannel

`func (o *CloudCalendarPost) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CloudCalendarPost) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CloudCalendarPost) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CloudCalendarPost) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudCalendarPost) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudCalendarPost) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudCalendarPost) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudCalendarPost) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetError

`func (o *CloudCalendarPost) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CloudCalendarPost) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CloudCalendarPost) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CloudCalendarPost) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *CloudCalendarPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudCalendarPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudCalendarPost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudCalendarPost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPublishedAt

`func (o *CloudCalendarPost) GetPublishedAt() int32`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *CloudCalendarPost) GetPublishedAtOk() (*int32, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *CloudCalendarPost) SetPublishedAt(v int32)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *CloudCalendarPost) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetScheduledAt

`func (o *CloudCalendarPost) GetScheduledAt() int32`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *CloudCalendarPost) GetScheduledAtOk() (*int32, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *CloudCalendarPost) SetScheduledAt(v int32)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *CloudCalendarPost) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetStatus

`func (o *CloudCalendarPost) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudCalendarPost) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudCalendarPost) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudCalendarPost) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *CloudCalendarPost) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudCalendarPost) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudCalendarPost) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudCalendarPost) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudCalendarPost) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudCalendarPost) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudCalendarPost) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudCalendarPost) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


