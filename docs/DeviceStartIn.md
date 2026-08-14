# DeviceStartIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Label names this connection so one user can hold several per provider (\&quot;work\&quot;, \&quot;personal\&quot;). Empty means \&quot;default\&quot;. 1-64 of [A-Za-z0-9._-]. | [optional] 
**Provider** | Pointer to **string** | Provider is the user-scoped provider&#39;s registry id, from the path. | [optional] 

## Methods

### NewDeviceStartIn

`func NewDeviceStartIn() *DeviceStartIn`

NewDeviceStartIn instantiates a new DeviceStartIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceStartInWithDefaults

`func NewDeviceStartInWithDefaults() *DeviceStartIn`

NewDeviceStartInWithDefaults instantiates a new DeviceStartIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *DeviceStartIn) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DeviceStartIn) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DeviceStartIn) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DeviceStartIn) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetProvider

`func (o *DeviceStartIn) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DeviceStartIn) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DeviceStartIn) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *DeviceStartIn) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


