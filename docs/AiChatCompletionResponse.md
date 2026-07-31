# AiChatCompletionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **int32** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Choices** | Pointer to [**[]AiChatChoice**](AiChatChoice.md) |  | [optional] 
**Usage** | Pointer to [**AiUsage**](AiUsage.md) |  | [optional] 

## Methods

### NewAiChatCompletionResponse

`func NewAiChatCompletionResponse() *AiChatCompletionResponse`

NewAiChatCompletionResponse instantiates a new AiChatCompletionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiChatCompletionResponseWithDefaults

`func NewAiChatCompletionResponseWithDefaults() *AiChatCompletionResponse`

NewAiChatCompletionResponseWithDefaults instantiates a new AiChatCompletionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiChatCompletionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiChatCompletionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiChatCompletionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiChatCompletionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *AiChatCompletionResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AiChatCompletionResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AiChatCompletionResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AiChatCompletionResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreated

`func (o *AiChatCompletionResponse) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AiChatCompletionResponse) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AiChatCompletionResponse) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AiChatCompletionResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModel

`func (o *AiChatCompletionResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiChatCompletionResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiChatCompletionResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AiChatCompletionResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetChoices

`func (o *AiChatCompletionResponse) GetChoices() []AiChatChoice`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *AiChatCompletionResponse) GetChoicesOk() (*[]AiChatChoice, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *AiChatCompletionResponse) SetChoices(v []AiChatChoice)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *AiChatCompletionResponse) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### GetUsage

`func (o *AiChatCompletionResponse) GetUsage() AiUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AiChatCompletionResponse) GetUsageOk() (*AiUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AiChatCompletionResponse) SetUsage(v AiUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *AiChatCompletionResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


