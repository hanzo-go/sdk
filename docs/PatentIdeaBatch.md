# PatentIdeaBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ideas** | Pointer to [**[]PatentPatentIdea**](PatentPatentIdea.md) | Ideas are the entries to append, at most 2000. | [optional] 

## Methods

### NewPatentIdeaBatch

`func NewPatentIdeaBatch() *PatentIdeaBatch`

NewPatentIdeaBatch instantiates a new PatentIdeaBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatentIdeaBatchWithDefaults

`func NewPatentIdeaBatchWithDefaults() *PatentIdeaBatch`

NewPatentIdeaBatchWithDefaults instantiates a new PatentIdeaBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdeas

`func (o *PatentIdeaBatch) GetIdeas() []PatentPatentIdea`

GetIdeas returns the Ideas field if non-nil, zero value otherwise.

### GetIdeasOk

`func (o *PatentIdeaBatch) GetIdeasOk() (*[]PatentPatentIdea, bool)`

GetIdeasOk returns a tuple with the Ideas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdeas

`func (o *PatentIdeaBatch) SetIdeas(v []PatentPatentIdea)`

SetIdeas sets Ideas field to given value.

### HasIdeas

`func (o *PatentIdeaBatch) HasIdeas() bool`

HasIdeas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


