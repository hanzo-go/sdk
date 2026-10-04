# AgentUpdateAgentIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **string** | Avatar and Emoji re-draw the agent. Sending either replaces the pair, so setting an image clears a glyph and \&quot;\&quot; for both goes back to the initial — there is no state where a row holds two answers. | [optional] 
**CapMicroUsd** | Pointer to **int64** | The budget, any part of it. Changing the period opens a new window. | [optional] 
**ComputeRef** | Pointer to **string** | ComputeRef re-binds (or, with \&quot;\&quot;, unbinds) the visor machine. Opaque here. | [optional] 
**Description** | Pointer to **string** | Description replaces the line other agents read in the tool catalogue. | [optional] 
**Emoji** | Pointer to **string** | Emoji re-draws the agent as a glyph. Sending either of the pair replaces BOTH, so setting a glyph clears an image and \&quot;\&quot; for both goes back to the initial — there is no state where a row holds two answers. | [optional] 
**ExecutionMode** | Pointer to **string** | ExecutionMode switches between one-shot and long-running. The RESULTING mode+schedule are validated together, so switching to long-running without a stored or supplied cron is refused rather than accepted into an agent the scheduler would skip forever. A switch INTO long-running counts against the per-org cap and can be a 409. | [optional] 
**Instructions** | Pointer to **string** | Instructions replaces the system prompt whole, up to 32 KiB. There is no append: a prompt is one text, and sending \&quot;\&quot; clears it. | [optional] 
**MaxTaskMicroUsd** | Pointer to **int64** | MaxTaskMicroUSD is the ceiling for a single run, in micro-USD. A session cannot exceed it even when the period cap still has room, so one runaway task cannot consume a month. | [optional] 
**Model** | Pointer to **string** | Model re-points the agent at another model, checked against the gateway&#39;s served catalogue exactly as create checks it. Empty STRING is refused — say nothing to keep the current one. Past runs keep the model that served them. | [optional] 
**Period** | Pointer to **string** | Period is the window the cap resets on: day, week or month. | [optional] 
**Ref** | Pointer to **string** | Ref is the agent to update — its public id or org-unique name, from the path. | [optional] 
**Schedule** | Pointer to **string** | Schedule replaces the cron. It is validated against the mode this update leaves behind, and dropped if that mode is one-shot. | [optional] 
**ServiceAccountId** | Pointer to **string** | ServiceAccountID re-points (or, with \&quot;\&quot;, clears) the IAM service account a scheduled run is billed as. Clearing it puts that spend back on the org. | [optional] 
**Tools** | Pointer to **[]string** | Tools replaces the whole allow-list, it does not add to it. Sending [] takes every tool away, which is the only way to say that. | [optional] 

## Methods

### NewAgentUpdateAgentIn

`func NewAgentUpdateAgentIn() *AgentUpdateAgentIn`

NewAgentUpdateAgentIn instantiates a new AgentUpdateAgentIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentUpdateAgentInWithDefaults

`func NewAgentUpdateAgentInWithDefaults() *AgentUpdateAgentIn`

NewAgentUpdateAgentInWithDefaults instantiates a new AgentUpdateAgentIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *AgentUpdateAgentIn) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *AgentUpdateAgentIn) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *AgentUpdateAgentIn) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *AgentUpdateAgentIn) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetCapMicroUsd

`func (o *AgentUpdateAgentIn) GetCapMicroUsd() int64`

GetCapMicroUsd returns the CapMicroUsd field if non-nil, zero value otherwise.

### GetCapMicroUsdOk

`func (o *AgentUpdateAgentIn) GetCapMicroUsdOk() (*int64, bool)`

GetCapMicroUsdOk returns a tuple with the CapMicroUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapMicroUsd

`func (o *AgentUpdateAgentIn) SetCapMicroUsd(v int64)`

SetCapMicroUsd sets CapMicroUsd field to given value.

### HasCapMicroUsd

`func (o *AgentUpdateAgentIn) HasCapMicroUsd() bool`

HasCapMicroUsd returns a boolean if a field has been set.

### GetComputeRef

`func (o *AgentUpdateAgentIn) GetComputeRef() string`

GetComputeRef returns the ComputeRef field if non-nil, zero value otherwise.

### GetComputeRefOk

`func (o *AgentUpdateAgentIn) GetComputeRefOk() (*string, bool)`

GetComputeRefOk returns a tuple with the ComputeRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRef

`func (o *AgentUpdateAgentIn) SetComputeRef(v string)`

SetComputeRef sets ComputeRef field to given value.

### HasComputeRef

`func (o *AgentUpdateAgentIn) HasComputeRef() bool`

HasComputeRef returns a boolean if a field has been set.

### GetDescription

`func (o *AgentUpdateAgentIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentUpdateAgentIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentUpdateAgentIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentUpdateAgentIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmoji

`func (o *AgentUpdateAgentIn) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *AgentUpdateAgentIn) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *AgentUpdateAgentIn) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *AgentUpdateAgentIn) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetExecutionMode

`func (o *AgentUpdateAgentIn) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *AgentUpdateAgentIn) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *AgentUpdateAgentIn) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *AgentUpdateAgentIn) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetInstructions

`func (o *AgentUpdateAgentIn) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AgentUpdateAgentIn) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AgentUpdateAgentIn) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AgentUpdateAgentIn) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetMaxTaskMicroUsd

`func (o *AgentUpdateAgentIn) GetMaxTaskMicroUsd() int64`

GetMaxTaskMicroUsd returns the MaxTaskMicroUsd field if non-nil, zero value otherwise.

### GetMaxTaskMicroUsdOk

`func (o *AgentUpdateAgentIn) GetMaxTaskMicroUsdOk() (*int64, bool)`

GetMaxTaskMicroUsdOk returns a tuple with the MaxTaskMicroUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskMicroUsd

`func (o *AgentUpdateAgentIn) SetMaxTaskMicroUsd(v int64)`

SetMaxTaskMicroUsd sets MaxTaskMicroUsd field to given value.

### HasMaxTaskMicroUsd

`func (o *AgentUpdateAgentIn) HasMaxTaskMicroUsd() bool`

HasMaxTaskMicroUsd returns a boolean if a field has been set.

### GetModel

`func (o *AgentUpdateAgentIn) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentUpdateAgentIn) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentUpdateAgentIn) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AgentUpdateAgentIn) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPeriod

`func (o *AgentUpdateAgentIn) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *AgentUpdateAgentIn) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *AgentUpdateAgentIn) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *AgentUpdateAgentIn) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetRef

`func (o *AgentUpdateAgentIn) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *AgentUpdateAgentIn) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *AgentUpdateAgentIn) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *AgentUpdateAgentIn) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSchedule

`func (o *AgentUpdateAgentIn) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AgentUpdateAgentIn) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AgentUpdateAgentIn) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AgentUpdateAgentIn) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *AgentUpdateAgentIn) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *AgentUpdateAgentIn) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *AgentUpdateAgentIn) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *AgentUpdateAgentIn) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### GetTools

`func (o *AgentUpdateAgentIn) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AgentUpdateAgentIn) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AgentUpdateAgentIn) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AgentUpdateAgentIn) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


