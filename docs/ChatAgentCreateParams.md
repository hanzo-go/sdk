# ChatAgentCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Instructions** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Tools** | Pointer to **[]string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 

## Methods

### NewChatAgentCreateParams

`func NewChatAgentCreateParams(name string, ) *ChatAgentCreateParams`

NewChatAgentCreateParams instantiates a new ChatAgentCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatAgentCreateParamsWithDefaults

`func NewChatAgentCreateParamsWithDefaults() *ChatAgentCreateParams`

NewChatAgentCreateParamsWithDefaults instantiates a new ChatAgentCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ChatAgentCreateParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChatAgentCreateParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChatAgentCreateParams) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ChatAgentCreateParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChatAgentCreateParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChatAgentCreateParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChatAgentCreateParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstructions

`func (o *ChatAgentCreateParams) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ChatAgentCreateParams) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ChatAgentCreateParams) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ChatAgentCreateParams) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetModel

`func (o *ChatAgentCreateParams) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatAgentCreateParams) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatAgentCreateParams) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ChatAgentCreateParams) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTools

`func (o *ChatAgentCreateParams) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ChatAgentCreateParams) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ChatAgentCreateParams) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *ChatAgentCreateParams) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetProvider

`func (o *ChatAgentCreateParams) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ChatAgentCreateParams) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ChatAgentCreateParams) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ChatAgentCreateParams) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


