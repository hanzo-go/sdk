# O11yO11yMetricStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description describes the metric. | [optional] 
**MetricName** | Pointer to **string** | MetricName is the metric&#39;s name. | [optional] 
**Samples** | Pointer to **int32** | Samples is how many samples the metric had. | [optional] 
**Timeseries** | Pointer to **int32** | TimeSeries is how many time series the metric had. | [optional] 
**Type** | Pointer to **string** | Type is the metric type. | [optional] 
**Unit** | Pointer to **string** | Unit is the metric&#39;s unit. | [optional] 

## Methods

### NewO11yO11yMetricStat

`func NewO11yO11yMetricStat() *O11yO11yMetricStat`

NewO11yO11yMetricStat instantiates a new O11yO11yMetricStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yMetricStatWithDefaults

`func NewO11yO11yMetricStatWithDefaults() *O11yO11yMetricStat`

NewO11yO11yMetricStatWithDefaults instantiates a new O11yO11yMetricStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *O11yO11yMetricStat) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yO11yMetricStat) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yO11yMetricStat) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yO11yMetricStat) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetricName

`func (o *O11yO11yMetricStat) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *O11yO11yMetricStat) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *O11yO11yMetricStat) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *O11yO11yMetricStat) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetSamples

`func (o *O11yO11yMetricStat) GetSamples() int32`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *O11yO11yMetricStat) GetSamplesOk() (*int32, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *O11yO11yMetricStat) SetSamples(v int32)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *O11yO11yMetricStat) HasSamples() bool`

HasSamples returns a boolean if a field has been set.

### GetTimeseries

`func (o *O11yO11yMetricStat) GetTimeseries() int32`

GetTimeseries returns the Timeseries field if non-nil, zero value otherwise.

### GetTimeseriesOk

`func (o *O11yO11yMetricStat) GetTimeseriesOk() (*int32, bool)`

GetTimeseriesOk returns a tuple with the Timeseries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseries

`func (o *O11yO11yMetricStat) SetTimeseries(v int32)`

SetTimeseries sets Timeseries field to given value.

### HasTimeseries

`func (o *O11yO11yMetricStat) HasTimeseries() bool`

HasTimeseries returns a boolean if a field has been set.

### GetType

`func (o *O11yO11yMetricStat) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yO11yMetricStat) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yO11yMetricStat) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yO11yMetricStat) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *O11yO11yMetricStat) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *O11yO11yMetricStat) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *O11yO11yMetricStat) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *O11yO11yMetricStat) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


