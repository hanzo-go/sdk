# O11ySigV4Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** |  | [optional] 
**ExternalID** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**RoleARN** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **interface{}** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**UseFIPSSTSEndpoint** | Pointer to **bool** |  | [optional] 

## Methods

### NewO11ySigV4Config

`func NewO11ySigV4Config() *O11ySigV4Config`

NewO11ySigV4Config instantiates a new O11ySigV4Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11ySigV4ConfigWithDefaults

`func NewO11ySigV4ConfigWithDefaults() *O11ySigV4Config`

NewO11ySigV4ConfigWithDefaults instantiates a new O11ySigV4Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *O11ySigV4Config) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *O11ySigV4Config) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *O11ySigV4Config) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *O11ySigV4Config) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetExternalID

`func (o *O11ySigV4Config) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *O11ySigV4Config) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *O11ySigV4Config) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *O11ySigV4Config) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetProfile

`func (o *O11ySigV4Config) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *O11ySigV4Config) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *O11ySigV4Config) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *O11ySigV4Config) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetRegion

`func (o *O11ySigV4Config) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *O11ySigV4Config) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *O11ySigV4Config) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *O11ySigV4Config) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRoleARN

`func (o *O11ySigV4Config) GetRoleARN() string`

GetRoleARN returns the RoleARN field if non-nil, zero value otherwise.

### GetRoleARNOk

`func (o *O11ySigV4Config) GetRoleARNOk() (*string, bool)`

GetRoleARNOk returns a tuple with the RoleARN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleARN

`func (o *O11ySigV4Config) SetRoleARN(v string)`

SetRoleARN sets RoleARN field to given value.

### HasRoleARN

`func (o *O11ySigV4Config) HasRoleARN() bool`

HasRoleARN returns a boolean if a field has been set.

### GetSecretKey

`func (o *O11ySigV4Config) GetSecretKey() interface{}`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *O11ySigV4Config) GetSecretKeyOk() (*interface{}, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *O11ySigV4Config) SetSecretKey(v interface{})`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *O11ySigV4Config) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### SetSecretKeyNil

`func (o *O11ySigV4Config) SetSecretKeyNil(b bool)`

 SetSecretKeyNil sets the value for SecretKey to be an explicit nil

### UnsetSecretKey
`func (o *O11ySigV4Config) UnsetSecretKey()`

UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
### GetServiceName

`func (o *O11ySigV4Config) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *O11ySigV4Config) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *O11ySigV4Config) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *O11ySigV4Config) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetUseFIPSSTSEndpoint

`func (o *O11ySigV4Config) GetUseFIPSSTSEndpoint() bool`

GetUseFIPSSTSEndpoint returns the UseFIPSSTSEndpoint field if non-nil, zero value otherwise.

### GetUseFIPSSTSEndpointOk

`func (o *O11ySigV4Config) GetUseFIPSSTSEndpointOk() (*bool, bool)`

GetUseFIPSSTSEndpointOk returns a tuple with the UseFIPSSTSEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFIPSSTSEndpoint

`func (o *O11ySigV4Config) SetUseFIPSSTSEndpoint(v bool)`

SetUseFIPSSTSEndpoint sets UseFIPSSTSEndpoint field to given value.

### HasUseFIPSSTSEndpoint

`func (o *O11ySigV4Config) HasUseFIPSSTSEndpoint() bool`

HasUseFIPSSTSEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


