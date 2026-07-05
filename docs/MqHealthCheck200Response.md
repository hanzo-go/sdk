# MqHealthCheck200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Uptime** | Pointer to **string** | Human-readable uptime duration. | [optional] 

## Methods

### NewMqHealthCheck200Response

`func NewMqHealthCheck200Response() *MqHealthCheck200Response`

NewMqHealthCheck200Response instantiates a new MqHealthCheck200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMqHealthCheck200ResponseWithDefaults

`func NewMqHealthCheck200ResponseWithDefaults() *MqHealthCheck200Response`

NewMqHealthCheck200ResponseWithDefaults instantiates a new MqHealthCheck200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MqHealthCheck200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MqHealthCheck200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MqHealthCheck200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MqHealthCheck200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *MqHealthCheck200Response) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MqHealthCheck200Response) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MqHealthCheck200Response) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MqHealthCheck200Response) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUptime

`func (o *MqHealthCheck200Response) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *MqHealthCheck200Response) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *MqHealthCheck200Response) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *MqHealthCheck200Response) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


