# O11yIntegrationConnectionTests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**O11yLogsConnectionTest**](O11yLogsConnectionTest.md) |  | [optional] 
**Metrics** | Pointer to **[]string** | Metric names expected to have been received for the integration. | [optional] 

## Methods

### NewO11yIntegrationConnectionTests

`func NewO11yIntegrationConnectionTests() *O11yIntegrationConnectionTests`

NewO11yIntegrationConnectionTests instantiates a new O11yIntegrationConnectionTests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yIntegrationConnectionTestsWithDefaults

`func NewO11yIntegrationConnectionTestsWithDefaults() *O11yIntegrationConnectionTests`

NewO11yIntegrationConnectionTestsWithDefaults instantiates a new O11yIntegrationConnectionTests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *O11yIntegrationConnectionTests) GetLogs() O11yLogsConnectionTest`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *O11yIntegrationConnectionTests) GetLogsOk() (*O11yLogsConnectionTest, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *O11yIntegrationConnectionTests) SetLogs(v O11yLogsConnectionTest)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *O11yIntegrationConnectionTests) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetMetrics

`func (o *O11yIntegrationConnectionTests) GetMetrics() []string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *O11yIntegrationConnectionTests) GetMetricsOk() (*[]string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *O11yIntegrationConnectionTests) SetMetrics(v []string)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *O11yIntegrationConnectionTests) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


