# ChatPostUserPluginsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginKey** | Pointer to **string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Auth** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewChatPostUserPluginsRequest

`func NewChatPostUserPluginsRequest() *ChatPostUserPluginsRequest`

NewChatPostUserPluginsRequest instantiates a new ChatPostUserPluginsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatPostUserPluginsRequestWithDefaults

`func NewChatPostUserPluginsRequestWithDefaults() *ChatPostUserPluginsRequest`

NewChatPostUserPluginsRequestWithDefaults instantiates a new ChatPostUserPluginsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginKey

`func (o *ChatPostUserPluginsRequest) GetPluginKey() string`

GetPluginKey returns the PluginKey field if non-nil, zero value otherwise.

### GetPluginKeyOk

`func (o *ChatPostUserPluginsRequest) GetPluginKeyOk() (*string, bool)`

GetPluginKeyOk returns a tuple with the PluginKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginKey

`func (o *ChatPostUserPluginsRequest) SetPluginKey(v string)`

SetPluginKey sets PluginKey field to given value.

### HasPluginKey

`func (o *ChatPostUserPluginsRequest) HasPluginKey() bool`

HasPluginKey returns a boolean if a field has been set.

### GetAction

`func (o *ChatPostUserPluginsRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ChatPostUserPluginsRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ChatPostUserPluginsRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ChatPostUserPluginsRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAuth

`func (o *ChatPostUserPluginsRequest) GetAuth() map[string]interface{}`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *ChatPostUserPluginsRequest) GetAuthOk() (*map[string]interface{}, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *ChatPostUserPluginsRequest) SetAuth(v map[string]interface{})`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *ChatPostUserPluginsRequest) HasAuth() bool`

HasAuth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


