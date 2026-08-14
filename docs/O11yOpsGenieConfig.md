# O11yOpsGenieConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierConfig** | Pointer to [**O11yNotifierConfig**](O11yNotifierConfig.md) |  | [optional] 
**Actions** | Pointer to **string** |  | [optional] 
**ApiKey** | Pointer to **interface{}** |  | [optional] 
**ApiKeyFile** | Pointer to **string** |  | [optional] 
**ApiUrl** | Pointer to **interface{}** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **map[string]string** |  | [optional] 
**Entity** | Pointer to **string** |  | [optional] 
**HttpConfig** | Pointer to [**O11yHTTPClientConfig**](O11yHTTPClientConfig.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **string** |  | [optional] 
**Responders** | Pointer to [**[]O11yOpsGenieConfigResponder**](O11yOpsGenieConfigResponder.md) |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**UpdateAlerts** | Pointer to **bool** |  | [optional] 

## Methods

### NewO11yOpsGenieConfig

`func NewO11yOpsGenieConfig() *O11yOpsGenieConfig`

NewO11yOpsGenieConfig instantiates a new O11yOpsGenieConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yOpsGenieConfigWithDefaults

`func NewO11yOpsGenieConfigWithDefaults() *O11yOpsGenieConfig`

NewO11yOpsGenieConfigWithDefaults instantiates a new O11yOpsGenieConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierConfig

`func (o *O11yOpsGenieConfig) GetNotifierConfig() O11yNotifierConfig`

GetNotifierConfig returns the NotifierConfig field if non-nil, zero value otherwise.

### GetNotifierConfigOk

`func (o *O11yOpsGenieConfig) GetNotifierConfigOk() (*O11yNotifierConfig, bool)`

GetNotifierConfigOk returns a tuple with the NotifierConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierConfig

`func (o *O11yOpsGenieConfig) SetNotifierConfig(v O11yNotifierConfig)`

SetNotifierConfig sets NotifierConfig field to given value.

### HasNotifierConfig

`func (o *O11yOpsGenieConfig) HasNotifierConfig() bool`

HasNotifierConfig returns a boolean if a field has been set.

### GetActions

`func (o *O11yOpsGenieConfig) GetActions() string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *O11yOpsGenieConfig) GetActionsOk() (*string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *O11yOpsGenieConfig) SetActions(v string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *O11yOpsGenieConfig) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetApiKey

`func (o *O11yOpsGenieConfig) GetApiKey() interface{}`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *O11yOpsGenieConfig) GetApiKeyOk() (*interface{}, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *O11yOpsGenieConfig) SetApiKey(v interface{})`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *O11yOpsGenieConfig) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *O11yOpsGenieConfig) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *O11yOpsGenieConfig) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetApiKeyFile

`func (o *O11yOpsGenieConfig) GetApiKeyFile() string`

GetApiKeyFile returns the ApiKeyFile field if non-nil, zero value otherwise.

### GetApiKeyFileOk

`func (o *O11yOpsGenieConfig) GetApiKeyFileOk() (*string, bool)`

GetApiKeyFileOk returns a tuple with the ApiKeyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyFile

`func (o *O11yOpsGenieConfig) SetApiKeyFile(v string)`

SetApiKeyFile sets ApiKeyFile field to given value.

### HasApiKeyFile

`func (o *O11yOpsGenieConfig) HasApiKeyFile() bool`

HasApiKeyFile returns a boolean if a field has been set.

### GetApiUrl

`func (o *O11yOpsGenieConfig) GetApiUrl() interface{}`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *O11yOpsGenieConfig) GetApiUrlOk() (*interface{}, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *O11yOpsGenieConfig) SetApiUrl(v interface{})`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *O11yOpsGenieConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### SetApiUrlNil

`func (o *O11yOpsGenieConfig) SetApiUrlNil(b bool)`

 SetApiUrlNil sets the value for ApiUrl to be an explicit nil

### UnsetApiUrl
`func (o *O11yOpsGenieConfig) UnsetApiUrl()`

UnsetApiUrl ensures that no value is present for ApiUrl, not even an explicit nil
### GetDescription

`func (o *O11yOpsGenieConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yOpsGenieConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yOpsGenieConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yOpsGenieConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetails

`func (o *O11yOpsGenieConfig) GetDetails() map[string]string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *O11yOpsGenieConfig) GetDetailsOk() (*map[string]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *O11yOpsGenieConfig) SetDetails(v map[string]string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *O11yOpsGenieConfig) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetEntity

`func (o *O11yOpsGenieConfig) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *O11yOpsGenieConfig) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *O11yOpsGenieConfig) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *O11yOpsGenieConfig) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetHttpConfig

`func (o *O11yOpsGenieConfig) GetHttpConfig() O11yHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *O11yOpsGenieConfig) GetHttpConfigOk() (*O11yHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *O11yOpsGenieConfig) SetHttpConfig(v O11yHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *O11yOpsGenieConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMessage

`func (o *O11yOpsGenieConfig) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *O11yOpsGenieConfig) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *O11yOpsGenieConfig) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *O11yOpsGenieConfig) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNote

`func (o *O11yOpsGenieConfig) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *O11yOpsGenieConfig) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *O11yOpsGenieConfig) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *O11yOpsGenieConfig) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPriority

`func (o *O11yOpsGenieConfig) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *O11yOpsGenieConfig) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *O11yOpsGenieConfig) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *O11yOpsGenieConfig) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetResponders

`func (o *O11yOpsGenieConfig) GetResponders() []O11yOpsGenieConfigResponder`

GetResponders returns the Responders field if non-nil, zero value otherwise.

### GetRespondersOk

`func (o *O11yOpsGenieConfig) GetRespondersOk() (*[]O11yOpsGenieConfigResponder, bool)`

GetRespondersOk returns a tuple with the Responders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponders

`func (o *O11yOpsGenieConfig) SetResponders(v []O11yOpsGenieConfigResponder)`

SetResponders sets Responders field to given value.

### HasResponders

`func (o *O11yOpsGenieConfig) HasResponders() bool`

HasResponders returns a boolean if a field has been set.

### GetSource

`func (o *O11yOpsGenieConfig) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *O11yOpsGenieConfig) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *O11yOpsGenieConfig) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *O11yOpsGenieConfig) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTags

`func (o *O11yOpsGenieConfig) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yOpsGenieConfig) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yOpsGenieConfig) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yOpsGenieConfig) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdateAlerts

`func (o *O11yOpsGenieConfig) GetUpdateAlerts() bool`

GetUpdateAlerts returns the UpdateAlerts field if non-nil, zero value otherwise.

### GetUpdateAlertsOk

`func (o *O11yOpsGenieConfig) GetUpdateAlertsOk() (*bool, bool)`

GetUpdateAlertsOk returns a tuple with the UpdateAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAlerts

`func (o *O11yOpsGenieConfig) SetUpdateAlerts(v bool)`

SetUpdateAlerts sets UpdateAlerts field to given value.

### HasUpdateAlerts

`func (o *O11yOpsGenieConfig) HasUpdateAlerts() bool`

HasUpdateAlerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


