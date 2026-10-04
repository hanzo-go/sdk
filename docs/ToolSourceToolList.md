# ToolSourceToolList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** | Source is the source these tools came from. | [optional] 
**Tools** | Pointer to [**[]ToolTool**](ToolTool.md) | Tools is the caller&#39;s tools from that source. Never null. | [optional] 

## Methods

### NewToolSourceToolList

`func NewToolSourceToolList() *ToolSourceToolList`

NewToolSourceToolList instantiates a new ToolSourceToolList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolSourceToolListWithDefaults

`func NewToolSourceToolListWithDefaults() *ToolSourceToolList`

NewToolSourceToolListWithDefaults instantiates a new ToolSourceToolList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *ToolSourceToolList) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ToolSourceToolList) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ToolSourceToolList) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ToolSourceToolList) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTools

`func (o *ToolSourceToolList) GetTools() []ToolTool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ToolSourceToolList) GetToolsOk() (*[]ToolTool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ToolSourceToolList) SetTools(v []ToolTool)`

SetTools sets Tools field to given value.

### HasTools

`func (o *ToolSourceToolList) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


