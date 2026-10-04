# GraphGraphAssertIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assertions** | [**[]GraphGraphFact**](GraphGraphFact.md) | Assertions is the batch. Each member is judged on its own: one refusal does not discard the rest, because a caller redelivering five facts must not lose four of them to one malformed fifth. | 

## Methods

### NewGraphGraphAssertIn

`func NewGraphGraphAssertIn(assertions []GraphGraphFact, ) *GraphGraphAssertIn`

NewGraphGraphAssertIn instantiates a new GraphGraphAssertIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphAssertInWithDefaults

`func NewGraphGraphAssertInWithDefaults() *GraphGraphAssertIn`

NewGraphGraphAssertInWithDefaults instantiates a new GraphGraphAssertIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssertions

`func (o *GraphGraphAssertIn) GetAssertions() []GraphGraphFact`

GetAssertions returns the Assertions field if non-nil, zero value otherwise.

### GetAssertionsOk

`func (o *GraphGraphAssertIn) GetAssertionsOk() (*[]GraphGraphFact, bool)`

GetAssertionsOk returns a tuple with the Assertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertions

`func (o *GraphGraphAssertIn) SetAssertions(v []GraphGraphFact)`

SetAssertions sets Assertions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


