# CloudSequenceView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sequence** | Pointer to [**CloudSequence**](CloudSequence.md) |  | [optional] 
**Steps** | Pointer to [**[]CloudStep**](CloudStep.md) | Steps are in send order (idx ascending); empty for a sequence with no messages yet, which enrolls fine and completes immediately. | [optional] 

## Methods

### NewCloudSequenceView

`func NewCloudSequenceView() *CloudSequenceView`

NewCloudSequenceView instantiates a new CloudSequenceView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSequenceViewWithDefaults

`func NewCloudSequenceViewWithDefaults() *CloudSequenceView`

NewCloudSequenceViewWithDefaults instantiates a new CloudSequenceView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSequence

`func (o *CloudSequenceView) GetSequence() CloudSequence`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *CloudSequenceView) GetSequenceOk() (*CloudSequence, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *CloudSequenceView) SetSequence(v CloudSequence)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *CloudSequenceView) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetSteps

`func (o *CloudSequenceView) GetSteps() []CloudStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *CloudSequenceView) GetStepsOk() (*[]CloudStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *CloudSequenceView) SetSteps(v []CloudStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *CloudSequenceView) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


