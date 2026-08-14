# O11yO11yMetricStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to [**[]O11yO11yMetricStat**](O11yO11yMetricStat.md) | Metrics are the counted metrics. | [optional] 
**Total** | Pointer to **int32** | Total is how many metrics matched, across all pages. | [optional] 

## Methods

### NewO11yO11yMetricStats

`func NewO11yO11yMetricStats() *O11yO11yMetricStats`

NewO11yO11yMetricStats instantiates a new O11yO11yMetricStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yMetricStatsWithDefaults

`func NewO11yO11yMetricStatsWithDefaults() *O11yO11yMetricStats`

NewO11yO11yMetricStatsWithDefaults instantiates a new O11yO11yMetricStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *O11yO11yMetricStats) GetMetrics() []O11yO11yMetricStat`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *O11yO11yMetricStats) GetMetricsOk() (*[]O11yO11yMetricStat, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *O11yO11yMetricStats) SetMetrics(v []O11yO11yMetricStat)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *O11yO11yMetricStats) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetTotal

`func (o *O11yO11yMetricStats) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *O11yO11yMetricStats) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *O11yO11yMetricStats) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *O11yO11yMetricStats) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


