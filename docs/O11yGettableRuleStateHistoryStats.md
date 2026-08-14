# O11yGettableRuleStateHistoryStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentAvgResolutionTime** | Pointer to **float32** |  | [optional] 
**CurrentAvgResolutionTimeSeries** | Pointer to [**O11yTimeSeries**](O11yTimeSeries.md) |  | [optional] 
**CurrentTriggersSeries** | Pointer to [**O11yTimeSeries**](O11yTimeSeries.md) |  | [optional] 
**PastAvgResolutionTime** | Pointer to **float32** |  | [optional] 
**PastAvgResolutionTimeSeries** | Pointer to [**O11yTimeSeries**](O11yTimeSeries.md) |  | [optional] 
**PastTriggersSeries** | Pointer to [**O11yTimeSeries**](O11yTimeSeries.md) |  | [optional] 
**TotalCurrentTriggers** | Pointer to **int32** |  | [optional] 
**TotalPastTriggers** | Pointer to **int32** |  | [optional] 

## Methods

### NewO11yGettableRuleStateHistoryStats

`func NewO11yGettableRuleStateHistoryStats() *O11yGettableRuleStateHistoryStats`

NewO11yGettableRuleStateHistoryStats instantiates a new O11yGettableRuleStateHistoryStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yGettableRuleStateHistoryStatsWithDefaults

`func NewO11yGettableRuleStateHistoryStatsWithDefaults() *O11yGettableRuleStateHistoryStats`

NewO11yGettableRuleStateHistoryStatsWithDefaults instantiates a new O11yGettableRuleStateHistoryStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentAvgResolutionTime

`func (o *O11yGettableRuleStateHistoryStats) GetCurrentAvgResolutionTime() float32`

GetCurrentAvgResolutionTime returns the CurrentAvgResolutionTime field if non-nil, zero value otherwise.

### GetCurrentAvgResolutionTimeOk

`func (o *O11yGettableRuleStateHistoryStats) GetCurrentAvgResolutionTimeOk() (*float32, bool)`

GetCurrentAvgResolutionTimeOk returns a tuple with the CurrentAvgResolutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvgResolutionTime

`func (o *O11yGettableRuleStateHistoryStats) SetCurrentAvgResolutionTime(v float32)`

SetCurrentAvgResolutionTime sets CurrentAvgResolutionTime field to given value.

### HasCurrentAvgResolutionTime

`func (o *O11yGettableRuleStateHistoryStats) HasCurrentAvgResolutionTime() bool`

HasCurrentAvgResolutionTime returns a boolean if a field has been set.

### GetCurrentAvgResolutionTimeSeries

`func (o *O11yGettableRuleStateHistoryStats) GetCurrentAvgResolutionTimeSeries() O11yTimeSeries`

GetCurrentAvgResolutionTimeSeries returns the CurrentAvgResolutionTimeSeries field if non-nil, zero value otherwise.

### GetCurrentAvgResolutionTimeSeriesOk

`func (o *O11yGettableRuleStateHistoryStats) GetCurrentAvgResolutionTimeSeriesOk() (*O11yTimeSeries, bool)`

GetCurrentAvgResolutionTimeSeriesOk returns a tuple with the CurrentAvgResolutionTimeSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvgResolutionTimeSeries

`func (o *O11yGettableRuleStateHistoryStats) SetCurrentAvgResolutionTimeSeries(v O11yTimeSeries)`

SetCurrentAvgResolutionTimeSeries sets CurrentAvgResolutionTimeSeries field to given value.

### HasCurrentAvgResolutionTimeSeries

`func (o *O11yGettableRuleStateHistoryStats) HasCurrentAvgResolutionTimeSeries() bool`

HasCurrentAvgResolutionTimeSeries returns a boolean if a field has been set.

### GetCurrentTriggersSeries

`func (o *O11yGettableRuleStateHistoryStats) GetCurrentTriggersSeries() O11yTimeSeries`

GetCurrentTriggersSeries returns the CurrentTriggersSeries field if non-nil, zero value otherwise.

### GetCurrentTriggersSeriesOk

`func (o *O11yGettableRuleStateHistoryStats) GetCurrentTriggersSeriesOk() (*O11yTimeSeries, bool)`

GetCurrentTriggersSeriesOk returns a tuple with the CurrentTriggersSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTriggersSeries

`func (o *O11yGettableRuleStateHistoryStats) SetCurrentTriggersSeries(v O11yTimeSeries)`

SetCurrentTriggersSeries sets CurrentTriggersSeries field to given value.

### HasCurrentTriggersSeries

`func (o *O11yGettableRuleStateHistoryStats) HasCurrentTriggersSeries() bool`

HasCurrentTriggersSeries returns a boolean if a field has been set.

