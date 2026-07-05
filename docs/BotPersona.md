# BotPersona

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**OwnerUserId** | Pointer to **string** |  | [optional] 
**LatestVersionId** | Pointer to **string** |  | [optional] 
**StatsDownloads** | Pointer to **int32** |  | [optional] 
**StatsStars** | Pointer to **int32** |  | [optional] 
**StatsVersions** | Pointer to **int32** |  | [optional] 
**StatsComments** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewBotPersona

`func NewBotPersona() *BotPersona`

NewBotPersona instantiates a new BotPersona object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotPersonaWithDefaults

`func NewBotPersonaWithDefaults() *BotPersona`

NewBotPersonaWithDefaults instantiates a new BotPersona object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BotPersona) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BotPersona) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BotPersona) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BotPersona) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSlug

`func (o *BotPersona) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BotPersona) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BotPersona) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *BotPersona) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDisplayName

`func (o *BotPersona) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BotPersona) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BotPersona) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BotPersona) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSummary

`func (o *BotPersona) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *BotPersona) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *BotPersona) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *BotPersona) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetOwnerUserId

`func (o *BotPersona) GetOwnerUserId() string`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *BotPersona) GetOwnerUserIdOk() (*string, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *BotPersona) SetOwnerUserId(v string)`

SetOwnerUserId sets OwnerUserId field to given value.

### HasOwnerUserId

`func (o *BotPersona) HasOwnerUserId() bool`

HasOwnerUserId returns a boolean if a field has been set.

### GetLatestVersionId

`func (o *BotPersona) GetLatestVersionId() string`

GetLatestVersionId returns the LatestVersionId field if non-nil, zero value otherwise.

### GetLatestVersionIdOk

`func (o *BotPersona) GetLatestVersionIdOk() (*string, bool)`

GetLatestVersionIdOk returns a tuple with the LatestVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersionId

`func (o *BotPersona) SetLatestVersionId(v string)`

SetLatestVersionId sets LatestVersionId field to given value.

### HasLatestVersionId

`func (o *BotPersona) HasLatestVersionId() bool`

HasLatestVersionId returns a boolean if a field has been set.

### GetStatsDownloads

`func (o *BotPersona) GetStatsDownloads() int32`

GetStatsDownloads returns the StatsDownloads field if non-nil, zero value otherwise.

### GetStatsDownloadsOk

`func (o *BotPersona) GetStatsDownloadsOk() (*int32, bool)`

GetStatsDownloadsOk returns a tuple with the StatsDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsDownloads

`func (o *BotPersona) SetStatsDownloads(v int32)`

SetStatsDownloads sets StatsDownloads field to given value.

### HasStatsDownloads

`func (o *BotPersona) HasStatsDownloads() bool`

HasStatsDownloads returns a boolean if a field has been set.

### GetStatsStars

`func (o *BotPersona) GetStatsStars() int32`

GetStatsStars returns the StatsStars field if non-nil, zero value otherwise.

### GetStatsStarsOk

`func (o *BotPersona) GetStatsStarsOk() (*int32, bool)`

GetStatsStarsOk returns a tuple with the StatsStars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsStars

`func (o *BotPersona) SetStatsStars(v int32)`

SetStatsStars sets StatsStars field to given value.

### HasStatsStars

`func (o *BotPersona) HasStatsStars() bool`

HasStatsStars returns a boolean if a field has been set.

### GetStatsVersions

`func (o *BotPersona) GetStatsVersions() int32`

GetStatsVersions returns the StatsVersions field if non-nil, zero value otherwise.

### GetStatsVersionsOk

`func (o *BotPersona) GetStatsVersionsOk() (*int32, bool)`

GetStatsVersionsOk returns a tuple with the StatsVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsVersions

`func (o *BotPersona) SetStatsVersions(v int32)`

SetStatsVersions sets StatsVersions field to given value.

### HasStatsVersions

`func (o *BotPersona) HasStatsVersions() bool`

HasStatsVersions returns a boolean if a field has been set.

### GetStatsComments

`func (o *BotPersona) GetStatsComments() int32`

GetStatsComments returns the StatsComments field if non-nil, zero value otherwise.

### GetStatsCommentsOk

`func (o *BotPersona) GetStatsCommentsOk() (*int32, bool)`

GetStatsCommentsOk returns a tuple with the StatsComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsComments

`func (o *BotPersona) SetStatsComments(v int32)`

SetStatsComments sets StatsComments field to given value.

### HasStatsComments

`func (o *BotPersona) HasStatsComments() bool`

HasStatsComments returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BotPersona) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BotPersona) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BotPersona) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BotPersona) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BotPersona) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BotPersona) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BotPersona) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BotPersona) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


