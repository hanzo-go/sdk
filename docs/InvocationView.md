# InvocationView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationMs** | Pointer to **int64** | how long it took | [optional] 
**Id** | Pointer to **string** | the invocation&#39;s handle | [optional] 
**Method** | Pointer to **string** | the HTTP method that triggered it | [optional] 
**Status** | Pointer to **string** | how the run ended: ok, error or timeout | [optional] 
**StatusCode** | Pointer to **int64** | Code is the status the function&#39;s OWN code answered with, which is not the status of the reply — a program can answer 500 through a healthy sandbox. | [optional] 
**Time** | Pointer to **string** | when it ran, RFC3339 | [optional] 

## Methods

### NewInvocationView

`func NewInvocationView() *InvocationView`

NewInvocationView instantiates a new InvocationView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvocationViewWithDefaults

`func NewInvocationViewWithDefaults() *InvocationView`

NewInvocationViewWithDefaults instantiates a new InvocationView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationMs

`func (o *InvocationView) GetDurationMs() int64`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *InvocationView) GetDurationMsOk() (*int64, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *InvocationView) SetDurationMs(v int64)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *InvocationView) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetId

`func (o *InvocationView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvocationView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvocationView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvocationView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethod

`func (o *InvocationView) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *InvocationView) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *InvocationView) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *InvocationView) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetStatus

`func (o *InvocationView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvocationView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvocationView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InvocationView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *InvocationView) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *InvocationView) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *InvocationView) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *InvocationView) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetTime

`func (o *InvocationView) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *InvocationView) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *InvocationView) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *InvocationView) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


