# BotIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Batch** | Pointer to **string** |  | [optional] 
**StatsDownloads** | Pointer to **int32** |  | [optional] 
**StatsStars** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewBotIntegration

`func NewBotIntegration() *BotIntegration`

NewBotIntegration instantiates a new BotIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotIntegrationWithDefaults

`func NewBotIntegrationWithDefaults() *BotIntegration`

NewBotIntegrationWithDefaults instantiates a new BotIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BotIntegration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BotIntegration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BotIntegration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BotIntegration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSlug

`func (o *BotIntegration) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BotIntegration) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BotIntegration) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *BotIntegration) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDisplayName

`func (o *BotIntegration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BotIntegration) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BotIntegration) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *BotIntegration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSummary

`func (o *BotIntegration) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *BotIntegration) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *BotIntegration) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *BotIntegration) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetBatch

`func (o *BotIntegration) GetBatch() string`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *BotIntegration) GetBatchOk() (*string, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *BotIntegration) SetBatch(v string)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *BotIntegration) HasBatch() bool`

HasBatch returns a boolean if a field has been set.

### GetStatsDownloads

`func (o *BotIntegration) GetStatsDownloads() int32`

GetStatsDownloads returns the StatsDownloads field if non-nil, zero value otherwise.

### GetStatsDownloadsOk

`func (o *BotIntegration) GetStatsDownloadsOk() (*int32, bool)`

GetStatsDownloadsOk returns a tuple with the StatsDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsDownloads

`func (o *BotIntegration) SetStatsDownloads(v int32)`

SetStatsDownloads sets StatsDownloads field to given value.

### HasStatsDownloads

`func (o *BotIntegration) HasStatsDownloads() bool`

HasStatsDownloads returns a boolean if a field has been set.

### GetStatsStars

`func (o *BotIntegration) GetStatsStars() int32`

GetStatsStars returns the StatsStars field if non-nil, zero value otherwise.

### GetStatsStarsOk

`func (o *BotIntegration) GetStatsStarsOk() (*int32, bool)`

GetStatsStarsOk returns a tuple with the StatsStars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatsStars

`func (o *BotIntegration) SetStatsStars(v int32)`

SetStatsStars sets StatsStars field to given value.

### HasStatsStars

`func (o *BotIntegration) HasStatsStars() bool`

HasStatsStars returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BotIntegration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BotIntegration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BotIntegration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BotIntegration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *BotIntegration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BotIntegration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BotIntegration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BotIntegration) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


