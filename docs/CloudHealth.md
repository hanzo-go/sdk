# CloudHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status is ok when the message plane answers, degraded otherwise. | [optional] 
**Uptime** | Pointer to **string** | Uptime is how long this surface has been mounted. | [optional] 
**Version** | Pointer to **string** | Version is the connected broker&#39;s server version; empty while degraded. | [optional] 

## Methods

### NewCloudHealth

`func NewCloudHealth() *CloudHealth`

NewCloudHealth instantiates a new CloudHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudHealthWithDefaults

`func NewCloudHealthWithDefaults() *CloudHealth`

NewCloudHealthWithDefaults instantiates a new CloudHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CloudHealth) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudHealth) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudHealth) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudHealth) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUptime

`func (o *CloudHealth) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *CloudHealth) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *CloudHealth) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *CloudHealth) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetVersion

`func (o *CloudHealth) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CloudHealth) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CloudHealth) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CloudHealth) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


