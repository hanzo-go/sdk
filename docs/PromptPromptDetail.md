# PromptPromptDetail

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
**VersionHistory** | Pointer to [**[]PromptVersionView**](PromptVersionView.md) | Versions is the history METADATA, newest first, capped at the last 100 — no bodies, so a long history cannot inflate this response. It always includes the current version as its first entry. | [optional] 

## Methods

### NewPromptPromptDetail

`func NewPromptPromptDetail() *PromptPromptDetail`

NewPromptPromptDetail instantiates a new PromptPromptDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptPromptDetailWithDefaults

`func NewPromptPromptDetailWithDefaults() *PromptPromptDetail`

NewPromptPromptDetailWithDefaults instantiates a new PromptPromptDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PromptPromptDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PromptPromptDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PromptPromptDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PromptPromptDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLabels

`func (o *PromptPromptDetail) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PromptPromptDetail) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PromptPromptDetail) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PromptPromptDetail) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *PromptPromptDetail) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *PromptPromptDetail) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *PromptPromptDetail) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *PromptPromptDetail) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *PromptPromptDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PromptPromptDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PromptPromptDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PromptPromptDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrompt

`func (o *PromptPromptDetail) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *PromptPromptDetail) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *PromptPromptDetail) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *PromptPromptDetail) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetTags

`func (o *PromptPromptDetail) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PromptPromptDetail) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PromptPromptDetail) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PromptPromptDetail) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *PromptPromptDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromptPromptDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromptPromptDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PromptPromptDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *PromptPromptDetail) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PromptPromptDetail) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PromptPromptDetail) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PromptPromptDetail) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionHistory

`func (o *PromptPromptDetail) GetVersionHistory() []PromptVersionView`

GetVersionHistory returns the VersionHistory field if non-nil, zero value otherwise.

### GetVersionHistoryOk

`func (o *PromptPromptDetail) GetVersionHistoryOk() (*[]PromptVersionView, bool)`

GetVersionHistoryOk returns a tuple with the VersionHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionHistory

`func (o *PromptPromptDetail) SetVersionHistory(v []PromptVersionView)`

SetVersionHistory sets VersionHistory field to given value.

### HasVersionHistory

`func (o *PromptPromptDetail) HasVersionHistory() bool`

HasVersionHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


