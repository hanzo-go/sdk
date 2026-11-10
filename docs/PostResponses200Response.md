# PostResponses200Response

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
**Routing** | Pointer to [**PostChat200ResponseAllOfRouting**](PostChat200ResponseAllOfRouting.md) |  | [optional] 

## Methods

### NewPostResponses200Response

`func NewPostResponses200Response() *PostResponses200Response`

NewPostResponses200Response instantiates a new PostResponses200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostResponses200ResponseWithDefaults

`func NewPostResponses200ResponseWithDefaults() *PostResponses200Response`

NewPostResponses200ResponseWithDefaults instantiates a new PostResponses200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PostResponses200Response) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PostResponses200Response) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PostResponses200Response) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PostResponses200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetError

`func (o *PostResponses200Response) GetError() interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PostResponses200Response) GetErrorOk() (*interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PostResponses200Response) SetError(v interface{})`

SetError sets Error field to given value.

### HasError

`func (o *PostResponses200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *PostResponses200Response) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *PostResponses200Response) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetId

`func (o *PostResponses200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostResponses200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostResponses200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PostResponses200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncompleteDetails

`func (o *PostResponses200Response) GetIncompleteDetails() interface{}`

GetIncompleteDetails returns the IncompleteDetails field if non-nil, zero value otherwise.

### GetIncompleteDetailsOk

`func (o *PostResponses200Response) GetIncompleteDetailsOk() (*interface{}, bool)`

GetIncompleteDetailsOk returns a tuple with the IncompleteDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompleteDetails

`func (o *PostResponses200Response) SetIncompleteDetails(v interface{})`

SetIncompleteDetails sets IncompleteDetails field to given value.

### HasIncompleteDetails

`func (o *PostResponses200Response) HasIncompleteDetails() bool`

HasIncompleteDetails returns a boolean if a field has been set.

### SetIncompleteDetailsNil

`func (o *PostResponses200Response) SetIncompleteDetailsNil(b bool)`

 SetIncompleteDetailsNil sets the value for IncompleteDetails to be an explicit nil

### UnsetIncompleteDetails
`func (o *PostResponses200Response) UnsetIncompleteDetails()`

UnsetIncompleteDetails ensures that no value is present for IncompleteDetails, not even an explicit nil
### GetInstructions

`func (o *PostResponses200Response) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *PostResponses200Response) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *PostResponses200Response) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *PostResponses200Response) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetMaxOutputTokens

`func (o *PostResponses200Response) GetMaxOutputTokens() int32`

GetMaxOutputTokens returns the MaxOutputTokens field if non-nil, zero value otherwise.

### GetMaxOutputTokensOk

`func (o *PostResponses200Response) GetMaxOutputTokensOk() (*int32, bool)`

GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOutputTokens

`func (o *PostResponses200Response) SetMaxOutputTokens(v int32)`

SetMaxOutputTokens sets MaxOutputTokens field to given value.

### HasMaxOutputTokens

`func (o *PostResponses200Response) HasMaxOutputTokens() bool`

HasMaxOutputTokens returns a boolean if a field has been set.

### GetMetadata

`func (o *PostResponses200Response) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PostResponses200Response) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PostResponses200Response) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PostResponses200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetModel

`func (o *PostResponses200Response) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PostResponses200Response) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PostResponses200Response) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *PostResponses200Response) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetObject

`func (o *PostResponses200Response) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PostResponses200Response) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PostResponses200Response) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *PostResponses200Response) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetOutput

