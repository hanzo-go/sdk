# AgentCodingPullIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base** | Pointer to **string** | Base is the branch to propose into. Empty is the branch the run started from. | [optional] 
**Body** | Pointer to **string** | Body is its description. Empty says who it is for, the task, the branch and the run it came from. | [optional] 
**Draft** | Pointer to **bool** | Draft opens it as a draft, which a reviewer cannot merge until it is ready. | [optional] 
**Session** | Pointer to **string** | Session is the run&#39;s handle, from the path. | [optional] 
**Title** | Pointer to **string** | Title is the pull request&#39;s title. Empty takes one from the run&#39;s task. | [optional] 

## Methods

### NewAgentCodingPullIn

`func NewAgentCodingPullIn() *AgentCodingPullIn`

NewAgentCodingPullIn instantiates a new AgentCodingPullIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCodingPullInWithDefaults

`func NewAgentCodingPullInWithDefaults() *AgentCodingPullIn`

NewAgentCodingPullInWithDefaults instantiates a new AgentCodingPullIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *AgentCodingPullIn) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *AgentCodingPullIn) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *AgentCodingPullIn) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *AgentCodingPullIn) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetBody

`func (o *AgentCodingPullIn) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *AgentCodingPullIn) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *AgentCodingPullIn) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *AgentCodingPullIn) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetDraft

`func (o *AgentCodingPullIn) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *AgentCodingPullIn) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *AgentCodingPullIn) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *AgentCodingPullIn) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetSession

`func (o *AgentCodingPullIn) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *AgentCodingPullIn) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *AgentCodingPullIn) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *AgentCodingPullIn) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetTitle

`func (o *AgentCodingPullIn) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AgentCodingPullIn) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AgentCodingPullIn) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AgentCodingPullIn) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


