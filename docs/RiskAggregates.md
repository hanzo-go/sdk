# RiskAggregates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bound** | Pointer to **int64** | Bound is the most they can hold. It is a per-organisation bound: at it, this organisation degrades and no other one notices. | [optional] 
**Forgotten** | Pointer to **int64** | Forgotten is how many of its own subjects have been dropped to stay inside that bound. Each one reads as inactive until it is active again. | [optional] 
**Saturated** | Pointer to **bool** | Saturated is whether the bound is binding right now. The two counts are its evidence; this is the state to act on. | [optional] 
**Subjects** | Pointer to **int64** | Subjects is how many of this organisation&#39;s subjects the aggregates hold. | [optional] 

## Methods

### NewRiskAggregates

`func NewRiskAggregates() *RiskAggregates`

NewRiskAggregates instantiates a new RiskAggregates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAggregatesWithDefaults

`func NewRiskAggregatesWithDefaults() *RiskAggregates`

NewRiskAggregatesWithDefaults instantiates a new RiskAggregates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBound

`func (o *RiskAggregates) GetBound() int64`

GetBound returns the Bound field if non-nil, zero value otherwise.

### GetBoundOk

`func (o *RiskAggregates) GetBoundOk() (*int64, bool)`

GetBoundOk returns a tuple with the Bound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBound

`func (o *RiskAggregates) SetBound(v int64)`

SetBound sets Bound field to given value.

### HasBound

`func (o *RiskAggregates) HasBound() bool`

HasBound returns a boolean if a field has been set.

### GetForgotten

`func (o *RiskAggregates) GetForgotten() int64`

GetForgotten returns the Forgotten field if non-nil, zero value otherwise.

### GetForgottenOk

`func (o *RiskAggregates) GetForgottenOk() (*int64, bool)`

GetForgottenOk returns a tuple with the Forgotten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForgotten

`func (o *RiskAggregates) SetForgotten(v int64)`

SetForgotten sets Forgotten field to given value.

### HasForgotten

`func (o *RiskAggregates) HasForgotten() bool`

HasForgotten returns a boolean if a field has been set.

### GetSaturated

`func (o *RiskAggregates) GetSaturated() bool`

GetSaturated returns the Saturated field if non-nil, zero value otherwise.

### GetSaturatedOk

`func (o *RiskAggregates) GetSaturatedOk() (*bool, bool)`

GetSaturatedOk returns a tuple with the Saturated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturated

`func (o *RiskAggregates) SetSaturated(v bool)`

SetSaturated sets Saturated field to given value.

### HasSaturated

`func (o *RiskAggregates) HasSaturated() bool`

HasSaturated returns a boolean if a field has been set.

### GetSubjects

`func (o *RiskAggregates) GetSubjects() int64`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *RiskAggregates) GetSubjectsOk() (*int64, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *RiskAggregates) SetSubjects(v int64)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *RiskAggregates) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


