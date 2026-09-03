# PromptDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | CreatedAt is when version 1 was written, RFC 3339 UTC. Appending a version does not move it. | [optional] 
**Labels** | Pointer to **[]string** | Labels is the current version&#39;s free-form taxonomy. &#x60;[]&#x60; when none, never null. | [optional] 
**LastUpdatedAt** | Pointer to **string** | UpdatedAt is when the current version was appended, RFC 3339 UTC. Equal to createdAt for a prompt that has only ever had one version. | [optional] 
**Name** | Pointer to **string** | Name is the prompt&#39;s org-unique handle and the URL segment it is addressed by. | [optional] 
**Prompt** | Pointer to **string** | Prompt is the CURRENT version&#39;s template body — the only content this service returns. Earlier versions are listed in versionHistory by number and date, and their bodies are not served in bulk. | [optional] 
**Tags** | Pointer to **[]string** | Tags is the second free-form taxonomy, same rules as Labels. | [optional] 
**Type** | Pointer to **string** | Type labels the current version&#39;s kind; \&quot;text\&quot; unless the creator said otherwise. | [optional] 
**Version** | Pointer to **int64** | Version is the current version number, starting at 1 and incremented by one on every create against an existing name. | [optional] 
**VersionHistory** | Pointer to [**[]VersionView**](VersionView.md) | Versions is the history METADATA, newest first, capped at the last 100 — no bodies, so a long history cannot inflate this response. It always includes the current version as its first entry. | [optional] 

## Methods

### NewPromptDetail

`func NewPromptDetail() *PromptDetail`

NewPromptDetail instantiates a new PromptDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptDetailWithDefaults

`func NewPromptDetailWithDefaults() *PromptDetail`

NewPromptDetailWithDefaults instantiates a new PromptDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PromptDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PromptDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PromptDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PromptDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLabels

`func (o *PromptDetail) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PromptDetail) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PromptDetail) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PromptDetail) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *PromptDetail) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *PromptDetail) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *PromptDetail) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *PromptDetail) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *PromptDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PromptDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PromptDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PromptDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrompt

`func (o *PromptDetail) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *PromptDetail) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *PromptDetail) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *PromptDetail) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetTags

`func (o *PromptDetail) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PromptDetail) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PromptDetail) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PromptDetail) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *PromptDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromptDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromptDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PromptDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *PromptDetail) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PromptDetail) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PromptDetail) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PromptDetail) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionHistory

`func (o *PromptDetail) GetVersionHistory() []VersionView`

GetVersionHistory returns the VersionHistory field if non-nil, zero value otherwise.

### GetVersionHistoryOk

`func (o *PromptDetail) GetVersionHistoryOk() (*[]VersionView, bool)`

GetVersionHistoryOk returns a tuple with the VersionHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionHistory

`func (o *PromptDetail) SetVersionHistory(v []VersionView)`

SetVersionHistory sets VersionHistory field to given value.

### HasVersionHistory

`func (o *PromptDetail) HasVersionHistory() bool`

HasVersionHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


