# SearchBatchView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | Pointer to **int32** |  | [optional] 
**Details** | Pointer to **map[string]interface{}** |  | [optional] 
**Stats** | Pointer to [**SearchBatchViewStats**](SearchBatchViewStats.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**Duration** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSearchBatchView

`func NewSearchBatchView() *SearchBatchView`

NewSearchBatchView instantiates a new SearchBatchView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBatchViewWithDefaults

`func NewSearchBatchViewWithDefaults() *SearchBatchView`

NewSearchBatchViewWithDefaults instantiates a new SearchBatchView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *SearchBatchView) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *SearchBatchView) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *SearchBatchView) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *SearchBatchView) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDetails

`func (o *SearchBatchView) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SearchBatchView) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SearchBatchView) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *SearchBatchView) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetStats

`func (o *SearchBatchView) GetStats() SearchBatchViewStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *SearchBatchView) GetStatsOk() (*SearchBatchViewStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *SearchBatchView) SetStats(v SearchBatchViewStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *SearchBatchView) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetStartedAt

`func (o *SearchBatchView) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *SearchBatchView) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *SearchBatchView) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *SearchBatchView) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *SearchBatchView) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *SearchBatchView) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *SearchBatchView) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *SearchBatchView) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetDuration

`func (o *SearchBatchView) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SearchBatchView) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SearchBatchView) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SearchBatchView) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetProgress

`func (o *SearchBatchView) GetProgress() map[string]interface{}`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *SearchBatchView) GetProgressOk() (*map[string]interface{}, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *SearchBatchView) SetProgress(v map[string]interface{})`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *SearchBatchView) HasProgress() bool`

HasProgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


