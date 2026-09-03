# O11yO11yErrorsListIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** | End is the window end, as a nanosecond epoch spelled as a string. | [optional] 
**ExceptionType** | Pointer to **string** | ExceptionType narrows to one exception type. | [optional] 
**Limit** | Pointer to **int64** | Limit caps how many exception groups come back. Required, non-zero. | [optional] 
**Offset** | Pointer to **int64** | Offset is how many groups to skip. | [optional] 
**Order** | Pointer to **string** | Order is the direction: ascending or descending. | [optional] 
**OrderParam** | Pointer to **string** | OrderParam is the column to order by, e.g. exceptionCount, lastSeen. | [optional] 
**ServiceName** | Pointer to **string** | ServiceName narrows to one reporting service. | [optional] 
**Start** | Pointer to **string** | Start is the window start, as a nanosecond epoch spelled as a string. | [optional] 
**Tags** | Pointer to [**[]O11yO11yTagQuery**](O11yO11yTagQuery.md) | Tags narrow the scan to spans carrying the given tag values. | [optional] 

## Methods

### NewO11yO11yErrorsListIn

`func NewO11yO11yErrorsListIn() *O11yO11yErrorsListIn`

NewO11yO11yErrorsListIn instantiates a new O11yO11yErrorsListIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yErrorsListInWithDefaults

`func NewO11yO11yErrorsListInWithDefaults() *O11yO11yErrorsListIn`

NewO11yO11yErrorsListInWithDefaults instantiates a new O11yO11yErrorsListIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *O11yO11yErrorsListIn) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yO11yErrorsListIn) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yO11yErrorsListIn) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *O11yO11yErrorsListIn) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetExceptionType

`func (o *O11yO11yErrorsListIn) GetExceptionType() string`

GetExceptionType returns the ExceptionType field if non-nil, zero value otherwise.

### GetExceptionTypeOk

`func (o *O11yO11yErrorsListIn) GetExceptionTypeOk() (*string, bool)`

GetExceptionTypeOk returns a tuple with the ExceptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionType

`func (o *O11yO11yErrorsListIn) SetExceptionType(v string)`

SetExceptionType sets ExceptionType field to given value.

### HasExceptionType

`func (o *O11yO11yErrorsListIn) HasExceptionType() bool`

HasExceptionType returns a boolean if a field has been set.

### GetLimit

`func (o *O11yO11yErrorsListIn) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *O11yO11yErrorsListIn) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *O11yO11yErrorsListIn) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *O11yO11yErrorsListIn) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *O11yO11yErrorsListIn) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *O11yO11yErrorsListIn) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *O11yO11yErrorsListIn) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *O11yO11yErrorsListIn) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrder

`func (o *O11yO11yErrorsListIn) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *O11yO11yErrorsListIn) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *O11yO11yErrorsListIn) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *O11yO11yErrorsListIn) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOrderParam

`func (o *O11yO11yErrorsListIn) GetOrderParam() string`

GetOrderParam returns the OrderParam field if non-nil, zero value otherwise.

### GetOrderParamOk

`func (o *O11yO11yErrorsListIn) GetOrderParamOk() (*string, bool)`

GetOrderParamOk returns a tuple with the OrderParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderParam

`func (o *O11yO11yErrorsListIn) SetOrderParam(v string)`

SetOrderParam sets OrderParam field to given value.

### HasOrderParam

`func (o *O11yO11yErrorsListIn) HasOrderParam() bool`

HasOrderParam returns a boolean if a field has been set.

### GetServiceName

`func (o *O11yO11yErrorsListIn) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *O11yO11yErrorsListIn) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *O11yO11yErrorsListIn) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *O11yO11yErrorsListIn) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetStart

`func (o *O11yO11yErrorsListIn) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yO11yErrorsListIn) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yO11yErrorsListIn) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *O11yO11yErrorsListIn) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTags

`func (o *O11yO11yErrorsListIn) GetTags() []O11yO11yTagQuery`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *O11yO11yErrorsListIn) GetTagsOk() (*[]O11yO11yTagQuery, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *O11yO11yErrorsListIn) SetTags(v []O11yO11yTagQuery)`

SetTags sets Tags field to given value.

### HasTags

`func (o *O11yO11yErrorsListIn) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


