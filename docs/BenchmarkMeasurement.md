# BenchmarkMeasurement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Benchmark** | Pointer to **string** | Benchmark is the catalog id this answer is about. | [optional] 
**Harness** | Pointer to **string** | Harness is the harness that measures this benchmark for the arena, by repository. Empty when none does, and then a measurement arrives only as a published run&#39;s items imported through POST /v1/benchmark/attempts. | [optional] 
**Latest** | Pointer to [**BenchmarkRunPoint**](BenchmarkRunPoint.md) | Latest is the model&#39;s latest run on this benchmark — its run id, when it was recorded, its accuracy and how many items it covered — read from the attempts the store already holds. Null when the store holds none. | [optional] 
**Note** | Pointer to **string** | Note says in a sentence what the answer means: a held measurement served rather than spent again, or nothing queued and why. | [optional] 

## Methods

### NewBenchmarkMeasurement

`func NewBenchmarkMeasurement() *BenchmarkMeasurement`

NewBenchmarkMeasurement instantiates a new BenchmarkMeasurement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarkMeasurementWithDefaults

`func NewBenchmarkMeasurementWithDefaults() *BenchmarkMeasurement`

NewBenchmarkMeasurementWithDefaults instantiates a new BenchmarkMeasurement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBenchmark

`func (o *BenchmarkMeasurement) GetBenchmark() string`

GetBenchmark returns the Benchmark field if non-nil, zero value otherwise.

### GetBenchmarkOk

`func (o *BenchmarkMeasurement) GetBenchmarkOk() (*string, bool)`

GetBenchmarkOk returns a tuple with the Benchmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmark

`func (o *BenchmarkMeasurement) SetBenchmark(v string)`

SetBenchmark sets Benchmark field to given value.

### HasBenchmark

`func (o *BenchmarkMeasurement) HasBenchmark() bool`

HasBenchmark returns a boolean if a field has been set.

### GetHarness

`func (o *BenchmarkMeasurement) GetHarness() string`

GetHarness returns the Harness field if non-nil, zero value otherwise.

### GetHarnessOk

`func (o *BenchmarkMeasurement) GetHarnessOk() (*string, bool)`

GetHarnessOk returns a tuple with the Harness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHarness

`func (o *BenchmarkMeasurement) SetHarness(v string)`

SetHarness sets Harness field to given value.

### HasHarness

`func (o *BenchmarkMeasurement) HasHarness() bool`

HasHarness returns a boolean if a field has been set.

### GetLatest

`func (o *BenchmarkMeasurement) GetLatest() BenchmarkRunPoint`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *BenchmarkMeasurement) GetLatestOk() (*BenchmarkRunPoint, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *BenchmarkMeasurement) SetLatest(v BenchmarkRunPoint)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *BenchmarkMeasurement) HasLatest() bool`

HasLatest returns a boolean if a field has been set.

### GetNote

`func (o *BenchmarkMeasurement) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *BenchmarkMeasurement) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *BenchmarkMeasurement) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *BenchmarkMeasurement) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


