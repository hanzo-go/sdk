# GraphGraphExtractOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Refused** | Pointer to [**[]GraphGraphTriple**](GraphGraphTriple.md) | Refused are the relations found that filing would refuse — a subject or object whose type is not the relation&#39;s declared domain or range, or a declaration from a caller who is not an admin — each carrying its reason. Ingest records none of them. | [optional] 
**Triples** | Pointer to [**[]GraphGraphTriple**](GraphGraphTriple.md) | Triples are the relations found that the caller may file, in the order the document states them. | [optional] 

## Methods

### NewGraphGraphExtractOut

`func NewGraphGraphExtractOut() *GraphGraphExtractOut`

NewGraphGraphExtractOut instantiates a new GraphGraphExtractOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphExtractOutWithDefaults

`func NewGraphGraphExtractOutWithDefaults() *GraphGraphExtractOut`

NewGraphGraphExtractOutWithDefaults instantiates a new GraphGraphExtractOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefused

`func (o *GraphGraphExtractOut) GetRefused() []GraphGraphTriple`

GetRefused returns the Refused field if non-nil, zero value otherwise.

### GetRefusedOk

`func (o *GraphGraphExtractOut) GetRefusedOk() (*[]GraphGraphTriple, bool)`

GetRefusedOk returns a tuple with the Refused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefused

`func (o *GraphGraphExtractOut) SetRefused(v []GraphGraphTriple)`

SetRefused sets Refused field to given value.

### HasRefused

`func (o *GraphGraphExtractOut) HasRefused() bool`

HasRefused returns a boolean if a field has been set.

### GetTriples

`func (o *GraphGraphExtractOut) GetTriples() []GraphGraphTriple`

GetTriples returns the Triples field if non-nil, zero value otherwise.

### GetTriplesOk

`func (o *GraphGraphExtractOut) GetTriplesOk() (*[]GraphGraphTriple, bool)`

GetTriplesOk returns a tuple with the Triples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriples

`func (o *GraphGraphExtractOut) SetTriples(v []GraphGraphTriple)`

SetTriples sets Triples field to given value.

### HasTriples

`func (o *GraphGraphExtractOut) HasTriples() bool`

HasTriples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


