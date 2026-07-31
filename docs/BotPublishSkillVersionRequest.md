# BotPublishSkillVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**Version** | **string** | Semver version | 
**Changelog** | **string** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**Files** | [**[]BotPublishSkillVersionRequestFilesInner**](BotPublishSkillVersionRequestFilesInner.md) |  | 

## Methods

### NewBotPublishSkillVersionRequest

`func NewBotPublishSkillVersionRequest(displayName string, version string, changelog string, files []BotPublishSkillVersionRequestFilesInner, ) *BotPublishSkillVersionRequest`

NewBotPublishSkillVersionRequest instantiates a new BotPublishSkillVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotPublishSkillVersionRequestWithDefaults

`func NewBotPublishSkillVersionRequestWithDefaults() *BotPublishSkillVersionRequest`

NewBotPublishSkillVersionRequestWithDefaults instantiates a new BotPublishSkillVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *BotPublishSkillVersionRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *BotPublishSkillVersionRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *BotPublishSkillVersionRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetVersion

`func (o *BotPublishSkillVersionRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BotPublishSkillVersionRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BotPublishSkillVersionRequest) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetChangelog

`func (o *BotPublishSkillVersionRequest) GetChangelog() string`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *BotPublishSkillVersionRequest) GetChangelogOk() (*string, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *BotPublishSkillVersionRequest) SetChangelog(v string)`

SetChangelog sets Changelog field to given value.


### GetTags

`func (o *BotPublishSkillVersionRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BotPublishSkillVersionRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BotPublishSkillVersionRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BotPublishSkillVersionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetFiles

`func (o *BotPublishSkillVersionRequest) GetFiles() []BotPublishSkillVersionRequestFilesInner`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *BotPublishSkillVersionRequest) GetFilesOk() (*[]BotPublishSkillVersionRequestFilesInner, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *BotPublishSkillVersionRequest) SetFiles(v []BotPublishSkillVersionRequestFilesInner)`

SetFiles sets Files field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


