# O11yTrace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TraceID** | Pointer to **string** |  | [optional] 
**RootServiceName** | Pointer to **string** |  | [optional] 
**RootTraceName** | Pointer to **string** |  | [optional] 
**StartTimeUnixNano** | Pointer to **string** |  | [optional] 
**DurationMs** | Pointer to **int32** |  | [optional] 
**SpanCount** | Pointer to **int32** |  | [optional] 
**ErrorCount** | Pointer to **int32** |  | [optional] 
**Spans** | Pointer to [**[]O11ySpan**](O11ySpan.md) |  | [optional] 

## Methods

### NewO11yTrace

`func NewO11yTrace() *O11yTrace`

NewO11yTrace instantiates a new O11yTrace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yTraceWithDefaults

`func NewO11yTraceWithDefaults() *O11yTrace`

NewO11yTraceWithDefaults instantiates a new O11yTrace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceID

`func (o *O11yTrace) GetTraceID() string`

GetTraceID returns the TraceID field if non-nil, zero value otherwise.

### GetTraceIDOk

`func (o *O11yTrace) GetTraceIDOk() (*string, bool)`

GetTraceIDOk returns a tuple with the TraceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceID

`func (o *O11yTrace) SetTraceID(v string)`

SetTraceID sets TraceID field to given value.

### HasTraceID

`func (o *O11yTrace) HasTraceID() bool`

HasTraceID returns a boolean if a field has been set.

### GetRootServiceName

`func (o *O11yTrace) GetRootServiceName() string`

GetRootServiceName returns the RootServiceName field if non-nil, zero value otherwise.

### GetRootServiceNameOk

`func (o *O11yTrace) GetRootServiceNameOk() (*string, bool)`

GetRootServiceNameOk returns a tuple with the RootServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootServiceName

`func (o *O11yTrace) SetRootServiceName(v string)`

SetRootServiceName sets RootServiceName field to given value.

### HasRootServiceName

`func (o *O11yTrace) HasRootServiceName() bool`

HasRootServiceName returns a boolean if a field has been set.

### GetRootTraceName

`func (o *O11yTrace) GetRootTraceName() string`

GetRootTraceName returns the RootTraceName field if non-nil, zero value otherwise.

### GetRootTraceNameOk

`func (o *O11yTrace) GetRootTraceNameOk() (*string, bool)`

GetRootTraceNameOk returns a tuple with the RootTraceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootTraceName

`func (o *O11yTrace) SetRootTraceName(v string)`

SetRootTraceName sets RootTraceName field to given value.

### HasRootTraceName

`func (o *O11yTrace) HasRootTraceName() bool`

HasRootTraceName returns a boolean if a field has been set.

### GetStartTimeUnixNano

`func (o *O11yTrace) GetStartTimeUnixNano() string`

GetStartTimeUnixNano returns the StartTimeUnixNano field if non-nil, zero value otherwise.

### GetStartTimeUnixNanoOk

`func (o *O11yTrace) GetStartTimeUnixNanoOk() (*string, bool)`

GetStartTimeUnixNanoOk returns a tuple with the StartTimeUnixNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUnixNano

`func (o *O11yTrace) SetStartTimeUnixNano(v string)`

SetStartTimeUnixNano sets StartTimeUnixNano field to given value.

### HasStartTimeUnixNano

`func (o *O11yTrace) HasStartTimeUnixNano() bool`

HasStartTimeUnixNano returns a boolean if a field has been set.

### GetDurationMs

`func (o *O11yTrace) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *O11yTrace) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *O11yTrace) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *O11yTrace) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetSpanCount

`func (o *O11yTrace) GetSpanCount() int32`

GetSpanCount returns the SpanCount field if non-nil, zero value otherwise.

### GetSpanCountOk

`func (o *O11yTrace) GetSpanCountOk() (*int32, bool)`

GetSpanCountOk returns a tuple with the SpanCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanCount

`func (o *O11yTrace) SetSpanCount(v int32)`

SetSpanCount sets SpanCount field to given value.

### HasSpanCount

`func (o *O11yTrace) HasSpanCount() bool`

HasSpanCount returns a boolean if a field has been set.

### GetErrorCount

`func (o *O11yTrace) GetErrorCount() int32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *O11yTrace) GetErrorCountOk() (*int32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *O11yTrace) SetErrorCount(v int32)`

SetErrorCount sets ErrorCount field to given value.

### HasErrorCount

`func (o *O11yTrace) HasErrorCount() bool`

HasErrorCount returns a boolean if a field has been set.

### GetSpans

`func (o *O11yTrace) GetSpans() []O11ySpan`

GetSpans returns the Spans field if non-nil, zero value otherwise.

### GetSpansOk

`func (o *O11yTrace) GetSpansOk() (*[]O11ySpan, bool)`

GetSpansOk returns a tuple with the Spans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpans

`func (o *O11yTrace) SetSpans(v []O11ySpan)`

SetSpans sets Spans field to given value.

### HasSpans

`func (o *O11yTrace) HasSpans() bool`

HasSpans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


