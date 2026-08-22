# SocialPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | AccountID / ExternalID / Error are server-managed publish results, set only by the publish path (never by a client update): the account a post was published through, the provider&#39;s returned external post id (for reconciliation), and the last failure reason. Empty until a publish attempt lands. AccountID is the connected account the post went out through. Absent until a publish succeeds. | [optional] 
**Channel** | Pointer to **string** | Channel is the network this post targets: x, facebook, instagram, linkedin, tiktok, youtube or threads. It is the channel whose CONNECTED accounts a publish fans out to.  Example: \&quot;x\&quot; | [optional] 
**Content** | Pointer to **string** | Content is the post&#39;s text, bounded at 8192 characters.  Example: \&quot;Shipping today.\&quot; | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the post was created, as a unix timestamp in seconds. | [optional] 
**Error** | Pointer to **string** | Error is why the last publish attempt failed, verbatim and bounded. Absent when no attempt has failed; cleared by a later success. | [optional] 
**ExternalId** | Pointer to **string** | ExternalID is the id the network returned for the published post, which is what reconciles this row against the post on the network. Absent until a publish succeeds. | [optional] 
**Id** | Pointer to **string** | ID is the post&#39;s identifier, minted on create and the id every later call addresses it by.  Example: \&quot;post_91ab20\&quot; | [optional] 
**Media** | Pointer to **[]string** | Media is the post&#39;s attached media as a list of URLs (images today; the composer&#39;s URL field now, an S3 picker later, populate it). Stored as a JSON array in the media TEXT column and ALWAYS serialized as an array (never null), so a client can rely on &#x60;media&#x60; being present. Bounded at the write layer (normMedia in social.go): each URL clipped to maxField, the list to maxMedia. | [optional] 
**ScheduleAt** | Pointer to **int32** | ScheduleAt is when the post is due, as a unix timestamp in SECONDS. 0 means unscheduled. It is meaningful only while the status is scheduled — a scheduled post whose time has arrived is published by the scheduler. | [optional] 
**Status** | Pointer to **string** | Status is the post&#39;s lifecycle state: draft, scheduled, published or failed. A fifth, transient publishing state exists while a publish attempt holds the claim; it is never settable from a request and a caller sees it only if it reads a post mid-attempt. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is when the post row last changed, as a unix timestamp in seconds. The listing is ordered by it, newest first. | [optional] 

## Methods

### NewSocialPost

`func NewSocialPost() *SocialPost`

NewSocialPost instantiates a new SocialPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialPostWithDefaults

`func NewSocialPostWithDefaults() *SocialPost`

NewSocialPostWithDefaults instantiates a new SocialPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SocialPost) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SocialPost) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SocialPost) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *SocialPost) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetChannel

`func (o *SocialPost) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *SocialPost) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *SocialPost) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *SocialPost) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetContent

`func (o *SocialPost) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SocialPost) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SocialPost) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SocialPost) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SocialPost) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SocialPost) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SocialPost) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SocialPost) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetError

`func (o *SocialPost) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SocialPost) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SocialPost) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SocialPost) HasError() bool`

HasError returns a boolean if a field has been set.

### GetExternalId

`func (o *SocialPost) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SocialPost) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SocialPost) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SocialPost) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *SocialPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialPost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialPost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMedia

`func (o *SocialPost) GetMedia() []string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *SocialPost) GetMediaOk() (*[]string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *SocialPost) SetMedia(v []string)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *SocialPost) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetScheduleAt

`func (o *SocialPost) GetScheduleAt() int32`

GetScheduleAt returns the ScheduleAt field if non-nil, zero value otherwise.

### GetScheduleAtOk

`func (o *SocialPost) GetScheduleAtOk() (*int32, bool)`

GetScheduleAtOk returns a tuple with the ScheduleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleAt

`func (o *SocialPost) SetScheduleAt(v int32)`

SetScheduleAt sets ScheduleAt field to given value.

### HasScheduleAt

`func (o *SocialPost) HasScheduleAt() bool`

HasScheduleAt returns a boolean if a field has been set.

### GetStatus

`func (o *SocialPost) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SocialPost) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SocialPost) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SocialPost) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SocialPost) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SocialPost) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SocialPost) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SocialPost) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


