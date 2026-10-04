# EventTimeseries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** | End is the window&#39;s exclusive upper bound, RFC3339 UTC. | [optional] 
**Interval** | Pointer to **string** | Interval is the bucket width: hour or day. | [optional] 
**Range** | Pointer to **string** | Range is the window that was actually applied: 24h, 7d, 30d or custom. | [optional] 
**Scope** | Pointer to [**EventScope**](EventScope.md) | Scope names the tenant these numbers belong to. | [optional] 
**Series** | Pointer to [**[]EventUsagePoint**](EventUsagePoint.md) | Series is one point per bucket, oldest first, with empty buckets zero-filled. | [optional] 
**Source** | Pointer to **string** | Source is the warehouse table the series read. | [optional] 
**Start** | Pointer to **string** | Start is the window&#39;s inclusive lower bound, RFC3339 UTC. | [optional] 

## Methods

### NewEventTimeseries

`func NewEventTimeseries() *EventTimeseries`

NewEventTimeseries instantiates a new EventTimeseries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTimeseriesWithDefaults

`func NewEventTimeseriesWithDefaults() *EventTimeseries`

NewEventTimeseriesWithDefaults instantiates a new EventTimeseries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *EventTimeseries) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *EventTimeseries) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *EventTimeseries) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *EventTimeseries) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetInterval

`func (o *EventTimeseries) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *EventTimeseries) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *EventTimeseries) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *EventTimeseries) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetRange

`func (o *EventTimeseries) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *EventTimeseries) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *EventTimeseries) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *EventTimeseries) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetScope

`func (o *EventTimeseries) GetScope() EventScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EventTimeseries) GetScopeOk() (*EventScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EventTimeseries) SetScope(v EventScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *EventTimeseries) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSeries

`func (o *EventTimeseries) GetSeries() []EventUsagePoint`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *EventTimeseries) GetSeriesOk() (*[]EventUsagePoint, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *EventTimeseries) SetSeries(v []EventUsagePoint)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *EventTimeseries) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetSource

`func (o *EventTimeseries) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EventTimeseries) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EventTimeseries) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *EventTimeseries) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStart

`func (o *EventTimeseries) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *EventTimeseries) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *EventTimeseries) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *EventTimeseries) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


