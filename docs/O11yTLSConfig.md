# O11yTLSConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ca** | Pointer to **string** | Text of the CA cert to use for the targets. | [optional] 
**CaFile** | Pointer to **string** | The CA cert to use for the targets. | [optional] 
**CaRef** | Pointer to **string** | CARef is the name of the secret within the secret manager to use as the CA cert for the targets. | [optional] 
**Cert** | Pointer to **string** | Text of the client cert file for the targets. | [optional] 
**CertFile** | Pointer to **string** | The client cert file for the targets. | [optional] 
**CertRef** | Pointer to **string** | CertRef is the name of the secret within the secret manager to use as the client cert for the targets. | [optional] 
**InsecureSkipVerify** | Pointer to **bool** | Disable target certificate validation. | [optional] 
**Key** | Pointer to **interface{}** |  | [optional] 
**KeyFile** | Pointer to **string** | The client key file for the targets. | [optional] 
**KeyRef** | Pointer to **string** | KeyRef is the name of the secret within the secret manager to use as the client key for the targets. | [optional] 
**MaxVersion** | Pointer to **interface{}** |  | [optional] 
**MinVersion** | Pointer to **interface{}** |  | [optional] 
**ServerName** | Pointer to **string** | Used to verify the hostname for the targets. | [optional] 

## Methods

### NewO11yTLSConfig

`func NewO11yTLSConfig() *O11yTLSConfig`

NewO11yTLSConfig instantiates a new O11yTLSConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yTLSConfigWithDefaults

`func NewO11yTLSConfigWithDefaults() *O11yTLSConfig`

NewO11yTLSConfigWithDefaults instantiates a new O11yTLSConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCa

`func (o *O11yTLSConfig) GetCa() string`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *O11yTLSConfig) GetCaOk() (*string, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *O11yTLSConfig) SetCa(v string)`

SetCa sets Ca field to given value.

### HasCa

`func (o *O11yTLSConfig) HasCa() bool`

HasCa returns a boolean if a field has been set.

### GetCaFile

`func (o *O11yTLSConfig) GetCaFile() string`

GetCaFile returns the CaFile field if non-nil, zero value otherwise.

### GetCaFileOk

`func (o *O11yTLSConfig) GetCaFileOk() (*string, bool)`

GetCaFileOk returns a tuple with the CaFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaFile

`func (o *O11yTLSConfig) SetCaFile(v string)`

SetCaFile sets CaFile field to given value.

### HasCaFile

`func (o *O11yTLSConfig) HasCaFile() bool`

HasCaFile returns a boolean if a field has been set.

### GetCaRef

`func (o *O11yTLSConfig) GetCaRef() string`

GetCaRef returns the CaRef field if non-nil, zero value otherwise.

### GetCaRefOk

`func (o *O11yTLSConfig) GetCaRefOk() (*string, bool)`

GetCaRefOk returns a tuple with the CaRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaRef

`func (o *O11yTLSConfig) SetCaRef(v string)`

SetCaRef sets CaRef field to given value.

### HasCaRef

`func (o *O11yTLSConfig) HasCaRef() bool`

HasCaRef returns a boolean if a field has been set.

### GetCert

`func (o *O11yTLSConfig) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *O11yTLSConfig) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *O11yTLSConfig) SetCert(v string)`

SetCert sets Cert field to given value.

### HasCert

`func (o *O11yTLSConfig) HasCert() bool`

HasCert returns a boolean if a field has been set.

### GetCertFile

`func (o *O11yTLSConfig) GetCertFile() string`

GetCertFile returns the CertFile field if non-nil, zero value otherwise.

### GetCertFileOk

`func (o *O11yTLSConfig) GetCertFileOk() (*string, bool)`

GetCertFileOk returns a tuple with the CertFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertFile

`func (o *O11yTLSConfig) SetCertFile(v string)`

SetCertFile sets CertFile field to given value.

### HasCertFile

`func (o *O11yTLSConfig) HasCertFile() bool`

HasCertFile returns a boolean if a field has been set.

### GetCertRef

`func (o *O11yTLSConfig) GetCertRef() string`

GetCertRef returns the CertRef field if non-nil, zero value otherwise.

### GetCertRefOk

`func (o *O11yTLSConfig) GetCertRefOk() (*string, bool)`

GetCertRefOk returns a tuple with the CertRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertRef

`func (o *O11yTLSConfig) SetCertRef(v string)`

SetCertRef sets CertRef field to given value.

### HasCertRef

`func (o *O11yTLSConfig) HasCertRef() bool`

HasCertRef returns a boolean if a field has been set.

### GetInsecureSkipVerify

`func (o *O11yTLSConfig) GetInsecureSkipVerify() bool`

GetInsecureSkipVerify returns the InsecureSkipVerify field if non-nil, zero value otherwise.

### GetInsecureSkipVerifyOk

`func (o *O11yTLSConfig) GetInsecureSkipVerifyOk() (*bool, bool)`

GetInsecureSkipVerifyOk returns a tuple with the InsecureSkipVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSkipVerify

`func (o *O11yTLSConfig) SetInsecureSkipVerify(v bool)`

SetInsecureSkipVerify sets InsecureSkipVerify field to given value.

### HasInsecureSkipVerify

`func (o *O11yTLSConfig) HasInsecureSkipVerify() bool`

HasInsecureSkipVerify returns a boolean if a field has been set.

### GetKey

`func (o *O11yTLSConfig) GetKey() interface{}`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *O11yTLSConfig) GetKeyOk() (*interface{}, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *O11yTLSConfig) SetKey(v interface{})`

