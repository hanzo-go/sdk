# AiHistoryDay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByTask** | Pointer to **map[string]int32** |  | [optional] 
**CostSavedIndex** | Pointer to **float32** |  | [optional] 
**CumulativeCostSaved** | Pointer to **float32** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**Events** | Pointer to **int32** |  | [optional] 
**LearnedShare** | Pointer to **float32** |  | [optional] 
**RewardRate** | Pointer to **float32** |  | [optional] 
**RewardedEvents** | Pointer to **int32** |  | [optional] 

## Methods

### NewAiHistoryDay

`func NewAiHistoryDay() *AiHistoryDay`

NewAiHistoryDay instantiates a new AiHistoryDay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiHistoryDayWithDefaults

`func NewAiHistoryDayWithDefaults() *AiHistoryDay`

NewAiHistoryDayWithDefaults instantiates a new AiHistoryDay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByTask

`func (o *AiHistoryDay) GetByTask() map[string]int32`

GetByTask returns the ByTask field if non-nil, zero value otherwise.

### GetByTaskOk

`func (o *AiHistoryDay) GetByTaskOk() (*map[string]int32, bool)`

GetByTaskOk returns a tuple with the ByTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByTask

`func (o *AiHistoryDay) SetByTask(v map[string]int32)`

SetByTask sets ByTask field to given value.

### HasByTask

`func (o *AiHistoryDay) HasByTask() bool`

HasByTask returns a boolean if a field has been set.

### GetCostSavedIndex

`func (o *AiHistoryDay) GetCostSavedIndex() float32`

GetCostSavedIndex returns the CostSavedIndex field if non-nil, zero value otherwise.

### GetCostSavedIndexOk

`func (o *AiHistoryDay) GetCostSavedIndexOk() (*float32, bool)`

GetCostSavedIndexOk returns a tuple with the CostSavedIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostSavedIndex

`func (o *AiHistoryDay) SetCostSavedIndex(v float32)`

SetCostSavedIndex sets CostSavedIndex field to given value.

### HasCostSavedIndex

`func (o *AiHistoryDay) HasCostSavedIndex() bool`

HasCostSavedIndex returns a boolean if a field has been set.

### GetCumulativeCostSaved

`func (o *AiHistoryDay) GetCumulativeCostSaved() float32`

GetCumulativeCostSaved returns the CumulativeCostSaved field if non-nil, zero value otherwise.

### GetCumulativeCostSavedOk

`func (o *AiHistoryDay) GetCumulativeCostSavedOk() (*float32, bool)`

GetCumulativeCostSavedOk returns a tuple with the CumulativeCostSaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeCostSaved

`func (o *AiHistoryDay) SetCumulativeCostSaved(v float32)`

SetCumulativeCostSaved sets CumulativeCostSaved field to given value.

### HasCumulativeCostSaved

`func (o *AiHistoryDay) HasCumulativeCostSaved() bool`

HasCumulativeCostSaved returns a boolean if a field has been set.

### GetDate

`func (o *AiHistoryDay) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AiHistoryDay) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AiHistoryDay) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *AiHistoryDay) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetEvents

`func (o *AiHistoryDay) GetEvents() int32`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AiHistoryDay) GetEventsOk() (*int32, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AiHistoryDay) SetEvents(v int32)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *AiHistoryDay) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetLearnedShare

`func (o *AiHistoryDay) GetLearnedShare() float32`

GetLearnedShare returns the LearnedShare field if non-nil, zero value otherwise.

### GetLearnedShareOk

`func (o *AiHistoryDay) GetLearnedShareOk() (*float32, bool)`

GetLearnedShareOk returns a tuple with the LearnedShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearnedShare

`func (o *AiHistoryDay) SetLearnedShare(v float32)`

SetLearnedShare sets LearnedShare field to given value.

### HasLearnedShare

`func (o *AiHistoryDay) HasLearnedShare() bool`

HasLearnedShare returns a boolean if a field has been set.

### GetRewardRate

`func (o *AiHistoryDay) GetRewardRate() float32`

GetRewardRate returns the RewardRate field if non-nil, zero value otherwise.

### GetRewardRateOk

`func (o *AiHistoryDay) GetRewardRateOk() (*float32, bool)`

GetRewardRateOk returns a tuple with the RewardRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardRate

`func (o *AiHistoryDay) SetRewardRate(v float32)`

SetRewardRate sets RewardRate field to given value.

### HasRewardRate

`func (o *AiHistoryDay) HasRewardRate() bool`

HasRewardRate returns a boolean if a field has been set.

### GetRewardedEvents

`func (o *AiHistoryDay) GetRewardedEvents() int32`

GetRewardedEvents returns the RewardedEvents field if non-nil, zero value otherwise.

### GetRewardedEventsOk

`func (o *AiHistoryDay) GetRewardedEventsOk() (*int32, bool)`

GetRewardedEventsOk returns a tuple with the RewardedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardedEvents

`func (o *AiHistoryDay) SetRewardedEvents(v int32)`

SetRewardedEvents sets RewardedEvents field to given value.

### HasRewardedEvents

`func (o *AiHistoryDay) HasRewardedEvents() bool`

HasRewardedEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


