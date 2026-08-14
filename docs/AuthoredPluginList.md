# AuthoredPluginList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plugins** | Pointer to [**[]AuthoredPlugin**](AuthoredPlugin.md) | Plugins is every plugin this org built, newest first, each carrying the TypeScript as authored. The bundled artifact is never rendered. | [optional] 

## Methods

### NewAuthoredPluginList

`func NewAuthoredPluginList() *AuthoredPluginList`

NewAuthoredPluginList instantiates a new AuthoredPluginList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthoredPluginListWithDefaults

`func NewAuthoredPluginListWithDefaults() *AuthoredPluginList`

NewAuthoredPluginListWithDefaults instantiates a new AuthoredPluginList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlugins

`func (o *AuthoredPluginList) GetPlugins() []AuthoredPlugin`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *AuthoredPluginList) GetPluginsOk() (*[]AuthoredPlugin, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *AuthoredPluginList) SetPlugins(v []AuthoredPlugin)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *AuthoredPluginList) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


