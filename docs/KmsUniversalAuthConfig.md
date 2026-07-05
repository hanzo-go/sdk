# KmsUniversalAuthConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityId** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**AccessTokenTrustedIps** | Pointer to [**[]KmsAttachUniversalAuthRequestAccessTokenTrustedIpsInner**](KmsAttachUniversalAuthRequestAccessTokenTrustedIpsInner.md) |  | [optional] 
**AccessTokenTTL** | Pointer to **int32** |  | [optional] 
**AccessTokenMaxTTL** | Pointer to **int32** |  | [optional] 
**AccessTokenNumUsesLimit** | Pointer to **int32** |  | [optional] 

## Methods

### NewKmsUniversalAuthConfig

`func NewKmsUniversalAuthConfig() *KmsUniversalAuthConfig`

NewKmsUniversalAuthConfig instantiates a new KmsUniversalAuthConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsUniversalAuthConfigWithDefaults

`func NewKmsUniversalAuthConfigWithDefaults() *KmsUniversalAuthConfig`

NewKmsUniversalAuthConfigWithDefaults instantiates a new KmsUniversalAuthConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityId

`func (o *KmsUniversalAuthConfig) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *KmsUniversalAuthConfig) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *KmsUniversalAuthConfig) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *KmsUniversalAuthConfig) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetClientId

`func (o *KmsUniversalAuthConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *KmsUniversalAuthConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *KmsUniversalAuthConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *KmsUniversalAuthConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetAccessTokenTrustedIps

`func (o *KmsUniversalAuthConfig) GetAccessTokenTrustedIps() []KmsAttachUniversalAuthRequestAccessTokenTrustedIpsInner`

GetAccessTokenTrustedIps returns the AccessTokenTrustedIps field if non-nil, zero value otherwise.

### GetAccessTokenTrustedIpsOk

`func (o *KmsUniversalAuthConfig) GetAccessTokenTrustedIpsOk() (*[]KmsAttachUniversalAuthRequestAccessTokenTrustedIpsInner, bool)`

GetAccessTokenTrustedIpsOk returns a tuple with the AccessTokenTrustedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenTrustedIps

`func (o *KmsUniversalAuthConfig) SetAccessTokenTrustedIps(v []KmsAttachUniversalAuthRequestAccessTokenTrustedIpsInner)`

SetAccessTokenTrustedIps sets AccessTokenTrustedIps field to given value.

### HasAccessTokenTrustedIps

`func (o *KmsUniversalAuthConfig) HasAccessTokenTrustedIps() bool`

HasAccessTokenTrustedIps returns a boolean if a field has been set.

### GetAccessTokenTTL

`func (o *KmsUniversalAuthConfig) GetAccessTokenTTL() int32`

GetAccessTokenTTL returns the AccessTokenTTL field if non-nil, zero value otherwise.

### GetAccessTokenTTLOk

`func (o *KmsUniversalAuthConfig) GetAccessTokenTTLOk() (*int32, bool)`

GetAccessTokenTTLOk returns a tuple with the AccessTokenTTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenTTL

`func (o *KmsUniversalAuthConfig) SetAccessTokenTTL(v int32)`

SetAccessTokenTTL sets AccessTokenTTL field to given value.

### HasAccessTokenTTL

`func (o *KmsUniversalAuthConfig) HasAccessTokenTTL() bool`

HasAccessTokenTTL returns a boolean if a field has been set.

### GetAccessTokenMaxTTL

`func (o *KmsUniversalAuthConfig) GetAccessTokenMaxTTL() int32`

GetAccessTokenMaxTTL returns the AccessTokenMaxTTL field if non-nil, zero value otherwise.

### GetAccessTokenMaxTTLOk

`func (o *KmsUniversalAuthConfig) GetAccessTokenMaxTTLOk() (*int32, bool)`

GetAccessTokenMaxTTLOk returns a tuple with the AccessTokenMaxTTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenMaxTTL

`func (o *KmsUniversalAuthConfig) SetAccessTokenMaxTTL(v int32)`

SetAccessTokenMaxTTL sets AccessTokenMaxTTL field to given value.

### HasAccessTokenMaxTTL

`func (o *KmsUniversalAuthConfig) HasAccessTokenMaxTTL() bool`

HasAccessTokenMaxTTL returns a boolean if a field has been set.

### GetAccessTokenNumUsesLimit

`func (o *KmsUniversalAuthConfig) GetAccessTokenNumUsesLimit() int32`

GetAccessTokenNumUsesLimit returns the AccessTokenNumUsesLimit field if non-nil, zero value otherwise.

### GetAccessTokenNumUsesLimitOk

`func (o *KmsUniversalAuthConfig) GetAccessTokenNumUsesLimitOk() (*int32, bool)`

GetAccessTokenNumUsesLimitOk returns a tuple with the AccessTokenNumUsesLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenNumUsesLimit

`func (o *KmsUniversalAuthConfig) SetAccessTokenNumUsesLimit(v int32)`

SetAccessTokenNumUsesLimit sets AccessTokenNumUsesLimit field to given value.

### HasAccessTokenNumUsesLimit

`func (o *KmsUniversalAuthConfig) HasAccessTokenNumUsesLimit() bool`

HasAccessTokenNumUsesLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


