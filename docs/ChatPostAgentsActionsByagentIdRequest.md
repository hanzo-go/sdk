# ChatPostAgentsActionsByagentIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Functions** | [**[]ChatFunctionTool**](ChatFunctionTool.md) |  | 
**ActionId** | Pointer to **string** |  | [optional] 
**Metadata** | [**ChatActionMetadata**](ChatActionMetadata.md) |  | 

## Methods

### NewChatPostAgentsActionsByagentIdRequest

`func NewChatPostAgentsActionsByagentIdRequest(functions []ChatFunctionTool, metadata ChatActionMetadata, ) *ChatPostAgentsActionsByagentIdRequest`

NewChatPostAgentsActionsByagentIdRequest instantiates a new ChatPostAgentsActionsByagentIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatPostAgentsActionsByagentIdRequestWithDefaults

`func NewChatPostAgentsActionsByagentIdRequestWithDefaults() *ChatPostAgentsActionsByagentIdRequest`

NewChatPostAgentsActionsByagentIdRequestWithDefaults instantiates a new ChatPostAgentsActionsByagentIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctions

`func (o *ChatPostAgentsActionsByagentIdRequest) GetFunctions() []ChatFunctionTool`

GetFunctions returns the Functions field if non-nil, zero value otherwise.

### GetFunctionsOk

`func (o *ChatPostAgentsActionsByagentIdRequest) GetFunctionsOk() (*[]ChatFunctionTool, bool)`

GetFunctionsOk returns a tuple with the Functions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctions

`func (o *ChatPostAgentsActionsByagentIdRequest) SetFunctions(v []ChatFunctionTool)`

SetFunctions sets Functions field to given value.


### GetActionId

`func (o *ChatPostAgentsActionsByagentIdRequest) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *ChatPostAgentsActionsByagentIdRequest) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *ChatPostAgentsActionsByagentIdRequest) SetActionId(v string)`

SetActionId sets ActionId field to given value.

### HasActionId

`func (o *ChatPostAgentsActionsByagentIdRequest) HasActionId() bool`

HasActionId returns a boolean if a field has been set.

### GetMetadata

`func (o *ChatPostAgentsActionsByagentIdRequest) GetMetadata() ChatActionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ChatPostAgentsActionsByagentIdRequest) GetMetadataOk() (*ChatActionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ChatPostAgentsActionsByagentIdRequest) SetMetadata(v ChatActionMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


