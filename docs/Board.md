# Board

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByModel** | Pointer to [**[]ModelStat**](ModelStat.md) | the top models by spend | [optional] 
**Latency** | Pointer to [**LatencyStat**](LatencyStat.md) | overall latency percentiles from the GenAI spans | [optional] 
**Other** | Pointer to [**ModelStat**](ModelStat.md) | the long tail beyond the top models, folded into one row | [optional] 
**Range** | Pointer to [**BoardRange**](BoardRange.md) | the window they were computed over, echoed back | [optional] 
**Scope** | Pointer to [**BoardScope**](BoardScope.md) | whose numbers these are | [optional] 
**Series** | Pointer to [**[]BoardPoint**](BoardPoint.md) | one gap-filled bucket per interval, so a chart never breaks | [optional] 
**Totals** | Pointer to [**BoardTotals**](BoardTotals.md) | the window&#39;s headline numbers | [optional] 

## Methods

### NewBoard

`func NewBoard() *Board`

NewBoard instantiates a new Board object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardWithDefaults

`func NewBoardWithDefaults() *Board`

NewBoardWithDefaults instantiates a new Board object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByModel

`func (o *Board) GetByModel() []ModelStat`

GetByModel returns the ByModel field if non-nil, zero value otherwise.

### GetByModelOk

`func (o *Board) GetByModelOk() (*[]ModelStat, bool)`

GetByModelOk returns a tuple with the ByModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByModel

`func (o *Board) SetByModel(v []ModelStat)`

SetByModel sets ByModel field to given value.

### HasByModel

`func (o *Board) HasByModel() bool`

HasByModel returns a boolean if a field has been set.

### GetLatency

`func (o *Board) GetLatency() LatencyStat`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *Board) GetLatencyOk() (*LatencyStat, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *Board) SetLatency(v LatencyStat)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *Board) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetOther

`func (o *Board) GetOther() ModelStat`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *Board) GetOtherOk() (*ModelStat, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *Board) SetOther(v ModelStat)`

SetOther sets Other field to given value.

### HasOther

`func (o *Board) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetRange

`func (o *Board) GetRange() BoardRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *Board) GetRangeOk() (*BoardRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *Board) SetRange(v BoardRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *Board) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetScope

`func (o *Board) GetScope() BoardScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Board) GetScopeOk() (*BoardScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Board) SetScope(v BoardScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Board) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSeries

`func (o *Board) GetSeries() []BoardPoint`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *Board) GetSeriesOk() (*[]BoardPoint, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *Board) SetSeries(v []BoardPoint)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *Board) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetTotals

`func (o *Board) GetTotals() BoardTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *Board) GetTotalsOk() (*BoardTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *Board) SetTotals(v BoardTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *Board) HasTotals() bool`

HasTotals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


