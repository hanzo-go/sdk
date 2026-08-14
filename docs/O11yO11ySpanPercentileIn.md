# O11yO11ySpanPercentileIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int32** | End is the window end, as epoch nanoseconds. | [optional] 
**Name** | **string** | Name is the span name whose peers are compared. Required. | 
**ResourceAttributes** | Pointer to **map[string]string** | ResourceAttributes narrow the peer group to spans carrying them all. | [optional] 
**ServiceName** | **string** | ServiceName is the service the span belongs to. Required. | 
**SpanDuration** | Pointer to **int32** | SpanDuration is the span&#39;s duration in nanoseconds. | [optional] 
**Start** | Pointer to **int32** | Start is the window start, as epoch nanoseconds. | [optional] 

## Methods

### NewO11yO11ySpanPercentileIn

`func NewO11yO11ySpanPercentileIn(name string, serviceName string, ) *O11yO11ySpanPercentileIn`

NewO11yO11ySpanPercentileIn instantiates a new O11yO11ySpanPercentileIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11ySpanPercentileInWithDefaults

`func NewO11yO11ySpanPercentileInWithDefaults() *O11yO11ySpanPercentileIn`

NewO11yO11ySpanPercentileInWithDefaults instantiates a new O11yO11ySpanPercentileIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *O11yO11ySpanPercentileIn) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yO11ySpanPercentileIn) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yO11ySpanPercentileIn) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *O11yO11ySpanPercentileIn) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetName

`func (o *O11yO11ySpanPercentileIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11ySpanPercentileIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11ySpanPercentileIn) SetName(v string)`

SetName sets Name field to given value.


### GetResourceAttributes

`func (o *O11yO11ySpanPercentileIn) GetResourceAttributes() map[string]string`

GetResourceAttributes returns the ResourceAttributes field if non-nil, zero value otherwise.

### GetResourceAttributesOk

`func (o *O11yO11ySpanPercentileIn) GetResourceAttributesOk() (*map[string]string, bool)`

GetResourceAttributesOk returns a tuple with the ResourceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttributes

`func (o *O11yO11ySpanPercentileIn) SetResourceAttributes(v map[string]string)`

SetResourceAttributes sets ResourceAttributes field to given value.

### HasResourceAttributes

`func (o *O11yO11ySpanPercentileIn) HasResourceAttributes() bool`

HasResourceAttributes returns a boolean if a field has been set.

### GetServiceName

`func (o *O11yO11ySpanPercentileIn) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *O11yO11ySpanPercentileIn) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *O11yO11ySpanPercentileIn) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetSpanDuration

`func (o *O11yO11ySpanPercentileIn) GetSpanDuration() int32`

GetSpanDuration returns the SpanDuration field if non-nil, zero value otherwise.

### GetSpanDurationOk

`func (o *O11yO11ySpanPercentileIn) GetSpanDurationOk() (*int32, bool)`

GetSpanDurationOk returns a tuple with the SpanDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanDuration

`func (o *O11yO11ySpanPercentileIn) SetSpanDuration(v int32)`

SetSpanDuration sets SpanDuration field to given value.

### HasSpanDuration

`func (o *O11yO11ySpanPercentileIn) HasSpanDuration() bool`

HasSpanDuration returns a boolean if a field has been set.

### GetStart

`func (o *O11yO11ySpanPercentileIn) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yO11ySpanPercentileIn) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yO11ySpanPercentileIn) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *O11yO11ySpanPercentileIn) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


