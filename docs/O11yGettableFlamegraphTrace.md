# O11yGettableFlamegraphTrace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTimestampMillis** | Pointer to **int64** |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**Spans** | Pointer to [**[][]O11yFlamegraphSpan**]([]O11yFlamegraphSpan.md) |  | [optional] 
**StartTimestampMillis** | Pointer to **int64** |  | [optional] 

## Methods

### NewO11yGettableFlamegraphTrace

`func NewO11yGettableFlamegraphTrace() *O11yGettableFlamegraphTrace`

NewO11yGettableFlamegraphTrace instantiates a new O11yGettableFlamegraphTrace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yGettableFlamegraphTraceWithDefaults

`func NewO11yGettableFlamegraphTraceWithDefaults() *O11yGettableFlamegraphTrace`

NewO11yGettableFlamegraphTraceWithDefaults instantiates a new O11yGettableFlamegraphTrace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTimestampMillis

`func (o *O11yGettableFlamegraphTrace) GetEndTimestampMillis() int64`

GetEndTimestampMillis returns the EndTimestampMillis field if non-nil, zero value otherwise.

### GetEndTimestampMillisOk

`func (o *O11yGettableFlamegraphTrace) GetEndTimestampMillisOk() (*int64, bool)`

GetEndTimestampMillisOk returns a tuple with the EndTimestampMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestampMillis

`func (o *O11yGettableFlamegraphTrace) SetEndTimestampMillis(v int64)`

SetEndTimestampMillis sets EndTimestampMillis field to given value.

### HasEndTimestampMillis

`func (o *O11yGettableFlamegraphTrace) HasEndTimestampMillis() bool`

HasEndTimestampMillis returns a boolean if a field has been set.

### GetHasMore

`func (o *O11yGettableFlamegraphTrace) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *O11yGettableFlamegraphTrace) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *O11yGettableFlamegraphTrace) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *O11yGettableFlamegraphTrace) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetSpans

`func (o *O11yGettableFlamegraphTrace) GetSpans() [][]O11yFlamegraphSpan`

GetSpans returns the Spans field if non-nil, zero value otherwise.

### GetSpansOk

`func (o *O11yGettableFlamegraphTrace) GetSpansOk() (*[][]O11yFlamegraphSpan, bool)`

GetSpansOk returns a tuple with the Spans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpans

`func (o *O11yGettableFlamegraphTrace) SetSpans(v [][]O11yFlamegraphSpan)`

SetSpans sets Spans field to given value.

### HasSpans

`func (o *O11yGettableFlamegraphTrace) HasSpans() bool`

HasSpans returns a boolean if a field has been set.

### GetStartTimestampMillis

`func (o *O11yGettableFlamegraphTrace) GetStartTimestampMillis() int64`

GetStartTimestampMillis returns the StartTimestampMillis field if non-nil, zero value otherwise.

### GetStartTimestampMillisOk

`func (o *O11yGettableFlamegraphTrace) GetStartTimestampMillisOk() (*int64, bool)`

GetStartTimestampMillisOk returns a tuple with the StartTimestampMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestampMillis

`func (o *O11yGettableFlamegraphTrace) SetStartTimestampMillis(v int64)`

SetStartTimestampMillis sets StartTimestampMillis field to given value.

### HasStartTimestampMillis

`func (o *O11yGettableFlamegraphTrace) HasStartTimestampMillis() bool`

HasStartTimestampMillis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


