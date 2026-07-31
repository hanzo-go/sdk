# PluginNameIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the app, from the path. | [optional] 
**Scope** | Pointer to **string** | Scope \&quot;host\&quot; applies here only; default \&quot;fleet\&quot; applies everywhere.  | [optional] 

## Methods

### NewPluginNameIn

`func NewPluginNameIn() *PluginNameIn`

NewPluginNameIn instantiates a new PluginNameIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginNameInWithDefaults

`func NewPluginNameInWithDefaults() *PluginNameIn`

NewPluginNameInWithDefaults instantiates a new PluginNameIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PluginNameIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PluginNameIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PluginNameIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PluginNameIn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScope

`func (o *PluginNameIn) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PluginNameIn) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PluginNameIn) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *PluginNameIn) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


