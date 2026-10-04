# RunnerNeed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | Pointer to **string** | Job is the depended-on job&#39;s id in the workflow. | [optional] 
**Outputs** | Pointer to [**[]RunnerPair**](RunnerPair.md) | Outputs are the values it published. | [optional] 
**Result** | Pointer to **string** | Result is how it finished. | [optional] 

## Methods

### NewRunnerNeed

`func NewRunnerNeed() *RunnerNeed`

NewRunnerNeed instantiates a new RunnerNeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerNeedWithDefaults

`func NewRunnerNeedWithDefaults() *RunnerNeed`

NewRunnerNeedWithDefaults instantiates a new RunnerNeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *RunnerNeed) GetJob() string`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *RunnerNeed) GetJobOk() (*string, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *RunnerNeed) SetJob(v string)`

SetJob sets Job field to given value.

### HasJob

`func (o *RunnerNeed) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetOutputs

`func (o *RunnerNeed) GetOutputs() []RunnerPair`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *RunnerNeed) GetOutputsOk() (*[]RunnerPair, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *RunnerNeed) SetOutputs(v []RunnerPair)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *RunnerNeed) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetResult

`func (o *RunnerNeed) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RunnerNeed) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RunnerNeed) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *RunnerNeed) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


