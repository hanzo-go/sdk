# FunctionsInvocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **string** | RFC3339 UTC | [optional] 
**DurationMs** | Pointer to **int64** |  | [optional] 

## Methods

### NewFunctionsInvocation

`func NewFunctionsInvocation() *FunctionsInvocation`

NewFunctionsInvocation instantiates a new FunctionsInvocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionsInvocationWithDefaults

`func NewFunctionsInvocationWithDefaults() *FunctionsInvocation`

NewFunctionsInvocationWithDefaults instantiates a new FunctionsInvocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FunctionsInvocation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FunctionsInvocation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FunctionsInvocation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FunctionsInvocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatusCode

`func (o *FunctionsInvocation) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *FunctionsInvocation) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *FunctionsInvocation) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *FunctionsInvocation) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetStatus

`func (o *FunctionsInvocation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FunctionsInvocation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FunctionsInvocation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FunctionsInvocation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMethod

`func (o *FunctionsInvocation) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *FunctionsInvocation) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *FunctionsInvocation) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *FunctionsInvocation) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetTime

`func (o *FunctionsInvocation) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *FunctionsInvocation) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *FunctionsInvocation) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *FunctionsInvocation) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetDurationMs

`func (o *FunctionsInvocation) GetDurationMs() int64`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *FunctionsInvocation) GetDurationMsOk() (*int64, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *FunctionsInvocation) SetDurationMs(v int64)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *FunctionsInvocation) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


