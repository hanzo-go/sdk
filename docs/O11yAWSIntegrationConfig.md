# O11yAWSIntegrationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnabledRegions** | Pointer to **[]string** |  | [optional] 
**TelemetryCollectionStrategy** | Pointer to [**O11yAWSTelemetryCollectionStrategy**](O11yAWSTelemetryCollectionStrategy.md) |  | [optional] 

## Methods

### NewO11yAWSIntegrationConfig

`func NewO11yAWSIntegrationConfig() *O11yAWSIntegrationConfig`

NewO11yAWSIntegrationConfig instantiates a new O11yAWSIntegrationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yAWSIntegrationConfigWithDefaults

`func NewO11yAWSIntegrationConfigWithDefaults() *O11yAWSIntegrationConfig`

NewO11yAWSIntegrationConfigWithDefaults instantiates a new O11yAWSIntegrationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabledRegions

`func (o *O11yAWSIntegrationConfig) GetEnabledRegions() []string`

GetEnabledRegions returns the EnabledRegions field if non-nil, zero value otherwise.

### GetEnabledRegionsOk

`func (o *O11yAWSIntegrationConfig) GetEnabledRegionsOk() (*[]string, bool)`

GetEnabledRegionsOk returns a tuple with the EnabledRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledRegions

`func (o *O11yAWSIntegrationConfig) SetEnabledRegions(v []string)`

SetEnabledRegions sets EnabledRegions field to given value.

### HasEnabledRegions

`func (o *O11yAWSIntegrationConfig) HasEnabledRegions() bool`

HasEnabledRegions returns a boolean if a field has been set.

### GetTelemetryCollectionStrategy

`func (o *O11yAWSIntegrationConfig) GetTelemetryCollectionStrategy() O11yAWSTelemetryCollectionStrategy`

GetTelemetryCollectionStrategy returns the TelemetryCollectionStrategy field if non-nil, zero value otherwise.

### GetTelemetryCollectionStrategyOk

`func (o *O11yAWSIntegrationConfig) GetTelemetryCollectionStrategyOk() (*O11yAWSTelemetryCollectionStrategy, bool)`

GetTelemetryCollectionStrategyOk returns a tuple with the TelemetryCollectionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetryCollectionStrategy

`func (o *O11yAWSIntegrationConfig) SetTelemetryCollectionStrategy(v O11yAWSTelemetryCollectionStrategy)`

SetTelemetryCollectionStrategy sets TelemetryCollectionStrategy field to given value.

### HasTelemetryCollectionStrategy

`func (o *O11yAWSIntegrationConfig) HasTelemetryCollectionStrategy() bool`

HasTelemetryCollectionStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


