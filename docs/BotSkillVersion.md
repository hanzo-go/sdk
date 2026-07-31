# BotSkillVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** | Semver version string | [optional] 
**Changelog** | Pointer to **string** |  | [optional] 
**ChangelogSource** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Sha256hash** | Pointer to **string** |  | [optional] 
**VtAnalysis** | Pointer to **map[string]interface{}** |  | [optional] 
**LlmAnalysis** | Pointer to **map[string]interface{}** |  | [optional] 
**Files** | Pointer to [**[]BotGetSkillVersionFiles200ResponseFilesInner**](BotGetSkillVersionFiles200ResponseFilesInner.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewBotSkillVersion

`func NewBotSkillVersion() *BotSkillVersion`

NewBotSkillVersion instantiates a new BotSkillVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotSkillVersionWithDefaults

`func NewBotSkillVersionWithDefaults() *BotSkillVersion`

NewBotSkillVersionWithDefaults instantiates a new BotSkillVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BotSkillVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BotSkillVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BotSkillVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BotSkillVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *BotSkillVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BotSkillVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BotSkillVersion) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BotSkillVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetChangelog

`func (o *BotSkillVersion) GetChangelog() string`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *BotSkillVersion) GetChangelogOk() (*string, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *BotSkillVersion) SetChangelog(v string)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *BotSkillVersion) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.

### GetChangelogSource

`func (o *BotSkillVersion) GetChangelogSource() string`

GetChangelogSource returns the ChangelogSource field if non-nil, zero value otherwise.

### GetChangelogSourceOk

`func (o *BotSkillVersion) GetChangelogSourceOk() (*string, bool)`

GetChangelogSourceOk returns a tuple with the ChangelogSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogSource

`func (o *BotSkillVersion) SetChangelogSource(v string)`

SetChangelogSource sets ChangelogSource field to given value.

### HasChangelogSource

`func (o *BotSkillVersion) HasChangelogSource() bool`

HasChangelogSource returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BotSkillVersion) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BotSkillVersion) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BotSkillVersion) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BotSkillVersion) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetSha256hash

`func (o *BotSkillVersion) GetSha256hash() string`

GetSha256hash returns the Sha256hash field if non-nil, zero value otherwise.

### GetSha256hashOk

`func (o *BotSkillVersion) GetSha256hashOk() (*string, bool)`

GetSha256hashOk returns a tuple with the Sha256hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256hash

`func (o *BotSkillVersion) SetSha256hash(v string)`

SetSha256hash sets Sha256hash field to given value.

### HasSha256hash

`func (o *BotSkillVersion) HasSha256hash() bool`

HasSha256hash returns a boolean if a field has been set.

### GetVtAnalysis

`func (o *BotSkillVersion) GetVtAnalysis() map[string]interface{}`

GetVtAnalysis returns the VtAnalysis field if non-nil, zero value otherwise.

### GetVtAnalysisOk

`func (o *BotSkillVersion) GetVtAnalysisOk() (*map[string]interface{}, bool)`

GetVtAnalysisOk returns a tuple with the VtAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVtAnalysis

`func (o *BotSkillVersion) SetVtAnalysis(v map[string]interface{})`

SetVtAnalysis sets VtAnalysis field to given value.

### HasVtAnalysis

`func (o *BotSkillVersion) HasVtAnalysis() bool`

HasVtAnalysis returns a boolean if a field has been set.

### GetLlmAnalysis

`func (o *BotSkillVersion) GetLlmAnalysis() map[string]interface{}`

GetLlmAnalysis returns the LlmAnalysis field if non-nil, zero value otherwise.

### GetLlmAnalysisOk

`func (o *BotSkillVersion) GetLlmAnalysisOk() (*map[string]interface{}, bool)`

GetLlmAnalysisOk returns a tuple with the LlmAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmAnalysis

`func (o *BotSkillVersion) SetLlmAnalysis(v map[string]interface{})`

SetLlmAnalysis sets LlmAnalysis field to given value.

### HasLlmAnalysis

`func (o *BotSkillVersion) HasLlmAnalysis() bool`

HasLlmAnalysis returns a boolean if a field has been set.

### GetFiles

`func (o *BotSkillVersion) GetFiles() []BotGetSkillVersionFiles200ResponseFilesInner`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *BotSkillVersion) GetFilesOk() (*[]BotGetSkillVersionFiles200ResponseFilesInner, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *BotSkillVersion) SetFiles(v []BotGetSkillVersionFiles200ResponseFilesInner)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *BotSkillVersion) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BotSkillVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BotSkillVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BotSkillVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BotSkillVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


