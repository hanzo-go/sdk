# PlatformEnvVarJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key is the variable&#39;s name in the container, which must match &#x60;^[A-Za-z_][A-Za-z0-9_]*$&#x60;. For a sealed value it is also the last segment of the KMS ref, so it is what identifies the value across a round trip. | [optional] 
**Secret** | Pointer to **bool** | Secret says the value lives in KMS and never in the database. A caller may only ADD secrecy: the server seals a value whose key or shape looks like a credential anyway (secretshape.go), so an entry can come back secret that was not sent that way. | [optional] 
**Value** | Pointer to **string** | Value is the plaintext, and it is WRITE-ONLY once the entry is secret: a sealed value reads back as \&quot;\&quot;, and sending \&quot;\&quot; again KEEPS what is sealed rather than wiping it. Only a non-empty value seals a new one. | [optional] 

## Methods

### NewPlatformEnvVarJSON

`func NewPlatformEnvVarJSON() *PlatformEnvVarJSON`

NewPlatformEnvVarJSON instantiates a new PlatformEnvVarJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformEnvVarJSONWithDefaults

`func NewPlatformEnvVarJSONWithDefaults() *PlatformEnvVarJSON`

NewPlatformEnvVarJSONWithDefaults instantiates a new PlatformEnvVarJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PlatformEnvVarJSON) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PlatformEnvVarJSON) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PlatformEnvVarJSON) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PlatformEnvVarJSON) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSecret

`func (o *PlatformEnvVarJSON) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PlatformEnvVarJSON) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PlatformEnvVarJSON) SetSecret(v bool)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PlatformEnvVarJSON) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetValue

`func (o *PlatformEnvVarJSON) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PlatformEnvVarJSON) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PlatformEnvVarJSON) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PlatformEnvVarJSON) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


