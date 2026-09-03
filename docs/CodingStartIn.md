# CodingStartIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to **string** |  | [optional] 
**AgentRef** | Pointer to **string** |  | [optional] 
**Base** | Pointer to **string** |  | [optional] 
**Desktop** | Pointer to **bool** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**Prompt** | Pointer to **string** |  | [optional] 
**ReplyChannel** | Pointer to **string** |  | [optional] 
**ReplyThread** | Pointer to **string** |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**TimeoutSeconds** | Pointer to **int64** |  | [optional] 
**Tool** | Pointer to **string** |  | [optional] 

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

`func (o *CodingStartIn) GetTimeoutSeconds() int64`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *CodingStartIn) GetTimeoutSecondsOk() (*int64, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *CodingStartIn) SetTimeoutSeconds(v int64)`

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


