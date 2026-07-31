# CloudHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addr** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** | Err is set when a peer could not be reached. Its plugins are then unknown, which is NOT the same as none, so the list stays empty and the drift below refuses to conclude anything from it. | [optional] 
**Host** | Pointer to **string** | Host is the pod&#39;s stable id, and Addr where it was reached. Self is true for the host that answered the request. | [optional] 
**Plugins** | Pointer to [**[]CloudStatus**](CloudStatus.md) |  | [optional] 
**Self** | Pointer to **bool** |  | [optional] 

## Methods

### NewCloudHost

`func NewCloudHost() *CloudHost`

NewCloudHost instantiates a new CloudHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudHostWithDefaults

`func NewCloudHostWithDefaults() *CloudHost`

NewCloudHostWithDefaults instantiates a new CloudHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddr

`func (o *CloudHost) GetAddr() string`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *CloudHost) GetAddrOk() (*string, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *CloudHost) SetAddr(v string)`

SetAddr sets Addr field to given value.

### HasAddr

`func (o *CloudHost) HasAddr() bool`

HasAddr returns a boolean if a field has been set.

### GetError

`func (o *CloudHost) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CloudHost) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CloudHost) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CloudHost) HasError() bool`

HasError returns a boolean if a field has been set.

### GetHost

`func (o *CloudHost) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CloudHost) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CloudHost) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CloudHost) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPlugins

`func (o *CloudHost) GetPlugins() []CloudStatus`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *CloudHost) GetPluginsOk() (*[]CloudStatus, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *CloudHost) SetPlugins(v []CloudStatus)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *CloudHost) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### GetSelf

`func (o *CloudHost) GetSelf() bool`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *CloudHost) GetSelfOk() (*bool, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *CloudHost) SetSelf(v bool)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *CloudHost) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


