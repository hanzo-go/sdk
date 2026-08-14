# O11yStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentAvgResolutionTime** | Pointer to **string** |  | [optional] 
**CurrentAvgResolutionTimeSeries** | Pointer to [**O11ySeries**](O11ySeries.md) |  | [optional] 
**CurrentTriggersSeries** | Pointer to [**O11ySeries**](O11ySeries.md) |  | [optional] 
**PastAvgResolutionTime** | Pointer to **string** |  | [optional] 
**PastAvgResolutionTimeSeries** | Pointer to [**O11ySeries**](O11ySeries.md) |  | [optional] 
**PastTriggersSeries** | Pointer to [**O11ySeries**](O11ySeries.md) |  | [optional] 
**TotalCurrentTriggers** | Pointer to **int32** |  | [optional] 
**TotalPastTriggers** | Pointer to **int32** |  | [optional] 

## Methods

### NewO11yStats

`func NewO11yStats() *O11yStats`

NewO11yStats instantiates a new O11yStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yStatsWithDefaults

`func NewO11yStatsWithDefaults() *O11yStats`

NewO11yStatsWithDefaults instantiates a new O11yStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentAvgResolutionTime

`func (o *O11yStats) GetCurrentAvgResolutionTime() string`

GetCurrentAvgResolutionTime returns the CurrentAvgResolutionTime field if non-nil, zero value otherwise.

### GetCurrentAvgResolutionTimeOk

`func (o *O11yStats) GetCurrentAvgResolutionTimeOk() (*string, bool)`

GetCurrentAvgResolutionTimeOk returns a tuple with the CurrentAvgResolutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvgResolutionTime

`func (o *O11yStats) SetCurrentAvgResolutionTime(v string)`

SetCurrentAvgResolutionTime sets CurrentAvgResolutionTime field to given value.

### HasCurrentAvgResolutionTime

`func (o *O11yStats) HasCurrentAvgResolutionTime() bool`

HasCurrentAvgResolutionTime returns a boolean if a field has been set.

### GetCurrentAvgResolutionTimeSeries

`func (o *O11yStats) GetCurrentAvgResolutionTimeSeries() O11ySeries`

GetCurrentAvgResolutionTimeSeries returns the CurrentAvgResolutionTimeSeries field if non-nil, zero value otherwise.

### GetCurrentAvgResolutionTimeSeriesOk

`func (o *O11yStats) GetCurrentAvgResolutionTimeSeriesOk() (*O11ySeries, bool)`

GetCurrentAvgResolutionTimeSeriesOk returns a tuple with the CurrentAvgResolutionTimeSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAvgResolutionTimeSeries

`func (o *O11yStats) SetCurrentAvgResolutionTimeSeries(v O11ySeries)`

SetCurrentAvgResolutionTimeSeries sets CurrentAvgResolutionTimeSeries field to given value.

### HasCurrentAvgResolutionTimeSeries

`func (o *O11yStats) HasCurrentAvgResolutionTimeSeries() bool`

HasCurrentAvgResolutionTimeSeries returns a boolean if a field has been set.

### GetCurrentTriggersSeries

`func (o *O11yStats) GetCurrentTriggersSeries() O11ySeries`

GetCurrentTriggersSeries returns the CurrentTriggersSeries field if non-nil, zero value otherwise.

### GetCurrentTriggersSeriesOk

`func (o *O11yStats) GetCurrentTriggersSeriesOk() (*O11ySeries, bool)`

GetCurrentTriggersSeriesOk returns a tuple with the CurrentTriggersSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTriggersSeries

`func (o *O11yStats) SetCurrentTriggersSeries(v O11ySeries)`

SetCurrentTriggersSeries sets CurrentTriggersSeries field to given value.

### HasCurrentTriggersSeries

`func (o *O11yStats) HasCurrentTriggersSeries() bool`

HasCurrentTriggersSeries returns a boolean if a field has been set.

### GetPastAvgResolutionTime

