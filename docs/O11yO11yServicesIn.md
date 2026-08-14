# O11yO11yServicesIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** | End is the window&#39;s end, epoch nanoseconds as a string. | [optional] 
**Start** | Pointer to **string** | Start is the window&#39;s start, epoch nanoseconds as a string. | [optional] 
**Tags** | Pointer to [**[]O11yO11yServiceTag**](O11yO11yServiceTag.md) | Tags narrow the spans counted, each a span-attribute predicate. | [optional] 

## Methods

### NewO11yO11yServicesIn

`func NewO11yO11yServicesIn() *O11yO11yServicesIn`

NewO11yO11yServicesIn instantiates a new O11yO11yServicesIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yServicesInWithDefaults

`func NewO11yO11yServicesInWithDefaults() *O11yO11yServicesIn`

NewO11yO11yServicesInWithDefaults instantiates a new O11yO11yServicesIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *O11yO11yServicesIn) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yO11yServicesIn) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yO11yServicesIn) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *O11yO11yServicesIn) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStart

`func (o *O11yO11yServicesIn) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yO11yServicesIn) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yO11yServicesIn) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *O11yO11yServicesIn) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTags

`func (o *O11yO11yServicesIn) GetTags() []O11yO11yServiceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yO11yServicesIn) GetTagsOk() (*[]O11yO11yServiceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yO11yServicesIn) SetTags(v []O11yO11yServiceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yO11yServicesIn) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


