# ChatPostConvosFork200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conversation** | Pointer to [**ChatConversation**](ChatConversation.md) |  | [optional] 
**Messages** | Pointer to [**[]ChatMessage**](ChatMessage.md) |  | [optional] 

## Methods

### NewChatPostConvosFork200Response

`func NewChatPostConvosFork200Response() *ChatPostConvosFork200Response`

NewChatPostConvosFork200Response instantiates a new ChatPostConvosFork200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatPostConvosFork200ResponseWithDefaults

`func NewChatPostConvosFork200ResponseWithDefaults() *ChatPostConvosFork200Response`

NewChatPostConvosFork200ResponseWithDefaults instantiates a new ChatPostConvosFork200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversation

`func (o *ChatPostConvosFork200Response) GetConversation() ChatConversation`

GetConversation returns the Conversation field if non-nil, zero value otherwise.

### GetConversationOk

`func (o *ChatPostConvosFork200Response) GetConversationOk() (*ChatConversation, bool)`

GetConversationOk returns a tuple with the Conversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversation

`func (o *ChatPostConvosFork200Response) SetConversation(v ChatConversation)`

SetConversation sets Conversation field to given value.

### HasConversation

`func (o *ChatPostConvosFork200Response) HasConversation() bool`

HasConversation returns a boolean if a field has been set.

### GetMessages

`func (o *ChatPostConvosFork200Response) GetMessages() []ChatMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ChatPostConvosFork200Response) GetMessagesOk() (*[]ChatMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ChatPostConvosFork200Response) SetMessages(v []ChatMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ChatPostConvosFork200Response) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


