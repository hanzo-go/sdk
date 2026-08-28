# OpenaiChatCompletionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Choices** | Pointer to [**[]OpenaiChatCompletionChoice**](OpenaiChatCompletionChoice.md) |  | [optional] 
**Created** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**PromptFilterResults** | Pointer to [**[]OpenaiPromptFilterResult**](OpenaiPromptFilterResult.md) |  | [optional] 
**SystemFingerprint** | Pointer to **string** |  | [optional] 
**Usage** | Pointer to [**OpenaiUsage**](OpenaiUsage.md) |  | [optional] 

## Methods

### NewOpenaiChatCompletionResponse

`func NewOpenaiChatCompletionResponse() *OpenaiChatCompletionResponse`

NewOpenaiChatCompletionResponse instantiates a new OpenaiChatCompletionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenaiChatCompletionResponseWithDefaults

`func NewOpenaiChatCompletionResponseWithDefaults() *OpenaiChatCompletionResponse`

NewOpenaiChatCompletionResponseWithDefaults instantiates a new OpenaiChatCompletionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoices

`func (o *OpenaiChatCompletionResponse) GetChoices() []OpenaiChatCompletionChoice`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *OpenaiChatCompletionResponse) GetChoicesOk() (*[]OpenaiChatCompletionChoice, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *OpenaiChatCompletionResponse) SetChoices(v []OpenaiChatCompletionChoice)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *OpenaiChatCompletionResponse) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### GetCreated

`func (o *OpenaiChatCompletionResponse) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OpenaiChatCompletionResponse) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OpenaiChatCompletionResponse) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OpenaiChatCompletionResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *OpenaiChatCompletionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenaiChatCompletionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenaiChatCompletionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OpenaiChatCompletionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *OpenaiChatCompletionResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OpenaiChatCompletionResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OpenaiChatCompletionResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OpenaiChatCompletionResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetObject

`func (o *OpenaiChatCompletionResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OpenaiChatCompletionResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OpenaiChatCompletionResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *OpenaiChatCompletionResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPromptFilterResults

`func (o *OpenaiChatCompletionResponse) GetPromptFilterResults() []OpenaiPromptFilterResult`

GetPromptFilterResults returns the PromptFilterResults field if non-nil, zero value otherwise.

### GetPromptFilterResultsOk

`func (o *OpenaiChatCompletionResponse) GetPromptFilterResultsOk() (*[]OpenaiPromptFilterResult, bool)`

GetPromptFilterResultsOk returns a tuple with the PromptFilterResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptFilterResults

`func (o *OpenaiChatCompletionResponse) SetPromptFilterResults(v []OpenaiPromptFilterResult)`

SetPromptFilterResults sets PromptFilterResults field to given value.

### HasPromptFilterResults

`func (o *OpenaiChatCompletionResponse) HasPromptFilterResults() bool`

HasPromptFilterResults returns a boolean if a field has been set.

### GetSystemFingerprint

`func (o *OpenaiChatCompletionResponse) GetSystemFingerprint() string`

GetSystemFingerprint returns the SystemFingerprint field if non-nil, zero value otherwise.

### GetSystemFingerprintOk

`func (o *OpenaiChatCompletionResponse) GetSystemFingerprintOk() (*string, bool)`

GetSystemFingerprintOk returns a tuple with the SystemFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemFingerprint

`func (o *OpenaiChatCompletionResponse) SetSystemFingerprint(v string)`

SetSystemFingerprint sets SystemFingerprint field to given value.

### HasSystemFingerprint

`func (o *OpenaiChatCompletionResponse) HasSystemFingerprint() bool`

HasSystemFingerprint returns a boolean if a field has been set.

### GetUsage

`func (o *OpenaiChatCompletionResponse) GetUsage() OpenaiUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *OpenaiChatCompletionResponse) GetUsageOk() (*OpenaiUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *OpenaiChatCompletionResponse) SetUsage(v OpenaiUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *OpenaiChatCompletionResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


