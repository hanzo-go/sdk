# ChatDeleteFilesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | [**[]ChatDeleteFilesRequestFilesInner**](ChatDeleteFilesRequestFilesInner.md) |  | 
**AgentId** | Pointer to **string** |  | [optional] 
**ToolResource** | Pointer to **string** |  | [optional] 
**AssistantId** | Pointer to **string** |  | [optional] 

## Methods

### NewChatDeleteFilesRequest

`func NewChatDeleteFilesRequest(files []ChatDeleteFilesRequestFilesInner, ) *ChatDeleteFilesRequest`

NewChatDeleteFilesRequest instantiates a new ChatDeleteFilesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatDeleteFilesRequestWithDefaults

`func NewChatDeleteFilesRequestWithDefaults() *ChatDeleteFilesRequest`

NewChatDeleteFilesRequestWithDefaults instantiates a new ChatDeleteFilesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *ChatDeleteFilesRequest) GetFiles() []ChatDeleteFilesRequestFilesInner`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ChatDeleteFilesRequest) GetFilesOk() (*[]ChatDeleteFilesRequestFilesInner, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ChatDeleteFilesRequest) SetFiles(v []ChatDeleteFilesRequestFilesInner)`

SetFiles sets Files field to given value.


### GetAgentId

`func (o *ChatDeleteFilesRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ChatDeleteFilesRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ChatDeleteFilesRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ChatDeleteFilesRequest) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetToolResource

`func (o *ChatDeleteFilesRequest) GetToolResource() string`

GetToolResource returns the ToolResource field if non-nil, zero value otherwise.

### GetToolResourceOk

`func (o *ChatDeleteFilesRequest) GetToolResourceOk() (*string, bool)`

GetToolResourceOk returns a tuple with the ToolResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolResource

`func (o *ChatDeleteFilesRequest) SetToolResource(v string)`

SetToolResource sets ToolResource field to given value.

### HasToolResource

`func (o *ChatDeleteFilesRequest) HasToolResource() bool`

HasToolResource returns a boolean if a field has been set.

### GetAssistantId

`func (o *ChatDeleteFilesRequest) GetAssistantId() string`

GetAssistantId returns the AssistantId field if non-nil, zero value otherwise.

### GetAssistantIdOk

`func (o *ChatDeleteFilesRequest) GetAssistantIdOk() (*string, bool)`

GetAssistantIdOk returns a tuple with the AssistantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantId

`func (o *ChatDeleteFilesRequest) SetAssistantId(v string)`

SetAssistantId sets AssistantId field to given value.

### HasAssistantId

`func (o *ChatDeleteFilesRequest) HasAssistantId() bool`

HasAssistantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


