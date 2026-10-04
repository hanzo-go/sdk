# LabelRiskLabelIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to [**[]LabelRiskLabelFact**](LabelRiskLabelFact.md) | Labels is the batch. Each member is judged on its own: one refusal does not discard the rest, because a webhook redelivering five disputes must not lose four of them to one malformed fifth. | [optional] 

## Methods

### NewLabelRiskLabelIn

`func NewLabelRiskLabelIn() *LabelRiskLabelIn`

NewLabelRiskLabelIn instantiates a new LabelRiskLabelIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelRiskLabelInWithDefaults

`func NewLabelRiskLabelInWithDefaults() *LabelRiskLabelIn`

NewLabelRiskLabelInWithDefaults instantiates a new LabelRiskLabelIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *LabelRiskLabelIn) GetLabels() []LabelRiskLabelFact`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *LabelRiskLabelIn) GetLabelsOk() (*[]LabelRiskLabelFact, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *LabelRiskLabelIn) SetLabels(v []LabelRiskLabelFact)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *LabelRiskLabelIn) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


