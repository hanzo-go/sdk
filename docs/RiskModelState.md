# RiskModelState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregates** | Pointer to [**RiskAggregates**](RiskAggregates.md) | Aggregates reports the pressure on this organisation&#39;s own sliding aggregates, and whether they have started forgetting subjects to stay inside their bound. | [optional] 
**Blind** | Pointer to **map[string]int32** | Blind counts, per feature, how often it took its neutral value for want of data. A feature blind on most traffic is not contributing whatever the inventory claims for it. | [optional] 
**Cut** | Pointer to **float32** | Cut is the threshold in force, derived from Stated as a quantile of the scores actually observed. | [optional] 
**Descends** | Pointer to **string** | Descends is the published value the working model grew out of: the newest one whose mass count it has reached or passed. Empty when nothing has been published yet.  It is DERIVED from the count and never stored, so an instant rollback is right for free — adopting an older value moves the count backward and this answers with that older value, where a stored pointer would be a second fact to keep in step. Read with Learned it is also the DRIFT: this model is Descends plus however many events the two counts differ by. | [optional] 
**Disposed** | Pointer to **int32** | Disposed is how many published values retention has taken. It is DERIVED from the lowest surviving sequence, so it cannot drift from what it describes, and it is reported because a retention that binds is a fact an operator must be able to read rather than a silence. | [optional] 
**Learned** | Pointer to **int32** | Learned is how many events the model has learned from. | [optional] 
**Live** | Pointer to **bool** | Live is false while the model is in shadow — scoring, learning and recording what it WOULD have alerted on, and changing no outcome. Shadow is the default for a new tenant. | [optional] 
**Policy** | Pointer to **int32** | Policy is the version of the decision regime this model is deciding under, from your organisation&#39;s own policy history (GET /v1/risk/policy). Every score cites it, so it is the join between a past decision and the appetite that produced its threshold. Zero means no regime has ever been stated and the default posture — shadow — is in force. | [optional] 
**Realised** | Pointer to **float32** | Realised is the share that actually was. Reading it beside Stated is what makes the appetite a measured commitment rather than an intention. | [optional] 
**Refused** | Pointer to **map[string]int32** | Refused counts events the model would not score, by reason. None of them was examined; a refusal is counted, never silent. | [optional] 
**Sample** | Pointer to **float32** | Sample is the share of below-the-line events retained for review, which is how the miss rate is measured rather than assumed. | [optional] 
**Saturated** | Pointer to **bool** | Saturated means no threshold can honour the stated appetite because too much of the stream scores in the top bucket, so the model is alerting on nothing — the one state that must never be mistaken for quiet. | [optional] 
**Shape** | Pointer to **string** | Shape is the model&#39;s identity, as &#x60;&lt;family&gt;:&lt;digest&gt;&#x60;: the KIND of model, and that family&#39;s own digest over the inventory in order and the detector&#39;s geometry parameters. It is what an auditor pins an alert to, because learned state is only meaningful against the space that produced it — and the family leads it because two families&#39; masses are not fitted differently, they are different kinds of number. | [optional] 
**Stated** | Pointer to **float32** | Stated is the share of the stream this organisation said may be examined. | [optional] 
**Surface** | Pointer to [**RiskSurface**](RiskSurface.md) | Surface reports what of the tenant&#39;s OWN event surface has been folded in. | [optional] 
**Tenant** | Pointer to **string** | Tenant is the qualified key the model is held under — the brand whose issuer vouched for the caller and the organisation it acts for. It is echoed so a reader can see the answer is its own and not a parameter it passed. | [optional] 
**Values** | Pointer to [**[]RiskModelValue**](RiskModelValue.md) | Values is your organisation&#39;s own published model values, newest first — every state it deliberately named, each addressed by its own content and immutable. This is what PUT /v1/risk/state/model names, so it is reported HERE rather than behind an address of its own: they are part of what a review of one model reads, and a list of names is a few hundred bytes.  Compare each one&#39;s &#x60;shape&#x60; with the &#x60;shape&#x60; above: equal means adopting it restores masses into the space this model already runs, and different means adopting it REPLANTS the model into the space that value describes — which is how the shape a search found becomes the shape you are running.  The working model is NOT in it. Publication is a boundary somebody marked; the state between two boundaries is in-process counters, and calling those a value would be a claim about reproducibility that nothing could honour. | [optional] 
**Warm** | Pointer to **bool** | Warm is whether that is enough for the model to have an opinion at all. Below it the model declines to score, which is an ordinary state and is not a clean bill of health. | [optional] 

