# CloudRegisterReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Actor** | Pointer to **string** |  | [optional] 
**Agent** | Pointer to **string** |  | [optional] 
**Cwd** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** | Execution context — where this session runs (all optional). | [optional] 
**ParentSessionId** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** | The readable build (provenance.go): which product this session builds, and whether its story may be read by the world. | [optional] 
**Provider** | Pointer to **string** | Account tag — the linked AI account this session ran under (login manager). | [optional] 
**Published** | Pointer to **bool** |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**TaskRunId** | Pointer to **string** |  | [optional] 
**TaskWorkflowId** | Pointer to **string** |  | [optional] 
**Terminal** | Pointer to **string** | Terminal is the URL this session&#39;s live terminal is published at, so the console can watch it. Optional — a session that publishes nothing is still a session. | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudRegisterReq

`func NewCloudRegisterReq() *CloudRegisterReq`

NewCloudRegisterReq instantiates a new CloudRegisterReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRegisterReqWithDefaults

`func NewCloudRegisterReqWithDefaults() *CloudRegisterReq`

NewCloudRegisterReqWithDefaults instantiates a new CloudRegisterReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudRegisterReq) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudRegisterReq) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudRegisterReq) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudRegisterReq) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetActor

`func (o *CloudRegisterReq) GetActor() string`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *CloudRegisterReq) GetActorOk() (*string, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *CloudRegisterReq) SetActor(v string)`

SetActor sets Actor field to given value.

### HasActor

`func (o *CloudRegisterReq) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetAgent

`func (o *CloudRegisterReq) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *CloudRegisterReq) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *CloudRegisterReq) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *CloudRegisterReq) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetCwd

`func (o *CloudRegisterReq) GetCwd() string`

GetCwd returns the Cwd field if non-nil, zero value otherwise.

### GetCwdOk

`func (o *CloudRegisterReq) GetCwdOk() (*string, bool)`

GetCwdOk returns a tuple with the Cwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwd

`func (o *CloudRegisterReq) SetCwd(v string)`

SetCwd sets Cwd field to given value.

### HasCwd

`func (o *CloudRegisterReq) HasCwd() bool`

HasCwd returns a boolean if a field has been set.

### GetHost

`func (o *CloudRegisterReq) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CloudRegisterReq) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CloudRegisterReq) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CloudRegisterReq) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetParentSessionId

`func (o *CloudRegisterReq) GetParentSessionId() string`

GetParentSessionId returns the ParentSessionId field if non-nil, zero value otherwise.

### GetParentSessionIdOk

`func (o *CloudRegisterReq) GetParentSessionIdOk() (*string, bool)`

GetParentSessionIdOk returns a tuple with the ParentSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSessionId

`func (o *CloudRegisterReq) SetParentSessionId(v string)`

SetParentSessionId sets ParentSessionId field to given value.

### HasParentSessionId

`func (o *CloudRegisterReq) HasParentSessionId() bool`

HasParentSessionId returns a boolean if a field has been set.

### GetProject

`func (o *CloudRegisterReq) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudRegisterReq) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudRegisterReq) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudRegisterReq) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProvider

`func (o *CloudRegisterReq) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudRegisterReq) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudRegisterReq) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudRegisterReq) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublished

`func (o *CloudRegisterReq) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *CloudRegisterReq) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *CloudRegisterReq) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *CloudRegisterReq) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetRepo

`func (o *CloudRegisterReq) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CloudRegisterReq) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CloudRegisterReq) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CloudRegisterReq) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetStatus

`func (o *CloudRegisterReq) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudRegisterReq) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudRegisterReq) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudRegisterReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTarget

`func (o *CloudRegisterReq) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CloudRegisterReq) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CloudRegisterReq) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *CloudRegisterReq) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTaskRunId

`func (o *CloudRegisterReq) GetTaskRunId() string`

GetTaskRunId returns the TaskRunId field if non-nil, zero value otherwise.

### GetTaskRunIdOk

`func (o *CloudRegisterReq) GetTaskRunIdOk() (*string, bool)`

GetTaskRunIdOk returns a tuple with the TaskRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRunId

`func (o *CloudRegisterReq) SetTaskRunId(v string)`

SetTaskRunId sets TaskRunId field to given value.

### HasTaskRunId

`func (o *CloudRegisterReq) HasTaskRunId() bool`

HasTaskRunId returns a boolean if a field has been set.

### GetTaskWorkflowId

`func (o *CloudRegisterReq) GetTaskWorkflowId() string`

GetTaskWorkflowId returns the TaskWorkflowId field if non-nil, zero value otherwise.

### GetTaskWorkflowIdOk

`func (o *CloudRegisterReq) GetTaskWorkflowIdOk() (*string, bool)`

GetTaskWorkflowIdOk returns a tuple with the TaskWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskWorkflowId

`func (o *CloudRegisterReq) SetTaskWorkflowId(v string)`

SetTaskWorkflowId sets TaskWorkflowId field to given value.

### HasTaskWorkflowId

`func (o *CloudRegisterReq) HasTaskWorkflowId() bool`

HasTaskWorkflowId returns a boolean if a field has been set.

### GetTerminal

`func (o *CloudRegisterReq) GetTerminal() string`

GetTerminal returns the Terminal field if non-nil, zero value otherwise.

### GetTerminalOk

`func (o *CloudRegisterReq) GetTerminalOk() (*string, bool)`

GetTerminalOk returns a tuple with the Terminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminal

`func (o *CloudRegisterReq) SetTerminal(v string)`

SetTerminal sets Terminal field to given value.

### HasTerminal

`func (o *CloudRegisterReq) HasTerminal() bool`

HasTerminal returns a boolean if a field has been set.

### GetTitle

`func (o *CloudRegisterReq) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudRegisterReq) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudRegisterReq) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudRegisterReq) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


