# KmsIssueCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FriendlyName** | **string** |  | 
**CommonName** | **string** |  | 
**AltNames** | Pointer to **string** | Comma-separated SANs | [optional] 
**Ttl** | Pointer to **string** | Duration string (e.g., 365d, 1y) | [optional] 

## Methods

### NewKmsIssueCertificateRequest

`func NewKmsIssueCertificateRequest(friendlyName string, commonName string, ) *KmsIssueCertificateRequest`

NewKmsIssueCertificateRequest instantiates a new KmsIssueCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsIssueCertificateRequestWithDefaults

`func NewKmsIssueCertificateRequestWithDefaults() *KmsIssueCertificateRequest`

NewKmsIssueCertificateRequestWithDefaults instantiates a new KmsIssueCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFriendlyName

`func (o *KmsIssueCertificateRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *KmsIssueCertificateRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *KmsIssueCertificateRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.


### GetCommonName

`func (o *KmsIssueCertificateRequest) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *KmsIssueCertificateRequest) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *KmsIssueCertificateRequest) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### GetAltNames

`func (o *KmsIssueCertificateRequest) GetAltNames() string`

GetAltNames returns the AltNames field if non-nil, zero value otherwise.

### GetAltNamesOk

`func (o *KmsIssueCertificateRequest) GetAltNamesOk() (*string, bool)`

GetAltNamesOk returns a tuple with the AltNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNames

`func (o *KmsIssueCertificateRequest) SetAltNames(v string)`

SetAltNames sets AltNames field to given value.

### HasAltNames

`func (o *KmsIssueCertificateRequest) HasAltNames() bool`

HasAltNames returns a boolean if a field has been set.

### GetTtl

`func (o *KmsIssueCertificateRequest) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *KmsIssueCertificateRequest) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *KmsIssueCertificateRequest) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *KmsIssueCertificateRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


