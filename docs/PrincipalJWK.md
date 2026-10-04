# PrincipalJWK

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to **string** | Alg is EdDSA. | [optional] 
**Crv** | Pointer to **string** | Crv is Ed25519. | [optional] 
**Kid** | Pointer to **string** | Kid names the key; a statement&#39;s header carries it. | [optional] 
**Kty** | Pointer to **string** | Kty is OKP, an octet key pair. | [optional] 
**Use** | Pointer to **string** | Use is sig. | [optional] 
**X** | Pointer to **string** | X is the public key, base64url without padding. | [optional] 

## Methods

### NewPrincipalJWK

`func NewPrincipalJWK() *PrincipalJWK`

NewPrincipalJWK instantiates a new PrincipalJWK object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalJWKWithDefaults

`func NewPrincipalJWKWithDefaults() *PrincipalJWK`

NewPrincipalJWKWithDefaults instantiates a new PrincipalJWK object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *PrincipalJWK) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *PrincipalJWK) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *PrincipalJWK) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *PrincipalJWK) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetCrv

`func (o *PrincipalJWK) GetCrv() string`

GetCrv returns the Crv field if non-nil, zero value otherwise.

### GetCrvOk

`func (o *PrincipalJWK) GetCrvOk() (*string, bool)`

GetCrvOk returns a tuple with the Crv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrv

`func (o *PrincipalJWK) SetCrv(v string)`

SetCrv sets Crv field to given value.

### HasCrv

`func (o *PrincipalJWK) HasCrv() bool`

HasCrv returns a boolean if a field has been set.

### GetKid

`func (o *PrincipalJWK) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *PrincipalJWK) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *PrincipalJWK) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *PrincipalJWK) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetKty

`func (o *PrincipalJWK) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *PrincipalJWK) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *PrincipalJWK) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *PrincipalJWK) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetUse

`func (o *PrincipalJWK) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *PrincipalJWK) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *PrincipalJWK) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *PrincipalJWK) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetX

`func (o *PrincipalJWK) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *PrincipalJWK) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *PrincipalJWK) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *PrincipalJWK) HasX() bool`

HasX returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


