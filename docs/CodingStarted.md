# CodingStarted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | Branch is the ref the run will push its work to, and the ONLY ref it is permitted to write. It exists before the work does, so it is safe to tell somebody where to look while the run is still going. | [optional] 
**Repo** | Pointer to **string** | Repo is the repository the run was admitted against, echoed back as the engine resolved it. | [optional] 
**Routed** | Pointer to **bool** | Routed says the run went to one of the org&#39;s own registered machines rather than to a sandbox in our cluster. False is the ordinary case. | [optional] 
**SessionId** | Pointer to **string** | SessionID is the run&#39;s handle: its durable record, and the id its live progress streams under at /v1/agents/sessions/{sessionId}/stream. Every later question about this run is asked with it. | [optional] 
**TargetId** | Pointer to **string** | TargetID names that machine when Routed is true, and is empty otherwise. | [optional] 

## Methods

### NewCodingStarted

`func NewCodingStarted() *CodingStarted`

NewCodingStarted instantiates a new CodingStarted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodingStartedWithDefaults

`func NewCodingStartedWithDefaults() *CodingStarted`

NewCodingStartedWithDefaults instantiates a new CodingStarted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *CodingStarted) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *CodingStarted) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *CodingStarted) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *CodingStarted) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetRepo

`func (o *CodingStarted) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *CodingStarted) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *CodingStarted) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *CodingStarted) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRouted

`func (o *CodingStarted) GetRouted() bool`

GetRouted returns the Routed field if non-nil, zero value otherwise.

### GetRoutedOk

`func (o *CodingStarted) GetRoutedOk() (*bool, bool)`

GetRoutedOk returns a tuple with the Routed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouted

`func (o *CodingStarted) SetRouted(v bool)`

SetRouted sets Routed field to given value.

### HasRouted

`func (o *CodingStarted) HasRouted() bool`

HasRouted returns a boolean if a field has been set.

### GetSessionId

`func (o *CodingStarted) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *CodingStarted) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *CodingStarted) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *CodingStarted) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetTargetId

`func (o *CodingStarted) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *CodingStarted) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *CodingStarted) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *CodingStarted) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


