# KmsListCertificates200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificates** | Pointer to [**[]KmsCertificate**](KmsCertificate.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewKmsListCertificates200Response

`func NewKmsListCertificates200Response() *KmsListCertificates200Response`

NewKmsListCertificates200Response instantiates a new KmsListCertificates200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmsListCertificates200ResponseWithDefaults

`func NewKmsListCertificates200ResponseWithDefaults() *KmsListCertificates200Response`

NewKmsListCertificates200ResponseWithDefaults instantiates a new KmsListCertificates200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificates

`func (o *KmsListCertificates200Response) GetCertificates() []KmsCertificate`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *KmsListCertificates200Response) GetCertificatesOk() (*[]KmsCertificate, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *KmsListCertificates200Response) SetCertificates(v []KmsCertificate)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *KmsListCertificates200Response) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetTotalCount

`func (o *KmsListCertificates200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *KmsListCertificates200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *KmsListCertificates200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *KmsListCertificates200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


