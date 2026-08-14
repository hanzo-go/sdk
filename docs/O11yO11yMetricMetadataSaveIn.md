# O11yO11yMetricMetadataSaveIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description describes the metric. | [optional] 
**IsMonotonic** | Pointer to **bool** | IsMonotonic marks a sum that only ever increases. | [optional] 
**MetricName** | **string** | MetricName is the metric to update. Required. | 
**Temporality** | Pointer to **string** | Temporality is delta or cumulative. | [optional] 
**Type** | Pointer to **string** | Type is the metric type, e.g. gauge, sum, histogram. | [optional] 
**Unit** | Pointer to **string** | Unit is the metric&#39;s unit. | [optional] 

## Methods

### NewO11yO11yMetricMetadataSaveIn

`func NewO11yO11yMetricMetadataSaveIn(metricName string, ) *O11yO11yMetricMetadataSaveIn`

NewO11yO11yMetricMetadataSaveIn instantiates a new O11yO11yMetricMetadataSaveIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yMetricMetadataSaveInWithDefaults

`func NewO11yO11yMetricMetadataSaveInWithDefaults() *O11yO11yMetricMetadataSaveIn`

NewO11yO11yMetricMetadataSaveInWithDefaults instantiates a new O11yO11yMetricMetadataSaveIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *O11yO11yMetricMetadataSaveIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yO11yMetricMetadataSaveIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yO11yMetricMetadataSaveIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yO11yMetricMetadataSaveIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsMonotonic

`func (o *O11yO11yMetricMetadataSaveIn) GetIsMonotonic() bool`

GetIsMonotonic returns the IsMonotonic field if non-nil, zero value otherwise.

### GetIsMonotonicOk

`func (o *O11yO11yMetricMetadataSaveIn) GetIsMonotonicOk() (*bool, bool)`

GetIsMonotonicOk returns a tuple with the IsMonotonic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMonotonic

`func (o *O11yO11yMetricMetadataSaveIn) SetIsMonotonic(v bool)`

SetIsMonotonic sets IsMonotonic field to given value.

### HasIsMonotonic

`func (o *O11yO11yMetricMetadataSaveIn) HasIsMonotonic() bool`

HasIsMonotonic returns a boolean if a field has been set.

### GetMetricName

`func (o *O11yO11yMetricMetadataSaveIn) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *O11yO11yMetricMetadataSaveIn) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *O11yO11yMetricMetadataSaveIn) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetTemporality

`func (o *O11yO11yMetricMetadataSaveIn) GetTemporality() string`

GetTemporality returns the Temporality field if non-nil, zero value otherwise.

### GetTemporalityOk

`func (o *O11yO11yMetricMetadataSaveIn) GetTemporalityOk() (*string, bool)`

GetTemporalityOk returns a tuple with the Temporality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporality

`func (o *O11yO11yMetricMetadataSaveIn) SetTemporality(v string)`

SetTemporality sets Temporality field to given value.

### HasTemporality

`func (o *O11yO11yMetricMetadataSaveIn) HasTemporality() bool`

HasTemporality returns a boolean if a field has been set.

### GetType

`func (o *O11yO11yMetricMetadataSaveIn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yO11yMetricMetadataSaveIn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yO11yMetricMetadataSaveIn) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yO11yMetricMetadataSaveIn) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *O11yO11yMetricMetadataSaveIn) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *O11yO11yMetricMetadataSaveIn) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *O11yO11yMetricMetadataSaveIn) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *O11yO11yMetricMetadataSaveIn) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


