# ToolBuildOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | Pointer to **int64** | Bytes is the size of the bundled CommonJS the runtime will execute. | [optional] 
**Generated** | Pointer to **bool** | Generated is whether a model wrote the source from a spec, rather than the caller posting the source itself. | [optional] 
**Plugin** | Pointer to [**ToolAuthoredPlugin**](ToolAuthoredPlugin.md) | Plugin is the plugin as stored, with its derived id and build time. | [optional] 

## Methods

### NewToolBuildOut

`func NewToolBuildOut() *ToolBuildOut`

NewToolBuildOut instantiates a new ToolBuildOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolBuildOutWithDefaults

`func NewToolBuildOutWithDefaults() *ToolBuildOut`

NewToolBuildOutWithDefaults instantiates a new ToolBuildOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *ToolBuildOut) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *ToolBuildOut) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *ToolBuildOut) SetBytes(v int64)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *ToolBuildOut) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetGenerated

`func (o *ToolBuildOut) GetGenerated() bool`

GetGenerated returns the Generated field if non-nil, zero value otherwise.

### GetGeneratedOk

`func (o *ToolBuildOut) GetGeneratedOk() (*bool, bool)`

GetGeneratedOk returns a tuple with the Generated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerated

`func (o *ToolBuildOut) SetGenerated(v bool)`

SetGenerated sets Generated field to given value.

### HasGenerated

`func (o *ToolBuildOut) HasGenerated() bool`

HasGenerated returns a boolean if a field has been set.

### GetPlugin

`func (o *ToolBuildOut) GetPlugin() ToolAuthoredPlugin`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *ToolBuildOut) GetPluginOk() (*ToolAuthoredPlugin, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *ToolBuildOut) SetPlugin(v ToolAuthoredPlugin)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *ToolBuildOut) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


