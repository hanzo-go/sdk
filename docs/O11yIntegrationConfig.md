# O11yIntegrationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnabledRegions** | Pointer to **[]string** | backward compatible | [optional] 
**Telemetry** | Pointer to [**O11yOldAWSCollectionStrategy**](O11yOldAWSCollectionStrategy.md) | backward compatible | [optional] 

## Methods

### NewO11yIntegrationConfig

`func NewO11yIntegrationConfig() *O11yIntegrationConfig`

NewO11yIntegrationConfig instantiates a new O11yIntegrationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yIntegrationConfigWithDefaults

`func NewO11yIntegrationConfigWithDefaults() *O11yIntegrationConfig`

NewO11yIntegrationConfigWithDefaults instantiates a new O11yIntegrationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabledRegions

`func (o *O11yIntegrationConfig) GetEnabledRegions() []string`

GetEnabledRegions returns the EnabledRegions field if non-nil, zero value otherwise.

### GetEnabledRegionsOk

`func (o *O11yIntegrationConfig) GetEnabledRegionsOk() (*[]string, bool)`

GetEnabledRegionsOk returns a tuple with the EnabledRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledRegions

`func (o *O11yIntegrationConfig) SetEnabledRegions(v []string)`

SetEnabledRegions sets EnabledRegions field to given value.

### HasEnabledRegions

`func (o *O11yIntegrationConfig) HasEnabledRegions() bool`

HasEnabledRegions returns a boolean if a field has been set.

### GetTelemetry

`func (o *O11yIntegrationConfig) GetTelemetry() O11yOldAWSCollectionStrategy`

GetTelemetry returns the Telemetry field if non-nil, zero value otherwise.

### GetTelemetryOk

`func (o *O11yIntegrationConfig) GetTelemetryOk() (*O11yOldAWSCollectionStrategy, bool)`

GetTelemetryOk returns a tuple with the Telemetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetry

`func (o *O11yIntegrationConfig) SetTelemetry(v O11yOldAWSCollectionStrategy)`

SetTelemetry sets Telemetry field to given value.

### HasTelemetry

`func (o *O11yIntegrationConfig) HasTelemetry() bool`

HasTelemetry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


