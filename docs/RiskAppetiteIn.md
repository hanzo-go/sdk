# RiskAppetiteIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Live** | Pointer to **bool** | Live turns the model out of shadow. It defaults to FALSE on every call, so going live is always an explicit act and never a side effect of changing a number.  Setting it requires an ADMIN of this organisation. Arming decides whether the model may change an outcome at all — a payment frozen, a grant refused — for every customer this organisation has, which is a governance act rather than a tuning one. Stating the appetite and the sample needs no admin. | [optional] 
**Review** | Pointer to **float64** | Review is the share of the stream that may be sent for examination, in (0, 0.5]. The alert threshold is derived from it as a quantile of the scores actually observed, so the level is governed rather than tuned. | [optional] 
**Sample** | Pointer to **float64** | Sample is the share of below-the-line events retained for review, in [0, 1]. It is the instrument that measures what the model missed; there are no labels, so nothing else can. | [optional] 

## Methods

### NewRiskAppetiteIn

`func NewRiskAppetiteIn() *RiskAppetiteIn`

NewRiskAppetiteIn instantiates a new RiskAppetiteIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAppetiteInWithDefaults

`func NewRiskAppetiteInWithDefaults() *RiskAppetiteIn`

NewRiskAppetiteInWithDefaults instantiates a new RiskAppetiteIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLive

`func (o *RiskAppetiteIn) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *RiskAppetiteIn) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *RiskAppetiteIn) SetLive(v bool)`

SetLive sets Live field to given value.

### HasLive

`func (o *RiskAppetiteIn) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetReview

`func (o *RiskAppetiteIn) GetReview() float64`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *RiskAppetiteIn) GetReviewOk() (*float64, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *RiskAppetiteIn) SetReview(v float64)`

SetReview sets Review field to given value.

### HasReview

`func (o *RiskAppetiteIn) HasReview() bool`

HasReview returns a boolean if a field has been set.

### GetSample

`func (o *RiskAppetiteIn) GetSample() float64`

GetSample returns the Sample field if non-nil, zero value otherwise.

### GetSampleOk

`func (o *RiskAppetiteIn) GetSampleOk() (*float64, bool)`

GetSampleOk returns a tuple with the Sample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSample

`func (o *RiskAppetiteIn) SetSample(v float64)`

SetSample sets Sample field to given value.

### HasSample

`func (o *RiskAppetiteIn) HasSample() bool`

HasSample returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


