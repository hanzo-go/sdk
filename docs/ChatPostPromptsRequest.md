# ChatPostPromptsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prompt** | **map[string]interface{}** |  | 
**Group** | [**ChatPostPromptsRequestGroup**](ChatPostPromptsRequestGroup.md) |  | 

## Methods

### NewChatPostPromptsRequest

`func NewChatPostPromptsRequest(prompt map[string]interface{}, group ChatPostPromptsRequestGroup, ) *ChatPostPromptsRequest`

NewChatPostPromptsRequest instantiates a new ChatPostPromptsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatPostPromptsRequestWithDefaults

`func NewChatPostPromptsRequestWithDefaults() *ChatPostPromptsRequest`

NewChatPostPromptsRequestWithDefaults instantiates a new ChatPostPromptsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrompt

`func (o *ChatPostPromptsRequest) GetPrompt() map[string]interface{}`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *ChatPostPromptsRequest) GetPromptOk() (*map[string]interface{}, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *ChatPostPromptsRequest) SetPrompt(v map[string]interface{})`

SetPrompt sets Prompt field to given value.


### GetGroup

`func (o *ChatPostPromptsRequest) GetGroup() ChatPostPromptsRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ChatPostPromptsRequest) GetGroupOk() (*ChatPostPromptsRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ChatPostPromptsRequest) SetGroup(v ChatPostPromptsRequestGroup)`

SetGroup sets Group field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


