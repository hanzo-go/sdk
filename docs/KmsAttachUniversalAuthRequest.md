# KmsAttachUniversalAuthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenTrustedIps** | Pointer to [**[]KmsAttachUniversalAuthRequestAccessTokenTrustedIpsInner**](KmsAttachUniversalAuthRequestAccessTokenTrustedIpsInner.md) |  | [optional] 
**AccessTokenTTL** | Pointer to **int32** |  | [optional] [default to 2592000]
**AccessTokenMaxTTL** | Pointer to **int32** |  | [optional] [default to 2592000]
**AccessTokenNumUsesLimit** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewKmsAttachUniversalAuthRequest

`func NewKmsAttachUniversalAuthRequest() *KmsAttachUniversalAuthRequest`

NewKmsAttachUniversalAuthRequest instantiates a new KmsAttachUniversalAuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsAttachUniversalAuthRequestWithDefaults

`func NewKmsAttachUniversalAuthRequestWithDefaults() *KmsAttachUniversalAuthRequest`

NewKmsAttachUniversalAuthRequestWithDefaults instantiates a new KmsAttachUniversalAuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenTrustedIps

`func (o *KmsAttachUniversalAuthRequest) GetAccessTokenTrustedIps() []KmsAttachUniversalAuthRequestAccessTokenTrustedIpsInner`

GetAccessTokenTrustedIps returns the AccessTokenTrustedIps field if non-nil, zero value otherwise.

### GetAccessTokenTrustedIpsOk

`func (o *KmsAttachUniversalAuthRequest) GetAccessTokenTrustedIpsOk() (*[]KmsAttachUniversalAuthRequestAccessTokenTrustedIpsInner, bool)`

GetAccessTokenTrustedIpsOk returns a tuple with the AccessTokenTrustedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenTrustedIps

`func (o *KmsAttachUniversalAuthRequest) SetAccessTokenTrustedIps(v []KmsAttachUniversalAuthRequestAccessTokenTrustedIpsInner)`

SetAccessTokenTrustedIps sets AccessTokenTrustedIps field to given value.

### HasAccessTokenTrustedIps

`func (o *KmsAttachUniversalAuthRequest) HasAccessTokenTrustedIps() bool`

HasAccessTokenTrustedIps returns a boolean if a field has been set.

### GetAccessTokenTTL

`func (o *KmsAttachUniversalAuthRequest) GetAccessTokenTTL() int32`

GetAccessTokenTTL returns the AccessTokenTTL field if non-nil, zero value otherwise.

### GetAccessTokenTTLOk

`func (o *KmsAttachUniversalAuthRequest) GetAccessTokenTTLOk() (*int32, bool)`

GetAccessTokenTTLOk returns a tuple with the AccessTokenTTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenTTL

`func (o *KmsAttachUniversalAuthRequest) SetAccessTokenTTL(v int32)`

SetAccessTokenTTL sets AccessTokenTTL field to given value.

### HasAccessTokenTTL

`func (o *KmsAttachUniversalAuthRequest) HasAccessTokenTTL() bool`

HasAccessTokenTTL returns a boolean if a field has been set.

### GetAccessTokenMaxTTL

`func (o *KmsAttachUniversalAuthRequest) GetAccessTokenMaxTTL() int32`

GetAccessTokenMaxTTL returns the AccessTokenMaxTTL field if non-nil, zero value otherwise.

### GetAccessTokenMaxTTLOk

`func (o *KmsAttachUniversalAuthRequest) GetAccessTokenMaxTTLOk() (*int32, bool)`

GetAccessTokenMaxTTLOk returns a tuple with the AccessTokenMaxTTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenMaxTTL

`func (o *KmsAttachUniversalAuthRequest) SetAccessTokenMaxTTL(v int32)`

SetAccessTokenMaxTTL sets AccessTokenMaxTTL field to given value.

### HasAccessTokenMaxTTL

`func (o *KmsAttachUniversalAuthRequest) HasAccessTokenMaxTTL() bool`

HasAccessTokenMaxTTL returns a boolean if a field has been set.

### GetAccessTokenNumUsesLimit

`func (o *KmsAttachUniversalAuthRequest) GetAccessTokenNumUsesLimit() int32`

GetAccessTokenNumUsesLimit returns the AccessTokenNumUsesLimit field if non-nil, zero value otherwise.

### GetAccessTokenNumUsesLimitOk

`func (o *KmsAttachUniversalAuthRequest) GetAccessTokenNumUsesLimitOk() (*int32, bool)`

GetAccessTokenNumUsesLimitOk returns a tuple with the AccessTokenNumUsesLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenNumUsesLimit

`func (o *KmsAttachUniversalAuthRequest) SetAccessTokenNumUsesLimit(v int32)`

SetAccessTokenNumUsesLimit sets AccessTokenNumUsesLimit field to given value.

### HasAccessTokenNumUsesLimit

`func (o *KmsAttachUniversalAuthRequest) HasAccessTokenNumUsesLimit() bool`

HasAccessTokenNumUsesLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


