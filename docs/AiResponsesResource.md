# AiResponsesResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **interface{}** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IncompleteDetails** | Pointer to **interface{}** |  | [optional] 
**Instructions** | Pointer to **string** |  | [optional] 
**MaxOutputTokens** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Output** | Pointer to **[]interface{}** |  | [optional] 
**ParallelToolCalls** | Pointer to **bool** |  | [optional] 
**PreviousResponseId** | Pointer to **interface{}** |  | [optional] 
**Reasoning** | Pointer to **interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Store** | Pointer to **bool** |  | [optional] 
**Temperature** | Pointer to **float32** |  | [optional] 
**Text** | Pointer to **interface{}** |  | [optional] 
**ToolChoice** | Pointer to **interface{}** |  | [optional] 
**Tools** | Pointer to [**[]AiResponsesTool**](AiResponsesTool.md) |  | [optional] 
**TopP** | Pointer to **float32** |  | [optional] 
**Usage** | Pointer to [**AiResponsesUsage**](AiResponsesUsage.md) |  | [optional] 

## Methods

### NewAiResponsesResource

`func NewAiResponsesResource() *AiResponsesResource`

NewAiResponsesResource instantiates a new AiResponsesResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiResponsesResourceWithDefaults

`func NewAiResponsesResourceWithDefaults() *AiResponsesResource`

NewAiResponsesResourceWithDefaults instantiates a new AiResponsesResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AiResponsesResource) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AiResponsesResource) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AiResponsesResource) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AiResponsesResource) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetError

`func (o *AiResponsesResource) GetError() interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AiResponsesResource) GetErrorOk() (*interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AiResponsesResource) SetError(v interface{})`

SetError sets Error field to given value.

### HasError

`func (o *AiResponsesResource) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *AiResponsesResource) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *AiResponsesResource) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetId

`func (o *AiResponsesResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiResponsesResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiResponsesResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiResponsesResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncompleteDetails

`func (o *AiResponsesResource) GetIncompleteDetails() interface{}`

GetIncompleteDetails returns the IncompleteDetails field if non-nil, zero value otherwise.

### GetIncompleteDetailsOk

`func (o *AiResponsesResource) GetIncompleteDetailsOk() (*interface{}, bool)`

GetIncompleteDetailsOk returns a tuple with the IncompleteDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompleteDetails

`func (o *AiResponsesResource) SetIncompleteDetails(v interface{})`

SetIncompleteDetails sets IncompleteDetails field to given value.

### HasIncompleteDetails

`func (o *AiResponsesResource) HasIncompleteDetails() bool`

HasIncompleteDetails returns a boolean if a field has been set.

### SetIncompleteDetailsNil

`func (o *AiResponsesResource) SetIncompleteDetailsNil(b bool)`

 SetIncompleteDetailsNil sets the value for IncompleteDetails to be an explicit nil

### UnsetIncompleteDetails
`func (o *AiResponsesResource) UnsetIncompleteDetails()`

UnsetIncompleteDetails ensures that no value is present for IncompleteDetails, not even an explicit nil
### GetInstructions

`func (o *AiResponsesResource) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AiResponsesResource) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AiResponsesResource) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AiResponsesResource) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetMaxOutputTokens

`func (o *AiResponsesResource) GetMaxOutputTokens() int32`

GetMaxOutputTokens returns the MaxOutputTokens field if non-nil, zero value otherwise.

### GetMaxOutputTokensOk

`func (o *AiResponsesResource) GetMaxOutputTokensOk() (*int32, bool)`

GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOutputTokens

`func (o *AiResponsesResource) SetMaxOutputTokens(v int32)`

SetMaxOutputTokens sets MaxOutputTokens field to given value.

### HasMaxOutputTokens

`func (o *AiResponsesResource) HasMaxOutputTokens() bool`

HasMaxOutputTokens returns a boolean if a field has been set.

### GetMetadata

`func (o *AiResponsesResource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AiResponsesResource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AiResponsesResource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AiResponsesResource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetModel

`func (o *AiResponsesResource) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiResponsesResource) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiResponsesResource) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AiResponsesResource) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetObject

`func (o *AiResponsesResource) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AiResponsesResource) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AiResponsesResource) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AiResponsesResource) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetOutput

