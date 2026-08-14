# LicensingJWK

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crv** | Pointer to **string** | Crv is always \&quot;Ed25519\&quot;. | [optional] 
**Kty** | Pointer to **string** | Kty is always \&quot;OKP\&quot;. | [optional] 
**Use** | Pointer to **string** | Use is always \&quot;sig\&quot;. | [optional] 
**X** | Pointer to **string** | X is the public key, base64url (the JWK convention). | [optional] 

## Methods

### NewLicensingJWK

`func NewLicensingJWK() *LicensingJWK`

NewLicensingJWK instantiates a new LicensingJWK object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensingJWKWithDefaults

`func NewLicensingJWKWithDefaults() *LicensingJWK`

NewLicensingJWKWithDefaults instantiates a new LicensingJWK object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrv

`func (o *LicensingJWK) GetCrv() string`

GetCrv returns the Crv field if non-nil, zero value otherwise.

### GetCrvOk

`func (o *LicensingJWK) GetCrvOk() (*string, bool)`

GetCrvOk returns a tuple with the Crv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrv

`func (o *LicensingJWK) SetCrv(v string)`

SetCrv sets Crv field to given value.

### HasCrv

`func (o *LicensingJWK) HasCrv() bool`

HasCrv returns a boolean if a field has been set.

### GetKty

`func (o *LicensingJWK) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *LicensingJWK) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *LicensingJWK) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *LicensingJWK) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetUse

`func (o *LicensingJWK) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *LicensingJWK) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *LicensingJWK) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *LicensingJWK) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetX

`func (o *LicensingJWK) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *LicensingJWK) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *LicensingJWK) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *LicensingJWK) HasX() bool`

HasX returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


