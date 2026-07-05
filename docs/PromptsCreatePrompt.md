# PromptsCreatePrompt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Org-unique handle, ^[A-Za-z0-9][A-Za-z0-9._-]{0,63}$ (not \&quot;metrics\&quot;/\&quot;new\&quot;/\&quot;catalog\&quot;) | 
**Type** | Pointer to **string** |  | [optional] [default to "text"]
**Prompt** | Pointer to **string** | Template body (max 64KiB) | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPromptsCreatePrompt

`func NewPromptsCreatePrompt(name string, ) *PromptsCreatePrompt`

NewPromptsCreatePrompt instantiates a new PromptsCreatePrompt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptsCreatePromptWithDefaults

`func NewPromptsCreatePromptWithDefaults() *PromptsCreatePrompt`

NewPromptsCreatePromptWithDefaults instantiates a new PromptsCreatePrompt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PromptsCreatePrompt) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PromptsCreatePrompt) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PromptsCreatePrompt) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *PromptsCreatePrompt) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromptsCreatePrompt) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromptsCreatePrompt) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PromptsCreatePrompt) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrompt

`func (o *PromptsCreatePrompt) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *PromptsCreatePrompt) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *PromptsCreatePrompt) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *PromptsCreatePrompt) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetLabels

`func (o *PromptsCreatePrompt) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PromptsCreatePrompt) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PromptsCreatePrompt) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PromptsCreatePrompt) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *PromptsCreatePrompt) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PromptsCreatePrompt) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PromptsCreatePrompt) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PromptsCreatePrompt) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


