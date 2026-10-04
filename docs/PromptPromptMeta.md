# PromptPromptMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **[]string** | Labels is the creator&#39;s free-form taxonomy, stored as given after trimming and de-duplication. Always present, &#x60;[]&#x60; when none — never null. | [optional] 
**LastUpdatedAt** | Pointer to **string** | LastUpdatedAt is when the newest version was appended, RFC 3339 UTC. Empty only if the record carries no timestamp at all. | [optional] 
**Name** | Pointer to **string** | Name is the prompt&#39;s org-unique handle and the URL segment it is fetched by: GET /v1/prompt/&lt;name&gt;. | [optional] 
**Tags** | Pointer to **[]string** | Tags is the second free-form taxonomy under the same rules as Labels. Nothing in this service interprets either; they are yours to organize by. | [optional] 
**Type** | Pointer to **string** | Type labels the template&#39;s kind, \&quot;text\&quot; unless the creator said otherwise. It is the CURRENT version&#39;s type; earlier versions may carry a different one. | [optional] 
**Versions** | Pointer to **[]int64** | Versions lists every version NUMBER this prompt has, newest first, capped at the last 100. The highest is the current one. (On a metrics row the same key is a count, not a list.) | [optional] 

## Methods

### NewPromptPromptMeta

`func NewPromptPromptMeta() *PromptPromptMeta`

NewPromptPromptMeta instantiates a new PromptPromptMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptPromptMetaWithDefaults

`func NewPromptPromptMetaWithDefaults() *PromptPromptMeta`

NewPromptPromptMetaWithDefaults instantiates a new PromptPromptMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *PromptPromptMeta) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PromptPromptMeta) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PromptPromptMeta) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PromptPromptMeta) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *PromptPromptMeta) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *PromptPromptMeta) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *PromptPromptMeta) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *PromptPromptMeta) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *PromptPromptMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PromptPromptMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PromptPromptMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PromptPromptMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *PromptPromptMeta) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PromptPromptMeta) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PromptPromptMeta) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PromptPromptMeta) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *PromptPromptMeta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromptPromptMeta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromptPromptMeta) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PromptPromptMeta) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersions

`func (o *PromptPromptMeta) GetVersions() []int64`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *PromptPromptMeta) GetVersionsOk() (*[]int64, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *PromptPromptMeta) SetVersions(v []int64)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *PromptPromptMeta) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


