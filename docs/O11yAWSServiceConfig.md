# O11yAWSServiceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**O11yAWSServiceLogsConfig**](O11yAWSServiceLogsConfig.md) |  | [optional] 
**Metrics** | Pointer to [**O11yAWSServiceMetricsConfig**](O11yAWSServiceMetricsConfig.md) |  | [optional] 

## Methods

### NewO11yAWSServiceConfig

`func NewO11yAWSServiceConfig() *O11yAWSServiceConfig`

NewO11yAWSServiceConfig instantiates a new O11yAWSServiceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAWSServiceConfigWithDefaults

`func NewO11yAWSServiceConfigWithDefaults() *O11yAWSServiceConfig`

NewO11yAWSServiceConfigWithDefaults instantiates a new O11yAWSServiceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *O11yAWSServiceConfig) GetLogs() O11yAWSServiceLogsConfig`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *O11yAWSServiceConfig) GetLogsOk() (*O11yAWSServiceLogsConfig, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *O11yAWSServiceConfig) SetLogs(v O11yAWSServiceLogsConfig)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *O11yAWSServiceConfig) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetMetrics

`func (o *O11yAWSServiceConfig) GetMetrics() O11yAWSServiceMetricsConfig`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *O11yAWSServiceConfig) GetMetricsOk() (*O11yAWSServiceMetricsConfig, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *O11yAWSServiceConfig) SetMetrics(v O11yAWSServiceMetricsConfig)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *O11yAWSServiceConfig) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


