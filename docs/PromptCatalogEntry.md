# PromptCatalogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **[]string** | Labels is the starter&#39;s second suggested taxonomy, same treatment as Tags. | [optional] 
**Name** | Pointer to **string** | Name is the starter&#39;s suggested handle. It is NOT taken in your org: the catalog is shared reference content, so this name is free until you import it, and posting it under a name you already use appends a version to yours. | [optional] 
**Prompt** | Pointer to **string** | Prompt is the starter&#39;s full template body, ready to POST as-is. Entries too large to create (over 64 KiB) are dropped from this list rather than offered. | [optional] 
**Tags** | Pointer to **[]string** | Tags is the starter&#39;s suggested taxonomy, carried through unchanged if you import it. | [optional] 
**Type** | Pointer to **string** | Type labels the template&#39;s kind, defaulted to \&quot;text\&quot; for entries that declare none. | [optional] 

## Methods

### NewPromptCatalogEntry

`func NewPromptCatalogEntry() *PromptCatalogEntry`

NewPromptCatalogEntry instantiates a new PromptCatalogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptCatalogEntryWithDefaults

`func NewPromptCatalogEntryWithDefaults() *PromptCatalogEntry`

NewPromptCatalogEntryWithDefaults instantiates a new PromptCatalogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *PromptCatalogEntry) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PromptCatalogEntry) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PromptCatalogEntry) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PromptCatalogEntry) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *PromptCatalogEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PromptCatalogEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PromptCatalogEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PromptCatalogEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrompt

`func (o *PromptCatalogEntry) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *PromptCatalogEntry) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *PromptCatalogEntry) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *PromptCatalogEntry) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetTags

`func (o *PromptCatalogEntry) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PromptCatalogEntry) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PromptCatalogEntry) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PromptCatalogEntry) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *PromptCatalogEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromptCatalogEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromptCatalogEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PromptCatalogEntry) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


