# AgentCreateAgentIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **string** | Avatar and Emoji are how the agent APPEARS. An image wins when both are given — it is the thing somebody made — and both empty leaves the agent drawn as its initial. Validated by iam/pkg/schema, the same rule a person&#39;s avatar passes, so the 96 KiB bound and the accepted URL forms are stated once for every subject that has a face. | [optional] 
**CapMicroUsd** | Pointer to **int64** | The budget, required. An agent is not creatable without one: a zero cap would mean no limit and no record of its spend, which is what a budget exists to remove. Integer micro-USD; the CLI converts dollars before sending. | [optional] 
**ComputeRef** | Pointer to **string** | ComputeRef optionally binds this bot to a visor machine. Opaque here, bounded at 256 characters, and not resolved — this package stores the reference and the binding&#39;s lifecycle belongs elsewhere. | [optional] 
**Description** | Pointer to **string** | Description is the one line published as the description of the &#x60;agent_&lt;name&gt;&#x60; tool, which is how another agent decides whether to call this one. Optional, and worth writing for exactly that reason. | [optional] 
**Emoji** | Pointer to **string** | Emoji is the single glyph shown when there is no image. An image WINS when both are given — it is the thing somebody made — and both empty leaves the agent drawn as its initial. | [optional] 
**ExecutionMode** | Pointer to **string** | ExecutionMode is one-shot or long-running. Empty takes one-shot, which runs only when something POSTs to it. long-running additionally requires Schedule, and counts against a per-org cap that answers 409 when it is full. | [optional] 
**Instructions** | Pointer to **string** | Instructions is the system prompt, up to 32 KiB, stored verbatim. This is what the model reads; Description is what other CALLERS read. | [optional] 
**MaxTaskMicroUsd** | Pointer to **int64** | MaxTaskMicroUSD is the ceiling for a single run, in micro-USD. A session cannot exceed it even when the period cap still has room, so one runaway task cannot consume a month. | [optional] 
**Model** | Pointer to **string** | Model names the model to run on. Omit it to take the deployment&#39;s configured default; name one and it is checked against the gateway&#39;s served catalogue here, so a model this deployment cannot serve is refused now rather than at the first run. Stored under our own name for it, whatever spelling arrives. | [optional] 
**Name** | Pointer to **string** | Name is the agent&#39;s org-unique handle and the only required field. It must match ^[A-Za-z0-9][A-Za-z0-9._-]{0,63}$, and a name already taken in this org is a 409 rather than an overwrite. It is permanent: no update route moves it. | [optional] 
**Parent** | Pointer to **string** | Parent is the agent spawning this one — its id or name, in the caller&#39;s org. The new agent answers to the org its parent answers to, and a parent outside the caller&#39;s org is not found: an agent is never spawned into a principal its creator does not answer to. Omitted for an agent a person defines. Set once; no update moves it. | [optional] 
**Period** | Pointer to **string** | Period is the window the cap resets on: day, week or month. | [optional] 
**Schedule** | Pointer to **string** | Schedule is the 5-field cron a long-running agent fires on, parsed here so a bad expression is a 400 and not an agent that silently never runs. Required with long-running; DISCARDED for one-shot rather than stored unused. | [optional] 
**ServiceAccountId** | Pointer to **string** | ServiceAccountID optionally names the IAM agent service account (&lt;org&gt;-&lt;agent&gt;) a scheduled run should be billed AS, so an autonomous run is attributable to a principal rather than only to the org. Same 256-character bound, also unresolved here. | [optional] 
**Tools** | Pointer to **[]string** | Tools are the tool names this agent may call. Omitted or empty grants NONE — that default is the agent&#39;s authority and is not widened anywhere. The single entry \&quot;*\&quot; means whatever the fleet&#39;s MCP server serves at the time of each run. | [optional] 

## Methods

### NewAgentCreateAgentIn

`func NewAgentCreateAgentIn() *AgentCreateAgentIn`

NewAgentCreateAgentIn instantiates a new AgentCreateAgentIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCreateAgentInWithDefaults

`func NewAgentCreateAgentInWithDefaults() *AgentCreateAgentIn`

