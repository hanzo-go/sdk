# OpenaiChatCompletionMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MultiContent** | Pointer to [**[]OpenaiChatMessagePart**](OpenaiChatMessagePart.md) |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**FunctionCall** | Pointer to [**OpenaiFunctionCall**](OpenaiFunctionCall.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ReasoningContent** | Pointer to **string** |  | [optional] 
**Refusal** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**ToolCallId** | Pointer to **string** |  | [optional] 
**ToolCalls** | Pointer to [**[]OpenaiToolCall**](OpenaiToolCall.md) |  | [optional] 

## Methods

### NewOpenaiChatCompletionMessage

`func NewOpenaiChatCompletionMessage() *OpenaiChatCompletionMessage`

NewOpenaiChatCompletionMessage instantiates a new OpenaiChatCompletionMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenaiChatCompletionMessageWithDefaults

`func NewOpenaiChatCompletionMessageWithDefaults() *OpenaiChatCompletionMessage`

NewOpenaiChatCompletionMessageWithDefaults instantiates a new OpenaiChatCompletionMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMultiContent

`func (o *OpenaiChatCompletionMessage) GetMultiContent() []OpenaiChatMessagePart`

GetMultiContent returns the MultiContent field if non-nil, zero value otherwise.

### GetMultiContentOk

`func (o *OpenaiChatCompletionMessage) GetMultiContentOk() (*[]OpenaiChatMessagePart, bool)`

GetMultiContentOk returns a tuple with the MultiContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiContent

`func (o *OpenaiChatCompletionMessage) SetMultiContent(v []OpenaiChatMessagePart)`

SetMultiContent sets MultiContent field to given value.

### HasMultiContent

`func (o *OpenaiChatCompletionMessage) HasMultiContent() bool`

HasMultiContent returns a boolean if a field has been set.

### GetContent

`func (o *OpenaiChatCompletionMessage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OpenaiChatCompletionMessage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OpenaiChatCompletionMessage) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *OpenaiChatCompletionMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetFunctionCall

`func (o *OpenaiChatCompletionMessage) GetFunctionCall() OpenaiFunctionCall`

GetFunctionCall returns the FunctionCall field if non-nil, zero value otherwise.

### GetFunctionCallOk

`func (o *OpenaiChatCompletionMessage) GetFunctionCallOk() (*OpenaiFunctionCall, bool)`

GetFunctionCallOk returns a tuple with the FunctionCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionCall

`func (o *OpenaiChatCompletionMessage) SetFunctionCall(v OpenaiFunctionCall)`

SetFunctionCall sets FunctionCall field to given value.

### HasFunctionCall

`func (o *OpenaiChatCompletionMessage) HasFunctionCall() bool`

HasFunctionCall returns a boolean if a field has been set.

### GetName

`func (o *OpenaiChatCompletionMessage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenaiChatCompletionMessage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenaiChatCompletionMessage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenaiChatCompletionMessage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReasoningContent

`func (o *OpenaiChatCompletionMessage) GetReasoningContent() string`

GetReasoningContent returns the ReasoningContent field if non-nil, zero value otherwise.

### GetReasoningContentOk

`func (o *OpenaiChatCompletionMessage) GetReasoningContentOk() (*string, bool)`

GetReasoningContentOk returns a tuple with the ReasoningContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoningContent

`func (o *OpenaiChatCompletionMessage) SetReasoningContent(v string)`

SetReasoningContent sets ReasoningContent field to given value.

### HasReasoningContent

`func (o *OpenaiChatCompletionMessage) HasReasoningContent() bool`

HasReasoningContent returns a boolean if a field has been set.

### GetRefusal

`func (o *OpenaiChatCompletionMessage) GetRefusal() string`

GetRefusal returns the Refusal field if non-nil, zero value otherwise.

### GetRefusalOk

`func (o *OpenaiChatCompletionMessage) GetRefusalOk() (*string, bool)`

GetRefusalOk returns a tuple with the Refusal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefusal

`func (o *OpenaiChatCompletionMessage) SetRefusal(v string)`

SetRefusal sets Refusal field to given value.

### HasRefusal

`func (o *OpenaiChatCompletionMessage) HasRefusal() bool`

HasRefusal returns a boolean if a field has been set.

### GetRole

`func (o *OpenaiChatCompletionMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OpenaiChatCompletionMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OpenaiChatCompletionMessage) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OpenaiChatCompletionMessage) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetToolCallId

`func (o *OpenaiChatCompletionMessage) GetToolCallId() string`

GetToolCallId returns the ToolCallId field if non-nil, zero value otherwise.

### GetToolCallIdOk

`func (o *OpenaiChatCompletionMessage) GetToolCallIdOk() (*string, bool)`

GetToolCallIdOk returns a tuple with the ToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallId

`func (o *OpenaiChatCompletionMessage) SetToolCallId(v string)`

SetToolCallId sets ToolCallId field to given value.

### HasToolCallId

`func (o *OpenaiChatCompletionMessage) HasToolCallId() bool`

HasToolCallId returns a boolean if a field has been set.

### GetToolCalls

`func (o *OpenaiChatCompletionMessage) GetToolCalls() []OpenaiToolCall`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *OpenaiChatCompletionMessage) GetToolCallsOk() (*[]OpenaiToolCall, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *OpenaiChatCompletionMessage) SetToolCalls(v []OpenaiToolCall)`

SetToolCalls sets ToolCalls field to given value.

### HasToolCalls

`func (o *OpenaiChatCompletionMessage) HasToolCalls() bool`

HasToolCalls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


