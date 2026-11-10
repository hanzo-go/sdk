# AgentAgentDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **string** | Avatar is an image the agent is drawn as — a link to one, or the bytes inline as a data URL, up to 96 KiB. Emoji is the one glyph a caller picked when they had no image. At most one is ever set; neither means the agent is drawn as its initial, the same way a person with no photo is. Both are iam/pkg/schema&#39;s Mark, so a face means the same thing on an agent as it does on a person or an org. Avatar is the agent&#39;s picture: an image URL, or the image itself inline as a data URL up to 96 KiB. Empty when the agent has no image. | [optional] 
**CapMicroUsd** | Pointer to **int64** | CapMicroUSD is the total this agent may spend within one Period, as an integer number of micro-USD (1,000,000 &#x3D; $1). It is required at creation: a cap of zero would mean no limit and no per-agent spend record at all. | [optional] 
**ComputeRef** | Pointer to **string** | ComputeRef is the visor machine this bot is bound to, opaque here: this package stores and echoes it, and the binding&#39;s lifecycle belongs elsewhere. Empty means unbound, which is what every one-shot agent is. | [optional] 
**ConsumedMicroUsd** | Pointer to **int64** | ConsumedMicroUSD is what has been spent in the current period, in micro-USD. It is settled from what the gateway reported, not from the quote. | [optional] 
**CreatedAt** | Pointer to **string** | CreatedAt is when the agent was defined, RFC 3339 in UTC to the second. | [optional] 
**Description** | Pointer to **string** | Description is the one line another agent reads when deciding whether to call this one: the tool catalogue publishes it as the description of &#x60;agent_&lt;name&gt;&#x60;, falling back to \&quot;agent &lt;name&gt;\&quot; when it is empty. It is not part of the prompt — Instructions is — so writing the behaviour here reaches the caller and not the model. | [optional] 
**Emoji** | Pointer to **string** | Emoji is the single glyph a caller picked when they had no image. At most one of avatar and emoji is ever set; neither means the agent is drawn as its initial, the same way a person with no photo is. | [optional] 
**ExecutionMode** | Pointer to **string** | ExecutionMode is one-shot or long-running, and it decides who may start this agent. one-shot runs only when something POSTs to it; long-running is additionally invoked by the scheduler on Schedule, once a minute against the cron. An org&#39;s long-running agents are capped, so a switch INTO it can be refused with 409. | [optional] 
**Id** | Pointer to **string** | ID is the agent&#39;s stable handle, minted here as \&quot;agent_\&quot; + 32 hex characters of crypto/rand. A caller cannot choose it, and it never changes — unlike Name, which is the other way to address the same agent. | [optional] 
**Instructions** | Pointer to **string** | Instructions is the agent&#39;s system prompt, verbatim, up to 32 KiB. It is the one field the list read withholds, because it is the agent&#39;s whole behaviour and a page of them would be a page of prompts. | [optional] 
**Lineage** | Pointer to **[]string** | Lineage is every agent above this one, nearest first — the parent, its parent, and so on to the agent a person defined. All of them are in this org, so the principal this agent answers to is theirs. | [optional] 
**MaxTaskMicroUsd** | Pointer to **int64** | MaxTaskMicroUSD is the ceiling for a single run, in micro-USD. A session cannot exceed it even when the period cap still has room, so one runaway task cannot consume a month. | [optional] 
**Model** | Pointer to **string** | Model is the Zen model this agent runs on, and it is always OUR name for it: writes normalize through cloud.ZenModel and the read normalizes again, so an upstream family name never leaves here even from a row written before that rule existed. A create that named none took the deployment&#39;s configured default, so this is where a caller learns which model it actually got. | [optional] 
**Name** | Pointer to **string** | Name is the agent&#39;s org-unique handle, matching ^[A-Za-z0-9][A-Za-z0-9._-]{0,63}$. It addresses the agent everywhere ID does, it is what a run row records, and it is the suffix of the &#x60;agent_&lt;name&gt;&#x60; tool other agents call this one by. Set once at create; no update route moves it, because moving it would orphan that history. | [optional] 
**Parent** | Pointer to **string** | Parent is the agent that spawned this one; absent for an agent a person defined. | [optional] 
**Period** | Pointer to **string** | Period is the window the cap resets on: day, week or month. | [optional] 
**RecentRuns** | Pointer to [**[]AgentAgentRunView**](AgentAgentRunView.md) | RecentRuns is the agent&#39;s 20 most recent executions, newest first. It is a window on the history, not the history: the count beside it is &#x60;runs&#x60;. | [optional] 
**Runs** | Pointer to **int64** | Runs is how many executions the org has recorded against this agent, counted at read time. The list and update reads count the WHOLE history; the detail read reports the size of the RecentRuns page it carries, which stops at 20 — so a detail row saying 20 means \&quot;at least 20\&quot;, not \&quot;exactly 20\&quot;. | [optional] 
**Schedule** | Pointer to **string** | Schedule is the 5-field cron the scheduler fires a long-running agent on, evaluated once a minute. Required for long-running and DROPPED for one-shot — a one-shot agent&#39;s schedule is not stored, so absence here is the mode&#39;s answer rather than a value nobody set. | [optional] 
**ServiceAccountId** | Pointer to **string** | ServiceAccountID is the IAM agent service account (&lt;org&gt;-&lt;agent&gt;) a scheduled run is billed AS. It is what makes an autonomous run attributable to a principal rather than only to the org; empty means the org itself wears the spend. | [optional] 
**Status** | Pointer to **string** | Status is the agent&#39;s readiness: \&quot;ready\&quot;, or \&quot;retired\&quot; for a built-in whose member has left Hanzo&#39;s team. A retired agent is left out of the list and still answers a get by its handle. Server-set; no route accepts it. | [optional] 
**Tools** | Pointer to **[]string** | Tools are the tool names this agent may call, and the list IS the authority: an agent that declares none gets none. The single entry \&quot;*\&quot; means whatever the fleet&#39;s MCP server serves at the moment of the run, resolved per run rather than frozen here, which is how the default assistant reaches subsystems that shipped after it was defined. Empty array, never null. | [optional] 
**UpdatedAt** | Pointer to **string** | UpdatedAt is the last time any field above was written, same format. It moves on an update to the DEFINITION and never on a run, so a busy agent nobody has edited keeps an old one. | [optional] 

## Methods

### NewAgentAgentDetail

`func NewAgentAgentDetail() *AgentAgentDetail`

NewAgentAgentDetail instantiates a new AgentAgentDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentAgentDetailWithDefaults

`func NewAgentAgentDetailWithDefaults() *AgentAgentDetail`

NewAgentAgentDetailWithDefaults instantiates a new AgentAgentDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *AgentAgentDetail) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *AgentAgentDetail) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *AgentAgentDetail) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *AgentAgentDetail) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetCapMicroUsd

