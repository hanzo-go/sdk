# O11yO11yQueryStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BytesScanned** | Pointer to **int32** | BytesScanned is how many bytes the query read. | [optional] 
**DurationMs** | Pointer to **int32** | DurationMS is how long the query took, in milliseconds. | [optional] 
**RowsScanned** | Pointer to **int32** | RowsScanned is how many rows the query read. | [optional] 
**StepIntervals** | Pointer to **map[string]int32** | StepIntervals is the step used per query, in seconds. | [optional] 

## Methods

### NewO11yO11yQueryStats

`func NewO11yO11yQueryStats() *O11yO11yQueryStats`

NewO11yO11yQueryStats instantiates a new O11yO11yQueryStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yQueryStatsWithDefaults

`func NewO11yO11yQueryStatsWithDefaults() *O11yO11yQueryStats`

NewO11yO11yQueryStatsWithDefaults instantiates a new O11yO11yQueryStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytesScanned

`func (o *O11yO11yQueryStats) GetBytesScanned() int32`

GetBytesScanned returns the BytesScanned field if non-nil, zero value otherwise.

### GetBytesScannedOk

`func (o *O11yO11yQueryStats) GetBytesScannedOk() (*int32, bool)`

GetBytesScannedOk returns a tuple with the BytesScanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesScanned

`func (o *O11yO11yQueryStats) SetBytesScanned(v int32)`

SetBytesScanned sets BytesScanned field to given value.

### HasBytesScanned

`func (o *O11yO11yQueryStats) HasBytesScanned() bool`

HasBytesScanned returns a boolean if a field has been set.

### GetDurationMs

`func (o *O11yO11yQueryStats) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *O11yO11yQueryStats) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *O11yO11yQueryStats) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *O11yO11yQueryStats) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetRowsScanned

`func (o *O11yO11yQueryStats) GetRowsScanned() int32`

GetRowsScanned returns the RowsScanned field if non-nil, zero value otherwise.

### GetRowsScannedOk

`func (o *O11yO11yQueryStats) GetRowsScannedOk() (*int32, bool)`

GetRowsScannedOk returns a tuple with the RowsScanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowsScanned

`func (o *O11yO11yQueryStats) SetRowsScanned(v int32)`

SetRowsScanned sets RowsScanned field to given value.

### HasRowsScanned

`func (o *O11yO11yQueryStats) HasRowsScanned() bool`

HasRowsScanned returns a boolean if a field has been set.

### GetStepIntervals

`func (o *O11yO11yQueryStats) GetStepIntervals() map[string]int32`

GetStepIntervals returns the StepIntervals field if non-nil, zero value otherwise.

### GetStepIntervalsOk

`func (o *O11yO11yQueryStats) GetStepIntervalsOk() (*map[string]int32, bool)`

GetStepIntervalsOk returns a tuple with the StepIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepIntervals

`func (o *O11yO11yQueryStats) SetStepIntervals(v map[string]int32)`

SetStepIntervals sets StepIntervals field to given value.

### HasStepIntervals

`func (o *O11yO11yQueryStats) HasStepIntervals() bool`

HasStepIntervals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


