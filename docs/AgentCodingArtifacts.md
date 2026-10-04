# AgentCodingArtifacts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | Pointer to [**[]AgentArtifact**](AgentArtifact.md) | Artifacts are the run&#39;s artifacts, stored files first. Never null. | [optional] 
**Saved** | Pointer to **string** | Saved is when the artifacts were saved, RFC 3339 in UTC; empty while the run has saved none. | [optional] 
**Session** | Pointer to **string** | Session is the run&#39;s handle. | [optional] 

## Methods

### NewAgentCodingArtifacts

`func NewAgentCodingArtifacts() *AgentCodingArtifacts`

NewAgentCodingArtifacts instantiates a new AgentCodingArtifacts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCodingArtifactsWithDefaults

`func NewAgentCodingArtifactsWithDefaults() *AgentCodingArtifacts`

NewAgentCodingArtifactsWithDefaults instantiates a new AgentCodingArtifacts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *AgentCodingArtifacts) GetArtifacts() []AgentArtifact`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *AgentCodingArtifacts) GetArtifactsOk() (*[]AgentArtifact, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *AgentCodingArtifacts) SetArtifacts(v []AgentArtifact)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *AgentCodingArtifacts) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetSaved

`func (o *AgentCodingArtifacts) GetSaved() string`

GetSaved returns the Saved field if non-nil, zero value otherwise.

### GetSavedOk

`func (o *AgentCodingArtifacts) GetSavedOk() (*string, bool)`

GetSavedOk returns a tuple with the Saved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaved

`func (o *AgentCodingArtifacts) SetSaved(v string)`

SetSaved sets Saved field to given value.

### HasSaved

`func (o *AgentCodingArtifacts) HasSaved() bool`

HasSaved returns a boolean if a field has been set.

### GetSession

`func (o *AgentCodingArtifacts) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *AgentCodingArtifacts) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *AgentCodingArtifacts) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *AgentCodingArtifacts) HasSession() bool`

HasSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


