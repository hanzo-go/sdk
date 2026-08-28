# AiRouterStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByModel** | Pointer to **map[string]int32** |  | [optional] 
**BySource** | Pointer to **map[string]int32** |  | [optional] 
**ByTask** | Pointer to [**map[string]AiTaskStats**](AiTaskStats.md) |  | [optional] 
**Cost** | Pointer to [**AiCostStats**](AiCostStats.md) |  | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**Quality** | Pointer to [**AiQualityStats**](AiQualityStats.md) |  | [optional] 
**Retrain** | Pointer to [**AiRetrainMeta**](AiRetrainMeta.md) |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**Throughput** | Pointer to [**AiThroughputStats**](AiThroughputStats.md) |  | [optional] 
**Window** | Pointer to [**AiStatsWindow**](AiStatsWindow.md) |  | [optional] 

## Methods

### NewAiRouterStats

`func NewAiRouterStats() *AiRouterStats`

NewAiRouterStats instantiates a new AiRouterStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiRouterStatsWithDefaults

`func NewAiRouterStatsWithDefaults() *AiRouterStats`

NewAiRouterStatsWithDefaults instantiates a new AiRouterStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByModel

`func (o *AiRouterStats) GetByModel() map[string]int32`

GetByModel returns the ByModel field if non-nil, zero value otherwise.

### GetByModelOk

`func (o *AiRouterStats) GetByModelOk() (*map[string]int32, bool)`

GetByModelOk returns a tuple with the ByModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByModel

`func (o *AiRouterStats) SetByModel(v map[string]int32)`

SetByModel sets ByModel field to given value.

### HasByModel

`func (o *AiRouterStats) HasByModel() bool`

HasByModel returns a boolean if a field has been set.

### GetBySource

`func (o *AiRouterStats) GetBySource() map[string]int32`

GetBySource returns the BySource field if non-nil, zero value otherwise.

### GetBySourceOk

`func (o *AiRouterStats) GetBySourceOk() (*map[string]int32, bool)`

GetBySourceOk returns a tuple with the BySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySource

`func (o *AiRouterStats) SetBySource(v map[string]int32)`

SetBySource sets BySource field to given value.

### HasBySource

`func (o *AiRouterStats) HasBySource() bool`

HasBySource returns a boolean if a field has been set.

### GetByTask

`func (o *AiRouterStats) GetByTask() map[string]AiTaskStats`

GetByTask returns the ByTask field if non-nil, zero value otherwise.

### GetByTaskOk

`func (o *AiRouterStats) GetByTaskOk() (*map[string]AiTaskStats, bool)`

GetByTaskOk returns a tuple with the ByTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByTask

`func (o *AiRouterStats) SetByTask(v map[string]AiTaskStats)`

SetByTask sets ByTask field to given value.

### HasByTask

`func (o *AiRouterStats) HasByTask() bool`

HasByTask returns a boolean if a field has been set.

### GetCost

`func (o *AiRouterStats) GetCost() AiCostStats`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *AiRouterStats) GetCostOk() (*AiCostStats, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *AiRouterStats) SetCost(v AiCostStats)`

SetCost sets Cost field to given value.

### HasCost

`func (o *AiRouterStats) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetOrg

`func (o *AiRouterStats) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AiRouterStats) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AiRouterStats) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AiRouterStats) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetQuality

`func (o *AiRouterStats) GetQuality() AiQualityStats`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *AiRouterStats) GetQualityOk() (*AiQualityStats, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *AiRouterStats) SetQuality(v AiQualityStats)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *AiRouterStats) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetRetrain

`func (o *AiRouterStats) GetRetrain() AiRetrainMeta`

GetRetrain returns the Retrain field if non-nil, zero value otherwise.

### GetRetrainOk

`func (o *AiRouterStats) GetRetrainOk() (*AiRetrainMeta, bool)`

GetRetrainOk returns a tuple with the Retrain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrain

`func (o *AiRouterStats) SetRetrain(v AiRetrainMeta)`

SetRetrain sets Retrain field to given value.

### HasRetrain

`func (o *AiRouterStats) HasRetrain() bool`

HasRetrain returns a boolean if a field has been set.

### GetScope

`func (o *AiRouterStats) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AiRouterStats) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AiRouterStats) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AiRouterStats) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetThroughput

`func (o *AiRouterStats) GetThroughput() AiThroughputStats`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *AiRouterStats) GetThroughputOk() (*AiThroughputStats, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *AiRouterStats) SetThroughput(v AiThroughputStats)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *AiRouterStats) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.

### GetWindow

`func (o *AiRouterStats) GetWindow() AiStatsWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *AiRouterStats) GetWindowOk() (*AiStatsWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *AiRouterStats) SetWindow(v AiStatsWindow)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *AiRouterStats) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


