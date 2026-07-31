# CloudNotifyHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** | Service names the subsystem answering — always \&quot;notify\&quot;. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;ok\&quot;; the route answers 200 whenever the subsystem is mounted. | [optional] 

## Methods

### NewCloudNotifyHealth

`func NewCloudNotifyHealth() *CloudNotifyHealth`

NewCloudNotifyHealth instantiates a new CloudNotifyHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudNotifyHealthWithDefaults

`func NewCloudNotifyHealthWithDefaults() *CloudNotifyHealth`

NewCloudNotifyHealthWithDefaults instantiates a new CloudNotifyHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *CloudNotifyHealth) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CloudNotifyHealth) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CloudNotifyHealth) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *CloudNotifyHealth) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *CloudNotifyHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudNotifyHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudNotifyHealth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudNotifyHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


