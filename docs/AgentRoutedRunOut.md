# AgentRoutedRunOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | Pointer to **string** | Base is the branch to start FROM. Empty means the repository&#39;s default — resolve it on the machine, since the machine is the one holding the clone. | [optional] 
**Branch** | Pointer to **string** | Branch is the ref the run must push its work to, and the ONLY one it is permitted to write: the forge&#39;s ref policy refuses anything else from this run&#39;s credential. It is decided at dispatch and exists before the work does. | [optional] 
**CloneUrl** | Pointer to **string** | CloneURL is how to fetch the repository. It carries NO credential — the machine authenticates with the git identity it already holds — which is why this whole shape is safe to hand to a claimed runner. | [optional] 
**Effort** | Pointer to **string** | Effort is how hard that model reasons: low, medium or high. Empty is the model&#39;s own default. | [optional] 
**Instructions** | Pointer to **string** | Instructions are the person&#39;s own instructions for their runs, as they set them in their settings: the machine&#39;s harness reads them as its user&#39;s, the way AGENTS.md or CLAUDE.md in its home is read. Empty when they set none. | [optional] 
**Model** | Pointer to **string** | Model is the model the person asked the run&#39;s agent to think with, by the id GET /v1/models lists. Empty is the machine&#39;s own default. | [optional] 
**Project** | Pointer to **string** | Project is the product slug the run is filed under, so the machine can tag what it produces. Empty when the dispatch named none. | [optional] 
**Prompt** | Pointer to **string** | Prompt is the task, in full, as the person wrote it. There is no second field for context. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository to work in and CloneURL is how to fetch it. | [optional] 
**SessionId** | Pointer to **string** | SessionID is the live session opened at dispatch; the machine streams its turns into it. | [optional] 
**TimeoutSeconds** | Pointer to **int64** | TimeoutSeconds bounds the run on the machine; 0 means the machine&#39;s own default. | [optional] 

## Methods

### NewAgentRoutedRunOut

`func NewAgentRoutedRunOut() *AgentRoutedRunOut`

NewAgentRoutedRunOut instantiates a new AgentRoutedRunOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRoutedRunOutWithDefaults

`func NewAgentRoutedRunOutWithDefaults() *AgentRoutedRunOut`

NewAgentRoutedRunOutWithDefaults instantiates a new AgentRoutedRunOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *AgentRoutedRunOut) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *AgentRoutedRunOut) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *AgentRoutedRunOut) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *AgentRoutedRunOut) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetBranch

`func (o *AgentRoutedRunOut) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *AgentRoutedRunOut) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *AgentRoutedRunOut) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *AgentRoutedRunOut) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetCloneUrl

`func (o *AgentRoutedRunOut) GetCloneUrl() string`

GetCloneUrl returns the CloneUrl field if non-nil, zero value otherwise.

### GetCloneUrlOk

`func (o *AgentRoutedRunOut) GetCloneUrlOk() (*string, bool)`

GetCloneUrlOk returns a tuple with the CloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneUrl

`func (o *AgentRoutedRunOut) SetCloneUrl(v string)`

SetCloneUrl sets CloneUrl field to given value.

### HasCloneUrl

`func (o *AgentRoutedRunOut) HasCloneUrl() bool`

HasCloneUrl returns a boolean if a field has been set.

### GetEffort

`func (o *AgentRoutedRunOut) GetEffort() string`

GetEffort returns the Effort field if non-nil, zero value otherwise.

### GetEffortOk

`func (o *AgentRoutedRunOut) GetEffortOk() (*string, bool)`

GetEffortOk returns a tuple with the Effort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffort

`func (o *AgentRoutedRunOut) SetEffort(v string)`

SetEffort sets Effort field to given value.

### HasEffort

`func (o *AgentRoutedRunOut) HasEffort() bool`

HasEffort returns a boolean if a field has been set.

### GetInstructions

`func (o *AgentRoutedRunOut) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AgentRoutedRunOut) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AgentRoutedRunOut) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AgentRoutedRunOut) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetModel

`func (o *AgentRoutedRunOut) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentRoutedRunOut) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentRoutedRunOut) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AgentRoutedRunOut) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetProject

`func (o *AgentRoutedRunOut) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AgentRoutedRunOut) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AgentRoutedRunOut) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *AgentRoutedRunOut) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPrompt

`func (o *AgentRoutedRunOut) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AgentRoutedRunOut) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AgentRoutedRunOut) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *AgentRoutedRunOut) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetRepo

`func (o *AgentRoutedRunOut) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *AgentRoutedRunOut) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *AgentRoutedRunOut) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *AgentRoutedRunOut) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetSessionId

`func (o *AgentRoutedRunOut) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *AgentRoutedRunOut) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *AgentRoutedRunOut) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *AgentRoutedRunOut) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *AgentRoutedRunOut) GetTimeoutSeconds() int64`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *AgentRoutedRunOut) GetTimeoutSecondsOk() (*int64, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *AgentRoutedRunOut) SetTimeoutSeconds(v int64)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *AgentRoutedRunOut) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


