# BotSkill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** | URL-safe unique identifier | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**OwnerUserId** | Pointer to **string** |  | [optional] 
**ForkOf** | Pointer to **string** |  | [optional] 
**LatestVersionId** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**Badges** | Pointer to **map[string]interface{}** |  | [optional] 
**Batch** | Pointer to **string** | Grouping key (e.g. \&quot;integration\&quot;) | [optional] 
**ModerationStatus** | Pointer to **string** |  | [optional] 
**Quality** | Pointer to **map[string]interface{}** |  | [optional] 
**StatsDownloads** | Pointer to **int32** |  | [optional] 
**StatsStars** | Pointer to **int32** |  | [optional] 
**StatsVersions** | Pointer to **int32** |  | [optional] 
**StatsComments** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**OwnerHandle** | Pointer to **string** |  | [optional] 
**OwnerImage** | Pointer to **string** |  | [optional] 

## Methods

### NewBotSkill

`func NewBotSkill() *BotSkill`

NewBotSkill instantiates a new BotSkill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotSkillWithDefaults

`func NewBotSkillWithDefaults() *BotSkill`

NewBotSkillWithDefaults instantiates a new BotSkill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BotSkill) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BotSkill) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BotSkill) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BotSkill) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSlug

`func (o *BotSkill) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BotSkill) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BotSkill) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *BotSkill) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDisplayName

`func (o *BotSkill) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BotSkill) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BotSkill) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BotSkill) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSummary

`func (o *BotSkill) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *BotSkill) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *BotSkill) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *BotSkill) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetOwnerUserId

`func (o *BotSkill) GetOwnerUserId() string`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *BotSkill) GetOwnerUserIdOk() (*string, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *BotSkill) SetOwnerUserId(v string)`

SetOwnerUserId sets OwnerUserId field to given value.

### HasOwnerUserId

`func (o *BotSkill) HasOwnerUserId() bool`

HasOwnerUserId returns a boolean if a field has been set.

### GetForkOf

`func (o *BotSkill) GetForkOf() string`

GetForkOf returns the ForkOf field if non-nil, zero value otherwise.

### GetForkOfOk

`func (o *BotSkill) GetForkOfOk() (*string, bool)`

GetForkOfOk returns a tuple with the ForkOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkOf

`func (o *BotSkill) SetForkOf(v string)`

SetForkOf sets ForkOf field to given value.

### HasForkOf

`func (o *BotSkill) HasForkOf() bool`

HasForkOf returns a boolean if a field has been set.

### GetLatestVersionId

`func (o *BotSkill) GetLatestVersionId() string`

GetLatestVersionId returns the LatestVersionId field if non-nil, zero value otherwise.

### GetLatestVersionIdOk

`func (o *BotSkill) GetLatestVersionIdOk() (*string, bool)`

GetLatestVersionIdOk returns a tuple with the LatestVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersionId

`func (o *BotSkill) SetLatestVersionId(v string)`

SetLatestVersionId sets LatestVersionId field to given value.

### HasLatestVersionId

`func (o *BotSkill) HasLatestVersionId() bool`

HasLatestVersionId returns a boolean if a field has been set.

### GetTags

`func (o *BotSkill) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BotSkill) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BotSkill) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *BotSkill) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetBadges

`func (o *BotSkill) GetBadges() map[string]interface{}`

GetBadges returns the Badges field if non-nil, zero value otherwise.

### GetBadgesOk

`func (o *BotSkill) GetBadgesOk() (*map[string]interface{}, bool)`

GetBadgesOk returns a tuple with the Badges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadges

`func (o *BotSkill) SetBadges(v map[string]interface{})`

SetBadges sets Badges field to given value.

### HasBadges

`func (o *BotSkill) HasBadges() bool`

HasBadges returns a boolean if a field has been set.

### GetBatch

`func (o *BotSkill) GetBatch() string`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *BotSkill) GetBatchOk() (*string, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *BotSkill) SetBatch(v string)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *BotSkill) HasBatch() bool`

HasBatch returns a boolean if a field has been set.

### GetModerationStatus

`func (o *BotSkill) GetModerationStatus() string`

GetModerationStatus returns the ModerationStatus field if non-nil, zero value otherwise.

### GetModerationStatusOk

`func (o *BotSkill) GetModerationStatusOk() (*string, bool)`

