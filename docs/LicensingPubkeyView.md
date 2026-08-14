# LicensingPubkeyView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to **string** | Alg is always \&quot;Ed25519\&quot;. | [optional] 
**Keys** | Pointer to [**[]LicensingJWK**](LicensingJWK.md) | Keys is the same key as a single-entry JWKS (OKP/Ed25519), for JWKS-shaped consumers. | [optional] 
**Provider** | Pointer to **string** | Provider names the KMS holding the private half (\&quot;local\&quot; | \&quot;aws\&quot; | ...). \&quot;local\&quot; means a development key — never trust it in production. | [optional] 
**PublicKey** | Pointer to **string** | PublicKey is the 32-byte Ed25519 public key, standard base64. This is the form the engine embeds for offline verification. | [optional] 
**Schema** | Pointer to **int32** | Schema is the license payload schema version this key signs. | [optional] 
**TokenFormat** | Pointer to **string** | TokenFormat states the wire layout so an implementer can verify a token without this service&#39;s source. | [optional] 

## Methods

### NewLicensingPubkeyView

`func NewLicensingPubkeyView() *LicensingPubkeyView`

NewLicensingPubkeyView instantiates a new LicensingPubkeyView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensingPubkeyViewWithDefaults

`func NewLicensingPubkeyViewWithDefaults() *LicensingPubkeyView`

NewLicensingPubkeyViewWithDefaults instantiates a new LicensingPubkeyView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *LicensingPubkeyView) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *LicensingPubkeyView) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *LicensingPubkeyView) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *LicensingPubkeyView) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetKeys

`func (o *LicensingPubkeyView) GetKeys() []LicensingJWK`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *LicensingPubkeyView) GetKeysOk() (*[]LicensingJWK, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *LicensingPubkeyView) SetKeys(v []LicensingJWK)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *LicensingPubkeyView) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetProvider

`func (o *LicensingPubkeyView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *LicensingPubkeyView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *LicensingPubkeyView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *LicensingPubkeyView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPublicKey

`func (o *LicensingPubkeyView) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *LicensingPubkeyView) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *LicensingPubkeyView) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *LicensingPubkeyView) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetSchema

`func (o *LicensingPubkeyView) GetSchema() int32`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *LicensingPubkeyView) GetSchemaOk() (*int32, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *LicensingPubkeyView) SetSchema(v int32)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *LicensingPubkeyView) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetTokenFormat

`func (o *LicensingPubkeyView) GetTokenFormat() string`

GetTokenFormat returns the TokenFormat field if non-nil, zero value otherwise.

### GetTokenFormatOk

`func (o *LicensingPubkeyView) GetTokenFormatOk() (*string, bool)`

GetTokenFormatOk returns a tuple with the TokenFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenFormat

`func (o *LicensingPubkeyView) SetTokenFormat(v string)`

SetTokenFormat sets TokenFormat field to given value.

### HasTokenFormat

`func (o *LicensingPubkeyView) HasTokenFormat() bool`

HasTokenFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


