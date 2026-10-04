# GraphGraphReadOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assertions** | Pointer to [**[]GraphWireFact**](GraphWireFact.md) | Assertions are the matching rows: a read&#39;s in the order they were written, oldest first; a search&#39;s best match first. Every version is here: this resolves nothing and withholds nothing, so a superseded claim and the one that superseded it both appear. | [optional] 
**Truncated** | Pointer to **bool** | Truncated says more assertions matched than came back: the limit, or the ceiling, stopped the read. Narrow it, or page by as_known, to see the rest; a search that stopped kept the best matches. | [optional] 

## Methods

### NewGraphGraphReadOut

`func NewGraphGraphReadOut() *GraphGraphReadOut`

NewGraphGraphReadOut instantiates a new GraphGraphReadOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphReadOutWithDefaults

`func NewGraphGraphReadOutWithDefaults() *GraphGraphReadOut`

NewGraphGraphReadOutWithDefaults instantiates a new GraphGraphReadOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssertions

`func (o *GraphGraphReadOut) GetAssertions() []GraphWireFact`

GetAssertions returns the Assertions field if non-nil, zero value otherwise.

### GetAssertionsOk

`func (o *GraphGraphReadOut) GetAssertionsOk() (*[]GraphWireFact, bool)`

GetAssertionsOk returns a tuple with the Assertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertions

`func (o *GraphGraphReadOut) SetAssertions(v []GraphWireFact)`

SetAssertions sets Assertions field to given value.

### HasAssertions

`func (o *GraphGraphReadOut) HasAssertions() bool`

HasAssertions returns a boolean if a field has been set.

### GetTruncated

`func (o *GraphGraphReadOut) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *GraphGraphReadOut) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *GraphGraphReadOut) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *GraphGraphReadOut) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


