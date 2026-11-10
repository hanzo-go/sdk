# BenchmarkSuite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Benchmarks** | **[]string** | Benchmarks are the catalog ids to ask about. At least one is required, and every id must be in the catalog. | 
**Model** | Pointer to **string** | Model is the model to ask about, by the id its attempts are stored under. Required. | [optional] 

## Methods

### NewBenchmarkSuite

`func NewBenchmarkSuite(benchmarks []string, ) *BenchmarkSuite`

NewBenchmarkSuite instantiates a new BenchmarkSuite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBenchmarkSuiteWithDefaults

`func NewBenchmarkSuiteWithDefaults() *BenchmarkSuite`

NewBenchmarkSuiteWithDefaults instantiates a new BenchmarkSuite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBenchmarks

`func (o *BenchmarkSuite) GetBenchmarks() []string`

GetBenchmarks returns the Benchmarks field if non-nil, zero value otherwise.

### GetBenchmarksOk

`func (o *BenchmarkSuite) GetBenchmarksOk() (*[]string, bool)`

GetBenchmarksOk returns a tuple with the Benchmarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmarks

`func (o *BenchmarkSuite) SetBenchmarks(v []string)`

SetBenchmarks sets Benchmarks field to given value.


### GetModel

`func (o *BenchmarkSuite) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BenchmarkSuite) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BenchmarkSuite) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *BenchmarkSuite) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


