# ChatRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message is the founder&#39;s question for the Business AI. Required; trimmed, and clipped to 4 KiB so a caller cannot amplify the AI prompt. | [optional] 

## Methods

### NewChatRequest

`func NewChatRequest() *ChatRequest`

NewChatRequest instantiates a new ChatRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatRequestWithDefaults

`func NewChatRequestWithDefaults() *ChatRequest`

NewChatRequestWithDefaults instantiates a new ChatRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ChatRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ChatRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ChatRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ChatRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


