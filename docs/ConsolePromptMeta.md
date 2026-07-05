# ConsolePromptMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Versions** | Pointer to **[]int32** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**LastConfig** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewConsolePromptMeta

`func NewConsolePromptMeta() *ConsolePromptMeta`

NewConsolePromptMeta instantiates a new ConsolePromptMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsolePromptMetaWithDefaults

`func NewConsolePromptMetaWithDefaults() *ConsolePromptMeta`

NewConsolePromptMetaWithDefaults instantiates a new ConsolePromptMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConsolePromptMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsolePromptMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsolePromptMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConsolePromptMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersions

`func (o *ConsolePromptMeta) GetVersions() []int32`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ConsolePromptMeta) GetVersionsOk() (*[]int32, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ConsolePromptMeta) SetVersions(v []int32)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ConsolePromptMeta) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetLabels

`func (o *ConsolePromptMeta) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ConsolePromptMeta) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ConsolePromptMeta) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ConsolePromptMeta) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *ConsolePromptMeta) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConsolePromptMeta) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConsolePromptMeta) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConsolePromptMeta) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *ConsolePromptMeta) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *ConsolePromptMeta) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *ConsolePromptMeta) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *ConsolePromptMeta) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetLastConfig

`func (o *ConsolePromptMeta) GetLastConfig() map[string]interface{}`

GetLastConfig returns the LastConfig field if non-nil, zero value otherwise.

### GetLastConfigOk

`func (o *ConsolePromptMeta) GetLastConfigOk() (*map[string]interface{}, bool)`

GetLastConfigOk returns a tuple with the LastConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConfig

`func (o *ConsolePromptMeta) SetLastConfig(v map[string]interface{})`

SetLastConfig sets LastConfig field to given value.

### HasLastConfig

`func (o *ConsolePromptMeta) HasLastConfig() bool`

HasLastConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


