# NotifyHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** | Service names the subsystem answering — always \&quot;notify\&quot;. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot;; the route answers 200 whenever the subsystem is mounted. | [optional] 

## Methods

### NewNotifyHealth

`func NewNotifyHealth() *NotifyHealth`

NewNotifyHealth instantiates a new NotifyHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyHealthWithDefaults

`func NewNotifyHealthWithDefaults() *NotifyHealth`

NewNotifyHealthWithDefaults instantiates a new NotifyHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *NotifyHealth) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *NotifyHealth) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *NotifyHealth) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *NotifyHealth) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *NotifyHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotifyHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotifyHealth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NotifyHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


