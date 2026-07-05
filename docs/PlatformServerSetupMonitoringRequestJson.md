# PlatformServerSetupMonitoringRequestJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | **string** |  | 
**MetricsConfig** | [**PlatformMetricsConfig**](PlatformMetricsConfig.md) |  | 

## Methods

### NewPlatformServerSetupMonitoringRequestJson

`func NewPlatformServerSetupMonitoringRequestJson(serverId string, metricsConfig PlatformMetricsConfig, ) *PlatformServerSetupMonitoringRequestJson`

NewPlatformServerSetupMonitoringRequestJson instantiates a new PlatformServerSetupMonitoringRequestJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformServerSetupMonitoringRequestJsonWithDefaults

`func NewPlatformServerSetupMonitoringRequestJsonWithDefaults() *PlatformServerSetupMonitoringRequestJson`

NewPlatformServerSetupMonitoringRequestJsonWithDefaults instantiates a new PlatformServerSetupMonitoringRequestJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *PlatformServerSetupMonitoringRequestJson) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *PlatformServerSetupMonitoringRequestJson) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *PlatformServerSetupMonitoringRequestJson) SetServerId(v string)`

SetServerId sets ServerId field to given value.


### GetMetricsConfig

`func (o *PlatformServerSetupMonitoringRequestJson) GetMetricsConfig() PlatformMetricsConfig`

GetMetricsConfig returns the MetricsConfig field if non-nil, zero value otherwise.

### GetMetricsConfigOk

`func (o *PlatformServerSetupMonitoringRequestJson) GetMetricsConfigOk() (*PlatformMetricsConfig, bool)`

GetMetricsConfigOk returns a tuple with the MetricsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsConfig

`func (o *PlatformServerSetupMonitoringRequestJson) SetMetricsConfig(v PlatformMetricsConfig)`

SetMetricsConfig sets MetricsConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


