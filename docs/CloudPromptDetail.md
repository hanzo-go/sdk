# CloudPromptDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**LastUpdatedAt** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Prompt** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**VersionHistory** | Pointer to [**[]CloudVersionView**](CloudVersionView.md) |  | [optional] 

## Methods

### NewCloudPromptDetail

`func NewCloudPromptDetail() *CloudPromptDetail`

NewCloudPromptDetail instantiates a new CloudPromptDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPromptDetailWithDefaults

`func NewCloudPromptDetailWithDefaults() *CloudPromptDetail`

NewCloudPromptDetailWithDefaults instantiates a new CloudPromptDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CloudPromptDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudPromptDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudPromptDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudPromptDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLabels

`func (o *CloudPromptDetail) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CloudPromptDetail) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CloudPromptDetail) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CloudPromptDetail) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *CloudPromptDetail) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *CloudPromptDetail) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *CloudPromptDetail) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *CloudPromptDetail) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *CloudPromptDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudPromptDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudPromptDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudPromptDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrompt

`func (o *CloudPromptDetail) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *CloudPromptDetail) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *CloudPromptDetail) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *CloudPromptDetail) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetTags

`func (o *CloudPromptDetail) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CloudPromptDetail) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CloudPromptDetail) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CloudPromptDetail) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *CloudPromptDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudPromptDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudPromptDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudPromptDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *CloudPromptDetail) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CloudPromptDetail) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CloudPromptDetail) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CloudPromptDetail) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionHistory

`func (o *CloudPromptDetail) GetVersionHistory() []CloudVersionView`

GetVersionHistory returns the VersionHistory field if non-nil, zero value otherwise.

### GetVersionHistoryOk

`func (o *CloudPromptDetail) GetVersionHistoryOk() (*[]CloudVersionView, bool)`

GetVersionHistoryOk returns a tuple with the VersionHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionHistory

`func (o *CloudPromptDetail) SetVersionHistory(v []CloudVersionView)`

SetVersionHistory sets VersionHistory field to given value.

### HasVersionHistory

`func (o *CloudPromptDetail) HasVersionHistory() bool`

HasVersionHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


