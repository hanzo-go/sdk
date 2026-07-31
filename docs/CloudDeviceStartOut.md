# CloudDeviceStartOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **string** | ExpiresAt is when the flow dies, RFC 3339 UTC. | [optional] 
**Flow** | Pointer to **string** | Flow is the id to poll with. | [optional] 
**Interval** | Pointer to **int32** | Interval is the seconds to wait between polls. | [optional] 
**UserCode** | Pointer to **string** | UserCode is the short code the user types at VerifyURL. | [optional] 
**VerifyUrl** | Pointer to **string** | VerifyURL is the page the user opens to enter UserCode. | [optional] 

## Methods

### NewCloudDeviceStartOut

`func NewCloudDeviceStartOut() *CloudDeviceStartOut`

NewCloudDeviceStartOut instantiates a new CloudDeviceStartOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDeviceStartOutWithDefaults

`func NewCloudDeviceStartOutWithDefaults() *CloudDeviceStartOut`

NewCloudDeviceStartOutWithDefaults instantiates a new CloudDeviceStartOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *CloudDeviceStartOut) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CloudDeviceStartOut) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CloudDeviceStartOut) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CloudDeviceStartOut) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFlow

`func (o *CloudDeviceStartOut) GetFlow() string`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *CloudDeviceStartOut) GetFlowOk() (*string, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *CloudDeviceStartOut) SetFlow(v string)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *CloudDeviceStartOut) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetInterval

`func (o *CloudDeviceStartOut) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CloudDeviceStartOut) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CloudDeviceStartOut) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CloudDeviceStartOut) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetUserCode

`func (o *CloudDeviceStartOut) GetUserCode() string`

GetUserCode returns the UserCode field if non-nil, zero value otherwise.

### GetUserCodeOk

`func (o *CloudDeviceStartOut) GetUserCodeOk() (*string, bool)`

GetUserCodeOk returns a tuple with the UserCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCode

`func (o *CloudDeviceStartOut) SetUserCode(v string)`

SetUserCode sets UserCode field to given value.

### HasUserCode

`func (o *CloudDeviceStartOut) HasUserCode() bool`

HasUserCode returns a boolean if a field has been set.

### GetVerifyUrl

`func (o *CloudDeviceStartOut) GetVerifyUrl() string`

GetVerifyUrl returns the VerifyUrl field if non-nil, zero value otherwise.

### GetVerifyUrlOk

`func (o *CloudDeviceStartOut) GetVerifyUrlOk() (*string, bool)`

GetVerifyUrlOk returns a tuple with the VerifyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyUrl

`func (o *CloudDeviceStartOut) SetVerifyUrl(v string)`

SetVerifyUrl sets VerifyUrl field to given value.

### HasVerifyUrl

`func (o *CloudDeviceStartOut) HasVerifyUrl() bool`

HasVerifyUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


