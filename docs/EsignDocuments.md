# EsignDocuments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Documents** | Pointer to [**[]EsignDocument**](EsignDocument.md) | Documents is the caller org&#39;s documents, newest first, capped at 200. There is no paging, so read it as the recent window rather than a complete export. | [optional] 

## Methods

### NewEsignDocuments

`func NewEsignDocuments() *EsignDocuments`

NewEsignDocuments instantiates a new EsignDocuments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignDocumentsWithDefaults

`func NewEsignDocumentsWithDefaults() *EsignDocuments`

NewEsignDocumentsWithDefaults instantiates a new EsignDocuments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocuments

`func (o *EsignDocuments) GetDocuments() []EsignDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *EsignDocuments) GetDocumentsOk() (*[]EsignDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *EsignDocuments) SetDocuments(v []EsignDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *EsignDocuments) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


