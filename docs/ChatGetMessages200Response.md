# ChatGetMessages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]ChatMessage**](ChatMessage.md) |  | [optional] 
**NextCursor** | Pointer to **string** |  | [optional] 

## Methods

### NewChatGetMessages200Response

`func NewChatGetMessages200Response() *ChatGetMessages200Response`

NewChatGetMessages200Response instantiates a new ChatGetMessages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatGetMessages200ResponseWithDefaults

`func NewChatGetMessages200ResponseWithDefaults() *ChatGetMessages200Response`

NewChatGetMessages200ResponseWithDefaults instantiates a new ChatGetMessages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *ChatGetMessages200Response) GetMessages() []ChatMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ChatGetMessages200Response) GetMessagesOk() (*[]ChatMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ChatGetMessages200Response) SetMessages(v []ChatMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ChatGetMessages200Response) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetNextCursor

`func (o *ChatGetMessages200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ChatGetMessages200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ChatGetMessages200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ChatGetMessages200Response) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


