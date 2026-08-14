# O11yO11yTrace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count is how many captured errors carried it. | [optional] 
**FirstSeen** | Pointer to **time.Time** | FirstSeen is when the earliest of them was recorded. | [optional] 
**LastSeen** | Pointer to **time.Time** | LastSeen is when the latest was. | [optional] 
**Message** | Pointer to **string** | Message is the latest error message seen on the trace. | [optional] 
**TraceId** | Pointer to **string** | TraceID is the trace id. | [optional] 

## Methods

### NewO11yO11yTrace

`func NewO11yO11yTrace() *O11yO11yTrace`

NewO11yO11yTrace instantiates a new O11yO11yTrace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yTraceWithDefaults

`func NewO11yO11yTraceWithDefaults() *O11yO11yTrace`

NewO11yO11yTraceWithDefaults instantiates a new O11yO11yTrace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *O11yO11yTrace) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *O11yO11yTrace) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *O11yO11yTrace) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *O11yO11yTrace) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetFirstSeen

`func (o *O11yO11yTrace) GetFirstSeen() time.Time`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *O11yO11yTrace) GetFirstSeenOk() (*time.Time, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *O11yO11yTrace) SetFirstSeen(v time.Time)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *O11yO11yTrace) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastSeen

`func (o *O11yO11yTrace) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *O11yO11yTrace) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *O11yO11yTrace) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *O11yO11yTrace) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMessage

`func (o *O11yO11yTrace) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *O11yO11yTrace) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *O11yO11yTrace) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *O11yO11yTrace) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTraceId

`func (o *O11yO11yTrace) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *O11yO11yTrace) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *O11yO11yTrace) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *O11yO11yTrace) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


