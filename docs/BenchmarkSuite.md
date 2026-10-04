# BenchmarkSuite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | Pointer to **int64** | Attempts is how many times to try each item; the harness&#39;s default applies when it is omitted. | [optional] 
**Benchmarks** | **[]string** | Benchmarks are the catalog ids to run. At least one is required, and every id must be in the catalog. | 
**Endpoint** | Pointer to **string** | Endpoint is your own chat-completions URL, for benchmarking a model this arena does not host. Either this or model is required. | [optional] 
**Model** | Pointer to **string** | Model is the catalog model id to run. Either this or endpoint is required. | [optional] 

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

### GetAttempts

`func (o *BenchmarkSuite) GetAttempts() int64`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *BenchmarkSuite) GetAttemptsOk() (*int64, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *BenchmarkSuite) SetAttempts(v int64)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *BenchmarkSuite) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

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


### GetEndpoint

`func (o *BenchmarkSuite) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *BenchmarkSuite) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *BenchmarkSuite) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *BenchmarkSuite) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

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


