# TrustDocuments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Documents** | Pointer to [**[]DocRow**](DocRow.md) | Documents is the list; a gated entry carries no address. | [optional] 

## Methods

### NewTrustDocuments

`func NewTrustDocuments() *TrustDocuments`

NewTrustDocuments instantiates a new TrustDocuments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustDocumentsWithDefaults

`func NewTrustDocumentsWithDefaults() *TrustDocuments`

NewTrustDocumentsWithDefaults instantiates a new TrustDocuments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocuments

`func (o *TrustDocuments) GetDocuments() []DocRow`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *TrustDocuments) GetDocumentsOk() (*[]DocRow, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *TrustDocuments) SetDocuments(v []DocRow)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *TrustDocuments) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


