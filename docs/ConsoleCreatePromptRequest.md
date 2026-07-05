# ConsoleCreatePromptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Prompt** | **interface{}** |  | 
**Type** | **string** |  | 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**CommitMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewConsoleCreatePromptRequest

`func NewConsoleCreatePromptRequest(name string, prompt interface{}, type_ string, ) *ConsoleCreatePromptRequest`

NewConsoleCreatePromptRequest instantiates a new ConsoleCreatePromptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleCreatePromptRequestWithDefaults

`func NewConsoleCreatePromptRequestWithDefaults() *ConsoleCreatePromptRequest`

NewConsoleCreatePromptRequestWithDefaults instantiates a new ConsoleCreatePromptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConsoleCreatePromptRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsoleCreatePromptRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsoleCreatePromptRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPrompt

`func (o *ConsoleCreatePromptRequest) GetPrompt() interface{}`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *ConsoleCreatePromptRequest) GetPromptOk() (*interface{}, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *ConsoleCreatePromptRequest) SetPrompt(v interface{})`

SetPrompt sets Prompt field to given value.


### SetPromptNil

`func (o *ConsoleCreatePromptRequest) SetPromptNil(b bool)`

 SetPromptNil sets the value for Prompt to be an explicit nil

### UnsetPrompt
`func (o *ConsoleCreatePromptRequest) UnsetPrompt()`

UnsetPrompt ensures that no value is present for Prompt, not even an explicit nil
### GetType

`func (o *ConsoleCreatePromptRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConsoleCreatePromptRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConsoleCreatePromptRequest) SetType(v string)`

SetType sets Type field to given value.


### GetConfig

`func (o *ConsoleCreatePromptRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ConsoleCreatePromptRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ConsoleCreatePromptRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ConsoleCreatePromptRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLabels

`func (o *ConsoleCreatePromptRequest) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ConsoleCreatePromptRequest) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ConsoleCreatePromptRequest) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ConsoleCreatePromptRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *ConsoleCreatePromptRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConsoleCreatePromptRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConsoleCreatePromptRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConsoleCreatePromptRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCommitMessage

`func (o *ConsoleCreatePromptRequest) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *ConsoleCreatePromptRequest) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *ConsoleCreatePromptRequest) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *ConsoleCreatePromptRequest) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


