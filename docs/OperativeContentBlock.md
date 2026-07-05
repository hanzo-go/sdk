# OperativeContentBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Thinking** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Tool use ID | [optional] 
**Name** | Pointer to **string** | Tool name (computer, bash, str_replace_editor) | [optional] 
**Input** | Pointer to **map[string]interface{}** | Tool input parameters | [optional] 
**ToolUseId** | Pointer to **string** | Reference to the tool_use block | [optional] 
**Content** | Pointer to **string** | Tool result content | [optional] 
**IsError** | Pointer to **bool** |  | [optional] 
**Base64Image** | Pointer to **string** | Base64-encoded PNG screenshot | [optional] 

## Methods

### NewOperativeContentBlock

`func NewOperativeContentBlock() *OperativeContentBlock`

NewOperativeContentBlock instantiates a new OperativeContentBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperativeContentBlockWithDefaults

`func NewOperativeContentBlockWithDefaults() *OperativeContentBlock`

NewOperativeContentBlockWithDefaults instantiates a new OperativeContentBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OperativeContentBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OperativeContentBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OperativeContentBlock) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OperativeContentBlock) HasType() bool`

HasType returns a boolean if a field has been set.

### GetText

`func (o *OperativeContentBlock) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *OperativeContentBlock) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *OperativeContentBlock) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *OperativeContentBlock) HasText() bool`

HasText returns a boolean if a field has been set.

### GetThinking

`func (o *OperativeContentBlock) GetThinking() string`

GetThinking returns the Thinking field if non-nil, zero value otherwise.

### GetThinkingOk

`func (o *OperativeContentBlock) GetThinkingOk() (*string, bool)`

GetThinkingOk returns a tuple with the Thinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinking

`func (o *OperativeContentBlock) SetThinking(v string)`

SetThinking sets Thinking field to given value.

### HasThinking

`func (o *OperativeContentBlock) HasThinking() bool`

HasThinking returns a boolean if a field has been set.

### GetId

`func (o *OperativeContentBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OperativeContentBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OperativeContentBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OperativeContentBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OperativeContentBlock) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OperativeContentBlock) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OperativeContentBlock) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OperativeContentBlock) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInput

`func (o *OperativeContentBlock) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *OperativeContentBlock) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *OperativeContentBlock) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *OperativeContentBlock) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetToolUseId

`func (o *OperativeContentBlock) GetToolUseId() string`

GetToolUseId returns the ToolUseId field if non-nil, zero value otherwise.

### GetToolUseIdOk

`func (o *OperativeContentBlock) GetToolUseIdOk() (*string, bool)`

GetToolUseIdOk returns a tuple with the ToolUseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolUseId

`func (o *OperativeContentBlock) SetToolUseId(v string)`

SetToolUseId sets ToolUseId field to given value.

### HasToolUseId

`func (o *OperativeContentBlock) HasToolUseId() bool`

HasToolUseId returns a boolean if a field has been set.

### GetContent

`func (o *OperativeContentBlock) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OperativeContentBlock) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OperativeContentBlock) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *OperativeContentBlock) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetIsError

`func (o *OperativeContentBlock) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *OperativeContentBlock) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *OperativeContentBlock) SetIsError(v bool)`

SetIsError sets IsError field to given value.

### HasIsError

`func (o *OperativeContentBlock) HasIsError() bool`

HasIsError returns a boolean if a field has been set.

### GetBase64Image

`func (o *OperativeContentBlock) GetBase64Image() string`

GetBase64Image returns the Base64Image field if non-nil, zero value otherwise.

### GetBase64ImageOk

`func (o *OperativeContentBlock) GetBase64ImageOk() (*string, bool)`

GetBase64ImageOk returns a tuple with the Base64Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64Image

`func (o *OperativeContentBlock) SetBase64Image(v string)`

SetBase64Image sets Base64Image field to given value.

### HasBase64Image

`func (o *OperativeContentBlock) HasBase64Image() bool`

HasBase64Image returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


