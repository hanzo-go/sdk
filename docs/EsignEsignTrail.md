# EsignEsignTrail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | Pointer to **string** | DocumentID is the document the trail belongs to. | [optional] 
**Entries** | Pointer to [**[]EsignEsignEvent**](EsignEsignEvent.md) | Entries is every recorded event in order, oldest first. It is append-only: nothing in this surface edits or removes an entry. | [optional] 

## Methods

### NewEsignEsignTrail

`func NewEsignEsignTrail() *EsignEsignTrail`

NewEsignEsignTrail instantiates a new EsignEsignTrail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignEsignTrailWithDefaults

`func NewEsignEsignTrailWithDefaults() *EsignEsignTrail`

NewEsignEsignTrailWithDefaults instantiates a new EsignEsignTrail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *EsignEsignTrail) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *EsignEsignTrail) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *EsignEsignTrail) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *EsignEsignTrail) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetEntries

`func (o *EsignEsignTrail) GetEntries() []EsignEsignEvent`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *EsignEsignTrail) GetEntriesOk() (*[]EsignEsignEvent, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *EsignEsignTrail) SetEntries(v []EsignEsignEvent)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *EsignEsignTrail) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


