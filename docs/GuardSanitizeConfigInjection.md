# GuardSanitizeConfigInjection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**BlockOnDetection** | Pointer to **bool** |  | [optional] [default to true]
**Sensitivity** | Pointer to **float32** | Detection threshold (0.0-1.0) | [optional] [default to 0.7]
**CustomPatterns** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGuardSanitizeConfigInjection

`func NewGuardSanitizeConfigInjection() *GuardSanitizeConfigInjection`

NewGuardSanitizeConfigInjection instantiates a new GuardSanitizeConfigInjection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardSanitizeConfigInjectionWithDefaults

`func NewGuardSanitizeConfigInjectionWithDefaults() *GuardSanitizeConfigInjection`

NewGuardSanitizeConfigInjectionWithDefaults instantiates a new GuardSanitizeConfigInjection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GuardSanitizeConfigInjection) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GuardSanitizeConfigInjection) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GuardSanitizeConfigInjection) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GuardSanitizeConfigInjection) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetBlockOnDetection

`func (o *GuardSanitizeConfigInjection) GetBlockOnDetection() bool`

GetBlockOnDetection returns the BlockOnDetection field if non-nil, zero value otherwise.

### GetBlockOnDetectionOk

`func (o *GuardSanitizeConfigInjection) GetBlockOnDetectionOk() (*bool, bool)`

GetBlockOnDetectionOk returns a tuple with the BlockOnDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockOnDetection

`func (o *GuardSanitizeConfigInjection) SetBlockOnDetection(v bool)`

SetBlockOnDetection sets BlockOnDetection field to given value.

### HasBlockOnDetection

`func (o *GuardSanitizeConfigInjection) HasBlockOnDetection() bool`

HasBlockOnDetection returns a boolean if a field has been set.

### GetSensitivity

`func (o *GuardSanitizeConfigInjection) GetSensitivity() float32`

GetSensitivity returns the Sensitivity field if non-nil, zero value otherwise.

### GetSensitivityOk

`func (o *GuardSanitizeConfigInjection) GetSensitivityOk() (*float32, bool)`

GetSensitivityOk returns a tuple with the Sensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitivity

`func (o *GuardSanitizeConfigInjection) SetSensitivity(v float32)`

SetSensitivity sets Sensitivity field to given value.

### HasSensitivity

`func (o *GuardSanitizeConfigInjection) HasSensitivity() bool`

HasSensitivity returns a boolean if a field has been set.

### GetCustomPatterns

`func (o *GuardSanitizeConfigInjection) GetCustomPatterns() []string`

GetCustomPatterns returns the CustomPatterns field if non-nil, zero value otherwise.

### GetCustomPatternsOk

`func (o *GuardSanitizeConfigInjection) GetCustomPatternsOk() (*[]string, bool)`

GetCustomPatternsOk returns a tuple with the CustomPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPatterns

`func (o *GuardSanitizeConfigInjection) SetCustomPatterns(v []string)`

SetCustomPatterns sets CustomPatterns field to given value.

### HasCustomPatterns

`func (o *GuardSanitizeConfigInjection) HasCustomPatterns() bool`

HasCustomPatterns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


