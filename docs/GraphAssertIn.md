# GraphAssertIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assertions** | Pointer to [**[]GraphFact**](GraphFact.md) | Assertions is the batch. Each member is judged on its own: one refusal does not discard the rest, because a caller redelivering five facts must not lose four of them to one malformed fifth. | [optional] 

## Methods

### NewGraphAssertIn

`func NewGraphAssertIn() *GraphAssertIn`

NewGraphAssertIn instantiates a new GraphAssertIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphAssertInWithDefaults

`func NewGraphAssertInWithDefaults() *GraphAssertIn`

NewGraphAssertInWithDefaults instantiates a new GraphAssertIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssertions

`func (o *GraphAssertIn) GetAssertions() []GraphFact`

GetAssertions returns the Assertions field if non-nil, zero value otherwise.

### GetAssertionsOk

`func (o *GraphAssertIn) GetAssertionsOk() (*[]GraphFact, bool)`

GetAssertionsOk returns a tuple with the Assertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertions

`func (o *GraphAssertIn) SetAssertions(v []GraphFact)`

SetAssertions sets Assertions field to given value.

### HasAssertions

`func (o *GraphAssertIn) HasAssertions() bool`

HasAssertions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


