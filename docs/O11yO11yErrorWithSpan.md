# O11yO11yErrorWithSpan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorId** | Pointer to **string** | ErrorID is the exception instance id. | [optional] 
**ExceptionEscaped** | Pointer to **bool** | ExceptionEscaped marks an exception that escaped its span uncaught. | [optional] 
**ExceptionMessage** | Pointer to **string** | ExceptionMsg is the exception&#39;s message. | [optional] 
**ExceptionStacktrace** | Pointer to **string** | ExceptionStacktrace is the captured stack trace. | [optional] 
**ExceptionType** | Pointer to **string** | ExceptionType is the exception&#39;s type. | [optional] 
**GroupID** | Pointer to **string** | GroupID is the exception group it belongs to. | [optional] 
**ServiceName** | Pointer to **string** | ServiceName is the service that reported it. | [optional] 
**SpanID** | Pointer to **string** | SpanID is the span it happened on. | [optional] 
**Timestamp** | Pointer to **time.Time** | Timestamp is when it happened. | [optional] 
**TraceID** | Pointer to **string** | TraceID is the trace the span belonged to. | [optional] 

## Methods

### NewO11yO11yErrorWithSpan

`func NewO11yO11yErrorWithSpan() *O11yO11yErrorWithSpan`

NewO11yO11yErrorWithSpan instantiates a new O11yO11yErrorWithSpan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yErrorWithSpanWithDefaults

`func NewO11yO11yErrorWithSpanWithDefaults() *O11yO11yErrorWithSpan`

NewO11yO11yErrorWithSpanWithDefaults instantiates a new O11yO11yErrorWithSpan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorId

`func (o *O11yO11yErrorWithSpan) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *O11yO11yErrorWithSpan) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *O11yO11yErrorWithSpan) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *O11yO11yErrorWithSpan) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.

### GetExceptionEscaped

`func (o *O11yO11yErrorWithSpan) GetExceptionEscaped() bool`

GetExceptionEscaped returns the ExceptionEscaped field if non-nil, zero value otherwise.

### GetExceptionEscapedOk

`func (o *O11yO11yErrorWithSpan) GetExceptionEscapedOk() (*bool, bool)`

GetExceptionEscapedOk returns a tuple with the ExceptionEscaped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionEscaped

`func (o *O11yO11yErrorWithSpan) SetExceptionEscaped(v bool)`

SetExceptionEscaped sets ExceptionEscaped field to given value.

### HasExceptionEscaped

`func (o *O11yO11yErrorWithSpan) HasExceptionEscaped() bool`

HasExceptionEscaped returns a boolean if a field has been set.

### GetExceptionMessage

`func (o *O11yO11yErrorWithSpan) GetExceptionMessage() string`

GetExceptionMessage returns the ExceptionMessage field if non-nil, zero value otherwise.

### GetExceptionMessageOk

`func (o *O11yO11yErrorWithSpan) GetExceptionMessageOk() (*string, bool)`

GetExceptionMessageOk returns a tuple with the ExceptionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionMessage

`func (o *O11yO11yErrorWithSpan) SetExceptionMessage(v string)`

SetExceptionMessage sets ExceptionMessage field to given value.

### HasExceptionMessage

`func (o *O11yO11yErrorWithSpan) HasExceptionMessage() bool`

HasExceptionMessage returns a boolean if a field has been set.

### GetExceptionStacktrace

`func (o *O11yO11yErrorWithSpan) GetExceptionStacktrace() string`

GetExceptionStacktrace returns the ExceptionStacktrace field if non-nil, zero value otherwise.

### GetExceptionStacktraceOk

`func (o *O11yO11yErrorWithSpan) GetExceptionStacktraceOk() (*string, bool)`

GetExceptionStacktraceOk returns a tuple with the ExceptionStacktrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionStacktrace

`func (o *O11yO11yErrorWithSpan) SetExceptionStacktrace(v string)`

SetExceptionStacktrace sets ExceptionStacktrace field to given value.

### HasExceptionStacktrace

`func (o *O11yO11yErrorWithSpan) HasExceptionStacktrace() bool`

HasExceptionStacktrace returns a boolean if a field has been set.

### GetExceptionType

`func (o *O11yO11yErrorWithSpan) GetExceptionType() string`

GetExceptionType returns the ExceptionType field if non-nil, zero value otherwise.

### GetExceptionTypeOk

`func (o *O11yO11yErrorWithSpan) GetExceptionTypeOk() (*string, bool)`

GetExceptionTypeOk returns a tuple with the ExceptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionType

`func (o *O11yO11yErrorWithSpan) SetExceptionType(v string)`

SetExceptionType sets ExceptionType field to given value.

### HasExceptionType

`func (o *O11yO11yErrorWithSpan) HasExceptionType() bool`

HasExceptionType returns a boolean if a field has been set.

### GetGroupID

`func (o *O11yO11yErrorWithSpan) GetGroupID() string`

GetGroupID returns the GroupID field if non-nil, zero value otherwise.

### GetGroupIDOk

`func (o *O11yO11yErrorWithSpan) GetGroupIDOk() (*string, bool)`

GetGroupIDOk returns a tuple with the GroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupID

`func (o *O11yO11yErrorWithSpan) SetGroupID(v string)`

SetGroupID sets GroupID field to given value.

### HasGroupID

`func (o *O11yO11yErrorWithSpan) HasGroupID() bool`

HasGroupID returns a boolean if a field has been set.

### GetServiceName

`func (o *O11yO11yErrorWithSpan) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *O11yO11yErrorWithSpan) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *O11yO11yErrorWithSpan) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *O11yO11yErrorWithSpan) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSpanID

`func (o *O11yO11yErrorWithSpan) GetSpanID() string`

GetSpanID returns the SpanID field if non-nil, zero value otherwise.

### GetSpanIDOk

`func (o *O11yO11yErrorWithSpan) GetSpanIDOk() (*string, bool)`

GetSpanIDOk returns a tuple with the SpanID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanID

`func (o *O11yO11yErrorWithSpan) SetSpanID(v string)`

SetSpanID sets SpanID field to given value.

### HasSpanID

`func (o *O11yO11yErrorWithSpan) HasSpanID() bool`

HasSpanID returns a boolean if a field has been set.

### GetTimestamp

`func (o *O11yO11yErrorWithSpan) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *O11yO11yErrorWithSpan) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *O11yO11yErrorWithSpan) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *O11yO11yErrorWithSpan) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTraceID

`func (o *O11yO11yErrorWithSpan) GetTraceID() string`

GetTraceID returns the TraceID field if non-nil, zero value otherwise.

### GetTraceIDOk

`func (o *O11yO11yErrorWithSpan) GetTraceIDOk() (*string, bool)`

GetTraceIDOk returns a tuple with the TraceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceID

`func (o *O11yO11yErrorWithSpan) SetTraceID(v string)`

SetTraceID sets TraceID field to given value.

### HasTraceID

`func (o *O11yO11yErrorWithSpan) HasTraceID() bool`

HasTraceID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


