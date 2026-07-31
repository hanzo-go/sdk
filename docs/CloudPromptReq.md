# CloudPromptReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **[]string** | Labels is free-form taxonomy, each up to 64 characters, capped at 32 entries. | [optional] 
**Name** | Pointer to **string** | Name is the org-unique handle AND the URL segment the prompt is addressed by: 1-64 characters matching ^[A-Za-z0-9][A-Za-z0-9._-]*$. \&quot;metrics\&quot;, \&quot;new\&quot; and \&quot;catalog\&quot; are reserved. A name that already exists appends a new version. | [optional] 
**Prompt** | Pointer to **string** | Prompt is the template body, capped at 64 KiB. It holds template text only — never a secret. | [optional] 
**Tags** | Pointer to **[]string** | Tags is free-form taxonomy under the same bounds as Labels. | [optional] 
**Type** | Pointer to **string** | Type labels the template&#39;s kind; defaults to \&quot;text\&quot;. | [optional] 

## Methods

### NewCloudPromptReq

`func NewCloudPromptReq() *CloudPromptReq`

NewCloudPromptReq instantiates a new CloudPromptReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPromptReqWithDefaults

`func NewCloudPromptReqWithDefaults() *CloudPromptReq`

NewCloudPromptReqWithDefaults instantiates a new CloudPromptReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *CloudPromptReq) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CloudPromptReq) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CloudPromptReq) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CloudPromptReq) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *CloudPromptReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudPromptReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudPromptReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudPromptReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrompt

`func (o *CloudPromptReq) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *CloudPromptReq) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *CloudPromptReq) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *CloudPromptReq) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetTags

`func (o *CloudPromptReq) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CloudPromptReq) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CloudPromptReq) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CloudPromptReq) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *CloudPromptReq) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudPromptReq) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudPromptReq) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudPromptReq) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


