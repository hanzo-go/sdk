# PluginHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | Host is the pod&#39;s stable id, and Addr where it was reached. Self is true for the host that answered the request.  | [optional] 
**Addr** | Pointer to **string** |  | [optional] 
**Self** | Pointer to **bool** |  | [optional] 
**Error** | Pointer to **string** | Err is set when a peer could not be reached. Its plugins are then unknown, which is NOT the same as none, so the list stays empty and the drift below refuses to conclude anything from it.  | [optional] 
**Plugins** | Pointer to [**[]PluginPluginStatus**](PluginPluginStatus.md) |  | [optional] 

## Methods

### NewPluginHost

`func NewPluginHost() *PluginHost`

NewPluginHost instantiates a new PluginHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginHostWithDefaults

`func NewPluginHostWithDefaults() *PluginHost`

NewPluginHostWithDefaults instantiates a new PluginHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *PluginHost) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PluginHost) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PluginHost) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *PluginHost) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetAddr

`func (o *PluginHost) GetAddr() string`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *PluginHost) GetAddrOk() (*string, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *PluginHost) SetAddr(v string)`

SetAddr sets Addr field to given value.

### HasAddr

`func (o *PluginHost) HasAddr() bool`

HasAddr returns a boolean if a field has been set.

### GetSelf

`func (o *PluginHost) GetSelf() bool`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *PluginHost) GetSelfOk() (*bool, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *PluginHost) SetSelf(v bool)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *PluginHost) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetError

`func (o *PluginHost) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PluginHost) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PluginHost) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *PluginHost) HasError() bool`

HasError returns a boolean if a field has been set.

### GetPlugins

`func (o *PluginHost) GetPlugins() []PluginPluginStatus`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *PluginHost) GetPluginsOk() (*[]PluginPluginStatus, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *PluginHost) SetPlugins(v []PluginPluginStatus)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *PluginHost) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


