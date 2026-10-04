# ProviderDevicePollOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | Pointer to [**ProviderConnectionView**](ProviderConnectionView.md) | Connection is the connected connector. Present only on \&quot;connected\&quot;. | [optional] 
**Interval** | Pointer to **int64** | Interval is the seconds to wait before the next poll. Present only while pending, and it may rise when the provider asks the client to slow down. | [optional] 
**Status** | Pointer to **string** | Status is the flow&#39;s state. \&quot;pending\&quot; means poll again after Interval. | [optional] 

## Methods

### NewProviderDevicePollOut

`func NewProviderDevicePollOut() *ProviderDevicePollOut`

NewProviderDevicePollOut instantiates a new ProviderDevicePollOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderDevicePollOutWithDefaults

`func NewProviderDevicePollOutWithDefaults() *ProviderDevicePollOut`

NewProviderDevicePollOutWithDefaults instantiates a new ProviderDevicePollOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *ProviderDevicePollOut) GetConnector() ProviderConnectionView`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ProviderDevicePollOut) GetConnectorOk() (*ProviderConnectionView, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ProviderDevicePollOut) SetConnector(v ProviderConnectionView)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *ProviderDevicePollOut) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetInterval

`func (o *ProviderDevicePollOut) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ProviderDevicePollOut) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ProviderDevicePollOut) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *ProviderDevicePollOut) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetStatus

`func (o *ProviderDevicePollOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProviderDevicePollOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProviderDevicePollOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProviderDevicePollOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


