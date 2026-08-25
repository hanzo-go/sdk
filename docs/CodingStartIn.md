# CodingStartIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to **string** | After names a previous run&#39;s session, and starts this one from where that one stopped instead of from the repository&#39;s default. It is how a follow-up instruction — \&quot;now add tests for it\&quot; — builds on work already done rather than beginning again on a fresh clone.  It sets the base and nothing else, so this run still writes its OWN branch. One run, one branch: a run that wrote back onto an earlier run&#39;s branch would break the rule the forge&#39;s ref policy is built on, and would leave two turns of work with one name to review.  A caller who already knows the branch may pass Base directly; this exists because the branch is derived from a session id and nobody should have to know how. Base wins if both are given. | [optional] 
**AgentRef** | Pointer to **string** | AgentRef names a configured agent to run as, which is how an org pins a harness, a model and a prompt to a name. Empty runs the default agent. | [optional] 
**Base** | Pointer to **string** | Base is the branch to start from. Empty takes the repository&#39;s default. The run never writes here — it writes the agent branch it answers with. | [optional] 
**Desktop** | Pointer to **bool** | Desktop asks for a run with a SCREEN — an image carrying an X server — for a task that has to drive a browser or another windowed program. False, the default, is a headless checkout, which is what writing code needs. | [optional] 
**Project** | Pointer to **string** | Project scopes the run to one board&#39;s work when the org keeps more than one. Empty is the org&#39;s default. | [optional] 
**Prompt** | Pointer to **string** | Prompt is the task, in the words you would use with a colleague who has the checkout open. It is the whole instruction: there is no second field for context, and a prompt that names files and the outcome it wants gets a run that does not have to guess either. | [optional] 
**ReplyChannel** | Pointer to **string** | ReplyChannel / ReplyThread are WHERE THE RUN NARRATES ITSELF, when the surface that started it has somewhere for it to talk. Empty means nobody is listening and the run simply does not narrate — which is the app surface&#39;s case, because /v1/agents/coding hands back a session id and the session stream is a better progress feed than any message could be.  It is an ADDRESS and not a token: the engine says \&quot;put this text there\&quot;, and the process that owns the workspace&#39;s bot credential is the one that actually posts. So a run reports into a Slack thread without the engine ever holding the token that could post anywhere else in that workspace. | [optional] 
**ReplyThread** | Pointer to **string** | ReplyThread narrows that address to one THREAD inside the channel: on Slack it is the parent message&#39;s ts, the same value a reply carries as thread_ts. Empty puts the run&#39;s status line at the top level of the channel instead.  The channel is what decides whether a run narrates at all, so this on its own addresses nothing — a thread with no ReplyChannel is a run nobody hears. | [optional] 
**Repo** | Pointer to **string** | Repo is what to work on, as &#x60;owner/name&#x60; in the caller&#39;s own org. The engine resolves the clone URL and the push credential from the org itself, so this says WHICH repository and never how to reach it. | [optional] 
**TargetId** | Pointer to **string** | TargetID routes the run to a registered machine the org has claimed instead of to a sandbox in our cluster. Empty runs it here, which is the usual case. | [optional] 
**TimeoutSeconds** | Pointer to **int32** | TimeoutSeconds bounds the whole run. Unset takes the default budget; a run that hits the bound is stopped and reports what it had done by then. | [optional] 
**Tool** | Pointer to **string** | Tool is which harness runs the prompt — dev | claude | codex | python | node — and Desktop is whether the run needs a screen. Both are empty by default, which is &#x60;dev&#x60; with no screen, and that default is what every caller gets until it says otherwise.  They are two fields because they are two questions. The harness decides what argv starts; the screen decides which image carries an X server. A caller may want claude WITH a browser it can see, and a single enum would have made that combination unsayable. | [optional] 

## Methods

### NewCodingStartIn

`func NewCodingStartIn() *CodingStartIn`

NewCodingStartIn instantiates a new CodingStartIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodingStartInWithDefaults

`func NewCodingStartInWithDefaults() *CodingStartIn`

NewCodingStartInWithDefaults instantiates a new CodingStartIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *CodingStartIn) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *CodingStartIn) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *CodingStartIn) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *CodingStartIn) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetAgentRef

`func (o *CodingStartIn) GetAgentRef() string`

GetAgentRef returns the AgentRef field if non-nil, zero value otherwise.

### GetAgentRefOk

`func (o *CodingStartIn) GetAgentRefOk() (*string, bool)`

GetAgentRefOk returns a tuple with the AgentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentRef

`func (o *CodingStartIn) SetAgentRef(v string)`

SetAgentRef sets AgentRef field to given value.

### HasAgentRef

`func (o *CodingStartIn) HasAgentRef() bool`

HasAgentRef returns a boolean if a field has been set.

### GetBase

`func (o *CodingStartIn) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *CodingStartIn) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *CodingStartIn) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *CodingStartIn) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetDesktop

`func (o *CodingStartIn) GetDesktop() bool`

GetDesktop returns the Desktop field if non-nil, zero value otherwise.

### GetDesktopOk

`func (o *CodingStartIn) GetDesktopOk() (*bool, bool)`

GetDesktopOk returns a tuple with the Desktop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktop

`func (o *CodingStartIn) SetDesktop(v bool)`

SetDesktop sets Desktop field to given value.

### HasDesktop

`func (o *CodingStartIn) HasDesktop() bool`

HasDesktop returns a boolean if a field has been set.

### GetProject

`func (o *CodingStartIn) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CodingStartIn) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CodingStartIn) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CodingStartIn) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPrompt

`func (o *CodingStartIn) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *CodingStartIn) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *CodingStartIn) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *CodingStartIn) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetReplyChannel

`func (o *CodingStartIn) GetReplyChannel() string`

GetReplyChannel returns the ReplyChannel field if non-nil, zero value otherwise.

### GetReplyChannelOk

`func (o *CodingStartIn) GetReplyChannelOk() (*string, bool)`

GetReplyChannelOk returns a tuple with the ReplyChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyChannel

`func (o *CodingStartIn) SetReplyChannel(v string)`

SetReplyChannel sets ReplyChannel field to given value.

### HasReplyChannel

`func (o *CodingStartIn) HasReplyChannel() bool`

HasReplyChannel returns a boolean if a field has been set.

### GetReplyThread

`func (o *CodingStartIn) GetReplyThread() string`

GetReplyThread returns the ReplyThread field if non-nil, zero value otherwise.

### GetReplyThreadOk

`func (o *CodingStartIn) GetReplyThreadOk() (*string, bool)`

GetReplyThreadOk returns a tuple with the ReplyThread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyThread

`func (o *CodingStartIn) SetReplyThread(v string)`

SetReplyThread sets ReplyThread field to given value.

### HasReplyThread

`func (o *CodingStartIn) HasReplyThread() bool`

HasReplyThread returns a boolean if a field has been set.

### GetRepo

`func (o *CodingStartIn) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CodingStartIn) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CodingStartIn) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CodingStartIn) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetTargetId

`func (o *CodingStartIn) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *CodingStartIn) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *CodingStartIn) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *CodingStartIn) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *CodingStartIn) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *CodingStartIn) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *CodingStartIn) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *CodingStartIn) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetTool

`func (o *CodingStartIn) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *CodingStartIn) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *CodingStartIn) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *CodingStartIn) HasTool() bool`

HasTool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


