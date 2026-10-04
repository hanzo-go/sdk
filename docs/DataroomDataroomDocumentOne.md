# DataroomDataroomDocumentOne

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Document** | Pointer to [**DataroomDataroomDocument**](DataroomDataroomDocument.md) | Document is the requested document&#39;s METADATA. Its bytes are a separate read, GET /v1/dataroom/documents/{id}/file. | [optional] 

## Methods

### NewDataroomDataroomDocumentOne

`func NewDataroomDataroomDocumentOne() *DataroomDataroomDocumentOne`

NewDataroomDataroomDocumentOne instantiates a new DataroomDataroomDocumentOne object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataroomDataroomDocumentOneWithDefaults

`func NewDataroomDataroomDocumentOneWithDefaults() *DataroomDataroomDocumentOne`

NewDataroomDataroomDocumentOneWithDefaults instantiates a new DataroomDataroomDocumentOne object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocument

`func (o *DataroomDataroomDocumentOne) GetDocument() DataroomDataroomDocument`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *DataroomDataroomDocumentOne) GetDocumentOk() (*DataroomDataroomDocument, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *DataroomDataroomDocumentOne) SetDocument(v DataroomDataroomDocument)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *DataroomDataroomDocumentOne) HasDocument() bool`

HasDocument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


