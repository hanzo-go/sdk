# ChatConversationListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conversations** | Pointer to [**[]ChatConversation**](ChatConversation.md) |  | [optional] 
**NextCursor** | Pointer to **string** |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 

## Methods

### NewChatConversationListResponse

`func NewChatConversationListResponse() *ChatConversationListResponse`

NewChatConversationListResponse instantiates a new ChatConversationListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatConversationListResponseWithDefaults

`func NewChatConversationListResponseWithDefaults() *ChatConversationListResponse`

NewChatConversationListResponseWithDefaults instantiates a new ChatConversationListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversations

`func (o *ChatConversationListResponse) GetConversations() []ChatConversation`

GetConversations returns the Conversations field if non-nil, zero value otherwise.

### GetConversationsOk

`func (o *ChatConversationListResponse) GetConversationsOk() (*[]ChatConversation, bool)`

GetConversationsOk returns a tuple with the Conversations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversations

`func (o *ChatConversationListResponse) SetConversations(v []ChatConversation)`

SetConversations sets Conversations field to given value.

### HasConversations

`func (o *ChatConversationListResponse) HasConversations() bool`

HasConversations returns a boolean if a field has been set.

### GetNextCursor

`func (o *ChatConversationListResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ChatConversationListResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ChatConversationListResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ChatConversationListResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetHasMore

`func (o *ChatConversationListResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ChatConversationListResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ChatConversationListResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *ChatConversationListResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


