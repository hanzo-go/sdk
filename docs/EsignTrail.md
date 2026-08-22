# EsignTrail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | Pointer to **string** | DocumentID is the document the trail belongs to. | [optional] 
**Entries** | Pointer to [**[]EsignEvent**](EsignEvent.md) | Entries is every recorded event in order, oldest first. It is append-only: nothing in this surface edits or removes an entry. | [optional] 

## Methods

### NewEsignTrail

`func NewEsignTrail() *EsignTrail`

NewEsignTrail instantiates a new EsignTrail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsignTrailWithDefaults

`func NewEsignTrailWithDefaults() *EsignTrail`

NewEsignTrailWithDefaults instantiates a new EsignTrail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *EsignTrail) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *EsignTrail) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *EsignTrail) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *EsignTrail) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetEntries

`func (o *EsignTrail) GetEntries() []EsignEvent`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *EsignTrail) GetEntriesOk() (*[]EsignEvent, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *EsignTrail) SetEntries(v []EsignEvent)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *EsignTrail) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


