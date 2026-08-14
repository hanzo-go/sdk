# O11yO11ySAMLConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeMapping** | Pointer to [**O11yO11yAttributeMapping**](O11yO11yAttributeMapping.md) | AttributeMapping names the assertion attributes to read identity from. | [optional] 
**InsecureSkipAuthNRequestsSigned** | Pointer to **bool** | InsecureSkipAuthNRequestsSigned skips signing outgoing AuthN requests, for IdPs that refuse signed ones. | [optional] 
**SamlCert** | Pointer to **string** | SamlCert is the IdP&#39;s signing certificate. | [optional] 
**SamlEntity** | Pointer to **string** | SamlEntity is the IdP&#39;s entityID. | [optional] 
**SamlIdp** | Pointer to **string** | SamlIdp is the IdP&#39;s single-sign-on endpoint. | [optional] 

## Methods

### NewO11yO11ySAMLConfig

`func NewO11yO11ySAMLConfig() *O11yO11ySAMLConfig`

NewO11yO11ySAMLConfig instantiates a new O11yO11ySAMLConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11ySAMLConfigWithDefaults

`func NewO11yO11ySAMLConfigWithDefaults() *O11yO11ySAMLConfig`

NewO11yO11ySAMLConfigWithDefaults instantiates a new O11yO11ySAMLConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeMapping

`func (o *O11yO11ySAMLConfig) GetAttributeMapping() O11yO11yAttributeMapping`

GetAttributeMapping returns the AttributeMapping field if non-nil, zero value otherwise.

### GetAttributeMappingOk

`func (o *O11yO11ySAMLConfig) GetAttributeMappingOk() (*O11yO11yAttributeMapping, bool)`

GetAttributeMappingOk returns a tuple with the AttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMapping

`func (o *O11yO11ySAMLConfig) SetAttributeMapping(v O11yO11yAttributeMapping)`

SetAttributeMapping sets AttributeMapping field to given value.

### HasAttributeMapping

`func (o *O11yO11ySAMLConfig) HasAttributeMapping() bool`

HasAttributeMapping returns a boolean if a field has been set.

### GetInsecureSkipAuthNRequestsSigned

`func (o *O11yO11ySAMLConfig) GetInsecureSkipAuthNRequestsSigned() bool`

GetInsecureSkipAuthNRequestsSigned returns the InsecureSkipAuthNRequestsSigned field if non-nil, zero value otherwise.

### GetInsecureSkipAuthNRequestsSignedOk

`func (o *O11yO11ySAMLConfig) GetInsecureSkipAuthNRequestsSignedOk() (*bool, bool)`

GetInsecureSkipAuthNRequestsSignedOk returns a tuple with the InsecureSkipAuthNRequestsSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSkipAuthNRequestsSigned

`func (o *O11yO11ySAMLConfig) SetInsecureSkipAuthNRequestsSigned(v bool)`

SetInsecureSkipAuthNRequestsSigned sets InsecureSkipAuthNRequestsSigned field to given value.

### HasInsecureSkipAuthNRequestsSigned

`func (o *O11yO11ySAMLConfig) HasInsecureSkipAuthNRequestsSigned() bool`

HasInsecureSkipAuthNRequestsSigned returns a boolean if a field has been set.

### GetSamlCert

`func (o *O11yO11ySAMLConfig) GetSamlCert() string`

GetSamlCert returns the SamlCert field if non-nil, zero value otherwise.

### GetSamlCertOk

`func (o *O11yO11ySAMLConfig) GetSamlCertOk() (*string, bool)`

GetSamlCertOk returns a tuple with the SamlCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlCert

`func (o *O11yO11ySAMLConfig) SetSamlCert(v string)`

SetSamlCert sets SamlCert field to given value.

### HasSamlCert

`func (o *O11yO11ySAMLConfig) HasSamlCert() bool`

HasSamlCert returns a boolean if a field has been set.

### GetSamlEntity

`func (o *O11yO11ySAMLConfig) GetSamlEntity() string`

GetSamlEntity returns the SamlEntity field if non-nil, zero value otherwise.

### GetSamlEntityOk

`func (o *O11yO11ySAMLConfig) GetSamlEntityOk() (*string, bool)`

GetSamlEntityOk returns a tuple with the SamlEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlEntity

`func (o *O11yO11ySAMLConfig) SetSamlEntity(v string)`

SetSamlEntity sets SamlEntity field to given value.

### HasSamlEntity

`func (o *O11yO11ySAMLConfig) HasSamlEntity() bool`

HasSamlEntity returns a boolean if a field has been set.

### GetSamlIdp

`func (o *O11yO11ySAMLConfig) GetSamlIdp() string`

GetSamlIdp returns the SamlIdp field if non-nil, zero value otherwise.

### GetSamlIdpOk

`func (o *O11yO11ySAMLConfig) GetSamlIdpOk() (*string, bool)`

GetSamlIdpOk returns a tuple with the SamlIdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlIdp

`func (o *O11yO11ySAMLConfig) SetSamlIdp(v string)`

SetSamlIdp sets SamlIdp field to given value.

### HasSamlIdp

`func (o *O11yO11ySAMLConfig) HasSamlIdp() bool`

HasSamlIdp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


