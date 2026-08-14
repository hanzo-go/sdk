# RiskLabelIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to [**[]RiskLabelFact**](RiskLabelFact.md) | Labels is the batch. Each member is judged on its own: one refusal does not discard the rest, because a webhook redelivering five disputes must not lose four of them to one malformed fifth. | [optional] 

## Methods

### NewRiskLabelIn

`func NewRiskLabelIn() *RiskLabelIn`

NewRiskLabelIn instantiates a new RiskLabelIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskLabelInWithDefaults

`func NewRiskLabelInWithDefaults() *RiskLabelIn`

NewRiskLabelInWithDefaults instantiates a new RiskLabelIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *RiskLabelIn) GetLabels() []RiskLabelFact`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RiskLabelIn) GetLabelsOk() (*[]RiskLabelFact, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RiskLabelIn) SetLabels(v []RiskLabelFact)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RiskLabelIn) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


