# CloudProfileMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Funnel** | Pointer to [**CloudFunnel**](CloudFunnel.md) |  | [optional] 
**LaunchProgress** | Pointer to [**CloudProgressView**](CloudProgressView.md) |  | [optional] 
**Records** | Pointer to **int32** |  | [optional] 
**RevenueCents** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudProfileMetrics

`func NewCloudProfileMetrics() *CloudProfileMetrics`

NewCloudProfileMetrics instantiates a new CloudProfileMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProfileMetricsWithDefaults

`func NewCloudProfileMetricsWithDefaults() *CloudProfileMetrics`

NewCloudProfileMetricsWithDefaults instantiates a new CloudProfileMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunnel

`func (o *CloudProfileMetrics) GetFunnel() CloudFunnel`

GetFunnel returns the Funnel field if non-nil, zero value otherwise.

### GetFunnelOk

`func (o *CloudProfileMetrics) GetFunnelOk() (*CloudFunnel, bool)`

GetFunnelOk returns a tuple with the Funnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnel

`func (o *CloudProfileMetrics) SetFunnel(v CloudFunnel)`

SetFunnel sets Funnel field to given value.

### HasFunnel

`func (o *CloudProfileMetrics) HasFunnel() bool`

HasFunnel returns a boolean if a field has been set.

### GetLaunchProgress

`func (o *CloudProfileMetrics) GetLaunchProgress() CloudProgressView`

GetLaunchProgress returns the LaunchProgress field if non-nil, zero value otherwise.

### GetLaunchProgressOk

`func (o *CloudProfileMetrics) GetLaunchProgressOk() (*CloudProgressView, bool)`

GetLaunchProgressOk returns a tuple with the LaunchProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchProgress

`func (o *CloudProfileMetrics) SetLaunchProgress(v CloudProgressView)`

SetLaunchProgress sets LaunchProgress field to given value.

### HasLaunchProgress

`func (o *CloudProfileMetrics) HasLaunchProgress() bool`

HasLaunchProgress returns a boolean if a field has been set.

### GetRecords

`func (o *CloudProfileMetrics) GetRecords() int32`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *CloudProfileMetrics) GetRecordsOk() (*int32, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *CloudProfileMetrics) SetRecords(v int32)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *CloudProfileMetrics) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetRevenueCents

`func (o *CloudProfileMetrics) GetRevenueCents() int32`

GetRevenueCents returns the RevenueCents field if non-nil, zero value otherwise.

### GetRevenueCentsOk

`func (o *CloudProfileMetrics) GetRevenueCentsOk() (*int32, bool)`

GetRevenueCentsOk returns a tuple with the RevenueCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueCents

`func (o *CloudProfileMetrics) SetRevenueCents(v int32)`

SetRevenueCents sets RevenueCents field to given value.

### HasRevenueCents

`func (o *CloudProfileMetrics) HasRevenueCents() bool`

HasRevenueCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


