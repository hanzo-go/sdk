# CloudSubsystemTotals

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

### NewCloudSubsystemTotals

`func NewCloudSubsystemTotals() *CloudSubsystemTotals`

NewCloudSubsystemTotals instantiates a new CloudSubsystemTotals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSubsystemTotalsWithDefaults

`func NewCloudSubsystemTotalsWithDefaults() *CloudSubsystemTotals`

NewCloudSubsystemTotalsWithDefaults instantiates a new CloudSubsystemTotals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *CloudSubsystemTotals) GetDisabled() int32`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *CloudSubsystemTotals) GetDisabledOk() (*int32, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *CloudSubsystemTotals) SetDisabled(v int32)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *CloudSubsystemTotals) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetEnabled

`func (o *CloudSubsystemTotals) GetEnabled() int32`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CloudSubsystemTotals) GetEnabledOk() (*int32, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CloudSubsystemTotals) SetEnabled(v int32)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CloudSubsystemTotals) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetErrorRate

`func (o *CloudSubsystemTotals) GetErrorRate() float32`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *CloudSubsystemTotals) GetErrorRateOk() (*float32, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *CloudSubsystemTotals) SetErrorRate(v float32)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *CloudSubsystemTotals) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetErrors

`func (o *CloudSubsystemTotals) GetErrors() int32`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CloudSubsystemTotals) GetErrorsOk() (*int32, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CloudSubsystemTotals) SetErrors(v int32)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CloudSubsystemTotals) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetReporting

`func (o *CloudSubsystemTotals) GetReporting() int32`

GetReporting returns the Reporting field if non-nil, zero value otherwise.

### GetReportingOk

`func (o *CloudSubsystemTotals) GetReportingOk() (*int32, bool)`

GetReportingOk returns a tuple with the Reporting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporting

`func (o *CloudSubsystemTotals) SetReporting(v int32)`

SetReporting sets Reporting field to given value.

### HasReporting

`func (o *CloudSubsystemTotals) HasReporting() bool`

HasReporting returns a boolean if a field has been set.

### GetRequests

`func (o *CloudSubsystemTotals) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *CloudSubsystemTotals) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *CloudSubsystemTotals) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *CloudSubsystemTotals) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetSubsystems

`func (o *CloudSubsystemTotals) GetSubsystems() int32`

GetSubsystems returns the Subsystems field if non-nil, zero value otherwise.

### GetSubsystemsOk

`func (o *CloudSubsystemTotals) GetSubsystemsOk() (*int32, bool)`

GetSubsystemsOk returns a tuple with the Subsystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsystems

`func (o *CloudSubsystemTotals) SetSubsystems(v int32)`

SetSubsystems sets Subsystems field to given value.

### HasSubsystems

`func (o *CloudSubsystemTotals) HasSubsystems() bool`

HasSubsystems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


