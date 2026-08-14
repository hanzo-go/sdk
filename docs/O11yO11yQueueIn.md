# O11yO11yQueueIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int32** | End is the window&#39;s end, epoch nanoseconds. | [optional] 
**EvalTime** | Pointer to **int32** | EvalTime bounds the span-evaluation scan, nanoseconds; only the span/evaluation view reads it. | [optional] 
**Start** | Pointer to **int32** | Start is the window&#39;s start, epoch nanoseconds. | [optional] 
**Variables** | Pointer to **map[string]string** | Variables name what the view drills into — topic, partition, service, consumer_group — keyed by the name the view expects. | [optional] 

## Methods

### NewO11yO11yQueueIn

`func NewO11yO11yQueueIn() *O11yO11yQueueIn`

NewO11yO11yQueueIn instantiates a new O11yO11yQueueIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yQueueInWithDefaults

`func NewO11yO11yQueueInWithDefaults() *O11yO11yQueueIn`

NewO11yO11yQueueInWithDefaults instantiates a new O11yO11yQueueIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *O11yO11yQueueIn) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *O11yO11yQueueIn) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *O11yO11yQueueIn) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *O11yO11yQueueIn) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetEvalTime

`func (o *O11yO11yQueueIn) GetEvalTime() int32`

GetEvalTime returns the EvalTime field if non-nil, zero value otherwise.

### GetEvalTimeOk

`func (o *O11yO11yQueueIn) GetEvalTimeOk() (*int32, bool)`

GetEvalTimeOk returns a tuple with the EvalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvalTime

`func (o *O11yO11yQueueIn) SetEvalTime(v int32)`

SetEvalTime sets EvalTime field to given value.

### HasEvalTime

`func (o *O11yO11yQueueIn) HasEvalTime() bool`

HasEvalTime returns a boolean if a field has been set.

### GetStart

`func (o *O11yO11yQueueIn) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *O11yO11yQueueIn) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *O11yO11yQueueIn) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *O11yO11yQueueIn) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetVariables

`func (o *O11yO11yQueueIn) GetVariables() map[string]string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *O11yO11yQueueIn) GetVariablesOk() (*map[string]string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *O11yO11yQueueIn) SetVariables(v map[string]string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *O11yO11yQueueIn) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


