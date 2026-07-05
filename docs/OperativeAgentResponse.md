# OperativeAgentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]OperativeMessage**](OperativeMessage.md) |  | [optional] 
**ToolResults** | Pointer to [**[]OperativeToolResult**](OperativeToolResult.md) |  | [optional] 

## Methods

### NewOperativeAgentResponse

`func NewOperativeAgentResponse() *OperativeAgentResponse`

NewOperativeAgentResponse instantiates a new OperativeAgentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperativeAgentResponseWithDefaults

`func NewOperativeAgentResponseWithDefaults() *OperativeAgentResponse`

NewOperativeAgentResponseWithDefaults instantiates a new OperativeAgentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *OperativeAgentResponse) GetMessages() []OperativeMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *OperativeAgentResponse) GetMessagesOk() (*[]OperativeMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *OperativeAgentResponse) SetMessages(v []OperativeMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *OperativeAgentResponse) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetToolResults

`func (o *OperativeAgentResponse) GetToolResults() []OperativeToolResult`

GetToolResults returns the ToolResults field if non-nil, zero value otherwise.

### GetToolResultsOk

`func (o *OperativeAgentResponse) GetToolResultsOk() (*[]OperativeToolResult, bool)`

GetToolResultsOk returns a tuple with the ToolResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolResults

`func (o *OperativeAgentResponse) SetToolResults(v []OperativeToolResult)`

SetToolResults sets ToolResults field to given value.

### HasToolResults

`func (o *OperativeAgentResponse) HasToolResults() bool`

HasToolResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