### GetPastAvgResolutionTime

`func (o *O11yGettableRuleStateHistoryStats) GetPastAvgResolutionTime() float32`

GetPastAvgResolutionTime returns the PastAvgResolutionTime field if non-nil, zero value otherwise.

### GetPastAvgResolutionTimeOk

`func (o *O11yGettableRuleStateHistoryStats) GetPastAvgResolutionTimeOk() (*float32, bool)`

GetPastAvgResolutionTimeOk returns a tuple with the PastAvgResolutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastAvgResolutionTime

`func (o *O11yGettableRuleStateHistoryStats) SetPastAvgResolutionTime(v float32)`

SetPastAvgResolutionTime sets PastAvgResolutionTime field to given value.

### HasPastAvgResolutionTime

`func (o *O11yGettableRuleStateHistoryStats) HasPastAvgResolutionTime() bool`

HasPastAvgResolutionTime returns a boolean if a field has been set.

### GetPastAvgResolutionTimeSeries

`func (o *O11yGettableRuleStateHistoryStats) GetPastAvgResolutionTimeSeries() O11yTimeSeries`

GetPastAvgResolutionTimeSeries returns the PastAvgResolutionTimeSeries field if non-nil, zero value otherwise.

### GetPastAvgResolutionTimeSeriesOk

`func (o *O11yGettableRuleStateHistoryStats) GetPastAvgResolutionTimeSeriesOk() (*O11yTimeSeries, bool)`

GetPastAvgResolutionTimeSeriesOk returns a tuple with the PastAvgResolutionTimeSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastAvgResolutionTimeSeries

`func (o *O11yGettableRuleStateHistoryStats) SetPastAvgResolutionTimeSeries(v O11yTimeSeries)`

SetPastAvgResolutionTimeSeries sets PastAvgResolutionTimeSeries field to given value.

### HasPastAvgResolutionTimeSeries

`func (o *O11yGettableRuleStateHistoryStats) HasPastAvgResolutionTimeSeries() bool`

HasPastAvgResolutionTimeSeries returns a boolean if a field has been set.

### GetPastTriggersSeries

`func (o *O11yGettableRuleStateHistoryStats) GetPastTriggersSeries() O11yTimeSeries`

GetPastTriggersSeries returns the PastTriggersSeries field if non-nil, zero value otherwise.

### GetPastTriggersSeriesOk

`func (o *O11yGettableRuleStateHistoryStats) GetPastTriggersSeriesOk() (*O11yTimeSeries, bool)`

GetPastTriggersSeriesOk returns a tuple with the PastTriggersSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastTriggersSeries

`func (o *O11yGettableRuleStateHistoryStats) SetPastTriggersSeries(v O11yTimeSeries)`

SetPastTriggersSeries sets PastTriggersSeries field to given value.

### HasPastTriggersSeries

`func (o *O11yGettableRuleStateHistoryStats) HasPastTriggersSeries() bool`

HasPastTriggersSeries returns a boolean if a field has been set.

### GetTotalCurrentTriggers

`func (o *O11yGettableRuleStateHistoryStats) GetTotalCurrentTriggers() int32`

GetTotalCurrentTriggers returns the TotalCurrentTriggers field if non-nil, zero value otherwise.

### GetTotalCurrentTriggersOk

`func (o *O11yGettableRuleStateHistoryStats) GetTotalCurrentTriggersOk() (*int32, bool)`

GetTotalCurrentTriggersOk returns a tuple with the TotalCurrentTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCurrentTriggers

`func (o *O11yGettableRuleStateHistoryStats) SetTotalCurrentTriggers(v int32)`

SetTotalCurrentTriggers sets TotalCurrentTriggers field to given value.

### HasTotalCurrentTriggers

`func (o *O11yGettableRuleStateHistoryStats) HasTotalCurrentTriggers() bool`

HasTotalCurrentTriggers returns a boolean if a field has been set.

### GetTotalPastTriggers

`func (o *O11yGettableRuleStateHistoryStats) GetTotalPastTriggers() int32`

GetTotalPastTriggers returns the TotalPastTriggers field if non-nil, zero value otherwise.

### GetTotalPastTriggersOk

`func (o *O11yGettableRuleStateHistoryStats) GetTotalPastTriggersOk() (*int32, bool)`

GetTotalPastTriggersOk returns a tuple with the TotalPastTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPastTriggers

`func (o *O11yGettableRuleStateHistoryStats) SetTotalPastTriggers(v int32)`

SetTotalPastTriggers sets TotalPastTriggers field to given value.

### HasTotalPastTriggers

`func (o *O11yGettableRuleStateHistoryStats) HasTotalPastTriggers() bool`

HasTotalPastTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


