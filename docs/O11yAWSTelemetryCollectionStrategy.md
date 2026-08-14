# O11yAWSTelemetryCollectionStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**O11yAWSLogsCollectionStrategy**](O11yAWSLogsCollectionStrategy.md) |  | [optional] 
**Metrics** | Pointer to [**O11yAWSMetricsCollectionStrategy**](O11yAWSMetricsCollectionStrategy.md) |  | [optional] 
**S3Buckets** | Pointer to **map[string][]string** | Only available in S3 Sync Service Type in AWS | [optional] 

## Methods

### NewO11yAWSTelemetryCollectionStrategy

`func NewO11yAWSTelemetryCollectionStrategy() *O11yAWSTelemetryCollectionStrategy`

NewO11yAWSTelemetryCollectionStrategy instantiates a new O11yAWSTelemetryCollectionStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAWSTelemetryCollectionStrategyWithDefaults

`func NewO11yAWSTelemetryCollectionStrategyWithDefaults() *O11yAWSTelemetryCollectionStrategy`

NewO11yAWSTelemetryCollectionStrategyWithDefaults instantiates a new O11yAWSTelemetryCollectionStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *O11yAWSTelemetryCollectionStrategy) GetLogs() O11yAWSLogsCollectionStrategy`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *O11yAWSTelemetryCollectionStrategy) GetLogsOk() (*O11yAWSLogsCollectionStrategy, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *O11yAWSTelemetryCollectionStrategy) SetLogs(v O11yAWSLogsCollectionStrategy)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *O11yAWSTelemetryCollectionStrategy) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetMetrics

`func (o *O11yAWSTelemetryCollectionStrategy) GetMetrics() O11yAWSMetricsCollectionStrategy`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *O11yAWSTelemetryCollectionStrategy) GetMetricsOk() (*O11yAWSMetricsCollectionStrategy, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *O11yAWSTelemetryCollectionStrategy) SetMetrics(v O11yAWSMetricsCollectionStrategy)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *O11yAWSTelemetryCollectionStrategy) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetS3Buckets

`func (o *O11yAWSTelemetryCollectionStrategy) GetS3Buckets() map[string][]string`

GetS3Buckets returns the S3Buckets field if non-nil, zero value otherwise.

### GetS3BucketsOk

`func (o *O11yAWSTelemetryCollectionStrategy) GetS3BucketsOk() (*map[string][]string, bool)`

GetS3BucketsOk returns a tuple with the S3Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Buckets

`func (o *O11yAWSTelemetryCollectionStrategy) SetS3Buckets(v map[string][]string)`

SetS3Buckets sets S3Buckets field to given value.

### HasS3Buckets

`func (o *O11yAWSTelemetryCollectionStrategy) HasS3Buckets() bool`

HasS3Buckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


