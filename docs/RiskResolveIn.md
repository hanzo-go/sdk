# RiskResolveIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Horizon** | Pointer to **int64** | Horizon is how many days an event must age before it may be resolved at all, and it is the whole of the no-leakage rule. 120 for the payment lane (past the Visa and Mastercard dispute windows), 14 for signup abuse. Unstated takes 120. | [optional] 
**Now** | Pointer to **string** | Now moves the observation instant BACKWARDS, RFC 3339. It exists so a BACKTEST can resolve labels as the plane stood at a past moment; without it, every backtest would score a model against knowledge that arrived after the decision it is being scored on. An instant after the server clock is refused: a backtest resolves the past, and a future one would declare unmatured events matured and hand a training set negatives for rows whose chargeback has not had time to arrive. | [optional] 
**Subjects** | Pointer to [**[]RiskLabelEvent**](RiskLabelEvent.md) | Subjects are the exact events being judged. Each carries its own event time, because the as-of that keeps the future out is derived from that instant plus the horizon — one as-of over a whole batch would give a January row six extra months of hindsight.  One entry per DISTINCT (kind, subject, at): naming an event twice answers once, because an event resolved twice would list its own winner as a contrary claim and would hand a materialiser duplicate training rows. | [optional] 

## Methods

### NewRiskResolveIn

`func NewRiskResolveIn() *RiskResolveIn`

NewRiskResolveIn instantiates a new RiskResolveIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskResolveInWithDefaults

`func NewRiskResolveInWithDefaults() *RiskResolveIn`

NewRiskResolveInWithDefaults instantiates a new RiskResolveIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHorizon

`func (o *RiskResolveIn) GetHorizon() int64`

GetHorizon returns the Horizon field if non-nil, zero value otherwise.

### GetHorizonOk

`func (o *RiskResolveIn) GetHorizonOk() (*int64, bool)`

GetHorizonOk returns a tuple with the Horizon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizon

`func (o *RiskResolveIn) SetHorizon(v int64)`

SetHorizon sets Horizon field to given value.

### HasHorizon

`func (o *RiskResolveIn) HasHorizon() bool`

HasHorizon returns a boolean if a field has been set.

### GetNow

`func (o *RiskResolveIn) GetNow() string`

GetNow returns the Now field if non-nil, zero value otherwise.

### GetNowOk

`func (o *RiskResolveIn) GetNowOk() (*string, bool)`

GetNowOk returns a tuple with the Now field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNow

`func (o *RiskResolveIn) SetNow(v string)`

SetNow sets Now field to given value.

### HasNow

`func (o *RiskResolveIn) HasNow() bool`

HasNow returns a boolean if a field has been set.

### GetSubjects

`func (o *RiskResolveIn) GetSubjects() []RiskLabelEvent`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *RiskResolveIn) GetSubjectsOk() (*[]RiskLabelEvent, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *RiskResolveIn) SetSubjects(v []RiskLabelEvent)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *RiskResolveIn) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