GetModerationStatusOk returns a tuple with the ModerationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationStatus

`func (o *BotSkill) SetModerationStatus(v string)`

SetModerationStatus sets ModerationStatus field to given value.

### HasModerationStatus

`func (o *BotSkill) HasModerationStatus() bool`

HasModerationStatus returns a boolean if a field has been set.

### GetQuality

`func (o *BotSkill) GetQuality() map[string]interface{}`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *BotSkill) GetQualityOk() (*map[string]interface{}, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *BotSkill) SetQuality(v map[string]interface{})`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *BotSkill) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetStatsDownloads

`func (o *BotSkill) GetStatsDownloads() int32`

GetStatsDownloads returns the StatsDownloads field if non-nil, zero value otherwise.

### GetStatsDownloadsOk

`func (o *BotSkill) GetStatsDownloadsOk() (*int32, bool)`

GetStatsDownloadsOk returns a tuple with the StatsDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsDownloads

`func (o *BotSkill) SetStatsDownloads(v int32)`

SetStatsDownloads sets StatsDownloads field to given value.

### HasStatsDownloads

`func (o *BotSkill) HasStatsDownloads() bool`

HasStatsDownloads returns a boolean if a field has been set.

### GetStatsStars

`func (o *BotSkill) GetStatsStars() int32`

GetStatsStars returns the StatsStars field if non-nil, zero value otherwise.

### GetStatsStarsOk

`func (o *BotSkill) GetStatsStarsOk() (*int32, bool)`

GetStatsStarsOk returns a tuple with the StatsStars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsStars

`func (o *BotSkill) SetStatsStars(v int32)`

SetStatsStars sets StatsStars field to given value.

### HasStatsStars

`func (o *BotSkill) HasStatsStars() bool`

HasStatsStars returns a boolean if a field has been set.

### GetStatsVersions

`func (o *BotSkill) GetStatsVersions() int32`

GetStatsVersions returns the StatsVersions field if non-nil, zero value otherwise.

### GetStatsVersionsOk

`func (o *BotSkill) GetStatsVersionsOk() (*int32, bool)`

GetStatsVersionsOk returns a tuple with the StatsVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsVersions

`func (o *BotSkill) SetStatsVersions(v int32)`

SetStatsVersions sets StatsVersions field to given value.

### HasStatsVersions

`func (o *BotSkill) HasStatsVersions() bool`

HasStatsVersions returns a boolean if a field has been set.

### GetStatsComments

`func (o *BotSkill) GetStatsComments() int32`

GetStatsComments returns the StatsComments field if non-nil, zero value otherwise.

### GetStatsCommentsOk

`func (o *BotSkill) GetStatsCommentsOk() (*int32, bool)`

GetStatsCommentsOk returns a tuple with the StatsComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsComments

`func (o *BotSkill) SetStatsComments(v int32)`

SetStatsComments sets StatsComments field to given value.

### HasStatsComments

`func (o *BotSkill) HasStatsComments() bool`

HasStatsComments returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BotSkill) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BotSkill) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BotSkill) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BotSkill) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BotSkill) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BotSkill) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BotSkill) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BotSkill) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOwnerHandle

`func (o *BotSkill) GetOwnerHandle() string`

GetOwnerHandle returns the OwnerHandle field if non-nil, zero value otherwise.

### GetOwnerHandleOk

`func (o *BotSkill) GetOwnerHandleOk() (*string, bool)`

GetOwnerHandleOk returns a tuple with the OwnerHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerHandle

`func (o *BotSkill) SetOwnerHandle(v string)`

SetOwnerHandle sets OwnerHandle field to given value.

### HasOwnerHandle

`func (o *BotSkill) HasOwnerHandle() bool`

HasOwnerHandle returns a boolean if a field has been set.

### GetOwnerImage

`func (o *BotSkill) GetOwnerImage() string`

GetOwnerImage returns the OwnerImage field if non-nil, zero value otherwise.

### GetOwnerImageOk

`func (o *BotSkill) GetOwnerImageOk() (*string, bool)`

GetOwnerImageOk returns a tuple with the OwnerImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerImage

`func (o *BotSkill) SetOwnerImage(v string)`

SetOwnerImage sets OwnerImage field to given value.

### HasOwnerImage

`func (o *BotSkill) HasOwnerImage() bool`

HasOwnerImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


