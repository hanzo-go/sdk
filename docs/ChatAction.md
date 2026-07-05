# ChatAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionId** | Pointer to **string** |  | [optional] 
**AgentId** | Pointer to **string** |  | [optional] 
**AssistantId** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**ChatActionMetadata**](ChatActionMetadata.md) |  | [optional] 

## Methods

### NewChatAction

`func NewChatAction() *ChatAction`

NewChatAction instantiates a new ChatAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatActionWithDefaults

`func NewChatActionWithDefaults() *ChatAction`

NewChatActionWithDefaults instantiates a new ChatAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionId

`func (o *ChatAction) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *ChatAction) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *ChatAction) SetActionId(v string)`

SetActionId sets ActionId field to given value.

### HasActionId

`func (o *ChatAction) HasActionId() bool`

HasActionId returns a boolean if a field has been set.

### GetAgentId

`func (o *ChatAction) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ChatAction) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ChatAction) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ChatAction) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetAssistantId

`func (o *ChatAction) GetAssistantId() string`

GetAssistantId returns the AssistantId field if non-nil, zero value otherwise.

### GetAssistantIdOk

`func (o *ChatAction) GetAssistantIdOk() (*string, bool)`

GetAssistantIdOk returns a tuple with the AssistantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantId

`func (o *ChatAction) SetAssistantId(v string)`

SetAssistantId sets AssistantId field to given value.

### HasAssistantId

`func (o *ChatAction) HasAssistantId() bool`

HasAssistantId returns a boolean if a field has been set.

### GetUser

`func (o *ChatAction) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ChatAction) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ChatAction) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ChatAction) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetMetadata

`func (o *ChatAction) GetMetadata() ChatActionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ChatAction) GetMetadataOk() (*ChatActionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ChatAction) SetMetadata(v ChatActionMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ChatAction) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


