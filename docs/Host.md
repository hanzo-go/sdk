# Host

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addr** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** | Err is set when a peer could not be reached. Its plugins are then unknown, which is NOT the same as none, so the list stays empty and the drift below refuses to conclude anything from it. | [optional] 
**Host** | Pointer to **string** | Host is the pod&#39;s stable id, and Addr where it was reached. Self is true for the host that answered the request. | [optional] 
**Plugins** | Pointer to [**[]Status**](Status.md) |  | [optional] 
**Self** | Pointer to **bool** |  | [optional] 

## Methods

### NewHost

`func NewHost() *Host`

NewHost instantiates a new Host object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostWithDefaults

`func NewHostWithDefaults() *Host`

NewHostWithDefaults instantiates a new Host object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddr

`func (o *Host) GetAddr() string`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *Host) GetAddrOk() (*string, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *Host) SetAddr(v string)`

SetAddr sets Addr field to given value.

### HasAddr

`func (o *Host) HasAddr() bool`

HasAddr returns a boolean if a field has been set.

### GetError

`func (o *Host) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Host) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Host) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Host) HasError() bool`

HasError returns a boolean if a field has been set.

### GetHost

`func (o *Host) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Host) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Host) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *Host) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPlugins

`func (o *Host) GetPlugins() []Status`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *Host) GetPluginsOk() (*[]Status, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *Host) SetPlugins(v []Status)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *Host) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### GetSelf

`func (o *Host) GetSelf() bool`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *Host) GetSelfOk() (*bool, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *Host) SetSelf(v bool)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *Host) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


