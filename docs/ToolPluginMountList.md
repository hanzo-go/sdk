# ToolPluginMountList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plugins** | Pointer to [**[]ToolPluginMount**](ToolPluginMount.md) | Plugins is every subsystem the composition root declared, filtered to the enabled ones unless all&#x3D;true. | [optional] 

## Methods

### NewToolPluginMountList

`func NewToolPluginMountList() *ToolPluginMountList`

NewToolPluginMountList instantiates a new ToolPluginMountList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolPluginMountListWithDefaults

`func NewToolPluginMountListWithDefaults() *ToolPluginMountList`

NewToolPluginMountListWithDefaults instantiates a new ToolPluginMountList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlugins

`func (o *ToolPluginMountList) GetPlugins() []ToolPluginMount`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *ToolPluginMountList) GetPluginsOk() (*[]ToolPluginMount, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *ToolPluginMountList) SetPlugins(v []ToolPluginMount)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *ToolPluginMountList) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


