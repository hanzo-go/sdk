# RiskRiskScoreOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | Pointer to **bool** | Alert is whether this would become evidence. It is false in shadow however high the score. | [optional] 
**Causes** | Pointer to [**[]RiskRiskCause**](RiskRiskCause.md) | Causes is the per-feature attribution, ordered by contribution. Each is a COUNTERFACTUAL on the model that produced the score — the coordinate moved to its neutral value and the event rescored — so the explanation is the same arithmetic the score came from. | [optional] 
**Cut** | Pointer to **float64** | Cut is the threshold in force, derived from the stated appetite as a quantile of the scores actually observed rather than fixed at a number. | [optional] 
**Policy** | Pointer to **int64** | Policy is the version of your organisation&#39;s decision regime this verdict was reached under, from its own policy history (GET /v1/risk/policy). Cut is derived from the appetite that version states, so it is the record that makes this decision reconstructible after the appetite is restated. Zero means no regime has ever been stated and the default posture — shadow — was in force. | [optional] 
**Refusal** | Pointer to **string** | Refusal names why the model declined, when it did. | [optional] 
**Score** | Pointer to **float64** | Score is where the event sits in the tenant&#39;s own density: 0 where its recent behaviour is densest, 1 where there is none of it. | [optional] 
**Scored** | Pointer to **bool** | Scored is false when the model declined, and Refusal says which refusal it was: warming, unusable or unidentified. None of them is a clean bill of health, which is why the refusal is stated rather than rendered as a score of zero. | [optional] 
**Shadow** | Pointer to **bool** | Shadow is whether the model is testing rather than deciding — scoring, learning and recording what it WOULD have alerted on, and changing no outcome. It is the default for a model no one has reviewed yet. | [optional] 
**Shape** | Pointer to **string** | Shape is the model space this verdict was reached in, as &#x60;&lt;family&gt;:&lt;digest&gt;&#x60;: the KIND of model, and that family&#39;s own digest over your organisation&#39;s feature inventory in order and the detector&#39;s geometry parameters. It is what pins an adverse decision to a model — a score is only meaningful against the space that produced it, and without this the only answer to \&quot;which model decided this\&quot; was \&quot;the one that was running\&quot;, which is not an answer.  The family leads it because everything after it is one family&#39;s arithmetic. Two spaces are the same space only if they are the same family, so comparing this with the &#x60;shape&#x60; on your model state or on a published value is a comparison that holds ACROSS families and not only inside one.  It names the SPACE, not the learned state, and that is deliberate. The masses at the instant of a score are in-process counters somewhere between two published values, so citing a published address here would claim that value produced this score — true only for the score taken the instant after a publication. This, the policy version and the event&#39;s own time are what IS true, and the published history&#39;s clock (GET /v1/risk/state) brackets the decision between two named values from there. | [optional] 
**Values** | Pointer to [**[]RiskRiskValue**](RiskRiskValue.md) | Values is every coordinate, including the ones that contributed nothing, so a reviewer sees what the model read and not only what it concluded. | [optional] 

## Methods

### NewRiskRiskScoreOut

`func NewRiskRiskScoreOut() *RiskRiskScoreOut`

NewRiskRiskScoreOut instantiates a new RiskRiskScoreOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRiskScoreOutWithDefaults

`func NewRiskRiskScoreOutWithDefaults() *RiskRiskScoreOut`

NewRiskRiskScoreOutWithDefaults instantiates a new RiskRiskScoreOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *RiskRiskScoreOut) GetAlert() bool`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *RiskRiskScoreOut) GetAlertOk() (*bool, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *RiskRiskScoreOut) SetAlert(v bool)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *RiskRiskScoreOut) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetCauses

`func (o *RiskRiskScoreOut) GetCauses() []RiskRiskCause`

GetCauses returns the Causes field if non-nil, zero value otherwise.

### GetCausesOk

`func (o *RiskRiskScoreOut) GetCausesOk() (*[]RiskRiskCause, bool)`

GetCausesOk returns a tuple with the Causes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauses

`func (o *RiskRiskScoreOut) SetCauses(v []RiskRiskCause)`

SetCauses sets Causes field to given value.

### HasCauses

`func (o *RiskRiskScoreOut) HasCauses() bool`

HasCauses returns a boolean if a field has been set.

### GetCut

`func (o *RiskRiskScoreOut) GetCut() float64`

GetCut returns the Cut field if non-nil, zero value otherwise.

### GetCutOk

`func (o *RiskRiskScoreOut) GetCutOk() (*float64, bool)`

GetCutOk returns a tuple with the Cut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCut

`func (o *RiskRiskScoreOut) SetCut(v float64)`

SetCut sets Cut field to given value.

### HasCut

`func (o *RiskRiskScoreOut) HasCut() bool`

HasCut returns a boolean if a field has been set.

### GetPolicy

`func (o *RiskRiskScoreOut) GetPolicy() int64`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *RiskRiskScoreOut) GetPolicyOk() (*int64, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *RiskRiskScoreOut) SetPolicy(v int64)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *RiskRiskScoreOut) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetRefusal

`func (o *RiskRiskScoreOut) GetRefusal() string`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *RiskRiskScoreOut) GetRefusalOk() (*string, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *RiskRiskScoreOut) SetRefusal(v string)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *RiskRiskScoreOut) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.

### GetScore

`func (o *RiskRiskScoreOut) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *RiskRiskScoreOut) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *RiskRiskScoreOut) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *RiskRiskScoreOut) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetScored

`func (o *RiskRiskScoreOut) GetScored() bool`

GetScored returns the Scored field if non-nil, zero value otherwise.

### GetScoredOk

`func (o *RiskRiskScoreOut) GetScoredOk() (*bool, bool)`

GetScoredOk returns a tuple with the Scored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScored

`func (o *RiskRiskScoreOut) SetScored(v bool)`

SetScored sets Scored field to given value.

### HasScored

`func (o *RiskRiskScoreOut) HasScored() bool`

HasScored returns a boolean if a field has been set.

### GetShadow

`func (o *RiskRiskScoreOut) GetShadow() bool`

GetShadow returns the Shadow field if non-nil, zero value otherwise.

### GetShadowOk

`func (o *RiskRiskScoreOut) GetShadowOk() (*bool, bool)`

GetShadowOk returns a tuple with the Shadow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadow

`func (o *RiskRiskScoreOut) SetShadow(v bool)`

SetShadow sets Shadow field to given value.

### HasShadow

`func (o *RiskRiskScoreOut) HasShadow() bool`

HasShadow returns a boolean if a field has been set.

### GetShape

`func (o *RiskRiskScoreOut) GetShape() string`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *RiskRiskScoreOut) GetShapeOk() (*string, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *RiskRiskScoreOut) SetShape(v string)`

SetShape sets Shape field to given value.

### HasShape

`func (o *RiskRiskScoreOut) HasShape() bool`

HasShape returns a boolean if a field has been set.

### GetValues

`func (o *RiskRiskScoreOut) GetValues() []RiskRiskValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *RiskRiskScoreOut) GetValuesOk() (*[]RiskRiskValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *RiskRiskScoreOut) SetValues(v []RiskRiskValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *RiskRiskScoreOut) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


