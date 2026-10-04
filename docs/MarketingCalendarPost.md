# MarketingCalendarPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Body is the post text. Required. | [optional] 
**Channel** | Pointer to **string** | Channel is the target network: x, facebook, instagram, linkedin, tiktok, youtube or threads. Required — a post must name where it goes. | [optional] 
**CreatedAt** | Pointer to **int64** | CreatedAt is unix seconds when the post was added, server-assigned and never rewritten. | [optional] 
**Error** | Pointer to **string** | Error is the exact reason the last publish attempt failed — the honest record behind a \&quot;failed\&quot; status, never a faked success. | [optional] 
**Id** | Pointer to **string** | ID is the server-assigned post id (\&quot;cal_\&quot; + 128 random bits). | [optional] 
**PublishedAt** | Pointer to **int64** | PublishedAt is when the publish succeeded; 0 until it does. | [optional] 
**ScheduledAt** | Pointer to **int64** | ScheduledAt is the unix publish time; 0 leaves the post a draft, and any value makes it \&quot;scheduled\&quot; for the durable sweep to pick up. | [optional] 
**Status** | Pointer to **string** | Status is draft, scheduled, published, failed or canceled. Server-owned. | [optional] 
**Title** | Pointer to **string** | Title is the post&#39;s internal label, capped at 1024 bytes. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is unix seconds of the last write, server-assigned. The durable sweep writes too — claiming a due post, publishing it and recording a failure each bump it — so this moves without anyone editing the post. | [optional] 

## Methods

### NewMarketingCalendarPost

`func NewMarketingCalendarPost() *MarketingCalendarPost`

NewMarketingCalendarPost instantiates a new MarketingCalendarPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingCalendarPostWithDefaults

`func NewMarketingCalendarPostWithDefaults() *MarketingCalendarPost`

NewMarketingCalendarPostWithDefaults instantiates a new MarketingCalendarPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *MarketingCalendarPost) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *MarketingCalendarPost) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *MarketingCalendarPost) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *MarketingCalendarPost) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetChannel

`func (o *MarketingCalendarPost) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *MarketingCalendarPost) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *MarketingCalendarPost) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *MarketingCalendarPost) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MarketingCalendarPost) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MarketingCalendarPost) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MarketingCalendarPost) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MarketingCalendarPost) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetError

`func (o *MarketingCalendarPost) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MarketingCalendarPost) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MarketingCalendarPost) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *MarketingCalendarPost) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *MarketingCalendarPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingCalendarPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingCalendarPost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingCalendarPost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPublishedAt

`func (o *MarketingCalendarPost) GetPublishedAt() int64`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *MarketingCalendarPost) GetPublishedAtOk() (*int64, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *MarketingCalendarPost) SetPublishedAt(v int64)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *MarketingCalendarPost) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetScheduledAt

`func (o *MarketingCalendarPost) GetScheduledAt() int64`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *MarketingCalendarPost) GetScheduledAtOk() (*int64, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *MarketingCalendarPost) SetScheduledAt(v int64)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *MarketingCalendarPost) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetStatus

`func (o *MarketingCalendarPost) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MarketingCalendarPost) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MarketingCalendarPost) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MarketingCalendarPost) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *MarketingCalendarPost) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MarketingCalendarPost) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MarketingCalendarPost) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MarketingCalendarPost) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MarketingCalendarPost) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MarketingCalendarPost) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MarketingCalendarPost) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MarketingCalendarPost) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


