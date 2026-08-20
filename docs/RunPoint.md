# RunPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **time.Time** | At is when the run was recorded. | [optional] 
**Delta** | Pointer to **float32** | Delta is the change in score from the previous run for this model, absent on the first. It is the number the whole surface exists to make visible. | [optional] 
**N** | Pointer to **int32** | N is how many items the run covered. Two runs are only comparable at the same n, which is why it travels with every point rather than being assumed. | [optional] 
**Run** | Pointer to **string** | Run is the measurement id these attempts were recorded under. | [optional] 
**Score** | Pointer to **float32** | Score is accuracy over the items this run covered, as a percentage. | [optional] 

## Methods

### NewRunPoint

`func NewRunPoint() *RunPoint`

NewRunPoint instantiates a new RunPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunPointWithDefaults

`func NewRunPointWithDefaults() *RunPoint`

NewRunPointWithDefaults instantiates a new RunPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *RunPoint) GetAt() time.Time`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *RunPoint) GetAtOk() (*time.Time, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *RunPoint) SetAt(v time.Time)`

SetAt sets At field to given value.

### HasAt

`func (o *RunPoint) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetDelta

`func (o *RunPoint) GetDelta() float32`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *RunPoint) GetDeltaOk() (*float32, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *RunPoint) SetDelta(v float32)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *RunPoint) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### GetN

`func (o *RunPoint) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *RunPoint) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *RunPoint) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *RunPoint) HasN() bool`

HasN returns a boolean if a field has been set.

### GetRun

`func (o *RunPoint) GetRun() string`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *RunPoint) GetRunOk() (*string, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *RunPoint) SetRun(v string)`

SetRun sets Run field to given value.

### HasRun

`func (o *RunPoint) HasRun() bool`

HasRun returns a boolean if a field has been set.

### GetScore

`func (o *RunPoint) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *RunPoint) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *RunPoint) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *RunPoint) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


