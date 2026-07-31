# GatewayChatCompletionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **int32** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Choices** | Pointer to [**[]GatewayChatCompletionResponseChoicesInner**](GatewayChatCompletionResponseChoicesInner.md) |  | [optional] 
**Usage** | Pointer to [**GatewayChatCompletionResponseUsage**](GatewayChatCompletionResponseUsage.md) |  | [optional] 

## Methods

### NewGatewayChatCompletionResponse

`func NewGatewayChatCompletionResponse() *GatewayChatCompletionResponse`

NewGatewayChatCompletionResponse instantiates a new GatewayChatCompletionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayChatCompletionResponseWithDefaults

`func NewGatewayChatCompletionResponseWithDefaults() *GatewayChatCompletionResponse`

NewGatewayChatCompletionResponseWithDefaults instantiates a new GatewayChatCompletionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GatewayChatCompletionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayChatCompletionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayChatCompletionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewayChatCompletionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *GatewayChatCompletionResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GatewayChatCompletionResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GatewayChatCompletionResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *GatewayChatCompletionResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreated

`func (o *GatewayChatCompletionResponse) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GatewayChatCompletionResponse) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GatewayChatCompletionResponse) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GatewayChatCompletionResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModel

`func (o *GatewayChatCompletionResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GatewayChatCompletionResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GatewayChatCompletionResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GatewayChatCompletionResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetChoices

`func (o *GatewayChatCompletionResponse) GetChoices() []GatewayChatCompletionResponseChoicesInner`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *GatewayChatCompletionResponse) GetChoicesOk() (*[]GatewayChatCompletionResponseChoicesInner, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *GatewayChatCompletionResponse) SetChoices(v []GatewayChatCompletionResponseChoicesInner)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *GatewayChatCompletionResponse) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### GetUsage

`func (o *GatewayChatCompletionResponse) GetUsage() GatewayChatCompletionResponseUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GatewayChatCompletionResponse) GetUsageOk() (*GatewayChatCompletionResponseUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GatewayChatCompletionResponse) SetUsage(v GatewayChatCompletionResponseUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *GatewayChatCompletionResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


