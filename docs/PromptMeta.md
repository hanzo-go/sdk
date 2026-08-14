# PromptMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **[]string** | Labels is the creator&#39;s free-form taxonomy, stored as given after trimming and de-duplication. Always present, &#x60;[]&#x60; when none — never null. | [optional] 
**LastUpdatedAt** | Pointer to **string** | LastUpdatedAt is when the newest version was appended, RFC 3339 UTC. Empty only if the record carries no timestamp at all. | [optional] 
**Name** | Pointer to **string** | Name is the prompt&#39;s org-unique handle and the URL segment it is fetched by: GET /v1/prompts/&lt;name&gt;. | [optional] 
**Tags** | Pointer to **[]string** | Tags is the second free-form taxonomy under the same rules as Labels. Nothing in this service interprets either; they are yours to organize by. | [optional] 
**Type** | Pointer to **string** | Type labels the template&#39;s kind, \&quot;text\&quot; unless the creator said otherwise. It is the CURRENT version&#39;s type; earlier versions may carry a different one. | [optional] 
**Versions** | Pointer to **[]int32** | Versions lists every version NUMBER this prompt has, newest first, capped at the last 100. The highest is the current one. (On a metrics row the same key is a count, not a list.) | [optional] 

## Methods

### NewPromptMeta

`func NewPromptMeta() *PromptMeta`

NewPromptMeta instantiates a new PromptMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptMetaWithDefaults

`func NewPromptMetaWithDefaults() *PromptMeta`

NewPromptMetaWithDefaults instantiates a new PromptMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *PromptMeta) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PromptMeta) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PromptMeta) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PromptMeta) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *PromptMeta) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *PromptMeta) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *PromptMeta) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *PromptMeta) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *PromptMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PromptMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PromptMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PromptMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *PromptMeta) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PromptMeta) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PromptMeta) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PromptMeta) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *PromptMeta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromptMeta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromptMeta) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PromptMeta) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersions

`func (o *PromptMeta) GetVersions() []int32`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *PromptMeta) GetVersionsOk() (*[]int32, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *PromptMeta) SetVersions(v []int32)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *PromptMeta) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


