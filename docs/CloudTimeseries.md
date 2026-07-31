# CloudTimeseries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** | End is the window&#39;s exclusive upper bound, RFC3339 UTC. | [optional] 
**Interval** | Pointer to **string** | Interval is the bucket width: hour or day. | [optional] 
**Range** | Pointer to **string** | Range is the window that was actually applied: 24h, 7d, 30d or custom. | [optional] 
**Scope** | Pointer to [**CloudScope**](CloudScope.md) | Scope names the tenant these numbers belong to. | [optional] 
**Series** | Pointer to [**[]CloudUsagePoint**](CloudUsagePoint.md) | Series is one point per bucket, oldest first, with empty buckets zero-filled. | [optional] 
**Source** | Pointer to **string** | Source is the warehouse table the series read. | [optional] 
**Start** | Pointer to **string** | Start is the window&#39;s inclusive lower bound, RFC3339 UTC. | [optional] 

## Methods

### NewCloudTimeseries

`func NewCloudTimeseries() *CloudTimeseries`

NewCloudTimeseries instantiates a new CloudTimeseries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTimeseriesWithDefaults

`func NewCloudTimeseriesWithDefaults() *CloudTimeseries`

NewCloudTimeseriesWithDefaults instantiates a new CloudTimeseries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *CloudTimeseries) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CloudTimeseries) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CloudTimeseries) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CloudTimeseries) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetInterval

`func (o *CloudTimeseries) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CloudTimeseries) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CloudTimeseries) SetInterval(v string)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CloudTimeseries) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetRange

`func (o *CloudTimeseries) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CloudTimeseries) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CloudTimeseries) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *CloudTimeseries) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetScope

`func (o *CloudTimeseries) GetScope() CloudScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CloudTimeseries) GetScopeOk() (*CloudScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CloudTimeseries) SetScope(v CloudScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CloudTimeseries) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSeries

`func (o *CloudTimeseries) GetSeries() []CloudUsagePoint`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *CloudTimeseries) GetSeriesOk() (*[]CloudUsagePoint, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *CloudTimeseries) SetSeries(v []CloudUsagePoint)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *CloudTimeseries) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### GetSource

`func (o *CloudTimeseries) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudTimeseries) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudTimeseries) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CloudTimeseries) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStart

`func (o *CloudTimeseries) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CloudTimeseries) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CloudTimeseries) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *CloudTimeseries) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


