# ChatAgentApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** | Only returned on creation | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**LastUsedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewChatAgentApiKey

`func NewChatAgentApiKey() *ChatAgentApiKey`

NewChatAgentApiKey instantiates a new ChatAgentApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatAgentApiKeyWithDefaults

`func NewChatAgentApiKeyWithDefaults() *ChatAgentApiKey`

NewChatAgentApiKeyWithDefaults instantiates a new ChatAgentApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChatAgentApiKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChatAgentApiKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChatAgentApiKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChatAgentApiKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ChatAgentApiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChatAgentApiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChatAgentApiKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChatAgentApiKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *ChatAgentApiKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ChatAgentApiKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ChatAgentApiKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ChatAgentApiKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChatAgentApiKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChatAgentApiKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChatAgentApiKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ChatAgentApiKey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *ChatAgentApiKey) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ChatAgentApiKey) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ChatAgentApiKey) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *ChatAgentApiKey) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


