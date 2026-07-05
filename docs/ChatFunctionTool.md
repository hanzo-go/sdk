# ChatFunctionTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Function** | Pointer to [**ChatFunctionToolFunction**](ChatFunctionToolFunction.md) |  | [optional] 

## Methods

### NewChatFunctionTool

`func NewChatFunctionTool() *ChatFunctionTool`

NewChatFunctionTool instantiates a new ChatFunctionTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatFunctionToolWithDefaults

`func NewChatFunctionToolWithDefaults() *ChatFunctionTool`

NewChatFunctionToolWithDefaults instantiates a new ChatFunctionTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChatFunctionTool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChatFunctionTool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChatFunctionTool) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ChatFunctionTool) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFunction

`func (o *ChatFunctionTool) GetFunction() ChatFunctionToolFunction`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *ChatFunctionTool) GetFunctionOk() (*ChatFunctionToolFunction, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *ChatFunctionTool) SetFunction(v ChatFunctionToolFunction)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *ChatFunctionTool) HasFunction() bool`

HasFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