## Methods

### NewRiskModelState

`func NewRiskModelState() *RiskModelState`

NewRiskModelState instantiates a new RiskModelState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskModelStateWithDefaults

`func NewRiskModelStateWithDefaults() *RiskModelState`

NewRiskModelStateWithDefaults instantiates a new RiskModelState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregates

`func (o *RiskModelState) GetAggregates() RiskAggregates`

GetAggregates returns the Aggregates field if non-nil, zero value otherwise.

### GetAggregatesOk

`func (o *RiskModelState) GetAggregatesOk() (*RiskAggregates, bool)`

GetAggregatesOk returns a tuple with the Aggregates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregates

`func (o *RiskModelState) SetAggregates(v RiskAggregates)`

SetAggregates sets Aggregates field to given value.

### HasAggregates

`func (o *RiskModelState) HasAggregates() bool`

HasAggregates returns a boolean if a field has been set.

### GetBlind

`func (o *RiskModelState) GetBlind() map[string]int32`

GetBlind returns the Blind field if non-nil, zero value otherwise.

### GetBlindOk

`func (o *RiskModelState) GetBlindOk() (*map[string]int32, bool)`

GetBlindOk returns a tuple with the Blind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlind

`func (o *RiskModelState) SetBlind(v map[string]int32)`

SetBlind sets Blind field to given value.

### HasBlind

`func (o *RiskModelState) HasBlind() bool`

HasBlind returns a boolean if a field has been set.

### GetCut

`func (o *RiskModelState) GetCut() float32`

GetCut returns the Cut field if non-nil, zero value otherwise.

### GetCutOk

`func (o *RiskModelState) GetCutOk() (*float32, bool)`

GetCutOk returns a tuple with the Cut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCut

`func (o *RiskModelState) SetCut(v float32)`

SetCut sets Cut field to given value.

### HasCut

`func (o *RiskModelState) HasCut() bool`

HasCut returns a boolean if a field has been set.

### GetDescends

`func (o *RiskModelState) GetDescends() string`

GetDescends returns the Descends field if non-nil, zero value otherwise.

### GetDescendsOk

`func (o *RiskModelState) GetDescendsOk() (*string, bool)`

GetDescendsOk returns a tuple with the Descends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescends

`func (o *RiskModelState) SetDescends(v string)`

SetDescends sets Descends field to given value.

### HasDescends

`func (o *RiskModelState) HasDescends() bool`

HasDescends returns a boolean if a field has been set.

### GetDisposed

`func (o *RiskModelState) GetDisposed() int32`

GetDisposed returns the Disposed field if non-nil, zero value otherwise.

### GetDisposedOk

`func (o *RiskModelState) GetDisposedOk() (*int32, bool)`

GetDisposedOk returns a tuple with the Disposed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposed

`func (o *RiskModelState) SetDisposed(v int32)`

SetDisposed sets Disposed field to given value.

### HasDisposed

`func (o *RiskModelState) HasDisposed() bool`

HasDisposed returns a boolean if a field has been set.

### GetLearned

`func (o *RiskModelState) GetLearned() int32`

GetLearned returns the Learned field if non-nil, zero value otherwise.

### GetLearnedOk

`func (o *RiskModelState) GetLearnedOk() (*int32, bool)`

GetLearnedOk returns a tuple with the Learned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLearned

`func (o *RiskModelState) SetLearned(v int32)`

SetLearned sets Learned field to given value.

### HasLearned

`func (o *RiskModelState) HasLearned() bool`

HasLearned returns a boolean if a field has been set.

### GetLive

`func (o *RiskModelState) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *RiskModelState) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *RiskModelState) SetLive(v bool)`

SetLive sets Live field to given value.

### HasLive

`func (o *RiskModelState) HasLive() bool`

HasLive returns a boolean if a field has been set.

### GetPolicy

`func (o *RiskModelState) GetPolicy() int32`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *RiskModelState) GetPolicyOk() (*int32, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *RiskModelState) SetPolicy(v int32)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *RiskModelState) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetRealised

`func (o *RiskModelState) GetRealised() float32`

GetRealised returns the Realised field if non-nil, zero value otherwise.

### GetRealisedOk

`func (o *RiskModelState) GetRealisedOk() (*float32, bool)`

GetRealisedOk returns a tuple with the Realised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealised

`func (o *RiskModelState) SetRealised(v float32)`

SetRealised sets Realised field to given value.

### HasRealised

`func (o *RiskModelState) HasRealised() bool`

HasRealised returns a boolean if a field has been set.

### GetRefused

`func (o *RiskModelState) GetRefused() map[string]int32`

GetRefused returns the Refused field if non-nil, zero value otherwise.

### GetRefusedOk

`func (o *RiskModelState) GetRefusedOk() (*map[string]int32, bool)`

GetRefusedOk returns a tuple with the Refused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefused

`func (o *RiskModelState) SetRefused(v map[string]int32)`

SetRefused sets Refused field to given value.

### HasRefused

`func (o *RiskModelState) HasRefused() bool`

HasRefused returns a boolean if a field has been set.

### GetSample

`func (o *RiskModelState) GetSample() float32`

GetSample returns the Sample field if non-nil, zero value otherwise.

### GetSampleOk

`func (o *RiskModelState) GetSampleOk() (*float32, bool)`

GetSampleOk returns a tuple with the Sample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSample

`func (o *RiskModelState) SetSample(v float32)`

SetSample sets Sample field to given value.

### HasSample

`func (o *RiskModelState) HasSample() bool`

HasSample returns a boolean if a field has been set.

### GetSaturated

`func (o *RiskModelState) GetSaturated() bool`

GetSaturated returns the Saturated field if non-nil, zero value otherwise.

### GetSaturatedOk

`func (o *RiskModelState) GetSaturatedOk() (*bool, bool)`

GetSaturatedOk returns a tuple with the Saturated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturated

`func (o *RiskModelState) SetSaturated(v bool)`

SetSaturated sets Saturated field to given value.

### HasSaturated

`func (o *RiskModelState) HasSaturated() bool`

HasSaturated returns a boolean if a field has been set.

### GetShape

`func (o *RiskModelState) GetShape() string`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *RiskModelState) GetShapeOk() (*string, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *RiskModelState) SetShape(v string)`

SetShape sets Shape field to given value.

### HasShape

`func (o *RiskModelState) HasShape() bool`

HasShape returns a boolean if a field has been set.

### GetStated

`func (o *RiskModelState) GetStated() float32`

GetStated returns the Stated field if non-nil, zero value otherwise.

### GetStatedOk

`func (o *RiskModelState) GetStatedOk() (*float32, bool)`

GetStatedOk returns a tuple with the Stated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStated

`func (o *RiskModelState) SetStated(v float32)`

SetStated sets Stated field to given value.

### HasStated

`func (o *RiskModelState) HasStated() bool`

HasStated returns a boolean if a field has been set.

### GetSurface

`func (o *RiskModelState) GetSurface() RiskSurface`

GetSurface returns the Surface field if non-nil, zero value otherwise.

### GetSurfaceOk

`func (o *RiskModelState) GetSurfaceOk() (*RiskSurface, bool)`

GetSurfaceOk returns a tuple with the Surface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurface

`func (o *RiskModelState) SetSurface(v RiskSurface)`

SetSurface sets Surface field to given value.

### HasSurface

`func (o *RiskModelState) HasSurface() bool`

HasSurface returns a boolean if a field has been set.

### GetTenant

`func (o *RiskModelState) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *RiskModelState) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *RiskModelState) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *RiskModelState) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetValues

`func (o *RiskModelState) GetValues() []RiskModelValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *RiskModelState) GetValuesOk() (*[]RiskModelValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *RiskModelState) SetValues(v []RiskModelValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *RiskModelState) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetWarm

`func (o *RiskModelState) GetWarm() bool`

GetWarm returns the Warm field if non-nil, zero value otherwise.

### GetWarmOk

`func (o *RiskModelState) GetWarmOk() (*bool, bool)`

GetWarmOk returns a tuple with the Warm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarm

`func (o *RiskModelState) SetWarm(v bool)`

SetWarm sets Warm field to given value.

### HasWarm

`func (o *RiskModelState) HasWarm() bool`

HasWarm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


