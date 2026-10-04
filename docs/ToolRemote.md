# ToolRemote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the server&#39;s id in the org, which also prefixes its tools on the tool plane, and names it to ToolsRelay. | [optional] 
**Tools** | Pointer to **[]string** | Tools are the tools of it an admin activated, each by the server&#39;s own name for it: all a run&#39;s harness may call of it. | [optional] 

## Methods

### NewToolRemote

`func NewToolRemote() *ToolRemote`

NewToolRemote instantiates a new ToolRemote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolRemoteWithDefaults

`func NewToolRemoteWithDefaults() *ToolRemote`

NewToolRemoteWithDefaults instantiates a new ToolRemote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ToolRemote) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolRemote) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolRemote) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ToolRemote) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTools

`func (o *ToolRemote) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ToolRemote) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ToolRemote) SetTools(v []string)`

SetTools sets Tools field to given value.

### HasTools

`func (o *ToolRemote) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


