# PostChat200Response

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
**Routing** | Pointer to [**PostChat200ResponseAllOfRouting**](PostChat200ResponseAllOfRouting.md) |  | [optional] 

## Methods

### NewPostChat200Response

`func NewPostChat200Response() *PostChat200Response`

NewPostChat200Response instantiates a new PostChat200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostChat200ResponseWithDefaults

`func NewPostChat200ResponseWithDefaults() *PostChat200Response`

NewPostChat200ResponseWithDefaults instantiates a new PostChat200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChoices

`func (o *PostChat200Response) GetChoices() []OpenaiChatCompletionChoice`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *PostChat200Response) GetChoicesOk() (*[]OpenaiChatCompletionChoice, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *PostChat200Response) SetChoices(v []OpenaiChatCompletionChoice)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *PostChat200Response) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### GetCreated

`func (o *PostChat200Response) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PostChat200Response) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PostChat200Response) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PostChat200Response) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *PostChat200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostChat200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostChat200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PostChat200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *PostChat200Response) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PostChat200Response) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PostChat200Response) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *PostChat200Response) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetObject

`func (o *PostChat200Response) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PostChat200Response) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PostChat200Response) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *PostChat200Response) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPromptFilterResults

`func (o *PostChat200Response) GetPromptFilterResults() []OpenaiPromptFilterResult`

GetPromptFilterResults returns the PromptFilterResults field if non-nil, zero value otherwise.

### GetPromptFilterResultsOk

`func (o *PostChat200Response) GetPromptFilterResultsOk() (*[]OpenaiPromptFilterResult, bool)`

GetPromptFilterResultsOk returns a tuple with the PromptFilterResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptFilterResults

`func (o *PostChat200Response) SetPromptFilterResults(v []OpenaiPromptFilterResult)`

SetPromptFilterResults sets PromptFilterResults field to given value.

### HasPromptFilterResults

`func (o *PostChat200Response) HasPromptFilterResults() bool`

HasPromptFilterResults returns a boolean if a field has been set.

### GetSystemFingerprint

`func (o *PostChat200Response) GetSystemFingerprint() string`

GetSystemFingerprint returns the SystemFingerprint field if non-nil, zero value otherwise.

### GetSystemFingerprintOk

`func (o *PostChat200Response) GetSystemFingerprintOk() (*string, bool)`

GetSystemFingerprintOk returns a tuple with the SystemFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemFingerprint

`func (o *PostChat200Response) SetSystemFingerprint(v string)`

SetSystemFingerprint sets SystemFingerprint field to given value.

### HasSystemFingerprint

`func (o *PostChat200Response) HasSystemFingerprint() bool`

HasSystemFingerprint returns a boolean if a field has been set.

### GetUsage

`func (o *PostChat200Response) GetUsage() OpenaiUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *PostChat200Response) GetUsageOk() (*OpenaiUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *PostChat200Response) SetUsage(v OpenaiUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *PostChat200Response) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetRouting

`func (o *PostChat200Response) GetRouting() PostChat200ResponseAllOfRouting`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *PostChat200Response) GetRoutingOk() (*PostChat200ResponseAllOfRouting, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *PostChat200Response) SetRouting(v PostChat200ResponseAllOfRouting)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *PostChat200Response) HasRouting() bool`

HasRouting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


