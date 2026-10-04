# GuideProfileMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Funnel** | Pointer to [**GuideFunnel**](GuideFunnel.md) | Funnel is what the org&#39;s analytics observed over the trailing window. Read its &#x60;available&#x60; first: an org with no analytics reports zeros here, and zero traffic and no measurement are different facts. | [optional] 
**LaunchProgress** | Pointer to [**GuideProgressView**](GuideProgressView.md) | LaunchProgress is the org&#39;s own position in the launch checklist, folded in so a profile carries both what the org has BUILT and what it has DONE. | [optional] 
**Records** | Pointer to **int64** | Records is how many business records the org holds — the volume that tells a real book of customers from an empty account. It feeds the &#x60;customers&#x60; signal, which crosses at a threshold rather than at one row. | [optional] 
**RevenueCents** | Pointer to **int64** | RevenueCents is the org&#39;s money OF RECORD — what its books say, in whole cents, never a float and never a display string. This is the number the scaling stage is decided on; funnel.revenue is the beacon&#39;s separate, unreconciled view of the same business. Zero when the org has none, and also zero when the books could not be read, which is why the &#x60;revenue&#x60; signal beside it is the thing to trust. | [optional] 

## Methods

### NewGuideProfileMetrics

`func NewGuideProfileMetrics() *GuideProfileMetrics`

NewGuideProfileMetrics instantiates a new GuideProfileMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuideProfileMetricsWithDefaults

`func NewGuideProfileMetricsWithDefaults() *GuideProfileMetrics`

NewGuideProfileMetricsWithDefaults instantiates a new GuideProfileMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunnel

`func (o *GuideProfileMetrics) GetFunnel() GuideFunnel`

GetFunnel returns the Funnel field if non-nil, zero value otherwise.

### GetFunnelOk

`func (o *GuideProfileMetrics) GetFunnelOk() (*GuideFunnel, bool)`

GetFunnelOk returns a tuple with the Funnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnel

`func (o *GuideProfileMetrics) SetFunnel(v GuideFunnel)`

SetFunnel sets Funnel field to given value.

### HasFunnel

`func (o *GuideProfileMetrics) HasFunnel() bool`

HasFunnel returns a boolean if a field has been set.

### GetLaunchProgress

`func (o *GuideProfileMetrics) GetLaunchProgress() GuideProgressView`

GetLaunchProgress returns the LaunchProgress field if non-nil, zero value otherwise.

### GetLaunchProgressOk

`func (o *GuideProfileMetrics) GetLaunchProgressOk() (*GuideProgressView, bool)`

GetLaunchProgressOk returns a tuple with the LaunchProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchProgress

`func (o *GuideProfileMetrics) SetLaunchProgress(v GuideProgressView)`

SetLaunchProgress sets LaunchProgress field to given value.

### HasLaunchProgress

`func (o *GuideProfileMetrics) HasLaunchProgress() bool`

HasLaunchProgress returns a boolean if a field has been set.

### GetRecords

`func (o *GuideProfileMetrics) GetRecords() int64`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *GuideProfileMetrics) GetRecordsOk() (*int64, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *GuideProfileMetrics) SetRecords(v int64)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *GuideProfileMetrics) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetRevenueCents

`func (o *GuideProfileMetrics) GetRevenueCents() int64`

GetRevenueCents returns the RevenueCents field if non-nil, zero value otherwise.

### GetRevenueCentsOk

`func (o *GuideProfileMetrics) GetRevenueCentsOk() (*int64, bool)`

GetRevenueCentsOk returns a tuple with the RevenueCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueCents

`func (o *GuideProfileMetrics) SetRevenueCents(v int64)`

SetRevenueCents sets RevenueCents field to given value.

### HasRevenueCents

`func (o *GuideProfileMetrics) HasRevenueCents() bool`

HasRevenueCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


