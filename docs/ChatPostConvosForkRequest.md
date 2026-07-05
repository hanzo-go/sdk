# ChatPostConvosForkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationId** | **string** |  | 
**MessageId** | **string** |  | 
**Option** | Pointer to **string** |  | [optional] 
**SplitAtTarget** | Pointer to **bool** |  | [optional] 
**LatestMessageId** | Pointer to **string** |  | [optional] 

## Methods

### NewChatPostConvosForkRequest

`func NewChatPostConvosForkRequest(conversationId string, messageId string, ) *ChatPostConvosForkRequest`

NewChatPostConvosForkRequest instantiates a new ChatPostConvosForkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatPostConvosForkRequestWithDefaults

`func NewChatPostConvosForkRequestWithDefaults() *ChatPostConvosForkRequest`

NewChatPostConvosForkRequestWithDefaults instantiates a new ChatPostConvosForkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversationId

`func (o *ChatPostConvosForkRequest) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *ChatPostConvosForkRequest) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *ChatPostConvosForkRequest) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetMessageId

`func (o *ChatPostConvosForkRequest) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ChatPostConvosForkRequest) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ChatPostConvosForkRequest) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetOption

`func (o *ChatPostConvosForkRequest) GetOption() string`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *ChatPostConvosForkRequest) GetOptionOk() (*string, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *ChatPostConvosForkRequest) SetOption(v string)`

SetOption sets Option field to given value.

### HasOption

`func (o *ChatPostConvosForkRequest) HasOption() bool`

HasOption returns a boolean if a field has been set.

### GetSplitAtTarget

`func (o *ChatPostConvosForkRequest) GetSplitAtTarget() bool`

GetSplitAtTarget returns the SplitAtTarget field if non-nil, zero value otherwise.

### GetSplitAtTargetOk

`func (o *ChatPostConvosForkRequest) GetSplitAtTargetOk() (*bool, bool)`

GetSplitAtTargetOk returns a tuple with the SplitAtTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitAtTarget

`func (o *ChatPostConvosForkRequest) SetSplitAtTarget(v bool)`

SetSplitAtTarget sets SplitAtTarget field to given value.

### HasSplitAtTarget

`func (o *ChatPostConvosForkRequest) HasSplitAtTarget() bool`

HasSplitAtTarget returns a boolean if a field has been set.

### GetLatestMessageId

`func (o *ChatPostConvosForkRequest) GetLatestMessageId() string`

GetLatestMessageId returns the LatestMessageId field if non-nil, zero value otherwise.

### GetLatestMessageIdOk

`func (o *ChatPostConvosForkRequest) GetLatestMessageIdOk() (*string, bool)`

GetLatestMessageIdOk returns a tuple with the LatestMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestMessageId

`func (o *ChatPostConvosForkRequest) SetLatestMessageId(v string)`

SetLatestMessageId sets LatestMessageId field to given value.

### HasLatestMessageId

`func (o *ChatPostConvosForkRequest) HasLatestMessageId() bool`

HasLatestMessageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


