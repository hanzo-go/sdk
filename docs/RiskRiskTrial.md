# RiskRiskTrial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerted** | Pointer to **int64** | Alerted is how many of those it would have raised. | [optional] 
**Curve** | Pointer to **[]float64** | Curve is the realised alert rate over successive tenths of the history — the learning curve, which says whether the shape settled or is still moving. | [optional] 
**Fit** | Pointer to **float64** | Fit ranks the shape, smaller being better: the relative miss of the stated appetite, plus flat penalties for never warming and for saturating, plus the share of coordinates that were blind. | [optional] 
**Learned** | Pointer to **int64** | Learned is how many events the shape learned from during the replay. | [optional] 
**Realised** | Pointer to **float64** | Realised is what that appetite actually produced. The distance between the two is what the search is searching over. | [optional] 
**Saturated** | Pointer to **bool** | Saturated is whether the appetite could not be honoured by any threshold, which is a shape that alerts on nothing and reads like a quiet one. | [optional] 
**Scored** | Pointer to **int64** | Scored is how many it was able to score. | [optional] 
**Stated** | Pointer to **float64** | Stated is the appetite the shape was tried at. | [optional] 
**Topology** | Pointer to [**RiskRiskTopology**](RiskRiskTopology.md) | Topology is the shape. | [optional] 
**Warm** | Pointer to **bool** | Warm is whether the shape learned enough to have an opinion at all over this organisation&#39;s whole history. | [optional] 

## Methods

### NewRiskRiskTrial

`func NewRiskRiskTrial() *RiskRiskTrial`

NewRiskRiskTrial instantiates a new RiskRiskTrial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRiskTrialWithDefaults

`func NewRiskRiskTrialWithDefaults() *RiskRiskTrial`

NewRiskRiskTrialWithDefaults instantiates a new RiskRiskTrial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerted

`func (o *RiskRiskTrial) GetAlerted() int64`

GetAlerted returns the Alerted field if non-nil, zero value otherwise.

### GetAlertedOk

`func (o *RiskRiskTrial) GetAlertedOk() (*int64, bool)`

GetAlertedOk returns a tuple with the Alerted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerted

`func (o *RiskRiskTrial) SetAlerted(v int64)`

SetAlerted sets Alerted field to given value.

### HasAlerted

`func (o *RiskRiskTrial) HasAlerted() bool`

HasAlerted returns a boolean if a field has been set.

### GetCurve

`func (o *RiskRiskTrial) GetCurve() []float64`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *RiskRiskTrial) GetCurveOk() (*[]float64, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *RiskRiskTrial) SetCurve(v []float64)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *RiskRiskTrial) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### GetFit

`func (o *RiskRiskTrial) GetFit() float64`

GetFit returns the Fit field if non-nil, zero value otherwise.

### GetFitOk

`func (o *RiskRiskTrial) GetFitOk() (*float64, bool)`

GetFitOk returns a tuple with the Fit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFit

`func (o *RiskRiskTrial) SetFit(v float64)`

SetFit sets Fit field to given value.

### HasFit

`func (o *RiskRiskTrial) HasFit() bool`

HasFit returns a boolean if a field has been set.

### GetLearned

`func (o *RiskRiskTrial) GetLearned() int64`

GetLearned returns the Learned field if non-nil, zero value otherwise.

### GetLearnedOk

`func (o *RiskRiskTrial) GetLearnedOk() (*int64, bool)`

GetLearnedOk returns a tuple with the Learned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearned

`func (o *RiskRiskTrial) SetLearned(v int64)`

SetLearned sets Learned field to given value.

### HasLearned

`func (o *RiskRiskTrial) HasLearned() bool`

HasLearned returns a boolean if a field has been set.

### GetRealised

`func (o *RiskRiskTrial) GetRealised() float64`

GetRealised returns the Realised field if non-nil, zero value otherwise.

### GetRealisedOk

`func (o *RiskRiskTrial) GetRealisedOk() (*float64, bool)`

GetRealisedOk returns a tuple with the Realised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealised

`func (o *RiskRiskTrial) SetRealised(v float64)`

SetRealised sets Realised field to given value.

### HasRealised

`func (o *RiskRiskTrial) HasRealised() bool`

HasRealised returns a boolean if a field has been set.

### GetSaturated

`func (o *RiskRiskTrial) GetSaturated() bool`

GetSaturated returns the Saturated field if non-nil, zero value otherwise.

### GetSaturatedOk

`func (o *RiskRiskTrial) GetSaturatedOk() (*bool, bool)`

GetSaturatedOk returns a tuple with the Saturated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturated

`func (o *RiskRiskTrial) SetSaturated(v bool)`

SetSaturated sets Saturated field to given value.

### HasSaturated

`func (o *RiskRiskTrial) HasSaturated() bool`

HasSaturated returns a boolean if a field has been set.

### GetScored

`func (o *RiskRiskTrial) GetScored() int64`

GetScored returns the Scored field if non-nil, zero value otherwise.

### GetScoredOk

`func (o *RiskRiskTrial) GetScoredOk() (*int64, bool)`

GetScoredOk returns a tuple with the Scored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScored

`func (o *RiskRiskTrial) SetScored(v int64)`

SetScored sets Scored field to given value.

### HasScored

`func (o *RiskRiskTrial) HasScored() bool`

HasScored returns a boolean if a field has been set.

### GetStated

`func (o *RiskRiskTrial) GetStated() float64`

GetStated returns the Stated field if non-nil, zero value otherwise.

### GetStatedOk

`func (o *RiskRiskTrial) GetStatedOk() (*float64, bool)`

GetStatedOk returns a tuple with the Stated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStated

`func (o *RiskRiskTrial) SetStated(v float64)`

SetStated sets Stated field to given value.

### HasStated

`func (o *RiskRiskTrial) HasStated() bool`

HasStated returns a boolean if a field has been set.

### GetTopology

`func (o *RiskRiskTrial) GetTopology() RiskRiskTopology`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *RiskRiskTrial) GetTopologyOk() (*RiskRiskTopology, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *RiskRiskTrial) SetTopology(v RiskRiskTopology)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *RiskRiskTrial) HasTopology() bool`

HasTopology returns a boolean if a field has been set.

### GetWarm

`func (o *RiskRiskTrial) GetWarm() bool`

GetWarm returns the Warm field if non-nil, zero value otherwise.

### GetWarmOk

`func (o *RiskRiskTrial) GetWarmOk() (*bool, bool)`

GetWarmOk returns a tuple with the Warm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarm

`func (o *RiskRiskTrial) SetWarm(v bool)`

SetWarm sets Warm field to given value.

### HasWarm

`func (o *RiskRiskTrial) HasWarm() bool`

HasWarm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


