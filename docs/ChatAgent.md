# ChatAgent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Instructions** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Tools** | Pointer to **[]string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to **map[string]interface{}** |  | [optional] 
**Author** | Pointer to **string** |  | [optional] 
**ProjectIds** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewChatAgent

`func NewChatAgent() *ChatAgent`

NewChatAgent instantiates a new ChatAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatAgentWithDefaults

`func NewChatAgentWithDefaults() *ChatAgent`

NewChatAgentWithDefaults instantiates a new ChatAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChatAgent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChatAgent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChatAgent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChatAgent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ChatAgent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChatAgent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChatAgent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChatAgent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ChatAgent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChatAgent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChatAgent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChatAgent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstructions

`func (o *ChatAgent) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ChatAgent) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ChatAgent) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ChatAgent) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetModel

`func (o *ChatAgent) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatAgent) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatAgent) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ChatAgent) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTools

`func (o *ChatAgent) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ChatAgent) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ChatAgent) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *ChatAgent) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetProvider

`func (o *ChatAgent) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ChatAgent) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ChatAgent) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ChatAgent) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetAvatar

`func (o *ChatAgent) GetAvatar() map[string]interface{}`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *ChatAgent) GetAvatarOk() (*map[string]interface{}, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *ChatAgent) SetAvatar(v map[string]interface{})`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *ChatAgent) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetAuthor

`func (o *ChatAgent) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ChatAgent) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ChatAgent) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ChatAgent) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetProjectIds

`func (o *ChatAgent) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *ChatAgent) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *ChatAgent) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *ChatAgent) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChatAgent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChatAgent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChatAgent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ChatAgent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ChatAgent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ChatAgent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ChatAgent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ChatAgent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


