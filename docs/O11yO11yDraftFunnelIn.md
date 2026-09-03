# O11yO11yDraftFunnelIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **int64** | EndTime is the end of the window, as a millisecond epoch. | [optional] 
**StartTime** | Pointer to **int64** | StartTime is the start of the window, as a millisecond epoch. | [optional] 
**StepEnd** | Pointer to **int64** | StepEnd is the step the transition runs to, 1-based. | [optional] 
**StepStart** | Pointer to **int64** | StepStart is the step the transition runs from, 1-based. Ignored by the reads that span the whole funnel. | [optional] 
**Steps** | Pointer to [**[]O11yFunnelStep**](O11yFunnelStep.md) | Steps are the funnel&#39;s steps, in order. At least two are needed. | [optional] 

## Methods

### NewO11yO11yDraftFunnelIn

`func NewO11yO11yDraftFunnelIn() *O11yO11yDraftFunnelIn`

NewO11yO11yDraftFunnelIn instantiates a new O11yO11yDraftFunnelIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDraftFunnelInWithDefaults

`func NewO11yO11yDraftFunnelInWithDefaults() *O11yO11yDraftFunnelIn`

NewO11yO11yDraftFunnelInWithDefaults instantiates a new O11yO11yDraftFunnelIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *O11yO11yDraftFunnelIn) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *O11yO11yDraftFunnelIn) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *O11yO11yDraftFunnelIn) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *O11yO11yDraftFunnelIn) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartTime

`func (o *O11yO11yDraftFunnelIn) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *O11yO11yDraftFunnelIn) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *O11yO11yDraftFunnelIn) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *O11yO11yDraftFunnelIn) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStepEnd

`func (o *O11yO11yDraftFunnelIn) GetStepEnd() int64`

GetStepEnd returns the StepEnd field if non-nil, zero value otherwise.

### GetStepEndOk

`func (o *O11yO11yDraftFunnelIn) GetStepEndOk() (*int64, bool)`

GetStepEndOk returns a tuple with the StepEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepEnd

`func (o *O11yO11yDraftFunnelIn) SetStepEnd(v int64)`

SetStepEnd sets StepEnd field to given value.

### HasStepEnd

`func (o *O11yO11yDraftFunnelIn) HasStepEnd() bool`

HasStepEnd returns a boolean if a field has been set.

### GetStepStart

`func (o *O11yO11yDraftFunnelIn) GetStepStart() int64`

GetStepStart returns the StepStart field if non-nil, zero value otherwise.

### GetStepStartOk

`func (o *O11yO11yDraftFunnelIn) GetStepStartOk() (*int64, bool)`

GetStepStartOk returns a tuple with the StepStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepStart

`func (o *O11yO11yDraftFunnelIn) SetStepStart(v int64)`

SetStepStart sets StepStart field to given value.

### HasStepStart

`func (o *O11yO11yDraftFunnelIn) HasStepStart() bool`

HasStepStart returns a boolean if a field has been set.

### GetSteps

`func (o *O11yO11yDraftFunnelIn) GetSteps() []O11yFunnelStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *O11yO11yDraftFunnelIn) GetStepsOk() (*[]O11yFunnelStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *O11yO11yDraftFunnelIn) SetSteps(v []O11yFunnelStep)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *O11yO11yDraftFunnelIn) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


