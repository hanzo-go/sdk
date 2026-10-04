# AgentCodingStartIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to **string** | After names a previous run&#39;s session, and starts this one from where that one stopped instead of from the repository&#39;s default. It is how a follow-up instruction — \&quot;now add tests for it\&quot; — builds on work already done rather than beginning again on a fresh clone: the follow-up resumes the earlier run&#39;s sandbox and works in its tree, files, installs and all, while that sandbox is kept; one reaped by then is cloned again from the earlier branch.  It sets the base and nothing else, so this run still writes its OWN branch. One run, one branch: a run that wrote back onto an earlier run&#39;s branch would break the rule the forge&#39;s ref policy is built on, and would leave two turns of work with one name to review.  A caller who already knows the branch may pass Base directly; this exists because the branch is derived from a session id and nobody should have to know how. Base wins if both are given. | [optional] 
**AgentRef** | Pointer to **string** | AgentRef names a configured agent to run as, which is how an org pins a harness, a model and a prompt to a name. Empty runs the default agent. | [optional] 
**Base** | Pointer to **string** | Base is the branch to start from. Empty takes the repository&#39;s default. The run never writes here — it writes the agent branch it answers with. | [optional] 
**Desktop** | Pointer to **bool** | Desktop asks for a run with a SCREEN — an image carrying an X server — for a task that has to drive a browser or another windowed program. False, the default, is a headless checkout, which is what writing code needs. | [optional] 
**Effort** | Pointer to **string** | Effort is how hard that model reasons: low, medium or high. Empty takes the model&#39;s own default. Taken by the dev and codex harnesses, like Model. | [optional] 
**Issue** | Pointer to **int64** | Issue is the number of the issue on Repo this run resolves. Its pull request says &#x60;Closes #&lt;n&gt;&#x60;, so merging the work closes the issue. Zero is a run no issue asked for. | [optional] 
**Mode** | Pointer to **string** | Mode is what the run may do to the repository: &#x60;build&#x60; edits, commits and pushes its branch; &#x60;plan&#x60; reads the repository and answers with a plan, writing nothing — its sandbox is handed a credential that can only read and no push step runs, and the plan arrives as the run&#39;s final status. &#x60;setup&#x60; is a plan put to one question: its agent explores the checkout, installs and checks what it finds, and ends with the codebase&#39;s environment, which is kept as the proposal at /v1/environment/{repo} until someone saves it. Empty is &#x60;build&#x60;. A plan or a setup runs in our sandbox only: a claimed machine clones and pushes with its own credential, so one routed to a machine is refused. | [optional] 
**Model** | Pointer to **string** | Model is the model the run&#39;s agent thinks with, by the id GET /v1/models lists. Empty takes the platform&#39;s coding model. The dev and codex harnesses take it; the others do not think with a model we name, so naming one for them is refused. | [optional] 
**Project** | Pointer to **string** | Project scopes the run to one board&#39;s work when the org keeps more than one. Empty is the org&#39;s default. | [optional] 
**Prompt** | Pointer to **string** | Prompt is the task, in the words you would use with a colleague who has the checkout open. It is the whole instruction: there is no second field for context, and a prompt that names files and the outcome it wants gets a run that does not have to guess either. | [optional] 
**Reply** | Pointer to [**AgentReply**](AgentReply.md) | Reply is WHERE THE RUN NARRATES ITSELF, when the surface that started it has somewhere for it to talk. Empty means nobody is listening and the run simply does not narrate — which is the app surface&#39;s case, because /v1/agent/coding hands back a session id and the session stream is a better progress feed than any message could be.  IT IS NOT A FIELD A MODEL FILLS IN. A run asked for in a conversation is dispatched with the address that conversation is happening at, written over whatever arrived here (apps/agents/replyto.go). A model naming a channel is a model naming a room it was not spoken to in, so the value it chose is discarded rather than argued with — which is the same rule that left the org and the subject off this type entirely. | [optional] 
**Repo** | Pointer to **string** | Repo is what to work on: a repository&#39;s name, &#x60;owner/name&#x60; in the caller&#39;s own org, or the site a repository serves (&#x60;hanzo.ai&#x60;), as the person said it. The engine resolves it against the org&#39;s own repositories — a name that matches none is refused with the nearest ones — and finds the clone URL and the push credential from the org itself, so this says WHICH repository and never how to reach it.  A repository on GitHub is named by GitHub&#39;s address, &#x60;github.com/owner/name&#x60; — or &#x60;owner/name&#x60; when the owner is not the org&#39;s own — and runs there directly: it is admitted when the caller&#39;s own connected GitHub may push to it, or when the caller administers the org and the org&#39;s installation holds it, and refused otherwise.  A repository the caller can read and not push to — a public repository, or one their GitHub reads — is cloned read-only: the run works, commits in its workspace, and pushes nothing; its change is kept as the run&#39;s artifacts.  Empty runs in an empty workspace: a fresh git repository in the sandbox, with nothing cloned and nothing pushed. | [optional] 
**TargetId** | Pointer to **string** | TargetID routes the run to a registered machine the org has claimed instead of to a sandbox in our cluster. Empty runs it here, which is the usual case. | [optional] 
**Task** | Pointer to **string** | Task is Prompt under the name a caller asking for work tends to reach for. The two are one field: Prompt wins when both are given. | [optional] 
**TimeoutSeconds** | Pointer to **int64** | TimeoutSeconds bounds the whole run. Unset takes the default budget; a run that hits the bound is stopped and reports what it had done by then. | [optional] 
**Tool** | Pointer to **string** | Tool is which harness runs the prompt — dev | claude | codex | python | node — and Desktop is whether the run needs a screen. Both are empty by default, which is &#x60;dev&#x60; with no screen, and that default is what every caller gets until it says otherwise.  They are two fields because they are two questions. The harness decides what argv starts; the screen decides which image carries an X server. A caller may want claude WITH a browser it can see, and a single enum would have made that combination unsayable. | [optional] 

