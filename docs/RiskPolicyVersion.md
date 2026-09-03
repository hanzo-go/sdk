# RiskPolicyVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is when it entered force, RFC 3339, from the server clock. | [optional] 
**By** | Pointer to **string** | By is the identity that stated it, stamped server-side from the validated principal at the moment it entered force. | [optional] 
**Live** | Pointer to **bool** | Live is whether the model was permitted to change an outcome under it. | [optional] 
**Review** | Pointer to **float64** | Review is the share of the stream the regime states may be examined. The threshold in force is derived from it, which is why a decision is only defensible against the version that produced it. | [optional] 
**Sample** | Pointer to **float64** | Sample is the share of below-the-line events the regime retains for review. | [optional] 
**Version** | Pointer to **int64** | Version names this regime in this organisation&#39;s history. | [optional] 

## Methods

### NewRiskPolicyVersion

`func NewRiskPolicyVersion() *RiskPolicyVersion`

NewRiskPolicyVersion instantiates a new RiskPolicyVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPolicyVersionWithDefaults

`func NewRiskPolicyVersionWithDefaults() *RiskPolicyVersion`

NewRiskPolicyVersionWithDefaults instantiates a new RiskPolicyVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *RiskPolicyVersion) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *RiskPolicyVersion) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *RiskPolicyVersion) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *RiskPolicyVersion) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetBy

`func (o *RiskPolicyVersion) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *RiskPolicyVersion) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *RiskPolicyVersion) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *RiskPolicyVersion) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetLive

`func (o *RiskPolicyVersion) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *RiskPolicyVersion) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *RiskPolicyVersion) SetLive(v bool)`

SetLive sets Live field to given value.

### HasLive

`func (o *RiskPolicyVersion) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetReview

`func (o *RiskPolicyVersion) GetReview() float64`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *RiskPolicyVersion) GetReviewOk() (*float64, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *RiskPolicyVersion) SetReview(v float64)`

SetReview sets Review field to given value.

### HasReview

`func (o *RiskPolicyVersion) HasReview() bool`

HasReview returns a boolean if a field has been set.

### GetSample

`func (o *RiskPolicyVersion) GetSample() float64`

GetSample returns the Sample field if non-nil, zero value otherwise.

### GetSampleOk

`func (o *RiskPolicyVersion) GetSampleOk() (*float64, bool)`

GetSampleOk returns a tuple with the Sample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSample

`func (o *RiskPolicyVersion) SetSample(v float64)`

SetSample sets Sample field to given value.

### HasSample

`func (o *RiskPolicyVersion) HasSample() bool`

HasSample returns a boolean if a field has been set.

### GetVersion

`func (o *RiskPolicyVersion) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RiskPolicyVersion) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RiskPolicyVersion) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RiskPolicyVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


