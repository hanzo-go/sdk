# O11yDataCollectedForIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**[]O11yIntegrationsCollectedLogAttribute**](O11yIntegrationsCollectedLogAttribute.md) |  | [optional] 
**Metrics** | Pointer to [**[]O11yIntegrationsCollectedMetric**](O11yIntegrationsCollectedMetric.md) |  | [optional] 

## Methods

### NewO11yDataCollectedForIntegration

`func NewO11yDataCollectedForIntegration() *O11yDataCollectedForIntegration`

NewO11yDataCollectedForIntegration instantiates a new O11yDataCollectedForIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yDataCollectedForIntegrationWithDefaults

`func NewO11yDataCollectedForIntegrationWithDefaults() *O11yDataCollectedForIntegration`

NewO11yDataCollectedForIntegrationWithDefaults instantiates a new O11yDataCollectedForIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *O11yDataCollectedForIntegration) GetLogs() []O11yIntegrationsCollectedLogAttribute`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *O11yDataCollectedForIntegration) GetLogsOk() (*[]O11yIntegrationsCollectedLogAttribute, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *O11yDataCollectedForIntegration) SetLogs(v []O11yIntegrationsCollectedLogAttribute)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *O11yDataCollectedForIntegration) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetMetrics

`func (o *O11yDataCollectedForIntegration) GetMetrics() []O11yIntegrationsCollectedMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *O11yDataCollectedForIntegration) GetMetricsOk() (*[]O11yIntegrationsCollectedMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *O11yDataCollectedForIntegration) SetMetrics(v []O11yIntegrationsCollectedMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *O11yDataCollectedForIntegration) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


