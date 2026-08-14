# SubsystemTotals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **int32** |  | [optional] 
**Enabled** | Pointer to **int32** |  | [optional] 
**ErrorRate** | Pointer to **float32** | percent (0..100) | [optional] 
**Errors** | Pointer to **int32** |  | [optional] 
**Reporting** | Pointer to **int32** | enabled AND served ≥1 traced request in the window | [optional] 
**Requests** | Pointer to **int32** |  | [optional] 
**Subsystems** | Pointer to **int32** |  | [optional] 

## Methods

### NewSubsystemTotals

`func NewSubsystemTotals() *SubsystemTotals`

NewSubsystemTotals instantiates a new SubsystemTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubsystemTotalsWithDefaults

`func NewSubsystemTotalsWithDefaults() *SubsystemTotals`

NewSubsystemTotalsWithDefaults instantiates a new SubsystemTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *SubsystemTotals) GetDisabled() int32`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SubsystemTotals) GetDisabledOk() (*int32, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SubsystemTotals) SetDisabled(v int32)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *SubsystemTotals) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetEnabled

`func (o *SubsystemTotals) GetEnabled() int32`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SubsystemTotals) GetEnabledOk() (*int32, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SubsystemTotals) SetEnabled(v int32)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SubsystemTotals) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetErrorRate

`func (o *SubsystemTotals) GetErrorRate() float32`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *SubsystemTotals) GetErrorRateOk() (*float32, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *SubsystemTotals) SetErrorRate(v float32)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *SubsystemTotals) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetErrors

`func (o *SubsystemTotals) GetErrors() int32`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SubsystemTotals) GetErrorsOk() (*int32, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SubsystemTotals) SetErrors(v int32)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SubsystemTotals) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetReporting

`func (o *SubsystemTotals) GetReporting() int32`

GetReporting returns the Reporting field if non-nil, zero value otherwise.

### GetReportingOk

`func (o *SubsystemTotals) GetReportingOk() (*int32, bool)`

GetReportingOk returns a tuple with the Reporting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporting

`func (o *SubsystemTotals) SetReporting(v int32)`

SetReporting sets Reporting field to given value.

### HasReporting

`func (o *SubsystemTotals) HasReporting() bool`

HasReporting returns a boolean if a field has been set.

### GetRequests

`func (o *SubsystemTotals) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *SubsystemTotals) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *SubsystemTotals) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *SubsystemTotals) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetSubsystems

`func (o *SubsystemTotals) GetSubsystems() int32`

GetSubsystems returns the Subsystems field if non-nil, zero value otherwise.

### GetSubsystemsOk

`func (o *SubsystemTotals) GetSubsystemsOk() (*int32, bool)`

GetSubsystemsOk returns a tuple with the Subsystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsystems

`func (o *SubsystemTotals) SetSubsystems(v int32)`

SetSubsystems sets Subsystems field to given value.

### HasSubsystems

`func (o *SubsystemTotals) HasSubsystems() bool`

HasSubsystems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


