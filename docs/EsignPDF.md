# EsignPDF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filename** | Pointer to **string** | Filename is the name to save it under, built from the title and marked _signed once it is sealed. | [optional] 
**Id** | Pointer to **string** | ID is the document this PDF was rendered from. | [optional] 
**PdfBase64** | Pointer to **string** | PdfBase64 is the PDF itself, base64-encoded. There is one field either way, so Sealed is what tells you which artifact you are holding. | [optional] 
**Sealed** | Pointer to **bool** | Sealed is whether this is the SEALED artifact — the field values rendered onto the page and a real x509 PKCS#7 signature applied — rather than the original upload. | [optional] 
**Status** | Pointer to **string** | Status is the document&#39;s state at the moment it was read. | [optional] 

## Methods

### NewEsignPDF

`func NewEsignPDF() *EsignPDF`

NewEsignPDF instantiates a new EsignPDF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignPDFWithDefaults

`func NewEsignPDFWithDefaults() *EsignPDF`

NewEsignPDFWithDefaults instantiates a new EsignPDF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilename

`func (o *EsignPDF) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *EsignPDF) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *EsignPDF) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *EsignPDF) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetId

`func (o *EsignPDF) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EsignPDF) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EsignPDF) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EsignPDF) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPdfBase64

`func (o *EsignPDF) GetPdfBase64() string`

GetPdfBase64 returns the PdfBase64 field if non-nil, zero value otherwise.

### GetPdfBase64Ok

`func (o *EsignPDF) GetPdfBase64Ok() (*string, bool)`

GetPdfBase64Ok returns a tuple with the PdfBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfBase64

`func (o *EsignPDF) SetPdfBase64(v string)`

SetPdfBase64 sets PdfBase64 field to given value.

### HasPdfBase64

`func (o *EsignPDF) HasPdfBase64() bool`

HasPdfBase64 returns a boolean if a field has been set.

### GetSealed

`func (o *EsignPDF) GetSealed() bool`

GetSealed returns the Sealed field if non-nil, zero value otherwise.

### GetSealedOk

`func (o *EsignPDF) GetSealedOk() (*bool, bool)`

GetSealedOk returns a tuple with the Sealed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealed

`func (o *EsignPDF) SetSealed(v bool)`

SetSealed sets Sealed field to given value.

### HasSealed

`func (o *EsignPDF) HasSealed() bool`

HasSealed returns a boolean if a field has been set.

### GetStatus

`func (o *EsignPDF) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EsignPDF) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EsignPDF) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EsignPDF) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


