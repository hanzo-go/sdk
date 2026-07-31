# CloudDevicePollOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | Pointer to [**CloudConnView**](CloudConnView.md) | Connector is the connected connector. Present only on \&quot;connected\&quot;. | [optional] 
**Interval** | Pointer to **int32** | Interval is the seconds to wait before the next poll. Present only while pending, and it may rise when the provider asks the client to slow down. | [optional] 
**Status** | Pointer to **string** | Status is the flow&#39;s state. \&quot;pending\&quot; means poll again after Interval. | [optional] 

## Methods

### NewCloudDevicePollOut

`func NewCloudDevicePollOut() *CloudDevicePollOut`

NewCloudDevicePollOut instantiates a new CloudDevicePollOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDevicePollOutWithDefaults

`func NewCloudDevicePollOutWithDefaults() *CloudDevicePollOut`

NewCloudDevicePollOutWithDefaults instantiates a new CloudDevicePollOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *CloudDevicePollOut) GetConnector() CloudConnView`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *CloudDevicePollOut) GetConnectorOk() (*CloudConnView, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *CloudDevicePollOut) SetConnector(v CloudConnView)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *CloudDevicePollOut) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetInterval

`func (o *CloudDevicePollOut) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CloudDevicePollOut) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CloudDevicePollOut) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CloudDevicePollOut) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetStatus

`func (o *CloudDevicePollOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudDevicePollOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudDevicePollOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudDevicePollOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


