# PluginMountList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plugins** | Pointer to [**[]PluginMount**](PluginMount.md) | Plugins is every subsystem the composition root declared, filtered to the enabled ones unless all&#x3D;true. | [optional] 

## Methods

### NewPluginMountList

`func NewPluginMountList() *PluginMountList`

NewPluginMountList instantiates a new PluginMountList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginMountListWithDefaults

`func NewPluginMountListWithDefaults() *PluginMountList`

NewPluginMountListWithDefaults instantiates a new PluginMountList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlugins

`func (o *PluginMountList) GetPlugins() []PluginMount`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *PluginMountList) GetPluginsOk() (*[]PluginMount, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *PluginMountList) SetPlugins(v []PluginMount)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *PluginMountList) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


