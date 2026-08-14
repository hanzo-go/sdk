# O11yGCPServiceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**O11yGCPServiceLogsConfig**](O11yGCPServiceLogsConfig.md) |  | [optional] 
**Metrics** | Pointer to [**O11yGCPServiceMetricsConfig**](O11yGCPServiceMetricsConfig.md) |  | [optional] 

## Methods

### NewO11yGCPServiceConfig

`func NewO11yGCPServiceConfig() *O11yGCPServiceConfig`

NewO11yGCPServiceConfig instantiates a new O11yGCPServiceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yGCPServiceConfigWithDefaults

`func NewO11yGCPServiceConfigWithDefaults() *O11yGCPServiceConfig`

NewO11yGCPServiceConfigWithDefaults instantiates a new O11yGCPServiceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *O11yGCPServiceConfig) GetLogs() O11yGCPServiceLogsConfig`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *O11yGCPServiceConfig) GetLogsOk() (*O11yGCPServiceLogsConfig, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *O11yGCPServiceConfig) SetLogs(v O11yGCPServiceLogsConfig)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *O11yGCPServiceConfig) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetMetrics

`func (o *O11yGCPServiceConfig) GetMetrics() O11yGCPServiceMetricsConfig`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *O11yGCPServiceConfig) GetMetricsOk() (*O11yGCPServiceMetricsConfig, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *O11yGCPServiceConfig) SetMetrics(v O11yGCPServiceMetricsConfig)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *O11yGCPServiceConfig) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