`func (o *O11yStats) GetPastAvgResolutionTime() string`

GetPastAvgResolutionTime returns the PastAvgResolutionTime field if non-nil, zero value otherwise.

### GetPastAvgResolutionTimeOk

`func (o *O11yStats) GetPastAvgResolutionTimeOk() (*string, bool)`

GetPastAvgResolutionTimeOk returns a tuple with the PastAvgResolutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastAvgResolutionTime

`func (o *O11yStats) SetPastAvgResolutionTime(v string)`

SetPastAvgResolutionTime sets PastAvgResolutionTime field to given value.

### HasPastAvgResolutionTime

`func (o *O11yStats) HasPastAvgResolutionTime() bool`

HasPastAvgResolutionTime returns a boolean if a field has been set.

### GetPastAvgResolutionTimeSeries

`func (o *O11yStats) GetPastAvgResolutionTimeSeries() O11ySeries`

GetPastAvgResolutionTimeSeries returns the PastAvgResolutionTimeSeries field if non-nil, zero value otherwise.

### GetPastAvgResolutionTimeSeriesOk

`func (o *O11yStats) GetPastAvgResolutionTimeSeriesOk() (*O11ySeries, bool)`

GetPastAvgResolutionTimeSeriesOk returns a tuple with the PastAvgResolutionTimeSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastAvgResolutionTimeSeries

`func (o *O11yStats) SetPastAvgResolutionTimeSeries(v O11ySeries)`

SetPastAvgResolutionTimeSeries sets PastAvgResolutionTimeSeries field to given value.

### HasPastAvgResolutionTimeSeries

`func (o *O11yStats) HasPastAvgResolutionTimeSeries() bool`

HasPastAvgResolutionTimeSeries returns a boolean if a field has been set.

### GetPastTriggersSeries

`func (o *O11yStats) GetPastTriggersSeries() O11ySeries`

GetPastTriggersSeries returns the PastTriggersSeries field if non-nil, zero value otherwise.

### GetPastTriggersSeriesOk

`func (o *O11yStats) GetPastTriggersSeriesOk() (*O11ySeries, bool)`

GetPastTriggersSeriesOk returns a tuple with the PastTriggersSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastTriggersSeries

`func (o *O11yStats) SetPastTriggersSeries(v O11ySeries)`

SetPastTriggersSeries sets PastTriggersSeries field to given value.

### HasPastTriggersSeries

`func (o *O11yStats) HasPastTriggersSeries() bool`

HasPastTriggersSeries returns a boolean if a field has been set.

### GetTotalCurrentTriggers

`func (o *O11yStats) GetTotalCurrentTriggers() int32`

GetTotalCurrentTriggers returns the TotalCurrentTriggers field if non-nil, zero value otherwise.

### GetTotalCurrentTriggersOk

`func (o *O11yStats) GetTotalCurrentTriggersOk() (*int32, bool)`

GetTotalCurrentTriggersOk returns a tuple with the TotalCurrentTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCurrentTriggers

`func (o *O11yStats) SetTotalCurrentTriggers(v int32)`

SetTotalCurrentTriggers sets TotalCurrentTriggers field to given value.

### HasTotalCurrentTriggers

`func (o *O11yStats) HasTotalCurrentTriggers() bool`

HasTotalCurrentTriggers returns a boolean if a field has been set.

### GetTotalPastTriggers

`func (o *O11yStats) GetTotalPastTriggers() int32`

GetTotalPastTriggers returns the TotalPastTriggers field if non-nil, zero value otherwise.

### GetTotalPastTriggersOk

`func (o *O11yStats) GetTotalPastTriggersOk() (*int32, bool)`

GetTotalPastTriggersOk returns a tuple with the TotalPastTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPastTriggers

`func (o *O11yStats) SetTotalPastTriggers(v int32)`

SetTotalPastTriggers sets TotalPastTriggers field to given value.

### HasTotalPastTriggers

`func (o *O11yStats) HasTotalPastTriggers() bool`

HasTotalPastTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


