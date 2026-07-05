# ChatChatCompletion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **int32** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Choices** | Pointer to [**[]ChatChatCompletionChoicesInner**](ChatChatCompletionChoicesInner.md) |  | [optional] 
**Usage** | Pointer to [**ChatChatCompletionUsage**](ChatChatCompletionUsage.md) |  | [optional] 

## Methods

### NewChatChatCompletion

`func NewChatChatCompletion() *ChatChatCompletion`

NewChatChatCompletion instantiates a new ChatChatCompletion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatChatCompletionWithDefaults

`func NewChatChatCompletionWithDefaults() *ChatChatCompletion`

NewChatChatCompletionWithDefaults instantiates a new ChatChatCompletion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChatChatCompletion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChatChatCompletion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChatChatCompletion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChatChatCompletion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *ChatChatCompletion) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChatChatCompletion) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChatChatCompletion) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ChatChatCompletion) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreated

`func (o *ChatChatCompletion) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ChatChatCompletion) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ChatChatCompletion) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ChatChatCompletion) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModel

`func (o *ChatChatCompletion) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatChatCompletion) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatChatCompletion) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ChatChatCompletion) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetChoices

`func (o *ChatChatCompletion) GetChoices() []ChatChatCompletionChoicesInner`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *ChatChatCompletion) GetChoicesOk() (*[]ChatChatCompletionChoicesInner, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *ChatChatCompletion) SetChoices(v []ChatChatCompletionChoicesInner)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *ChatChatCompletion) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### GetUsage

`func (o *ChatChatCompletion) GetUsage() ChatChatCompletionUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ChatChatCompletion) GetUsageOk() (*ChatChatCompletionUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ChatChatCompletion) SetUsage(v ChatChatCompletionUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ChatChatCompletion) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


