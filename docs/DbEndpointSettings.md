# DbEndpointSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoscalingLimitMinCu** | Pointer to **float32** | Minimum compute units (0.25 CU &#x3D; 0.25 vCPU, 1 GB RAM) | [optional] [default to 0.25]
**AutoscalingLimitMaxCu** | Pointer to **float32** | Maximum compute units | [optional] [default to 4]
**SuspendTimeoutSeconds** | Pointer to **int32** | Seconds of inactivity before suspending (0 &#x3D; never) | [optional] [default to 300]
**PgSettings** | Pointer to **map[string]string** | PostgreSQL configuration overrides | [optional] 

## Methods

### NewDbEndpointSettings

`func NewDbEndpointSettings() *DbEndpointSettings`

NewDbEndpointSettings instantiates a new DbEndpointSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbEndpointSettingsWithDefaults

`func NewDbEndpointSettingsWithDefaults() *DbEndpointSettings`

NewDbEndpointSettingsWithDefaults instantiates a new DbEndpointSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoscalingLimitMinCu

`func (o *DbEndpointSettings) GetAutoscalingLimitMinCu() float32`

GetAutoscalingLimitMinCu returns the AutoscalingLimitMinCu field if non-nil, zero value otherwise.

### GetAutoscalingLimitMinCuOk

`func (o *DbEndpointSettings) GetAutoscalingLimitMinCuOk() (*float32, bool)`

GetAutoscalingLimitMinCuOk returns a tuple with the AutoscalingLimitMinCu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingLimitMinCu

`func (o *DbEndpointSettings) SetAutoscalingLimitMinCu(v float32)`

SetAutoscalingLimitMinCu sets AutoscalingLimitMinCu field to given value.

### HasAutoscalingLimitMinCu

`func (o *DbEndpointSettings) HasAutoscalingLimitMinCu() bool`

HasAutoscalingLimitMinCu returns a boolean if a field has been set.

### GetAutoscalingLimitMaxCu

`func (o *DbEndpointSettings) GetAutoscalingLimitMaxCu() float32`

GetAutoscalingLimitMaxCu returns the AutoscalingLimitMaxCu field if non-nil, zero value otherwise.

### GetAutoscalingLimitMaxCuOk

`func (o *DbEndpointSettings) GetAutoscalingLimitMaxCuOk() (*float32, bool)`

GetAutoscalingLimitMaxCuOk returns a tuple with the AutoscalingLimitMaxCu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalingLimitMaxCu

`func (o *DbEndpointSettings) SetAutoscalingLimitMaxCu(v float32)`

SetAutoscalingLimitMaxCu sets AutoscalingLimitMaxCu field to given value.

### HasAutoscalingLimitMaxCu

`func (o *DbEndpointSettings) HasAutoscalingLimitMaxCu() bool`

HasAutoscalingLimitMaxCu returns a boolean if a field has been set.

### GetSuspendTimeoutSeconds

`func (o *DbEndpointSettings) GetSuspendTimeoutSeconds() int32`

GetSuspendTimeoutSeconds returns the SuspendTimeoutSeconds field if non-nil, zero value otherwise.

### GetSuspendTimeoutSecondsOk

`func (o *DbEndpointSettings) GetSuspendTimeoutSecondsOk() (*int32, bool)`

GetSuspendTimeoutSecondsOk returns a tuple with the SuspendTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendTimeoutSeconds

`func (o *DbEndpointSettings) SetSuspendTimeoutSeconds(v int32)`

SetSuspendTimeoutSeconds sets SuspendTimeoutSeconds field to given value.

### HasSuspendTimeoutSeconds

`func (o *DbEndpointSettings) HasSuspendTimeoutSeconds() bool`

HasSuspendTimeoutSeconds returns a boolean if a field has been set.

### GetPgSettings

`func (o *DbEndpointSettings) GetPgSettings() map[string]string`

GetPgSettings returns the PgSettings field if non-nil, zero value otherwise.

### GetPgSettingsOk

`func (o *DbEndpointSettings) GetPgSettingsOk() (*map[string]string, bool)`

GetPgSettingsOk returns a tuple with the PgSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgSettings

`func (o *DbEndpointSettings) SetPgSettings(v map[string]string)`

SetPgSettings sets PgSettings field to given value.

### HasPgSettings

`func (o *DbEndpointSettings) HasPgSettings() bool`

HasPgSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


