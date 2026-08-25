# CreateAgentIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputeRef** | Pointer to **string** | ComputeRef optionally binds this bot to a visor machine. Opaque here, bounded at 256 characters, and not resolved — this package stores the reference and the binding&#39;s lifecycle belongs elsewhere. | [optional] 
**Description** | Pointer to **string** | Description is the one line published as the description of the &#x60;agent_&lt;name&gt;&#x60; tool, which is how another agent decides whether to call this one. Optional, and worth writing for exactly that reason. | [optional] 
**ExecutionMode** | Pointer to **string** | ExecutionMode is one-shot or long-running. Empty takes one-shot, which runs only when something POSTs to it. long-running additionally requires Schedule, and counts against a per-org cap that answers 409 when it is full. | [optional] 
**Instructions** | Pointer to **string** | Instructions is the system prompt, up to 32 KiB, stored verbatim. This is what the model reads; Description is what other CALLERS read. | [optional] 
**Model** | Pointer to **string** | Model names the model to run on. Omit it to take the deployment&#39;s configured default; name one and it is checked against the gateway&#39;s served catalogue here, so a model this deployment cannot serve is refused now rather than at the first run. Stored under our own name for it, whatever spelling arrives. | [optional] 
**Name** | Pointer to **string** | Name is the agent&#39;s org-unique handle and the only required field. It must match ^[A-Za-z0-9][A-Za-z0-9._-]{0,63}$, and a name already taken in this org is a 409 rather than an overwrite. It is permanent: no update route moves it. | [optional] 
**Schedule** | Pointer to **string** | Schedule is the 5-field cron a long-running agent fires on, parsed here so a bad expression is a 400 and not an agent that silently never runs. Required with long-running; DISCARDED for one-shot rather than stored unused. | [optional] 
**ServiceAccountId** | Pointer to **string** | ServiceAccountID optionally names the IAM agent service account (&lt;org&gt;-&lt;agent&gt;) a scheduled run should be billed AS, so an autonomous run is attributable to a principal rather than only to the org. Same 256-character bound, also unresolved here. | [optional] 
**Tools** | Pointer to **[]string** | Tools are the tool names this agent may call. Omitted or empty grants NONE — that default is the agent&#39;s authority and is not widened anywhere. The single entry \&quot;*\&quot; means whatever the fleet&#39;s MCP server serves at the time of each run. | [optional] 

## Methods

### NewCreateAgentIn

`func NewCreateAgentIn() *CreateAgentIn`

NewCreateAgentIn instantiates a new CreateAgentIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAgentInWithDefaults

`func NewCreateAgentInWithDefaults() *CreateAgentIn`

NewCreateAgentInWithDefaults instantiates a new CreateAgentIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeRef

`func (o *CreateAgentIn) GetComputeRef() string`

GetComputeRef returns the ComputeRef field if non-nil, zero value otherwise.

### GetComputeRefOk

`func (o *CreateAgentIn) GetComputeRefOk() (*string, bool)`

GetComputeRefOk returns a tuple with the ComputeRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRef

`func (o *CreateAgentIn) SetComputeRef(v string)`

SetComputeRef sets ComputeRef field to given value.

### HasComputeRef

`func (o *CreateAgentIn) HasComputeRef() bool`

HasComputeRef returns a boolean if a field has been set.

### GetDescription

`func (o *CreateAgentIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAgentIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAgentIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAgentIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExecutionMode

`func (o *CreateAgentIn) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *CreateAgentIn) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *CreateAgentIn) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *CreateAgentIn) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetInstructions

`func (o *CreateAgentIn) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *CreateAgentIn) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *CreateAgentIn) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *CreateAgentIn) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetModel

`func (o *CreateAgentIn) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CreateAgentIn) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CreateAgentIn) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CreateAgentIn) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *CreateAgentIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAgentIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAgentIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateAgentIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchedule

`func (o *CreateAgentIn) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CreateAgentIn) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CreateAgentIn) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CreateAgentIn) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *CreateAgentIn) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *CreateAgentIn) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *CreateAgentIn) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *CreateAgentIn) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### GetTools

`func (o *CreateAgentIn) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *CreateAgentIn) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *CreateAgentIn) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *CreateAgentIn) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


