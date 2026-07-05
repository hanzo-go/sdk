# PromptsPromptMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Versions** | Pointer to **[]int32** | Version numbers, newest first | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPromptsPromptMeta

`func NewPromptsPromptMeta() *PromptsPromptMeta`

NewPromptsPromptMeta instantiates a new PromptsPromptMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptsPromptMetaWithDefaults

`func NewPromptsPromptMetaWithDefaults() *PromptsPromptMeta`

NewPromptsPromptMetaWithDefaults instantiates a new PromptsPromptMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PromptsPromptMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PromptsPromptMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PromptsPromptMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PromptsPromptMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PromptsPromptMeta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromptsPromptMeta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromptsPromptMeta) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PromptsPromptMeta) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersions

`func (o *PromptsPromptMeta) GetVersions() []int32`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *PromptsPromptMeta) GetVersionsOk() (*[]int32, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *PromptsPromptMeta) SetVersions(v []int32)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *PromptsPromptMeta) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetLabels

`func (o *PromptsPromptMeta) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PromptsPromptMeta) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PromptsPromptMeta) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PromptsPromptMeta) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *PromptsPromptMeta) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PromptsPromptMeta) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PromptsPromptMeta) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PromptsPromptMeta) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *PromptsPromptMeta) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *PromptsPromptMeta) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *PromptsPromptMeta) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *PromptsPromptMeta) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


