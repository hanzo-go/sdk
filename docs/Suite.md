# Suite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | Pointer to **int32** | Attempts is how many times to try each item; the harness&#39;s default applies when it is omitted. | [optional] 
**Benchmarks** | **[]string** | Benchmarks are the catalog ids to run. At least one is required, and every id must be in the catalog. | 
**Endpoint** | Pointer to **string** | Endpoint is your own chat-completions URL, for benchmarking a model this arena does not host. Either this or model is required. | [optional] 
**Model** | Pointer to **string** | Model is the catalog model id to run. Either this or endpoint is required. | [optional] 

## Methods

### NewSuite

`func NewSuite(benchmarks []string, ) *Suite`

NewSuite instantiates a new Suite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuiteWithDefaults

`func NewSuiteWithDefaults() *Suite`

NewSuiteWithDefaults instantiates a new Suite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *Suite) GetAttempts() int32`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *Suite) GetAttemptsOk() (*int32, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *Suite) SetAttempts(v int32)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *Suite) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetBenchmarks

`func (o *Suite) GetBenchmarks() []string`

GetBenchmarks returns the Benchmarks field if non-nil, zero value otherwise.

### GetBenchmarksOk

`func (o *Suite) GetBenchmarksOk() (*[]string, bool)`

GetBenchmarksOk returns a tuple with the Benchmarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmarks

`func (o *Suite) SetBenchmarks(v []string)`

SetBenchmarks sets Benchmarks field to given value.


### GetEndpoint

`func (o *Suite) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *Suite) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *Suite) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *Suite) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetModel

`func (o *Suite) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Suite) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Suite) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Suite) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


