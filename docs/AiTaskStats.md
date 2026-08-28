# AiTaskStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to **int32** |  | [optional] 
**Models** | Pointer to **map[string]int32** |  | [optional] 

## Methods

### NewAiTaskStats

`func NewAiTaskStats() *AiTaskStats`

NewAiTaskStats instantiates a new AiTaskStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiTaskStatsWithDefaults

`func NewAiTaskStatsWithDefaults() *AiTaskStats`

NewAiTaskStatsWithDefaults instantiates a new AiTaskStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *AiTaskStats) GetEvents() int32`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AiTaskStats) GetEventsOk() (*int32, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AiTaskStats) SetEvents(v int32)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *AiTaskStats) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetModels

`func (o *AiTaskStats) GetModels() map[string]int32`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *AiTaskStats) GetModelsOk() (*map[string]int32, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *AiTaskStats) SetModels(v map[string]int32)`

SetModels sets Models field to given value.

### HasModels

`func (o *AiTaskStats) HasModels() bool`

HasModels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


