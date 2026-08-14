# O11yO11yOperationsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** | End is the window&#39;s end, epoch nanoseconds as a string. | [optional] 
**Limit** | Pointer to **int32** | Limit caps how many operations come back. | [optional] 
**Service** | Pointer to **string** | Service is the service whose operations are read. | [optional] 
**Start** | Pointer to **string** | Start is the window&#39;s start, epoch nanoseconds as a string. | [optional] 
**Tags** | Pointer to [**[]O11yO11yServiceTag**](O11yO11yServiceTag.md) | Tags narrow the spans counted, each a span-attribute predicate. | [optional] 

## Methods

### NewO11yO11yOperationsIn

`func NewO11yO11yOperationsIn() *O11yO11yOperationsIn`

NewO11yO11yOperationsIn instantiates a new O11yO11yOperationsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yOperationsInWithDefaults

`func NewO11yO11yOperationsInWithDefaults() *O11yO11yOperationsIn`

NewO11yO11yOperationsInWithDefaults instantiates a new O11yO11yOperationsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *O11yO11yOperationsIn) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yO11yOperationsIn) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yO11yOperationsIn) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *O11yO11yOperationsIn) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetLimit

`func (o *O11yO11yOperationsIn) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *O11yO11yOperationsIn) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *O11yO11yOperationsIn) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *O11yO11yOperationsIn) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetService

`func (o *O11yO11yOperationsIn) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *O11yO11yOperationsIn) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *O11yO11yOperationsIn) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *O11yO11yOperationsIn) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStart

`func (o *O11yO11yOperationsIn) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yO11yOperationsIn) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yO11yOperationsIn) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *O11yO11yOperationsIn) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTags

`func (o *O11yO11yOperationsIn) GetTags() []O11yO11yServiceTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yO11yOperationsIn) GetTagsOk() (*[]O11yO11yServiceTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yO11yOperationsIn) SetTags(v []O11yO11yServiceTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yO11yOperationsIn) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


