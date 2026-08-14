# O11yO11yErrorsCountIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** | End is the window end, as a nanosecond epoch spelled as a string. | [optional] 
**ExceptionType** | Pointer to **string** | ExceptionType narrows to one exception type. | [optional] 
**ServiceName** | Pointer to **string** | ServiceName narrows to one reporting service. | [optional] 
**Start** | Pointer to **string** | Start is the window start, as a nanosecond epoch spelled as a string. | [optional] 
**Tags** | Pointer to [**[]O11yO11yTagQuery**](O11yO11yTagQuery.md) | Tags narrow the scan to spans carrying the given tag values. | [optional] 

## Methods

### NewO11yO11yErrorsCountIn

`func NewO11yO11yErrorsCountIn() *O11yO11yErrorsCountIn`

NewO11yO11yErrorsCountIn instantiates a new O11yO11yErrorsCountIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yErrorsCountInWithDefaults

`func NewO11yO11yErrorsCountInWithDefaults() *O11yO11yErrorsCountIn`

NewO11yO11yErrorsCountInWithDefaults instantiates a new O11yO11yErrorsCountIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *O11yO11yErrorsCountIn) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yO11yErrorsCountIn) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yO11yErrorsCountIn) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *O11yO11yErrorsCountIn) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetExceptionType

`func (o *O11yO11yErrorsCountIn) GetExceptionType() string`

GetExceptionType returns the ExceptionType field if non-nil, zero value otherwise.

### GetExceptionTypeOk

`func (o *O11yO11yErrorsCountIn) GetExceptionTypeOk() (*string, bool)`

GetExceptionTypeOk returns a tuple with the ExceptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionType

`func (o *O11yO11yErrorsCountIn) SetExceptionType(v string)`

SetExceptionType sets ExceptionType field to given value.

### HasExceptionType

`func (o *O11yO11yErrorsCountIn) HasExceptionType() bool`

HasExceptionType returns a boolean if a field has been set.

### GetServiceName

`func (o *O11yO11yErrorsCountIn) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *O11yO11yErrorsCountIn) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *O11yO11yErrorsCountIn) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *O11yO11yErrorsCountIn) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetStart

`func (o *O11yO11yErrorsCountIn) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yO11yErrorsCountIn) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yO11yErrorsCountIn) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *O11yO11yErrorsCountIn) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTags

`func (o *O11yO11yErrorsCountIn) GetTags() []O11yO11yTagQuery`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yO11yErrorsCountIn) GetTagsOk() (*[]O11yO11yTagQuery, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yO11yErrorsCountIn) SetTags(v []O11yO11yTagQuery)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yO11yErrorsCountIn) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


