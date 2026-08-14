# O11yAzureTelemetryCollectionStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**O11yAzureLogsCollectionStrategy**](O11yAzureLogsCollectionStrategy.md) |  | [optional] 
**Metrics** | Pointer to **map[string]interface{}** |  | [optional] 
**ResourceProvider** | Pointer to **string** | https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/resource-providers-and-types | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yAzureTelemetryCollectionStrategy

`func NewO11yAzureTelemetryCollectionStrategy() *O11yAzureTelemetryCollectionStrategy`

NewO11yAzureTelemetryCollectionStrategy instantiates a new O11yAzureTelemetryCollectionStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAzureTelemetryCollectionStrategyWithDefaults

`func NewO11yAzureTelemetryCollectionStrategyWithDefaults() *O11yAzureTelemetryCollectionStrategy`

NewO11yAzureTelemetryCollectionStrategyWithDefaults instantiates a new O11yAzureTelemetryCollectionStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *O11yAzureTelemetryCollectionStrategy) GetLogs() O11yAzureLogsCollectionStrategy`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *O11yAzureTelemetryCollectionStrategy) GetLogsOk() (*O11yAzureLogsCollectionStrategy, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *O11yAzureTelemetryCollectionStrategy) SetLogs(v O11yAzureLogsCollectionStrategy)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *O11yAzureTelemetryCollectionStrategy) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetMetrics

`func (o *O11yAzureTelemetryCollectionStrategy) GetMetrics() map[string]interface{}`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *O11yAzureTelemetryCollectionStrategy) GetMetricsOk() (*map[string]interface{}, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *O11yAzureTelemetryCollectionStrategy) SetMetrics(v map[string]interface{})`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *O11yAzureTelemetryCollectionStrategy) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetResourceProvider

`func (o *O11yAzureTelemetryCollectionStrategy) GetResourceProvider() string`

GetResourceProvider returns the ResourceProvider field if non-nil, zero value otherwise.

### GetResourceProviderOk

`func (o *O11yAzureTelemetryCollectionStrategy) GetResourceProviderOk() (*string, bool)`

GetResourceProviderOk returns a tuple with the ResourceProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceProvider

`func (o *O11yAzureTelemetryCollectionStrategy) SetResourceProvider(v string)`

SetResourceProvider sets ResourceProvider field to given value.

### HasResourceProvider

`func (o *O11yAzureTelemetryCollectionStrategy) HasResourceProvider() bool`

HasResourceProvider returns a boolean if a field has been set.

### GetResourceType

`func (o *O11yAzureTelemetryCollectionStrategy) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *O11yAzureTelemetryCollectionStrategy) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *O11yAzureTelemetryCollectionStrategy) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *O11yAzureTelemetryCollectionStrategy) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


