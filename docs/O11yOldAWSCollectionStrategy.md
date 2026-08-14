# O11yOldAWSCollectionStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsLogs** | Pointer to [**O11yOldAWSLogsStrategy**](O11yOldAWSLogsStrategy.md) |  | [optional] 
**AwsMetrics** | Pointer to [**O11yOldAWSMetricsStrategy**](O11yOldAWSMetricsStrategy.md) |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**S3Buckets** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewO11yOldAWSCollectionStrategy

`func NewO11yOldAWSCollectionStrategy() *O11yOldAWSCollectionStrategy`

NewO11yOldAWSCollectionStrategy instantiates a new O11yOldAWSCollectionStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yOldAWSCollectionStrategyWithDefaults

`func NewO11yOldAWSCollectionStrategyWithDefaults() *O11yOldAWSCollectionStrategy`

NewO11yOldAWSCollectionStrategyWithDefaults instantiates a new O11yOldAWSCollectionStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsLogs

`func (o *O11yOldAWSCollectionStrategy) GetAwsLogs() O11yOldAWSLogsStrategy`

GetAwsLogs returns the AwsLogs field if non-nil, zero value otherwise.

### GetAwsLogsOk

`func (o *O11yOldAWSCollectionStrategy) GetAwsLogsOk() (*O11yOldAWSLogsStrategy, bool)`

GetAwsLogsOk returns a tuple with the AwsLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsLogs

`func (o *O11yOldAWSCollectionStrategy) SetAwsLogs(v O11yOldAWSLogsStrategy)`

SetAwsLogs sets AwsLogs field to given value.

### HasAwsLogs

`func (o *O11yOldAWSCollectionStrategy) HasAwsLogs() bool`

HasAwsLogs returns a boolean if a field has been set.

### GetAwsMetrics

`func (o *O11yOldAWSCollectionStrategy) GetAwsMetrics() O11yOldAWSMetricsStrategy`

GetAwsMetrics returns the AwsMetrics field if non-nil, zero value otherwise.

### GetAwsMetricsOk

`func (o *O11yOldAWSCollectionStrategy) GetAwsMetricsOk() (*O11yOldAWSMetricsStrategy, bool)`

GetAwsMetricsOk returns a tuple with the AwsMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsMetrics

`func (o *O11yOldAWSCollectionStrategy) SetAwsMetrics(v O11yOldAWSMetricsStrategy)`

SetAwsMetrics sets AwsMetrics field to given value.

### HasAwsMetrics

`func (o *O11yOldAWSCollectionStrategy) HasAwsMetrics() bool`

HasAwsMetrics returns a boolean if a field has been set.

### GetProvider

`func (o *O11yOldAWSCollectionStrategy) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *O11yOldAWSCollectionStrategy) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *O11yOldAWSCollectionStrategy) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *O11yOldAWSCollectionStrategy) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetS3Buckets

`func (o *O11yOldAWSCollectionStrategy) GetS3Buckets() map[string][]string`

GetS3Buckets returns the S3Buckets field if non-nil, zero value otherwise.

### GetS3BucketsOk

`func (o *O11yOldAWSCollectionStrategy) GetS3BucketsOk() (*map[string][]string, bool)`

GetS3BucketsOk returns a tuple with the S3Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Buckets

`func (o *O11yOldAWSCollectionStrategy) SetS3Buckets(v map[string][]string)`

SetS3Buckets sets S3Buckets field to given value.

### HasS3Buckets

`func (o *O11yOldAWSCollectionStrategy) HasS3Buckets() bool`

HasS3Buckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


