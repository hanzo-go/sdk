# BotGetPersona200ResponsePersona

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
**LatestVersion** | Pointer to [**BotPersonaVersion**](BotPersonaVersion.md) |  | [optional] 
**Owner** | Pointer to [**BotGetPersona200ResponsePersonaAllOfOwner**](BotGetPersona200ResponsePersonaAllOfOwner.md) |  | [optional] 

## Methods

### NewBotGetPersona200ResponsePersona

`func NewBotGetPersona200ResponsePersona() *BotGetPersona200ResponsePersona`

NewBotGetPersona200ResponsePersona instantiates a new BotGetPersona200ResponsePersona object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotGetPersona200ResponsePersonaWithDefaults

`func NewBotGetPersona200ResponsePersonaWithDefaults() *BotGetPersona200ResponsePersona`

NewBotGetPersona200ResponsePersonaWithDefaults instantiates a new BotGetPersona200ResponsePersona object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BotGetPersona200ResponsePersona) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BotGetPersona200ResponsePersona) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BotGetPersona200ResponsePersona) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BotGetPersona200ResponsePersona) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSlug

`func (o *BotGetPersona200ResponsePersona) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BotGetPersona200ResponsePersona) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BotGetPersona200ResponsePersona) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *BotGetPersona200ResponsePersona) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDisplayName

`func (o *BotGetPersona200ResponsePersona) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BotGetPersona200ResponsePersona) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BotGetPersona200ResponsePersona) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BotGetPersona200ResponsePersona) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSummary

`func (o *BotGetPersona200ResponsePersona) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *BotGetPersona200ResponsePersona) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *BotGetPersona200ResponsePersona) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *BotGetPersona200ResponsePersona) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetOwnerUserId

`func (o *BotGetPersona200ResponsePersona) GetOwnerUserId() string`

GetOwnerUserId returns the OwnerUserId field if non-nil, zero value otherwise.

### GetOwnerUserIdOk

`func (o *BotGetPersona200ResponsePersona) GetOwnerUserIdOk() (*string, bool)`

GetOwnerUserIdOk returns a tuple with the OwnerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserId

`func (o *BotGetPersona200ResponsePersona) SetOwnerUserId(v string)`

SetOwnerUserId sets OwnerUserId field to given value.

### HasOwnerUserId

`func (o *BotGetPersona200ResponsePersona) HasOwnerUserId() bool`

HasOwnerUserId returns a boolean if a field has been set.

### GetLatestVersionId

`func (o *BotGetPersona200ResponsePersona) GetLatestVersionId() string`

GetLatestVersionId returns the LatestVersionId field if non-nil, zero value otherwise.

### GetLatestVersionIdOk

`func (o *BotGetPersona200ResponsePersona) GetLatestVersionIdOk() (*string, bool)`

GetLatestVersionIdOk returns a tuple with the LatestVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersionId

`func (o *BotGetPersona200ResponsePersona) SetLatestVersionId(v string)`

SetLatestVersionId sets LatestVersionId field to given value.

### HasLatestVersionId

`func (o *BotGetPersona200ResponsePersona) HasLatestVersionId() bool`

HasLatestVersionId returns a boolean if a field has been set.

### GetStatsDownloads

`func (o *BotGetPersona200ResponsePersona) GetStatsDownloads() int32`

GetStatsDownloads returns the StatsDownloads field if non-nil, zero value otherwise.

### GetStatsDownloadsOk

`func (o *BotGetPersona200ResponsePersona) GetStatsDownloadsOk() (*int32, bool)`

GetStatsDownloadsOk returns a tuple with the StatsDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsDownloads

`func (o *BotGetPersona200ResponsePersona) SetStatsDownloads(v int32)`

SetStatsDownloads sets StatsDownloads field to given value.

### HasStatsDownloads

`func (o *BotGetPersona200ResponsePersona) HasStatsDownloads() bool`

HasStatsDownloads returns a boolean if a field has been set.

### GetStatsStars

`func (o *BotGetPersona200ResponsePersona) GetStatsStars() int32`

GetStatsStars returns the StatsStars field if non-nil, zero value otherwise.

### GetStatsStarsOk

`func (o *BotGetPersona200ResponsePersona) GetStatsStarsOk() (*int32, bool)`

GetStatsStarsOk returns a tuple with the StatsStars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsStars

`func (o *BotGetPersona200ResponsePersona) SetStatsStars(v int32)`

SetStatsStars sets StatsStars field to given value.

### HasStatsStars

`func (o *BotGetPersona200ResponsePersona) HasStatsStars() bool`

HasStatsStars returns a boolean if a field has been set.

### GetStatsVersions

`func (o *BotGetPersona200ResponsePersona) GetStatsVersions() int32`

GetStatsVersions returns the StatsVersions field if non-nil, zero value otherwise.

### GetStatsVersionsOk

`func (o *BotGetPersona200ResponsePersona) GetStatsVersionsOk() (*int32, bool)`

GetStatsVersionsOk returns a tuple with the StatsVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsVersions

`func (o *BotGetPersona200ResponsePersona) SetStatsVersions(v int32)`

SetStatsVersions sets StatsVersions field to given value.

### HasStatsVersions

`func (o *BotGetPersona200ResponsePersona) HasStatsVersions() bool`

HasStatsVersions returns a boolean if a field has been set.

### GetStatsComments

`func (o *BotGetPersona200ResponsePersona) GetStatsComments() int32`

GetStatsComments returns the StatsComments field if non-nil, zero value otherwise.

### GetStatsCommentsOk

`func (o *BotGetPersona200ResponsePersona) GetStatsCommentsOk() (*int32, bool)`

GetStatsCommentsOk returns a tuple with the StatsComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsComments

`func (o *BotGetPersona200ResponsePersona) SetStatsComments(v int32)`

SetStatsComments sets StatsComments field to given value.

### HasStatsComments

`func (o *BotGetPersona200ResponsePersona) HasStatsComments() bool`

HasStatsComments returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BotGetPersona200ResponsePersona) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BotGetPersona200ResponsePersona) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BotGetPersona200ResponsePersona) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BotGetPersona200ResponsePersona) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BotGetPersona200ResponsePersona) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BotGetPersona200ResponsePersona) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BotGetPersona200ResponsePersona) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BotGetPersona200ResponsePersona) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLatestVersion

`func (o *BotGetPersona200ResponsePersona) GetLatestVersion() BotPersonaVersion`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *BotGetPersona200ResponsePersona) GetLatestVersionOk() (*BotPersonaVersion, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *BotGetPersona200ResponsePersona) SetLatestVersion(v BotPersonaVersion)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *BotGetPersona200ResponsePersona) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetOwner

`func (o *BotGetPersona200ResponsePersona) GetOwner() BotGetPersona200ResponsePersonaAllOfOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BotGetPersona200ResponsePersona) GetOwnerOk() (*BotGetPersona200ResponsePersonaAllOfOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BotGetPersona200ResponsePersona) SetOwner(v BotGetPersona200ResponsePersonaAllOfOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BotGetPersona200ResponsePersona) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