`func (o *PostResponses200Response) GetOutput() []interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *PostResponses200Response) GetOutputOk() (*[]interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *PostResponses200Response) SetOutput(v []interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *PostResponses200Response) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetParallelToolCalls

`func (o *PostResponses200Response) GetParallelToolCalls() bool`

GetParallelToolCalls returns the ParallelToolCalls field if non-nil, zero value otherwise.

### GetParallelToolCallsOk

`func (o *PostResponses200Response) GetParallelToolCallsOk() (*bool, bool)`

GetParallelToolCallsOk returns a tuple with the ParallelToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelToolCalls

`func (o *PostResponses200Response) SetParallelToolCalls(v bool)`

SetParallelToolCalls sets ParallelToolCalls field to given value.

### HasParallelToolCalls

`func (o *PostResponses200Response) HasParallelToolCalls() bool`

HasParallelToolCalls returns a boolean if a field has been set.

### GetPreviousResponseId

`func (o *PostResponses200Response) GetPreviousResponseId() interface{}`

GetPreviousResponseId returns the PreviousResponseId field if non-nil, zero value otherwise.

### GetPreviousResponseIdOk

`func (o *PostResponses200Response) GetPreviousResponseIdOk() (*interface{}, bool)`

GetPreviousResponseIdOk returns a tuple with the PreviousResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousResponseId

`func (o *PostResponses200Response) SetPreviousResponseId(v interface{})`

SetPreviousResponseId sets PreviousResponseId field to given value.

### HasPreviousResponseId

`func (o *PostResponses200Response) HasPreviousResponseId() bool`

HasPreviousResponseId returns a boolean if a field has been set.

### SetPreviousResponseIdNil

`func (o *PostResponses200Response) SetPreviousResponseIdNil(b bool)`

 SetPreviousResponseIdNil sets the value for PreviousResponseId to be an explicit nil

### UnsetPreviousResponseId
`func (o *PostResponses200Response) UnsetPreviousResponseId()`

UnsetPreviousResponseId ensures that no value is present for PreviousResponseId, not even an explicit nil
### GetReasoning

`func (o *PostResponses200Response) GetReasoning() interface{}`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *PostResponses200Response) GetReasoningOk() (*interface{}, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *PostResponses200Response) SetReasoning(v interface{})`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *PostResponses200Response) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### SetReasoningNil

`func (o *PostResponses200Response) SetReasoningNil(b bool)`

 SetReasoningNil sets the value for Reasoning to be an explicit nil

### UnsetReasoning
`func (o *PostResponses200Response) UnsetReasoning()`

UnsetReasoning ensures that no value is present for Reasoning, not even an explicit nil
### GetStatus

`func (o *PostResponses200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostResponses200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostResponses200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostResponses200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStore

`func (o *PostResponses200Response) GetStore() bool`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *PostResponses200Response) GetStoreOk() (*bool, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *PostResponses200Response) SetStore(v bool)`

SetStore sets Store field to given value.

### HasStore

`func (o *PostResponses200Response) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetTemperature

`func (o *PostResponses200Response) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *PostResponses200Response) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *PostResponses200Response) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *PostResponses200Response) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetText

`func (o *PostResponses200Response) GetText() interface{}`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostResponses200Response) GetTextOk() (*interface{}, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostResponses200Response) SetText(v interface{})`

SetText sets Text field to given value.

### HasText

`func (o *PostResponses200Response) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *PostResponses200Response) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *PostResponses200Response) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetToolChoice

`func (o *PostResponses200Response) GetToolChoice() interface{}`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *PostResponses200Response) GetToolChoiceOk() (*interface{}, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *PostResponses200Response) SetToolChoice(v interface{})`

SetToolChoice sets ToolChoice field to given value.

### HasToolChoice

`func (o *PostResponses200Response) HasToolChoice() bool`

HasToolChoice returns a boolean if a field has been set.

### SetToolChoiceNil

`func (o *PostResponses200Response) SetToolChoiceNil(b bool)`

 SetToolChoiceNil sets the value for ToolChoice to be an explicit nil

### UnsetToolChoice
`func (o *PostResponses200Response) UnsetToolChoice()`

UnsetToolChoice ensures that no value is present for ToolChoice, not even an explicit nil
### GetTools

`func (o *PostResponses200Response) GetTools() []AiResponsesTool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *PostResponses200Response) GetToolsOk() (*[]AiResponsesTool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *PostResponses200Response) SetTools(v []AiResponsesTool)`

SetTools sets Tools field to given value.

### HasTools

`func (o *PostResponses200Response) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetTopP

`func (o *PostResponses200Response) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *PostResponses200Response) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *PostResponses200Response) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *PostResponses200Response) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### GetUsage

`func (o *PostResponses200Response) GetUsage() AiResponsesUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *PostResponses200Response) GetUsageOk() (*AiResponsesUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *PostResponses200Response) SetUsage(v AiResponsesUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *PostResponses200Response) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetRouting

`func (o *PostResponses200Response) GetRouting() PostChat200ResponseAllOfRouting`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *PostResponses200Response) GetRoutingOk() (*PostChat200ResponseAllOfRouting, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *PostResponses200Response) SetRouting(v PostChat200ResponseAllOfRouting)`

SetRouting sets Routing field to given value.

### HasRouting

`func (o *PostResponses200Response) HasRouting() bool`

HasRouting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


