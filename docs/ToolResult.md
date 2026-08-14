# ToolResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the tool that ran. | [optional] 
**Result** | Pointer to **map[string]interface{}** | Result is the tool&#39;s own output, verbatim — its shape is the tool&#39;s, not this plane&#39;s. | [optional] 

## Methods

### NewToolResult

`func NewToolResult() *ToolResult`

NewToolResult instantiates a new ToolResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolResultWithDefaults

`func NewToolResultWithDefaults() *ToolResult`

NewToolResultWithDefaults instantiates a new ToolResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ToolResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ToolResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResult

`func (o *ToolResult) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ToolResult) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ToolResult) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *ToolResult) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