SetKey sets Key field to given value.

### HasKey

`func (o *O11yTLSConfig) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *O11yTLSConfig) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *O11yTLSConfig) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetKeyFile

`func (o *O11yTLSConfig) GetKeyFile() string`

GetKeyFile returns the KeyFile field if non-nil, zero value otherwise.

### GetKeyFileOk

`func (o *O11yTLSConfig) GetKeyFileOk() (*string, bool)`

GetKeyFileOk returns a tuple with the KeyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFile

`func (o *O11yTLSConfig) SetKeyFile(v string)`

SetKeyFile sets KeyFile field to given value.

### HasKeyFile

`func (o *O11yTLSConfig) HasKeyFile() bool`

HasKeyFile returns a boolean if a field has been set.

### GetKeyRef

`func (o *O11yTLSConfig) GetKeyRef() string`

GetKeyRef returns the KeyRef field if non-nil, zero value otherwise.

### GetKeyRefOk

`func (o *O11yTLSConfig) GetKeyRefOk() (*string, bool)`

GetKeyRefOk returns a tuple with the KeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRef

`func (o *O11yTLSConfig) SetKeyRef(v string)`

SetKeyRef sets KeyRef field to given value.

### HasKeyRef

`func (o *O11yTLSConfig) HasKeyRef() bool`

HasKeyRef returns a boolean if a field has been set.

### GetMaxVersion

`func (o *O11yTLSConfig) GetMaxVersion() interface{}`

GetMaxVersion returns the MaxVersion field if non-nil, zero value otherwise.

### GetMaxVersionOk

`func (o *O11yTLSConfig) GetMaxVersionOk() (*interface{}, bool)`

GetMaxVersionOk returns a tuple with the MaxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersion

`func (o *O11yTLSConfig) SetMaxVersion(v interface{})`

SetMaxVersion sets MaxVersion field to given value.

### HasMaxVersion

`func (o *O11yTLSConfig) HasMaxVersion() bool`

HasMaxVersion returns a boolean if a field has been set.

### SetMaxVersionNil

`func (o *O11yTLSConfig) SetMaxVersionNil(b bool)`

 SetMaxVersionNil sets the value for MaxVersion to be an explicit nil

### UnsetMaxVersion
`func (o *O11yTLSConfig) UnsetMaxVersion()`

UnsetMaxVersion ensures that no value is present for MaxVersion, not even an explicit nil
### GetMinVersion

`func (o *O11yTLSConfig) GetMinVersion() interface{}`

GetMinVersion returns the MinVersion field if non-nil, zero value otherwise.

### GetMinVersionOk

`func (o *O11yTLSConfig) GetMinVersionOk() (*interface{}, bool)`

GetMinVersionOk returns a tuple with the MinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVersion

`func (o *O11yTLSConfig) SetMinVersion(v interface{})`

SetMinVersion sets MinVersion field to given value.

### HasMinVersion

`func (o *O11yTLSConfig) HasMinVersion() bool`

HasMinVersion returns a boolean if a field has been set.

### SetMinVersionNil

`func (o *O11yTLSConfig) SetMinVersionNil(b bool)`

 SetMinVersionNil sets the value for MinVersion to be an explicit nil

### UnsetMinVersion
`func (o *O11yTLSConfig) UnsetMinVersion()`

UnsetMinVersion ensures that no value is present for MinVersion, not even an explicit nil
### GetServerName

`func (o *O11yTLSConfig) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *O11yTLSConfig) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *O11yTLSConfig) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *O11yTLSConfig) HasServerName() bool`

HasServerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


