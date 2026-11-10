# PatentPatentSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algo** | Pointer to **string** | Algo names the scheme, e.g. ed25519. Empty means the record does not say. | [optional] 
**KeyId** | Pointer to **string** | KeyID names the key that signed. It is a reference, not public bytes. | [optional] 
**Value** | Pointer to **string** | Value is the signature, base64. | [optional] 
**Verified** | Pointer to **bool** | Verified is the SERVER&#39;s answer to whether it checked this signature. It is false on a deployment with no verifier wired, and a record carrying an unchecked signature says so rather than borrowing credibility. | [optional] 

## Methods

### NewPatentPatentSignature

`func NewPatentPatentSignature() *PatentPatentSignature`

NewPatentPatentSignature instantiates a new PatentPatentSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentPatentSignatureWithDefaults

`func NewPatentPatentSignatureWithDefaults() *PatentPatentSignature`

NewPatentPatentSignatureWithDefaults instantiates a new PatentPatentSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgo

`func (o *PatentPatentSignature) GetAlgo() string`

GetAlgo returns the Algo field if non-nil, zero value otherwise.

### GetAlgoOk

`func (o *PatentPatentSignature) GetAlgoOk() (*string, bool)`

GetAlgoOk returns a tuple with the Algo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgo

`func (o *PatentPatentSignature) SetAlgo(v string)`

SetAlgo sets Algo field to given value.

### HasAlgo

`func (o *PatentPatentSignature) HasAlgo() bool`

HasAlgo returns a boolean if a field has been set.

### GetKeyId

`func (o *PatentPatentSignature) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *PatentPatentSignature) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *PatentPatentSignature) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *PatentPatentSignature) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetValue

`func (o *PatentPatentSignature) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatentPatentSignature) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatentPatentSignature) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PatentPatentSignature) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVerified

`func (o *PatentPatentSignature) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *PatentPatentSignature) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *PatentPatentSignature) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *PatentPatentSignature) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


