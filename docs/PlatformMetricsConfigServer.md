# PlatformMetricsConfigServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefreshRate** | Pointer to **int32** |  | [optional] 
**RetentionDays** | Pointer to **int32** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**UrlCallback** | Pointer to **string** |  | [optional] 
**CronJob** | Pointer to **string** |  | [optional] 
**Thresholds** | Pointer to [**PlatformMetricsConfigServerThresholds**](PlatformMetricsConfigServerThresholds.md) |  | [optional] 

## Methods

### NewPlatformMetricsConfigServer

`func NewPlatformMetricsConfigServer() *PlatformMetricsConfigServer`

NewPlatformMetricsConfigServer instantiates a new PlatformMetricsConfigServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformMetricsConfigServerWithDefaults

`func NewPlatformMetricsConfigServerWithDefaults() *PlatformMetricsConfigServer`

NewPlatformMetricsConfigServerWithDefaults instantiates a new PlatformMetricsConfigServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefreshRate

`func (o *PlatformMetricsConfigServer) GetRefreshRate() int32`

GetRefreshRate returns the RefreshRate field if non-nil, zero value otherwise.

### GetRefreshRateOk

`func (o *PlatformMetricsConfigServer) GetRefreshRateOk() (*int32, bool)`

GetRefreshRateOk returns a tuple with the RefreshRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshRate

`func (o *PlatformMetricsConfigServer) SetRefreshRate(v int32)`

SetRefreshRate sets RefreshRate field to given value.

### HasRefreshRate

`func (o *PlatformMetricsConfigServer) HasRefreshRate() bool`

HasRefreshRate returns a boolean if a field has been set.

### GetRetentionDays

`func (o *PlatformMetricsConfigServer) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *PlatformMetricsConfigServer) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *PlatformMetricsConfigServer) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.

### HasRetentionDays

`func (o *PlatformMetricsConfigServer) HasRetentionDays() bool`

HasRetentionDays returns a boolean if a field has been set.

### GetPort

`func (o *PlatformMetricsConfigServer) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PlatformMetricsConfigServer) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PlatformMetricsConfigServer) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PlatformMetricsConfigServer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetToken

`func (o *PlatformMetricsConfigServer) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PlatformMetricsConfigServer) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PlatformMetricsConfigServer) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PlatformMetricsConfigServer) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUrlCallback

`func (o *PlatformMetricsConfigServer) GetUrlCallback() string`

GetUrlCallback returns the UrlCallback field if non-nil, zero value otherwise.

### GetUrlCallbackOk

`func (o *PlatformMetricsConfigServer) GetUrlCallbackOk() (*string, bool)`

GetUrlCallbackOk returns a tuple with the UrlCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlCallback

`func (o *PlatformMetricsConfigServer) SetUrlCallback(v string)`

SetUrlCallback sets UrlCallback field to given value.

### HasUrlCallback

`func (o *PlatformMetricsConfigServer) HasUrlCallback() bool`

HasUrlCallback returns a boolean if a field has been set.

### GetCronJob

`func (o *PlatformMetricsConfigServer) GetCronJob() string`

GetCronJob returns the CronJob field if non-nil, zero value otherwise.

### GetCronJobOk

`func (o *PlatformMetricsConfigServer) GetCronJobOk() (*string, bool)`

GetCronJobOk returns a tuple with the CronJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronJob

`func (o *PlatformMetricsConfigServer) SetCronJob(v string)`

SetCronJob sets CronJob field to given value.

### HasCronJob

`func (o *PlatformMetricsConfigServer) HasCronJob() bool`

HasCronJob returns a boolean if a field has been set.

### GetThresholds

`func (o *PlatformMetricsConfigServer) GetThresholds() PlatformMetricsConfigServerThresholds`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *PlatformMetricsConfigServer) GetThresholdsOk() (*PlatformMetricsConfigServerThresholds, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *PlatformMetricsConfigServer) SetThresholds(v PlatformMetricsConfigServerThresholds)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *PlatformMetricsConfigServer) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


