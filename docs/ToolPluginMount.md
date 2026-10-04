# ToolPluginMount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enabled is whether this subsystem is switched on in this deployment. | [optional] 
**Name** | Pointer to **string** | Name is the subsystem&#39;s name, the same label a traced request resolves to. | [optional] 
**Prefixes** | Pointer to **[]string** | Prefixes are the URL prefixes this subsystem serves. | [optional] 

## Methods

### NewToolPluginMount

`func NewToolPluginMount() *ToolPluginMount`

NewToolPluginMount instantiates a new ToolPluginMount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolPluginMountWithDefaults

`func NewToolPluginMountWithDefaults() *ToolPluginMount`

NewToolPluginMountWithDefaults instantiates a new ToolPluginMount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ToolPluginMount) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ToolPluginMount) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ToolPluginMount) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ToolPluginMount) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *ToolPluginMount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolPluginMount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolPluginMount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ToolPluginMount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrefixes

`func (o *ToolPluginMount) GetPrefixes() []string`

GetPrefixes returns the Prefixes field if non-nil, zero value otherwise.

### GetPrefixesOk

`func (o *ToolPluginMount) GetPrefixesOk() (*[]string, bool)`

GetPrefixesOk returns a tuple with the Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixes

`func (o *ToolPluginMount) SetPrefixes(v []string)`

SetPrefixes sets Prefixes field to given value.

### HasPrefixes

`func (o *ToolPluginMount) HasPrefixes() bool`

HasPrefixes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


