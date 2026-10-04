# MarketingSequenceView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sequence** | Pointer to [**MarketingSequence**](MarketingSequence.md) | Sequence is the definition itself — the same record create and the list return. Its status is the one that decides whether enroll is accepted. | [optional] 
**Steps** | Pointer to [**[]MarketingStep**](MarketingStep.md) | Steps are in send order (idx ascending); empty for a sequence with no messages yet, which enrolls fine and completes immediately. | [optional] 

## Methods

### NewMarketingSequenceView

`func NewMarketingSequenceView() *MarketingSequenceView`

NewMarketingSequenceView instantiates a new MarketingSequenceView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingSequenceViewWithDefaults

`func NewMarketingSequenceViewWithDefaults() *MarketingSequenceView`

NewMarketingSequenceViewWithDefaults instantiates a new MarketingSequenceView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSequence

`func (o *MarketingSequenceView) GetSequence() MarketingSequence`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *MarketingSequenceView) GetSequenceOk() (*MarketingSequence, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *MarketingSequenceView) SetSequence(v MarketingSequence)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *MarketingSequenceView) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSteps

`func (o *MarketingSequenceView) GetSteps() []MarketingStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *MarketingSequenceView) GetStepsOk() (*[]MarketingStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *MarketingSequenceView) SetSteps(v []MarketingStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *MarketingSequenceView) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


