# AnalyticsView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Funnel** | Pointer to [**Funnel**](Funnel.md) | Funnel is the org&#39;s trailing-30-day traffic → signups → orders from the shared analytics warehouse; available is false when it has emitted nothing. | [optional] 
**Recommendations** | Pointer to **[]string** | Recommendations are the next-best GTM actions derived from that funnel. | [optional] 

## Methods

### NewAnalyticsView

`func NewAnalyticsView() *AnalyticsView`

NewAnalyticsView instantiates a new AnalyticsView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsViewWithDefaults

`func NewAnalyticsViewWithDefaults() *AnalyticsView`

NewAnalyticsViewWithDefaults instantiates a new AnalyticsView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunnel

`func (o *AnalyticsView) GetFunnel() Funnel`

GetFunnel returns the Funnel field if non-nil, zero value otherwise.

### GetFunnelOk

`func (o *AnalyticsView) GetFunnelOk() (*Funnel, bool)`

GetFunnelOk returns a tuple with the Funnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnel

`func (o *AnalyticsView) SetFunnel(v Funnel)`

SetFunnel sets Funnel field to given value.

### HasFunnel

`func (o *AnalyticsView) HasFunnel() bool`

HasFunnel returns a boolean if a field has been set.

### GetRecommendations

`func (o *AnalyticsView) GetRecommendations() []string`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *AnalyticsView) GetRecommendationsOk() (*[]string, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *AnalyticsView) SetRecommendations(v []string)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *AnalyticsView) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


