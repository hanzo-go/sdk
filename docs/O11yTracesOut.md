# O11yTracesOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count is how many traces this page carries. | [optional] 
**Limit** | Pointer to **int32** | Limit is the page cap actually applied, after clamping. | [optional] 
**SinceSec** | Pointer to **int32** | SinceSec is the window actually read, in seconds, after clamping. | [optional] 
**Traces** | Pointer to [**[]O11yTraceRow**](O11yTraceRow.md) | Traces are the caller org&#39;s traces, most recently active first. | [optional] 

## Methods

### NewO11yTracesOut

`func NewO11yTracesOut() *O11yTracesOut`

NewO11yTracesOut instantiates a new O11yTracesOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yTracesOutWithDefaults

`func NewO11yTracesOutWithDefaults() *O11yTracesOut`

NewO11yTracesOutWithDefaults instantiates a new O11yTracesOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *O11yTracesOut) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *O11yTracesOut) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *O11yTracesOut) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *O11yTracesOut) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLimit

`func (o *O11yTracesOut) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *O11yTracesOut) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *O11yTracesOut) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *O11yTracesOut) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetSinceSec

`func (o *O11yTracesOut) GetSinceSec() int32`

GetSinceSec returns the SinceSec field if non-nil, zero value otherwise.

### GetSinceSecOk

`func (o *O11yTracesOut) GetSinceSecOk() (*int32, bool)`

GetSinceSecOk returns a tuple with the SinceSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinceSec

`func (o *O11yTracesOut) SetSinceSec(v int32)`

SetSinceSec sets SinceSec field to given value.

### HasSinceSec

`func (o *O11yTracesOut) HasSinceSec() bool`

HasSinceSec returns a boolean if a field has been set.

### GetTraces

`func (o *O11yTracesOut) GetTraces() []O11yTraceRow`

GetTraces returns the Traces field if non-nil, zero value otherwise.

### GetTracesOk

`func (o *O11yTracesOut) GetTracesOk() (*[]O11yTraceRow, bool)`

GetTracesOk returns a tuple with the Traces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraces

`func (o *O11yTracesOut) SetTraces(v []O11yTraceRow)`

SetTraces sets Traces field to given value.

### HasTraces

`func (o *O11yTracesOut) HasTraces() bool`

HasTraces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


