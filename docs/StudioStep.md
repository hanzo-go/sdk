# StudioStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | Pointer to **int64** | Attempt is how many times the step has started. | [optional] 
**CostCents** | Pointer to **int64** |  | [optional] 
**Error** | Pointer to [**StudioProblem**](StudioProblem.md) |  | [optional] 
**Executor** | Pointer to **string** | Executor is where the step runs: gateway, studio.render, render or flow.run. | [optional] 
**Finished** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Items** | Pointer to **int64** | Items is how many items a step with each fanned out to. | [optional] 
**Message** | Pointer to **string** | Message is what the step is doing now, for a person. | [optional] 
**Op** | Pointer to **string** |  | [optional] 
**Outputs** | Pointer to **[]string** | Outputs are the Library item ids the step made. | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 
**Started** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | Status is pending, queued, running, succeeded, failed, canceled or skipped. | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewStudioStep

`func NewStudioStep() *StudioStep`

NewStudioStep instantiates a new StudioStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioStepWithDefaults

`func NewStudioStepWithDefaults() *StudioStep`

NewStudioStepWithDefaults instantiates a new StudioStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempt

`func (o *StudioStep) GetAttempt() int64`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *StudioStep) GetAttemptOk() (*int64, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *StudioStep) SetAttempt(v int64)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *StudioStep) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.

### GetCostCents

`func (o *StudioStep) GetCostCents() int64`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *StudioStep) GetCostCentsOk() (*int64, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *StudioStep) SetCostCents(v int64)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *StudioStep) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### GetError

`func (o *StudioStep) GetError() StudioProblem`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StudioStep) GetErrorOk() (*StudioProblem, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StudioStep) SetError(v StudioProblem)`

SetError sets Error field to given value.

### HasError

`func (o *StudioStep) HasError() bool`

HasError returns a boolean if a field has been set.

### GetExecutor

`func (o *StudioStep) GetExecutor() string`

GetExecutor returns the Executor field if non-nil, zero value otherwise.

### GetExecutorOk

`func (o *StudioStep) GetExecutorOk() (*string, bool)`

GetExecutorOk returns a tuple with the Executor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutor

`func (o *StudioStep) SetExecutor(v string)`

SetExecutor sets Executor field to given value.

### HasExecutor

`func (o *StudioStep) HasExecutor() bool`

HasExecutor returns a boolean if a field has been set.

### GetFinished

`func (o *StudioStep) GetFinished() string`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *StudioStep) GetFinishedOk() (*string, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *StudioStep) SetFinished(v string)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *StudioStep) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetId

`func (o *StudioStep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StudioStep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StudioStep) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StudioStep) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *StudioStep) GetItems() int64`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *StudioStep) GetItemsOk() (*int64, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *StudioStep) SetItems(v int64)`

SetItems sets Items field to given value.

### HasItems

`func (o *StudioStep) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMessage

`func (o *StudioStep) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StudioStep) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StudioStep) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *StudioStep) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOp

`func (o *StudioStep) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *StudioStep) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *StudioStep) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *StudioStep) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetOutputs

`func (o *StudioStep) GetOutputs() []string`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *StudioStep) GetOutputsOk() (*[]string, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *StudioStep) SetOutputs(v []string)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *StudioStep) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetProgress

`func (o *StudioStep) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *StudioStep) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *StudioStep) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *StudioStep) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetStarted

`func (o *StudioStep) GetStarted() string`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *StudioStep) GetStartedOk() (*string, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *StudioStep) SetStarted(v string)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *StudioStep) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetStatus

`func (o *StudioStep) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StudioStep) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StudioStep) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StudioStep) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetText

`func (o *StudioStep) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *StudioStep) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *StudioStep) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *StudioStep) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTitle

`func (o *StudioStep) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StudioStep) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StudioStep) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *StudioStep) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


