# EsignEsignSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Document** | Pointer to [**EsignEsignState**](EsignEsignState.md) | Document is what is being signed. | [optional] 
**Fields** | Pointer to [**[]EsignEsignField**](EsignEsignField.md) | Fields is only the fields this recipient must fill — never another party&#39;s, so the layout a signer sees cannot reveal what anyone else was asked for. | [optional] 
**PdfBase64** | Pointer to **string** | PdfBase64 is the PDF to display, base64-encoded. Null when the document&#39;s stored bytes are missing. | [optional] 
**Recipient** | Pointer to [**EsignEsignSigner**](EsignEsignSigner.md) | Recipient is who the token says you are. | [optional] 

## Methods

### NewEsignEsignSession

`func NewEsignEsignSession() *EsignEsignSession`

NewEsignEsignSession instantiates a new EsignEsignSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignEsignSessionWithDefaults

`func NewEsignEsignSessionWithDefaults() *EsignEsignSession`

NewEsignEsignSessionWithDefaults instantiates a new EsignEsignSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocument

`func (o *EsignEsignSession) GetDocument() EsignEsignState`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *EsignEsignSession) GetDocumentOk() (*EsignEsignState, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *EsignEsignSession) SetDocument(v EsignEsignState)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *EsignEsignSession) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### GetFields

`func (o *EsignEsignSession) GetFields() []EsignEsignField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *EsignEsignSession) GetFieldsOk() (*[]EsignEsignField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *EsignEsignSession) SetFields(v []EsignEsignField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *EsignEsignSession) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetPdfBase64

`func (o *EsignEsignSession) GetPdfBase64() string`

GetPdfBase64 returns the PdfBase64 field if non-nil, zero value otherwise.

### GetPdfBase64Ok

`func (o *EsignEsignSession) GetPdfBase64Ok() (*string, bool)`

GetPdfBase64Ok returns a tuple with the PdfBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfBase64

`func (o *EsignEsignSession) SetPdfBase64(v string)`

SetPdfBase64 sets PdfBase64 field to given value.

### HasPdfBase64

`func (o *EsignEsignSession) HasPdfBase64() bool`

HasPdfBase64 returns a boolean if a field has been set.

### GetRecipient

`func (o *EsignEsignSession) GetRecipient() EsignEsignSigner`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *EsignEsignSession) GetRecipientOk() (*EsignEsignSigner, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *EsignEsignSession) SetRecipient(v EsignEsignSigner)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *EsignEsignSession) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


