# AgentAgentRunView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actor** | Pointer to **string** | Actor is the \&quot;org/sub\&quot; identity the run was executed and billed AS. Empty means there was no PERSON — a schedule or a service token — which is a different fact from \&quot;we do not know\&quot;, and the difference is what an audit asks about. | [optional] 
**Agent** | Pointer to **string** | What an operator needs to answer \&quot;what ran, for whom, and what did it do\&quot; — and, through traceId, to leave this record for the waterfall of the very same run rather than a search that hopefully lands near it.  Agent is on the row because the org-wide feed lists runs across agents, and a run that cannot name its agent is an orphan in exactly the view built to make sense of many of them. Every field is omitempty: a run recorded before these columns existed reports absence rather than a zero it never measured. | [optional] 
**CompletionTokens** | Pointer to **int64** | CompletionTokens is the same measurement for what the model produced, on the same final completion. It is a count of TOKENS, not of turns and not of money. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when the run finished, RFC 3339 in UTC to the second — the duration above already says how long it had been going. | [optional] 
**DurationMs** | Pointer to **int64** | DurationMs is wall-clock milliseconds around the completion, including a failover&#39;s retries. It is time SPENT, not time billed. | [optional] 
**Error** | Pointer to **string** | Error is why an \&quot;ok\&quot;-less run failed, as the failing call reported it. Empty on every successful run. | [optional] 
**Id** | Pointer to **string** | ID is the run&#39;s handle, minted as \&quot;run_\&quot; + 32 hex characters. It is the key the metering ledger records this run&#39;s per-round token spend under, so it is how a bill and a run are joined. | [optional] 
**Input** | Pointer to **string** | Input is the text the run was given, verbatim. | [optional] 
**MicroUsd** | Pointer to **int64** | MicroUSD is what the run spent, integer micro-USD. Absent when nothing was metered. | [optional] 
**Model** | Pointer to **string** | Model is the model that actually SERVED this run, which is not always the one the agent is defined on — a failover records what answered. Normalized to our name on the way out; the stored row is left exactly as it happened, because a run is a record and rewriting it would be worse than the name it carries. | [optional] 
**Output** | Pointer to **string** | Output is what the model produced. Empty on an error run, and empty is also a legitimate answer from a run that succeeded with nothing to say — Status is what separates those. | [optional] 
**PromptTokens** | Pointer to **int64** | PromptTokens is what the gateway reported for the run&#39;s FINAL completion, and only that one — a tool loop&#39;s earlier rounds are the metering ledger&#39;s account, joined by this run&#39;s id. Reading it as the run&#39;s total spend undercounts a loop. | [optional] 
**Status** | Pointer to **string** | Status is the run&#39;s outcome, and there are exactly two: \&quot;ok\&quot; when the model answered, \&quot;error\&quot; when it did not. It is written when the run ends, so no row here is in flight. | [optional] 
**ToolCalls** | Pointer to **int64** | ToolCalls is how many tool dispatches the run made — a count of ACTIONS, which is a different measurement from the token counts above and from the turns a build reports. Zero is a run that answered straight from the model. | [optional] 
**TraceId** | Pointer to **string** | TraceID is the trace this run IS, so the record and its spans are one thing to move between: it opens the waterfall for THIS run rather than a search that lands near it. Empty when the process had no tracer, never a fabricated id. | [optional] 

## Methods

### NewAgentAgentRunView

`func NewAgentAgentRunView() *AgentAgentRunView`

NewAgentAgentRunView instantiates a new AgentAgentRunView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentAgentRunViewWithDefaults

`func NewAgentAgentRunViewWithDefaults() *AgentAgentRunView`

NewAgentAgentRunViewWithDefaults instantiates a new AgentAgentRunView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActor

`func (o *AgentAgentRunView) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *AgentAgentRunView) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *AgentAgentRunView) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *AgentAgentRunView) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetAgent

`func (o *AgentAgentRunView) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *AgentAgentRunView) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *AgentAgentRunView) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *AgentAgentRunView) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetCompletionTokens

`func (o *AgentAgentRunView) GetCompletionTokens() int64`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *AgentAgentRunView) GetCompletionTokensOk() (*int64, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *AgentAgentRunView) SetCompletionTokens(v int64)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *AgentAgentRunView) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentAgentRunView) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentAgentRunView) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentAgentRunView) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentAgentRunView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDurationMs

`func (o *AgentAgentRunView) GetDurationMs() int64`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *AgentAgentRunView) GetDurationMsOk() (*int64, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *AgentAgentRunView) SetDurationMs(v int64)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *AgentAgentRunView) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetError

`func (o *AgentAgentRunView) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AgentAgentRunView) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AgentAgentRunView) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AgentAgentRunView) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *AgentAgentRunView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentAgentRunView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentAgentRunView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentAgentRunView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInput

`func (o *AgentAgentRunView) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *AgentAgentRunView) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *AgentAgentRunView) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *AgentAgentRunView) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetMicroUsd

`func (o *AgentAgentRunView) GetMicroUsd() int64`

GetMicroUsd returns the MicroUsd field if non-nil, zero value otherwise.

### GetMicroUsdOk

`func (o *AgentAgentRunView) GetMicroUsdOk() (*int64, bool)`

GetMicroUsdOk returns a tuple with the MicroUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroUsd

`func (o *AgentAgentRunView) SetMicroUsd(v int64)`

SetMicroUsd sets MicroUsd field to given value.

### HasMicroUsd

`func (o *AgentAgentRunView) HasMicroUsd() bool`

HasMicroUsd returns a boolean if a field has been set.

### GetModel

`func (o *AgentAgentRunView) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentAgentRunView) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentAgentRunView) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AgentAgentRunView) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOutput

`func (o *AgentAgentRunView) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *AgentAgentRunView) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *AgentAgentRunView) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *AgentAgentRunView) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetPromptTokens

`func (o *AgentAgentRunView) GetPromptTokens() int64`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *AgentAgentRunView) GetPromptTokensOk() (*int64, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *AgentAgentRunView) SetPromptTokens(v int64)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *AgentAgentRunView) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetStatus

`func (o *AgentAgentRunView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentAgentRunView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentAgentRunView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentAgentRunView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToolCalls

`func (o *AgentAgentRunView) GetToolCalls() int64`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *AgentAgentRunView) GetToolCallsOk() (*int64, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *AgentAgentRunView) SetToolCalls(v int64)`

SetToolCalls sets ToolCalls field to given value.

### HasToolCalls

`func (o *AgentAgentRunView) HasToolCalls() bool`

HasToolCalls returns a boolean if a field has been set.

### GetTraceId

`func (o *AgentAgentRunView) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *AgentAgentRunView) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *AgentAgentRunView) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *AgentAgentRunView) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


