# KmsCreateCertificateAuthorityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**Type** | **string** |  | 
**FriendlyName** | **string** |  | 
**CommonName** | **string** |  | 
**Organization** | Pointer to **string** |  | [optional] 
**Ou** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Province** | Pointer to **string** |  | [optional] 
**Locality** | Pointer to **string** |  | [optional] 
**KeyAlgorithm** | **string** |  | 
**MaxPathLength** | Pointer to **int32** |  | [optional] 
**NotAfter** | Pointer to **string** | Duration (e.g., \&quot;10y\&quot;) | [optional] 

## Methods

### NewKmsCreateCertificateAuthorityRequest

`func NewKmsCreateCertificateAuthorityRequest(projectId string, type_ string, friendlyName string, commonName string, keyAlgorithm string, ) *KmsCreateCertificateAuthorityRequest`

NewKmsCreateCertificateAuthorityRequest instantiates a new KmsCreateCertificateAuthorityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsCreateCertificateAuthorityRequestWithDefaults

`func NewKmsCreateCertificateAuthorityRequestWithDefaults() *KmsCreateCertificateAuthorityRequest`

NewKmsCreateCertificateAuthorityRequestWithDefaults instantiates a new KmsCreateCertificateAuthorityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *KmsCreateCertificateAuthorityRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *KmsCreateCertificateAuthorityRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *KmsCreateCertificateAuthorityRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetType

`func (o *KmsCreateCertificateAuthorityRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KmsCreateCertificateAuthorityRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KmsCreateCertificateAuthorityRequest) SetType(v string)`

SetType sets Type field to given value.


### GetFriendlyName

`func (o *KmsCreateCertificateAuthorityRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *KmsCreateCertificateAuthorityRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *KmsCreateCertificateAuthorityRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.


### GetCommonName

`func (o *KmsCreateCertificateAuthorityRequest) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *KmsCreateCertificateAuthorityRequest) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *KmsCreateCertificateAuthorityRequest) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### GetOrganization

`func (o *KmsCreateCertificateAuthorityRequest) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *KmsCreateCertificateAuthorityRequest) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *KmsCreateCertificateAuthorityRequest) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *KmsCreateCertificateAuthorityRequest) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOu

`func (o *KmsCreateCertificateAuthorityRequest) GetOu() string`

GetOu returns the Ou field if non-nil, zero value otherwise.

### GetOuOk

`func (o *KmsCreateCertificateAuthorityRequest) GetOuOk() (*string, bool)`

GetOuOk returns a tuple with the Ou field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOu

`func (o *KmsCreateCertificateAuthorityRequest) SetOu(v string)`

SetOu sets Ou field to given value.

### HasOu

`func (o *KmsCreateCertificateAuthorityRequest) HasOu() bool`

HasOu returns a boolean if a field has been set.

### GetCountry

`func (o *KmsCreateCertificateAuthorityRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *KmsCreateCertificateAuthorityRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *KmsCreateCertificateAuthorityRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *KmsCreateCertificateAuthorityRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetProvince

`func (o *KmsCreateCertificateAuthorityRequest) GetProvince() string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *KmsCreateCertificateAuthorityRequest) GetProvinceOk() (*string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *KmsCreateCertificateAuthorityRequest) SetProvince(v string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *KmsCreateCertificateAuthorityRequest) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### GetLocality

`func (o *KmsCreateCertificateAuthorityRequest) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *KmsCreateCertificateAuthorityRequest) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *KmsCreateCertificateAuthorityRequest) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *KmsCreateCertificateAuthorityRequest) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetKeyAlgorithm

`func (o *KmsCreateCertificateAuthorityRequest) GetKeyAlgorithm() string`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *KmsCreateCertificateAuthorityRequest) GetKeyAlgorithmOk() (*string, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *KmsCreateCertificateAuthorityRequest) SetKeyAlgorithm(v string)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.


### GetMaxPathLength

`func (o *KmsCreateCertificateAuthorityRequest) GetMaxPathLength() int32`

GetMaxPathLength returns the MaxPathLength field if non-nil, zero value otherwise.

### GetMaxPathLengthOk

`func (o *KmsCreateCertificateAuthorityRequest) GetMaxPathLengthOk() (*int32, bool)`

GetMaxPathLengthOk returns a tuple with the MaxPathLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPathLength

`func (o *KmsCreateCertificateAuthorityRequest) SetMaxPathLength(v int32)`

SetMaxPathLength sets MaxPathLength field to given value.

### HasMaxPathLength

`func (o *KmsCreateCertificateAuthorityRequest) HasMaxPathLength() bool`

HasMaxPathLength returns a boolean if a field has been set.

### GetNotAfter

`func (o *KmsCreateCertificateAuthorityRequest) GetNotAfter() string`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *KmsCreateCertificateAuthorityRequest) GetNotAfterOk() (*string, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *KmsCreateCertificateAuthorityRequest) SetNotAfter(v string)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *KmsCreateCertificateAuthorityRequest) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


