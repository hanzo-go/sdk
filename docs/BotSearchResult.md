# BotSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**OwnerHandle** | Pointer to **string** |  | [optional] 
**OwnerImage** | Pointer to **string** |  | [optional] 
**StatsDownloads** | Pointer to **int32** |  | [optional] 
**StatsStars** | Pointer to **int32** |  | [optional] 
**Score** | Pointer to **float32** | Relevance score (1.0 &#x3D; exact match, 0 &#x3D; lexical only) | [optional] 

## Methods

### NewBotSearchResult

`func NewBotSearchResult() *BotSearchResult`

NewBotSearchResult instantiates a new BotSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotSearchResultWithDefaults

`func NewBotSearchResultWithDefaults() *BotSearchResult`

NewBotSearchResultWithDefaults instantiates a new BotSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BotSearchResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BotSearchResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BotSearchResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BotSearchResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSlug

`func (o *BotSearchResult) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BotSearchResult) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BotSearchResult) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *BotSearchResult) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDisplayName

`func (o *BotSearchResult) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BotSearchResult) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BotSearchResult) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BotSearchResult) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSummary

`func (o *BotSearchResult) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *BotSearchResult) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *BotSearchResult) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *BotSearchResult) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetOwnerHandle

`func (o *BotSearchResult) GetOwnerHandle() string`

GetOwnerHandle returns the OwnerHandle field if non-nil, zero value otherwise.

### GetOwnerHandleOk

`func (o *BotSearchResult) GetOwnerHandleOk() (*string, bool)`

GetOwnerHandleOk returns a tuple with the OwnerHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerHandle

`func (o *BotSearchResult) SetOwnerHandle(v string)`

SetOwnerHandle sets OwnerHandle field to given value.

### HasOwnerHandle

`func (o *BotSearchResult) HasOwnerHandle() bool`

HasOwnerHandle returns a boolean if a field has been set.

### GetOwnerImage

`func (o *BotSearchResult) GetOwnerImage() string`

GetOwnerImage returns the OwnerImage field if non-nil, zero value otherwise.

### GetOwnerImageOk

`func (o *BotSearchResult) GetOwnerImageOk() (*string, bool)`

GetOwnerImageOk returns a tuple with the OwnerImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerImage

`func (o *BotSearchResult) SetOwnerImage(v string)`

SetOwnerImage sets OwnerImage field to given value.

### HasOwnerImage

`func (o *BotSearchResult) HasOwnerImage() bool`

HasOwnerImage returns a boolean if a field has been set.

### GetStatsDownloads

`func (o *BotSearchResult) GetStatsDownloads() int32`

GetStatsDownloads returns the StatsDownloads field if non-nil, zero value otherwise.

### GetStatsDownloadsOk

`func (o *BotSearchResult) GetStatsDownloadsOk() (*int32, bool)`

GetStatsDownloadsOk returns a tuple with the StatsDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsDownloads

`func (o *BotSearchResult) SetStatsDownloads(v int32)`

SetStatsDownloads sets StatsDownloads field to given value.

### HasStatsDownloads

`func (o *BotSearchResult) HasStatsDownloads() bool`

HasStatsDownloads returns a boolean if a field has been set.

### GetStatsStars

`func (o *BotSearchResult) GetStatsStars() int32`

GetStatsStars returns the StatsStars field if non-nil, zero value otherwise.

### GetStatsStarsOk

`func (o *BotSearchResult) GetStatsStarsOk() (*int32, bool)`

GetStatsStarsOk returns a tuple with the StatsStars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsStars

`func (o *BotSearchResult) SetStatsStars(v int32)`

SetStatsStars sets StatsStars field to given value.

### HasStatsStars

`func (o *BotSearchResult) HasStatsStars() bool`

HasStatsStars returns a boolean if a field has been set.

### GetScore

`func (o *BotSearchResult) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *BotSearchResult) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *BotSearchResult) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *BotSearchResult) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


