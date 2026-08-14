# RiskTrial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerted** | Pointer to **int32** | Alerted is how many of those it would have raised. | [optional] 
**Curve** | Pointer to **[]float32** | Curve is the realised alert rate over successive tenths of the history — the learning curve, which says whether the shape settled or is still moving. | [optional] 
**Fit** | Pointer to **float32** | Fit ranks the shape, smaller being better: the relative miss of the stated appetite, plus flat penalties for never warming and for saturating, plus the share of coordinates that were blind. | [optional] 
**Learned** | Pointer to **int32** | Learned is how many events the shape learned from during the replay. | [optional] 
**Realised** | Pointer to **float32** | Realised is what that appetite actually produced. The distance between the two is what the search is searching over. | [optional] 
**Saturated** | Pointer to **bool** | Saturated is whether the appetite could not be honoured by any threshold, which is a shape that alerts on nothing and reads like a quiet one. | [optional] 
**Scored** | Pointer to **int32** | Scored is how many it was able to score. | [optional] 
**Stated** | Pointer to **float32** | Stated is the appetite the shape was tried at. | [optional] 
**Topology** | Pointer to [**RiskTopology**](RiskTopology.md) | Topology is the shape. | [optional] 
**Warm** | Pointer to **bool** | Warm is whether the shape learned enough to have an opinion at all over this organisation&#39;s whole history. | [optional] 

## Methods

### NewRiskTrial

`func NewRiskTrial() *RiskTrial`

NewRiskTrial instantiates a new RiskTrial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskTrialWithDefaults

`func NewRiskTrialWithDefaults() *RiskTrial`

NewRiskTrialWithDefaults instantiates a new RiskTrial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerted

`func (o *RiskTrial) GetAlerted() int32`

GetAlerted returns the Alerted field if non-nil, zero value otherwise.

### GetAlertedOk

`func (o *RiskTrial) GetAlertedOk() (*int32, bool)`

GetAlertedOk returns a tuple with the Alerted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerted

`func (o *RiskTrial) SetAlerted(v int32)`

SetAlerted sets Alerted field to given value.

### HasAlerted

`func (o *RiskTrial) HasAlerted() bool`

HasAlerted returns a boolean if a field has been set.

### GetCurve

`func (o *RiskTrial) GetCurve() []float32`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *RiskTrial) GetCurveOk() (*[]float32, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *RiskTrial) SetCurve(v []float32)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *RiskTrial) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### GetFit

`func (o *RiskTrial) GetFit() float32`

GetFit returns the Fit field if non-nil, zero value otherwise.

### GetFitOk

`func (o *RiskTrial) GetFitOk() (*float32, bool)`

GetFitOk returns a tuple with the Fit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFit

`func (o *RiskTrial) SetFit(v float32)`

SetFit sets Fit field to given value.

### HasFit

`func (o *RiskTrial) HasFit() bool`

HasFit returns a boolean if a field has been set.

### GetLearned

`func (o *RiskTrial) GetLearned() int32`

GetLearned returns the Learned field if non-nil, zero value otherwise.

### GetLearnedOk

`func (o *RiskTrial) GetLearnedOk() (*int32, bool)`

GetLearnedOk returns a tuple with the Learned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearned

`func (o *RiskTrial) SetLearned(v int32)`

SetLearned sets Learned field to given value.

### HasLearned

`func (o *RiskTrial) HasLearned() bool`

HasLearned returns a boolean if a field has been set.

### GetRealised

`func (o *RiskTrial) GetRealised() float32`

GetRealised returns the Realised field if non-nil, zero value otherwise.

### GetRealisedOk

`func (o *RiskTrial) GetRealisedOk() (*float32, bool)`

GetRealisedOk returns a tuple with the Realised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealised

`func (o *RiskTrial) SetRealised(v float32)`

SetRealised sets Realised field to given value.

### HasRealised

`func (o *RiskTrial) HasRealised() bool`

HasRealised returns a boolean if a field has been set.

### GetSaturated

`func (o *RiskTrial) GetSaturated() bool`

GetSaturated returns the Saturated field if non-nil, zero value otherwise.

### GetSaturatedOk

`func (o *RiskTrial) GetSaturatedOk() (*bool, bool)`

GetSaturatedOk returns a tuple with the Saturated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturated

`func (o *RiskTrial) SetSaturated(v bool)`

SetSaturated sets Saturated field to given value.

### HasSaturated

`func (o *RiskTrial) HasSaturated() bool`

HasSaturated returns a boolean if a field has been set.

### GetScored

`func (o *RiskTrial) GetScored() int32`

GetScored returns the Scored field if non-nil, zero value otherwise.

### GetScoredOk

`func (o *RiskTrial) GetScoredOk() (*int32, bool)`

GetScoredOk returns a tuple with the Scored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScored

`func (o *RiskTrial) SetScored(v int32)`

SetScored sets Scored field to given value.

### HasScored

`func (o *RiskTrial) HasScored() bool`

HasScored returns a boolean if a field has been set.

### GetStated

`func (o *RiskTrial) GetStated() float32`

GetStated returns the Stated field if non-nil, zero value otherwise.

### GetStatedOk

`func (o *RiskTrial) GetStatedOk() (*float32, bool)`

GetStatedOk returns a tuple with the Stated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStated

`func (o *RiskTrial) SetStated(v float32)`

SetStated sets Stated field to given value.

### HasStated

`func (o *RiskTrial) HasStated() bool`

HasStated returns a boolean if a field has been set.

### GetTopology

`func (o *RiskTrial) GetTopology() RiskTopology`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *RiskTrial) GetTopologyOk() (*RiskTopology, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *RiskTrial) SetTopology(v RiskTopology)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *RiskTrial) HasTopology() bool`

HasTopology returns a boolean if a field has been set.

### GetWarm

`func (o *RiskTrial) GetWarm() bool`

GetWarm returns the Warm field if non-nil, zero value otherwise.

### GetWarmOk

`func (o *RiskTrial) GetWarmOk() (*bool, bool)`

GetWarmOk returns a tuple with the Warm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarm

`func (o *RiskTrial) SetWarm(v bool)`

SetWarm sets Warm field to given value.

### HasWarm

`func (o *RiskTrial) HasWarm() bool`

HasWarm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


