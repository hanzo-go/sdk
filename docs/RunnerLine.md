# RunnerLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Content is the line itself, without its terminator. | [optional] 
**Time** | Pointer to **int64** | Time is when the line was written, in unix nanoseconds. | [optional] 

## Methods

### NewRunnerLine

`func NewRunnerLine() *RunnerLine`

NewRunnerLine instantiates a new RunnerLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerLineWithDefaults

`func NewRunnerLineWithDefaults() *RunnerLine`

NewRunnerLineWithDefaults instantiates a new RunnerLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *RunnerLine) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *RunnerLine) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *RunnerLine) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *RunnerLine) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetTime

`func (o *RunnerLine) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *RunnerLine) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *RunnerLine) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *RunnerLine) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


