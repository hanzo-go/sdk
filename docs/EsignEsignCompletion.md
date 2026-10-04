# EsignEsignCompletion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentStatus** | Pointer to **string** | DocumentStatus is COMPLETED when this was the last signature and the document sealed here, PENDING while others have still to sign. | [optional] 
**RecipientId** | Pointer to **string** | RecipientID is the recipient who finished. | [optional] 
**Sealed** | Pointer to **bool** | Sealed is whether the document sealed on this call — the field values rendered onto the PDF and a real x509 PKCS#7 signature applied. | [optional] 

## Methods

### NewEsignEsignCompletion

`func NewEsignEsignCompletion() *EsignEsignCompletion`

NewEsignEsignCompletion instantiates a new EsignEsignCompletion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignEsignCompletionWithDefaults

`func NewEsignEsignCompletionWithDefaults() *EsignEsignCompletion`

NewEsignEsignCompletionWithDefaults instantiates a new EsignEsignCompletion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentStatus

`func (o *EsignEsignCompletion) GetDocumentStatus() string`

GetDocumentStatus returns the DocumentStatus field if non-nil, zero value otherwise.

### GetDocumentStatusOk

`func (o *EsignEsignCompletion) GetDocumentStatusOk() (*string, bool)`

GetDocumentStatusOk returns a tuple with the DocumentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentStatus

`func (o *EsignEsignCompletion) SetDocumentStatus(v string)`

SetDocumentStatus sets DocumentStatus field to given value.

### HasDocumentStatus

`func (o *EsignEsignCompletion) HasDocumentStatus() bool`

HasDocumentStatus returns a boolean if a field has been set.

### GetRecipientId

`func (o *EsignEsignCompletion) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *EsignEsignCompletion) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *EsignEsignCompletion) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *EsignEsignCompletion) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### GetSealed

`func (o *EsignEsignCompletion) GetSealed() bool`

GetSealed returns the Sealed field if non-nil, zero value otherwise.

### GetSealedOk

`func (o *EsignEsignCompletion) GetSealedOk() (*bool, bool)`

GetSealedOk returns a tuple with the Sealed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealed

`func (o *EsignEsignCompletion) SetSealed(v bool)`

SetSealed sets Sealed field to given value.

### HasSealed

`func (o *EsignEsignCompletion) HasSealed() bool`

HasSealed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


