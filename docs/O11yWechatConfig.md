# O11yWechatConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierConfig** | Pointer to [**O11yNotifierConfig**](O11yNotifierConfig.md) |  | [optional] 
**AgentId** | Pointer to **string** |  | [optional] 
**ApiSecret** | Pointer to **interface{}** |  | [optional] 
**ApiSecretFile** | Pointer to **string** |  | [optional] 
**ApiUrl** | Pointer to **interface{}** |  | [optional] 
**CorpId** | Pointer to **string** |  | [optional] 
**HttpConfig** | Pointer to [**O11yHTTPClientConfig**](O11yHTTPClientConfig.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**MessageType** | Pointer to **string** |  | [optional] 
**ToParty** | Pointer to **string** |  | [optional] 
**ToTag** | Pointer to **string** |  | [optional] 
**ToUser** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yWechatConfig

`func NewO11yWechatConfig() *O11yWechatConfig`

NewO11yWechatConfig instantiates a new O11yWechatConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yWechatConfigWithDefaults

`func NewO11yWechatConfigWithDefaults() *O11yWechatConfig`

NewO11yWechatConfigWithDefaults instantiates a new O11yWechatConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierConfig

`func (o *O11yWechatConfig) GetNotifierConfig() O11yNotifierConfig`

GetNotifierConfig returns the NotifierConfig field if non-nil, zero value otherwise.

### GetNotifierConfigOk

`func (o *O11yWechatConfig) GetNotifierConfigOk() (*O11yNotifierConfig, bool)`

GetNotifierConfigOk returns a tuple with the NotifierConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierConfig

`func (o *O11yWechatConfig) SetNotifierConfig(v O11yNotifierConfig)`

SetNotifierConfig sets NotifierConfig field to given value.

### HasNotifierConfig

`func (o *O11yWechatConfig) HasNotifierConfig() bool`

HasNotifierConfig returns a boolean if a field has been set.

### GetAgentId

`func (o *O11yWechatConfig) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *O11yWechatConfig) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *O11yWechatConfig) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *O11yWechatConfig) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetApiSecret

`func (o *O11yWechatConfig) GetApiSecret() interface{}`

GetApiSecret returns the ApiSecret field if non-nil, zero value otherwise.

### GetApiSecretOk

`func (o *O11yWechatConfig) GetApiSecretOk() (*interface{}, bool)`

GetApiSecretOk returns a tuple with the ApiSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSecret

`func (o *O11yWechatConfig) SetApiSecret(v interface{})`

SetApiSecret sets ApiSecret field to given value.

### HasApiSecret

`func (o *O11yWechatConfig) HasApiSecret() bool`

HasApiSecret returns a boolean if a field has been set.

### SetApiSecretNil

`func (o *O11yWechatConfig) SetApiSecretNil(b bool)`

 SetApiSecretNil sets the value for ApiSecret to be an explicit nil

### UnsetApiSecret
`func (o *O11yWechatConfig) UnsetApiSecret()`

UnsetApiSecret ensures that no value is present for ApiSecret, not even an explicit nil
### GetApiSecretFile

`func (o *O11yWechatConfig) GetApiSecretFile() string`

GetApiSecretFile returns the ApiSecretFile field if non-nil, zero value otherwise.

### GetApiSecretFileOk

`func (o *O11yWechatConfig) GetApiSecretFileOk() (*string, bool)`

GetApiSecretFileOk returns a tuple with the ApiSecretFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSecretFile

`func (o *O11yWechatConfig) SetApiSecretFile(v string)`

SetApiSecretFile sets ApiSecretFile field to given value.

### HasApiSecretFile

`func (o *O11yWechatConfig) HasApiSecretFile() bool`

HasApiSecretFile returns a boolean if a field has been set.

### GetApiUrl

`func (o *O11yWechatConfig) GetApiUrl() interface{}`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *O11yWechatConfig) GetApiUrlOk() (*interface{}, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *O11yWechatConfig) SetApiUrl(v interface{})`

SetApiUrl sets ApiUrl field to given value.

### HasApiUrl

`func (o *O11yWechatConfig) HasApiUrl() bool`

HasApiUrl returns a boolean if a field has been set.

### SetApiUrlNil

`func (o *O11yWechatConfig) SetApiUrlNil(b bool)`

 SetApiUrlNil sets the value for ApiUrl to be an explicit nil

### UnsetApiUrl
`func (o *O11yWechatConfig) UnsetApiUrl()`

UnsetApiUrl ensures that no value is present for ApiUrl, not even an explicit nil
### GetCorpId

`func (o *O11yWechatConfig) GetCorpId() string`

GetCorpId returns the CorpId field if non-nil, zero value otherwise.

### GetCorpIdOk

`func (o *O11yWechatConfig) GetCorpIdOk() (*string, bool)`

GetCorpIdOk returns a tuple with the CorpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorpId

`func (o *O11yWechatConfig) SetCorpId(v string)`

SetCorpId sets CorpId field to given value.

### HasCorpId

`func (o *O11yWechatConfig) HasCorpId() bool`

HasCorpId returns a boolean if a field has been set.

### GetHttpConfig

`func (o *O11yWechatConfig) GetHttpConfig() O11yHTTPClientConfig`

GetHttpConfig returns the HttpConfig field if non-nil, zero value otherwise.

### GetHttpConfigOk

`func (o *O11yWechatConfig) GetHttpConfigOk() (*O11yHTTPClientConfig, bool)`

GetHttpConfigOk returns a tuple with the HttpConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConfig

`func (o *O11yWechatConfig) SetHttpConfig(v O11yHTTPClientConfig)`

SetHttpConfig sets HttpConfig field to given value.

### HasHttpConfig

`func (o *O11yWechatConfig) HasHttpConfig() bool`

HasHttpConfig returns a boolean if a field has been set.

### GetMessage

`func (o *O11yWechatConfig) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *O11yWechatConfig) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *O11yWechatConfig) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *O11yWechatConfig) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageType

`func (o *O11yWechatConfig) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *O11yWechatConfig) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *O11yWechatConfig) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *O11yWechatConfig) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetToParty

`func (o *O11yWechatConfig) GetToParty() string`

GetToParty returns the ToParty field if non-nil, zero value otherwise.

### GetToPartyOk

`func (o *O11yWechatConfig) GetToPartyOk() (*string, bool)`

GetToPartyOk returns a tuple with the ToParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToParty

`func (o *O11yWechatConfig) SetToParty(v string)`

SetToParty sets ToParty field to given value.

### HasToParty

`func (o *O11yWechatConfig) HasToParty() bool`

HasToParty returns a boolean if a field has been set.

### GetToTag

`func (o *O11yWechatConfig) GetToTag() string`

GetToTag returns the ToTag field if non-nil, zero value otherwise.

### GetToTagOk

`func (o *O11yWechatConfig) GetToTagOk() (*string, bool)`

GetToTagOk returns a tuple with the ToTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTag

`func (o *O11yWechatConfig) SetToTag(v string)`

SetToTag sets ToTag field to given value.

### HasToTag

`func (o *O11yWechatConfig) HasToTag() bool`

HasToTag returns a boolean if a field has been set.

### GetToUser

`func (o *O11yWechatConfig) GetToUser() string`

GetToUser returns the ToUser field if non-nil, zero value otherwise.

### GetToUserOk

`func (o *O11yWechatConfig) GetToUserOk() (*string, bool)`

GetToUserOk returns a tuple with the ToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUser

`func (o *O11yWechatConfig) SetToUser(v string)`

SetToUser sets ToUser field to given value.

### HasToUser

`func (o *O11yWechatConfig) HasToUser() bool`

HasToUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


