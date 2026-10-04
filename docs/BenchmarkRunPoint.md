# BenchmarkRunPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **time.Time** | At is when the run was recorded. | [optional] 
**Delta** | Pointer to **float64** | Delta is the change in score from the previous run for this model, absent on the first. It is the number the whole surface exists to make visible. | [optional] 
**N** | Pointer to **int64** | N is how many items the run covered. Two runs are only comparable at the same n, which is why it travels with every point rather than being assumed. | [optional] 
**Run** | Pointer to **string** | Run is the measurement id these attempts were recorded under. | [optional] 
**Score** | Pointer to **float64** | Score is accuracy over the items this run covered, as a percentage. | [optional] 

## Methods

### NewBenchmarkRunPoint

`func NewBenchmarkRunPoint() *BenchmarkRunPoint`

NewBenchmarkRunPoint instantiates a new BenchmarkRunPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarkRunPointWithDefaults

`func NewBenchmarkRunPointWithDefaults() *BenchmarkRunPoint`

NewBenchmarkRunPointWithDefaults instantiates a new BenchmarkRunPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *BenchmarkRunPoint) GetAt() time.Time`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *BenchmarkRunPoint) GetAtOk() (*time.Time, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *BenchmarkRunPoint) SetAt(v time.Time)`

SetAt sets At field to given value.

### HasAt

`func (o *BenchmarkRunPoint) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetDelta

`func (o *BenchmarkRunPoint) GetDelta() float64`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *BenchmarkRunPoint) GetDeltaOk() (*float64, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *BenchmarkRunPoint) SetDelta(v float64)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *BenchmarkRunPoint) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### GetN

`func (o *BenchmarkRunPoint) GetN() int64`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *BenchmarkRunPoint) GetNOk() (*int64, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *BenchmarkRunPoint) SetN(v int64)`

SetN sets N field to given value.

### HasN

`func (o *BenchmarkRunPoint) HasN() bool`

HasN returns a boolean if a field has been set.

### GetRun

`func (o *BenchmarkRunPoint) GetRun() string`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *BenchmarkRunPoint) GetRunOk() (*string, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *BenchmarkRunPoint) SetRun(v string)`

SetRun sets Run field to given value.

### HasRun

`func (o *BenchmarkRunPoint) HasRun() bool`

HasRun returns a boolean if a field has been set.

### GetScore

`func (o *BenchmarkRunPoint) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *BenchmarkRunPoint) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *BenchmarkRunPoint) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *BenchmarkRunPoint) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


