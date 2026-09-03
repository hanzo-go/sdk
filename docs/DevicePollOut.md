# DevicePollOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | Pointer to [**ConnView**](ConnView.md) | Connection is the connected connector. Present only on \&quot;connected\&quot;. | [optional] 
**Interval** | Pointer to **int64** | Interval is the seconds to wait before the next poll. Present only while pending, and it may rise when the provider asks the client to slow down. | [optional] 
**Status** | Pointer to **string** | Status is the flow&#39;s state. \&quot;pending\&quot; means poll again after Interval. | [optional] 

## Methods

### NewDevicePollOut

`func NewDevicePollOut() *DevicePollOut`

NewDevicePollOut instantiates a new DevicePollOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicePollOutWithDefaults

`func NewDevicePollOutWithDefaults() *DevicePollOut`

NewDevicePollOutWithDefaults instantiates a new DevicePollOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *DevicePollOut) GetConnector() ConnView`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *DevicePollOut) GetConnectorOk() (*ConnView, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *DevicePollOut) SetConnector(v ConnView)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *DevicePollOut) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetInterval

`func (o *DevicePollOut) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *DevicePollOut) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *DevicePollOut) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *DevicePollOut) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetStatus

`func (o *DevicePollOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DevicePollOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DevicePollOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DevicePollOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


