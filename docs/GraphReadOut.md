# GraphReadOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assertions** | Pointer to [**[]WireFact**](WireFact.md) | Assertions are the matching rows in the order they were written, oldest first. Every version is here: this read resolves nothing and withholds nothing, so a superseded claim and the one that superseded it both appear. | [optional] 

## Methods

### NewGraphReadOut

`func NewGraphReadOut() *GraphReadOut`

NewGraphReadOut instantiates a new GraphReadOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphReadOutWithDefaults

`func NewGraphReadOutWithDefaults() *GraphReadOut`

NewGraphReadOutWithDefaults instantiates a new GraphReadOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssertions

`func (o *GraphReadOut) GetAssertions() []WireFact`

GetAssertions returns the Assertions field if non-nil, zero value otherwise.

### GetAssertionsOk

`func (o *GraphReadOut) GetAssertionsOk() (*[]WireFact, bool)`

GetAssertionsOk returns a tuple with the Assertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertions

`func (o *GraphReadOut) SetAssertions(v []WireFact)`

SetAssertions sets Assertions field to given value.

### HasAssertions

`func (o *GraphReadOut) HasAssertions() bool`

HasAssertions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


