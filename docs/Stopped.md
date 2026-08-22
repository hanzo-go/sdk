# Stopped

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stopped** | Pointer to **int32** | Stopped counts the commands that were still running and were interrupted. Zero says the sandbox was idle, not that the stop failed — see above. | [optional] 

## Methods

### NewStopped

`func NewStopped() *Stopped`

NewStopped instantiates a new Stopped object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoppedWithDefaults

`func NewStoppedWithDefaults() *Stopped`

NewStoppedWithDefaults instantiates a new Stopped object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStopped

`func (o *Stopped) GetStopped() int32`

GetStopped returns the Stopped field if non-nil, zero value otherwise.

### GetStoppedOk

`func (o *Stopped) GetStoppedOk() (*int32, bool)`

GetStoppedOk returns a tuple with the Stopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopped

`func (o *Stopped) SetStopped(v int32)`

SetStopped sets Stopped field to given value.

### HasStopped

`func (o *Stopped) HasStopped() bool`

HasStopped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


