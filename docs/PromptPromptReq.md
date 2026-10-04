# PromptPromptReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **[]string** | Labels is free-form taxonomy, each up to 64 characters, capped at 32 entries. | [optional] 
**Name** | Pointer to **string** | Name is the org-unique handle AND the URL segment the prompt is addressed by: 1-64 characters matching ^[A-Za-z0-9][A-Za-z0-9._-]*$. \&quot;metrics\&quot;, \&quot;new\&quot; and \&quot;catalog\&quot; are reserved. A name that already exists appends a new version. | [optional] 
**Prompt** | Pointer to **string** | Prompt is the template body, capped at 64 KiB. It holds template text only — never a secret. | [optional] 
**Tags** | Pointer to **[]string** | Tags is free-form taxonomy under the same bounds as Labels. | [optional] 
**Type** | Pointer to **string** | Type labels the template&#39;s kind; defaults to \&quot;text\&quot;. | [optional] 

## Methods

### NewPromptPromptReq

`func NewPromptPromptReq() *PromptPromptReq`

NewPromptPromptReq instantiates a new PromptPromptReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptPromptReqWithDefaults

`func NewPromptPromptReqWithDefaults() *PromptPromptReq`

NewPromptPromptReqWithDefaults instantiates a new PromptPromptReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *PromptPromptReq) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PromptPromptReq) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PromptPromptReq) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PromptPromptReq) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *PromptPromptReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PromptPromptReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PromptPromptReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PromptPromptReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrompt

`func (o *PromptPromptReq) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *PromptPromptReq) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *PromptPromptReq) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *PromptPromptReq) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetTags

`func (o *PromptPromptReq) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PromptPromptReq) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PromptPromptReq) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PromptPromptReq) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *PromptPromptReq) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromptPromptReq) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromptPromptReq) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PromptPromptReq) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