`func (o *AgentAgentDetail) GetCapMicroUsd() int64`

GetCapMicroUsd returns the CapMicroUsd field if non-nil, zero value otherwise.

### GetCapMicroUsdOk

`func (o *AgentAgentDetail) GetCapMicroUsdOk() (*int64, bool)`

GetCapMicroUsdOk returns a tuple with the CapMicroUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapMicroUsd

`func (o *AgentAgentDetail) SetCapMicroUsd(v int64)`

SetCapMicroUsd sets CapMicroUsd field to given value.

### HasCapMicroUsd

`func (o *AgentAgentDetail) HasCapMicroUsd() bool`

HasCapMicroUsd returns a boolean if a field has been set.

### GetComputeRef

`func (o *AgentAgentDetail) GetComputeRef() string`

GetComputeRef returns the ComputeRef field if non-nil, zero value otherwise.

### GetComputeRefOk

`func (o *AgentAgentDetail) GetComputeRefOk() (*string, bool)`

GetComputeRefOk returns a tuple with the ComputeRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRef

`func (o *AgentAgentDetail) SetComputeRef(v string)`

SetComputeRef sets ComputeRef field to given value.

### HasComputeRef

`func (o *AgentAgentDetail) HasComputeRef() bool`

HasComputeRef returns a boolean if a field has been set.

### GetConsumedMicroUsd

`func (o *AgentAgentDetail) GetConsumedMicroUsd() int64`

GetConsumedMicroUsd returns the ConsumedMicroUsd field if non-nil, zero value otherwise.

### GetConsumedMicroUsdOk

`func (o *AgentAgentDetail) GetConsumedMicroUsdOk() (*int64, bool)`

GetConsumedMicroUsdOk returns a tuple with the ConsumedMicroUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedMicroUsd

