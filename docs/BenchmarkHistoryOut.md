# BenchmarkHistoryOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Benchmark** | Pointer to **string** | Benchmark is the catalog id these histories are about. | [optional] 
**Data** | Pointer to [**[]BenchmarkModelHistory**](BenchmarkModelHistory.md) | Data is one entry per model, ordered by model name. | [optional] 
**Total** | Pointer to **int64** | Total is how many models Data holds. | [optional] 

## Methods

### NewBenchmarkHistoryOut

`func NewBenchmarkHistoryOut() *BenchmarkHistoryOut`

NewBenchmarkHistoryOut instantiates a new BenchmarkHistoryOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarkHistoryOutWithDefaults

`func NewBenchmarkHistoryOutWithDefaults() *BenchmarkHistoryOut`

NewBenchmarkHistoryOutWithDefaults instantiates a new BenchmarkHistoryOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBenchmark

`func (o *BenchmarkHistoryOut) GetBenchmark() string`

GetBenchmark returns the Benchmark field if non-nil, zero value otherwise.

### GetBenchmarkOk

`func (o *BenchmarkHistoryOut) GetBenchmarkOk() (*string, bool)`

GetBenchmarkOk returns a tuple with the Benchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmark

`func (o *BenchmarkHistoryOut) SetBenchmark(v string)`

SetBenchmark sets Benchmark field to given value.

### HasBenchmark

`func (o *BenchmarkHistoryOut) HasBenchmark() bool`

HasBenchmark returns a boolean if a field has been set.

### GetData

`func (o *BenchmarkHistoryOut) GetData() []BenchmarkModelHistory`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BenchmarkHistoryOut) GetDataOk() (*[]BenchmarkModelHistory, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BenchmarkHistoryOut) SetData(v []BenchmarkModelHistory)`

SetData sets Data field to given value.

### HasData

`func (o *BenchmarkHistoryOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *BenchmarkHistoryOut) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BenchmarkHistoryOut) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BenchmarkHistoryOut) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BenchmarkHistoryOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


