# O11ySpan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TraceID** | Pointer to **string** |  | [optional] 
**SpanID** | Pointer to **string** |  | [optional] 
**ParentSpanID** | Pointer to **string** |  | [optional] 
**OperationName** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**StartTimeUnixNano** | Pointer to **string** |  | [optional] 
**DurationNano** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]string** |  | [optional] 
**Events** | Pointer to [**[]O11ySpanEventsInner**](O11ySpanEventsInner.md) |  | [optional] 

## Methods

### NewO11ySpan

`func NewO11ySpan() *O11ySpan`

NewO11ySpan instantiates a new O11ySpan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11ySpanWithDefaults

`func NewO11ySpanWithDefaults() *O11ySpan`

NewO11ySpanWithDefaults instantiates a new O11ySpan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceID

`func (o *O11ySpan) GetTraceID() string`

GetTraceID returns the TraceID field if non-nil, zero value otherwise.

### GetTraceIDOk

`func (o *O11ySpan) GetTraceIDOk() (*string, bool)`

GetTraceIDOk returns a tuple with the TraceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceID

`func (o *O11ySpan) SetTraceID(v string)`

SetTraceID sets TraceID field to given value.

### HasTraceID

`func (o *O11ySpan) HasTraceID() bool`

HasTraceID returns a boolean if a field has been set.

### GetSpanID

`func (o *O11ySpan) GetSpanID() string`

GetSpanID returns the SpanID field if non-nil, zero value otherwise.

### GetSpanIDOk

`func (o *O11ySpan) GetSpanIDOk() (*string, bool)`

GetSpanIDOk returns a tuple with the SpanID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanID

`func (o *O11ySpan) SetSpanID(v string)`

SetSpanID sets SpanID field to given value.

### HasSpanID

`func (o *O11ySpan) HasSpanID() bool`

HasSpanID returns a boolean if a field has been set.

### GetParentSpanID

`func (o *O11ySpan) GetParentSpanID() string`

GetParentSpanID returns the ParentSpanID field if non-nil, zero value otherwise.

### GetParentSpanIDOk

`func (o *O11ySpan) GetParentSpanIDOk() (*string, bool)`

GetParentSpanIDOk returns a tuple with the ParentSpanID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpanID

`func (o *O11ySpan) SetParentSpanID(v string)`

SetParentSpanID sets ParentSpanID field to given value.

### HasParentSpanID

`func (o *O11ySpan) HasParentSpanID() bool`

HasParentSpanID returns a boolean if a field has been set.

### GetOperationName

`func (o *O11ySpan) GetOperationName() string`

GetOperationName returns the OperationName field if non-nil, zero value otherwise.

### GetOperationNameOk

`func (o *O11ySpan) GetOperationNameOk() (*string, bool)`

GetOperationNameOk returns a tuple with the OperationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationName

`func (o *O11ySpan) SetOperationName(v string)`

SetOperationName sets OperationName field to given value.

### HasOperationName

`func (o *O11ySpan) HasOperationName() bool`

HasOperationName returns a boolean if a field has been set.

### GetServiceName

`func (o *O11ySpan) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *O11ySpan) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *O11ySpan) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *O11ySpan) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetStartTimeUnixNano

`func (o *O11ySpan) GetStartTimeUnixNano() string`

GetStartTimeUnixNano returns the StartTimeUnixNano field if non-nil, zero value otherwise.

### GetStartTimeUnixNanoOk

`func (o *O11ySpan) GetStartTimeUnixNanoOk() (*string, bool)`

GetStartTimeUnixNanoOk returns a tuple with the StartTimeUnixNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUnixNano

`func (o *O11ySpan) SetStartTimeUnixNano(v string)`

SetStartTimeUnixNano sets StartTimeUnixNano field to given value.

### HasStartTimeUnixNano

`func (o *O11ySpan) HasStartTimeUnixNano() bool`

HasStartTimeUnixNano returns a boolean if a field has been set.

### GetDurationNano

`func (o *O11ySpan) GetDurationNano() string`

GetDurationNano returns the DurationNano field if non-nil, zero value otherwise.

### GetDurationNanoOk

`func (o *O11ySpan) GetDurationNanoOk() (*string, bool)`

GetDurationNanoOk returns a tuple with the DurationNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationNano

`func (o *O11ySpan) SetDurationNano(v string)`

SetDurationNano sets DurationNano field to given value.

### HasDurationNano

`func (o *O11ySpan) HasDurationNano() bool`

HasDurationNano returns a boolean if a field has been set.

### GetStatus

`func (o *O11ySpan) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *O11ySpan) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *O11ySpan) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *O11ySpan) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAttributes

`func (o *O11ySpan) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *O11ySpan) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *O11ySpan) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *O11ySpan) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetEvents

`func (o *O11ySpan) GetEvents() []O11ySpanEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *O11ySpan) GetEventsOk() (*[]O11ySpanEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *O11ySpan) SetEvents(v []O11ySpanEventsInner)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *O11ySpan) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


