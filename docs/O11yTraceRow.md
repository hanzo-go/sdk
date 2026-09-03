# O11yTraceRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationMs** | Pointer to **float64** | DurationMs is End minus Start in milliseconds: the trace&#39;s wall clock, not the sum of its spans, which double-counts everything concurrent. | [optional] 
**End** | Pointer to **string** | End is the latest span end, RFC3339 with nanoseconds, in UTC. | [optional] 
**NumSpans** | Pointer to **int64** | NumSpans is how many spans the trace carries. | [optional] 
**Start** | Pointer to **string** | Start is the earliest span start, RFC3339 with nanoseconds, in UTC. | [optional] 
**TraceId** | Pointer to **string** | TraceID is the trace&#39;s id — the {traceId} of the detail read. | [optional] 

## Methods

### NewO11yTraceRow

`func NewO11yTraceRow() *O11yTraceRow`

NewO11yTraceRow instantiates a new O11yTraceRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yTraceRowWithDefaults

`func NewO11yTraceRowWithDefaults() *O11yTraceRow`

NewO11yTraceRowWithDefaults instantiates a new O11yTraceRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationMs

`func (o *O11yTraceRow) GetDurationMs() float64`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *O11yTraceRow) GetDurationMsOk() (*float64, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *O11yTraceRow) SetDurationMs(v float64)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *O11yTraceRow) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetEnd

`func (o *O11yTraceRow) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yTraceRow) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yTraceRow) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *O11yTraceRow) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetNumSpans

`func (o *O11yTraceRow) GetNumSpans() int64`

GetNumSpans returns the NumSpans field if non-nil, zero value otherwise.

### GetNumSpansOk

`func (o *O11yTraceRow) GetNumSpansOk() (*int64, bool)`

GetNumSpansOk returns a tuple with the NumSpans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSpans

`func (o *O11yTraceRow) SetNumSpans(v int64)`

SetNumSpans sets NumSpans field to given value.

### HasNumSpans

`func (o *O11yTraceRow) HasNumSpans() bool`

HasNumSpans returns a boolean if a field has been set.

### GetStart

`func (o *O11yTraceRow) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yTraceRow) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yTraceRow) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *O11yTraceRow) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetTraceId

`func (o *O11yTraceRow) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *O11yTraceRow) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *O11yTraceRow) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *O11yTraceRow) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


