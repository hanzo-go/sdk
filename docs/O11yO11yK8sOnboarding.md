# O11yO11yK8sOnboarding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DidSendClusterMetrics** | Pointer to **bool** | DidSendClusterMetrics says whether cluster metrics have arrived. | [optional] 
**DidSendNodeMetrics** | Pointer to **bool** | DidSendNodeMetrics says whether node metrics have arrived. | [optional] 
**DidSendPodMetrics** | Pointer to **bool** | DidSendPodMetrics says whether pod metrics have arrived. | [optional] 
**IsSendingOptionalPodMetrics** | Pointer to **bool** | IsSendingOptionalPodMetrics says whether optional pod metrics are flowing. | [optional] 
**IsSendingRequiredMetadata** | Pointer to [**[]O11yO11yPodOnboarding**](O11yO11yPodOnboarding.md) | IsSendingRequiredMetadata reports, per pod, which required metadata labels are present. | [optional] 

## Methods

### NewO11yO11yK8sOnboarding

`func NewO11yO11yK8sOnboarding() *O11yO11yK8sOnboarding`

NewO11yO11yK8sOnboarding instantiates a new O11yO11yK8sOnboarding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yK8sOnboardingWithDefaults

`func NewO11yO11yK8sOnboardingWithDefaults() *O11yO11yK8sOnboarding`

NewO11yO11yK8sOnboardingWithDefaults instantiates a new O11yO11yK8sOnboarding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDidSendClusterMetrics

`func (o *O11yO11yK8sOnboarding) GetDidSendClusterMetrics() bool`

GetDidSendClusterMetrics returns the DidSendClusterMetrics field if non-nil, zero value otherwise.

### GetDidSendClusterMetricsOk

`func (o *O11yO11yK8sOnboarding) GetDidSendClusterMetricsOk() (*bool, bool)`

GetDidSendClusterMetricsOk returns a tuple with the DidSendClusterMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDidSendClusterMetrics

`func (o *O11yO11yK8sOnboarding) SetDidSendClusterMetrics(v bool)`

SetDidSendClusterMetrics sets DidSendClusterMetrics field to given value.

### HasDidSendClusterMetrics

`func (o *O11yO11yK8sOnboarding) HasDidSendClusterMetrics() bool`

HasDidSendClusterMetrics returns a boolean if a field has been set.

### GetDidSendNodeMetrics

`func (o *O11yO11yK8sOnboarding) GetDidSendNodeMetrics() bool`

GetDidSendNodeMetrics returns the DidSendNodeMetrics field if non-nil, zero value otherwise.

### GetDidSendNodeMetricsOk

`func (o *O11yO11yK8sOnboarding) GetDidSendNodeMetricsOk() (*bool, bool)`

GetDidSendNodeMetricsOk returns a tuple with the DidSendNodeMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDidSendNodeMetrics

`func (o *O11yO11yK8sOnboarding) SetDidSendNodeMetrics(v bool)`

SetDidSendNodeMetrics sets DidSendNodeMetrics field to given value.

### HasDidSendNodeMetrics

`func (o *O11yO11yK8sOnboarding) HasDidSendNodeMetrics() bool`

HasDidSendNodeMetrics returns a boolean if a field has been set.

### GetDidSendPodMetrics

`func (o *O11yO11yK8sOnboarding) GetDidSendPodMetrics() bool`

GetDidSendPodMetrics returns the DidSendPodMetrics field if non-nil, zero value otherwise.

### GetDidSendPodMetricsOk

`func (o *O11yO11yK8sOnboarding) GetDidSendPodMetricsOk() (*bool, bool)`

GetDidSendPodMetricsOk returns a tuple with the DidSendPodMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDidSendPodMetrics

`func (o *O11yO11yK8sOnboarding) SetDidSendPodMetrics(v bool)`

SetDidSendPodMetrics sets DidSendPodMetrics field to given value.

### HasDidSendPodMetrics

`func (o *O11yO11yK8sOnboarding) HasDidSendPodMetrics() bool`

HasDidSendPodMetrics returns a boolean if a field has been set.

### GetIsSendingOptionalPodMetrics

`func (o *O11yO11yK8sOnboarding) GetIsSendingOptionalPodMetrics() bool`

GetIsSendingOptionalPodMetrics returns the IsSendingOptionalPodMetrics field if non-nil, zero value otherwise.

### GetIsSendingOptionalPodMetricsOk

`func (o *O11yO11yK8sOnboarding) GetIsSendingOptionalPodMetricsOk() (*bool, bool)`

GetIsSendingOptionalPodMetricsOk returns a tuple with the IsSendingOptionalPodMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSendingOptionalPodMetrics

`func (o *O11yO11yK8sOnboarding) SetIsSendingOptionalPodMetrics(v bool)`

SetIsSendingOptionalPodMetrics sets IsSendingOptionalPodMetrics field to given value.

### HasIsSendingOptionalPodMetrics

`func (o *O11yO11yK8sOnboarding) HasIsSendingOptionalPodMetrics() bool`

HasIsSendingOptionalPodMetrics returns a boolean if a field has been set.

### GetIsSendingRequiredMetadata

`func (o *O11yO11yK8sOnboarding) GetIsSendingRequiredMetadata() []O11yO11yPodOnboarding`

GetIsSendingRequiredMetadata returns the IsSendingRequiredMetadata field if non-nil, zero value otherwise.

### GetIsSendingRequiredMetadataOk

`func (o *O11yO11yK8sOnboarding) GetIsSendingRequiredMetadataOk() (*[]O11yO11yPodOnboarding, bool)`

GetIsSendingRequiredMetadataOk returns a tuple with the IsSendingRequiredMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSendingRequiredMetadata

`func (o *O11yO11yK8sOnboarding) SetIsSendingRequiredMetadata(v []O11yO11yPodOnboarding)`

SetIsSendingRequiredMetadata sets IsSendingRequiredMetadata field to given value.

### HasIsSendingRequiredMetadata

`func (o *O11yO11yK8sOnboarding) HasIsSendingRequiredMetadata() bool`

HasIsSendingRequiredMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


