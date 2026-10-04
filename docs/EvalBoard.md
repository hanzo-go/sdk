# EvalBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByModel** | Pointer to [**[]EvalModelStat**](EvalModelStat.md) | the top models by spend | [optional] 
**Latency** | Pointer to [**EvalLatencyStat**](EvalLatencyStat.md) | overall latency percentiles from the GenAI spans | [optional] 
**Other** | Pointer to [**EvalModelStat**](EvalModelStat.md) | the long tail beyond the top models, folded into one row | [optional] 
**Range** | Pointer to [**EvalBoardRange**](EvalBoardRange.md) | the window they were computed over, echoed back | [optional] 
**Scope** | Pointer to [**EvalBoardScope**](EvalBoardScope.md) | whose numbers these are | [optional] 
**Series** | Pointer to [**[]EvalBoardPoint**](EvalBoardPoint.md) | one gap-filled bucket per interval, so a chart never breaks | [optional] 
**Totals** | Pointer to [**EvalBoardTotals**](EvalBoardTotals.md) | the window&#39;s headline numbers | [optional] 

## Methods

### NewEvalBoard

`func NewEvalBoard() *EvalBoard`

NewEvalBoard instantiates a new EvalBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvalBoardWithDefaults

`func NewEvalBoardWithDefaults() *EvalBoard`

NewEvalBoardWithDefaults instantiates a new EvalBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByModel

`func (o *EvalBoard) GetByModel() []EvalModelStat`

GetByModel returns the ByModel field if non-nil, zero value otherwise.

### GetByModelOk

`func (o *EvalBoard) GetByModelOk() (*[]EvalModelStat, bool)`

GetByModelOk returns a tuple with the ByModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByModel

`func (o *EvalBoard) SetByModel(v []EvalModelStat)`

SetByModel sets ByModel field to given value.

### HasByModel

`func (o *EvalBoard) HasByModel() bool`

HasByModel returns a boolean if a field has been set.

### GetLatency

`func (o *EvalBoard) GetLatency() EvalLatencyStat`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *EvalBoard) GetLatencyOk() (*EvalLatencyStat, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *EvalBoard) SetLatency(v EvalLatencyStat)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *EvalBoard) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetOther

`func (o *EvalBoard) GetOther() EvalModelStat`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *EvalBoard) GetOtherOk() (*EvalModelStat, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *EvalBoard) SetOther(v EvalModelStat)`

SetOther sets Other field to given value.

### HasOther

`func (o *EvalBoard) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetRange

`func (o *EvalBoard) GetRange() EvalBoardRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *EvalBoard) GetRangeOk() (*EvalBoardRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *EvalBoard) SetRange(v EvalBoardRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *EvalBoard) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetScope

`func (o *EvalBoard) GetScope() EvalBoardScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EvalBoard) GetScopeOk() (*EvalBoardScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EvalBoard) SetScope(v EvalBoardScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *EvalBoard) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSeries

`func (o *EvalBoard) GetSeries() []EvalBoardPoint`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *EvalBoard) GetSeriesOk() (*[]EvalBoardPoint, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *EvalBoard) SetSeries(v []EvalBoardPoint)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *EvalBoard) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetTotals

`func (o *EvalBoard) GetTotals() EvalBoardTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *EvalBoard) GetTotalsOk() (*EvalBoardTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *EvalBoard) SetTotals(v EvalBoardTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *EvalBoard) HasTotals() bool`

HasTotals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


