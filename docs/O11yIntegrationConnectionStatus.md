# O11yIntegrationConnectionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**O11ySignalConnectionStatus**](O11ySignalConnectionStatus.md) |  | [optional] 
**Metrics** | Pointer to [**O11ySignalConnectionStatus**](O11ySignalConnectionStatus.md) |  | [optional] 

## Methods

### NewO11yIntegrationConnectionStatus

`func NewO11yIntegrationConnectionStatus() *O11yIntegrationConnectionStatus`

NewO11yIntegrationConnectionStatus instantiates a new O11yIntegrationConnectionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yIntegrationConnectionStatusWithDefaults

`func NewO11yIntegrationConnectionStatusWithDefaults() *O11yIntegrationConnectionStatus`

NewO11yIntegrationConnectionStatusWithDefaults instantiates a new O11yIntegrationConnectionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *O11yIntegrationConnectionStatus) GetLogs() O11ySignalConnectionStatus`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *O11yIntegrationConnectionStatus) GetLogsOk() (*O11ySignalConnectionStatus, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *O11yIntegrationConnectionStatus) SetLogs(v O11ySignalConnectionStatus)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *O11yIntegrationConnectionStatus) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetMetrics

`func (o *O11yIntegrationConnectionStatus) GetMetrics() O11ySignalConnectionStatus`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *O11yIntegrationConnectionStatus) GetMetricsOk() (*O11ySignalConnectionStatus, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *O11yIntegrationConnectionStatus) SetMetrics(v O11ySignalConnectionStatus)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *O11yIntegrationConnectionStatus) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


