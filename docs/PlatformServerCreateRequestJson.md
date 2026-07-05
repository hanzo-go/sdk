# PlatformServerCreateRequestJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**IpAddress** | **string** |  | 
**Port** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**SshKeyId** | **string** |  | 

## Methods

### NewPlatformServerCreateRequestJson

`func NewPlatformServerCreateRequestJson(name string, ipAddress string, sshKeyId string, ) *PlatformServerCreateRequestJson`

NewPlatformServerCreateRequestJson instantiates a new PlatformServerCreateRequestJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformServerCreateRequestJsonWithDefaults

`func NewPlatformServerCreateRequestJsonWithDefaults() *PlatformServerCreateRequestJson`

NewPlatformServerCreateRequestJsonWithDefaults instantiates a new PlatformServerCreateRequestJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlatformServerCreateRequestJson) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlatformServerCreateRequestJson) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlatformServerCreateRequestJson) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PlatformServerCreateRequestJson) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlatformServerCreateRequestJson) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlatformServerCreateRequestJson) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlatformServerCreateRequestJson) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpAddress

`func (o *PlatformServerCreateRequestJson) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *PlatformServerCreateRequestJson) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *PlatformServerCreateRequestJson) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetPort

`func (o *PlatformServerCreateRequestJson) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PlatformServerCreateRequestJson) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PlatformServerCreateRequestJson) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PlatformServerCreateRequestJson) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *PlatformServerCreateRequestJson) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PlatformServerCreateRequestJson) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PlatformServerCreateRequestJson) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PlatformServerCreateRequestJson) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetSshKeyId

`func (o *PlatformServerCreateRequestJson) GetSshKeyId() string`

GetSshKeyId returns the SshKeyId field if non-nil, zero value otherwise.

### GetSshKeyIdOk

`func (o *PlatformServerCreateRequestJson) GetSshKeyIdOk() (*string, bool)`

GetSshKeyIdOk returns a tuple with the SshKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyId

`func (o *PlatformServerCreateRequestJson) SetSshKeyId(v string)`

SetSshKeyId sets SshKeyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


