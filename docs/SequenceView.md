# SequenceView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sequence** | Pointer to [**Sequence**](Sequence.md) | Sequence is the definition itself — the same record create and the list return. Its status is the one that decides whether enroll is accepted. | [optional] 
**Steps** | Pointer to [**[]Step**](Step.md) | Steps are in send order (idx ascending); empty for a sequence with no messages yet, which enrolls fine and completes immediately. | [optional] 

## Methods

### NewSequenceView

`func NewSequenceView() *SequenceView`

NewSequenceView instantiates a new SequenceView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSequenceViewWithDefaults

`func NewSequenceViewWithDefaults() *SequenceView`

NewSequenceViewWithDefaults instantiates a new SequenceView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSequence

`func (o *SequenceView) GetSequence() Sequence`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *SequenceView) GetSequenceOk() (*Sequence, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *SequenceView) SetSequence(v Sequence)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *SequenceView) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSteps

`func (o *SequenceView) GetSteps() []Step`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *SequenceView) GetStepsOk() (*[]Step, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *SequenceView) SetSteps(v []Step)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *SequenceView) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


