# TaxCertification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**By** | Pointer to **string** | By is who signed it: the IAM user legal recorded completing the signature. | [optional] 
**Document** | Pointer to **string** | Document is the /v1/legal document that carries the signature. | [optional] 
**Signed** | Pointer to **int64** | Signed is when legal reported the document signed, unix seconds. | [optional] 
**Signer** | Pointer to **string** | Signer is the IAM user the signature was opened for — the admin who asked. Only that user&#39;s completion certifies: a completion anyone else recorded is not their signature, and a fresh request is opened in its place. | [optional] 
**Status** | Pointer to **string** | Status is none, pending (a signature is open) or certified. | [optional] 
**Version** | Pointer to **int64** | Version is the profile version the signature covers. Changing any W-9 fact makes a new version, and the certification no longer covers it. | [optional] 

## Methods

### NewTaxCertification

`func NewTaxCertification() *TaxCertification`

NewTaxCertification instantiates a new TaxCertification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxCertificationWithDefaults

`func NewTaxCertificationWithDefaults() *TaxCertification`

NewTaxCertificationWithDefaults instantiates a new TaxCertification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBy

`func (o *TaxCertification) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *TaxCertification) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *TaxCertification) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *TaxCertification) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetDocument

`func (o *TaxCertification) GetDocument() string`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *TaxCertification) GetDocumentOk() (*string, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *TaxCertification) SetDocument(v string)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *TaxCertification) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetSigned

`func (o *TaxCertification) GetSigned() int64`

GetSigned returns the Signed field if non-nil, zero value otherwise.

### GetSignedOk

`func (o *TaxCertification) GetSignedOk() (*int64, bool)`

GetSignedOk returns a tuple with the Signed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigned

`func (o *TaxCertification) SetSigned(v int64)`

SetSigned sets Signed field to given value.

### HasSigned

`func (o *TaxCertification) HasSigned() bool`

HasSigned returns a boolean if a field has been set.

### GetSigner

`func (o *TaxCertification) GetSigner() string`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *TaxCertification) GetSignerOk() (*string, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *TaxCertification) SetSigner(v string)`

SetSigner sets Signer field to given value.

### HasSigner

`func (o *TaxCertification) HasSigner() bool`

HasSigner returns a boolean if a field has been set.

### GetStatus

`func (o *TaxCertification) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaxCertification) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaxCertification) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaxCertification) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *TaxCertification) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TaxCertification) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TaxCertification) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TaxCertification) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


