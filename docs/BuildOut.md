# BuildOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | Pointer to **int64** | Bytes is the size of the bundled CommonJS the runtime will execute. | [optional] 
**Generated** | Pointer to **bool** | Generated is whether a model wrote the source from a spec, rather than the caller posting the source itself. | [optional] 
**Plugin** | Pointer to [**AuthoredPlugin**](AuthoredPlugin.md) | Plugin is the plugin as stored, with its derived id and build time. | [optional] 

## Methods

### NewBuildOut

`func NewBuildOut() *BuildOut`

NewBuildOut instantiates a new BuildOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildOutWithDefaults

`func NewBuildOutWithDefaults() *BuildOut`

NewBuildOutWithDefaults instantiates a new BuildOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *BuildOut) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *BuildOut) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *BuildOut) SetBytes(v int64)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *BuildOut) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetGenerated

`func (o *BuildOut) GetGenerated() bool`

GetGenerated returns the Generated field if non-nil, zero value otherwise.

### GetGeneratedOk

`func (o *BuildOut) GetGeneratedOk() (*bool, bool)`

GetGeneratedOk returns a tuple with the Generated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerated

`func (o *BuildOut) SetGenerated(v bool)`

SetGenerated sets Generated field to given value.

### HasGenerated

`func (o *BuildOut) HasGenerated() bool`

HasGenerated returns a boolean if a field has been set.

### GetPlugin

`func (o *BuildOut) GetPlugin() AuthoredPlugin`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *BuildOut) GetPluginOk() (*AuthoredPlugin, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *BuildOut) SetPlugin(v AuthoredPlugin)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *BuildOut) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


