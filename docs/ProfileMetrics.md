# ProfileMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Funnel** | Pointer to [**Funnel**](Funnel.md) |  | [optional] 
**LaunchProgress** | Pointer to [**ProgressView**](ProgressView.md) |  | [optional] 
**Records** | Pointer to **int32** |  | [optional] 
**RevenueCents** | Pointer to **int32** |  | [optional] 

## Methods

### NewProfileMetrics

`func NewProfileMetrics() *ProfileMetrics`

NewProfileMetrics instantiates a new ProfileMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileMetricsWithDefaults

`func NewProfileMetricsWithDefaults() *ProfileMetrics`

NewProfileMetricsWithDefaults instantiates a new ProfileMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunnel

`func (o *ProfileMetrics) GetFunnel() Funnel`

GetFunnel returns the Funnel field if non-nil, zero value otherwise.

### GetFunnelOk

`func (o *ProfileMetrics) GetFunnelOk() (*Funnel, bool)`

GetFunnelOk returns a tuple with the Funnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnel

`func (o *ProfileMetrics) SetFunnel(v Funnel)`

SetFunnel sets Funnel field to given value.

### HasFunnel

`func (o *ProfileMetrics) HasFunnel() bool`

HasFunnel returns a boolean if a field has been set.

### GetLaunchProgress

`func (o *ProfileMetrics) GetLaunchProgress() ProgressView`

GetLaunchProgress returns the LaunchProgress field if non-nil, zero value otherwise.

### GetLaunchProgressOk

`func (o *ProfileMetrics) GetLaunchProgressOk() (*ProgressView, bool)`

GetLaunchProgressOk returns a tuple with the LaunchProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchProgress

`func (o *ProfileMetrics) SetLaunchProgress(v ProgressView)`

SetLaunchProgress sets LaunchProgress field to given value.

### HasLaunchProgress

`func (o *ProfileMetrics) HasLaunchProgress() bool`

HasLaunchProgress returns a boolean if a field has been set.

### GetRecords

`func (o *ProfileMetrics) GetRecords() int32`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *ProfileMetrics) GetRecordsOk() (*int32, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *ProfileMetrics) SetRecords(v int32)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *ProfileMetrics) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetRevenueCents

`func (o *ProfileMetrics) GetRevenueCents() int32`

GetRevenueCents returns the RevenueCents field if non-nil, zero value otherwise.

### GetRevenueCentsOk

`func (o *ProfileMetrics) GetRevenueCentsOk() (*int32, bool)`

GetRevenueCentsOk returns a tuple with the RevenueCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueCents

`func (o *ProfileMetrics) SetRevenueCents(v int32)`

SetRevenueCents sets RevenueCents field to given value.

### HasRevenueCents

`func (o *ProfileMetrics) HasRevenueCents() bool`

HasRevenueCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


