# RiskRiskAggregates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bound** | Pointer to **int64** | Bound is the most they can hold. It is a per-organisation bound: at it, this organisation degrades and no other one notices. | [optional] 
**Forgotten** | Pointer to **int64** | Forgotten is how many of its own subjects have been dropped to stay inside that bound. Each one reads as inactive until it is active again. | [optional] 
**Saturated** | Pointer to **bool** | Saturated is whether the bound is binding right now. The two counts are its evidence; this is the state to act on. | [optional] 
**Subjects** | Pointer to **int64** | Subjects is how many of this organisation&#39;s subjects the aggregates hold. | [optional] 

## Methods

### NewRiskRiskAggregates

`func NewRiskRiskAggregates() *RiskRiskAggregates`

NewRiskRiskAggregates instantiates a new RiskRiskAggregates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRiskAggregatesWithDefaults

`func NewRiskRiskAggregatesWithDefaults() *RiskRiskAggregates`

NewRiskRiskAggregatesWithDefaults instantiates a new RiskRiskAggregates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBound

`func (o *RiskRiskAggregates) GetBound() int64`

GetBound returns the Bound field if non-nil, zero value otherwise.

### GetBoundOk

`func (o *RiskRiskAggregates) GetBoundOk() (*int64, bool)`

GetBoundOk returns a tuple with the Bound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBound

`func (o *RiskRiskAggregates) SetBound(v int64)`

SetBound sets Bound field to given value.

### HasBound

`func (o *RiskRiskAggregates) HasBound() bool`

HasBound returns a boolean if a field has been set.

### GetForgotten

`func (o *RiskRiskAggregates) GetForgotten() int64`

GetForgotten returns the Forgotten field if non-nil, zero value otherwise.

### GetForgottenOk

`func (o *RiskRiskAggregates) GetForgottenOk() (*int64, bool)`

GetForgottenOk returns a tuple with the Forgotten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgotten

`func (o *RiskRiskAggregates) SetForgotten(v int64)`

SetForgotten sets Forgotten field to given value.

### HasForgotten

`func (o *RiskRiskAggregates) HasForgotten() bool`

HasForgotten returns a boolean if a field has been set.

### GetSaturated

`func (o *RiskRiskAggregates) GetSaturated() bool`

GetSaturated returns the Saturated field if non-nil, zero value otherwise.

### GetSaturatedOk

`func (o *RiskRiskAggregates) GetSaturatedOk() (*bool, bool)`

GetSaturatedOk returns a tuple with the Saturated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturated

`func (o *RiskRiskAggregates) SetSaturated(v bool)`

SetSaturated sets Saturated field to given value.

### HasSaturated

`func (o *RiskRiskAggregates) HasSaturated() bool`

HasSaturated returns a boolean if a field has been set.

### GetSubjects

`func (o *RiskRiskAggregates) GetSubjects() int64`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *RiskRiskAggregates) GetSubjectsOk() (*int64, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *RiskRiskAggregates) SetSubjects(v int64)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *RiskRiskAggregates) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


