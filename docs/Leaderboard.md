# Leaderboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Benchmark** | Pointer to **string** | Benchmark is the catalog id these rows are about. | [optional] 
**Rows** | Pointer to [**[]LeaderRow**](LeaderRow.md) | Rows is one per model, ordered by measured accuracy descending. | [optional] 

## Methods

### NewLeaderboard

`func NewLeaderboard() *Leaderboard`

NewLeaderboard instantiates a new Leaderboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardWithDefaults

`func NewLeaderboardWithDefaults() *Leaderboard`

NewLeaderboardWithDefaults instantiates a new Leaderboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBenchmark

`func (o *Leaderboard) GetBenchmark() string`

GetBenchmark returns the Benchmark field if non-nil, zero value otherwise.

### GetBenchmarkOk

`func (o *Leaderboard) GetBenchmarkOk() (*string, bool)`

GetBenchmarkOk returns a tuple with the Benchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmark

`func (o *Leaderboard) SetBenchmark(v string)`

SetBenchmark sets Benchmark field to given value.

### HasBenchmark

`func (o *Leaderboard) HasBenchmark() bool`

HasBenchmark returns a boolean if a field has been set.

### GetRows

`func (o *Leaderboard) GetRows() []LeaderRow`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *Leaderboard) GetRowsOk() (*[]LeaderRow, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *Leaderboard) SetRows(v []LeaderRow)`

SetRows sets Rows field to given value.

### HasRows

`func (o *Leaderboard) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


