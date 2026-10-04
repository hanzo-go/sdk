# RunnerTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**RunnerContext**](RunnerContext.md) | Context is what the job evaluates github.* against. | [optional] 
**Id** | Pointer to **int64** | ID identifies this task on every later state and log report. | [optional] 
**Needs** | Pointer to [**[]RunnerNeed**](RunnerNeed.md) | Needs is what the jobs this one depends on produced. | [optional] 
**Secrets** | Pointer to [**[]RunnerPair**](RunnerPair.md) | Secrets are the values masked out of the log and exposed as secrets.*. | [optional] 
**Vars** | Pointer to [**[]RunnerPair**](RunnerPair.md) | Vars are the values exposed as vars.*. | [optional] 
**Workflow** | Pointer to **string** | Workflow is the workflow document verbatim. The forge decides THAT a job runs; act, on the runner, decides what the document means. | [optional] 

## Methods

### NewRunnerTask

`func NewRunnerTask() *RunnerTask`

NewRunnerTask instantiates a new RunnerTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerTaskWithDefaults

`func NewRunnerTaskWithDefaults() *RunnerTask`

NewRunnerTaskWithDefaults instantiates a new RunnerTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *RunnerTask) GetContext() RunnerContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *RunnerTask) GetContextOk() (*RunnerContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *RunnerTask) SetContext(v RunnerContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *RunnerTask) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *RunnerTask) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunnerTask) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunnerTask) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RunnerTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNeeds

`func (o *RunnerTask) GetNeeds() []RunnerNeed`

GetNeeds returns the Needs field if non-nil, zero value otherwise.

### GetNeedsOk

`func (o *RunnerTask) GetNeedsOk() (*[]RunnerNeed, bool)`

GetNeedsOk returns a tuple with the Needs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeeds

`func (o *RunnerTask) SetNeeds(v []RunnerNeed)`

SetNeeds sets Needs field to given value.

### HasNeeds

`func (o *RunnerTask) HasNeeds() bool`

HasNeeds returns a boolean if a field has been set.

### GetSecrets

`func (o *RunnerTask) GetSecrets() []RunnerPair`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *RunnerTask) GetSecretsOk() (*[]RunnerPair, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *RunnerTask) SetSecrets(v []RunnerPair)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *RunnerTask) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetVars

`func (o *RunnerTask) GetVars() []RunnerPair`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *RunnerTask) GetVarsOk() (*[]RunnerPair, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *RunnerTask) SetVars(v []RunnerPair)`

SetVars sets Vars field to given value.

### HasVars

`func (o *RunnerTask) HasVars() bool`

HasVars returns a boolean if a field has been set.

### GetWorkflow

`func (o *RunnerTask) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *RunnerTask) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *RunnerTask) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *RunnerTask) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