`func (o *AgentAgentDetail) SetConsumedMicroUsd(v int64)`

SetConsumedMicroUsd sets ConsumedMicroUsd field to given value.

### HasConsumedMicroUsd

`func (o *AgentAgentDetail) HasConsumedMicroUsd() bool`

HasConsumedMicroUsd returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AgentAgentDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AgentAgentDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AgentAgentDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AgentAgentDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *AgentAgentDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentAgentDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentAgentDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentAgentDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmoji

`func (o *AgentAgentDetail) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *AgentAgentDetail) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *AgentAgentDetail) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *AgentAgentDetail) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetExecutionMode

`func (o *AgentAgentDetail) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *AgentAgentDetail) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *AgentAgentDetail) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *AgentAgentDetail) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetId

`func (o *AgentAgentDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentAgentDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentAgentDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AgentAgentDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstructions

`func (o *AgentAgentDetail) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AgentAgentDetail) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AgentAgentDetail) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AgentAgentDetail) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetLineage

`func (o *AgentAgentDetail) GetLineage() []string`

GetLineage returns the Lineage field if non-nil, zero value otherwise.

### GetLineageOk

`func (o *AgentAgentDetail) GetLineageOk() (*[]string, bool)`

GetLineageOk returns a tuple with the Lineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineage

`func (o *AgentAgentDetail) SetLineage(v []string)`

SetLineage sets Lineage field to given value.

### HasLineage

`func (o *AgentAgentDetail) HasLineage() bool`

HasLineage returns a boolean if a field has been set.

### GetMaxTaskMicroUsd

`func (o *AgentAgentDetail) GetMaxTaskMicroUsd() int64`

GetMaxTaskMicroUsd returns the MaxTaskMicroUsd field if non-nil, zero value otherwise.

### GetMaxTaskMicroUsdOk

`func (o *AgentAgentDetail) GetMaxTaskMicroUsdOk() (*int64, bool)`

GetMaxTaskMicroUsdOk returns a tuple with the MaxTaskMicroUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskMicroUsd

`func (o *AgentAgentDetail) SetMaxTaskMicroUsd(v int64)`

SetMaxTaskMicroUsd sets MaxTaskMicroUsd field to given value.

### HasMaxTaskMicroUsd

`func (o *AgentAgentDetail) HasMaxTaskMicroUsd() bool`

HasMaxTaskMicroUsd returns a boolean if a field has been set.

### GetModel

`func (o *AgentAgentDetail) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentAgentDetail) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentAgentDetail) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AgentAgentDetail) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *AgentAgentDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentAgentDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentAgentDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentAgentDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *AgentAgentDetail) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AgentAgentDetail) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AgentAgentDetail) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AgentAgentDetail) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPeriod

`func (o *AgentAgentDetail) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *AgentAgentDetail) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *AgentAgentDetail) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *AgentAgentDetail) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetRecentRuns

`func (o *AgentAgentDetail) GetRecentRuns() []AgentAgentRunView`

GetRecentRuns returns the RecentRuns field if non-nil, zero value otherwise.

### GetRecentRunsOk

`func (o *AgentAgentDetail) GetRecentRunsOk() (*[]AgentAgentRunView, bool)`

GetRecentRunsOk returns a tuple with the RecentRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentRuns

`func (o *AgentAgentDetail) SetRecentRuns(v []AgentAgentRunView)`

SetRecentRuns sets RecentRuns field to given value.

### HasRecentRuns

`func (o *AgentAgentDetail) HasRecentRuns() bool`

HasRecentRuns returns a boolean if a field has been set.

### GetRuns

`func (o *AgentAgentDetail) GetRuns() int64`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *AgentAgentDetail) GetRunsOk() (*int64, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *AgentAgentDetail) SetRuns(v int64)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *AgentAgentDetail) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### GetSchedule

`func (o *AgentAgentDetail) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AgentAgentDetail) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AgentAgentDetail) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AgentAgentDetail) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *AgentAgentDetail) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *AgentAgentDetail) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *AgentAgentDetail) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *AgentAgentDetail) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### GetStatus

`func (o *AgentAgentDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentAgentDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentAgentDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentAgentDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTools

`func (o *AgentAgentDetail) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AgentAgentDetail) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AgentAgentDetail) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AgentAgentDetail) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AgentAgentDetail) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AgentAgentDetail) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AgentAgentDetail) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AgentAgentDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


