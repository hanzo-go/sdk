# PluginPluginInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Plugin name from the manifest entry. | [optional] 
**Kind** | Pointer to **string** | Plugin kind selecting how it was built and mounted. | [optional] 
**Prefix** | Pointer to **string** | Mount point at which the plugin&#39;s routes are served. | [optional] 

## Methods

### NewPluginPluginInfo

`func NewPluginPluginInfo() *PluginPluginInfo`

NewPluginPluginInfo instantiates a new PluginPluginInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginPluginInfoWithDefaults

`func NewPluginPluginInfoWithDefaults() *PluginPluginInfo`

NewPluginPluginInfoWithDefaults instantiates a new PluginPluginInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PluginPluginInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PluginPluginInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PluginPluginInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PluginPluginInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKind

`func (o *PluginPluginInfo) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PluginPluginInfo) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PluginPluginInfo) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PluginPluginInfo) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPrefix

`func (o *PluginPluginInfo) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PluginPluginInfo) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PluginPluginInfo) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *PluginPluginInfo) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


