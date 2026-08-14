# O11yAzureServiceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**O11yAzureServiceLogsConfig**](O11yAzureServiceLogsConfig.md) |  | [optional] 
**Metrics** | Pointer to [**O11yAzureServiceMetricsConfig**](O11yAzureServiceMetricsConfig.md) |  | [optional] 

## Methods

### NewO11yAzureServiceConfig

`func NewO11yAzureServiceConfig() *O11yAzureServiceConfig`

NewO11yAzureServiceConfig instantiates a new O11yAzureServiceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAzureServiceConfigWithDefaults

`func NewO11yAzureServiceConfigWithDefaults() *O11yAzureServiceConfig`

NewO11yAzureServiceConfigWithDefaults instantiates a new O11yAzureServiceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *O11yAzureServiceConfig) GetLogs() O11yAzureServiceLogsConfig`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *O11yAzureServiceConfig) GetLogsOk() (*O11yAzureServiceLogsConfig, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *O11yAzureServiceConfig) SetLogs(v O11yAzureServiceLogsConfig)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *O11yAzureServiceConfig) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetMetrics

`func (o *O11yAzureServiceConfig) GetMetrics() O11yAzureServiceMetricsConfig`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *O11yAzureServiceConfig) GetMetricsOk() (*O11yAzureServiceMetricsConfig, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *O11yAzureServiceConfig) SetMetrics(v O11yAzureServiceMetricsConfig)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *O11yAzureServiceConfig) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