NewAgentCreateAgentInWithDefaults instantiates a new AgentCreateAgentIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *AgentCreateAgentIn) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *AgentCreateAgentIn) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *AgentCreateAgentIn) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *AgentCreateAgentIn) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetCapMicroUsd

`func (o *AgentCreateAgentIn) GetCapMicroUsd() int64`

GetCapMicroUsd returns the CapMicroUsd field if non-nil, zero value otherwise.

### GetCapMicroUsdOk

`func (o *AgentCreateAgentIn) GetCapMicroUsdOk() (*int64, bool)`

GetCapMicroUsdOk returns a tuple with the CapMicroUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapMicroUsd

`func (o *AgentCreateAgentIn) SetCapMicroUsd(v int64)`

SetCapMicroUsd sets CapMicroUsd field to given value.

### HasCapMicroUsd

`func (o *AgentCreateAgentIn) HasCapMicroUsd() bool`

HasCapMicroUsd returns a boolean if a field has been set.

### GetComputeRef

`func (o *AgentCreateAgentIn) GetComputeRef() string`

GetComputeRef returns the ComputeRef field if non-nil, zero value otherwise.

### GetComputeRefOk

`func (o *AgentCreateAgentIn) GetComputeRefOk() (*string, bool)`

GetComputeRefOk returns a tuple with the ComputeRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRef

`func (o *AgentCreateAgentIn) SetComputeRef(v string)`

SetComputeRef sets ComputeRef field to given value.

### HasComputeRef

`func (o *AgentCreateAgentIn) HasComputeRef() bool`

HasComputeRef returns a boolean if a field has been set.

### GetDescription

`func (o *AgentCreateAgentIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentCreateAgentIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentCreateAgentIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentCreateAgentIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmoji

`func (o *AgentCreateAgentIn) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *AgentCreateAgentIn) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *AgentCreateAgentIn) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *AgentCreateAgentIn) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetExecutionMode

`func (o *AgentCreateAgentIn) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *AgentCreateAgentIn) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *AgentCreateAgentIn) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *AgentCreateAgentIn) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetInstructions

`func (o *AgentCreateAgentIn) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AgentCreateAgentIn) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AgentCreateAgentIn) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AgentCreateAgentIn) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetMaxTaskMicroUsd

`func (o *AgentCreateAgentIn) GetMaxTaskMicroUsd() int64`

GetMaxTaskMicroUsd returns the MaxTaskMicroUsd field if non-nil, zero value otherwise.

### GetMaxTaskMicroUsdOk

`func (o *AgentCreateAgentIn) GetMaxTaskMicroUsdOk() (*int64, bool)`

GetMaxTaskMicroUsdOk returns a tuple with the MaxTaskMicroUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskMicroUsd

`func (o *AgentCreateAgentIn) SetMaxTaskMicroUsd(v int64)`

SetMaxTaskMicroUsd sets MaxTaskMicroUsd field to given value.

### HasMaxTaskMicroUsd

`func (o *AgentCreateAgentIn) HasMaxTaskMicroUsd() bool`

HasMaxTaskMicroUsd returns a boolean if a field has been set.

### GetModel

`func (o *AgentCreateAgentIn) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentCreateAgentIn) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentCreateAgentIn) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AgentCreateAgentIn) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *AgentCreateAgentIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentCreateAgentIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentCreateAgentIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentCreateAgentIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *AgentCreateAgentIn) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AgentCreateAgentIn) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AgentCreateAgentIn) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AgentCreateAgentIn) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPeriod

`func (o *AgentCreateAgentIn) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *AgentCreateAgentIn) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *AgentCreateAgentIn) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *AgentCreateAgentIn) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetSchedule

`func (o *AgentCreateAgentIn) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AgentCreateAgentIn) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AgentCreateAgentIn) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AgentCreateAgentIn) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *AgentCreateAgentIn) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *AgentCreateAgentIn) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *AgentCreateAgentIn) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *AgentCreateAgentIn) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### GetTools

`func (o *AgentCreateAgentIn) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AgentCreateAgentIn) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AgentCreateAgentIn) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AgentCreateAgentIn) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


