# ToolToolList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tools** | Pointer to [**[]ToolTool**](ToolTool.md) | Tools is every tool the caller may see, deduplicated by name with source precedence applied. | [optional] 

## Methods

### NewToolToolList

`func NewToolToolList() *ToolToolList`

NewToolToolList instantiates a new ToolToolList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolToolListWithDefaults

`func NewToolToolListWithDefaults() *ToolToolList`

NewToolToolListWithDefaults instantiates a new ToolToolList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTools

`func (o *ToolToolList) GetTools() []ToolTool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ToolToolList) GetToolsOk() (*[]ToolTool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ToolToolList) SetTools(v []ToolTool)`

SetTools sets Tools field to given value.

### HasTools

`func (o *ToolToolList) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


