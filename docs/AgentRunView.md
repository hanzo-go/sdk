# AgentRunView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** |  | [optional] 
**Agent** | Pointer to **string** | What an operator needs to answer \&quot;what ran, for whom, and what did it do\&quot; — and, through traceId, to leave this record for the waterfall of the very same run rather than a search that hopefully lands near it.  Agent is on the row because the org-wide feed lists runs across agents, and a run that cannot name its agent is an orphan in exactly the view built to make sense of many of them. Every field is omitempty: a run recorded before these columns existed reports absence rather than a zero it never measured. | [optional] 
**CompletionTokens** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**DurationMs** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Input** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Output** | Pointer to **string** |  | [optional] 
**PromptTokens** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ToolCalls** | Pointer to **int32** |  | [optional] 
**TraceId** | Pointer to **string** |  | [optional] 

## Methods

### NewAgentRunView

`func NewAgentRunView() *AgentRunView`

NewAgentRunView instantiates a new AgentRunView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRunViewWithDefaults

`func NewAgentRunViewWithDefaults() *AgentRunView`

NewAgentRunViewWithDefaults instantiates a new AgentRunView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *AgentRunView) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *AgentRunView) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *AgentRunView) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *AgentRunView) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetAgent

`func (o *AgentRunView) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *AgentRunView) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *AgentRunView) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *AgentRunView) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetCompletionTokens

`func (o *AgentRunView) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *AgentRunView) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *AgentRunView) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *AgentRunView) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentRunView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentRunView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentRunView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentRunView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDurationMs

`func (o *AgentRunView) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *AgentRunView) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *AgentRunView) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *AgentRunView) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetError

`func (o *AgentRunView) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AgentRunView) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AgentRunView) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AgentRunView) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *AgentRunView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentRunView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentRunView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentRunView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInput

`func (o *AgentRunView) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *AgentRunView) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *AgentRunView) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *AgentRunView) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetModel

`func (o *AgentRunView) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentRunView) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentRunView) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AgentRunView) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOutput

`func (o *AgentRunView) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *AgentRunView) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *AgentRunView) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *AgentRunView) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetPromptTokens

`func (o *AgentRunView) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *AgentRunView) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *AgentRunView) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *AgentRunView) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetStatus

`func (o *AgentRunView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentRunView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentRunView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentRunView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToolCalls

`func (o *AgentRunView) GetToolCalls() int32`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *AgentRunView) GetToolCallsOk() (*int32, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *AgentRunView) SetToolCalls(v int32)`

SetToolCalls sets ToolCalls field to given value.

### HasToolCalls

`func (o *AgentRunView) HasToolCalls() bool`

HasToolCalls returns a boolean if a field has been set.

### GetTraceId

`func (o *AgentRunView) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *AgentRunView) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *AgentRunView) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *AgentRunView) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


