# O11yO11yDependencyGraphIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **string** | End is the window end, as epoch nanoseconds. Required. | 
**Start** | **string** | Start is the window start, as epoch nanoseconds. Required. | 
**Tags** | Pointer to [**[]O11yO11yTagFilter**](O11yO11yTagFilter.md) | Tags narrow the graph to spans matching every condition. | [optional] 

## Methods

### NewO11yO11yDependencyGraphIn

`func NewO11yO11yDependencyGraphIn(end string, start string, ) *O11yO11yDependencyGraphIn`

NewO11yO11yDependencyGraphIn instantiates a new O11yO11yDependencyGraphIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yDependencyGraphInWithDefaults

`func NewO11yO11yDependencyGraphInWithDefaults() *O11yO11yDependencyGraphIn`

NewO11yO11yDependencyGraphInWithDefaults instantiates a new O11yO11yDependencyGraphIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *O11yO11yDependencyGraphIn) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yO11yDependencyGraphIn) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yO11yDependencyGraphIn) SetEnd(v string)`

SetEnd sets End field to given value.


### GetStart

`func (o *O11yO11yDependencyGraphIn) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yO11yDependencyGraphIn) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yO11yDependencyGraphIn) SetStart(v string)`

SetStart sets Start field to given value.


### GetTags

`func (o *O11yO11yDependencyGraphIn) GetTags() []O11yO11yTagFilter`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yO11yDependencyGraphIn) GetTagsOk() (*[]O11yO11yTagFilter, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yO11yDependencyGraphIn) SetTags(v []O11yO11yTagFilter)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yO11yDependencyGraphIn) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


