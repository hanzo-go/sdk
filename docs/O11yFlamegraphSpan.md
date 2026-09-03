# O11yFlamegraphSpan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**DurationNano** | Pointer to **int32** |  | [optional] 
**Event** | Pointer to [**[]O11yEvent**](O11yEvent.md) |  | [optional] 
**HasError** | Pointer to **bool** |  | [optional] 
**Level** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentSpanId** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **map[string]string** |  | [optional] 
**SpanId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **int32** |  | [optional] 

## Methods

### NewO11yFlamegraphSpan

`func NewO11yFlamegraphSpan() *O11yFlamegraphSpan`

NewO11yFlamegraphSpan instantiates a new O11yFlamegraphSpan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yFlamegraphSpanWithDefaults

`func NewO11yFlamegraphSpanWithDefaults() *O11yFlamegraphSpan`

NewO11yFlamegraphSpanWithDefaults instantiates a new O11yFlamegraphSpan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *O11yFlamegraphSpan) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *O11yFlamegraphSpan) GetAttributesOk() (*map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *O11yFlamegraphSpan) SetAttributes(v map[string]map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *O11yFlamegraphSpan) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDurationNano

`func (o *O11yFlamegraphSpan) GetDurationNano() int32`

GetDurationNano returns the DurationNano field if non-nil, zero value otherwise.

### GetDurationNanoOk

`func (o *O11yFlamegraphSpan) GetDurationNanoOk() (*int32, bool)`

GetDurationNanoOk returns a tuple with the DurationNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationNano

`func (o *O11yFlamegraphSpan) SetDurationNano(v int32)`

SetDurationNano sets DurationNano field to given value.

### HasDurationNano

`func (o *O11yFlamegraphSpan) HasDurationNano() bool`

HasDurationNano returns a boolean if a field has been set.

### GetEvent

`func (o *O11yFlamegraphSpan) GetEvent() []O11yEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *O11yFlamegraphSpan) GetEventOk() (*[]O11yEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *O11yFlamegraphSpan) SetEvent(v []O11yEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *O11yFlamegraphSpan) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetHasError

`func (o *O11yFlamegraphSpan) GetHasError() bool`

GetHasError returns the HasError field if non-nil, zero value otherwise.

### GetHasErrorOk

`func (o *O11yFlamegraphSpan) GetHasErrorOk() (*bool, bool)`

GetHasErrorOk returns a tuple with the HasError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasError

`func (o *O11yFlamegraphSpan) SetHasError(v bool)`

SetHasError sets HasError field to given value.

### HasHasError

`func (o *O11yFlamegraphSpan) HasHasError() bool`

HasHasError returns a boolean if a field has been set.

### GetLevel

`func (o *O11yFlamegraphSpan) GetLevel() int64`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *O11yFlamegraphSpan) GetLevelOk() (*int64, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *O11yFlamegraphSpan) SetLevel(v int64)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *O11yFlamegraphSpan) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetName

`func (o *O11yFlamegraphSpan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yFlamegraphSpan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yFlamegraphSpan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yFlamegraphSpan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentSpanId

`func (o *O11yFlamegraphSpan) GetParentSpanId() string`

GetParentSpanId returns the ParentSpanId field if non-nil, zero value otherwise.

### GetParentSpanIdOk

`func (o *O11yFlamegraphSpan) GetParentSpanIdOk() (*string, bool)`

GetParentSpanIdOk returns a tuple with the ParentSpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpanId

`func (o *O11yFlamegraphSpan) SetParentSpanId(v string)`

SetParentSpanId sets ParentSpanId field to given value.

### HasParentSpanId

`func (o *O11yFlamegraphSpan) HasParentSpanId() bool`

HasParentSpanId returns a boolean if a field has been set.

### GetResource

`func (o *O11yFlamegraphSpan) GetResource() map[string]string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *O11yFlamegraphSpan) GetResourceOk() (*map[string]string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *O11yFlamegraphSpan) SetResource(v map[string]string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *O11yFlamegraphSpan) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSpanId

`func (o *O11yFlamegraphSpan) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *O11yFlamegraphSpan) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *O11yFlamegraphSpan) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.

### HasSpanId

`func (o *O11yFlamegraphSpan) HasSpanId() bool`

HasSpanId returns a boolean if a field has been set.

### GetTimestamp

`func (o *O11yFlamegraphSpan) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *O11yFlamegraphSpan) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *O11yFlamegraphSpan) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *O11yFlamegraphSpan) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


