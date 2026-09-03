# DeviceStartOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **string** | ExpiresAt is when the flow dies, RFC 3339 UTC. | [optional] 
**Flow** | Pointer to **string** | Flow is the id to poll with. | [optional] 
**Interval** | Pointer to **int64** | Interval is the seconds to wait between polls. | [optional] 
**UserCode** | Pointer to **string** | UserCode is the short code the user types at VerifyURL. | [optional] 
**VerifyUrl** | Pointer to **string** | VerifyURL is the page the user opens to enter UserCode. | [optional] 

## Methods

### NewDeviceStartOut

`func NewDeviceStartOut() *DeviceStartOut`

NewDeviceStartOut instantiates a new DeviceStartOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceStartOutWithDefaults

`func NewDeviceStartOutWithDefaults() *DeviceStartOut`

NewDeviceStartOutWithDefaults instantiates a new DeviceStartOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *DeviceStartOut) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *DeviceStartOut) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *DeviceStartOut) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *DeviceStartOut) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFlow

`func (o *DeviceStartOut) GetFlow() string`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *DeviceStartOut) GetFlowOk() (*string, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *DeviceStartOut) SetFlow(v string)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *DeviceStartOut) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetInterval

`func (o *DeviceStartOut) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *DeviceStartOut) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *DeviceStartOut) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *DeviceStartOut) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetUserCode

`func (o *DeviceStartOut) GetUserCode() string`

GetUserCode returns the UserCode field if non-nil, zero value otherwise.

### GetUserCodeOk

`func (o *DeviceStartOut) GetUserCodeOk() (*string, bool)`

GetUserCodeOk returns a tuple with the UserCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCode

`func (o *DeviceStartOut) SetUserCode(v string)`

SetUserCode sets UserCode field to given value.

### HasUserCode

`func (o *DeviceStartOut) HasUserCode() bool`

HasUserCode returns a boolean if a field has been set.

### GetVerifyUrl

`func (o *DeviceStartOut) GetVerifyUrl() string`

GetVerifyUrl returns the VerifyUrl field if non-nil, zero value otherwise.

### GetVerifyUrlOk

`func (o *DeviceStartOut) GetVerifyUrlOk() (*string, bool)`

GetVerifyUrlOk returns a tuple with the VerifyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyUrl

`func (o *DeviceStartOut) SetVerifyUrl(v string)`

SetVerifyUrl sets VerifyUrl field to given value.

### HasVerifyUrl

`func (o *DeviceStartOut) HasVerifyUrl() bool`

HasVerifyUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