`func (o *AiResponsesResource) GetOutput() []interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *AiResponsesResource) GetOutputOk() (*[]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *AiResponsesResource) SetOutput(v []interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *AiResponsesResource) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetParallelToolCalls

`func (o *AiResponsesResource) GetParallelToolCalls() bool`

GetParallelToolCalls returns the ParallelToolCalls field if non-nil, zero value otherwise.

### GetParallelToolCallsOk

`func (o *AiResponsesResource) GetParallelToolCallsOk() (*bool, bool)`

GetParallelToolCallsOk returns a tuple with the ParallelToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelToolCalls

`func (o *AiResponsesResource) SetParallelToolCalls(v bool)`

SetParallelToolCalls sets ParallelToolCalls field to given value.

### HasParallelToolCalls

`func (o *AiResponsesResource) HasParallelToolCalls() bool`

HasParallelToolCalls returns a boolean if a field has been set.

### GetPreviousResponseId

`func (o *AiResponsesResource) GetPreviousResponseId() interface{}`

GetPreviousResponseId returns the PreviousResponseId field if non-nil, zero value otherwise.

### GetPreviousResponseIdOk

`func (o *AiResponsesResource) GetPreviousResponseIdOk() (*interface{}, bool)`

GetPreviousResponseIdOk returns a tuple with the PreviousResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousResponseId

`func (o *AiResponsesResource) SetPreviousResponseId(v interface{})`

SetPreviousResponseId sets PreviousResponseId field to given value.

### HasPreviousResponseId

`func (o *AiResponsesResource) HasPreviousResponseId() bool`

HasPreviousResponseId returns a boolean if a field has been set.

### SetPreviousResponseIdNil

`func (o *AiResponsesResource) SetPreviousResponseIdNil(b bool)`

 SetPreviousResponseIdNil sets the value for PreviousResponseId to be an explicit nil

### UnsetPreviousResponseId
`func (o *AiResponsesResource) UnsetPreviousResponseId()`

UnsetPreviousResponseId ensures that no value is present for PreviousResponseId, not even an explicit nil
### GetReasoning

`func (o *AiResponsesResource) GetReasoning() interface{}`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *AiResponsesResource) GetReasoningOk() (*interface{}, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *AiResponsesResource) SetReasoning(v interface{})`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *AiResponsesResource) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### SetReasoningNil

`func (o *AiResponsesResource) SetReasoningNil(b bool)`

 SetReasoningNil sets the value for Reasoning to be an explicit nil

### UnsetReasoning
`func (o *AiResponsesResource) UnsetReasoning()`

UnsetReasoning ensures that no value is present for Reasoning, not even an explicit nil
### GetStatus

`func (o *AiResponsesResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AiResponsesResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AiResponsesResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AiResponsesResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStore

`func (o *AiResponsesResource) GetStore() bool`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *AiResponsesResource) GetStoreOk() (*bool, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *AiResponsesResource) SetStore(v bool)`

SetStore sets Store field to given value.

### HasStore

`func (o *AiResponsesResource) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetTemperature

`func (o *AiResponsesResource) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *AiResponsesResource) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *AiResponsesResource) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *AiResponsesResource) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetText

`func (o *AiResponsesResource) GetText() interface{}`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AiResponsesResource) GetTextOk() (*interface{}, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AiResponsesResource) SetText(v interface{})`

SetText sets Text field to given value.

### HasText

`func (o *AiResponsesResource) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *AiResponsesResource) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *AiResponsesResource) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetToolChoice

`func (o *AiResponsesResource) GetToolChoice() interface{}`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *AiResponsesResource) GetToolChoiceOk() (*interface{}, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *AiResponsesResource) SetToolChoice(v interface{})`

SetToolChoice sets ToolChoice field to given value.

### HasToolChoice

`func (o *AiResponsesResource) HasToolChoice() bool`

HasToolChoice returns a boolean if a field has been set.

### SetToolChoiceNil

`func (o *AiResponsesResource) SetToolChoiceNil(b bool)`

 SetToolChoiceNil sets the value for ToolChoice to be an explicit nil

### UnsetToolChoice
`func (o *AiResponsesResource) UnsetToolChoice()`

UnsetToolChoice ensures that no value is present for ToolChoice, not even an explicit nil
### GetTools

`func (o *AiResponsesResource) GetTools() []AiResponsesTool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AiResponsesResource) GetToolsOk() (*[]AiResponsesTool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AiResponsesResource) SetTools(v []AiResponsesTool)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AiResponsesResource) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetTopP

`func (o *AiResponsesResource) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *AiResponsesResource) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *AiResponsesResource) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *AiResponsesResource) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### GetUsage

`func (o *AiResponsesResource) GetUsage() AiResponsesUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AiResponsesResource) GetUsageOk() (*AiResponsesUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AiResponsesResource) SetUsage(v AiResponsesUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *AiResponsesResource) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


