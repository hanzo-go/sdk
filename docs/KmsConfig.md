# KmsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **string** | APIBase is this subsystem&#39;s own prefix, &#x60;/v1/kms&#x60;. | [optional] 
**Brand** | Pointer to **string** | Brand is the deployment&#39;s brand, so the console renders as the right product. | [optional] 
**Issuer** | Pointer to **string** | Issuer is the OIDC issuer the console authenticates against. | [optional] 
**LoginPath** | Pointer to **string** | LoginPath is the credential exchange&#39;s address. | [optional] 

## Methods

### NewKmsConfig

`func NewKmsConfig() *KmsConfig`

NewKmsConfig instantiates a new KmsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsConfigWithDefaults

`func NewKmsConfigWithDefaults() *KmsConfig`

NewKmsConfigWithDefaults instantiates a new KmsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *KmsConfig) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *KmsConfig) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *KmsConfig) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *KmsConfig) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### GetBrand

`func (o *KmsConfig) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *KmsConfig) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *KmsConfig) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *KmsConfig) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetIssuer

`func (o *KmsConfig) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *KmsConfig) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *KmsConfig) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *KmsConfig) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetLoginPath

`func (o *KmsConfig) GetLoginPath() string`

GetLoginPath returns the LoginPath field if non-nil, zero value otherwise.

### GetLoginPathOk

`func (o *KmsConfig) GetLoginPathOk() (*string, bool)`

GetLoginPathOk returns a tuple with the LoginPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginPath

`func (o *KmsConfig) SetLoginPath(v string)`

SetLoginPath sets LoginPath field to given value.

### HasLoginPath

`func (o *KmsConfig) HasLoginPath() bool`

HasLoginPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