## Methods

### NewAgentCodingStartIn

`func NewAgentCodingStartIn() *AgentCodingStartIn`

NewAgentCodingStartIn instantiates a new AgentCodingStartIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCodingStartInWithDefaults

`func NewAgentCodingStartInWithDefaults() *AgentCodingStartIn`

NewAgentCodingStartInWithDefaults instantiates a new AgentCodingStartIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *AgentCodingStartIn) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *AgentCodingStartIn) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *AgentCodingStartIn) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *AgentCodingStartIn) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetAgentRef

`func (o *AgentCodingStartIn) GetAgentRef() string`

GetAgentRef returns the AgentRef field if non-nil, zero value otherwise.

### GetAgentRefOk

`func (o *AgentCodingStartIn) GetAgentRefOk() (*string, bool)`

GetAgentRefOk returns a tuple with the AgentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentRef

`func (o *AgentCodingStartIn) SetAgentRef(v string)`

SetAgentRef sets AgentRef field to given value.

### HasAgentRef

`func (o *AgentCodingStartIn) HasAgentRef() bool`

HasAgentRef returns a boolean if a field has been set.

### GetBase

`func (o *AgentCodingStartIn) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *AgentCodingStartIn) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *AgentCodingStartIn) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *AgentCodingStartIn) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetDesktop

`func (o *AgentCodingStartIn) GetDesktop() bool`

GetDesktop returns the Desktop field if non-nil, zero value otherwise.

### GetDesktopOk

`func (o *AgentCodingStartIn) GetDesktopOk() (*bool, bool)`

GetDesktopOk returns a tuple with the Desktop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktop

`func (o *AgentCodingStartIn) SetDesktop(v bool)`

SetDesktop sets Desktop field to given value.

### HasDesktop

`func (o *AgentCodingStartIn) HasDesktop() bool`

HasDesktop returns a boolean if a field has been set.

### GetEffort

`func (o *AgentCodingStartIn) GetEffort() string`

GetEffort returns the Effort field if non-nil, zero value otherwise.

### GetEffortOk

`func (o *AgentCodingStartIn) GetEffortOk() (*string, bool)`

GetEffortOk returns a tuple with the Effort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffort

`func (o *AgentCodingStartIn) SetEffort(v string)`

SetEffort sets Effort field to given value.

### HasEffort

`func (o *AgentCodingStartIn) HasEffort() bool`

HasEffort returns a boolean if a field has been set.

### GetIssue

`func (o *AgentCodingStartIn) GetIssue() int64`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *AgentCodingStartIn) GetIssueOk() (*int64, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *AgentCodingStartIn) SetIssue(v int64)`

SetIssue sets Issue field to given value.

### HasIssue

`func (o *AgentCodingStartIn) HasIssue() bool`

HasIssue returns a boolean if a field has been set.

### GetMode

`func (o *AgentCodingStartIn) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AgentCodingStartIn) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AgentCodingStartIn) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *AgentCodingStartIn) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetModel

`func (o *AgentCodingStartIn) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentCodingStartIn) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentCodingStartIn) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AgentCodingStartIn) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetProject

`func (o *AgentCodingStartIn) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AgentCodingStartIn) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AgentCodingStartIn) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *AgentCodingStartIn) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPrompt

`func (o *AgentCodingStartIn) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AgentCodingStartIn) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AgentCodingStartIn) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *AgentCodingStartIn) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetReply

`func (o *AgentCodingStartIn) GetReply() AgentReply`

GetReply returns the Reply field if non-nil, zero value otherwise.

### GetReplyOk

`func (o *AgentCodingStartIn) GetReplyOk() (*AgentReply, bool)`

GetReplyOk returns a tuple with the Reply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReply

`func (o *AgentCodingStartIn) SetReply(v AgentReply)`

SetReply sets Reply field to given value.

### HasReply

`func (o *AgentCodingStartIn) HasReply() bool`

HasReply returns a boolean if a field has been set.

### GetRepo

`func (o *AgentCodingStartIn) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *AgentCodingStartIn) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *AgentCodingStartIn) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *AgentCodingStartIn) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetTargetId

`func (o *AgentCodingStartIn) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *AgentCodingStartIn) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *AgentCodingStartIn) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *AgentCodingStartIn) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetTask

`func (o *AgentCodingStartIn) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *AgentCodingStartIn) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *AgentCodingStartIn) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *AgentCodingStartIn) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *AgentCodingStartIn) GetTimeoutSeconds() int64`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *AgentCodingStartIn) GetTimeoutSecondsOk() (*int64, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *AgentCodingStartIn) SetTimeoutSeconds(v int64)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *AgentCodingStartIn) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetTool

`func (o *AgentCodingStartIn) GetTool() string`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *AgentCodingStartIn) GetToolOk() (*string, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *AgentCodingStartIn) SetTool(v string)`

SetTool sets Tool field to given value.

### HasTool

`func (o *AgentCodingStartIn) HasTool() bool`

HasTool returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


