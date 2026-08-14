# O11yO11yMetricSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description describes the metric. | [optional] 
**IsMonotonic** | Pointer to **bool** | IsMonotonic marks a sum that only ever increases. | [optional] 
**MetricName** | Pointer to **string** | MetricName is the metric&#39;s name. | [optional] 
**Temporality** | Pointer to **string** | Temporality is delta or cumulative. | [optional] 
**Type** | Pointer to **string** | Type is the metric type, e.g. gauge, sum, histogram. | [optional] 
**Unit** | Pointer to **string** | Unit is the metric&#39;s unit. | [optional] 

## Methods

### NewO11yO11yMetricSummary

`func NewO11yO11yMetricSummary() *O11yO11yMetricSummary`

NewO11yO11yMetricSummary instantiates a new O11yO11yMetricSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yMetricSummaryWithDefaults

`func NewO11yO11yMetricSummaryWithDefaults() *O11yO11yMetricSummary`

NewO11yO11yMetricSummaryWithDefaults instantiates a new O11yO11yMetricSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *O11yO11yMetricSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yO11yMetricSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yO11yMetricSummary) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yO11yMetricSummary) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsMonotonic

`func (o *O11yO11yMetricSummary) GetIsMonotonic() bool`

GetIsMonotonic returns the IsMonotonic field if non-nil, zero value otherwise.

### GetIsMonotonicOk

`func (o *O11yO11yMetricSummary) GetIsMonotonicOk() (*bool, bool)`

GetIsMonotonicOk returns a tuple with the IsMonotonic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonotonic

`func (o *O11yO11yMetricSummary) SetIsMonotonic(v bool)`

SetIsMonotonic sets IsMonotonic field to given value.

### HasIsMonotonic

`func (o *O11yO11yMetricSummary) HasIsMonotonic() bool`

HasIsMonotonic returns a boolean if a field has been set.

### GetMetricName

`func (o *O11yO11yMetricSummary) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *O11yO11yMetricSummary) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *O11yO11yMetricSummary) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *O11yO11yMetricSummary) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetTemporality

`func (o *O11yO11yMetricSummary) GetTemporality() string`

GetTemporality returns the Temporality field if non-nil, zero value otherwise.

### GetTemporalityOk

`func (o *O11yO11yMetricSummary) GetTemporalityOk() (*string, bool)`

GetTemporalityOk returns a tuple with the Temporality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporality

`func (o *O11yO11yMetricSummary) SetTemporality(v string)`

SetTemporality sets Temporality field to given value.

### HasTemporality

`func (o *O11yO11yMetricSummary) HasTemporality() bool`

HasTemporality returns a boolean if a field has been set.

### GetType

`func (o *O11yO11yMetricSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yO11yMetricSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yO11yMetricSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yO11yMetricSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *O11yO11yMetricSummary) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *O11yO11yMetricSummary) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *O11yO11yMetricSummary) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *O11yO11yMetricSummary) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


