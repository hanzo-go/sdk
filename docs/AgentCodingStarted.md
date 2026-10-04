# AgentCodingStarted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | Branch is the ref the run will push its work to, and the ONLY ref it is permitted to write. It exists before the work does, so it is safe to tell somebody where to look while the run is still going. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository the run was admitted against, echoed back as the engine resolved it. | [optional] 
**Routed** | Pointer to **bool** | Routed says the run went to one of the org&#39;s own registered machines rather than to a sandbox in our cluster. False is the ordinary case. | [optional] 
**SessionId** | Pointer to **string** | SessionID is the run&#39;s handle: its durable record, and the root its live progress streams under at /v1/agent/sessions/stream?root&#x3D;&lt;sessionId&gt;. Every later question about this run is asked with it. | [optional] 
**TargetId** | Pointer to **string** | TargetID names that machine when Routed is true, and is empty otherwise. | [optional] 

## Methods

### NewAgentCodingStarted

`func NewAgentCodingStarted() *AgentCodingStarted`

NewAgentCodingStarted instantiates a new AgentCodingStarted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCodingStartedWithDefaults

`func NewAgentCodingStartedWithDefaults() *AgentCodingStarted`

NewAgentCodingStartedWithDefaults instantiates a new AgentCodingStarted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *AgentCodingStarted) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *AgentCodingStarted) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *AgentCodingStarted) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *AgentCodingStarted) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetRepo

`func (o *AgentCodingStarted) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *AgentCodingStarted) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *AgentCodingStarted) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *AgentCodingStarted) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRouted

`func (o *AgentCodingStarted) GetRouted() bool`

GetRouted returns the Routed field if non-nil, zero value otherwise.

### GetRoutedOk

`func (o *AgentCodingStarted) GetRoutedOk() (*bool, bool)`

GetRoutedOk returns a tuple with the Routed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouted

`func (o *AgentCodingStarted) SetRouted(v bool)`

SetRouted sets Routed field to given value.

### HasRouted

`func (o *AgentCodingStarted) HasRouted() bool`

HasRouted returns a boolean if a field has been set.

### GetSessionId

`func (o *AgentCodingStarted) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *AgentCodingStarted) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *AgentCodingStarted) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *AgentCodingStarted) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetTargetId

`func (o *AgentCodingStarted) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *AgentCodingStarted) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *AgentCodingStarted) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *AgentCodingStarted) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


