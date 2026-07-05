# PromptsPromptDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Prompt** | Pointer to **string** | Current version content | [optional] 
**Version** | Pointer to **int32** | Current version number | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**VersionHistory** | Pointer to [**[]PromptsVersionMeta**](PromptsVersionMeta.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPromptsPromptDetail

`func NewPromptsPromptDetail() *PromptsPromptDetail`

NewPromptsPromptDetail instantiates a new PromptsPromptDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptsPromptDetailWithDefaults

`func NewPromptsPromptDetailWithDefaults() *PromptsPromptDetail`

NewPromptsPromptDetailWithDefaults instantiates a new PromptsPromptDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PromptsPromptDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PromptsPromptDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PromptsPromptDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PromptsPromptDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PromptsPromptDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromptsPromptDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromptsPromptDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PromptsPromptDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrompt

`func (o *PromptsPromptDetail) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *PromptsPromptDetail) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *PromptsPromptDetail) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *PromptsPromptDetail) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetVersion

`func (o *PromptsPromptDetail) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PromptsPromptDetail) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PromptsPromptDetail) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PromptsPromptDetail) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLabels

`func (o *PromptsPromptDetail) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PromptsPromptDetail) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PromptsPromptDetail) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PromptsPromptDetail) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *PromptsPromptDetail) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PromptsPromptDetail) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PromptsPromptDetail) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PromptsPromptDetail) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionHistory

`func (o *PromptsPromptDetail) GetVersionHistory() []PromptsVersionMeta`

GetVersionHistory returns the VersionHistory field if non-nil, zero value otherwise.

### GetVersionHistoryOk

`func (o *PromptsPromptDetail) GetVersionHistoryOk() (*[]PromptsVersionMeta, bool)`

GetVersionHistoryOk returns a tuple with the VersionHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionHistory

`func (o *PromptsPromptDetail) SetVersionHistory(v []PromptsVersionMeta)`

SetVersionHistory sets VersionHistory field to given value.

### HasVersionHistory

`func (o *PromptsPromptDetail) HasVersionHistory() bool`

HasVersionHistory returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PromptsPromptDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PromptsPromptDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PromptsPromptDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PromptsPromptDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *PromptsPromptDetail) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *PromptsPromptDetail) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *PromptsPromptDetail) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *PromptsPromptDetail) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


